package simapp

import (
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	group "github.com/cosmos/cosmos-sdk/x/group/module"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	ibc "github.com/cosmos/ibc-go/v6/modules/core"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/persistenceOne/persistence-sdk/v2/x/epochs"
	"github.com/persistenceOne/persistence-sdk/v2/x/halving"
	"github.com/persistenceOne/persistence-sdk/v2/x/ibchooker"
	"github.com/persistenceOne/persistence-sdk/v2/x/interchainquery"
)

func TestSimAppExportAndBlockedAddrs(t *testing.T) {
	encCfg := MakeTestEncodingConfig()
	db := dbm.NewMemDB()
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	app := NewSimappWithCustomOptions(t, false, SetupOptions{
		Logger:             logger,
		DB:                 db,
		InvCheckPeriod:     0,
		EncConfig:          encCfg,
		HomePath:           DefaultNodeHome,
		SkipUpgradeHeights: map[int64]bool{},
		AppOpts:            EmptyAppOptions{},
	})

	for acc := range maccPerms {
		require.True(
			t,
			app.BankKeeper.BlockedAddr(app.AccountKeeper.GetModuleAddress(acc)),
			"ensure that blocked addresses are properly set in bank keeper",
		)
	}

	app.Commit()

	logger2 := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	// Making a new app object with the db, so that initchain hasn't been called
	app2 := NewSimApp(logger2, db, nil, true, map[int64]bool{}, DefaultNodeHome, 0, encCfg, EmptyAppOptions{})
	_, err := app2.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}

func TestGetMaccPerms(t *testing.T) {
	dup := GetMaccPerms()
	require.Equal(t, maccPerms, dup, "duplicated module account permissions differed from actual module account permissions")
}

