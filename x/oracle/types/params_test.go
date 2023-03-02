package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParamKeyTable(t *testing.T) {
	require.NotNil(t, ParamKeyTable())
}

func TestValidateVotePeriod(t *testing.T) {
	err := validateVotePeriod("invalidUint64")
	require.ErrorContains(t, err, "invalid parameter type: string")

	err = validateVotePeriod(uint64(0))
	require.ErrorContains(t, err, "vote period must be positive: 0")

	err = validateVotePeriod(uint64(10))
	require.Nil(t, err)
}

func TestValidateVoteThreshold(t *testing.T) {
	testCases := []struct {
		name      string
		threshold interface{}
		errMsg    string
	}{
		{"fail: invalid parameter", "invalidSdkType", "invalid parameter type"},
		{"fail: negative", sdk.MustNewDecFromStr("-1"), "threshold must be"},
		{"fail: zero", sdk.ZeroDec(), "threshold must be"},
		{"fail: less than 0.33", sdk.MustNewDecFromStr("0.3"), "threshold must be"},
		{"fail: equal 0.33", sdk.MustNewDecFromStr("0.33"), "threshold must be"},
		{"fail: more than 1", sdk.MustNewDecFromStr("1.1"), "threshold must be"},
		{"fail: more than 1", sdk.MustNewDecFromStr("10"), "threshold must be"},
		{"fail: max precision 2", sdk.MustNewDecFromStr("0.333"), "maximum 2 decimals"},
		{"fail: max precision 2", sdk.MustNewDecFromStr("0.401"), "maximum 2 decimals"},
		{"fail: max precision 2", sdk.MustNewDecFromStr("0.409"), "maximum 2 decimals"},
		{"fail: max precision 2", sdk.MustNewDecFromStr("0.4009"), "maximum 2 decimals"},
		{"fail: max precision 2", sdk.MustNewDecFromStr("0.999"), "maximum 2 decimals"},

		{"ok: 1", sdk.MustNewDecFromStr("1"), ""},
		{"ok: 0.34", sdk.MustNewDecFromStr("0.34"), ""},
		{"ok: 0.99", sdk.MustNewDecFromStr("0.99"), ""},
	}

	for _, tc := range testCases {
		err := ValidateVoteThreshold(tc.threshold)
		if tc.errMsg == "" {
			require.NoError(t, err, "test_case", tc.name)
		} else {
			require.ErrorContains(t, err, tc.errMsg, tc.name)
		}
	}
}

func TestValidateRewardBand(t *testing.T) {
	err := validateRewardBand("invalidSdkType")
	require.ErrorContains(t, err, "invalid parameter type: string")

	err = validateRewardBand(sdk.MustNewDecFromStr("-0.31"))
	require.ErrorContains(t, err, "reward band must be positive: -0.310000000000000000")

	err = validateRewardBand(sdk.MustNewDecFromStr("40.0"))
	require.ErrorContains(t, err, "reward band is too large: 40.000000000000000000")

	err = validateRewardBand(sdk.OneDec())
	require.Nil(t, err)
}

func TestValidateRewardDistributionWindow(t *testing.T) {
	err := validateRewardDistributionWindow("invalidUint64")
	require.ErrorContains(t, err, "invalid parameter type: string")

	err = validateRewardDistributionWindow(uint64(0))
	require.ErrorContains(t, err, "reward distribution window must be positive: 0")

	err = validateRewardDistributionWindow(uint64(10))
	require.Nil(t, err)
}

func TestValidateAcceptList(t *testing.T) {
	err := validateAcceptList("invalidUint64")
	require.ErrorContains(t, err, "invalid parameter type: string")

	err = validateAcceptList(DenomList{
		{BaseDenom: ""},
	})
	require.ErrorContains(t, err, "oracle parameter AcceptList Denom must have BaseDenom")

	err = validateAcceptList(DenomList{
		{BaseDenom: denomPersistence.BaseDenom, SymbolDenom: ""},
	})
	require.ErrorContains(t, err, "oracle parameter AcceptList Denom must have SymbolDenom")

	err = validateAcceptList(DenomList{
		{BaseDenom: denomPersistence.BaseDenom, SymbolDenom: denomPersistence.SymbolDenom},
	})
	require.Nil(t, err)
}

func TestValidateSlashFraction(t *testing.T) {
	err := validateSlashFraction("invalidSdkType")
	require.ErrorContains(t, err, "invalid parameter type: string")

	err = validateSlashFraction(sdk.MustNewDecFromStr("-0.31"))
	require.ErrorContains(t, err, "slash fraction must be positive: -0.310000000000000000")

	err = validateSlashFraction(sdk.MustNewDecFromStr("40.0"))
	require.ErrorContains(t, err, "slash fraction is too large: 40.000000000000000000")

	err = validateSlashFraction(sdk.OneDec())
	require.Nil(t, err)
}

func TestParamsEqual(t *testing.T) {
	p1 := DefaultParams()
	err := p1.Validate()
	require.NoError(t, err)

	// minus vote period
	p1.VotePeriod = 0
	err = p1.Validate()
	require.Error(t, err)

	// small vote threshold
	p2 := DefaultParams()
	p2.VoteThreshold = sdk.ZeroDec()
	err = p2.Validate()
	require.Error(t, err)

	// negative reward band
	p3 := DefaultParams()
	p3.RewardBand = sdk.NewDecWithPrec(-1, 2)
	err = p3.Validate()
	require.Error(t, err)

	// negative slash fraction
	p4 := DefaultParams()
	p4.SlashFraction = sdk.NewDec(-1)
	err = p4.Validate()
	require.Error(t, err)

	// negative min valid per window
	p5 := DefaultParams()
	p5.MinValidPerWindow = sdk.NewDec(-1)
	err = p5.Validate()
	require.Error(t, err)

	// small slash window
	p6 := DefaultParams()
	p6.SlashWindow = 0
	err = p6.Validate()
	require.Error(t, err)

	// small distribution window
	p7 := DefaultParams()
	p7.RewardDistributionWindow = 0
	err = p7.Validate()
	require.Error(t, err)

	// empty name
	p10 := DefaultParams()
	p10.AcceptList[0].BaseDenom = ""
	p10.AcceptList[0].SymbolDenom = "ATOM"
	err = p10.Validate()
	require.Error(t, err)

	// empty
	p11 := DefaultParams()
	p11.AcceptList[0].BaseDenom = "uatom"
	p11.AcceptList[0].SymbolDenom = ""
	err = p11.Validate()
	require.Error(t, err)

	p13 := DefaultParams()
	require.NotNil(t, p13.ParamSetPairs())
	require.NotNil(t, p13.String())
}
