// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/identities/transactions/mutate/tx.proto

package mutate

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	"github.com/persistenceOne/persistenceSDK/schema/proto/types"
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
	From                  github_com_cosmos_cosmos_sdk_types.AccAddress                        `protobuf:"bytes,1,opt,name=from,proto3,customtype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"from"  valid:"required~required field from missing"`
	FromID                types.ID                                                             `protobuf:"bytes,2,opt,name=fromID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"fromID"  valid:"required~required field fromID missing"`
	IdentityID            types.ID                                                             `protobuf:"bytes,3,opt,name=identityID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"identityID"  valid:"required~required field identityID missing"`
	MutableMetaProperties github_com_persistenceOne_persistenceSDK_schema_types.MetaProperties `protobuf:"bytes,4,opt,name=mutableMetaProperties,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.MetaProperties" json:"mutableMetaProperties"  valid:"required~required field mutableMetaProperties missing"`
	MutableProperties     github_com_persistenceOne_persistenceSDK_schema_types.Properties     `protobuf:"bytes,5,opt,name=mutableProperties,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.Properties" json:"mutableProperties"  valid:"required~required field mutableProperties missing"`
}

func (m message) Reset()         { m = message{} }
func (m message) String() string { return proto.CompactTextString(&m) }
func (message) ProtoMessage()    {}
func (*message) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7df2b182a1ea9e8, []int{0}
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
	proto.RegisterFile("proto/identities/transactions/mutate/tx.proto", fileDescriptor_d7df2b182a1ea9e8)
}

