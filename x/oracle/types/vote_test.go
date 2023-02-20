package types

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const valAddr = "addr1_______________"

var valAddrStr = sdk.ValAddress(valAddr).String()

func TestAggregateExchangeRatePrevoteString(t *testing.T) {
	addr := sdk.ValAddress(valAddr)
	aggregateVoteHash := GetAggregateVoteHash("salt", "XPRT:100,ATOM:100", addr)
	aggregateExchangeRatePreVote := NewAggregateExchangeRatePrevote(
		aggregateVoteHash,
		addr,
		100,
	)

	expected := fmt.Sprintf("hash: f66eefc3f015d1ec70a2aa27e634810875275940\nvoter: %s\nsubmit_block: 100\n", valAddrStr)
	require.Equal(t, expected, aggregateExchangeRatePreVote.String())
}

func TestAggregateExchangeRateVoteString(t *testing.T) {
	aggregateExchangeRatePreVote := NewAggregateExchangeRateVote(
		ExchangeRateTuples{
			NewExchangeRateTuple(PersistenceDenom, sdk.OneDec()),
		},
		sdk.ValAddress(valAddr),
	)

	expected := fmt.Sprintf("exchange_rate_tuples:\n    - denom: uxprt\n      exchange_rate: \"1.000000000000000000\"\nvoter: %s\n", valAddrStr)
	require.Equal(t, expected, aggregateExchangeRatePreVote.String())
}

func TestExchangeRateTuplesString(t *testing.T) {
	exchangeRateTuple := NewExchangeRateTuple(PersistenceDenom, sdk.OneDec())
	require.Equal(t, exchangeRateTuple.String(), "denom: uxprt\nexchange_rate: \"1.000000000000000000\"\n")

	exchangeRateTuples := ExchangeRateTuples{
		exchangeRateTuple,
		NewExchangeRateTuple(AtomDenom, sdk.SmallestDec()),
	}
	require.Equal(t, "- denom: uxprt\n  exchange_rate: \"1.000000000000000000\"\n- denom: ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2\n  exchange_rate: \"0.000000000000000001\"\n", exchangeRateTuples.String())
}

func TestParseExchangeRateTuples(t *testing.T) {
	valid := "uxprt:123.0,uatom:123.123"
	_, err := ParseExchangeRateTuples(valid)
	require.NoError(t, err)

	duplicatedDenom := "uxprt:100.0,uatom:123.123,uatom:121233.123"
	_, err = ParseExchangeRateTuples(duplicatedDenom)
	require.Error(t, err)

	invalidCoins := "123.123"
	_, err = ParseExchangeRateTuples(invalidCoins)
	require.Error(t, err)

	invalidCoinsWithValid := "uxprt:123.0,123.1"
	_, err = ParseExchangeRateTuples(invalidCoinsWithValid)
	require.Error(t, err)

	zeroCoinsWithValid := "uxprt:0.0,uatom:123.1"
	_, err = ParseExchangeRateTuples(zeroCoinsWithValid)
	require.Error(t, err)

	negativeCoinsWithValid := "uxprt:-1234.5,uatom:123.1"
	_, err = ParseExchangeRateTuples(negativeCoinsWithValid)
	require.EqualError(t, err, ErrInvalidOraclePrice.Error())

	multiplePricesPerRate := "uxprt:123: uxprt:456,uusdc:789"
	_, err = ParseExchangeRateTuples(multiplePricesPerRate)
	require.Error(t, err)

	res, err := ParseExchangeRateTuples("")
	require.Nil(t, err)
	require.Nil(t, res)
}
