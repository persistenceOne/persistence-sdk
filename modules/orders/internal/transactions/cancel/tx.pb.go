// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/persistence/assets/orders/transactions/cancel/v1beta1/tx.proto

package cancel

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_types "github.com/persistenceOne/persistenceSDK/schema/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type message struct {
	From    github_com_cosmos_cosmos_sdk_types.AccAddress            `protobuf:"bytes,1,opt,name=from,proto3,customtype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"from"`
	FromID  github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,2,opt,name=from_iD,json=fromID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"from_iD"`
	OrderID github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,3,opt,name=orderID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"orderID"`
}

func (m message) Reset()         { m = message{} }
func (m message) String() string { return proto.CompactTextString(&m) }
func (message) ProtoMessage()    {}
func (*message) Descriptor() ([]byte, []int) {
	return fileDescriptor_11c1b7db2f1fe2ff, []int{0}
}
func (m *message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *message) XXX_Size() int {
	return m.Size()
}
func (m *message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func init() {
	proto.RegisterType((*message)(nil), "persistence.assets.orders.transactions.cancel.v1beta1.message")
}

func init() {
	proto.RegisterFile("proto/persistence/assets/orders/transactions/cancel/v1beta1/tx.proto", fileDescriptor_11c1b7db2f1fe2ff)
}

var fileDescriptor_11c1b7db2f1fe2ff = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x93, 0x2a, 0xad, 0x66, 0xcc, 0x54, 0x1c, 0x12, 0x71, 0x10, 0x41, 0x7a, 0x47, 0x91,
	0x82, 0x38, 0xd9, 0x12, 0x87, 0xe0, 0x20, 0xd4, 0xc9, 0x2e, 0xe5, 0x7a, 0x79, 0xa6, 0xc1, 0xe6,
	0x2e, 0xdc, 0x7b, 0x15, 0xfb, 0x0d, 0x1c, 0x5d, 0xdc, 0xfb, 0x71, 0x3a, 0x76, 0x14, 0x87, 0x22,
	0xed, 0xe2, 0xc7, 0x90, 0x26, 0x51, 0x33, 0x38, 0x3a, 0xbd, 0xbb, 0xe1, 0xf7, 0xbb, 0xfb, 0xbf,
	0xbf, 0x13, 0x64, 0x46, 0x93, 0xe6, 0x19, 0x18, 0x4c, 0x90, 0x40, 0x49, 0xe0, 0x02, 0x11, 0x08,
	0xb9, 0x36, 0x11, 0x18, 0xe4, 0x64, 0x84, 0x42, 0x21, 0x29, 0xd1, 0x0a, 0xb9, 0x14, 0x4a, 0xc2,
	0x84, 0x3f, 0xb6, 0x47, 0x40, 0xa2, 0xcd, 0xe9, 0x89, 0xe5, 0xb8, 0xdb, 0xa9, 0xf0, 0xac, 0xe0,
	0x59, 0xc1, 0xb3, 0x2a, 0xcf, 0x0a, 0x9e, 0x95, 0xfc, 0xc1, 0x31, 0x8d, 0x13, 0x13, 0x0d, 0x33,
	0x61, 0x68, 0xc6, 0x8b, 0x8f, 0xc4, 0x3a, 0xd6, 0xbf, 0xa7, 0x42, 0x7f, 0xf4, 0x5a, 0x73, 0x1a,
	0x29, 0x20, 0x8a, 0x18, 0xdc, 0xd0, 0xd9, 0xbd, 0x37, 0x3a, 0x6d, 0xda, 0x87, 0xf6, 0xc9, 0x7e,
	0xaf, 0xb3, 0x58, 0xf9, 0xd6, 0xfb, 0xca, 0x6f, 0xc5, 0x09, 0x8d, 0xa7, 0x23, 0x26, 0x75, 0xca,
	0xa5, 0xc6, 0x54, 0x63, 0x39, 0x5a, 0x18, 0x3d, 0x70, 0x9a, 0x65, 0x80, 0xac, 0x2b, 0x65, 0x37,
	0x8a, 0x0c, 0x20, 0xf6, 0x73, 0x85, 0x7b, 0xe7, 0x34, 0xb6, 0x73, 0x98, 0x04, 0xcd, 0x5a, 0x6e,
	0xbb, 0x2c, 0x6d, 0xe7, 0x15, 0x5b, 0x25, 0xd9, 0x8d, 0x82, 0xea, 0xf5, 0x36, 0xb8, 0xe6, 0x28,
	0xc7, 0x90, 0x8a, 0xf2, 0x81, 0x30, 0xe8, 0xd7, 0xb7, 0xc2, 0x30, 0x70, 0x07, 0x4e, 0x23, 0xcf,
	0x1f, 0x06, 0xcd, 0x9d, 0x7f, 0x52, 0x7f, 0x0b, 0x2f, 0xf6, 0x9e, 0xe7, 0xbe, 0xf5, 0x39, 0xf7,
	0xad, 0xde, 0xd5, 0x62, 0xed, 0xd9, 0xcb, 0xb5, 0x67, 0x7f, 0xac, 0x3d, 0xfb, 0x65, 0xe3, 0x59,
	0xcb, 0x8d, 0x67, 0xbd, 0x6d, 0x3c, 0x6b, 0x70, 0x9a, 0xea, 0x68, 0x3a, 0x81, 0x9f, 0x16, 0x13,
	0x45, 0x60, 0x94, 0x98, 0xfc, 0x55, 0xe7, 0xa8, 0x9e, 0x6f, 0xf9, 0xec, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0x8b, 0xf0, 0x46, 0x46, 0x0c, 0x02, 0x00, 0x00,
}

func (m *message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.OrderID.Size()
		i -= size
		if _, err := m.OrderID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.FromID.Size()
		i -= size
		if _, err := m.FromID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.FromID.Size()
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.OrderID.Size()
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = github_com_cosmos_cosmos_sdk_types.AccAddress(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FromID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OrderID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)