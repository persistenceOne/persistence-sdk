// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/assets/internal/transactions/burn/message.proto

package burn

import (
	fmt "fmt"
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

type Message struct {
	From    string                                                   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID  github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"from_i_d"`
	AssetID github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,3,opt,name=asset_i_d,json=assetID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"asset_i_d"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2c6284c4487d28c, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
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
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Message)(nil), "modules.assets.internal.transactions.burn.Message")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/assets/internal/transactions/burn/message.proto", fileDescriptor_a2c6284c4487d28c)
}

var fileDescriptor_a2c6284c4487d28c = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x2f, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e, 0xc9, 0xd6, 0xcf, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0x4f, 0x2c, 0x2e, 0x4e, 0x2d, 0x29, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d,
	0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x29, 0x4a, 0xcc, 0x2b, 0x4e, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0x2b,
	0xd6, 0x4f, 0x2a, 0x2d, 0xca, 0xd3, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xd2, 0x84, 0x6a, 0xd4, 0x83, 0x68, 0xd4, 0x83, 0x69, 0xd4, 0x43, 0xd6,
	0xa8, 0x07, 0xd2, 0x28, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0xd6, 0xa5, 0x0f, 0x62, 0x41, 0x0c,
	0x50, 0x7a, 0xc8, 0xc8, 0xc5, 0xee, 0x0b, 0x31, 0x52, 0x48, 0x88, 0x8b, 0x25, 0xad, 0x28, 0x3f,
	0x57, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0x8a, 0xe2, 0xe2, 0x00, 0xd1, 0xf1,
	0x99, 0xf1, 0x29, 0x12, 0x4c, 0x20, 0x71, 0x27, 0x87, 0x13, 0xf7, 0xe4, 0x19, 0x6e, 0xdd, 0x93,
	0xb7, 0x48, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x47, 0xf2, 0x8e, 0x7f,
	0x5e, 0x2a, 0x32, 0x37, 0xd8, 0xc5, 0x5b, 0xbf, 0x38, 0x39, 0x23, 0x35, 0x37, 0x51, 0xbf, 0xa4,
	0xb2, 0x20, 0xb5, 0x58, 0xcf, 0xd3, 0x25, 0x88, 0x0d, 0x64, 0xa2, 0xa7, 0x8b, 0x50, 0x0c, 0x17,
	0x27, 0xd8, 0xd9, 0x60, 0xc3, 0x99, 0xa9, 0x64, 0x38, 0x3b, 0xd8, 0x48, 0x4f, 0x17, 0x2b, 0x8e,
	0x8e, 0x05, 0xf2, 0x0c, 0x2f, 0x16, 0xc8, 0x33, 0x38, 0xe5, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1,
	0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70,
	0xe3, 0xb1, 0x1c, 0x43, 0x54, 0x10, 0xd1, 0xd6, 0x10, 0x1d, 0x41, 0x49, 0x6c, 0xe0, 0x80, 0x35,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x30, 0x64, 0x21, 0x4c, 0xe4, 0x01, 0x00, 0x00,
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.AssetID.Size()
		i -= size
		if _, err := m.AssetID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.FromID.Size()
		i -= size
		if _, err := m.FromID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = m.FromID.Size()
	n += 1 + l + sovMessage(uint64(l))
	l = m.AssetID.Size()
	n += 1 + l + sovMessage(uint64(l))
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
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
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AssetID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)