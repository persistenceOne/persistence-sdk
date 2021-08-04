// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/assets/transactions/mutate/tx.proto

package mutate

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_test_types "github.com/persistenceOne/persistenceSDK/schema/test_types"
	github_com_persistenceOne_persistenceSDK_schema_types "github.com/persistenceOne/persistenceSDK/schema/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
	From                  github_com_cosmos_cosmos_sdk_types.AccAddress                        `protobuf:"bytes,1,opt,name=from,proto3,customtype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"from"`
	FromID                github_com_persistenceOne_persistenceSDK_schema_test_types.ID        `protobuf:"bytes,2,opt,name=fromID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/test_types.ID" json:"fromID"`
	AssetID               github_com_persistenceOne_persistenceSDK_schema_test_types.ID        `protobuf:"bytes,3,opt,name=assetID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/test_types.ID" json:"assetID"`
	MutableMetaProperties github_com_persistenceOne_persistenceSDK_schema_types.MetaProperties `protobuf:"bytes,4,opt,name=mutableMetaProperties,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.MetaProperties" json:"mutableMetaProperties"`
	MutableProperties     github_com_persistenceOne_persistenceSDK_schema_types.Properties     `protobuf:"bytes,5,opt,name=mutableProperties,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.Properties" json:"mutableProperties"`
}

func (m message) Reset()         { m = message{} }
func (m message) String() string { return proto.CompactTextString(&m) }
func (message) ProtoMessage()    {}
func (*message) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cd41e0a206bbc04, []int{0}
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
	proto.RegisterType((*message)(nil), "base.message")
}

func init() {
	proto.RegisterFile("proto/assets/transactions/mutate/tx.proto", fileDescriptor_7cd41e0a206bbc04)
}

