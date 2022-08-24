package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v3/modules/core/exported"
	"github.com/persistenceOne/persistence-sdk/utils"
)

var _ IBCHandshakeHooks = MultiIBCHandshakeHooks{}

// MultiIBCHandshakeHooks combine multiple ibc transfer hooks, all hook functions are run in array sequence
type MultiIBCHandshakeHooks []IBCHandshakeHooks

func NewMultiStakingHooks(hooks ...IBCHandshakeHooks) MultiIBCHandshakeHooks {
	return hooks
}

func (h MultiIBCHandshakeHooks) OnRecvPacket(ctx sdk.Context, packet types.Packet, relayer sdk.AccAddress, transferAck exported.Acknowledgement) {
	for i := range h {
		panicCatchingOnRecvPacketFnHook(ctx, h[i].OnRecvPacket, packet, relayer, transferAck)
	}
}

func (h MultiIBCHandshakeHooks) OnAcknowledgementPacket(ctx sdk.Context, packet types.Packet, acknowledgement []byte, relayer sdk.AccAddress, transferAckErr error) {
	for i := range h {
		panicCatchingOnAcknowledgementPacketFnHook(ctx, h[i].OnAcknowledgementPacket, packet, acknowledgement, relayer, transferAckErr)
	}
}

func (h MultiIBCHandshakeHooks) OnTimeoutPacket(ctx sdk.Context, packet types.Packet, relayer sdk.AccAddress, transferTimeoutErr error) {
	for i := range h {
		panicCatchingOnTimeoutPacketFnHook(ctx, h[i].OnTimeoutPacket, packet, relayer, transferTimeoutErr)
	}
}

// Panic catching Fn implementations
// We want to be using cacheContext in case of failure not to write entire code block
func panicCatchingOnRecvPacketFnHook(
	ctx sdk.Context,
	onRecvPacketFn func(ctx sdk.Context, packet types.Packet, relayer sdk.AccAddress, transferAck exported.Acknowledgement),
	packet types.Packet,
	relayer sdk.AccAddress,
	transferAck exported.Acknowledgement,
) {
	defer func() {
		if recovErr := recover(); recovErr != nil {
			utils.PrintPanicRecoveryError(ctx, recovErr)
		}
	}()

	cacheCtx, write := ctx.CacheContext()
	onRecvPacketFn(cacheCtx, packet, relayer, transferAck)
	write()
}

func panicCatchingOnAcknowledgementPacketFnHook(
	ctx sdk.Context,
	onAcknowledgementPacketFn func(ctx sdk.Context, packet types.Packet, acknowledgement []byte, relayer sdk.AccAddress, transferAckErr error),
	packet types.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
	transferAckErr error,
) {
	defer func() {
		if recovErr := recover(); recovErr != nil {
			utils.PrintPanicRecoveryError(ctx, recovErr)
		}
	}()

	cacheCtx, write := ctx.CacheContext()
	onAcknowledgementPacketFn(cacheCtx, packet, acknowledgement, relayer, transferAckErr)
	write()
}

func panicCatchingOnTimeoutPacketFnHook(
	ctx sdk.Context,
	onTimeoutPacketFn func(ctx sdk.Context, packet types.Packet, relayer sdk.AccAddress, transferTimeoutErr error),
	packet types.Packet,
	relayer sdk.AccAddress,
	transferTimeoutErr error,
) {
	defer func() {
		if recovErr := recover(); recovErr != nil {
			utils.PrintPanicRecoveryError(ctx, recovErr)
		}
	}()

	cacheCtx, write := ctx.CacheContext()
	onTimeoutPacketFn(cacheCtx, packet, relayer, transferTimeoutErr)
	write()
}
