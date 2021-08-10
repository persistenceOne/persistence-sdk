/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"bytes"
	"fmt"
	"io"
	mathBits "math/bits"
	"sort"

	"github.com/persistenceOne/persistenceSDK/constants/errors"

	"github.com/persistenceOne/persistenceSDK/schema/types"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type sortedAccAddresses []sdkTypes.AccAddress

var _ types.SortedList = (*sortedAccAddresses)(nil)


//legacy methods and functions

func (accAddresses sortedAccAddresses) Len() int {
	return len(accAddresses)
}
func (accAddresses sortedAccAddresses) Less(i, j int) bool {
	return bytes.Compare(accAddresses[i], accAddresses[j]) < 0
}
func (accAddresses sortedAccAddresses) Swap(i, j int) {
	accAddresses[i], accAddresses[j] = accAddresses[j], accAddresses[i]
}
func (accAddresses sortedAccAddresses) Sort() types.SortedList {
	sort.Sort(accAddresses)
	return accAddresses
}
func (accAddresses sortedAccAddresses) Insert(i interface{}) types.SortedList {
	accAddress := accAddressFromInterface(i)
	if accAddresses.Search(accAddress) != accAddresses.Len() {
		return accAddresses
	}

	index := sort.Search(
		accAddresses.Len(),
		func(i int) bool {
			return bytes.Compare(accAddresses[i].Bytes(), accAddress.Bytes()) < 0
		},
	)

	newAccAddresses := append(accAddresses, sdkTypes.AccAddress{})
	copy(newAccAddresses[index+1:], newAccAddresses[index:])
	newAccAddresses[index] = accAddress

	return newAccAddresses
}
func (accAddresses sortedAccAddresses) Delete(i interface{}) types.SortedList {
	accAddress := accAddressFromInterface(i)
	index := accAddresses.Search(accAddress)

	if index == accAddresses.Len() {
		return accAddresses
	}

	return append(accAddresses[:index], accAddresses[index+1:]...)
}
func (accAddresses sortedAccAddresses) Search(i interface{}) int {
	accAddress := accAddressFromInterface(i)

	return sort.Search(
		accAddresses.Len(),
		func(i int) bool {
			return bytes.Equal(accAddresses[i].Bytes(), accAddress.Bytes())
		},
	)
}
func accAddressFromInterface(i interface{}) sdkTypes.AccAddress {
	switch value := i.(type) {
	case sdkTypes.AccAddress:
		return value
	default:
		panic(errors.IncorrectFormat)
	}
}

//protoInterface methods and functions


func (accAddresses sortedAccAddresses) MarshalTo(dAtA []byte) (int, error) {
	size := accAddresses.Len()
	return accAddresses.MarshalToSizedBuffer(dAtA[:size])
}

func (accAddresses sortedAccAddresses) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSortedAccAddresses
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: sorted_acc_addresses: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: sorted_acc_addresses: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSortedAccAddresses
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSortedAccAddresses
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSortedAccAddresses
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := accAddresses[accAddresses.Len()-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSortedAccAddresses(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSortedAccAddresses
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (accAddresses sortedAccAddresses) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if accAddresses.Len() > 0 {
		for iNdEx := accAddresses.Len() - 1; iNdEx >= 0; iNdEx-- {
			{
				size := len(accAddresses[iNdEx])
				i -= size
				if _, err := accAddresses[iNdEx].Marshal(); err != nil {
					return 0, err
				}
				i = encodeVarintSortedAccAddresses(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintSortedAccAddresses(dAtA []byte, offset int, v uint64) int {
	offset -= sovSortedAccAddresses(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}


func sovSortedAccAddresses(x uint64) (n int) {
	return (mathBits.Len64(x|1) + 6) / 7
}
func sozSortedAccAddresses(x uint64) (n int) {
	return sovSortedAccAddresses(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}


func skipSortedAccAddresses(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSortedAccAddresses
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSortedAccAddresses
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSortedAccAddresses
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSortedAccAddresses
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSortedAccAddresses
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSortedAccAddresses
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}


var (
	ErrInvalidLengthSortedAccAddresses        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSortedAccAddresses          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSortedAccAddresses = fmt.Errorf("proto: unexpected end of group")
)
