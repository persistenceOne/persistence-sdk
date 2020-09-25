package define

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/persistenceOne/persistenceSDK/modules/classifications"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/define"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/super"
	"github.com/persistenceOne/persistenceSDK/modules/metas"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
	"reflect"
	"testing"
	"time"
)

func MakeTestCodec() *codec.Codec {
	var cdc = codec.New()
	params.RegisterCodec(cdc)
	return cdc
}

type TestKeepers struct {
	IdentitiesKeeper helpers.TransactionKeeper
}

func CreateTestInput(t *testing.T) (types.Context, TestKeepers) {

	keyParams := types.NewKVStoreKey(params.StoreKey)
	tkeyParams := types.NewTransientStoreKey(params.TStoreKey)

	keyIdentity := mapper.Mapper.GetKVStoreKey()
	keyMeta := metas.Module.GetKVStoreKey()
	keyClassification := classifications.Module.GetKVStoreKey()

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyIdentity, types.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyMeta, types.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyClassification, types.StoreTypeIAVL, db)
	err := ms.LoadLatestVersion()
	require.Nil(t, err)

	ctx := types.NewContext(ms, abci.Header{
		Height: 2,
		Time:   time.Date(2020, time.April, 22, 12, 0, 0, 0, time.UTC),
	}, false, log.NewNopLogger())

	paramsKeeper := params.NewKeeper(MakeTestCodec(), keyParams, tkeyParams)

	metasModule := metas.Module.Initialize(paramsKeeper.Subspace(metas.Module.GetDefaultParamspace()))
	classificationsModule := classifications.Module.Initialize(
		paramsKeeper.Subspace(classifications.Module.GetDefaultParamspace()),
		metasModule.GetAuxiliary(scrub.AuxiliaryName),
	)
	maintainersModule := maintainers.Module.Initialize(paramsKeeper.Subspace(maintainers.Module.GetDefaultParamspace()))

	keepers := TestKeepers{
		IdentitiesKeeper: initializeTransactionKeeper(mapper.Mapper,
			[]interface{}{metasModule.GetAuxiliary(scrub.AuxiliaryName),
				classificationsModule.GetAuxiliary(define.AuxiliaryName),
				maintainersModule.GetAuxiliary(super.AuxiliaryName),
			}),
	}
	//
	//mapper.NewIdentities(mapper.Mapper, ctx).Add(mapper.NewIdentity(base.NewID("addr"),
	//	[]types.AccAddress{types.AccAddress("addr")}, []types.AccAddress{}, )
	return ctx, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {
	ctx, keepers := CreateTestInput(t)

	type fields struct {
		keeper helpers.TransactionKeeper
	}
	type args struct {
		context types.Context
		msg     types.Msg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.TransactionResponse
	}{
		// TODO: Add test cases.
		{
			name:   "Expected case",
			fields: fields{keeper: keepers.IdentitiesKeeper},
			args: args{context: ctx, msg: message{
				From:                types.AccAddress("addr"),
				FromID:              nil,
				ImmutableMetaTraits: nil,
				ImmutableTraits:     nil,
				MutableMetaTraits:   nil,
				MutableTraits:       nil,
			}},
			want: transactionResponse{
				Success: true,
				Error:   nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.keeper.Transact(tt.args.context, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transact() = %v, want %v", got, tt.want)
			}
		})
	}
}
