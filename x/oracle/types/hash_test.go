package types

import (
	"encoding/hex"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestAggregateVoteHash(t *testing.T) {
	addrs := []sdk.AccAddress{
		sdk.AccAddress([]byte("addr1_______________")),
	}

	aggregateVoteHash := GetAggregateVoteHash("salt", "XPRT:100,ATOM:100", sdk.ValAddress(addrs[0]))
	hexStr := hex.EncodeToString(aggregateVoteHash)
	aggregateVoteHashRes, err := AggregateVoteHashFromHexString(hexStr)
	require.NoError(t, err)
	require.Equal(t, true, aggregateVoteHash.Equal(aggregateVoteHashRes))
	require.Equal(t, true, AggregateVoteHash([]byte{}).Empty())

	got, _ := yaml.Marshal(&aggregateVoteHash)
	require.Equal(t, aggregateVoteHash.String()+"\n", string(got))

	res := AggregateVoteHash{}
	testMarshal(t, &aggregateVoteHash, &res, aggregateVoteHash.MarshalJSON, (&res).UnmarshalJSON)
	testMarshal(t, &aggregateVoteHash, &res, aggregateVoteHash.Marshal, (&res).Unmarshal)
}

func testMarshal(t *testing.T, original, res interface{}, marshal func() ([]byte, error), unmarshal func([]byte) error) {
	bz, err := marshal()
	require.NoError(t, err)
	err = unmarshal(bz)
	require.NoError(t, err)
	require.EqualValues(t, original, res)
}
