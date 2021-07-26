/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

type Height interface {
	Get() int64
	IsGreaterThan(Height) bool
	Equals(Height) bool

	//New Addition
	Size() int
	MarshalTo([]byte) (int, error)
	Unmarshal([]byte) error
}
