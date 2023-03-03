package chainclient

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
)

var (
	ErrTimedOut       = errors.New("tx timed out")
	ErrQueueClosed    = errors.New("queue is closed")
	ErrEnqueueTimeout = errors.New("enqueue timeout")
	ErrReadOnly       = errors.New("client is in read-only mode")
)

type ChainClient interface {
	ClientContext() client.Context
	FromAddress() sdk.AccAddress
	CanSignTransactions() bool

	GetAccountNumberSequence() (accNum uint64, accSeq uint64)
	QueryClient() *grpc.ClientConn

	SimulateMsg(ctx context.Context, msgs ...sdk.Msg) (*txtypes.SimulateResponse, error)

	AsyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error)
	SyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error)
	QueueBroadcastMsg(msgs ...sdk.Msg) error

	Close()
}

type chainClient struct {
	clientContext client.Context
	conn          *grpc.ClientConn

	doneC chan bool
	msgC  chan sdk.Msg

	txFactory tx.Factory
	syncMux   *sync.RWMutex

	accNum uint64
	accSeq uint64

	closed  int64
	canSign bool

	txClient txtypes.ServiceClient
	logger   zerolog.Logger
}

// NewCosmosClient creates a new gRPC client that communicates with gRPC server at protoAddr.
// protoAddr must be in form "tcp://127.0.0.1:8080" or "unix:///tmp/test.sock", protocol is required.
func NewChainClient(
	clientContext client.Context,
	protoAddr string,
	options ...ClientOption,
) (ChainClient, error) {
	opts := DefaultOptions()
	for _, opt := range options {
		if err := opt(opts); err != nil {
			err = errors.Wrap(err, "error in client option")
			return nil, err
		}
	}

	txFactory := NewTxFactory(clientContext)
	if len(opts.GasPrices) > 0 {
		txFactory = txFactory.WithGasPrices(opts.GasPrices)
	}

	var (
		conn *grpc.ClientConn
		err  error
	)

	if opts.TLSCert != nil {
		conn, err = grpc.Dial(protoAddr, grpc.WithTransportCredentials(opts.TLSCert), grpc.WithContextDialer(DialerFunc))
	} else {
		conn, err = grpc.Dial(protoAddr, grpc.WithInsecure(), grpc.WithContextDialer(DialerFunc))
	}
	if err != nil {
		err = errors.Wrapf(err, "failed to connect to the gRPC: %s", protoAddr)
		return nil, err
	}

	if clientContext.Client != nil && !clientContext.Client.IsRunning() {
		err = clientContext.Client.Start()
		if err != nil {
			return nil, errors.Wrap(err, "failed to start client")
		}
	}

	cc := &chainClient{
		clientContext: clientContext,
		conn:          conn,

		txFactory: txFactory,
		canSign:   clientContext.Keyring != nil,
		syncMux:   new(sync.RWMutex),

		msgC:  make(chan sdk.Msg, msgCommitBatchSizeLimit),
		doneC: make(chan bool, 1),

		txClient: txtypes.NewServiceClient(conn),
		logger: log.With().Str(
			"module", "persistence-sdk",
		).Logger(),
	}

	if cc.canSign {
		var err error

		cc.accNum, cc.accSeq, err = cc.txFactory.AccountRetriever().GetAccountNumberSequence(clientContext, clientContext.GetFromAddress())
		if err != nil {
			err = errors.Wrap(err, "failed to get initial account num and seq")
			return nil, err
		}

		go cc.runBatchBroadcast()
		go cc.syncTimeoutHeight()
	}

	return cc, nil
}

func (c *chainClient) syncNonce() {
	num, seq, err := c.txFactory.AccountRetriever().GetAccountNumberSequence(c.clientContext, c.clientContext.GetFromAddress())
	if err != nil {
		c.logger.Warn().Err(err).Msg("failed to get account seq")
		return
	} else if num != c.accNum {
		c.logger.Panic().
			Uint64("expected", c.accNum).
			Uint64("actual", num).Msg("account number changed during nonce sync")
	}

	c.accSeq = seq
}

const (
	defaultTimeoutHeight             = 20
	defaultTimeoutHeightSyncInterval = 10 * time.Second
)

