package test_types

import (
	"bytes"
)

var _  = (*ID)(nil)

func (id ID) String() string {
return id.IdString

}
func (id ID) Bytes() []byte {
return []byte(id.IdString)
}
func (id ID) Equals(compareID ID) bool {
return bytes.Equal(id.Bytes(), compareID.Bytes())
}

func NewID(idString string) ID {
return ID{IdString: idString}
}