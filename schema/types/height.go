/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import "github.com/gogo/protobuf/proto"

type Height interface {
	proto.Message
	Get() int64
	IsGreaterThan(Height) bool
	Equals(Height) bool
}
