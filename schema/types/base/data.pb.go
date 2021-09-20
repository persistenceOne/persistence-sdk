// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/schema/types/base/data.proto

package base

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	"github.com/persistenceOne/persistenceSDK/schema/types"
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

type Data struct {
	// Types that are valid to be assigned to Data:
	//	*Data_DecData
	//	*Data_HeightData
	//	*Data_IdData
	//	*Data_StringData
	//	*Data_AccAddressData
	//	*Data_ListData
	Data types.Data `protobuf_oneof:"data"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return m.Data.String() }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b342118f875caf9, []int{0}
}
func (m *Data) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Data.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return m.Size()
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

type isData_Data interface {
	isData_Data()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Data_DecData struct {
	DecData *DecData `protobuf:"bytes,1,opt,name=dec_data,json=decData,proto3,oneof" json:"dec_data,omitempty"`
}
type Data_HeightData struct {
	HeightData *HeightData `protobuf:"bytes,2,opt,name=height_data,json=heightData,proto3,oneof" json:"height_data,omitempty"`
}
type Data_IdData struct {
	IdData *IDData `protobuf:"bytes,3,opt,name=id_data,json=idData,proto3,oneof" json:"id_data,omitempty"`
}
type Data_StringData struct {
	StringData *StringData `protobuf:"bytes,4,opt,name=string_data,json=stringData,proto3,oneof" json:"string_data,omitempty"`
}
type Data_AccAddressData struct {
	AccAddressData *AccAddressData `protobuf:"bytes,5,opt,name=acc_address_data,json=accAddressData,proto3,oneof" json:"acc_address_data,omitempty"`
}
type Data_ListData struct {
	ListData ListData `protobuf:"bytes,6,opt,name=list_data,json=listData,proto3,oneof,customtype=github.com/persistenceOne/persistenceSDK/schema/types/base.ListData" json:"list_data,omitempty"`
}

func (*Data_DecData) isData_Data()        {}
func (*Data_HeightData) isData_Data()     {}
func (*Data_IdData) isData_Data()         {}
func (*Data_StringData) isData_Data()     {}
func (*Data_AccAddressData) isData_Data() {}
func (*Data_ListData) isData_Data()       {}

func (m *Data) GetData() types.Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Data) GetDecData() *DecData {
	if x, ok := m.GetData().(*Data_DecData); ok {
		return x.DecData
	}
	return nil
}

func (m *Data) GetHeightData() *HeightData {
	if x, ok := m.GetData().(*Data_HeightData); ok {
		return x.HeightData
	}
	return nil
}

func (m *Data) GetIdData() *IDData {
	if x, ok := m.GetData().(*Data_IdData); ok {
		return x.IdData
	}
	return nil
}

func (m *Data) GetStringData() *StringData {
	if x, ok := m.GetData().(*Data_StringData); ok {
		return x.StringData
	}
	return nil
}

func (m *Data) GetAccAddressData() *AccAddressData {
	if x, ok := m.GetData().(*Data_AccAddressData); ok {
		return x.AccAddressData
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Data) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Data_DecData)(nil),
		(*Data_HeightData)(nil),
		(*Data_IdData)(nil),
		(*Data_StringData)(nil),
		(*Data_AccAddressData)(nil),
		(*Data_ListData)(nil),
	}
}

func init() {
	proto.RegisterType((*Data)(nil), "persistence_sdk.schema.types.base.Data")
}

func init() {
	proto.RegisterFile("persistence_sdk/schema/types/base/data.proto", fileDescriptor_9b342118f875caf9)
}

var fileDescriptor_9b342118f875caf9 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0x5b, 0xad, 0xdd, 0x96, 0x81, 0x48, 0xf1, 0x20, 0x3b, 0xd4, 0xe9, 0x69, 0x8a, 0xa6,
	0x38, 0xc1, 0x83, 0xb7, 0xcd, 0xc2, 0x2a, 0x0a, 0xca, 0xe6, 0x49, 0x90, 0x91, 0x25, 0xb1, 0x0d,
	0x6e, 0xeb, 0x68, 0xe2, 0xc1, 0x7f, 0xe0, 0xd1, 0x9f, 0xe5, 0x71, 0x47, 0xf1, 0x20, 0xb2, 0xfd,
	0x11, 0x49, 0x53, 0xd7, 0x4e, 0x0f, 0xad, 0xde, 0x12, 0xf8, 0xde, 0xe7, 0x0d, 0x4f, 0x3e, 0x70,
	0x30, 0xa1, 0x11, 0x67, 0x5c, 0xd0, 0x31, 0xa6, 0x7d, 0x4e, 0x1e, 0x1c, 0x8e, 0x03, 0x3a, 0x42,
	0x8e, 0x78, 0x9a, 0x50, 0xee, 0x0c, 0x10, 0xa7, 0x0e, 0x41, 0x02, 0xc1, 0x49, 0x14, 0x8a, 0xd0,
	0xda, 0xf9, 0x31, 0x0d, 0xd5, 0x34, 0x8c, 0xa7, 0xa1, 0x9c, 0xae, 0x6d, 0xfa, 0xa1, 0x1f, 0xc6,
	0xd3, 0x8e, 0x3c, 0xa9, 0x60, 0xcd, 0x29, 0x50, 0x43, 0xb1, 0xbb, 0x68, 0xaa, 0x35, 0xf3, 0x03,
	0x01, 0x65, 0x7e, 0x20, 0x32, 0x19, 0x98, 0x9f, 0x61, 0xe4, 0x6f, 0x1d, 0x5c, 0x44, 0x6c, 0xec,
	0x67, 0x32, 0x27, 0xf9, 0x19, 0x84, 0x71, 0x8b, 0x90, 0x88, 0x72, 0x9e, 0xe6, 0x76, 0x9f, 0x0d,
	0x60, 0xc8, 0xab, 0xd5, 0x01, 0x65, 0x42, 0x71, 0x5f, 0x4a, 0xdd, 0xd2, 0xeb, 0x7a, 0xa3, 0xda,
	0xdc, 0x87, 0xb9, 0x56, 0xa1, 0xab, 0xe4, 0x78, 0x5a, 0xb7, 0x94, 0x78, 0xb2, 0xae, 0x41, 0x55,
	0x19, 0x50, 0xac, 0x95, 0x98, 0x75, 0x58, 0x80, 0xe5, 0x2d, 0xbc, 0x79, 0x5a, 0x17, 0xa4, 0x16,
	0x2d, 0x17, 0x94, 0x18, 0x51, 0xb4, 0xd5, 0x98, 0xb6, 0x57, 0x80, 0x76, 0xee, 0x26, 0x24, 0x53,
	0xb9, 0x95, 0xef, 0x52, 0xd6, 0x14, 0xc9, 0x28, 0xfc, 0xae, 0xde, 0xc2, 0xb5, 0x7c, 0x57, 0x6a,
	0xde, 0xba, 0x03, 0x1b, 0x08, 0xe3, 0x3e, 0x52, 0x52, 0x15, 0x76, 0x2d, 0xc6, 0x1e, 0x15, 0xc0,
	0xb6, 0x96, 0xbe, 0xc3, 0xd3, 0xba, 0xeb, 0xcb, 0x1f, 0x64, 0xdd, 0x83, 0xca, 0x90, 0xf1, 0x44,
	0xa3, 0x59, 0xd7, 0x1b, 0x95, 0x76, 0xe7, 0xfd, 0x63, 0xfb, 0xcc, 0x67, 0x22, 0x78, 0x1c, 0x40,
	0x1c, 0x8e, 0xb2, 0xdb, 0x7b, 0x35, 0xa6, 0xd9, 0x6b, 0xcf, 0xbd, 0xf8, 0xbd, 0x02, 0xf0, 0x92,
	0xf1, 0x6f, 0xc1, 0xe5, 0x61, 0x72, 0x6e, 0x9b, 0xc0, 0x90, 0x15, 0xed, 0x9b, 0xd7, 0x99, 0xad,
	0x4f, 0x67, 0xb6, 0xfe, 0x39, 0xb3, 0xf5, 0x97, 0xb9, 0xad, 0x4d, 0xe7, 0xb6, 0xf6, 0x36, 0xb7,
	0xb5, 0xdb, 0xd3, 0xff, 0x57, 0x0e, 0xcc, 0x78, 0xcf, 0x8e, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xfc, 0xd5, 0x1c, 0xdb, 0xd1, 0x03, 0x00, 0x00,
}

func (m *Data) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Data) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Data) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		{
			size := m.Data.Size()
			i -= size
			if _, err := m.Data.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Data_DecData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Data_DecData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DecData != nil {
		{
			size, err := m.DecData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Data_HeightData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Data_HeightData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HeightData != nil {
		{
			size, err := m.HeightData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Data_IdData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Data_IdData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.IdData != nil {
		{
			size, err := m.IdData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *Data_StringData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Data_StringData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.StringData != nil {
		{
			size, err := m.StringData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *Data_AccAddressData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Data_AccAddressData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.AccAddressData != nil {
		{
			size, err := m.AccAddressData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *Data_ListData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Data_ListData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	{
		size := m.ListData.Size()
		i -= size
		if _, err := m.ListData.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintData(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	return len(dAtA) - i, nil
}
func encodeVarintData(dAtA []byte, offset int, v uint64) int {
	offset -= sovData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Data) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		n += m.Data.Size()
	}
	return n
}

func (m *Data_DecData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DecData != nil {
		l = m.DecData.Size()
		n += 1 + l + sovData(uint64(l))
	}
	return n
}
func (m *Data_HeightData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HeightData != nil {
		l = m.HeightData.Size()
		n += 1 + l + sovData(uint64(l))
	}
	return n
}
func (m *Data_IdData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IdData != nil {
		l = m.IdData.Size()
		n += 1 + l + sovData(uint64(l))
	}
	return n
}
func (m *Data_StringData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StringData != nil {
		l = m.StringData.Size()
		n += 1 + l + sovData(uint64(l))
	}
	return n
}
func (m *Data_AccAddressData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AccAddressData != nil {
		l = m.AccAddressData.Size()
		n += 1 + l + sovData(uint64(l))
	}
	return n
}
func (m *Data_ListData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ListData.Size()
	n += 1 + l + sovData(uint64(l))
	return n
}

func sovData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozData(x uint64) (n int) {
	return sovData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Data) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowData
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
			return fmt.Errorf("proto: Data: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Data: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DecData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DecData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Data = &Data_DecData{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeightData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &HeightData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Data = &Data_HeightData{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &IDData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Data = &Data_IdData{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &StringData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Data = &Data_StringData{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccAddressData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &AccAddressData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Data = &Data_AccAddressData{v}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var vv ListData
			v := &vv
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Data = &Data_ListData{*v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthData
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
func skipData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowData
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
					return 0, ErrIntOverflowData
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
					return 0, ErrIntOverflowData
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
				return 0, ErrInvalidLengthData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupData = fmt.Errorf("proto: unexpected end of group")
)