package keeper_test

import (
	"testing"
	"time"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"

	"github.com/persistenceOne/persistence-sdk/v4/simapp"
	"github.com/persistenceOne/persistence-sdk/v4/utils"
	"github.com/persistenceOne/persistence-sdk/v4/x/epochs/types"
)

// This test is responsible for testing how epochs increment based off
// of their initial conditions, and subsequent block height / times.
func (suite *KeeperTestSuite) TestEpochInfoBeginBlockChanges() {
	block1Time := time.Unix(1656907200, 0).UTC()

	const defaultIdentifier = "hourly"

	const defaultDuration = time.Hour
	// eps is short for epsilon - in this case a negligible amount of time.
	const eps = time.Nanosecond

	tests := map[string]struct {
		// if identifier, duration is not set, we make it defaultIdentifier and defaultDuration.
		// EpochCountingStarted, if unspecified, is inferred by CurrentEpoch == 0
		// StartTime is inferred to be block1Time if left blank.
		initialEpochInfo     types.EpochInfo
		blockHeightTimePairs map[int]time.Time
		expEpochInfo         types.EpochInfo
	}{
		"First block running at exactly start time sets epoch tick": {
			initialEpochInfo: types.EpochInfo{StartTime: block1Time, CurrentEpoch: 0, CurrentEpochStartTime: time.Time{}},
			expEpochInfo:     types.EpochInfo{StartTime: block1Time, CurrentEpoch: 1, CurrentEpochStartTime: block1Time, CurrentEpochStartHeight: 1},
		},
		"First block run sets start time, subsequent blocks within timer interval do not cause timer tick": {
			initialEpochInfo:     types.EpochInfo{StartTime: block1Time, CurrentEpoch: 0, CurrentEpochStartTime: time.Time{}},
			blockHeightTimePairs: map[int]time.Time{2: block1Time.Add(time.Second), 3: block1Time.Add(time.Minute), 4: block1Time.Add(30 * time.Minute)},
			expEpochInfo:         types.EpochInfo{StartTime: block1Time, CurrentEpoch: 1, CurrentEpochStartTime: block1Time, CurrentEpochStartHeight: 1},
		},
		"Second block at exactly timer interval later does not tick": {
			initialEpochInfo:     types.EpochInfo{StartTime: block1Time, CurrentEpoch: 0, CurrentEpochStartTime: time.Time{}},
			blockHeightTimePairs: map[int]time.Time{2: block1Time.Add(defaultDuration)},
			expEpochInfo:         types.EpochInfo{StartTime: block1Time, CurrentEpoch: 1, CurrentEpochStartTime: block1Time, CurrentEpochStartHeight: 1},
		},
		"Second block at timer interval + epsilon later does tick": {
			initialEpochInfo:     types.EpochInfo{StartTime: block1Time, CurrentEpoch: 0, CurrentEpochStartTime: time.Time{}},
			blockHeightTimePairs: map[int]time.Time{2: block1Time.Add(defaultDuration).Add(eps)},
			expEpochInfo:         types.EpochInfo{StartTime: block1Time, CurrentEpoch: 2, CurrentEpochStartTime: block1Time.Add(time.Hour), CurrentEpochStartHeight: 2},
		},
		"Downtime recovery (many intervals), first block causes 1 tick and sets current start time 1 interval ahead": {
			initialEpochInfo:     types.EpochInfo{StartTime: block1Time, CurrentEpoch: 0, CurrentEpochStartTime: time.Time{}},
			blockHeightTimePairs: map[int]time.Time{2: block1Time.Add(24 * time.Hour)},
			expEpochInfo:         types.EpochInfo{StartTime: block1Time, CurrentEpoch: 2, CurrentEpochStartTime: block1Time.Add(time.Hour), CurrentEpochStartHeight: 2},
		},
		"Downtime recovery (many intervals), second block is at tick 2, w/ start time 2 intervals ahead": {
			initialEpochInfo:     types.EpochInfo{StartTime: block1Time, CurrentEpoch: 0, CurrentEpochStartTime: time.Time{}},
			blockHeightTimePairs: map[int]time.Time{2: block1Time.Add(24 * time.Hour), 3: block1Time.Add(24 * time.Hour).Add(eps)},
			expEpochInfo:         types.EpochInfo{StartTime: block1Time, CurrentEpoch: 3, CurrentEpochStartTime: block1Time.Add(2 * time.Hour), CurrentEpochStartHeight: 3},
		},
		"Many blocks between first and second tick": {
			initialEpochInfo:     types.EpochInfo{StartTime: block1Time, CurrentEpoch: 1, CurrentEpochStartTime: block1Time},
			blockHeightTimePairs: map[int]time.Time{2: block1Time.Add(time.Second), 3: block1Time.Add(2 * time.Second), 4: block1Time.Add(time.Hour).Add(eps)},
			expEpochInfo:         types.EpochInfo{StartTime: block1Time, CurrentEpoch: 2, CurrentEpochStartTime: block1Time.Add(time.Hour), CurrentEpochStartHeight: 4},
		},
		"Distinct identifier and duration still works": {
			initialEpochInfo:     types.EpochInfo{Identifier: "hello", Duration: time.Minute, StartTime: block1Time, CurrentEpoch: 0, CurrentEpochStartTime: time.Time{}},
			blockHeightTimePairs: map[int]time.Time{2: block1Time.Add(time.Second), 3: block1Time.Add(time.Minute).Add(eps)},
			expEpochInfo:         types.EpochInfo{Identifier: "hello", Duration: time.Minute, StartTime: block1Time, CurrentEpoch: 2, CurrentEpochStartTime: block1Time.Add(time.Minute), CurrentEpochStartHeight: 3},
		},
		"StartTime in future won't get ticked on first block": {
			initialEpochInfo: types.EpochInfo{StartTime: block1Time.Add(time.Second), CurrentEpoch: 0, CurrentEpochStartTime: time.Time{}},
			// currentEpochStartHeight is 1 because thats when the timer was created on-chain
			expEpochInfo: types.EpochInfo{StartTime: block1Time.Add(time.Second), CurrentEpoch: 0, CurrentEpochStartTime: time.Time{}, CurrentEpochStartHeight: 1},
		},
		"StartTime in past will get ticked on first block": {
			initialEpochInfo: types.EpochInfo{StartTime: block1Time.Add(-time.Second), CurrentEpoch: 0, CurrentEpochStartTime: time.Time{}},
			expEpochInfo:     types.EpochInfo{StartTime: block1Time.Add(-time.Second), CurrentEpoch: 1, CurrentEpochStartTime: block1Time.Add(-time.Second), CurrentEpochStartHeight: 1},
		},
	}
	for name, test := range tests {
		//nolint:scopelint,testfile
		suite.Run(name, func() {
			suite.SetupTest()
			suite.Ctx = suite.Ctx.WithBlockHeight(1).WithBlockTime(block1Time)
			initialEpoch := initializeBlankEpochInfoFields(test.initialEpochInfo, defaultIdentifier, defaultDuration)
			err := suite.App.EpochsKeeper.AddEpochInfo(suite.Ctx, initialEpoch)
			suite.Require().NoError(err)
			suite.App.EpochsKeeper.BeginBlocker(suite.Ctx)

			// get sorted heights
			heights := maps.Keys(test.blockHeightTimePairs)
			utils.SortSlice(heights)
			for _, h := range heights {
				// for each height in order, run begin block
				suite.Ctx = suite.Ctx.WithBlockHeight(int64(h)).WithBlockTime(test.blockHeightTimePairs[h])
				suite.App.EpochsKeeper.BeginBlocker(suite.Ctx)
			}
			expEpoch := initializeBlankEpochInfoFields(test.expEpochInfo, initialEpoch.Identifier, initialEpoch.Duration)
			actEpoch := suite.App.EpochsKeeper.GetEpochInfo(suite.Ctx, initialEpoch.Identifier)
			suite.Require().Equal(expEpoch, actEpoch)
		})
	}
}

