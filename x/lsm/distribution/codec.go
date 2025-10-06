package distribution

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary x/halving interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgWithdrawTokenizeShareRecordReward{}, "cosmos-sdk/distr/MsgWithdrawTokenizeShareRecordReward", nil)
	cdc.RegisterConcrete(&MsgWithdrawAllTokenizeShareRecordReward{}, "cosmos-sdk/distr/MsgWithdrawAllTokenizeShareRecordReward", nil)
}

// RegisterInterfaces registers the x/halving interfaces types with the interface registry
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdrawTokenizeShareRecordReward{},
		&MsgWithdrawAllTokenizeShareRecordReward{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_Lsm_serviceDesc)
}
