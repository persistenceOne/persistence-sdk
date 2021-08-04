/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
)

type Data interface {
	String() string

	GetTypeID() test_types.ID

	ZeroValue() Data

	GenerateHashID() test_types.ID

	AsAccAddress() (sdkTypes.AccAddress, error)
	AsAccAddressList() ([]sdkTypes.AccAddress, error)
	AsString() (string, error)
	AsDec() (sdkTypes.Dec, error)
	AsHeight() (Height, error)
	AsID() (test_types.ID, error)

	Get() interface{}

	Equal(Data) bool
}
