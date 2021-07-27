package test_oneof

type BaseReq interface {
	GetReqDetails() string
	GenerateHash() []byte
}
