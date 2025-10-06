package staking

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary x/halving interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgUnbondValidator{}, "cosmos-sdk/MsgUnbondValidator", nil)
	cdc.RegisterConcrete(&MsgTokenizeShares{}, "cosmos-sdk/MsgTokenizeShares", nil)
	cdc.RegisterConcrete(&MsgRedeemTokensForShares{}, "cosmos-sdk/MsgRedeemTokensForShares", nil)
	cdc.RegisterConcrete(&MsgTransferTokenizeShareRecord{}, "cosmos-sdk/MsgTransferTokenizeShareRecord", nil)
	cdc.RegisterConcrete(&MsgDisableTokenizeShares{}, "cosmos-sdk/MsgDisableTokenizeShares", nil)
	cdc.RegisterConcrete(&MsgEnableTokenizeShares{}, "cosmos-sdk/MsgEnableTokenizeShares", nil)
	cdc.RegisterConcrete(&MsgValidatorBond{}, "cosmos-sdk/MsgValidatorBond", nil)

}

// RegisterInterfaces registers the x/halving interfaces types with the interface registry
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnbondValidator{},
		&MsgTokenizeShares{},
		&MsgRedeemTokensForShares{},
		&MsgTransferTokenizeShareRecord{},
		&MsgDisableTokenizeShares{},
		&MsgEnableTokenizeShares{},
		&MsgValidatorBond{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_Lsm_serviceDesc)
}
