package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v6/modules/core/exported"

	"github.com/persistenceOne/persistence-sdk/v2/utils"
)

var _ IBCHandshakeHooks = MultiIBCHandshakeHooks{}

// MultiIBCHandshakeHooks combine multiple ibc transfer hooks, all hook functions are run in array sequence
type MultiIBCHandshakeHooks []IBCHandshakeHooks

func NewMultiStakingHooks(hooks ...IBCHandshakeHooks) MultiIBCHandshakeHooks {
	return hooks
}

func (h MultiIBCHandshakeHooks) OnRecvPacket(ctx sdk.Context, packet types.Packet, relayer sdk.AccAddress, transferAck exported.Acknowledgement) error {
	for i := range h {
		wrappedHookFn := func(ctx sdk.Context) error {
			//nolint:scopelint
			return h[i].OnRecvPacket(ctx, packet, relayer, transferAck)
		}

		err := utils.ApplyFuncIfNoError(ctx, wrappedHookFn)
		if err != nil {
			ctx.Logger().Error("Error occurred in calling OnRecvPacket hooks, ", "err: ", err, "module:", ModuleName, "index:", i)
		}
	}

	return nil
}

func (h MultiIBCHandshakeHooks) OnAcknowledgementPacket(ctx sdk.Context, packet types.Packet, acknowledgement []byte, relayer sdk.AccAddress, transferAckErr error) error {
	for i := range h {
		wrappedHookFn := func(ctx sdk.Context) error {
			//nolint:scopelint
			return h[i].OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer, transferAckErr)
		}

		err := utils.ApplyFuncIfNoError(ctx, wrappedHookFn)
		if err != nil {
			ctx.Logger().Error("Error occurred in calling OnAcknowledgementPacket hooks, ", "err: ", err, "module:", ModuleName, "index:", i)
		}
	}

	return nil
}

func (h MultiIBCHandshakeHooks) OnTimeoutPacket(ctx sdk.Context, packet types.Packet, relayer sdk.AccAddress, transferTimeoutErr error) error {
	for i := range h {
		wrappedHookFn := func(ctx sdk.Context) error {
			//nolint:scopelint
			return h[i].OnTimeoutPacket(ctx, packet, relayer, transferTimeoutErr)
		}

		err := utils.ApplyFuncIfNoError(ctx, wrappedHookFn)
		if err != nil {
			ctx.Logger().Error("Error occurred in calling OnTimeoutPacket hooks, ", "err: ", err, "module:", ModuleName, "index:", i)
		}
	}

	return nil
}