var fileDescriptor_7cd41e0a206bbc04 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xbf, 0xab, 0xd3, 0x50,
	0x14, 0x80, 0x13, 0x5f, 0x7c, 0x4f, 0x2f, 0x2e, 0x5e, 0x10, 0x62, 0x87, 0x44, 0x1c, 0x44, 0x91,
	0x97, 0x0b, 0x8a, 0x8b, 0x20, 0xf8, 0x9e, 0x29, 0x18, 0xb5, 0x28, 0x71, 0x13, 0xa4, 0xdc, 0x24,
	0xc7, 0x34, 0xd8, 0xe4, 0xc6, 0x7b, 0x4e, 0xc4, 0xae, 0x4e, 0x8e, 0xee, 0x2e, 0xfd, 0x73, 0x3a,
	0x76, 0x14, 0x87, 0x22, 0xed, 0xe2, 0x9f, 0x21, 0xbd, 0x49, 0x31, 0x62, 0xc1, 0x1f, 0x38, 0xdd,
	0x84, 0x7b, 0xce, 0xf7, 0x91, 0xf0, 0xb1, 0x1b, 0xb5, 0x56, 0xa4, 0x84, 0x44, 0x04, 0x42, 0x41,
	0x5a, 0x56, 0x28, 0x53, 0x2a, 0x54, 0x85, 0xa2, 0x6c, 0x48, 0x12, 0x08, 0x7a, 0x17, 0x98, 0x19,
	0xee, 0x24, 0x12, 0x61, 0x70, 0x8d, 0x26, 0x85, 0xce, 0xc6, 0xb5, 0xd4, 0x34, 0x13, 0xed, 0x72,
	0xae, 0x72, 0xf5, 0xe3, 0xa9, 0x9d, 0x1e, 0x04, 0xbf, 0x05, 0x6b, 0x78, 0xd3, 0x00, 0x52, 0x37,
	0x2f, 0xfe, 0x60, 0x1e, 0x6b, 0x55, 0x21, 0xb4, 0x0b, 0x57, 0x3f, 0x39, 0xec, 0xa8, 0x04, 0x44,
	0x99, 0x03, 0x8f, 0x98, 0xf3, 0x4a, 0xab, 0xd2, 0xb5, 0xaf, 0xd8, 0xd7, 0xcf, 0x9f, 0xde, 0x59,
	0xac, 0x7c, 0xeb, 0xcb, 0xca, 0x3f, 0xce, 0x0b, 0x9a, 0x34, 0x49, 0x90, 0xaa, 0x52, 0xa4, 0x0a,
	0x4b, 0x85, 0xdd, 0x71, 0x8c, 0xd9, 0x6b, 0x41, 0xb3, 0x1a, 0x30, 0x38, 0x49, 0xd3, 0x93, 0x2c,
	0xd3, 0x80, 0x18, 0x1b, 0x04, 0x7f, 0xc9, 0x0e, 0xb7, 0x67, 0x14, 0xba, 0x67, 0x0c, 0x6c, 0xd8,
	0xc1, 0xee, 0xf5, 0x60, 0x35, 0x68, 0x2c, 0x90, 0xa0, 0x4a, 0xe1, 0x69, 0x05, 0xfd, 0xd7, 0xe7,
	0xe1, 0x63, 0x81, 0xe9, 0x04, 0x4a, 0x29, 0x08, 0x90, 0xc6, 0xad, 0x24, 0x0a, 0xe3, 0x0e, 0xca,
	0xc7, 0xec, 0xc8, 0x7c, 0x62, 0x14, 0xba, 0x07, 0xff, 0x93, 0xbf, 0xa3, 0xf2, 0xf7, 0x36, 0xbb,
	0xb4, 0xfd, 0x61, 0xc9, 0x14, 0x46, 0x40, 0xf2, 0x99, 0x56, 0x35, 0x68, 0x2a, 0x00, 0x5d, 0xc7,
	0xf8, 0x9e, 0x74, 0xbe, 0xf0, 0xaf, 0x7d, 0x46, 0xf5, 0x33, 0x33, 0xde, 0xaf, 0xe2, 0x6f, 0xd9,
	0xc5, 0xee, 0xa2, 0xe7, 0x3f, 0x6b, 0xfc, 0x0f, 0x3b, 0xff, 0xfd, 0x7f, 0xf3, 0xf7, 0xdc, 0xbf,
	0x2a, 0xee, 0x9e, 0xfb, 0x30, 0xf7, 0xad, 0x6f, 0x73, 0xdf, 0xba, 0xf5, 0x88, 0x1d, 0x8c, 0x30,
	0xe7, 0x0f, 0xd8, 0x85, 0xb6, 0xa8, 0x91, 0x69, 0x88, 0xbb, 0xc1, 0x36, 0xe2, 0xa0, 0x97, 0x57,
	0xdc, 0x56, 0x38, 0xb8, 0xbc, 0xe7, 0xa6, 0xed, 0xed, 0x74, 0xb8, 0x58, 0x7b, 0xf6, 0x72, 0xed,
	0xd9, 0x5f, 0xd7, 0x9e, 0xfd, 0x71, 0xe3, 0x59, 0xcb, 0x8d, 0x67, 0x7d, 0xde, 0x78, 0xd6, 0x8b,
	0x9b, 0xa5, 0xca, 0x9a, 0x29, 0xe0, 0x2e, 0xdb, 0xa2, 0x22, 0xd0, 0x95, 0x9c, 0xee, 0xeb, 0x37,
	0x39, 0x34, 0xdd, 0xde, 0xfe, 0x1e, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x31, 0x33, 0x28, 0x73, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	AssetsMutate(ctx context.Context, in *transactionRequest, opts ...grpc.CallOption) (*transactionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AssetsMutate(ctx context.Context, in *transactionRequest, opts ...grpc.CallOption) (*transactionResponse, error) {
	out := new(transactionResponse)
	err := c.cc.Invoke(ctx, "/base.Msg/assetsMutate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	AssetsMutate(context.Context, *transactionRequest) (*transactionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AssetsMutate(ctx context.Context, req *transactionRequest) (*transactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssetsMutate not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AssetsMutate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(transactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AssetsMutate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.Msg/AssetsMutate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AssetsMutate(ctx, req.(*transactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "base.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "assetsMutate",
			Handler:    _Msg_AssetsMutate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/assets/transactions/mutate/tx.proto",
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
		size := m.MutableProperties.Size()
		i -= size
		if _, err := m.MutableProperties.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.MutableMetaProperties.Size()
		i -= size
		if _, err := m.MutableMetaProperties.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.AssetID.Size()
		i -= size
		if _, err := m.AssetID.MarshalTo(dAtA[i:]); err != nil {
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
	{
		size := len(m.From)
		i -= size
		if _, err := m.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	n += 1 + l + sovTx(uint64(l))
	l = m.FromID.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.AssetID.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.MutableMetaProperties.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.MutableProperties.Size()
	n += 1 + l + sovTx(uint64(l))
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
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
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
			if err := m.AssetID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutableMetaProperties", wireType)
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
			if err := m.MutableMetaProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutableProperties", wireType)
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
			if err := m.MutableProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
