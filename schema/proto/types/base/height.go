package base

import (
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
)

var _ protoTypes.Height = (*height)(nil)


func (height height) Get() int64 { return height.Value }
func (height height) IsGreaterThan(compareHeight protoTypes.Height) bool {
	return height.Get() > compareHeight.Get()
}
func (height height) Equals(compareHeight protoTypes.Height) bool {
	return height.Get() == compareHeight.Get()
}
func NewHeight(value int64) protoTypes.Height {
	return height{Value: value}
}
