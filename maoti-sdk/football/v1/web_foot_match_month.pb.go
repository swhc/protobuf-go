// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: web_foot_match_month.proto

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

//指数-列表
type WebFootMatchMonthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId int64 `protobuf:"varint,1,opt,name=teamId,proto3" json:"teamId,omitempty"` //联赛id
}

func (x *WebFootMatchMonthRequest) Reset() {
	*x = WebFootMatchMonthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_match_month_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootMatchMonthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootMatchMonthRequest) ProtoMessage() {}

func (x *WebFootMatchMonthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_match_month_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootMatchMonthRequest.ProtoReflect.Descriptor instead.
func (*WebFootMatchMonthRequest) Descriptor() ([]byte, []int) {
	return file_web_foot_match_month_proto_rawDescGZIP(), []int{0}
}

func (x *WebFootMatchMonthRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

type WebFootMatchMonthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []string `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *WebFootMatchMonthResponse) Reset() {
	*x = WebFootMatchMonthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_match_month_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootMatchMonthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootMatchMonthResponse) ProtoMessage() {}

func (x *WebFootMatchMonthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_match_month_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootMatchMonthResponse.ProtoReflect.Descriptor instead.
func (*WebFootMatchMonthResponse) Descriptor() ([]byte, []int) {
	return file_web_foot_match_month_proto_rawDescGZIP(), []int{1}
}

func (x *WebFootMatchMonthResponse) GetList() []string {
	if x != nil {
		return x.List
	}
	return nil
}

var File_web_foot_match_month_proto protoreflect.FileDescriptor

var file_web_foot_match_month_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x77, 0x65, 0x62, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x18,
	0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x6f, 0x6e, 0x74,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x61, 0x6d,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64,
	0x22, 0x2f, 0x0a, 0x19, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_web_foot_match_month_proto_rawDescOnce sync.Once
	file_web_foot_match_month_proto_rawDescData = file_web_foot_match_month_proto_rawDesc
)

func file_web_foot_match_month_proto_rawDescGZIP() []byte {
	file_web_foot_match_month_proto_rawDescOnce.Do(func() {
		file_web_foot_match_month_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_foot_match_month_proto_rawDescData)
	})
	return file_web_foot_match_month_proto_rawDescData
}

var file_web_foot_match_month_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_web_foot_match_month_proto_goTypes = []interface{}{
	(*WebFootMatchMonthRequest)(nil),  // 0: WebFootMatchMonthRequest
	(*WebFootMatchMonthResponse)(nil), // 1: WebFootMatchMonthResponse
}
var file_web_foot_match_month_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_web_foot_match_month_proto_init() }
func file_web_foot_match_month_proto_init() {
	if File_web_foot_match_month_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_foot_match_month_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootMatchMonthRequest); i {
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
		file_web_foot_match_month_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootMatchMonthResponse); i {
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
			RawDescriptor: file_web_foot_match_month_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_foot_match_month_proto_goTypes,
		DependencyIndexes: file_web_foot_match_month_proto_depIdxs,
		MessageInfos:      file_web_foot_match_month_proto_msgTypes,
	}.Build()
	File_web_foot_match_month_proto = out.File
	file_web_foot_match_month_proto_rawDesc = nil
	file_web_foot_match_month_proto_goTypes = nil
	file_web_foot_match_month_proto_depIdxs = nil
}
