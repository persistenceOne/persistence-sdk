/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
	testBase "github.com/persistenceOne/persistenceSDK/schema/test_types/base"

	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

type idData struct {
	Value test_types.ID `json:"value"`
}

var _ types.Data = (*idData)(nil)

func (idData idData) String() string {
	return idData.Value.String()
}
func (idData idData) ZeroValue() types.Data {
	return NewIDData(testBase.NewID(""))
}
func (idData idData) GetTypeID() test_types.ID {
	return testBase.NewID("I")
}
func (idData idData) GenerateHashID() test_types.ID {
	return testBase.NewID(meta.Hash(idData.Value.String()))
}
func (idData idData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := accAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.EntityNotFound
}
func (idData idData) AsAccAddressList() ([]sdkTypes.AccAddress, error) {
	zeroValue, _ := accAddressListData{}.ZeroValue().AsAccAddressList()
	return zeroValue, errors.IncorrectFormat
}
func (idData idData) AsString() (string, error) {
	zeroValue, _ := stringData{}.ZeroValue().AsString()
	return zeroValue, errors.IncorrectFormat
}
func (idData idData) AsDec() (sdkTypes.Dec, error) {
	zeroValue, _ := decData{}.ZeroValue().AsDec()
	return zeroValue, errors.IncorrectFormat
}
func (idData idData) AsHeight() (types.Height, error) {
	zeroValue, _ := heightData{}.ZeroValue().AsHeight()
	return zeroValue, errors.IncorrectFormat
}
func (idData idData) AsID() (test_types.ID, error) {
	return idData.Value, nil
}
func (idData idData) Get() interface{} {
	return idData.Value
}
func (idData idData) Equal(data types.Data) bool {
	compareIDData, Error := idDataFromInterface(data)
	if Error != nil {
		return false
	}

	return idData.Value.Equals(compareIDData.Value)
}
func idDataFromInterface(data types.Data) (idData, error) {
	switch value := data.(type) {
	case idData:
		return value, nil
	default:
		return idData{}, errors.MetaDataError
	}
}

func NewIDData(value test_types.ID) types.Data {
	return idData{
		Value: value,
	}
}

func ReadIDData(idData string) (types.Data, error) {
	return NewIDData(testBase.NewID(idData)), nil
}
