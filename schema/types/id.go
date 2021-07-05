/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import "github.com/gogo/protobuf/proto"

type ID interface {
	proto.Message
	String() string
	Bytes() []byte
	Equals(ID) bool
}
