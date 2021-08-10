package types

type Error interface {
	ProtoInterface
	error
}