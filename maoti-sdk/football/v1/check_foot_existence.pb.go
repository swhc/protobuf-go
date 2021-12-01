// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: check_foot_existence.proto

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
type CheckFootExistenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckType int64   `protobuf:"varint,2,opt,name=CheckType,proto3" json:"CheckType,omitempty"` //1：联赛 2：比赛 3：队伍 4：赛季
	Ids       []int64 `protobuf:"varint,3,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *CheckFootExistenceRequest) Reset() {
	*x = CheckFootExistenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_foot_existence_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckFootExistenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckFootExistenceRequest) ProtoMessage() {}

func (x *CheckFootExistenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_check_foot_existence_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckFootExistenceRequest.ProtoReflect.Descriptor instead.
func (*CheckFootExistenceRequest) Descriptor() ([]byte, []int) {
	return file_check_foot_existence_proto_rawDescGZIP(), []int{0}
}

func (x *CheckFootExistenceRequest) GetCheckType() int64 {
	if x != nil {
		return x.CheckType
	}
	return 0
}

func (x *CheckFootExistenceRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type CheckFootExistenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOK bool `protobuf:"varint,1,opt,name=IsOK,proto3" json:"IsOK,omitempty"` //总数
}

func (x *CheckFootExistenceResponse) Reset() {
	*x = CheckFootExistenceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_foot_existence_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckFootExistenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckFootExistenceResponse) ProtoMessage() {}

func (x *CheckFootExistenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_check_foot_existence_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckFootExistenceResponse.ProtoReflect.Descriptor instead.
func (*CheckFootExistenceResponse) Descriptor() ([]byte, []int) {
	return file_check_foot_existence_proto_rawDescGZIP(), []int{1}
}

func (x *CheckFootExistenceResponse) GetIsOK() bool {
	if x != nil {
		return x.IsOK
	}
	return false
}

var File_check_foot_existence_proto protoreflect.FileDescriptor

var file_check_foot_existence_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x65, 0x78, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x19,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x6f, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x30, 0x0a, 0x1a, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x46, 0x6f, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x73, 0x4f, 0x4b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x49, 0x73, 0x4f, 0x4b, 0x42, 0x07, 0x5a, 0x05, 0x2e,
	0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_check_foot_existence_proto_rawDescOnce sync.Once
	file_check_foot_existence_proto_rawDescData = file_check_foot_existence_proto_rawDesc
)

func file_check_foot_existence_proto_rawDescGZIP() []byte {
	file_check_foot_existence_proto_rawDescOnce.Do(func() {
		file_check_foot_existence_proto_rawDescData = protoimpl.X.CompressGZIP(file_check_foot_existence_proto_rawDescData)
	})
	return file_check_foot_existence_proto_rawDescData
}

var file_check_foot_existence_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_check_foot_existence_proto_goTypes = []interface{}{
	(*CheckFootExistenceRequest)(nil),  // 0: CheckFootExistenceRequest
	(*CheckFootExistenceResponse)(nil), // 1: CheckFootExistenceResponse
}
var file_check_foot_existence_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_check_foot_existence_proto_init() }
func file_check_foot_existence_proto_init() {
	if File_check_foot_existence_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_check_foot_existence_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckFootExistenceRequest); i {
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
		file_check_foot_existence_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckFootExistenceResponse); i {
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
			RawDescriptor: file_check_foot_existence_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_check_foot_existence_proto_goTypes,
		DependencyIndexes: file_check_foot_existence_proto_depIdxs,
		MessageInfos:      file_check_foot_existence_proto_msgTypes,
	}.Build()
	File_check_foot_existence_proto = out.File
	file_check_foot_existence_proto_rawDesc = nil
	file_check_foot_existence_proto_goTypes = nil
	file_check_foot_existence_proto_depIdxs = nil
}