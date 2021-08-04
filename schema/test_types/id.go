package test_types

type ID interface {
	String() string
	Bytes() []byte
	Equals(compareID ID) bool
	Size() int
	MarshalTo([]byte) (int, error)
	Unmarshal([]byte) error
	MarshalToSizedBuffer([]byte) (int, error)

}


