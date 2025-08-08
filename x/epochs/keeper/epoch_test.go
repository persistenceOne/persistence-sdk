package keeper_test

import (
	"time"

	"github.com/persistenceOne/persistence-sdk/v4/x/epochs/types"
)

func (suite *KeeperTestSuite) TestAddEpochInfo() {
	defaultIdentifier := "default_add_epoch_info_id"
	defaultDuration := time.Hour
	startBlockHeight := int64(100)
	startBlockTime := time.Unix(1656907200, 0).UTC()
	tests := map[string]struct {
		addedEpochInfo types.EpochInfo
		expErr         bool
		expEpochInfo   types.EpochInfo
	}{
		"simple_add": {
			addedEpochInfo: types.EpochInfo{
				Identifier:              defaultIdentifier,
				StartTime:               time.Time{},
				Duration:                defaultDuration,
				CurrentEpoch:            0,
				CurrentEpochStartHeight: 0,
				CurrentEpochStartTime:   time.Time{},
				EpochCountingStarted:    false,
			},
			expErr: false,
			expEpochInfo: types.EpochInfo{
				Identifier:              defaultIdentifier,
				StartTime:               startBlockTime,
				Duration:                defaultDuration,
				CurrentEpoch:            0,
				CurrentEpochStartHeight: startBlockHeight,
				CurrentEpochStartTime:   time.Time{},
				EpochCountingStarted:    false,
			},
		},
		"zero_duration": {
			addedEpochInfo: types.EpochInfo{
				Identifier:              defaultIdentifier,
				StartTime:               time.Time{},
				Duration:                time.Duration(0),
				CurrentEpoch:            0,
				CurrentEpochStartHeight: 0,
				CurrentEpochStartTime:   time.Time{},
				EpochCountingStarted:    false,
			},
			expErr: true,
		},
	}

	for name, test := range tests {
		//nolint:scopelint,testfile
		suite.Run(name, func() {
			suite.SetupTest()
			ctx := suite.Ctx.WithBlockHeight(startBlockHeight).WithBlockTime(startBlockTime)
			err := suite.EpochsKeeper.AddEpochInfo(ctx, test.addedEpochInfo)
			if !test.expErr {
				suite.Require().NoError(err)
				actualEpochInfo := suite.EpochsKeeper.GetEpochInfo(ctx, test.addedEpochInfo.Identifier)
				suite.Require().Equal(test.expEpochInfo, actualEpochInfo)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestDuplicateAddEpochInfo() {
	suite.SetupTest()
	ctx := suite.Ctx

	identifier := "duplicate_add_epoch_info"
	epochInfo := types.NewGenesisEpochInfo(identifier, time.Hour*24*30)
	err := suite.EpochsKeeper.AddEpochInfo(ctx, epochInfo)
	suite.Require().NoError(err)
	err = suite.EpochsKeeper.AddEpochInfo(ctx, epochInfo)
	suite.Require().Error(err)
}

func (suite *KeeperTestSuite) TestEpochLifeCycle() {
	suite.SetupTest()
	ctx := suite.Ctx

	epochInfo := types.NewGenesisEpochInfo("monthly", time.Hour*24*30)
	err := suite.EpochsKeeper.AddEpochInfo(ctx, epochInfo)
	suite.Require().NoError(err)
	epochInfoSaved := suite.EpochsKeeper.GetEpochInfo(ctx, "monthly")
	// setup expected epoch info
	expectedEpochInfo := epochInfo
	expectedEpochInfo.StartTime = ctx.BlockTime()
	expectedEpochInfo.CurrentEpochStartHeight = ctx.BlockHeight()
	suite.Require().Equal(expectedEpochInfo, epochInfoSaved)

	allEpochs := suite.EpochsKeeper.AllEpochInfos(ctx)
	suite.Require().Len(allEpochs, 4)
	suite.Require().Equal(allEpochs[0].Identifier, "day") // alphabetical order
	suite.Require().Equal(allEpochs[1].Identifier, "hour")
	suite.Require().Equal(allEpochs[2].Identifier, "monthly")
	suite.Require().Equal(allEpochs[3].Identifier, "week")
}
