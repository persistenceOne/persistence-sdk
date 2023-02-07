package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	channelTypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v6/modules/core/exported"
)

func (k Keeper) OnRecvPacket(ctx sdk.Context, packet channelTypes.Packet, relayer sdk.AccAddress, transferAck exported.Acknowledgement) {
	if k.hooks != nil {
		_ = k.hooks.OnRecvPacket(ctx, packet, relayer, transferAck)
	}
}

func (k Keeper) OnAcknowledgementPacket(ctx sdk.Context, packet channelTypes.Packet, acknowledgement []byte, relayer sdk.AccAddress, transferAckErr error) {
	if k.hooks != nil {
		_ = k.hooks.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer, transferAckErr)
	}
}

func (k Keeper) OnTimeoutPacket(ctx sdk.Context, packet channelTypes.Packet, relayer sdk.AccAddress, transferTimeoutErr error) {
	if k.hooks != nil {
		_ = k.hooks.OnTimeoutPacket(ctx, packet, relayer, transferTimeoutErr)
	}
}