func (c *chainClient) syncTimeoutHeight() {
	for {
		ctx := context.Background()
		block, err := c.clientContext.Client.Block(ctx, nil)
		if err != nil {
			c.logger.Warn().Err(err).Msg("failed to get current block")
			return
		}

		c.syncMux.Lock()
		c.txFactory.WithTimeoutHeight(uint64(block.Block.Height) + defaultTimeoutHeight)
		c.syncMux.Unlock()

		time.Sleep(defaultTimeoutHeightSyncInterval)
	}
}

// prepareFactory ensures the account defined by ctx.GetFromAddress() exists and
// if the account number and/or the account sequence number are zero (not set),
// they will be queried for and set on the provided Factory. A new Factory with
// the updated fields will be returned.
func (c *chainClient) prepareFactory(clientCtx client.Context, txf tx.Factory) (tx.Factory, error) {
	from := clientCtx.GetFromAddress()

	if err := txf.AccountRetriever().EnsureExists(clientCtx, from); err != nil {
		return txf, err
	}

	initNum, initSeq := txf.AccountNumber(), txf.Sequence()
	if initNum == 0 || initSeq == 0 {
		num, seq, err := txf.AccountRetriever().GetAccountNumberSequence(clientCtx, from)
		if err != nil {
			return txf, err
		}

		if initNum == 0 {
			txf = txf.WithAccountNumber(num)
		}

		if initSeq == 0 {
			txf = txf.WithSequence(seq)
		}
	}

	return txf, nil
}

func (c *chainClient) GetAccountNumberSequence() (accNum uint64, accSeq uint64) {
	c.syncMux.RLock()
	defer c.syncMux.RUnlock()

	return c.accNum, c.accSeq
}

func (c *chainClient) QueryClient() *grpc.ClientConn {
	return c.conn
}

func (c *chainClient) ClientContext() client.Context {
	return c.clientContext
}

func (c *chainClient) CanSignTransactions() bool {
	return c.canSign
}

func (c *chainClient) FromAddress() sdk.AccAddress {
	if !c.canSign {
		return sdk.AccAddress{}
	}

	return c.clientContext.FromAddress
}

func (c *chainClient) Close() {
	if !c.canSign {
		return
	}

	if atomic.CompareAndSwapInt64(&c.closed, 0, 1) {
		close(c.msgC)
	}

	<-c.doneC

	if c.conn != nil {
		c.conn.Close()
	}
}

// SimulateMsg runs a simulation and returns gas estimate.
func (c *chainClient) SimulateMsg(
	ctx context.Context,
	msgs ...sdk.Msg,
) (*txtypes.SimulateResponse, error) {
	c.syncMux.RLock()
	txFactory := c.txFactory // a copy
	txFactory = txFactory.WithSequence(c.accSeq)
	txFactory = txFactory.WithAccountNumber(c.accNum)
	c.syncMux.RUnlock()

	txf, err := c.prepareFactory(c.clientContext, txFactory)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare factory for the client")
		return nil, err
	}

	simTxBytes, err := tx.BuildSimTx(txf, msgs...)
	if err != nil {
		err = errors.Wrap(err, "failed to build SimTx bytes")
		return nil, err
	}

	simRes, err := c.txClient.Simulate(ctx, &txtypes.SimulateRequest{
		TxBytes: simTxBytes,
	})
	if err != nil {
		err = errors.Wrap(err, "failed to run SimulateRequest")
		return nil, err
	}

	return simRes, nil
}

// SyncBroadcastMsg sends Tx to chain and waits until Tx is included in block.
func (c *chainClient) SyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error) {
	c.syncMux.Lock()
	defer c.syncMux.Unlock()

	c.txFactory = c.txFactory.WithSequence(c.accSeq)
	c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
	res, err := c.broadcastTx(context.TODO(), c.clientContext, c.txFactory, true, msgs...)

	if err != nil {
		if strings.Contains(err.Error(), "account sequence mismatch") {
			c.syncNonce()
			c.txFactory = c.txFactory.WithSequence(c.accSeq)
			c.txFactory = c.txFactory.WithAccountNumber(c.accNum)

			c.logger.Debug().Uint64("nonce", c.accSeq).Msg("retrying broadcastTx with nonce")

			res, err = c.broadcastTx(context.TODO(), c.clientContext, c.txFactory, true, msgs...)
		}
		if err != nil {
			resJSON, _ := json.MarshalIndent(res, "", "\t")
			c.logger.Warn().
				Int("size", len(msgs)).
				Str("res_json", string(resJSON)).
				Err(err).
				Msg("failed to commit msg batch")

			return nil, errors.WithStack(err)
		}
	}

	c.accSeq++

	return res, nil
}

