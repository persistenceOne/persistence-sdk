package cli_test

import (
	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/suite"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"

	"github.com/persistenceOne/persistence-sdk/v4/simapp"
	"github.com/persistenceOne/persistence-sdk/v4/x/epochs/client/cli"
	"github.com/persistenceOne/persistence-sdk/v4/x/epochs/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	s.cfg = network.DefaultConfig(simapp.NewTestNetworkFixture)

	var err error
	s.network, err = network.New(s.T(), "", s.cfg)
	s.Require().NoError(err)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestGetCmdCurrentEpoch() {
	val := s.network.Validators[0]

	testCases := []struct {
		name       string
		identifier string
		expectErr  bool
		respType   proto.Message
	}{
		{
			"query weekly epoch number",
			"weekly",
			false, &types.QueryCurrentEpochResponse{},
		},
		{
			"query unavailable epoch number",
			"unavailable",
			false, &types.QueryCurrentEpochResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdCurrentEpoch()
			clientCtx := val.ClientCtx

			args := []string{
				tc.identifier,
			}

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetCmdEpochsInfos() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		expectErr bool
		respType  proto.Message
	}{
		{
			"query epoch infos",
			false, &types.QueryEpochsInfoResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdCurrentEpoch()
			clientCtx := val.ClientCtx

			args := []string{}

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}
