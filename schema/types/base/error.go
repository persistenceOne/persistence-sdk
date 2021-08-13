package base

import "github.com/persistenceOne/persistenceSDK/schema/types"

var _ types.Error = (*TxError)(nil)

func (m TxError) Error() string {
	panic("implement me")
}
