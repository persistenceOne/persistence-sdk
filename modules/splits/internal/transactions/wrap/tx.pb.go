// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/splits/transactions/wrap/v1beta1/tx.proto

package wrap

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

type Message struct {
	From   github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=from,proto3,customtype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"from"`
	FromID types.ID                                      `protobuf:"bytes,2,opt,name=from_iD,json=fromID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"from_iD"`
	Coins  github_com_cosmos_cosmos_sdk_types.Coins      `protobuf:"bytes,5,opt,name=coins,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_098357e27cff17ae, []int{0}
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
	proto.RegisterType((*Message)(nil), "base.message")
}

func init() {
	proto.RegisterFile("proto/splits/transactions/wrap/v1beta1/tx.proto", fileDescriptor_098357e27cff17ae)
}

var fileDescriptor_098357e27cff17ae = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbd, 0xce, 0xd3, 0x30,
	0x14, 0x86, 0x93, 0xd2, 0x1f, 0xf0, 0x98, 0x29, 0x74, 0x48, 0x10, 0x03, 0x2a, 0x43, 0x63, 0x0a,
	0x74, 0x41, 0x62, 0x68, 0x1b, 0x10, 0x15, 0x42, 0x48, 0x61, 0x40, 0x62, 0xa0, 0x72, 0x9c, 0x43,
	0x6a, 0xd1, 0xd8, 0xc6, 0xc7, 0x05, 0x7a, 0x03, 0x88, 0x91, 0x4b, 0xe8, 0xe5, 0x74, 0xec, 0x88,
	0x18, 0x2a, 0xd4, 0x2e, 0x5c, 0x06, 0xca, 0x0f, 0xa2, 0xd2, 0xf7, 0x0d, 0x9d, 0x4e, 0x22, 0x9f,
	0xe7, 0xb1, 0xf5, 0xbe, 0x84, 0x6a, 0xa3, 0xac, 0xa2, 0xa8, 0x57, 0xc2, 0x22, 0xb5, 0x86, 0x49,
	0x64, 0xdc, 0x0a, 0x25, 0x91, 0x7e, 0x31, 0x4c, 0xd3, 0xcf, 0xa3, 0x14, 0x2c, 0x1b, 0x51, 0xfb,
	0x35, 0xaa, 0x36, 0xbd, 0x76, 0xca, 0x10, 0xfa, 0xf7, 0xec, 0x52, 0x98, 0x6c, 0xa1, 0x99, 0xb1,
	0x9b, 0x46, 0x91, 0xab, 0x5c, 0xfd, 0xff, 0xaa, 0xb7, 0xfb, 0x8f, 0x2f, 0xd4, 0x1b, 0xf8, 0xb4,
	0x06, 0xb4, 0x0d, 0x35, 0xbe, 0x98, 0x42, 0xad, 0x24, 0x42, 0x8d, 0xdd, 0xfd, 0xd6, 0x22, 0xbd,
	0x02, 0x10, 0x59, 0x0e, 0xde, 0x9c, 0xb4, 0x3f, 0x18, 0x55, 0xf8, 0xee, 0x1d, 0x77, 0x70, 0x6b,
	0x3a, 0xde, 0x1d, 0x42, 0xe7, 0xd7, 0x21, 0x1c, 0xe6, 0xc2, 0x2e, 0xd7, 0x69, 0xc4, 0x55, 0x41,
	0xb9, 0xc2, 0x42, 0x61, 0x33, 0x86, 0x98, 0x7d, 0xa4, 0x76, 0xa3, 0x01, 0xa3, 0x09, 0xe7, 0x93,
	0x2c, 0x33, 0x80, 0x98, 0x54, 0x0a, 0xef, 0x3d, 0xe9, 0x95, 0x73, 0x21, 0x62, 0xbf, 0x55, 0xd9,
	0x9e, 0x35, 0xb6, 0xa7, 0x67, 0x36, 0x0d, 0x06, 0x05, 0x5a, 0x90, 0x1c, 0x5e, 0x4b, 0x38, 0xff,
	0x7d, 0x13, 0xbf, 0xa4, 0xc8, 0x97, 0x50, 0x30, 0x6a, 0x01, 0xed, 0xa2, 0xbe, 0x65, 0x1e, 0x27,
	0xdd, 0xd2, 0x3a, 0x8f, 0xbd, 0xe7, 0xa4, 0xc3, 0x95, 0x90, 0xe8, 0x77, 0x2a, 0xfb, 0x83, 0xc6,
	0x3e, 0xb8, 0xe0, 0xad, 0xb3, 0x92, 0x4b, 0x6a, 0xfc, 0xc9, 0xcd, 0xef, 0xdb, 0xd0, 0xf9, 0xb3,
	0x0d, 0x9d, 0x87, 0x2f, 0xc8, 0x8d, 0x57, 0x98, 0x7b, 0x13, 0x42, 0xea, 0x08, 0xdf, 0x1a, 0xa6,
	0x3d, 0x3f, 0x2a, 0x9b, 0x8b, 0xce, 0xd2, 0x4c, 0xea, 0xd0, 0xfb, 0xb7, 0xaf, 0x39, 0xa9, 0x83,
	0x9d, 0xce, 0x76, 0xc7, 0xc0, 0xdd, 0x1f, 0x03, 0xf7, 0xf7, 0x31, 0x70, 0x7f, 0x9c, 0x02, 0x67,
	0x7f, 0x0a, 0x9c, 0x9f, 0xa7, 0xc0, 0x79, 0x77, 0xbf, 0x50, 0xd9, 0x7a, 0x05, 0xf8, 0xaf, 0x25,
	0x21, 0x2d, 0x18, 0xc9, 0x56, 0x57, 0xeb, 0x4a, 0xbb, 0x55, 0x3d, 0x8f, 0xfe, 0x06, 0x00, 0x00,
	0xff, 0xff, 0xc9, 0xcf, 0xd4, 0xc8, 0x6c, 0x02, 0x00, 0x00,
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
	SplitsWrap(ctx context.Context, in *transactionRequest, opts ...grpc.CallOption) (*transactionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SplitsWrap(ctx context.Context, in *transactionRequest, opts ...grpc.CallOption) (*transactionResponse, error) {
	out := new(transactionResponse)
	err := c.cc.Invoke(ctx, "/base.Msg/splitsWrap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SplitsWrap(context.Context, *transactionRequest) (*transactionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SplitsWrap(ctx context.Context, req *transactionRequest) (*transactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SplitsWrap not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SplitsWrap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(transactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SplitsWrap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.Msg/SplitsWrap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SplitsWrap(ctx, req.(*transactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "base.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "splitsWrap",
			Handler:    _Msg_SplitsWrap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/splits/transactions/wrap/v1beta1/tx.proto",
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
		size := m.Coins.Len()
		i -= size
		if _, err := m.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
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
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	n += 1 + l + sovTx(uint64(l))
	l = m.FromID.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.Coins.Len()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
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
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
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
			if err := m.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
