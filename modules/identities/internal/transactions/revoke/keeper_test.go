/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package revoke

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptoCodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/persistenceOne/persistenceSDK/schema/applications"
	testBase "github.com/persistenceOne/persistenceSDK/schema/test_types/base"
	tendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	"reflect"
	"testing"

	"github.com/persistenceOne/persistenceSDK/constants/test"

	tendermintDB "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/conform"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
)

type TestKeepers struct {
	MaintainersKeeper helpers.TransactionKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers) {

	var Codec = codec.NewLegacyAmino()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterLegacyAminoCodec(Codec)
	cryptoCodec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vestingTypes.RegisterLegacyAminoCodec(Codec)
	Codec.Seal()

	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)
	std.RegisterLegacyAminoCodec(Codec)
	std.RegisterInterfaces(interfaceRegistry)

	encodingConfig := applications.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             Codec,
	}

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	paramsKeeper := paramsKeeper.NewKeeper(
		encodingConfig.Marshaler,
		encodingConfig.Amino,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	Parameters := parameters.Prototype().Initialize(paramsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	Error := commitMultiStore.LoadLatestVersion()
	require.Nil(t, Error)

	context := sdkTypes.NewContext(commitMultiStore, tendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	conformAuxiliary := conform.AuxiliaryMock.Initialize(Mapper, Parameters)
	verifyAuxiliary := verify.AuxiliaryMock.Initialize(Mapper, Parameters)
	keepers := TestKeepers{
		MaintainersKeeper: keeperPrototype().Initialize(Mapper, Parameters,
			[]interface{}{verifyAuxiliary, conformAuxiliary}).(helpers.TransactionKeeper),
	}

	return context, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {

	context, keepers := CreateTestInput(t)
	_, Error := base.ReadProperties("maintainedProperties:S|maintainedProperties")
	require.Equal(t, nil, Error)
	_, Error = base.ReadProperties("conformError:S|mockError")
	require.Equal(t, nil, Error)
	defaultAddr := sdkTypes.AccAddress("addr")
	verifyMockErrorAddress := sdkTypes.AccAddress("verifyError")
	defaultIdentityID :=testBase.NewID("fromIdentityID")
	fakeFromID := testBase.NewID("fakeFromID")
	toID := testBase.NewID("toID")
	toID2 := testBase.NewID("toID2")
	classificationID := testBase.NewID("ClassificationID")

	t.Run("PositiveCase", func(t *testing.T) {
		want := newTransactionResponse(nil)
		require.Panics(t, func() {
			if got := keepers.MaintainersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, toID, classificationID)); !reflect.DeepEqual(got, want) {
				t.Errorf("Transact() = %v, want %v", got, want)
			}
		})
	})

	t.Run("NegativeCase - non-maintainer adding maintainer", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.NotAuthorized)
		require.Panics(t, func() {
			if got := keepers.MaintainersKeeper.Transact(context, newMessage(defaultAddr, fakeFromID, toID, classificationID)); !reflect.DeepEqual(got, want) {
				t.Errorf("Transact() = %v, want %v", got, want)
			}
		})
	})

	t.Run("NegativeCase - verify identity fail", func(t *testing.T) {
		want := newTransactionResponse(test.MockError)
		if got := keepers.MaintainersKeeper.Transact(context, newMessage(verifyMockErrorAddress, defaultIdentityID, toID, classificationID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
	t.Run("NegativeCase - ReAdd same maintainer", func(t *testing.T) {
		want := newTransactionResponse(errors.EntityAlreadyExists)
		require.Panics(t, func() {
			if got := keepers.MaintainersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, toID, classificationID)); !reflect.DeepEqual(got, want) {
				t.Errorf("Transact() = %v, want %v", got, want)
			}
		})
	})
	t.Run("NegativeCase - conform mock error", func(t *testing.T) {
		want := newTransactionResponse(test.MockError)
		require.Panics(t, func() {
			if got := keepers.MaintainersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, toID2, classificationID)); !reflect.DeepEqual(got, want) {
				t.Errorf("Transact() = %v, want %v", got, want)
			}
		})
	})
}
