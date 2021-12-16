// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: check_filtration.proto

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

//赛程
type FiltrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckType int64   `protobuf:"varint,2,opt,name=CheckType,proto3" json:"CheckType,omitempty"` //1：联赛 2：比赛 3：队伍 4：赛季
	Ids       []int64 `protobuf:"varint,3,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Language  string  `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"` //请求语言  1:zh  2:en
}

func (x *FiltrationRequest) Reset() {
	*x = FiltrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_filtration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FiltrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FiltrationRequest) ProtoMessage() {}

func (x *FiltrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_check_filtration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FiltrationRequest.ProtoReflect.Descriptor instead.
func (*FiltrationRequest) Descriptor() ([]byte, []int) {
	return file_check_filtration_proto_rawDescGZIP(), []int{0}
}

func (x *FiltrationRequest) GetCheckType() int64 {
	if x != nil {
		return x.CheckType
	}
	return 0
}

func (x *FiltrationRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *FiltrationRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type FiltrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,3,rep,packed,name=ids,proto3" json:"ids,omitempty"` //返回存在的id数组
}

func (x *FiltrationResponse) Reset() {
	*x = FiltrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_filtration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FiltrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FiltrationResponse) ProtoMessage() {}

func (x *FiltrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_check_filtration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FiltrationResponse.ProtoReflect.Descriptor instead.
func (*FiltrationResponse) Descriptor() ([]byte, []int) {
	return file_check_filtration_proto_rawDescGZIP(), []int{1}
}

func (x *FiltrationResponse) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

var File_check_filtration_proto protoreflect.FileDescriptor

var file_check_filtration_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x26, 0x0a, 0x12, 0x46, 0x69, 0x6c,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64,
	0x73, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_check_filtration_proto_rawDescOnce sync.Once
	file_check_filtration_proto_rawDescData = file_check_filtration_proto_rawDesc
)

func file_check_filtration_proto_rawDescGZIP() []byte {
	file_check_filtration_proto_rawDescOnce.Do(func() {
		file_check_filtration_proto_rawDescData = protoimpl.X.CompressGZIP(file_check_filtration_proto_rawDescData)
	})
	return file_check_filtration_proto_rawDescData
}

var file_check_filtration_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_check_filtration_proto_goTypes = []interface{}{
	(*FiltrationRequest)(nil),  // 0: FiltrationRequest
	(*FiltrationResponse)(nil), // 1: FiltrationResponse
}
var file_check_filtration_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_check_filtration_proto_init() }
func file_check_filtration_proto_init() {
	if File_check_filtration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_check_filtration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FiltrationRequest); i {
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
		file_check_filtration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FiltrationResponse); i {
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
			RawDescriptor: file_check_filtration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_check_filtration_proto_goTypes,
		DependencyIndexes: file_check_filtration_proto_depIdxs,
		MessageInfos:      file_check_filtration_proto_msgTypes,
	}.Build()
	File_check_filtration_proto = out.File
	file_check_filtration_proto_rawDesc = nil
	file_check_filtration_proto_goTypes = nil
	file_check_filtration_proto_depIdxs = nil
}