package v046_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/simapp"

	v046 "github.com/persistenceOne/persistence-sdk/v2/x/lsnative/staking/migrations/v046"
	"github.com/persistenceOne/persistence-sdk/v2/x/lsnative/staking/types"
)

func TestMigrateJSON(t *testing.T) {
	encodingConfig := simapp.MakeTestEncodingConfig()
	clientCtx := client.Context{}.
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithCodec(encodingConfig.Codec)

	oldState := types.DefaultGenesisState()

	newState, err := v046.MigrateJSON(*oldState)
	require.NoError(t, err)

	bz, err := clientCtx.Codec.MarshalJSON(&newState)
	require.NoError(t, err)

	// Indent the JSON bz correctly.
	var jsonObj map[string]interface{}
	err = json.Unmarshal(bz, &jsonObj)
	require.NoError(t, err)
	indentedBz, err := json.MarshalIndent(jsonObj, "", "\t")
	require.NoError(t, err)

	// Make sure about new params MinCommissionRate & ExemptionFactor.
	expected := `{
	"delegations": [],
	"exported": false,
	"last_tokenize_share_record_id": "0",
	"last_total_power": "0",
	"last_validator_powers": [],
	"params": {
		"bond_denom": "stake",
		"exemption_factor": "-1.000000000000000000",
		"historical_entries": 10000,
		"max_entries": 7,
		"max_validators": 100,
		"min_commission_rate": "0.000000000000000000",
		"unbonding_time": "1814400s"
	},
	"redelegations": [],
	"tokenize_share_records": [],
	"unbonding_delegations": [],
	"validators": []
}`

	require.Equal(t, expected, string(indentedBz))
}
