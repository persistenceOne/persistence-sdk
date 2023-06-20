package types

import (
	"cosmossdk.io/errors"
	"github.com/cometbft/cometbft/crypto/tmhash"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgDelegateFeedConsent{}
	_ sdk.Msg = &MsgAggregateExchangeRatePrevote{}
	_ sdk.Msg = &MsgAggregateExchangeRateVote{}
)

// Messages types constants
const (
	TypeMsgDelegateFeedConsent          = "delegate_feeder"
	TypeMsgAggregateExchangeRatePrevote = "aggregate_exchange_rate_prevote"
	TypeMsgAggregateExchangeRateVote    = "aggregate_exchange_rate_vote"
	TypeMsgAddFundsToRewardPool         = "add_funds_to_reward_pool"
)

func NewMsgAggregateExchangeRatePrevote(
	hash AggregateVoteHash,
	feeder sdk.AccAddress,
	validator sdk.ValAddress,
) *MsgAggregateExchangeRatePrevote {
	return &MsgAggregateExchangeRatePrevote{
		Hash:      hash.String(),
		Feeder:    feeder.String(),
		Validator: validator.String(),
	}
}

// Type implements sdk.Msg
func (msg MsgAggregateExchangeRatePrevote) Type() string { return TypeMsgAggregateExchangeRatePrevote }

// GetSignBytes implements sdk.Msg
func (msg MsgAggregateExchangeRatePrevote) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgAggregateExchangeRatePrevote) GetSigners() []sdk.AccAddress {
	feeder, _ := sdk.AccAddressFromBech32(msg.Feeder)
	return []sdk.AccAddress{feeder}
}

// ValidateBasic Implements sdk.Msg
func (msg MsgAggregateExchangeRatePrevote) ValidateBasic() error {
	_, err := AggregateVoteHashFromHexString(msg.Hash)
	if err != nil {
		return errors.Wrapf(ErrInvalidHash, "invalid vote hash (%s)", err)
	}

	// HEX encoding doubles the hash length
	if len(msg.Hash) != tmhash.TruncatedSize*2 {
		return ErrInvalidHashLength
	}

	_, err = sdk.AccAddressFromBech32(msg.Feeder)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid feeder address (%s)", err)
	}

	_, err = sdk.ValAddressFromBech32(msg.Validator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid operator address (%s)", err)
	}

	return nil
}

func NewMsgAggregateExchangeRateVote(
	salt string,
	exchangeRates string,
	feeder sdk.AccAddress,
	validator sdk.ValAddress,
) *MsgAggregateExchangeRateVote {
	return &MsgAggregateExchangeRateVote{
		Salt:          salt,
		ExchangeRates: exchangeRates,
		Feeder:        feeder.String(),
		Validator:     validator.String(),
	}
}

// Type implements sdk.Msg
func (msg MsgAggregateExchangeRateVote) Type() string { return TypeMsgAggregateExchangeRateVote }

// GetSignBytes implements sdk.Msg
func (msg MsgAggregateExchangeRateVote) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgAggregateExchangeRateVote) GetSigners() []sdk.AccAddress {
	feeder, _ := sdk.AccAddressFromBech32(msg.Feeder)
	return []sdk.AccAddress{feeder}
}

// ValidateBasic implements sdk.Msg
func (msg MsgAggregateExchangeRateVote) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Feeder)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid feeder address (%s)", err)
	}

	_, err = sdk.ValAddressFromBech32(msg.Validator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid operator address (%s)", err)
	}

	if l := len(msg.ExchangeRates); l == 0 {
		return errors.Wrap(sdkerrors.ErrUnknownRequest, "must provide at least one oracle exchange rate")
	} else if l > 4096 {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "exchange rates string can not exceed 4096 characters")
	}

	exchangeRates, err := ParseExchangeRateTuples(msg.ExchangeRates)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, "failed to parse exchange rates string cause: "+err.Error())
	}

	for _, exchangeRate := range exchangeRates {
		// check overflow bit length
		if exchangeRate.ExchangeRate.BigInt().BitLen() > 255+sdk.DecimalPrecisionBits {
			return errors.Wrap(ErrInvalidExchangeRate, "overflow")
		}
	}

	if len(msg.Salt) != 64 {
		return ErrInvalidSaltLength
	}

	_, err = AggregateVoteHashFromHexString(msg.Salt)
	if err != nil {
		return errors.Wrap(ErrInvalidSaltFormat, "salt must be a valid hex string")
	}

	return nil
}

// NewMsgDelegateFeedConsent creates a MsgDelegateFeedConsent instance
func NewMsgDelegateFeedConsent(operatorAddress sdk.ValAddress, feederAddress sdk.AccAddress) *MsgDelegateFeedConsent {
	return &MsgDelegateFeedConsent{
		Operator: operatorAddress.String(),
		Delegate: feederAddress.String(),
	}
}

// Type implements sdk.Msg
func (msg MsgDelegateFeedConsent) Type() string { return TypeMsgDelegateFeedConsent }

// GetSignBytes implements sdk.Msg
func (msg MsgDelegateFeedConsent) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgDelegateFeedConsent) GetSigners() []sdk.AccAddress {
	operator, _ := sdk.ValAddressFromBech32(msg.Operator)
	return []sdk.AccAddress{sdk.AccAddress(operator)}
}

// ValidateBasic implements sdk.Msg
func (msg MsgDelegateFeedConsent) ValidateBasic() error {
	_, err := sdk.ValAddressFromBech32(msg.Operator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid operator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Delegate)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid delegate address (%s)", err)
	}

	return nil
}

func NewMsgAddFundsToRewardPool(from sdk.AccAddress, funds sdk.Coins) *MsgAddFundsToRewardPool {
	return &MsgAddFundsToRewardPool{
		From:  from.String(),
		Funds: funds,
	}
}

// Type implements sdk.Msg
func (m MsgAddFundsToRewardPool) Type() string { return TypeMsgAddFundsToRewardPool }

// GetSignBytes implements sdk.Msg
func (m MsgAddFundsToRewardPool) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners implements sdk.Msg
func (m MsgAddFundsToRewardPool) GetSigners() []sdk.AccAddress {
	operator, _ := sdk.ValAddressFromBech32(m.From)
	return []sdk.AccAddress{sdk.AccAddress(operator)}
}

// ValidateBasic implements sdk.Msg
func (m MsgAddFundsToRewardPool) ValidateBasic() error {
	_, err := sdk.ValAddressFromBech32(m.From)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid operator address (%s)", err)
	}

	if coins := sdk.NewCoins(m.Funds...); coins.Empty() {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, err.Error())
	}

	return nil
}
