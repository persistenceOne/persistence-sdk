package keeper

import (
	context "context"
	sdkslashing "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/persistenceOne/persistence-sdk/v2/x/lsnative/slashing/types"
)

type SdkMsgHandler interface {
	sdkslashing.MsgServer
}

// NewSdkMsgHandlerWrapper returns SdkMsgHandler that implements MsgServer for
// vanilla sdk slashing keeper handler, allowing us to route legacy messages to a forked keeper.
func NewSdkMsgHandlerWrapper(handler types.MsgServer) SdkMsgHandler {
	return sdkMsgHandlerWrapper{
		handler: handler,
	}
}

var _ SdkMsgHandler = sdkMsgHandlerWrapper{}

type sdkMsgHandlerWrapper struct {
	handler types.MsgServer
}

// Unjail implements SdkMsgHandler
func (h sdkMsgHandlerWrapper) Unjail(
	ctx context.Context,
	sdkMsg *sdkslashing.MsgUnjail,
) (*sdkslashing.MsgUnjailResponse, error) {
	msg := types.MsgUnjail(*sdkMsg)

	res, err := h.handler.Unjail(ctx, &msg)
	if err != nil {
		return nil, err
	}

	sdkResp := (sdkslashing.MsgUnjailResponse)(*res)

	return &sdkResp, nil
}
