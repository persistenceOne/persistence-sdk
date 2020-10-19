/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type identity struct {
	ID                       types.ID              `json:"id" valid:"required~required field id missing"`
	ProvisionedAddressList   []sdkTypes.AccAddress `json:"provisionedAddressList" valid:"required~required field provisionedAddressList missing"`
	UnprovisionedAddressList []sdkTypes.AccAddress `json:"unprovisionedAddressList" valid:"required~required field unprovisionedAddressList missing"`
	Immutables               types.Immutables      `json:"immutables" valid:"required~required field immutables missing"`
	Mutables                 types.Mutables        `json:"mutables" valid:"required~required field mutables missing"`
}

var _ mappables.InterIdentity = (*identity)(nil)

func (identity identity) GetID() types.ID { return identity.ID }

func (identity identity) GetProvisionedAddressList() []sdkTypes.AccAddress {
	return identity.ProvisionedAddressList
}
func (identity identity) GetUnprovisionedAddressList() []sdkTypes.AccAddress {
	return identity.UnprovisionedAddressList
}
func (identity identity) ProvisionAddress(accAddress sdkTypes.AccAddress) mappables.InterIdentity {
	identity.ProvisionedAddressList = append(identity.ProvisionedAddressList, accAddress)
	return identity
}
func (identity identity) UnprovisionAddress(accAddress sdkTypes.AccAddress) mappables.InterIdentity {
	for i, provisionedAddress := range identity.ProvisionedAddressList {
		if provisionedAddress.Equals(accAddress) {
			identity.ProvisionedAddressList = append(identity.ProvisionedAddressList[:i], identity.ProvisionedAddressList[i+1:]...)
			identity.UnprovisionedAddressList = append(identity.UnprovisionedAddressList, accAddress)
			return identity
		}
	}
	return identity
}
func (identity identity) GetImmutables() types.Immutables { return identity.Immutables }
func (identity identity) GetMutables() types.Mutables     { return identity.Mutables }
func (identity identity) IsProvisioned(accAddress sdkTypes.AccAddress) bool {
	for _, provisionedAddress := range identity.ProvisionedAddressList {
		if provisionedAddress.Equals(accAddress) {
			return true
		}
	}
	return false
}
func (identity identity) IsUnprovisioned(accAddress sdkTypes.AccAddress) bool {
	for _, unprovisionedAddress := range identity.UnprovisionedAddressList {
		if unprovisionedAddress.Equals(accAddress) {
			return true
		}
	}
	return false
}
func (identity identity) GetKey() helpers.Key {
	return key.New(identity.ID)
}
func (identity) RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(identity{}, constants.ProjectRoute+"/"+"identity", nil)
}

func NewIdentity(identityID types.ID, provisionedAddressList []sdkTypes.AccAddress, unprovisionedAddressList []sdkTypes.AccAddress, immutables types.Immutables, mutables types.Mutables) mappables.InterIdentity {
	return identity{
		ID:                       identityID,
		ProvisionedAddressList:   provisionedAddressList,
		UnprovisionedAddressList: unprovisionedAddressList,
		Immutables:               immutables,
		Mutables:                 mutables,
	}
}