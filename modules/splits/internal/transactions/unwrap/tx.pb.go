// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/splits/transactions/unwrap/v1beta1/tx.proto

package unwrap

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	"github.com/persistenceOne/persistenceSDK/schema/proto/types"
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
	From      github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=from,proto3,customtype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"from" valid:"required~required field from missing"`
	FromID    types.ID                                      `protobuf:"bytes,2,opt,name=from_iD,json=fromID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"from_iD" valid:"required~required field from_iD missing"`
	OwnableID types.ID                                      `protobuf:"bytes,4,opt,name=ownable_iD,json=ownableID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"ownable_iD" valid:"required~required ownable_iD from missing"`
	Value     github_com_cosmos_cosmos_sdk_types.Int        `protobuf:"bytes,5,opt,name=value,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"value" valid:"required~required value from missing"`
}

func (m message) Reset()         { m = message{} }
func (m message) String() string { return proto.CompactTextString(&m) }
func (message) ProtoMessage()    {}
func (*message) Descriptor() ([]byte, []int) {
	return fileDescriptor_d83165358ef093b8, []int{0}
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
	proto.RegisterFile("proto/splits/transactions/unwrap/v1beta1/tx.proto", fileDescriptor_d83165358ef093b8)
}

var fileDescriptor_d83165358ef093b8 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x93, 0xb5, 0xbb, 0xeb, 0xce, 0x31, 0xa7, 0xd8, 0x43, 0x22, 0x1e, 0x16, 0x41, 0x36,
	0x43, 0x15, 0x15, 0x04, 0x0f, 0x5d, 0xb2, 0x60, 0x10, 0x11, 0xb2, 0x37, 0x0f, 0x96, 0x49, 0xf2,
	0x4c, 0x83, 0xc9, 0x4c, 0x9c, 0xf7, 0xd2, 0xda, 0x6f, 0xe0, 0xd1, 0x8f, 0xd0, 0x8f, 0xd3, 0x63,
	0x8f, 0xe2, 0xa1, 0x48, 0x7b, 0x11, 0x3f, 0x85, 0x24, 0x13, 0xb1, 0x87, 0x1e, 0x2a, 0x78, 0x7a,
	0x33, 0xcc, 0xfb, 0xff, 0xfe, 0x30, 0xff, 0x3f, 0x1b, 0xd5, 0x5a, 0x91, 0xe2, 0x58, 0x97, 0x05,
	0x21, 0x27, 0x2d, 0x24, 0x8a, 0x94, 0x0a, 0x25, 0x91, 0x37, 0x72, 0xae, 0x45, 0xcd, 0x67, 0xa3,
	0x04, 0x48, 0x8c, 0x38, 0x7d, 0x0e, 0xba, 0x5d, 0x67, 0x90, 0x08, 0x84, 0xe1, 0x25, 0x4d, 0x0b,
	0x9d, 0x4d, 0x6a, 0xa1, 0x69, 0xc1, 0x0d, 0x24, 0x57, 0xb9, 0xfa, 0x7b, 0x32, 0xdb, 0xc3, 0x67,
	0x47, 0x1b, 0x68, 0xf8, 0xd4, 0x00, 0x52, 0xaf, 0x7b, 0xfe, 0x0f, 0x3a, 0xac, 0x95, 0x44, 0x30,
	0xc2, 0x07, 0xbf, 0x4e, 0xd8, 0x79, 0x05, 0x88, 0x22, 0x07, 0x27, 0x62, 0x83, 0x0f, 0x5a, 0x55,
	0xae, 0x7d, 0xdf, 0x7e, 0x78, 0x71, 0xfd, 0x74, 0xb5, 0xf1, 0xad, 0xef, 0x1b, 0xff, 0x2a, 0x2f,
	0x68, 0xda, 0x24, 0x41, 0xaa, 0x2a, 0x9e, 0x2a, 0xac, 0x14, 0xf6, 0xe3, 0x0a, 0xb3, 0x8f, 0x9c,
	0x16, 0x35, 0x60, 0x30, 0x4e, 0xd3, 0x71, 0x96, 0x69, 0x40, 0x8c, 0x3b, 0x84, 0xf3, 0x9e, 0x9d,
	0xb7, 0x73, 0x52, 0x84, 0xee, 0x49, 0x47, 0xbb, 0xe9, 0x69, 0x2f, 0xf7, 0x68, 0x35, 0x68, 0x2c,
	0x90, 0x40, 0xa6, 0xf0, 0x56, 0xc2, 0xfe, 0xf5, 0x36, 0x7c, 0xcd, 0x31, 0x9d, 0x42, 0x25, 0x38,
	0x01, 0xd2, 0xc4, 0xb8, 0x44, 0x61, 0x7c, 0xd6, 0x52, 0xa3, 0xd0, 0xc9, 0x18, 0x53, 0x73, 0x29,
	0x92, 0x12, 0x5a, 0x8b, 0xc1, 0xff, 0xb4, 0xb8, 0xe8, 0xc1, 0x51, 0xe8, 0x84, 0xec, 0x74, 0x26,
	0xca, 0x06, 0xdc, 0xd3, 0xce, 0x20, 0xe8, 0x0d, 0x2e, 0x8f, 0xf8, 0x91, 0x48, 0x52, 0x6c, 0xc4,
	0x2f, 0xee, 0x7e, 0x59, 0xfa, 0xd6, 0xcf, 0xa5, 0x6f, 0x3d, 0x7e, 0xc5, 0xee, 0xbc, 0xc1, 0xdc,
	0x19, 0x33, 0x66, 0x82, 0xba, 0x05, 0x99, 0x39, 0x6e, 0xd0, 0x36, 0x24, 0xd8, 0xcb, 0x2c, 0x36,
	0xd1, 0x0e, 0xef, 0x1d, 0x78, 0x31, 0xe1, 0x5d, 0xdf, 0xac, 0xb6, 0x9e, 0xbd, 0xde, 0x7a, 0xf6,
	0x8f, 0xad, 0x67, 0x7f, 0xdd, 0x79, 0xd6, 0x7a, 0xe7, 0x59, 0xdf, 0x76, 0x9e, 0xf5, 0xee, 0x51,
	0xa5, 0xb2, 0xa6, 0x04, 0xfc, 0xd3, 0x85, 0x42, 0x12, 0x68, 0x29, 0xca, 0x43, 0xa5, 0x48, 0xce,
	0xba, 0x12, 0x3c, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xea, 0x04, 0x92, 0x6c, 0xd8, 0x02, 0x00,
	0x00,
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
	SplitsSend(ctx context.Context, in *transactionRequest, opts ...grpc.CallOption) (*transactionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SplitsSend(ctx context.Context, in *transactionRequest, opts ...grpc.CallOption) (*transactionResponse, error) {
	out := new(transactionResponse)
	err := c.cc.Invoke(ctx, "/base.Msg/splitsSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SplitsSend(context.Context, *transactionRequest) (*transactionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SplitsSend(ctx context.Context, req *transactionRequest) (*transactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SplitsSend not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SplitsSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(transactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SplitsSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.Msg/SplitsSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SplitsSend(ctx, req.(*transactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "base.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "splitsSend",
			Handler:    _Msg_SplitsSend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/splits/transactions/unwrap/v1beta1/tx.proto",
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
		size := m.Value.Size()
		i -= size
		if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.OwnableID.Size()
		i -= size
		if _, err := m.OwnableID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
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
	n += 1 + l + sovTx(uint64(l))
	l = m.FromID.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.OwnableID.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.Value.Size()
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnableID", wireType)
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
			if err := m.OwnableID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
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
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
