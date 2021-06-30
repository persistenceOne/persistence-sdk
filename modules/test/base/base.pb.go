// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: test_proto/base.proto

package base

import (
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	// Types that are assignable to Req:
	//	*TokenRequest_SendToken
	//	*TokenRequest_RequestToken
	Req isTokenRequest_Req `protobuf_oneof:"req"`
}

func (x *TokenRequest) Reset() {
	*x = TokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRequest) ProtoMessage() {}

func (x *TokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRequest.ProtoReflect.Descriptor instead.
func (*TokenRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_base_proto_rawDescGZIP(), []int{0}
}

func (x *TokenRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (m *TokenRequest) GetReq() isTokenRequest_Req {
	if m != nil {
		return m.Req
	}
	return nil
}

func (x *TokenRequest) GetSendToken() *SendToken {
	if x, ok := x.GetReq().(*TokenRequest_SendToken); ok {
		return x.SendToken
	}
	return nil
}

func (x *TokenRequest) GetRequestToken() *RequestToken {
	if x, ok := x.GetReq().(*TokenRequest_RequestToken); ok {
		return x.RequestToken
	}
	return nil
}

type isTokenRequest_Req interface {
	isTokenRequest_Req()
}

type TokenRequest_SendToken struct {
	SendToken *SendToken `protobuf:"bytes,2,opt,name=sendToken,proto3,oneof"`
}

type TokenRequest_RequestToken struct {
	RequestToken *RequestToken `protobuf:"bytes,3,opt,name=requestToken,proto3,oneof"`
}

func (*TokenRequest_SendToken) isTokenRequest_Req() {}

func (*TokenRequest_RequestToken) isTokenRequest_Req() {}

var File_test_proto_base_proto protoreflect.FileDescriptor

var file_test_proto_base_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x26, 0x74, 0x68, 0x69, 0x72,
	0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x0c, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x3c, 0x0a, 0x09,
	0x73, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x48, 0x00, 0x52,
	0x09, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x45, 0x0a, 0x0c, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x3a, 0x0c, 0x88, 0xa0, 0x1f, 0x00, 0x98, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x42,
	0x05, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x42, 0x13, 0x5a, 0x11, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_test_proto_base_proto_rawDescOnce sync.Once
	file_test_proto_base_proto_rawDescData = file_test_proto_base_proto_rawDesc
)

func file_test_proto_base_proto_rawDescGZIP() []byte {
	file_test_proto_base_proto_rawDescOnce.Do(func() {
		file_test_proto_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_base_proto_rawDescData)
	})
	return file_test_proto_base_proto_rawDescData
}

var file_test_proto_base_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_test_proto_base_proto_goTypes = []interface{}{
	(*TokenRequest)(nil), // 0: test.base.v1beta1.tokenRequest
	(*SendToken)(nil),    // 1: test.base.v1beta1.sendToken
	(*RequestToken)(nil), // 2: test.base.v1beta1.requestToken
}
var file_test_proto_base_proto_depIdxs = []int32{
	1, // 0: test.base.v1beta1.tokenRequest.sendToken:type_name -> test.base.v1beta1.sendToken
	2, // 1: test.base.v1beta1.tokenRequest.requestToken:type_name -> test.base.v1beta1.requestToken
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_test_proto_base_proto_init() }
func file_test_proto_base_proto_init() {
	if File_test_proto_base_proto != nil {
		return
	}
	file_test_proto_test_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_test_proto_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_test_proto_base_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TokenRequest_SendToken)(nil),
		(*TokenRequest_RequestToken)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_proto_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_proto_base_proto_goTypes,
		DependencyIndexes: file_test_proto_base_proto_depIdxs,
		MessageInfos:      file_test_proto_base_proto_msgTypes,
	}.Build()
	File_test_proto_base_proto = out.File
	file_test_proto_base_proto_rawDesc = nil
	file_test_proto_base_proto_goTypes = nil
	file_test_proto_base_proto_depIdxs = nil
}
