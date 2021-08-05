package types

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type Data interface {
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

	ProtoInterface
}
