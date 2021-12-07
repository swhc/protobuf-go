// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: foot_datatype_count.proto

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

type DataTypeCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckType int64   `protobuf:"varint,2,opt,name=checkType,proto3" json:"checkType,omitempty"` //1：联赛 2：比赛 3：队伍 4：赛季
	Ids       []int64 `protobuf:"varint,3,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Language  string  `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"` //请求语言  1:zh  2:en
}

func (x *DataTypeCountRequest) Reset() {
	*x = DataTypeCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_datatype_count_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataTypeCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataTypeCountRequest) ProtoMessage() {}

func (x *DataTypeCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foot_datatype_count_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataTypeCountRequest.ProtoReflect.Descriptor instead.
func (*DataTypeCountRequest) Descriptor() ([]byte, []int) {
	return file_foot_datatype_count_proto_rawDescGZIP(), []int{0}
}

func (x *DataTypeCountRequest) GetCheckType() int64 {
	if x != nil {
		return x.CheckType
	}
	return 0
}

func (x *DataTypeCountRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *DataTypeCountRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type DataTypeCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"` //数量
}

func (x *DataTypeCountResponse) Reset() {
	*x = DataTypeCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_datatype_count_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataTypeCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataTypeCountResponse) ProtoMessage() {}

func (x *DataTypeCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_foot_datatype_count_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataTypeCountResponse.ProtoReflect.Descriptor instead.
func (*DataTypeCountResponse) Descriptor() ([]byte, []int) {
	return file_foot_datatype_count_proto_rawDescGZIP(), []int{1}
}

func (x *DataTypeCountResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_foot_datatype_count_proto protoreflect.FileDescriptor

var file_foot_datatype_count_proto_rawDesc = []byte{
	0x0a, 0x19, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x14, 0x44,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22,
	0x2d, 0x0a, 0x15, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x07,
	0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foot_datatype_count_proto_rawDescOnce sync.Once
	file_foot_datatype_count_proto_rawDescData = file_foot_datatype_count_proto_rawDesc
)

func file_foot_datatype_count_proto_rawDescGZIP() []byte {
	file_foot_datatype_count_proto_rawDescOnce.Do(func() {
		file_foot_datatype_count_proto_rawDescData = protoimpl.X.CompressGZIP(file_foot_datatype_count_proto_rawDescData)
	})
	return file_foot_datatype_count_proto_rawDescData
}

var file_foot_datatype_count_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_foot_datatype_count_proto_goTypes = []interface{}{
	(*DataTypeCountRequest)(nil),  // 0: DataTypeCountRequest
	(*DataTypeCountResponse)(nil), // 1: DataTypeCountResponse
}
var file_foot_datatype_count_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_foot_datatype_count_proto_init() }
func file_foot_datatype_count_proto_init() {
	if File_foot_datatype_count_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foot_datatype_count_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataTypeCountRequest); i {
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
		file_foot_datatype_count_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataTypeCountResponse); i {
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
			RawDescriptor: file_foot_datatype_count_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foot_datatype_count_proto_goTypes,
		DependencyIndexes: file_foot_datatype_count_proto_depIdxs,
		MessageInfos:      file_foot_datatype_count_proto_msgTypes,
	}.Build()
	File_foot_datatype_count_proto = out.File
	file_foot_datatype_count_proto_rawDesc = nil
	file_foot_datatype_count_proto_goTypes = nil
	file_foot_datatype_count_proto_depIdxs = nil
}
