package base

import "github.com/persistenceOne/persistenceSDK/schema/proto/types"

var _ types.Error =(*txError)(nil)

func (m *txError) Error() string {
	return m.error
}
