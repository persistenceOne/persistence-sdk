package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

var _ test_types.Data = (*decData)(nil)

func (decData decData) String() string {
	return decData.Value.String()
}
func (decData decData) GetTypeID() test_types.ID {
	return NewID("D")
}
func (decData decData) ZeroValue() test_types.Data {
	return NewDecData(sdkTypes.ZeroDec())
}
func (decData decData) GenerateHashID() test_types.ID {
	if decData.Equal(decData.ZeroValue()) {
		return NewID("")
	}

	return NewID(meta.Hash(decData.Value.String()))
}
func (decData decData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := accAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.IncorrectFormat
}
func (decData decData) AsAccAddressList() ([]sdkTypes.AccAddress, error) {
	zeroValue, _ := accAddressListData{}.ZeroValue().AsAccAddressList()
	return zeroValue, errors.IncorrectFormat
}
func (decData decData) AsString() (string, error) {
	zeroValue, _ := stringData{}.ZeroValue().AsString()
	return zeroValue, errors.IncorrectFormat
}
func (decData decData) AsDec() (sdkTypes.Dec, error) {
	return decData.Value, nil
}
func (decData decData) AsHeight() (types.Height, error) {
	zeroValue, _ := heightData{}.ZeroValue().AsHeight()
	return zeroValue, errors.IncorrectFormat
}
func (decData decData) AsID() (types.ID, error) {
	zeroValue, _ := idData{}.ZeroValue().AsID()
	return zeroValue, errors.IncorrectFormat
}
func (decData decData) Get() interface{} {
	return decData.Value
}
func (decData decData) Equal(data types.Data) bool {
	compareDecData, Error := decDataFromInterface(data)
	if Error != nil {
		return false
	}

	return decData.Value.Equal(compareDecData.Value)
}
func decDataFromInterface(data types.Data) (decData, error) {
	switch value := data.(type) {
	case decData:
		return value, nil
	default:
		return decData{}, errors.MetaDataError
	}
}

func NewDecData(value sdkTypes.Dec) types.Data {
	return decData{
		Value: value,
	}
}

func ReadDecData(dataString string) (types.Data, error) {
	if dataString == "" {
		return decData{}.ZeroValue(), nil
	}

	dec, Error := sdkTypes.NewDecFromStr(dataString)
	if Error != nil {
		return decData{}.ZeroValue(), Error
	}

	return NewDecData(dec), nil
}
