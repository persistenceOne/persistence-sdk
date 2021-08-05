/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package define

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cryptoCodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	parKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	base2 "github.com/persistenceOne/persistenceSDK/schema/proto/types/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"
	"reflect"
	"testing"
)

type TestKeepers struct {
	ClassificationsKeeper helpers.AuxiliaryKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers) {
	var Codec = codec.NewLegacyAmino()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterLegacyAminoCodec(Codec)
	cryptoCodec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vestingTypes.RegisterLegacyAminoCodec(Codec)
	Codec.Seal()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	paramsKeeper := parKeeper.NewKeeper(
		nil,
		codec.NewLegacyAmino(),
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

	keepers := TestKeepers{
		ClassificationsKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{}).(helpers.AuxiliaryKeeper),
	}

	return context, keepers
}

func Test_Auxiliary_Keeper_Help(t *testing.T) {

	context, keepers := CreateTestInput(t)

	immutableProperties := base.NewProperties(base.NewProperty(base2.NewID("ID2"), base.NewFact(base.NewStringData("Data2"))))
	mutableProperties := base.NewProperties(base.NewProperty(base2.NewID("ID1"), base.NewFact(base.NewStringData("Data1"))))

	classificationID := key.NewClassificationID(base2.NewID(context.ChainID()), immutableProperties, mutableProperties)

	testClassificationID := key.NewClassificationID(base2.NewID(context.ChainID()), base.NewProperties(), base.NewProperties())

	keepers.ClassificationsKeeper.(auxiliaryKeeper).mapper.NewCollection(context).Add(mappable.NewClassification(testClassificationID, base.NewProperties(), base.NewProperties()))

	t.Run("PositiveCase", func(t *testing.T) {
		want := newAuxiliaryResponse(base2.NewID(classificationID.String()), nil)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(immutableProperties, mutableProperties)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Classification already present", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(base2.NewID(testClassificationID.String()), errors.EntityAlreadyExists)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(base.NewProperties(), base.NewProperties())); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Max Property Count", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(nil, errors.InvalidRequest)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(base.NewProperties(base.NewProperty(base2.NewID("ID1"), base.NewFact(base.NewStringData("Data1"))), base.NewProperty(base2.NewID("ID2"), base.NewFact(base.NewStringData("Data2"))), base.NewProperty(base2.NewID("ID3"), base.NewFact(base.NewStringData("Data3"))), base.NewProperty(base2.NewID("ID4"), base.NewFact(base.NewStringData("Data4"))), base.NewProperty(base2.NewID("ID5"), base.NewFact(base.NewStringData("Data5"))), base.NewProperty(base2.NewID("ID6"), base.NewFact(base.NewStringData("Data6"))), base.NewProperty(base2.NewID("ID7"), base.NewFact(base.NewStringData("Data7"))), base.NewProperty(base2.NewID("ID8"), base.NewFact(base.NewStringData("Data8"))), base.NewProperty(base2.NewID("ID9"), base.NewFact(base.NewStringData("Data9"))), base.NewProperty(base2.NewID("ID10"), base.NewFact(base.NewStringData("Data10"))), base.NewProperty(base2.NewID("ID9"), base.NewFact(base.NewStringData("Data9"))), base.NewProperty(base2.NewID("ID10"), base.NewFact(base.NewStringData("Data10")))), base.NewProperties(base.NewProperty(base2.NewID("ID1"), base.NewFact(base.NewStringData("Data1"))), base.NewProperty(base2.NewID("ID2"), base.NewFact(base.NewStringData("Data2"))), base.NewProperty(base2.NewID("ID3"), base.NewFact(base.NewStringData("Data3"))), base.NewProperty(base2.NewID("ID4"), base.NewFact(base.NewStringData("Data4"))), base.NewProperty(base2.NewID("ID5"), base.NewFact(base.NewStringData("Data5"))), base.NewProperty(base2.NewID("ID6"), base.NewFact(base.NewStringData("Data6"))), base.NewProperty(base2.NewID("ID7"), base.NewFact(base.NewStringData("Data7"))), base.NewProperty(base2.NewID("ID8"), base.NewFact(base.NewStringData("Data8"))), base.NewProperty(base2.NewID("ID9"), base.NewFact(base.NewStringData("Data9"))), base.NewProperty(base2.NewID("ID10"), base.NewFact(base.NewStringData("Data10"))), base.NewProperty(base2.NewID("ID9"), base.NewFact(base.NewStringData("Data9"))), base.NewProperty(base2.NewID("ID10"), base.NewFact(base.NewStringData("Data10")))))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Duplicate Immutable Property", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(nil, errors.InvalidRequest)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(base.NewProperties(base.NewProperty(base2.NewID("ID1"), base.NewFact(base.NewStringData("Data1"))), base.NewProperty(base2.NewID("ID1"), base.NewFact(base.NewStringData("Data2"))), base.NewProperty(base2.NewID("ID3"), base.NewFact(base.NewStringData("Data3"))), base.NewProperty(base2.NewID("ID4"), base.NewFact(base.NewStringData("Data4"))), base.NewProperty(base2.NewID("ID5"), base.NewFact(base.NewStringData("Data5"))), base.NewProperty(base2.NewID("ID6"), base.NewFact(base.NewStringData("Data6"))), base.NewProperty(base2.NewID("ID7"), base.NewFact(base.NewStringData("Data7"))), base.NewProperty(base2.NewID("ID8"), base.NewFact(base.NewStringData("Data8"))), base.NewProperty(base2.NewID("ID9"), base.NewFact(base.NewStringData("Data9"))), base.NewProperty(base2.NewID("ID10"), base.NewFact(base.NewStringData("Data10")))), base.NewProperties(base.NewProperty(base2.NewID("ID11"), base.NewFact(base.NewStringData("Data11"))), base.NewProperty(base2.NewID("ID12"), base.NewFact(base.NewStringData("Data12"))), base.NewProperty(base2.NewID("ID13"), base.NewFact(base.NewStringData("Data13"))), base.NewProperty(base2.NewID("ID14"), base.NewFact(base.NewStringData("Data14"))), base.NewProperty(base2.NewID("ID15"), base.NewFact(base.NewStringData("Data15"))), base.NewProperty(base2.NewID("ID16"), base.NewFact(base.NewStringData("Data16"))), base.NewProperty(base2.NewID("ID17"), base.NewFact(base.NewStringData("Data17"))), base.NewProperty(base2.NewID("ID18"), base.NewFact(base.NewStringData("Data18"))), base.NewProperty(base2.NewID("ID19"), base.NewFact(base.NewStringData("Data19"))), base.NewProperty(base2.NewID("ID20"), base.NewFact(base.NewStringData("Data20")))))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Duplicate Immutable and Mutable Property", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(nil, errors.InvalidRequest)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(base.NewProperties(base.NewProperty(base2.NewID("ID1"), base.NewFact(base.NewStringData("Data1"))), base.NewProperty(base2.NewID("ID2"), base.NewFact(base.NewStringData("Data2"))), base.NewProperty(base2.NewID("ID3"), base.NewFact(base.NewStringData("Data3"))), base.NewProperty(base2.NewID("ID4"), base.NewFact(base.NewStringData("Data4"))), base.NewProperty(base2.NewID("ID5"), base.NewFact(base.NewStringData("Data5"))), base.NewProperty(base2.NewID("ID6"), base.NewFact(base.NewStringData("Data6"))), base.NewProperty(base2.NewID("ID7"), base.NewFact(base.NewStringData("Data7"))), base.NewProperty(base2.NewID("ID8"), base.NewFact(base.NewStringData("Data8"))), base.NewProperty(base2.NewID("ID9"), base.NewFact(base.NewStringData("Data9"))), base.NewProperty(base2.NewID("ID10"), base.NewFact(base.NewStringData("Data10")))), base.NewProperties(base.NewProperty(base2.NewID("ID1"), base.NewFact(base.NewStringData("Data11"))), base.NewProperty(base2.NewID("ID12"), base.NewFact(base.NewStringData("Data12"))), base.NewProperty(base2.NewID("ID13"), base.NewFact(base.NewStringData("Data13"))), base.NewProperty(base2.NewID("ID14"), base.NewFact(base.NewStringData("Data14"))), base.NewProperty(base2.NewID("ID15"), base.NewFact(base.NewStringData("Data15"))), base.NewProperty(base2.NewID("ID16"), base.NewFact(base.NewStringData("Data16"))), base.NewProperty(base2.NewID("ID17"), base.NewFact(base.NewStringData("Data17"))), base.NewProperty(base2.NewID("ID18"), base.NewFact(base.NewStringData("Data18"))), base.NewProperty(base2.NewID("ID19"), base.NewFact(base.NewStringData("Data19"))), base.NewProperty(base2.NewID("ID20"), base.NewFact(base.NewStringData("Data20")))))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
}
