// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/assets/internal/queries/asset/query.proto

package asset

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
	AssetID github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,1,opt,name=asset_i_d,json=assetID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"asset_i_d" valid:"required~required AssetID missing"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a52141441ac620ee, []int{0}
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
	return fileDescriptor_a52141441ac620ee, []int{1}
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
	proto.RegisterType((*QueryRequest)(nil), "modules.assets.internal.queries.asset.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "modules.assets.internal.queries.asset.QueryResponse")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/assets/internal/queries/asset/query.proto", fileDescriptor_a52141441ac620ee)
}

var fileDescriptor_a52141441ac620ee = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbf, 0xea, 0xd3, 0x50,
	0x14, 0xc7, 0xef, 0xf5, 0xa7, 0xb6, 0x0d, 0xba, 0x84, 0x0e, 0xc1, 0x21, 0x37, 0x04, 0x84, 0xba,
	0xe4, 0x0e, 0x2e, 0xd2, 0xc9, 0x96, 0x0e, 0x16, 0x51, 0x31, 0x6e, 0x22, 0x94, 0x34, 0x39, 0x24,
	0x17, 0x93, 0xdc, 0xf4, 0x9e, 0x1b, 0xa1, 0x8b, 0x73, 0xdd, 0xfa, 0x08, 0x2d, 0xf8, 0x06, 0xbe,
	0x44, 0xc7, 0x8e, 0xe2, 0x50, 0xa4, 0x5d, 0x9c, 0x7d, 0x02, 0xc9, 0x3f, 0xe8, 0x58, 0xb7, 0xf3,
	0xfd, 0x5e, 0x3e, 0xe7, 0x7c, 0x0f, 0xf7, 0x18, 0x93, 0x02, 0x14, 0x0a, 0xd4, 0x90, 0x87, 0xb0,
	0xc0, 0xe8, 0x33, 0xcf, 0x64, 0x54, 0xa6, 0x80, 0x3c, 0x40, 0x04, 0x8d, 0x5c, 0xe4, 0x1a, 0x54,
	0x1e, 0xa4, 0x7c, 0x55, 0x82, 0x12, 0x9d, 0x5f, 0xab, 0xb5, 0x57, 0x28, 0xa9, 0xa5, 0xf9, 0xb4,
	0x45, 0xbc, 0x06, 0xf1, 0x3a, 0xc4, 0x6b, 0x91, 0xc6, 0x7f, 0x32, 0x8c, 0x65, 0x2c, 0x6b, 0x82,
	0x57, 0x55, 0x03, 0xbb, 0xdf, 0xa9, 0xf1, 0xe8, 0x7d, 0xd5, 0xcc, 0x87, 0x55, 0x09, 0xa8, 0xcd,
	0x6f, 0xd4, 0x18, 0xd4, 0xc0, 0x42, 0x2c, 0x22, 0x8b, 0x3a, 0x74, 0x34, 0x98, 0x66, 0x87, 0x13,
	0x23, 0xbf, 0x4e, 0xec, 0x45, 0x2c, 0x74, 0x52, 0x2e, 0xbd, 0x50, 0x66, 0xfc, 0x2a, 0xf7, 0xbb,
	0x1c, 0xae, 0xe5, 0x87, 0xd9, 0x6b, 0x8e, 0x61, 0x02, 0x59, 0xc0, 0xf5, 0xba, 0x00, 0xf4, 0xe6,
	0xb3, 0xbf, 0x27, 0xf6, 0xec, 0x4b, 0x90, 0x8a, 0x68, 0xec, 0x2a, 0x58, 0x95, 0x42, 0x41, 0xf4,
	0xb5, 0x2b, 0x9c, 0x49, 0x35, 0x6d, 0x3e, 0x73, 0x32, 0x81, 0x28, 0xf2, 0xd8, 0xf5, 0x7b, 0x41,
	0xe3, 0x8c, 0xfb, 0x9b, 0x1d, 0x23, 0x7f, 0x76, 0x8c, 0xb8, 0x3f, 0xa8, 0xf1, 0xb8, 0x8d, 0x89,
	0x85, 0xcc, 0x11, 0x4c, 0xcb, 0xe8, 0x61, 0x19, 0x86, 0x80, 0x58, 0x87, 0xec, 0xfb, 0x9d, 0x34,
	0x87, 0xc6, 0x03, 0x50, 0x4a, 0x2a, 0xeb, 0x5e, 0x15, 0xde, 0x6f, 0x84, 0xf9, 0xc9, 0xb8, 0x9f,
	0x0a, 0xd4, 0xd6, 0x9d, 0x73, 0x37, 0x1a, 0x4c, 0x5f, 0xb5, 0x1b, 0xbd, 0xfc, 0xdf, 0x8d, 0x12,
	0x48, 0xab, 0x07, 0xef, 0x4d, 0x50, 0x14, 0xc1, 0x32, 0x05, 0xbf, 0xee, 0x3a, 0x1e, 0x76, 0x49,
	0x37, 0x7b, 0x46, 0xb6, 0x7b, 0x46, 0x76, 0x7b, 0x46, 0xa6, 0xc9, 0xe1, 0x6c, 0xd3, 0xe3, 0xd9,
	0xa6, 0xbf, 0xcf, 0x36, 0xdd, 0x5e, 0x6c, 0x72, 0xbc, 0xd8, 0xe4, 0xe7, 0xc5, 0x26, 0x1f, 0xdf,
	0xde, 0x3c, 0xf7, 0xa6, 0x7b, 0x58, 0x3e, 0xac, 0x7f, 0xf3, 0xf9, 0xbf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xdd, 0x9e, 0x85, 0x7f, 0x4f, 0x02, 0x00, 0x00,
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
		size := m.AssetID.Size()
		i -= size
		if _, err := m.AssetID.MarshalTo(dAtA[i:]); err != nil {
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
	l = m.AssetID.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
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
			if err := m.AssetID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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