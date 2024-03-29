package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdkstaking "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the oracle MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (ms msgServer) AggregateExchangeRatePrevote(
	goCtx context.Context,
	msg *types.MsgAggregateExchangeRatePrevote,
) (*types.MsgAggregateExchangeRatePrevoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valAddr, err := sdk.ValAddressFromBech32(msg.Validator)
	if err != nil {
		return nil, err
	}

	feederAddr, err := sdk.AccAddressFromBech32(msg.Feeder)
	if err != nil {
		return nil, err
	}

	if err := ms.ValidateFeeder(ctx, valAddr, feederAddr); err != nil {
		return nil, err
	}

	// Ensure prevote wasn't already submitted
	if ms.HasAggregateExchangeRatePrevote(ctx, valAddr) {
		return nil, types.ErrExistingPrevote
	}

	// Convert hex string to votehash
	voteHash, err := types.AggregateVoteHashFromHexString(msg.Hash)
	if err != nil {
		return nil, errors.Wrap(types.ErrInvalidHash, err.Error())
	}

	aggregatePrevote := types.NewAggregateExchangeRatePrevote(voteHash, valAddr, uint64(ctx.BlockHeight()))
	ms.SetAggregateExchangeRatePrevote(ctx, valAddr, aggregatePrevote)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAggregatePrevote,
			sdk.NewAttribute(types.EventAttrKeyVoter, msg.Validator),
		),
	})

	return &types.MsgAggregateExchangeRatePrevoteResponse{}, nil
}

func (ms msgServer) AggregateExchangeRateVote(
	goCtx context.Context,
	msg *types.MsgAggregateExchangeRateVote,
) (*types.MsgAggregateExchangeRateVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valAddr, err := sdk.ValAddressFromBech32(msg.Validator)
	if err != nil {
		return nil, err
	}

	feederAddr, err := sdk.AccAddressFromBech32(msg.Feeder)
	if err != nil {
		return nil, err
	}

	if err := ms.ValidateFeeder(ctx, valAddr, feederAddr); err != nil {
		return nil, err
	}

	params := ms.GetParams(ctx)

	aggregatePrevote, err := ms.GetAggregateExchangeRatePrevote(ctx, valAddr)
	if err != nil {
		return nil, errors.Wrap(types.ErrNoAggregatePrevote, msg.Validator)
	}

	// Check a msg is submitted proper period
	if (uint64(ctx.BlockHeight())/params.VotePeriod)-(aggregatePrevote.SubmitBlock/params.VotePeriod) != 1 {
		return nil, types.ErrRevealPeriodMissMatch
	}

	exchangeRateTuples, err := types.ParseExchangeRateTuples(msg.ExchangeRates)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidCoins, err.Error())
	}

	// Verify a exchange rate with aggregate prevote hash
	hash := types.GetAggregateVoteHash(msg.Salt, msg.ExchangeRates, valAddr)
	if aggregatePrevote.Hash != hash.String() {
		return nil, errors.Wrapf(types.ErrVerificationFailed, "must be given %s not %s", aggregatePrevote.Hash, hash)
	}

	// Filter out rates which aren't included in the AcceptList
	filteredTuples := types.ExchangeRateTuples{}

	for _, tuple := range exchangeRateTuples {
		if params.AcceptList.Contains(tuple.Denom) {
			filteredTuples = append(filteredTuples, tuple)
		}
	}

	// Move aggregate prevote to aggregate vote with given exchange rates
	ms.SetAggregateExchangeRateVote(ctx, valAddr, types.NewAggregateExchangeRateVote(filteredTuples, valAddr))
	ms.DeleteAggregateExchangeRatePrevote(ctx, valAddr)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAggregateVote,
			sdk.NewAttribute(types.EventAttrKeyVoter, msg.Validator),
			sdk.NewAttribute(types.EventAttrKeyExchangeRates, msg.ExchangeRates),
		),
	})

	return &types.MsgAggregateExchangeRateVoteResponse{}, nil
}

func (ms msgServer) DelegateFeedConsent(
	goCtx context.Context,
	msg *types.MsgDelegateFeedConsent,
) (*types.MsgDelegateFeedConsentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	operatorAddr, err := sdk.ValAddressFromBech32(msg.Operator)
	if err != nil {
		return nil, err
	}

	delegateAddr, err := sdk.AccAddressFromBech32(msg.Delegate)
	if err != nil {
		return nil, err
	}

	val := ms.StakingKeeper.Validator(ctx, operatorAddr)
	if val == nil {
		return nil, errors.Wrap(sdkstaking.ErrNoValidatorFound, msg.Operator)
	}

	ms.SetFeederDelegation(ctx, operatorAddr, delegateAddr)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeFeedDelegate,
			sdk.NewAttribute(types.EventAttrKeyFeeder, msg.Delegate),
		),
	})

	return &types.MsgDelegateFeedConsentResponse{}, nil
}

// AddFundsToRewardPool add funds to the reward pool
func (ms msgServer) AddFundsToRewardPool(
	goCtx context.Context,
	msg *types.MsgAddFundsToRewardPool) (*types.MsgAddFundsToRewardPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	if err = ms.FundRewardPool(ctx, fromAddr, msg.Funds); err != nil {
		return nil, err
	}

	return &types.MsgAddFundsToRewardPoolResponse{}, nil
}
