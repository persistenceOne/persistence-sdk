package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v3/modules/core/exported"
)

type IBCHandshakeHooks interface {
	// Do we care about `relayer sdk.AccAddress` argument here?
	OnRecvPacket(ctx sdk.Context, packet types.Packet, relayer sdk.AccAddress, transferAck exported.Acknowledgement) error
	OnAcknowledgementPacket(ctx sdk.Context, packet types.Packet, acknowledgement []byte, relayer sdk.AccAddress, transferAckErr error) error
	OnTimeoutPacket(ctx sdk.Context, packet types.Packet, relayer sdk.AccAddress, transferTimeoutErr error) error
}