// initializeBlankEpochInfoFields set identifier, duration and epochCountingStarted if blank in epoch
func initializeBlankEpochInfoFields(epoch types.EpochInfo, identifier string, duration time.Duration) types.EpochInfo {
	if epoch.Identifier == "" {
		epoch.Identifier = identifier
	}

	if epoch.Duration == time.Duration(0) {
		epoch.Duration = duration
	}

	epoch.EpochCountingStarted = (epoch.CurrentEpoch != 0)

	return epoch
}

func TestEpochStartingOneMonthAfterInitGenesis(t *testing.T) {
	app := simapp.Setup(t, false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	// On init genesis, default epochs information is set
	// To check init genesis again, should make it fresh status
	epochInfos := app.EpochsKeeper.AllEpochInfos(ctx)
	for _, epochInfo := range epochInfos {
		app.EpochsKeeper.DeleteEpochInfo(ctx, epochInfo.Identifier)
	}

	now := time.Now()
	week := time.Hour * 24 * 7
	month := time.Hour * 24 * 30
	initialBlockHeight := int64(1)
	ctx = ctx.WithBlockHeight(initialBlockHeight).WithBlockTime(now)

	app.EpochsKeeper.InitGenesis(ctx, types.GenesisState{
		Epochs: []types.EpochInfo{
			{
				Identifier:              "monthly",
				StartTime:               now.Add(month),
				Duration:                time.Hour * 24 * 30,
				CurrentEpoch:            0,
				CurrentEpochStartHeight: ctx.BlockHeight(),
				CurrentEpochStartTime:   time.Time{},
				EpochCountingStarted:    false,
			},
		},
	})

	// epoch not started yet
	epochInfo := app.EpochsKeeper.GetEpochInfo(ctx, "monthly")
	require.Equal(t, epochInfo.CurrentEpoch, int64(0))
	require.Equal(t, epochInfo.CurrentEpochStartHeight, initialBlockHeight)
	require.Equal(t, epochInfo.CurrentEpochStartTime, time.Time{})
	require.Equal(t, epochInfo.EpochCountingStarted, false)

	// after 1 week
	ctx = ctx.WithBlockHeight(2).WithBlockTime(now.Add(week))
	app.EpochsKeeper.BeginBlocker(ctx)

	// epoch not started yet
	epochInfo = app.EpochsKeeper.GetEpochInfo(ctx, "monthly")
	require.Equal(t, epochInfo.CurrentEpoch, int64(0))
	require.Equal(t, epochInfo.CurrentEpochStartHeight, initialBlockHeight)
	require.Equal(t, epochInfo.CurrentEpochStartTime, time.Time{})
	require.Equal(t, epochInfo.EpochCountingStarted, false)

	// after 1 month
	ctx = ctx.WithBlockHeight(3).WithBlockTime(now.Add(month))
	app.EpochsKeeper.BeginBlocker(ctx)

	// epoch started
	epochInfo = app.EpochsKeeper.GetEpochInfo(ctx, "monthly")
	require.Equal(t, epochInfo.CurrentEpoch, int64(1))
	require.Equal(t, epochInfo.CurrentEpochStartHeight, ctx.BlockHeight())
	require.Equal(t, epochInfo.CurrentEpochStartTime.UTC().String(), now.Add(month).UTC().String())
	require.Equal(t, epochInfo.EpochCountingStarted, true)
}
