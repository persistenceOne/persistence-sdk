package base

import "github.com/persistenceOne/persistenceSDK/schema/test_types"

var _ test_types.Height = (*height)(nil)


func (height height) Get() int64 { return height.Value }
func (height height) IsGreaterThan(compareHeight test_types.Height) bool {
	return height.Get() > compareHeight.Get()
}
func (height height) Equals(compareHeight test_types.Height) bool {
	return height.Get() == compareHeight.Get()
}
func NewHeight(value int64) test_types.Height {
	return height{Value: value}
}
