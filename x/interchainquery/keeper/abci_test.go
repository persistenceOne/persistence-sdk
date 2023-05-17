package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/persistenceOne/persistence-sdk/v2/ibctesting"
	"github.com/persistenceOne/persistence-sdk/v2/x/interchainquery/keeper"
	stakingtypes "github.com/persistenceOne/persistence-sdk/v2/x/lsnative/staking/types"
)

func (suite *KeeperTestSuite) TestEndBlocker() {
	validators := suite.GetSimApp(suite.chainB).StakingKeeper.GetBondedValidatorsByPower(suite.chainB.GetContext())

	qvr := stakingtypes.QueryValidatorsResponse{
		Validators: ibctesting.SdkValidatorsToValidators(validators),
	}

	bondedQuery := stakingtypes.QueryValidatorsRequest{Status: stakingtypes.BondStatusBonded}
	bz, err := bondedQuery.Marshal()
	suite.NoError(err)

	id := keeper.GenerateQueryHash(suite.path.EndpointB.ConnectionID, suite.chainB.ChainID, "cosmos.staking.v1beta1.Query/Validators", bz, "")

	query := suite.GetSimApp(suite.chainA).InterchainQueryKeeper.NewQuery(
		suite.chainA.GetContext(),
		"",
		suite.path.EndpointB.ConnectionID,
		suite.chainB.ChainID,
		"cosmos.staking.v1beta1.Query/Validators",
		bz,
		sdk.NewInt(200),
		"",
		0,
	)

	// set the query
	suite.GetSimApp(suite.chainA).InterchainQueryKeeper.SetQuery(suite.chainA.GetContext(), *query)

	// call end blocker
	suite.GetSimApp(suite.chainA).InterchainQueryKeeper.EndBlocker(suite.chainA.GetContext())

	err = suite.GetSimApp(suite.chainA).InterchainQueryKeeper.SetDatapointForID(
		suite.chainA.GetContext(),
		id,
		suite.GetSimApp(suite.chainB).AppCodec().MustMarshalJSON(&qvr),
		sdk.NewInt(suite.chainB.CurrentHeader.Height),
	)
	suite.NoError(err)

	dataPoint, err := suite.GetSimApp(suite.chainA).InterchainQueryKeeper.GetDatapointForID(suite.chainA.GetContext(), id)
	suite.NoError(err)
	suite.NotNil(dataPoint)

	// set the query
	suite.GetSimApp(suite.chainA).InterchainQueryKeeper.DeleteQuery(suite.chainA.GetContext(), id)

	// call end blocker
	suite.GetSimApp(suite.chainA).InterchainQueryKeeper.EndBlocker(suite.chainA.GetContext())
}
