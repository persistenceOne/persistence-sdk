// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/assets/transactions/renumerate/tx.proto

package renumerate

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_test_types "github.com/persistenceOne/persistenceSDK/schema/test_types"
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
	From    github_com_cosmos_cosmos_sdk_types.AccAddress                 `protobuf:"bytes,1,opt,name=from,proto3,customtype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"from"`
	FromID  github_com_persistenceOne_persistenceSDK_schema_test_types.ID `protobuf:"bytes,2,opt,name=fromID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/test_types.ID" json:"fromID"`
	AssetID github_com_persistenceOne_persistenceSDK_schema_test_types.ID `protobuf:"bytes,3,opt,name=assetID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/test_types.ID" json:"assetID"`
}

func (m message) Reset()         { m = message{} }
func (m message) String() string { return proto.CompactTextString(&m) }
func (message) ProtoMessage()    {}
func (*message) Descriptor() ([]byte, []int) {
	return fileDescriptor_63ef39e78dc74a70, []int{0}
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
	proto.RegisterFile("proto/assets/transactions/renumerate/tx.proto", fileDescriptor_63ef39e78dc74a70)
}

var fileDescriptor_63ef39e78dc74a70 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xb1, 0x6a, 0xe3, 0x40,
	0x10, 0x86, 0x25, 0xdb, 0xd8, 0x77, 0xaa, 0x0e, 0x55, 0x3a, 0x17, 0xd2, 0x71, 0xc5, 0x71, 0x8d,
	0xb5, 0x60, 0x73, 0xcd, 0x41, 0x0a, 0x1b, 0x85, 0x20, 0x42, 0x12, 0x50, 0xba, 0x40, 0x30, 0x6b,
	0x69, 0x22, 0x8b, 0x58, 0xbb, 0xca, 0xce, 0x08, 0xe2, 0x37, 0x48, 0x99, 0x27, 0x08, 0x7e, 0x1c,
	0x97, 0x2e, 0x43, 0x0a, 0x13, 0xec, 0x26, 0x8f, 0x11, 0x2c, 0xc9, 0x89, 0x8b, 0x04, 0x5c, 0xa4,
	0x1a, 0x89, 0x9d, 0xff, 0x5b, 0xf6, 0xff, 0x8c, 0x4e, 0xa6, 0x24, 0x49, 0xc6, 0x11, 0x81, 0x90,
	0x91, 0xe2, 0x02, 0x79, 0x48, 0x89, 0x14, 0xc8, 0x14, 0x88, 0x3c, 0x05, 0xc5, 0x09, 0x18, 0xdd,
	0xba, 0xc5, 0x9e, 0xd9, 0x18, 0x71, 0x84, 0xf6, 0x1f, 0x1a, 0x27, 0x2a, 0x1a, 0x66, 0x5c, 0xd1,
	0x94, 0x95, 0x80, 0x58, 0xc6, 0xf2, 0xfd, 0xab, 0xdc, 0x6e, 0x77, 0xf7, 0x82, 0x2b, 0xb8, 0xc9,
	0x01, 0xa9, 0xca, 0xf4, 0xf6, 0xcc, 0x60, 0x26, 0x05, 0x42, 0x19, 0xfa, 0xfd, 0x50, 0x33, 0x5a,
	0x29, 0x20, 0xf2, 0x18, 0x4c, 0xdf, 0x68, 0x5c, 0x29, 0x99, 0x5a, 0xfa, 0x2f, 0xfd, 0xef, 0xf7,
	0xc1, 0xbf, 0xf9, 0xd2, 0xd1, 0x9e, 0x96, 0x4e, 0x27, 0x4e, 0x68, 0x9c, 0x8f, 0xdc, 0x50, 0xa6,
	0x2c, 0x94, 0x98, 0x4a, 0xac, 0x46, 0x07, 0xa3, 0x6b, 0x46, 0xd3, 0x0c, 0xd0, 0xed, 0x87, 0x61,
	0x3f, 0x8a, 0x14, 0x20, 0x06, 0x05, 0xc2, 0xbc, 0x34, 0x9a, 0x9b, 0xe9, 0x7b, 0x56, 0xad, 0x80,
	0x1d, 0x56, 0xb0, 0x83, 0x1d, 0x58, 0x06, 0x0a, 0x13, 0x24, 0x10, 0x21, 0x9c, 0x09, 0xd8, 0xfd,
	0x3d, 0xf7, 0x8e, 0x19, 0x86, 0x63, 0x48, 0x39, 0x23, 0x40, 0x1a, 0x96, 0x97, 0xf8, 0x5e, 0x50,
	0x41, 0xcd, 0xa1, 0xd1, 0x2a, 0x9e, 0xe9, 0x7b, 0x56, 0xfd, 0x2b, 0xf9, 0x5b, 0xea, 0xff, 0x6f,
	0x77, 0x33, 0x47, 0x7b, 0x99, 0x39, 0x5a, 0xf7, 0xd4, 0xa8, 0x9f, 0x60, 0x6c, 0x1e, 0x19, 0x3f,
	0xca, 0x62, 0x83, 0xb7, 0x2a, 0x4d, 0xcb, 0xdd, 0x38, 0x75, 0x77, 0x9a, 0x0e, 0x4a, 0x21, 0xed,
	0x9f, 0x1f, 0x9c, 0x94, 0xb5, 0x0f, 0xfc, 0xf9, 0xca, 0xd6, 0x17, 0x2b, 0x5b, 0x7f, 0x5e, 0xd9,
	0xfa, 0xfd, 0xda, 0xd6, 0x16, 0x6b, 0x5b, 0x7b, 0x5c, 0xdb, 0xda, 0x05, 0x4b, 0x65, 0x94, 0x4f,
	0x00, 0xb7, 0x06, 0x13, 0x41, 0xa0, 0x04, 0x9f, 0x7c, 0xa6, 0x72, 0xd4, 0x2c, 0x14, 0xf6, 0x5e,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xee, 0x24, 0xf3, 0x8f, 0x8a, 0x02, 0x00, 0x00,
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
	AssetsRenumerate(ctx context.Context, in *transactionRequest, opts ...grpc.CallOption) (*transactionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AssetsRenumerate(ctx context.Context, in *transactionRequest, opts ...grpc.CallOption) (*transactionResponse, error) {
	out := new(transactionResponse)
	err := c.cc.Invoke(ctx, "/base.Msg/assetsRenumerate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	AssetsRenumerate(context.Context, *transactionRequest) (*transactionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AssetsRenumerate(ctx context.Context, req *transactionRequest) (*transactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssetsRenumerate not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AssetsRenumerate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(transactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AssetsRenumerate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.Msg/AssetsRenumerate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AssetsRenumerate(ctx, req.(*transactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "base.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "assetsRenumerate",
			Handler:    _Msg_AssetsRenumerate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/assets/transactions/renumerate/tx.proto",
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
