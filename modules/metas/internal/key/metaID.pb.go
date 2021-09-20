// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/metas/internal/key/metaID.proto

package key

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	base "github.com/persistenceOne/persistenceSDK/schema/types/base"
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

type MetaID struct {
	TypeID base.ID `protobuf:"bytes,1,opt,name=type_i_d,json=typeID,proto3" json:"type_i_d"`
	HashID base.ID `protobuf:"bytes,2,opt,name=hash_i_d,json=hashID,proto3" json:"hash_i_d"`
}

func (m *MetaID) Reset()      { *m = MetaID{} }
func (*MetaID) ProtoMessage() {}
func (*MetaID) Descriptor() ([]byte, []int) {
	return fileDescriptor_872133bbc3312930, []int{0}
}
func (m *MetaID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetaID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetaID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MetaID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaID.Merge(m, src)
}
func (m *MetaID) XXX_Size() int {
	return m.Size()
}
func (m *MetaID) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaID.DiscardUnknown(m)
}

var xxx_messageInfo_MetaID proto.InternalMessageInfo

func (m *MetaID) GetTypeID() base.ID {
	if m != nil {
		return m.TypeID
	}
	return base.ID{}
}

func (m *MetaID) GetHashID() base.ID {
	if m != nil {
		return m.HashID
	}
	return base.ID{}
}

func init() {
	proto.RegisterType((*MetaID)(nil), "persistence_sdk.modules.metas.internal.key.MetaID")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/metas/internal/key/metaID.proto", fileDescriptor_872133bbc3312930)
}

var fileDescriptor_872133bbc3312930 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2f, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e, 0xc9, 0xd6, 0xcf, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0xcf, 0x4d, 0x2d, 0x49, 0x2c, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca,
	0x4b, 0xcc, 0xd1, 0xcf, 0x4e, 0xad, 0x04, 0x0b, 0x79, 0xba, 0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4,
	0x0b, 0x69, 0xa1, 0x69, 0xd4, 0x83, 0x6a, 0xd4, 0x03, 0x6b, 0xd4, 0x83, 0x69, 0xd4, 0xcb, 0x4e,
	0xad, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x6b, 0xd3, 0x07, 0xb1, 0x20, 0x26, 0x48, 0xa1,
	0x9b, 0xa0, 0x5f, 0x9c, 0x9c, 0x91, 0x9a, 0x9b, 0xa8, 0x5f, 0x52, 0x59, 0x90, 0x5a, 0xac, 0x9f,
	0x94, 0x58, 0x9c, 0xaa, 0x9f, 0x99, 0x02, 0x51, 0xab, 0x34, 0x87, 0x91, 0x8b, 0xcd, 0x17, 0x6c,
	0xbd, 0x90, 0x2b, 0x17, 0x07, 0x48, 0x45, 0x7c, 0x66, 0x7c, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0xb7, 0x91, 0xaa, 0x1e, 0xba, 0x5b, 0x20, 0x26, 0xe9, 0x81, 0x4d, 0xd2, 0x03, 0x99, 0xa4, 0xe7,
	0xe9, 0xe2, 0xc4, 0x72, 0xe2, 0x9e, 0x3c, 0x43, 0x10, 0x1b, 0x48, 0x10, 0x62, 0x4c, 0x46, 0x62,
	0x71, 0x06, 0xd8, 0x18, 0x26, 0x32, 0x8c, 0x01, 0x69, 0xf6, 0x74, 0xb1, 0x62, 0x99, 0xb1, 0x40,
	0x9e, 0xc1, 0x29, 0xf6, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63,
	0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x9c, 0xd3,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x91, 0x8c, 0xf7, 0xcf, 0x4b, 0x45,
	0xe6, 0x06, 0xbb, 0x78, 0xe3, 0x09, 0xf8, 0x24, 0x36, 0x70, 0x20, 0x18, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x78, 0xd7, 0xf1, 0xa5, 0xad, 0x01, 0x00, 0x00,
}

func (m *MetaID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetaID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetaID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.HashID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMetaID(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.TypeID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMetaID(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMetaID(dAtA []byte, offset int, v uint64) int {
	offset -= sovMetaID(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MetaID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TypeID.Size()
	n += 1 + l + sovMetaID(uint64(l))
	l = m.HashID.Size()
	n += 1 + l + sovMetaID(uint64(l))
	return n
}

func sovMetaID(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMetaID(x uint64) (n int) {
	return sovMetaID(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MetaID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetaID
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
			return fmt.Errorf("proto: MetaID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetaID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetaID
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetaID
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetaID
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TypeID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetaID
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetaID
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetaID
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HashID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetaID(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetaID
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
func skipMetaID(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetaID
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
					return 0, ErrIntOverflowMetaID
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
					return 0, ErrIntOverflowMetaID
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
				return 0, ErrInvalidLengthMetaID
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMetaID
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMetaID
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMetaID        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetaID          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMetaID = fmt.Errorf("proto: unexpected end of group")
)