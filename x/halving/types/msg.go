/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceCore contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgUpdateParams = "update_params"
)

var (
	_ sdk.Msg = &MsgUpdateParams{}
)

// NewMsgUpdateParams creates a new MsgUpdateParams instance
func NewMsgUpdateParams(authority string, params Params) *MsgUpdateParams {
	return &MsgUpdateParams{
		Authority: authority,
		Params:    params,
	}
}

// Route returns the message route
func (msg *MsgUpdateParams) Route() string {
	return RouterKey
}

// Type returns the message type
func (msg *MsgUpdateParams) Type() string {
	return TypeMsgUpdateParams
}

// GetSigners returns the expected signers for a MsgUpdateParams message
func (msg *MsgUpdateParams) GetSigners() []sdk.AccAddress {
	authority, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{authority}
}

// GetSignBytes returns the sign bytes for a MsgUpdateParams message
func (msg *MsgUpdateParams) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic performs basic MsgUpdateParams message validation
func (msg *MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}

	if err := msg.Params.Validate(); err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "invalid params: %s", err)
	}

	return nil
}
