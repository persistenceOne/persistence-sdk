package simapp

import (
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/persistenceOne/persistence-sdk/v6/x/lsm/distribution"
	"github.com/persistenceOne/persistence-sdk/v6/x/lsm/staking"

	"github.com/persistenceOne/persistence-sdk/v6/simapp/params"
	interchainquerytypes "github.com/persistenceOne/persistence-sdk/v6/x/interchainquery/types"
	oracletypes "github.com/persistenceOne/persistence-sdk/v6/x/oracle/types"
)

// MakeTestEncodingConfig creates an EncodingConfig for testing. This function
// should be used only in tests or when creating a new app instance (NewApp*()).
// App user shouldn't create new codecs - use the app.AppCodec instead.
// [DEPRECATED]
func MakeTestEncodingConfig() params.EncodingConfig {
	encodingConfig := params.MakeTestEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	// Register deprecated module codecs directly for historical data compatibility
	// This replaces the need for interchainquery.AppModuleBasic and oracle.AppModuleBasic in ModuleBasics
	interchainquerytypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	interchainquerytypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	oracletypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	oracletypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	//for parsing lsm msgs
	distribution.RegisterLegacyAminoCodec(encodingConfig.Amino)
	distribution.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	staking.RegisterLegacyAminoCodec(encodingConfig.Amino)
	staking.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	return encodingConfig
}
