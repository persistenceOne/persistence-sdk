// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/persistence/assets/splits/transactions/unwrap/v1beta1/request.proto

package unwrap

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	test_types "github.com/persistenceOne/persistenceSDK/schema/test_types"
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

type transactionRequest struct {
	BaseReq   test_types.BaseReq `protobuf:"bytes,1,opt,name=base_req,json=baseReq,proto3" json:"base_req"`
	FromID    string             `protobuf:"bytes,2,opt,name=from_iD,json=fromID,proto3" json:"from_iD,omitempty"`
	OwnableID string             `protobuf:"bytes,4,opt,name=ownable_iD,json=ownableID,proto3" json:"ownable_iD,omitempty"`
	Value     string             `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *transactionRequest) Reset()         { *m = transactionRequest{} }
func (m *transactionRequest) String() string { return proto.CompactTextString(m) }
func (*transactionRequest) ProtoMessage()    {}
func (*transactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_93a33951dfdbcc05, []int{0}
}
func (m *transactionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *transactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *transactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionRequest.Merge(m, src)
}
func (m *transactionRequest) XXX_Size() int {
	return m.Size()
}
func (m *transactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionRequest proto.InternalMessageInfo

func (m *transactionRequest) GetFromID() string {
	if m != nil {
		return m.FromID
	}
	return ""
}

func (m *transactionRequest) GetOwnableID() string {
	if m != nil {
		return m.OwnableID
	}
	return ""
}

func (m *transactionRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*transactionRequest)(nil), "persistence.assets.splits.transactions.unwrap.v1beta1.transactionRequest")
}

func init() {
	proto.RegisterFile("proto/persistence/assets/splits/transactions/unwrap/v1beta1/request.proto", fileDescriptor_93a33951dfdbcc05)
}

var fileDescriptor_93a33951dfdbcc05 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x93, 0x8f, 0xfe, 0xf9, 0x1a, 0x77, 0x83, 0x60, 0x29, 0x98, 0x16, 0x17, 0x52, 0x28,
	0xce, 0x50, 0x8b, 0x1b, 0x97, 0xa5, 0x2e, 0xb2, 0x0d, 0xae, 0xdc, 0x84, 0x49, 0x7b, 0xad, 0x81,
	0x74, 0x26, 0x99, 0x3b, 0x69, 0xf1, 0x2d, 0x7c, 0x04, 0x1f, 0xc2, 0x87, 0xa8, 0xbb, 0xe2, 0xca,
	0x95, 0x48, 0xbb, 0xf1, 0x31, 0x24, 0x99, 0x58, 0x03, 0x75, 0x97, 0x93, 0x73, 0xe6, 0x37, 0xf7,
	0xcc, 0x75, 0xbc, 0x44, 0x49, 0x2d, 0x59, 0x02, 0x0a, 0x23, 0xd4, 0x20, 0xa6, 0xc0, 0x38, 0x22,
	0x68, 0x64, 0x98, 0xc4, 0x91, 0x46, 0xa6, 0x15, 0x17, 0xc8, 0xa7, 0x3a, 0x92, 0x02, 0x59, 0x26,
	0x56, 0x8a, 0x27, 0x6c, 0x39, 0x0c, 0x41, 0xf3, 0x21, 0x53, 0x90, 0x66, 0x80, 0x9a, 0x16, 0x0c,
	0x72, 0x55, 0x81, 0x50, 0x03, 0xa1, 0x06, 0x42, 0xab, 0x10, 0x6a, 0x20, 0xb4, 0x84, 0x74, 0xce,
	0xf5, 0x43, 0xa4, 0x66, 0x41, 0xc2, 0x95, 0x7e, 0x64, 0x66, 0x9a, 0xb9, 0x9c, 0xcb, 0xdf, 0x2f,
	0x83, 0xef, 0x0c, 0x0e, 0x73, 0x53, 0x89, 0x0b, 0x89, 0x41, 0x55, 0x94, 0xe1, 0xd1, 0x61, 0x2d,
	0xe3, 0xb3, 0x90, 0x23, 0xec, 0x0b, 0xe4, 0x42, 0x41, 0x6a, 0x0e, 0x9d, 0xbd, 0xda, 0x0e, 0xa9,
	0x4c, 0xea, 0x9b, 0x76, 0xc4, 0x73, 0xfe, 0xe7, 0xb9, 0x40, 0x41, 0xda, 0xb6, 0x7b, 0x76, 0xff,
	0xe8, 0xb2, 0x4f, 0xab, 0x55, 0xcb, 0x8b, 0xf3, 0xcc, 0x4f, 0x29, 0x3a, 0xe6, 0x08, 0x3e, 0xa4,
	0xe3, 0xda, 0xfa, 0xa3, 0x6b, 0xf9, 0xcd, 0xd0, 0x48, 0x72, 0xe2, 0x34, 0xef, 0x95, 0x5c, 0x04,
	0xd1, 0xa4, 0xfd, 0xaf, 0x67, 0xf7, 0x5b, 0x7e, 0x23, 0x97, 0xde, 0x84, 0x9c, 0x3a, 0x8e, 0x5c,
	0x09, 0x1e, 0xc6, 0x90, 0x7b, 0xb5, 0xc2, 0x6b, 0x95, 0x7f, 0xbc, 0x09, 0x39, 0x76, 0xea, 0x4b,
	0x1e, 0x67, 0xd0, 0xae, 0x17, 0x8e, 0x11, 0xd7, 0x9d, 0xaf, 0xe7, 0xae, 0xf5, 0xf6, 0x72, 0x41,
	0x6e, 0x0f, 0x86, 0x1e, 0xdf, 0xac, 0xb7, 0xae, 0xbd, 0xd9, 0xba, 0xf6, 0xe7, 0xd6, 0xb5, 0x9f,
	0x76, 0xae, 0xb5, 0xd9, 0xb9, 0xd6, 0xfb, 0xce, 0xb5, 0xee, 0x06, 0x0b, 0x39, 0xcb, 0x62, 0xd8,
	0x2f, 0x38, 0x12, 0x1a, 0x94, 0xe0, 0xf1, 0x5f, 0x9b, 0x0e, 0x1b, 0xc5, 0xcb, 0x8c, 0xbe, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x06, 0x80, 0xa0, 0xdb, 0x27, 0x02, 0x00, 0x00,
}

func (m *transactionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *transactionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *transactionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.OwnableID) > 0 {
		i -= len(m.OwnableID)
		copy(dAtA[i:], m.OwnableID)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.OwnableID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.FromID) > 0 {
		i -= len(m.FromID)
		copy(dAtA[i:], m.FromID)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.FromID)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.BaseReq.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintRequest(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *transactionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.BaseReq.Size()
	n += 1 + l + sovRequest(uint64(l))
	l = len(m.FromID)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.OwnableID)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	return n
}

func sovRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRequest(x uint64) (n int) {
	return sovRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *transactionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequest
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
			return fmt.Errorf("proto: transactionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: transactionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseReq", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseReq.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnableID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnableID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequest
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
func skipRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRequest
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
					return 0, ErrIntOverflowRequest
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
					return 0, ErrIntOverflowRequest
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
				return 0, ErrInvalidLengthRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRequest = fmt.Errorf("proto: unexpected end of group")
)