package base

import (
	"bytes"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
	"sort"

	"github.com/persistenceOne/persistenceSDK/constants/errors"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type SortedAccAddresses []sdkTypes.AccAddress

var _ test_types.SortedList = (*SortedAccAddresses)(nil)

func (accAddresses SortedAccAddresses) Len() int {
	return len(accAddresses)
}
func (accAddresses SortedAccAddresses) Less(i, j int) bool {
	return bytes.Compare(accAddresses[i], accAddresses[j]) < 0
}
func (accAddresses SortedAccAddresses) Swap(i, j int) {
	accAddresses[i], accAddresses[j] = accAddresses[j], accAddresses[i]
}
func (accAddresses SortedAccAddresses) Sort() test_types.SortedList {
	sort.Sort(accAddresses)
	return accAddresses
}
func (accAddresses SortedAccAddresses) Insert(i interface{}) test_types.SortedList {
	accAddress := accAddressFromInterface(i)
	if accAddresses.Search(accAddress) != accAddresses.Len() {
		return accAddresses
	}

	index := sort.Search(
		accAddresses.Len(),
		func(i int) bool {
			return bytes.Compare(accAddresses[i].Bytes(), accAddress.Bytes()) < 0
		},
	)

	newAccAddresses := append(accAddresses, sdkTypes.AccAddress{})
	copy(newAccAddresses[index+1:], newAccAddresses[index:])
	newAccAddresses[index] = accAddress

	return newAccAddresses
}
func (accAddresses SortedAccAddresses) Delete(i interface{}) test_types.SortedList {
	accAddress := accAddressFromInterface(i)
	index := accAddresses.Search(accAddress)

	if index == accAddresses.Len() {
		return accAddresses
	}

	return append(accAddresses[:index], accAddresses[index+1:]...)
}
func (accAddresses SortedAccAddresses) Search(i interface{}) int {
	accAddress := accAddressFromInterface(i)

	return sort.Search(
		accAddresses.Len(),
		func(i int) bool {
			return bytes.Equal(accAddresses[i].Bytes(), accAddress.Bytes())
		},
	)
}

func (accAddresses SortedAccAddresses) MarshalTo(i []byte) (interface{}, interface{}) {
	
}

func accAddressFromInterface(i interface{}) sdkTypes.AccAddress {
	switch value := i.(type) {
	case sdkTypes.AccAddress:
		return value
	default:
		panic(errors.IncorrectFormat)
	}
}

