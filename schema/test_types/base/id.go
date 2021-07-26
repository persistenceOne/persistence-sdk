package base

import (
	"bytes"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
)

var _ test_types.ID = (*ID)(nil)

func (id ID) String() string {
	return id.IdString

}
func (id ID) Bytes() []byte {
	return []byte(id.IdString)
}
func (id ID) Equals(compareID test_types.ID) bool {
	return bytes.Equal(id.Bytes(), compareID.Bytes())
}

func NewID(idString string) test_types.ID {
	return ID{IdString: idString}
}
