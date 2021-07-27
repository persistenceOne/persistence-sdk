/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

type Data interface {
	proto.Message
	String() string

	GetTypeID() ID

	ZeroValue() Data

	GenerateHashID() ID

	AsAccAddress() (sdkTypes.AccAddress, error)
	AsAccAddressList() ([]sdkTypes.AccAddress, error)
	AsString() (string, error)
	AsDec() (sdkTypes.Dec, error)
	AsHeight() (Height, error)
	AsID() (ID, error)

	Get() interface{}

	Equal(Data) bool
}
