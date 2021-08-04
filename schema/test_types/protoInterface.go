package test_types

type ProtoInterface interface {

Size() int
MarshalTo([]byte) (int, error)
Unmarshal([]byte) error
MarshalToSizedBuffer([]byte) (int, error)
}
