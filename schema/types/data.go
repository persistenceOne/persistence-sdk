/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
)

type Data interface {
	String() string

	GetTypeID() protoTypes.ID

	ZeroValue() Data

	GenerateHashID() protoTypes.ID

	AsAccAddress() (sdkTypes.AccAddress, error)
	AsAccAddressList() ([]sdkTypes.AccAddress, error)
	AsString() (string, error)
	AsDec() (sdkTypes.Dec, error)
	AsHeight() (Height, error)
	AsID() (protoTypes.ID, error)

	Get() interface{}

	Equal(Data) bool
}
