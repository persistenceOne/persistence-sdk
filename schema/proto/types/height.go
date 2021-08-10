package types

type Height interface {
	Get() int64
	IsGreaterThan(Height) bool
	Equals(Height) bool
	ProtoInterface

}