var fileDescriptor_d7df2b182a1ea9e8 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x31, 0x8b, 0xd4, 0x40,
	0x14, 0x80, 0x13, 0x6f, 0x3d, 0x75, 0x2a, 0x1d, 0x10, 0xe2, 0x16, 0x59, 0xb1, 0x10, 0x9b, 0xcd,
	0xc0, 0x1d, 0x36, 0x82, 0xe0, 0x1d, 0x11, 0x0d, 0xba, 0x2a, 0x6b, 0x27, 0xc8, 0x31, 0x9b, 0x3c,
	0xb3, 0x83, 0x9b, 0x99, 0x38, 0xef, 0x45, 0xdc, 0xd6, 0xca, 0xd2, 0x9f, 0x70, 0xfe, 0x9b, 0x2b,
	0xaf, 0x14, 0x8b, 0x43, 0x76, 0x1b, 0x7f, 0x86, 0x64, 0x32, 0x70, 0x11, 0xb7, 0x58, 0xe5, 0xaa,
	0x49, 0x98, 0x79, 0xdf, 0xc7, 0x0c, 0x1f, 0x1b, 0xd7, 0xd6, 0x90, 0x11, 0xaa, 0x00, 0x4d, 0x8a,
	0x14, 0xa0, 0x20, 0x2b, 0x35, 0xca, 0x9c, 0x94, 0xd1, 0x28, 0xaa, 0x86, 0x24, 0x81, 0xa0, 0x4f,
	0x89, 0x3b, 0xc7, 0x07, 0x33, 0x89, 0x30, 0xbc, 0x4b, 0x73, 0x65, 0x8b, 0xa3, 0x5a, 0x5a, 0x5a,
	0x8a, 0x0e, 0x50, 0x9a, 0xd2, 0x9c, 0x7f, 0x75, 0xa7, 0x87, 0x7b, 0x5b, 0xc1, 0x2d, 0x7c, 0x68,
	0x00, 0xc9, 0xcf, 0xec, 0x6f, 0x39, 0x83, 0xb5, 0xd1, 0x08, 0xdd, 0xd0, 0x9d, 0x6f, 0x03, 0x76,
	0xa5, 0x02, 0x44, 0x59, 0x02, 0xcf, 0xd8, 0xe0, 0x9d, 0x35, 0x55, 0x14, 0xde, 0x0e, 0xef, 0x5d,
	0x3b, 0xbc, 0x7f, 0x72, 0x36, 0x0a, 0x7e, 0x9c, 0x8d, 0xc6, 0xa5, 0xa2, 0x79, 0x33, 0x4b, 0x72,
	0x53, 0x89, 0xdc, 0x60, 0x65, 0xd0, 0x2f, 0x63, 0x2c, 0xde, 0x0b, 0x5a, 0xd6, 0x80, 0xc9, 0x41,
	0x9e, 0x1f, 0x14, 0x85, 0x05, 0xc4, 0xa9, 0x43, 0xf0, 0xb7, 0x6c, 0xb7, 0x5d, 0xb3, 0x34, 0xba,
	0xe4, 0x60, 0x8f, 0x3d, 0xec, 0x61, 0x0f, 0x56, 0x83, 0x45, 0x85, 0x04, 0x3a, 0x87, 0x97, 0x1a,
	0xfa, 0xbf, 0xaf, 0xd3, 0x67, 0x02, 0xf3, 0x39, 0x54, 0x52, 0x10, 0x20, 0x1d, 0x75, 0x92, 0x2c,
	0x9d, 0x7a, 0x28, 0x07, 0xc6, 0xfc, 0x35, 0x97, 0x59, 0x1a, 0xed, 0x5c, 0xa4, 0xa2, 0x07, 0xe6,
	0x9f, 0x43, 0x76, 0xb3, 0x7d, 0xb6, 0xd9, 0x02, 0x26, 0x40, 0xf2, 0x95, 0x35, 0x35, 0xd8, 0xf6,
	0x65, 0xa3, 0x81, 0x53, 0x3e, 0xf7, 0xca, 0xf4, 0x9f, 0x95, 0xce, 0xf6, 0x27, 0x73, 0xba, 0x59,
	0xc5, 0x3f, 0xb2, 0x1b, 0x7e, 0xa3, 0xe7, 0xbf, 0xec, 0xfc, 0x4f, 0xbd, 0xff, 0xd1, 0xff, 0xf9,
	0x7b, 0xee, 0xbf, 0x15, 0x0f, 0xae, 0x7e, 0x39, 0x1e, 0x05, 0xbf, 0x8e, 0x47, 0xc1, 0xde, 0x0b,
	0xb6, 0x33, 0xc1, 0x92, 0x3f, 0x61, 0xd7, 0xcf, 0xdb, 0x9a, 0xb8, 0x9a, 0x78, 0x94, 0xb4, 0x59,
	0x27, 0xbd, 0xd0, 0xa6, 0x5d, 0x93, 0xc3, 0x5b, 0x1b, 0x76, 0xba, 0xf2, 0x0e, 0xb3, 0x93, 0x55,
	0x1c, 0x9e, 0xae, 0xe2, 0xf0, 0xe7, 0x2a, 0x0e, 0xbf, 0xae, 0xe3, 0xe0, 0x74, 0x1d, 0x07, 0xdf,
	0xd7, 0x71, 0xf0, 0x46, 0x54, 0xa6, 0x68, 0x16, 0x80, 0xfd, 0x88, 0x95, 0x26, 0xb0, 0x5a, 0x2e,
	0x36, 0xd5, 0x3c, 0xdb, 0x75, 0x15, 0xef, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x34, 0x58, 0x36,
	0xcf, 0x8d, 0x03, 0x00, 0x00,
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
	IdentitiesMutate(ctx context.Context, in *transactionRequest, opts ...grpc.CallOption) (*transactionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) IdentitiesMutate(ctx context.Context, in *transactionRequest, opts ...grpc.CallOption) (*transactionResponse, error) {
	out := new(transactionResponse)
	err := c.cc.Invoke(ctx, "/base.Msg/identitiesMutate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	IdentitiesMutate(context.Context, *transactionRequest) (*transactionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) IdentitiesMutate(ctx context.Context, req *transactionRequest) (*transactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IdentitiesMutate not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_IdentitiesMutate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(transactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).IdentitiesMutate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.Msg/IdentitiesMutate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).IdentitiesMutate(ctx, req.(*transactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "base.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "identitiesMutate",
			Handler:    _Msg_IdentitiesMutate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/identities/transactions/mutate/tx.proto",
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
		size := m.IdentityID.Size()
		i -= size
		if _, err := m.IdentityID.MarshalTo(dAtA[i:]); err != nil {
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
	l = m.IdentityID.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
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
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
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
				return fmt.Errorf("proto: wrong wireType = %d for field IdentityID", wireType)
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
			if err := m.IdentityID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
