package types

type Height interface {
	Get() int64
	IsGreaterThan(Height) bool
	Equals(Height) bool
	MarshalTo([]byte) (int, error)
	Unmarshal([]byte) error
	MarshalToSizedBuffer([]byte) (int, error)
	Size() int

}
