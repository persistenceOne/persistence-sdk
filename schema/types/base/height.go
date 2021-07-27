/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import "github.com/persistenceOne/persistenceSDK/schema/types"

//type height struct {
//	Value int64 `json:"height"`
//}

var _ types.Height = &Height{}

func (height *Height) Get() int64 {
	return height.Value
}

func (height *Height) IsGreaterThan(compareHeight types.Height) bool {
	return height.Get() > compareHeight.Get()
}

func (height *Height) Equals(compareHeight types.Height) bool {
	return height.Get() == compareHeight.Get()
}

func NewHeight(value int64) Height {
	return Height{Value: value}
}

//func (height height) Get() int64 { return height.Value }
//func (height height) IsGreaterThan(compareHeight types.Height) bool {
//	return height.Get() > compareHeight.Get()
//}
//func (height height) Equals(compareHeight types.Height) bool {
//	return height.Get() == compareHeight.Get()
//}
//func NewHeight(value int64) types.Height {
//	return height{Value: value}
//}
