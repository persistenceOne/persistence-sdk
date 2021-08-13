package types

type Error interface {
	Size() int
	MarshalTo([]byte) (int, error)
	Unmarshal([]byte) error
	MarshalToSizedBuffer([]byte) (int, error)
	error
}
