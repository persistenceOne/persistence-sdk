package test

type BaseReq interface {
	GetReqDetails() string
	GenerateHash() []byte

}
