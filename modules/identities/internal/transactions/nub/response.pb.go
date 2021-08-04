// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/identities/transactions/nub/response.proto

package nub

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

type transactionResponse struct {
	Success bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   error `protobuf:"bytes,2,opt,name=error,proto3,customtype=error" json:"error"`
}

func (m transactionResponse) Reset()         { m = transactionResponse{} }
func (m transactionResponse) String() string { return proto.CompactTextString(&m) }
func (transactionResponse) ProtoMessage()    {}
func (*transactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5be042a366756653, []int{0}
}
func (m *transactionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *transactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *transactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResponse.Merge(m, src)
}
func (m *transactionResponse) XXX_Size() int {
	return m.Size()
}
func (m *transactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResponse proto.InternalMessageInfo

func (m *transactionResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*transactionResponse)(nil), "base.transactionResponse")
}

func init() {
	proto.RegisterFile("proto/identities/transactions/nub/response.proto", fileDescriptor_5be042a366756653)
}

var fileDescriptor_5be042a366756653 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xc9, 0x4c, 0x2d, 0xd6, 0x2f, 0x29, 0x4a,
	0xcc, 0x2b, 0x4e, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0xcf, 0x2b, 0x4d, 0xd2, 0x2f, 0x4a,
	0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x03, 0x2b, 0x15, 0x62, 0x49, 0x4a, 0x2c, 0x4e, 0x95,
	0x52, 0x2b, 0xc9, 0xc8, 0x2c, 0x4a, 0x89, 0x2f, 0x48, 0x2c, 0x2a, 0xa9, 0xd4, 0x87, 0x98, 0x91,
	0x9e, 0x9f, 0x9e, 0x8f, 0x60, 0x41, 0x54, 0x4b, 0x69, 0x63, 0xaa, 0x4b, 0xce, 0x2f, 0xce, 0xcd,
	0x2f, 0x8e, 0x47, 0xe6, 0x40, 0x14, 0x2b, 0x15, 0x72, 0x09, 0x23, 0xd9, 0x1e, 0x04, 0xb5, 0x57,
	0x48, 0x82, 0x8b, 0xbd, 0xb8, 0x34, 0x39, 0x39, 0xb5, 0xb8, 0x58, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x23, 0x08, 0xc6, 0x15, 0x52, 0xe6, 0x62, 0x4d, 0x2d, 0x2a, 0xca, 0x2f, 0x92, 0x60, 0x52, 0x60,
	0xd4, 0xe0, 0x74, 0xe2, 0x3d, 0x71, 0x4f, 0x9e, 0xe1, 0xd6, 0x3d, 0x79, 0x88, 0x60, 0x10, 0x84,
	0xb2, 0x92, 0x7e, 0xb1, 0x40, 0x9e, 0xe1, 0xd2, 0x16, 0x5d, 0xe1, 0x10, 0x4c, 0xb3, 0x9d, 0xdc,
	0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5,
	0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x27, 0x37, 0x3f, 0xa5, 0x34,
	0x27, 0xb5, 0x18, 0x39, 0x6c, 0x32, 0xf3, 0x4a, 0x52, 0x8b, 0xf2, 0x12, 0x73, 0x30, 0x02, 0x29,
	0x89, 0x0d, 0xec, 0x03, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xb2, 0x55, 0x84, 0x50,
	0x01, 0x00, 0x00,
}

func (m *transactionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *transactionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *transactionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := len(m.Error.Error())
		i -= size
		if _, err := m.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintResponse(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintResponse(dAtA []byte, offset int, v uint64) int {
	offset -= sovResponse(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *transactionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	l = len(m.Error.Error())
	n += 1 + l + sovResponse(uint64(l))
	return n
}

func sovResponse(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResponse(x uint64) (n int) {
	return sovResponse(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *transactionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResponse
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
			return fmt.Errorf("proto: transactionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: transactionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponse
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponse
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
				return ErrInvalidLengthResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResponse(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResponse
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
func skipResponse(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResponse
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
					return 0, ErrIntOverflowResponse
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
					return 0, ErrIntOverflowResponse
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
				return 0, ErrInvalidLengthResponse
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResponse
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResponse
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResponse        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResponse          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResponse = fmt.Errorf("proto: unexpected end of group")
)
