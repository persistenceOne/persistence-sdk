package base

import (
	"bytes"
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
)

var _ protoTypes.ID = (*ID)(nil)

func (id ID) String() string {
	return id.IdString

}
func (id ID) Bytes() []byte {
	return []byte(id.IdString)
}
func (id ID) Equals(compareID protoTypes.ID) bool {
	return bytes.Equal(id.Bytes(), compareID.Bytes())
}

func NewID(idString string) protoTypes.ID {
	return ID{IdString: idString}
}