func TestRunMigrations(t *testing.T) {
	db := dbm.NewMemDB()
	encCfg := MakeTestEncodingConfig()
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	app := NewSimApp(logger, db, nil, true, map[int64]bool{}, DefaultNodeHome, 0, encCfg, EmptyAppOptions{})

	// Create a new baseapp and configurator for the purpose of this test.
	bApp := baseapp.NewBaseApp(appName, logger, db, encCfg.TxConfig.TxDecoder())
	bApp.SetCommitMultiStoreTracer(nil)
	bApp.SetInterfaceRegistry(encCfg.InterfaceRegistry)
	app.BaseApp = bApp
	app.configurator = module.NewConfigurator(app.appCodec, bApp.MsgServiceRouter(), app.GRPCQueryRouter())

	// We register all modules on the Configurator, except x/bank. x/bank will
	// serve as the test subject on which we run the migration tests.
	//
	// The loop below is the same as calling `RegisterServices` on
	// ModuleManager, except that we skip x/bank.
	for _, module := range app.mm.Modules {
		if module.Name() == banktypes.ModuleName {
			continue
		}

		module.RegisterServices(app.configurator)
	}

	// Initialize the chain
	app.InitChain(abci.RequestInitChain{})
	app.Commit()

	testCases := []struct { //nolint:maligned,testfile
		name         string
		moduleName   string
		fromVersion  uint64
		expRegErr    bool // errors while registering migration
		expRegErrMsg string
		expRunErr    bool // errors while running migration
		expRunErrMsg string
		expCalled    int
	}{
		{
			"cannot register migration for version 0",
			"bank", 0,
			true, "module migration versions should start at 1: invalid version", false, "", 0,
		},
		{
			"throws error on RunMigrations if no migration registered for bank",
			"", 1,
			false, "", true, "no migrations found for module bank: not found", 0,
		},
		{
			"can register 1->2 migration handler for x/bank, cannot run migration",
			"bank", 1,
			false, "", true, "no migration found for module bank from version 2 to version 3: not found", 0,
		},
		{
			"can register 2->3 migration handler for x/bank, can run migration",
			"bank", 2,
			false, "", false, "", 1,
		},
		{
			"cannot register migration handler for same module & fromVersion",
			"bank", 1,
			true, "another migration for module bank and version 1 already exists: internal logic error", false, "", 0,
		},
	}

	for _, tc := range testCases {
		tc2 := tc
		t.Run(tc2.name, func(t *testing.T) {
			var err error

			// Since it's very hard to test actual in-place store migrations in
			// tests (due to the difficulty of maintaining multiple versions of a
			// module), we're just testing here that the migration logic is
			// called.
			called := 0

			if tc2.moduleName != "" {
				// Register migration for module from version `fromVersion` to `fromVersion+1`.
				err = app.configurator.RegisterMigration(tc2.moduleName, tc2.fromVersion, func(sdk.Context) error {
					called++

					return nil
				})

				if tc2.expRegErr {
					require.EqualError(t, err, tc2.expRegErrMsg)

					return
				}
			}
			require.NoError(t, err)

			// Run migrations only for bank. That's why we put the initial
			// version for bank as 1, and for all other modules, we put as
			// their latest ConsensusVersion.
			_, err = app.mm.RunMigrations(
				app.NewContext(true, tmproto.Header{Height: app.LastBlockHeight()}), app.configurator,
				module.VersionMap{
					"bank":            1,
					"auth":            auth.AppModule{}.ConsensusVersion(),
					"authz":           authzmodule.AppModule{}.ConsensusVersion(),
					"staking":         staking.AppModule{}.ConsensusVersion(),
					"mint":            mint.AppModule{}.ConsensusVersion(),
					"distribution":    distribution.AppModule{}.ConsensusVersion(),
					"slashing":        slashing.AppModule{}.ConsensusVersion(),
					"gov":             gov.AppModule{}.ConsensusVersion(),
					"group":           group.AppModule{}.ConsensusVersion(),
					"params":          params.AppModule{}.ConsensusVersion(),
					"upgrade":         upgrade.AppModule{}.ConsensusVersion(),
					"vesting":         vesting.AppModule{}.ConsensusVersion(),
					"feegrant":        feegrantmodule.AppModule{}.ConsensusVersion(),
					"evidence":        evidence.AppModule{}.ConsensusVersion(),
					"crisis":          crisis.AppModule{}.ConsensusVersion(),
					"genutil":         genutil.AppModule{}.ConsensusVersion(),
					"capability":      capability.AppModule{}.ConsensusVersion(),
					"epochs":          epochs.AppModule{}.ConsensusVersion(),
					"halving":         halving.AppModule{}.ConsensusVersion(),
					"ibc":             ibc.AppModule{}.ConsensusVersion(),
					"interchainquery": interchainquery.AppModule{}.ConsensusVersion(),
					"ibchooker":       ibchooker.AppModule{}.ConsensusVersion(),
				},
			)
			if tc2.expRunErr {
				require.EqualError(t, err, tc2.expRunErrMsg)
			} else {
				require.NoError(t, err)
				// Make sure bank's migration is called.
				require.Equal(t, tc2.expCalled, called)
			}
		})
	}
}

func TestUpgradeStateOnGenesis(t *testing.T) {
	encCfg := MakeTestEncodingConfig()
	db := dbm.NewMemDB()
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	app := NewSimappWithCustomOptions(t, false, SetupOptions{
		Logger:             logger,
		DB:                 db,
		InvCheckPeriod:     0,
		EncConfig:          encCfg,
		HomePath:           DefaultNodeHome,
		SkipUpgradeHeights: map[int64]bool{},
		AppOpts:            EmptyAppOptions{},
	})

	// make sure the upgrade keeper has version map in state
	ctx := app.NewContext(false, tmproto.Header{})
	vm := app.UpgradeKeeper.GetModuleVersionMap(ctx)

	for v, i := range app.mm.Modules {
		require.Equal(t, vm[v], i.ConsensusVersion())
	}
}
