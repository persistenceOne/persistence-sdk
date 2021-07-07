package test_types

import (
	"bytes"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ types.ID = (*ID)(nil)

func (id ID) String() string {
return id.IdString

}
func (id ID) Bytes() []byte {
return []byte(id.IdString)
}
func (id ID) Equals(compareID types.ID) bool {
return bytes.Equal(id.Bytes(), compareID.Bytes())
}

func NewID(idString string) types.ID {
return ID{IdString: idString}
}