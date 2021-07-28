/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

type Property interface {
	GetID() ID
	GetFact() Fact

	//New Addition
	Size() int
	MarshalTo([]byte) (int, error)
	Unmarshal([]byte) error
}