// AsyncBroadcastMsg sends Tx to chain and doesn't wait until Tx is included in block. This method
// cannot be used for rapid Tx sending, it is expected that you wait for transaction status with
// external tools. If you want sdk to wait for it, use SyncBroadcastMsg.
func (c *chainClient) AsyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error) {
	c.syncMux.Lock()
	defer c.syncMux.Unlock()

	c.txFactory = c.txFactory.WithSequence(c.accSeq)
	c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
	res, err := c.broadcastTx(context.TODO(), c.clientContext, c.txFactory, false, msgs...)
	if err != nil {
		if strings.Contains(err.Error(), "account sequence mismatch") {
			c.syncNonce()
			c.txFactory = c.txFactory.WithSequence(c.accSeq)
			c.txFactory = c.txFactory.WithAccountNumber(c.accNum)

			c.logger.Debug().Uint64("nonce", c.accSeq).Msg("retrying broadcastTx with nonce")

			res, err = c.broadcastTx(context.TODO(), c.clientContext, c.txFactory, false, msgs...)
		}
		if err != nil {
			resJSON, _ := json.MarshalIndent(res, "", "\t")
			c.logger.Warn().
				Int("size", len(msgs)).
				Str("res_json", string(resJSON)).
				Err(err).
				Msg("failed to commit msg batch")

			return nil, errors.WithStack(err)
		}
	}

	c.accSeq++

	return res, nil
}

const (
	defaultBroadcastStatusPoll = 100 * time.Millisecond
	defaultBroadcastTimeout    = 40 * time.Second
)

func (c *chainClient) broadcastTx(
	ctx context.Context,
	clientCtx client.Context,
	txf tx.Factory,
	await bool,
	msgs ...sdk.Msg,
) (*txtypes.BroadcastTxResponse, error) {
	txf, err := c.prepareFactory(clientCtx, txf)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare Txx factory")
		return nil, err
	}

	if clientCtx.Simulate {
		simTxBytes, err := tx.BuildSimTx(txf, msgs...)
		if err != nil {
			err = errors.Wrap(err, "failed to build SimTx bytes")
			return nil, err
		}

		simRes, err := c.txClient.Simulate(ctx, &txtypes.SimulateRequest{TxBytes: simTxBytes})
		if err != nil {
			err = errors.Wrap(err, "failed to run SimulateRequest")
			return nil, err
		}

		adjustedGas := uint64(txf.GasAdjustment() * float64(simRes.GasInfo.GasUsed))
		txf = txf.WithGas(adjustedGas)
	}

	txn, err := tx.BuildUnsignedTx(txf, msgs...)
	if err != nil {
		err = errors.Wrap(err, "failed to BuildUnsignedTx")
		return nil, err
	}

	txn.SetFeeGranter(clientCtx.GetFeeGranterAddress())
	err = tx.Sign(txf, clientCtx.GetFromName(), txn, true)
	if err != nil {
		err = errors.Wrap(err, "failed to Sign Tx")
		return nil, err
	}

	txBytes, err := clientCtx.TxConfig.TxEncoder()(txn.GetTx())
	if err != nil {
		err = errors.Wrap(err, "failed TxEncoder to encode Tx")
		return nil, err
	}

	res, err := c.txClient.BroadcastTx(ctx, &txtypes.BroadcastTxRequest{
		TxBytes: txBytes,
		Mode:    txtypes.BroadcastMode_BROADCAST_MODE_SYNC,
	})
	if !await || err != nil {
		return res, err
	}

	awaitCtx, cancelFn := context.WithTimeout(context.Background(), defaultBroadcastTimeout)
	defer cancelFn()

	txHash, _ := hex.DecodeString(res.TxResponse.TxHash)
	t := time.NewTimer(defaultBroadcastStatusPoll)

	for {
		select {
		case <-awaitCtx.Done():
			err := errors.Wrapf(ErrTimedOut, "%s", res.TxResponse.TxHash)
			t.Stop()
			return nil, err
		case <-t.C:
			resultTx, err := clientCtx.Client.Tx(awaitCtx, txHash, false)
			if err != nil {
				if errRes := client.CheckTendermintError(err, txBytes); errRes != nil {
					return &txtypes.BroadcastTxResponse{TxResponse: errRes}, err
				}

				t.Reset(defaultBroadcastStatusPoll)
				continue

			} else if resultTx.Height > 0 {
				resResultTx := sdk.NewResponseResultTx(resultTx, res.TxResponse.Tx, res.TxResponse.Timestamp)
				res = &txtypes.BroadcastTxResponse{TxResponse: resResultTx}
				t.Stop()
				return res, err
			}

			t.Reset(defaultBroadcastStatusPoll)
		}
	}
}

