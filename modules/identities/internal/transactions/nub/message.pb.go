// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/identities/internal/transactions/nub/message.proto

package nub

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_types "github.com/persistenceOne/persistenceSDK/schema/types"
	github_com_persistenceOne_persistenceSDK_schema_types_base "github.com/persistenceOne/persistenceSDK/schema/types/base"
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
	From  github_com_persistenceOne_persistenceSDK_schema_types_base.AccAddress `protobuf:"bytes,1,opt,name=from,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types/base.AccAddress" json:"from"`
	NubID github_com_persistenceOne_persistenceSDK_schema_types.ID              `protobuf:"bytes,2,opt,name=nub_i_d,json=nubID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"nub_i_d"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a91bc736e1ead03, []int{0}
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
	proto.RegisterType((*Message)(nil), "modules.identities.internal.transactions.nub.Message")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/identities/internal/transactions/nub/message.proto", fileDescriptor_4a91bc736e1ead03)
}

var fileDescriptor_4a91bc736e1ead03 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xbd, 0x4a, 0x73, 0x41,
	0x10, 0x86, 0x77, 0x3f, 0x3e, 0x8d, 0x9e, 0x32, 0x58, 0x04, 0x8b, 0x3d, 0x62, 0x65, 0x21, 0xbb,
	0x85, 0x8d, 0x58, 0x99, 0x10, 0x8b, 0x28, 0x41, 0x88, 0x08, 0x62, 0x13, 0xf6, 0x67, 0x3c, 0x59,
	0xcc, 0xd9, 0x0d, 0x3b, 0xbb, 0x85, 0x77, 0x60, 0xe9, 0x25, 0xe4, 0x72, 0x52, 0xa6, 0x14, 0x8b,
	0x20, 0x49, 0xe3, 0x65, 0x48, 0x4e, 0x22, 0xa6, 0x54, 0xbb, 0x77, 0x8a, 0xe7, 0x7d, 0x86, 0x99,
	0xec, 0x72, 0x04, 0x01, 0x2d, 0x46, 0x70, 0x1a, 0xfa, 0x68, 0x1e, 0x45, 0xe9, 0x4d, 0x1a, 0x02,
	0x0a, 0x6b, 0xc0, 0x45, 0x1b, 0xed, 0x32, 0xba, 0x08, 0xc1, 0xc9, 0xa1, 0x88, 0x41, 0x3a, 0x94,
	0x3a, 0x5a, 0xef, 0x50, 0xb8, 0xa4, 0x44, 0x09, 0x88, 0xb2, 0x00, 0x3e, 0x0a, 0x3e, 0xfa, 0xfa,
	0xf1, 0x9a, 0xe5, 0xdf, 0x2c, 0xff, 0x62, 0xf9, 0x26, 0xcb, 0x5d, 0x52, 0xfb, 0x7b, 0x85, 0x2f,
	0x7c, 0x05, 0x8a, 0x65, 0x5a, 0x75, 0x1c, 0x4e, 0x69, 0x56, 0xeb, 0xae, 0x5a, 0xeb, 0x32, 0xfb,
	0xff, 0x10, 0x7c, 0xd9, 0xa0, 0x07, 0xf4, 0x68, 0xb7, 0xd5, 0x9d, 0xcc, 0x72, 0xf2, 0x36, 0xcb,
	0x2f, 0x0a, 0x1b, 0x07, 0x49, 0x71, 0xed, 0x4b, 0xb1, 0xb1, 0xfc, 0xb5, 0x83, 0xcd, 0xf1, 0xa6,
	0x7d, 0x25, 0x50, 0x0f, 0xa0, 0x94, 0x22, 0x3e, 0x8d, 0x00, 0x85, 0x92, 0x08, 0xbc, 0xa9, 0x75,
	0xd3, 0x98, 0x00, 0x88, 0xbd, 0xaa, 0xba, 0x7e, 0x97, 0xd5, 0x5c, 0x52, 0x7d, 0xdb, 0x37, 0x8d,
	0x7f, 0x95, 0xe5, 0x7c, 0x6d, 0x39, 0xfd, 0x93, 0x85, 0x77, 0xda, 0xbd, 0x2d, 0x97, 0x54, 0xa7,
	0x7d, 0xb6, 0xf3, 0x3c, 0xce, 0xc9, 0xc7, 0x38, 0x27, 0x2d, 0x3f, 0x99, 0x33, 0x3a, 0x9d, 0x33,
	0xfa, 0x3e, 0x67, 0xf4, 0x65, 0xc1, 0xc8, 0x74, 0xc1, 0xc8, 0xeb, 0x82, 0x91, 0xfb, 0xdb, 0x1f,
	0x4b, 0x7e, 0xf3, 0x15, 0xb5, 0x5d, 0x9d, 0xf2, 0xe4, 0x33, 0x00, 0x00, 0xff, 0xff, 0x06, 0xd8,
	0x6e, 0x49, 0xdc, 0x01, 0x00, 0x00,
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
		size := m.NubID.Size()
		i -= size
		if _, err := m.NubID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.From.Size()
		i -= size
		if _, err := m.From.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	l = m.From.Size()
	n += 1 + l + sovMessage(uint64(l))
	l = m.NubID.Size()
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
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NubID", wireType)
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
			if err := m.NubID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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