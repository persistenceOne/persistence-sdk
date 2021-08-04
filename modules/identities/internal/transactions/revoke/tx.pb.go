// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/identities/transactions/revoke/tx.proto

package revoke

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
	From             github_com_cosmos_cosmos_sdk_types.AccAddress                 `protobuf:"bytes,1,opt,name=from,proto3,customtype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"from"  valid:"required~required field from missing"`
	FromID           github_com_persistenceOne_persistenceSDK_schema_test_types.ID `protobuf:"bytes,2,opt,name=fromID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/test_types.ID" json:"fromID"  valid:"required~required field fromID missing"`
	ToID             github_com_persistenceOne_persistenceSDK_schema_test_types.ID `protobuf:"bytes,3,opt,name=toID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/test_types.ID" json:"toID"  valid:"required~required field toID missing"`
	ClassificationID github_com_persistenceOne_persistenceSDK_schema_test_types.ID `protobuf:"bytes,4,opt,name=classificationID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/test_types.ID" json:"classificationID"  valid:"required~required field classificationID missing"`
}

func (m message) Reset()         { m = message{} }
func (m message) String() string { return proto.CompactTextString(&m) }
func (message) ProtoMessage()    {}
func (*message) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dfde099f3f0c40e, []int{0}
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
	proto.RegisterFile("proto/identities/transactions/revoke/tx.proto", fileDescriptor_6dfde099f3f0c40e)
}

