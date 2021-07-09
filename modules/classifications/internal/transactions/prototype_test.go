package transactions

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/test_helpers"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRegister(t *testing.T) {
	registry := types.NewInterfaceRegistry()
	registry.RegisterInterface("Transactions", (*helpers.Transactions)(nil))
	require.NotPanics(t, func() {
		registry.RegisterImplementations((*helpers.Transactions)(nil), &test_helpers.Transactions{})
	})
}

func TestPrototype(t *testing.T) {
	registry := test_helpers.NewTestInterfaceRegistry()
	anyarr := make([]*types.Any, 2)
	//cli := &base.CliCommand{Use: "NoUse", Short:"No short"}
	//spot1 , err := types.NewAnyWithValue(cli)
	spot2, err := types.NewAnyWithValue(&test_helpers.Transaction{Name: "Arham"})
	fmt.Println(spot2)
	spot3, err := types.NewAnyWithValue(&test_helpers.Transaction{Name: "Chordia"})
	//fmt.Println(spot1)
	require.NoError(t, err)
	anyarr[0] = spot2
	anyarr[1] = spot3
	hasAny := test_helpers.Transactions{
		TransactionList: anyarr,
	}
	//fmt.Println(hasAny.GetList()[0].GetName())
	//fmt.Println(hasAny.GetList()[1].GetName())
	bz, err := hasAny.Marshal()
	//fmt.Println(hasAny)
	require.NoError(t, err)

	var hasAny2 test_helpers.Transactions
	err = hasAny2.Unmarshal(bz)
	//fmt.Println(hasAny2)
	require.NoError(t, err)

	//err = types.UnpackInterfaces(hasAny2, registry)
	//fmt.Println(hasAny2)
	err = hasAny2.UnpackInterfaces(registry)
	require.NoError(t, err)

	fmt.Println(hasAny2.TransactionList[0].GetCachedValue())
	fmt.Println(hasAny2.TransactionList[1].GetCachedValue())
	//require.NoError(t, err)
	//require.Equal(t, hasAny, hasAny2)
	require.Equal(t, Prototype(), test_helpers.NewTransactions())
}
