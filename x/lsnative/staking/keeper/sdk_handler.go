package keeper

import (
	context "context"
	sdkstaking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/persistenceOne/persistence-sdk/v2/x/lsnative/staking/types"
)

type SdkMsgHandler interface {
	sdkstaking.MsgServer
}

// NewSdkMsgHandlerWrapper returns SdkMsgHandler that implements MsgServer for
// vanilla sdk staking keeper handler, allowing us to route legacy messages to a forked keeper.
func NewSdkMsgHandlerWrapper(handler types.MsgServer) SdkMsgHandler {
	return sdkMsgHandlerWrapper{
		handler: handler,
	}
}

var _ SdkMsgHandler = sdkMsgHandlerWrapper{}

type sdkMsgHandlerWrapper struct {
	handler types.MsgServer
}

// BeginRedelegate implements SdkMsgHandler
func (h sdkMsgHandlerWrapper) BeginRedelegate(
	ctx context.Context,
	sdkMsg *sdkstaking.MsgBeginRedelegate,
) (*sdkstaking.MsgBeginRedelegateResponse, error) {
	msg := types.MsgBeginRedelegate(*sdkMsg)

	res, err := h.handler.BeginRedelegate(ctx, &msg)
	if err != nil {
		return nil, err
	}

	sdkResp := (sdkstaking.MsgBeginRedelegateResponse)(*res)

	return &sdkResp, nil
}

// CancelUnbondingDelegation implements SdkMsgHandler
func (h sdkMsgHandlerWrapper) CancelUnbondingDelegation(
	ctx context.Context,
	sdkMsg *sdkstaking.MsgCancelUnbondingDelegation,
) (*sdkstaking.MsgCancelUnbondingDelegationResponse, error) {
	msg := types.MsgCancelUnbondingDelegation(*sdkMsg)

	res, err := h.handler.CancelUnbondingDelegation(ctx, &msg)
	if err != nil {
		return nil, err
	}

	sdkResp := (sdkstaking.MsgCancelUnbondingDelegationResponse)(*res)

	return &sdkResp, nil
}

// CreateValidator implements SdkMsgHandler
func (h sdkMsgHandlerWrapper) CreateValidator(
	ctx context.Context,
	sdkMsg *sdkstaking.MsgCreateValidator,
) (*sdkstaking.MsgCreateValidatorResponse, error) {
	msg := types.MsgCreateValidator{
		Description:      types.Description(sdkMsg.Description),
		Commission:       types.CommissionRates(sdkMsg.Commission),
		DelegatorAddress: sdkMsg.DelegatorAddress,
		ValidatorAddress: sdkMsg.ValidatorAddress,
		Pubkey:           sdkMsg.Pubkey,
		Value:            sdkMsg.Value,

		// MinSelfDelegation omitted
	}

	res, err := h.handler.CreateValidator(ctx, &msg)
	if err != nil {
		return nil, err
	}

	sdkResp := (sdkstaking.MsgCreateValidatorResponse)(*res)

	return &sdkResp, nil
}

// Delegate implements SdkMsgHandler
func (h sdkMsgHandlerWrapper) Delegate(
	ctx context.Context,
	sdkMsg *sdkstaking.MsgDelegate,
) (*sdkstaking.MsgDelegateResponse, error) {
	msg := types.MsgDelegate(*sdkMsg)

	res, err := h.handler.Delegate(ctx, &msg)
	if err != nil {
		return nil, err
	}

	sdkResp := (sdkstaking.MsgDelegateResponse)(*res)

	return &sdkResp, nil
}

// EditValidator implements SdkMsgHandler
func (h sdkMsgHandlerWrapper) EditValidator(
	ctx context.Context,
	sdkMsg *sdkstaking.MsgEditValidator,
) (*sdkstaking.MsgEditValidatorResponse, error) {
	msg := types.MsgEditValidator{
		Description:      types.Description(sdkMsg.Description),
		ValidatorAddress: sdkMsg.ValidatorAddress,
		CommissionRate:   sdkMsg.CommissionRate,

		// MinSelfDelegation omitted
	}

	res, err := h.handler.EditValidator(ctx, &msg)
	if err != nil {
		return nil, err
	}

	sdkResp := (sdkstaking.MsgEditValidatorResponse)(*res)

	return &sdkResp, nil
}

// Undelegate implements SdkMsgHandler
func (h sdkMsgHandlerWrapper) Undelegate(
	ctx context.Context,
	sdkMsg *sdkstaking.MsgUndelegate,
) (*sdkstaking.MsgUndelegateResponse, error) {
	msg := types.MsgUndelegate(*sdkMsg)

	res, err := h.handler.Undelegate(ctx, &msg)
	if err != nil {
		return nil, err
	}

	sdkResp := (sdkstaking.MsgUndelegateResponse)(*res)

	return &sdkResp, nil
}