const (
	defaultMsgEnqueueTimeout = 10 * time.Second
)

// QueueBroadcastMsg enqueues a list of messages. Messages will added to the queue
// and grouped into Txns in chunks. Use this method to mass broadcast Txns with efficiency.
func (c *chainClient) QueueBroadcastMsg(msgs ...sdk.Msg) error {
	if !c.canSign {
		return ErrReadOnly
	} else if atomic.LoadInt64(&c.closed) == 1 {
		return ErrQueueClosed
	}

	t := time.NewTimer(defaultMsgEnqueueTimeout)
	for _, msg := range msgs {
		select {
		case <-t.C:
			return ErrEnqueueTimeout
		case c.msgC <- msg:
		}
	}
	t.Stop()

	return nil
}

const (
	msgCommitBatchSizeLimit = 1024
	msgCommitBatchFlushTime = 500 * time.Millisecond
)

func (c *chainClient) runBatchBroadcast() {
	flushTimer := time.NewTimer(msgCommitBatchFlushTime)
	msgBatch := make([]sdk.Msg, 0, msgCommitBatchSizeLimit)

	submitBatch := func(toSubmit []sdk.Msg) {
		c.syncMux.Lock()
		defer c.syncMux.Unlock()

		c.txFactory = c.txFactory.WithSequence(c.accSeq)
		c.txFactory = c.txFactory.WithAccountNumber(c.accNum)

		c.logger.Debug().Uint64("nonce", c.accSeq).Msg("broadcastTx with nonce")

		res, err := c.broadcastTx(context.TODO(), c.clientContext, c.txFactory, true, toSubmit...)
		if err != nil {
			if strings.Contains(err.Error(), "account sequence mismatch") {
				c.syncNonce()
				c.txFactory = c.txFactory.WithSequence(c.accSeq)
				c.txFactory = c.txFactory.WithAccountNumber(c.accNum)

				c.logger.Debug().Uint64("nonce", c.accSeq).Msg("retrying broadcastTx with nonce")

				res, err = c.broadcastTx(context.TODO(), c.clientContext, c.txFactory, true, toSubmit...)
			}
			if err != nil {
				resJSON, _ := json.MarshalIndent(res, "", "\t")
				c.logger.Warn().
					Int("size", len(toSubmit)).
					Str("res_json", string(resJSON)).
					Err(err).
					Msg("failed to commit msg batch")

				return
			}
		}

		if res.TxResponse.Code != 0 {
			err = errors.Errorf("error %d (%s): %s", res.TxResponse.Code, res.TxResponse.Codespace, res.TxResponse.RawLog)
			c.logger.Warn().
				Int("size", len(toSubmit)).
				Str("tx_hash", string(res.TxResponse.TxHash)).
				Err(err).
				Msg("failed to commit msg batch")
		} else {
			c.logger.Debug().
				Str("tx_hash", string(res.TxResponse.TxHash)).
				Int64("height", res.TxResponse.Height).
				Msg("msg batch committed successfully at height")
		}

		c.accSeq++
		c.logger.Debug().Uint64("new_nonce", c.accSeq).Msg("nonce incremented")

	}

	for {
		select {
		case msg, ok := <-c.msgC:
			if !ok {
				// exit required
				if len(msgBatch) > 0 {
					submitBatch(msgBatch)
				}

				close(c.doneC)
				return
			}

			msgBatch = append(msgBatch, msg)

			if len(msgBatch) >= msgCommitBatchSizeLimit {
				toSubmit := msgBatch
				msgBatch = msgBatch[:0]
				flushTimer.Reset(msgCommitBatchFlushTime)

				submitBatch(toSubmit)
			}
		case <-flushTimer.C:
			if len(msgBatch) > 0 {
				toSubmit := msgBatch
				msgBatch = msgBatch[:0]
				flushTimer.Reset(msgCommitBatchFlushTime)
				submitBatch(toSubmit)
			} else {
				flushTimer.Reset(msgCommitBatchFlushTime)
			}
		}
	}
}
