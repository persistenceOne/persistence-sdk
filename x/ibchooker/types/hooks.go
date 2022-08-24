package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v3/modules/core/exported"
)

var _ IBCHandshakeHooks = MultiIBCHandshakeHooks{}

// MultiIBCHandshakeHooks combine multiple ibc transfer hooks, all hook functions are run in array sequence
type MultiIBCHandshakeHooks []IBCHandshakeHooks

func NewMultiStakingHooks(hooks ...IBCHandshakeHooks) MultiIBCHandshakeHooks {
	return hooks
}

func (h MultiIBCHandshakeHooks) OnRecvPacket(ctx sdk.Context, packet types.Packet, relayer sdk.AccAddress, transferAck exported.Acknowledgement) {
	for i := range h {
		h[i].OnRecvPacket(ctx, packet, relayer, transferAck)
	}
}

func (h MultiIBCHandshakeHooks) OnAcknowledgementPacket(ctx sdk.Context, packet types.Packet, acknowledgement []byte, relayer sdk.AccAddress, transferAckErr error) {
	for i := range h {
		h[i].OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer, transferAckErr)
	}
}

func (h MultiIBCHandshakeHooks) OnTimeoutPacket(ctx sdk.Context, packet types.Packet, relayer sdk.AccAddress, transferTimeoutErr error) {
	for i := range h {
		h[i].OnTimeoutPacket(ctx, packet, relayer, transferTimeoutErr)
	}
}
