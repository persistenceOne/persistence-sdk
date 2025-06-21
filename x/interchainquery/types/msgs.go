package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// interchainquery message types
const (
	TypeMsgSubmitQueryResponse = "submitqueryresponse"
)

var _ sdk.Msg = &MsgSubmitQueryResponse{}

// Type Implements Msg.
func (msg MsgSubmitQueryResponse) Type() string { return TypeMsgSubmitQueryResponse }

// ValidateBasic Implements Msg.
func (msg MsgSubmitQueryResponse) ValidateBasic() error {
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgSubmitQueryResponse) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners Implements Msg.
func (msg MsgSubmitQueryResponse) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.FromAddress)
	return []sdk.AccAddress{fromAddress}
}
