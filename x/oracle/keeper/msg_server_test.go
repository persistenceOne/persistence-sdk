package keeper_test

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/types"
)

// generateSalt generates a random salt, size length/2,  as a HEX encoded string.
func generateSalt(length int) (string, error) {
	if length == 0 {
		return "", fmt.Errorf("failed to generate salt: zero length")
	}

	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func (s *KeeperTestSuite) TestMsgServer_AggregateExchangeRatePrevote() {
	ctx := s.ctx
	addr, valAddr := s.accAddresses[0], s.valAddresses[0]

	exchangeRatesStr := "123.2:ATOM"
	salt, err := generateSalt(32)
	s.Require().NoError(err)

	hash := types.GetAggregateVoteHash(salt, exchangeRatesStr, valAddr)

	invalidHash := &types.MsgAggregateExchangeRatePrevote{
		Hash:      "invalid_hash",
		Feeder:    addr.String(),
		Validator: valAddr.String(),
	}

	invalidFeeder := &types.MsgAggregateExchangeRatePrevote{
		Hash:      hash.String(),
		Feeder:    "invalid_feeder",
		Validator: valAddr.String(),
	}

	invalidValidator := &types.MsgAggregateExchangeRatePrevote{
		Hash:      hash.String(),
		Feeder:    addr.String(),
		Validator: "invalid_val",
	}

	validMsg := &types.MsgAggregateExchangeRatePrevote{
		Hash:      hash.String(),
		Feeder:    addr.String(),
		Validator: valAddr.String(),
	}

	_, err = s.msgServer.AggregateExchangeRatePrevote(sdk.WrapSDKContext(ctx), invalidHash)
	s.Require().ErrorContains(err, types.ErrInvalidHash.Error())

	_, err = s.msgServer.AggregateExchangeRatePrevote(sdk.WrapSDKContext(ctx), invalidFeeder)
	s.Require().Error(err)

	_, err = s.msgServer.AggregateExchangeRatePrevote(sdk.WrapSDKContext(ctx), invalidValidator)
	s.Require().Error(err)

	_, err = s.msgServer.AggregateExchangeRatePrevote(sdk.WrapSDKContext(ctx), validMsg)
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) TestMsgServer_AggregateExchangeRateVote() {
	app, ctx := s.app, s.ctx
	addr, valAddr := s.accAddresses[0], s.valAddresses[0]

	ratesStr := "xprt:123.2"
	ratesStrInvalidCoin := "xprt:123.2,badcoin:234.5"

	salt, err := generateSalt(32)
	s.Require().NoError(err)

	hash := types.GetAggregateVoteHash(salt, ratesStr, valAddr)
	hashInvalidRate := types.GetAggregateVoteHash(salt, ratesStrInvalidCoin, valAddr)

	prevoteMsg := types.NewMsgAggregateExchangeRatePrevote(hash, addr, valAddr)
	voteMsg := types.NewMsgAggregateExchangeRateVote(salt, ratesStr, addr, valAddr)
	voteMsgInvalidRate := types.NewMsgAggregateExchangeRateVote(salt, ratesStrInvalidCoin, addr, valAddr)
	voteMsgInvalidSalt := types.NewMsgAggregateExchangeRateVote("invalid_salt", ratesStr, addr, valAddr)

	// Flattened acceptList symbols to make checks easier
	oracleParams := app.OracleKeeper.GetParams(ctx)
	acceptList := oracleParams.AcceptList
	votePeriod := oracleParams.VotePeriod

	acceptListFlat := make([]string, len(acceptList))
	for i, v := range acceptList {
		acceptListFlat[i] = v.SymbolDenom
	}

	// No existing prevote
	_, err = s.msgServer.AggregateExchangeRateVote(sdk.WrapSDKContext(ctx), voteMsg)
	s.Require().EqualError(err, sdkerrors.Wrap(types.ErrNoAggregatePrevote, valAddr.String()).Error())

	_, err = s.msgServer.AggregateExchangeRatePrevote(sdk.WrapSDKContext(ctx), prevoteMsg)
	s.Require().NoError(err)

	_, err = s.msgServer.AggregateExchangeRatePrevote(sdk.WrapSDKContext(ctx), prevoteMsg)
	s.Require().EqualError(err, types.ErrExistingPrevote.Error())

	// Reveal period mismatch
	_, err = s.msgServer.AggregateExchangeRateVote(sdk.WrapSDKContext(ctx), voteMsg)
	s.Require().EqualError(err, types.ErrRevealPeriodMissMatch.Error())

	// Valid
	app.OracleKeeper.SetAggregateExchangeRatePrevote(
		ctx,
		valAddr,
		types.NewAggregateExchangeRatePrevote(
			hash, valAddr, initialHeight-votePeriod,
		))

	_, err = s.msgServer.AggregateExchangeRateVote(sdk.WrapSDKContext(ctx), voteMsg)
	s.Require().NoError(err)

	vote, err := app.OracleKeeper.GetAggregateExchangeRateVote(ctx, valAddr)
	s.Require().Nil(err)

	for _, v := range vote.ExchangeRateTuples {
		s.Require().Contains(acceptListFlat, strings.ToLower(v.Denom))
	}

	// Valid, but with an exchange rate which isn't in AcceptList
	app.OracleKeeper.SetAggregateExchangeRatePrevote(
		ctx,
		valAddr,
		types.NewAggregateExchangeRatePrevote(
			hashInvalidRate, valAddr, initialHeight-votePeriod,
		))

	_, err = s.msgServer.AggregateExchangeRateVote(sdk.WrapSDKContext(ctx), voteMsgInvalidRate)
	s.Require().NoError(err)

	vote, err = app.OracleKeeper.GetAggregateExchangeRateVote(ctx, valAddr)
	s.Require().NoError(err)

	for _, v := range vote.ExchangeRateTuples {
		s.Require().Contains(acceptListFlat, strings.ToLower(v.Denom))
	}

	// Valid pre-vote but invalid salt
	app.OracleKeeper.SetAggregateExchangeRatePrevote(
		ctx,
		valAddr,
		types.NewAggregateExchangeRatePrevote(
			hash, valAddr, initialHeight-votePeriod,
		))

	_, err = s.msgServer.AggregateExchangeRateVote(sdk.WrapSDKContext(ctx), voteMsgInvalidSalt)
	s.Require().ErrorContains(err, types.ErrVerificationFailed.Error())
}

func (s *KeeperTestSuite) TestMsgServer_DelegateFeedConsent() {
	app, ctx := s.app, s.ctx
	valAddr := s.valAddresses[0]

	feederAddr := sdk.AccAddress([]byte("addr________________"))
	feederAcc := app.AccountKeeper.NewAccountWithAddress(ctx, feederAddr)
	app.AccountKeeper.SetAccount(ctx, feederAcc)

	_, err := s.msgServer.DelegateFeedConsent(sdk.WrapSDKContext(ctx), types.NewMsgDelegateFeedConsent(valAddr, feederAddr))
	s.Require().NoError(err)
}
