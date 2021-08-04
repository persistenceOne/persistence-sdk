package test_types


type Signatures interface {
	Get(ID) Signature

	GetList() []Signature

	Add(Signature) Signatures
	Remove(Signature) Signatures
	Mutate(Signature) Signatures

	MarshalTo([]byte) (int, error)
	Unmarshal([]byte) error
	MarshalToSizedBuffer([]byte) (int, error)
	Size() int
}
