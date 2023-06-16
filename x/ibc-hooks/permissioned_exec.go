package ibc_hooks

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/CosmWasm/wasmd/x/wasm/types"
)

type permissionedExecutor struct {
	keeper *wasmkeeper.PermissionedKeeper
}

type PermissionedExecutor interface {
	ExecuteContract(goCtx context.Context, msg *types.MsgExecuteContract) (*types.MsgExecuteContractResponse, error)
}

// NewPermissionedMsgExecutor constructor for ExecuteContract msg handler using permissioned keeper
//
// Since x/wasm 0.40 expects normal *wasmkeeper.Keeper in its MsgServerImpl, we use this shim instead
func NewPermissionedMsgExecutor(k *wasmkeeper.PermissionedKeeper) PermissionedExecutor {
	return &permissionedExecutor{keeper: k}
}

func (e *permissionedExecutor) ExecuteContract(goCtx context.Context, msg *types.MsgExecuteContract) (*types.MsgExecuteContractResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, errorsmod.Wrap(err, "sender")
	}
	contractAddr, err := sdk.AccAddressFromBech32(msg.Contract)
	if err != nil {
		return nil, errorsmod.Wrap(err, "contract")
	}

	data, err := e.keeper.Execute(ctx, contractAddr, senderAddr, msg.Msg, msg.Funds)
	if err != nil {
		return nil, err
	}

	return &types.MsgExecuteContractResponse{
		Data: data,
	}, nil
}
