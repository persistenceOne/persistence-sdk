package base

import "github.com/persistenceOne/persistenceSDK/modules/test"

var _ test.BaseReq = (*TokenRequest)(nil)

func (x *TokenRequest) GetReqDetails() string {
	panic("implement me")
}

func (x *TokenRequest) GenerateHash() []byte {
	panic("implement me")
}
