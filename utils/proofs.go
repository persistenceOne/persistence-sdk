package utils

import (
	"fmt"
	"net/url"

	"cosmossdk.io/errors"
	"github.com/cometbft/cometbft/proto/tendermint/crypto"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	commitmenttypes "github.com/cosmos/ibc-go/v7/modules/core/23-commitment/types"
	ibcKeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"
)

func convertProof(cdc codec.BinaryCodec, proofOps *crypto.ProofOps) ([]byte, error) {
	if proofOps == nil {
		return nil, fmt.Errorf("unable to validate proof. No proof submitted")
	}

	merkleProof, err := commitmenttypes.ConvertProofs(proofOps)
	if err != nil {
		return nil, errors.Wrap(err, "error converting proofs")
	}

	proof, err := cdc.Marshal(&merkleProof)
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal merkle proof")
	}

	return proof, nil
}

func ValidateProofOps(
	ctx sdk.Context, ibcKeeper *ibcKeeper.Keeper,
	connectionID string, chainID string,
	height int64, module string, key []byte,
	data []byte, proofOps *crypto.ProofOps,
) error {
	cdc := ibcKeeper.Codec()
	proof, err := convertProof(cdc, proofOps)
	if err != nil {
		return err
	}

	path := commitmenttypes.NewMerklePath([]string{module, url.PathEscape(string(key))}...)
	connection, found := ibcKeeper.ConnectionKeeper.GetConnection(ctx, connectionID)
	if !found {
		return fmt.Errorf("connection %s not found", connectionID)
	}

	clientState, found := ibcKeeper.ClientKeeper.GetClientState(ctx, connection.ClientId)
	if !found {
		return fmt.Errorf("unable to fetch client state")
	}

	clientStore := ibcKeeper.ClientKeeper.ClientStore(ctx, connection.ClientId)
	csHeight := clienttypes.NewHeight(clienttypes.ParseChainID(chainID), uint64(height)+1)

	if len(data) != 0 {
		if err := clientState.VerifyMembership(
			ctx, clientStore, cdc, csHeight,
			0, 0, // skip delay period checks for non-packet processing verification
			proof, path, data,
		); err != nil {
			return errors.Wrap(err, "unable to verify inclusion proof")
		}
	} else {
		if err := clientState.VerifyNonMembership(
			ctx, clientStore, cdc, csHeight,
			0, 0, // skip delay period checks for non-packet processing verification
			proof, path,
		); err != nil {
			return errors.Wrap(err, "unable to verify non-inclusion proof")
		}
	}

	return nil
}
