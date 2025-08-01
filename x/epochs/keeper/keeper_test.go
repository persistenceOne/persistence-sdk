package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/persistenceOne/persistence-sdk/v4/simapp"
	"github.com/persistenceOne/persistence-sdk/v4/x/epochs/types"
)

type KeeperTestSuite struct {
	simapp.KeeperTestHelper
	queryClient types.QueryClient
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	suite.queryClient = types.NewQueryClient(suite.QueryHelper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
