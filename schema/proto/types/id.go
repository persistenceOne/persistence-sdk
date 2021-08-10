package types

type ID interface {
	String() string
	Bytes() []byte
	Equals(compareID ID) bool
	ProtoInterface

}


