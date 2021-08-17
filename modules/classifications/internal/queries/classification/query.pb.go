// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/classifications/internal/queries/classification/query.proto

package classification

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_helpers "github.com/persistenceOne/persistenceSDK/schema/helpers"
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

type QueryRequest struct {
	ClassificationID github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,1,opt,name=classification_i_d,json=classificationID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"classification_i_d" valid:"required~required ClassificationID missing"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddac3d6de870e613, []int{0}
}
func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

type QueryResponse struct {
	Success bool                                                               `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string                                                             `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	List    []github_com_persistenceOne_persistenceSDK_schema_helpers.Mappable `protobuf:"bytes,3,rep,name=list,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/helpers.Mappable" json:"list"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddac3d6de870e613, []int{1}
}
func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResponse.Unmarshal(m, b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
}
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return xxx_messageInfo_QueryResponse.Size(m)
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryRequest)(nil), "modules.classifications.internal.queries.classification.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "modules.classifications.internal.queries.classification.QueryResponse")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/classifications/internal/queries/classification/query.proto", fileDescriptor_ddac3d6de870e613)
}

var fileDescriptor_ddac3d6de870e613 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x3f, 0xaf, 0xd3, 0x30,
	0x10, 0xc0, 0x6d, 0x1e, 0x7f, 0xde, 0x8b, 0x40, 0x42, 0x51, 0x87, 0x88, 0x21, 0x8e, 0x32, 0xbd,
	0x29, 0x96, 0x60, 0x00, 0x75, 0x42, 0xa5, 0x03, 0x4f, 0x08, 0x21, 0xf2, 0x36, 0x84, 0x54, 0x52,
	0xe7, 0x48, 0x2d, 0x92, 0x38, 0xf5, 0x39, 0x95, 0xca, 0xc0, 0xdc, 0xb1, 0x1f, 0xa1, 0x95, 0xf8,
	0x06, 0x7c, 0x00, 0xd6, 0x8e, 0x1d, 0x11, 0x43, 0x85, 0xda, 0x85, 0x99, 0x4f, 0x80, 0x92, 0x26,
	0x52, 0x5b, 0x16, 0x60, 0xbb, 0xf3, 0x9d, 0x7f, 0x77, 0x3f, 0xd9, 0xd6, 0x75, 0x01, 0x1a, 0x25,
	0x1a, 0xc8, 0x05, 0x0c, 0x30, 0xfe, 0xc0, 0x33, 0x15, 0x97, 0x29, 0x20, 0x17, 0x69, 0x84, 0x28,
	0xdf, 0x4b, 0x11, 0x19, 0xa9, 0x72, 0xe4, 0x32, 0x37, 0xa0, 0xf3, 0x28, 0xe5, 0xe3, 0x12, 0xb4,
	0xfc, 0xa3, 0xa1, 0x3e, 0x9e, 0x06, 0x85, 0x56, 0x46, 0xd9, 0x8f, 0x1b, 0x48, 0x70, 0x02, 0x09,
	0x5a, 0x48, 0xd0, 0x40, 0x4e, 0x1a, 0x1e, 0x74, 0x12, 0x95, 0xa8, 0x9a, 0xc1, 0xab, 0x68, 0x8f,
	0xf3, 0xbf, 0x52, 0xeb, 0xee, 0xeb, 0x0a, 0x1f, 0xc2, 0xb8, 0x04, 0x34, 0xf6, 0x67, 0x6a, 0xd9,
	0xc7, 0x37, 0x07, 0x72, 0x10, 0x3b, 0xd4, 0xa3, 0x97, 0x17, 0xbd, 0xc9, 0x6a, 0xc3, 0xc8, 0xf7,
	0x0d, 0x7b, 0x92, 0x48, 0x33, 0x2a, 0x87, 0x81, 0x50, 0x19, 0x3f, 0x90, 0x7c, 0x95, 0xc3, 0x61,
	0x7a, 0xdd, 0x7f, 0xc1, 0x51, 0x8c, 0x20, 0x8b, 0xb8, 0x99, 0x16, 0x80, 0xc1, 0x55, 0xff, 0xd7,
	0x86, 0x3d, 0x9c, 0x44, 0xa9, 0x8c, 0xbb, 0xbe, 0x86, 0x71, 0x29, 0x35, 0xc4, 0x9f, 0xda, 0xc0,
	0x7b, 0x76, 0x34, 0xf6, 0xaa, 0xef, 0x65, 0x12, 0x51, 0xe6, 0x89, 0x1f, 0xde, 0x17, 0x27, 0xa5,
	0xee, 0xf9, 0x6c, 0xc1, 0xc8, 0xcf, 0x05, 0x23, 0xfe, 0x17, 0x6a, 0xdd, 0x6b, 0x0c, 0xb0, 0x50,
	0x39, 0x82, 0xed, 0x58, 0x77, 0xb0, 0x14, 0x02, 0x10, 0xeb, 0xb5, 0xcf, 0xc3, 0x36, 0xb5, 0x3b,
	0xd6, 0x2d, 0xd0, 0x5a, 0x69, 0xe7, 0x46, 0xa5, 0x13, 0xee, 0x13, 0xfb, 0xad, 0x75, 0x33, 0x95,
	0x68, 0x9c, 0x33, 0xef, 0xec, 0xf2, 0xa2, 0xf7, 0xbc, 0x71, 0x7c, 0xfa, 0xaf, 0x8e, 0x23, 0x48,
	0xab, 0x42, 0xf0, 0x32, 0x2a, 0x8a, 0x68, 0x98, 0x42, 0x58, 0x53, 0xbb, 0x9d, 0x76, 0xd3, 0xd9,
	0x92, 0x91, 0xf9, 0x92, 0x91, 0xc5, 0x92, 0x91, 0xde, 0xc7, 0xd5, 0xd6, 0xa5, 0xeb, 0xad, 0x4b,
	0x7f, 0x6c, 0x5d, 0x3a, 0xdf, 0xb9, 0x64, 0xbd, 0x73, 0xc9, 0xb7, 0x9d, 0x4b, 0xde, 0xbc, 0xfb,
	0xeb, 0xb9, 0xff, 0xf9, 0x9d, 0x86, 0xb7, 0xeb, 0xa7, 0x7f, 0xf4, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0xfc, 0xe2, 0x1b, 0x00, 0xa0, 0x02, 0x00, 0x00,
}

func (m *QueryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ClassificationID.Size()
		i -= size
		if _, err := m.ClassificationID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ClassificationID.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClassificationID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)