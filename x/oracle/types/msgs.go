package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	return nil
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
	return nil
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
	return nil
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
	return nil
}
