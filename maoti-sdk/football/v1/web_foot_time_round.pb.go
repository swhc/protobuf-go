// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: web_foot_time_round.proto

package v1

import (
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

type WebFootTimeRoundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language string `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"` //请求语言
	ComId    int64  `protobuf:"varint,2,opt,name=comId,proto3" json:"comId,omitempty"`      //联赛id
}

func (x *WebFootTimeRoundRequest) Reset() {
	*x = WebFootTimeRoundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_time_round_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootTimeRoundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootTimeRoundRequest) ProtoMessage() {}

func (x *WebFootTimeRoundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_time_round_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootTimeRoundRequest.ProtoReflect.Descriptor instead.
func (*WebFootTimeRoundRequest) Descriptor() ([]byte, []int) {
	return file_web_foot_time_round_proto_rawDescGZIP(), []int{0}
}

func (x *WebFootTimeRoundRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *WebFootTimeRoundRequest) GetComId() int64 {
	if x != nil {
		return x.ComId
	}
	return 0
}

var File_web_foot_time_round_proto protoreflect.FileDescriptor

var file_web_foot_time_round_proto_rawDesc = []byte{
	0x0a, 0x19, 0x77, 0x65, 0x62, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x17, 0x57,
	0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x63, 0x6f, 0x6d, 0x49, 0x64, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web_foot_time_round_proto_rawDescOnce sync.Once
	file_web_foot_time_round_proto_rawDescData = file_web_foot_time_round_proto_rawDesc
)

func file_web_foot_time_round_proto_rawDescGZIP() []byte {
	file_web_foot_time_round_proto_rawDescOnce.Do(func() {
		file_web_foot_time_round_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_foot_time_round_proto_rawDescData)
	})
	return file_web_foot_time_round_proto_rawDescData
}

var file_web_foot_time_round_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_web_foot_time_round_proto_goTypes = []interface{}{
	(*WebFootTimeRoundRequest)(nil), // 0: WebFootTimeRoundRequest
}
var file_web_foot_time_round_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_web_foot_time_round_proto_init() }
func file_web_foot_time_round_proto_init() {
	if File_web_foot_time_round_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_foot_time_round_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootTimeRoundRequest); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_web_foot_time_round_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_foot_time_round_proto_goTypes,
		DependencyIndexes: file_web_foot_time_round_proto_depIdxs,
		MessageInfos:      file_web_foot_time_round_proto_msgTypes,
	}.Build()
	File_web_foot_time_round_proto = out.File
	file_web_foot_time_round_proto_rawDesc = nil
	file_web_foot_time_round_proto_goTypes = nil
	file_web_foot_time_round_proto_depIdxs = nil
}
