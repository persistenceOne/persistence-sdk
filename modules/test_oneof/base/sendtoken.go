package base

import "github.com/persistenceOne/persistenceSDK/modules/test_oneof"

var _ test_oneof.BaseReq = (*TokenRequest)(nil)

func (x *TokenRequest) GetReqDetails() string {
	return "Hello"
}

func (x *TokenRequest) GenerateHash() []byte {

	panic("implement me")
}
