package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v3/modules/core/exported"
)

var _ IBCTransferHooks = MultiIBCTransferHooks{}

// combine multiple ibc transfer hooks, all hook functions are run in array sequence
type MultiIBCTransferHooks []IBCTransferHooks

func NewMultiStakingHooks(hooks ...IBCTransferHooks) MultiIBCTransferHooks {
	return hooks
}

func (h MultiIBCTransferHooks) OnRecvPacket(ctx sdk.Context, packet types.Packet, relayer sdk.AccAddress, success bool, transferAck exported.Acknowledgement) {
	for i := range h {
		h[i].OnRecvPacket(ctx, packet, relayer, success, transferAck)
	}
}

func (h MultiIBCTransferHooks) OnAcknowledgementPacket(ctx sdk.Context, packet types.Packet, acknowledgement []byte, relayer sdk.AccAddress, transferAckErr error) {
	for i := range h {
		h[i].OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer, transferAckErr)
	}
}

func (h MultiIBCTransferHooks) OnTimeoutPacket(ctx sdk.Context, packet types.Packet, relayer sdk.AccAddress, transferTimeoutErr error) {
	for i := range h {
		h[i].OnTimeoutPacket(ctx, packet, relayer, transferTimeoutErr)
	}
}