var fileDescriptor_6dfde099f3f0c40e = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x13, 0x1b, 0x56, 0x9d, 0xd3, 0x92, 0x53, 0xec, 0x21, 0x11, 0x0f, 0xe2, 0xa5, 0x19,
	0xd8, 0xc5, 0x8b, 0xe0, 0x61, 0x97, 0x88, 0x04, 0x51, 0x21, 0x9e, 0x14, 0x64, 0x99, 0x4e, 0x5e,
	0xd3, 0xa1, 0xcd, 0x4c, 0x3a, 0xef, 0x55, 0xec, 0x37, 0x10, 0xbc, 0xf8, 0x11, 0xfa, 0x71, 0x7a,
	0xec, 0x51, 0x3c, 0x14, 0x69, 0x2f, 0x7e, 0x0c, 0x69, 0x12, 0x68, 0x60, 0x7b, 0xe8, 0xa1, 0xa7,
	0x37, 0xc3, 0xbc, 0xf7, 0x7b, 0x30, 0xbf, 0x3f, 0x1b, 0x54, 0xd6, 0x90, 0xe1, 0x2a, 0x07, 0x4d,
	0x8a, 0x14, 0x20, 0x27, 0x2b, 0x34, 0x0a, 0x49, 0xca, 0x68, 0xe4, 0x16, 0xbe, 0x99, 0x09, 0x70,
	0xfa, 0x1e, 0xd7, 0x7d, 0xbe, 0x37, 0x14, 0x08, 0xfd, 0xe7, 0x34, 0x56, 0x36, 0xbf, 0xab, 0x84,
	0xa5, 0x05, 0x6f, 0x00, 0x85, 0x29, 0xcc, 0xe1, 0xd4, 0x74, 0xf7, 0xaf, 0x4e, 0x82, 0x5b, 0x98,
	0xcd, 0x01, 0xa9, 0x9d, 0xb9, 0x3e, 0x71, 0x06, 0x2b, 0xa3, 0x11, 0x9a, 0xa1, 0x67, 0x3f, 0x7b,
	0xec, 0x61, 0x09, 0x88, 0xa2, 0x00, 0x3f, 0x65, 0xde, 0xc8, 0x9a, 0x32, 0x70, 0x9f, 0xba, 0x2f,
	0x1e, 0xdf, 0xbe, 0x5c, 0x6d, 0x22, 0xe7, 0xcf, 0x26, 0x1a, 0x14, 0x8a, 0xc6, 0xf3, 0x61, 0x2c,
	0x4d, 0xc9, 0xa5, 0xc1, 0xd2, 0x60, 0x5b, 0x06, 0x98, 0x4f, 0x38, 0x2d, 0x2a, 0xc0, 0xf8, 0x46,
	0xca, 0x9b, 0x3c, 0xb7, 0x80, 0x98, 0xd5, 0x08, 0xff, 0x2b, 0xbb, 0xd8, 0xd7, 0x34, 0x09, 0x1e,
	0xd4, 0xb0, 0x37, 0x2d, 0xec, 0x75, 0x07, 0x56, 0x81, 0x45, 0x85, 0x04, 0x5a, 0xc2, 0x47, 0x0d,
	0xdd, 0xeb, 0xa7, 0xe4, 0x1d, 0x47, 0x39, 0x86, 0x52, 0x70, 0x02, 0xa4, 0xbb, 0x66, 0x49, 0x9a,
	0x64, 0x2d, 0xd4, 0xff, 0xcc, 0x3c, 0x32, 0x69, 0x12, 0xf4, 0xce, 0x09, 0xaf, 0x91, 0xfe, 0x8c,
	0x5d, 0xca, 0xa9, 0x40, 0x54, 0x23, 0x25, 0xc5, 0xfe, 0xe7, 0xd2, 0x24, 0xf0, 0xce, 0xb9, 0xe6,
	0x1e, 0xfe, 0xd5, 0xa3, 0x1f, 0xcb, 0xc8, 0xf9, 0xb7, 0x8c, 0x9c, 0xab, 0x0f, 0xac, 0xf7, 0x1e,
	0x0b, 0xff, 0x2d, 0xbb, 0x3c, 0x58, 0xcc, 0x6a, 0x6f, 0x7e, 0x10, 0xef, 0x03, 0x14, 0x77, 0x94,
	0x66, 0x8d, 0xfd, 0xfe, 0x93, 0x23, 0x2f, 0x8d, 0xe3, 0xdb, 0x74, 0xb5, 0x0d, 0xdd, 0xf5, 0x36,
	0x74, 0xff, 0x6e, 0x43, 0xf7, 0xd7, 0x2e, 0x74, 0xd6, 0xbb, 0xd0, 0xf9, 0xbd, 0x0b, 0x9d, 0x2f,
	0xbc, 0x34, 0xf9, 0x7c, 0x0a, 0xd8, 0x8d, 0x8b, 0xd2, 0x04, 0x56, 0x8b, 0xe9, 0xb1, 0xdc, 0x0c,
	0x2f, 0xea, 0xbc, 0x5c, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x65, 0x1e, 0x18, 0xf7, 0x02,
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
	IdentitiesRevoke(ctx context.Context, in *transactionRequest, opts ...grpc.CallOption) (*transactionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) IdentitiesRevoke(ctx context.Context, in *transactionRequest, opts ...grpc.CallOption) (*transactionResponse, error) {
	out := new(transactionResponse)
	err := c.cc.Invoke(ctx, "/base.Msg/identitiesRevoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	IdentitiesRevoke(context.Context, *transactionRequest) (*transactionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) IdentitiesRevoke(ctx context.Context, req *transactionRequest) (*transactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IdentitiesRevoke not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_IdentitiesRevoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(transactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).IdentitiesRevoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.Msg/IdentitiesRevoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).IdentitiesRevoke(ctx, req.(*transactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "base.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "identitiesRevoke",
			Handler:    _Msg_IdentitiesRevoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/identities/transactions/revoke/tx.proto",
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
		size := m.ClassificationID.Size()
		i -= size
		if _, err := m.ClassificationID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.ToID.Size()
		i -= size
		if _, err := m.ToID.MarshalTo(dAtA[i:]); err != nil {
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
	l = m.ToID.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.ClassificationID.Size()
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
			if err := m.ToID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationID", wireType)
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
			if err := m.ClassificationID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
