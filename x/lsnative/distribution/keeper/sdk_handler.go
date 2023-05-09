package keeper

import (
	context "context"

	sdkdistr "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/persistenceOne/persistence-sdk/v2/x/lsnative/distribution/types"
)

type SdkMsgHandler interface {
	sdkdistr.MsgServer
}

// NewSdkMsgHandlerWrapper returns SdkMsgHandler that implements MsgServer for
// vanilla sdk distribution keeper handler, allowing us to route legacy messages to a forked keeper.
func NewSdkMsgHandlerWrapper(handler types.MsgServer) SdkMsgHandler {
	return sdkMsgHandlerWrapper{
		handler: handler,
	}
}

var _ SdkMsgHandler = sdkMsgHandlerWrapper{}

type sdkMsgHandlerWrapper struct {
	handler types.MsgServer
}

// FundCommunityPool implements SdkMsgHandler
func (h sdkMsgHandlerWrapper) FundCommunityPool(
	ctx context.Context,
	sdkMsg *sdkdistr.MsgFundCommunityPool,
) (*sdkdistr.MsgFundCommunityPoolResponse, error) {
	msg := types.MsgFundCommunityPool(*sdkMsg)

	res, err := h.handler.FundCommunityPool(ctx, &msg)
	if err != nil {
		return nil, err
	}

	sdkResp := (sdkdistr.MsgFundCommunityPoolResponse)(*res)

	return &sdkResp, nil
}

// SetWithdrawAddress implements SdkMsgHandler
func (h sdkMsgHandlerWrapper) SetWithdrawAddress(
	ctx context.Context,
	sdkMsg *sdkdistr.MsgSetWithdrawAddress,
) (*sdkdistr.MsgSetWithdrawAddressResponse, error) {
	msg := types.MsgSetWithdrawAddress(*sdkMsg)

	res, err := h.handler.SetWithdrawAddress(ctx, &msg)
	if err != nil {
		return nil, err
	}

	sdkResp := (sdkdistr.MsgSetWithdrawAddressResponse)(*res)

	return &sdkResp, nil
}

// WithdrawDelegatorReward implements SdkMsgHandler
func (h sdkMsgHandlerWrapper) WithdrawDelegatorReward(
	ctx context.Context,
	sdkMsg *sdkdistr.MsgWithdrawDelegatorReward,
) (*sdkdistr.MsgWithdrawDelegatorRewardResponse, error) {
	msg := types.MsgWithdrawDelegatorReward(*sdkMsg)

	res, err := h.handler.WithdrawDelegatorReward(ctx, &msg)
	if err != nil {
		return nil, err
	}

	sdkResp := (sdkdistr.MsgWithdrawDelegatorRewardResponse)(*res)

	return &sdkResp, nil
}

// WithdrawValidatorCommission implements SdkMsgHandler
func (h sdkMsgHandlerWrapper) WithdrawValidatorCommission(
	ctx context.Context,
	sdkMsg *sdkdistr.MsgWithdrawValidatorCommission,
) (*sdkdistr.MsgWithdrawValidatorCommissionResponse, error) {
	msg := types.MsgWithdrawValidatorCommission(*sdkMsg)

	res, err := h.handler.WithdrawValidatorCommission(ctx, &msg)
	if err != nil {
		return nil, err
	}

	sdkResp := (sdkdistr.MsgWithdrawValidatorCommissionResponse)(*res)

	return &sdkResp, nil
}
