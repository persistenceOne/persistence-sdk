package types_test

import (
	"testing"

	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestDenomString(t *testing.T) {
	testCases := []struct {
		denom       types.Denom
		expectedStr string
	}{
		{
			denom:       types.DenomPersistence,
			expectedStr: "base_denom: uxprt\nsymbol_denom: XPRT\nexponent: 6\n",
		},
		{
			denom:       types.DenomAtom,
			expectedStr: "base_denom: ibc/4A17832B26BF318D052563EFFE677C1DE11DF8CE104F00204860F3E3439818B2\nsymbol_denom: ATOM\nexponent: 6\n",
		},
	}

	for _, testCase := range testCases {
		require.Equal(t, testCase.expectedStr, testCase.denom.String())
	}
}

func TestDenomEqual(t *testing.T) {
	testCases := []struct {
		denom         types.Denom
		denomCompared types.Denom
		equal         bool
	}{
		{
			denom:         types.DenomPersistence,
			denomCompared: types.DenomPersistence,
			equal:         true,
		},
		{
			denom:         types.DenomPersistence,
			denomCompared: types.DenomAtom,
			equal:         false,
		},
		{
			denom:         types.DenomAtom,
			denomCompared: types.DenomAtom,
			equal:         true,
		},
		{
			denom:         types.DenomAtom,
			denomCompared: types.DenomPersistence,
			equal:         false,
		},
	}

	for _, testCase := range testCases {
		require.Equal(t, testCase.equal, testCase.denom.Equal(&testCase.denomCompared))
	}
}

func TestDenomListString(t *testing.T) {
	testCases := []struct {
		denomList   types.DenomList
		expectedStr string
	}{
		{
			denomList:   types.DenomList{types.DenomPersistence},
			expectedStr: "base_denom: uxprt\nsymbol_denom: XPRT\nexponent: 6",
		},
		{
			denomList:   types.DenomList{types.DenomPersistence, types.DenomAtom},
			expectedStr: "base_denom: uxprt\nsymbol_denom: XPRT\nexponent: 6\n\nbase_denom: ibc/4A17832B26BF318D052563EFFE677C1DE11DF8CE104F00204860F3E3439818B2\nsymbol_denom: ATOM\nexponent: 6",
		},
	}

	for _, testCase := range testCases {
		require.Equal(t, testCase.expectedStr, testCase.denomList.String())
	}
}

func TestDenomListContains(t *testing.T) {
	testCases := []struct {
		denomList    types.DenomList
		denomSymbol  string
		symbolInList bool
	}{
		{
			denomList:    types.DenomList{types.DenomPersistence},
			denomSymbol:  types.DenomPersistence.SymbolDenom,
			symbolInList: true,
		},
		{
			denomList:    types.DenomList{types.DenomPersistence},
			denomSymbol:  types.DenomAtom.SymbolDenom,
			symbolInList: false,
		},
		{
			denomList:    types.DenomList{types.DenomPersistence, types.DenomAtom},
			denomSymbol:  types.DenomAtom.SymbolDenom,
			symbolInList: true,
		},
	}

	for _, testCase := range testCases {
		require.Equal(t, testCase.symbolInList, testCase.denomList.Contains(testCase.denomSymbol))
	}
}
