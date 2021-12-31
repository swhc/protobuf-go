// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: collect_match.proto

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

type CollectMatchRequestParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchIds    []int64 `protobuf:"varint,1,rep,packed,name=MatchIds,proto3" json:"MatchIds,omitempty"`
	DeviceToken string  `protobuf:"bytes,2,opt,name=DeviceToken,proto3" json:"DeviceToken,omitempty"`
	ActionType  int64   `protobuf:"varint,3,opt,name=ActionType,proto3" json:"ActionType,omitempty"` // 1-收藏 2-取消收藏
}

func (x *CollectMatchRequestParam) Reset() {
	*x = CollectMatchRequestParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_match_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectMatchRequestParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectMatchRequestParam) ProtoMessage() {}

func (x *CollectMatchRequestParam) ProtoReflect() protoreflect.Message {
	mi := &file_collect_match_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectMatchRequestParam.ProtoReflect.Descriptor instead.
func (*CollectMatchRequestParam) Descriptor() ([]byte, []int) {
	return file_collect_match_proto_rawDescGZIP(), []int{0}
}

func (x *CollectMatchRequestParam) GetMatchIds() []int64 {
	if x != nil {
		return x.MatchIds
	}
	return nil
}

func (x *CollectMatchRequestParam) GetDeviceToken() string {
	if x != nil {
		return x.DeviceToken
	}
	return ""
}

func (x *CollectMatchRequestParam) GetActionType() int64 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type PushMsgMatchRequestParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchId   int64  `protobuf:"varint,1,opt,name=MatchId,proto3" json:"MatchId,omitempty"`     // 比赛自增id
	SportType int64  `protobuf:"varint,2,opt,name=SportType,proto3" json:"SportType,omitempty"` // 比赛类型 1-足球 2-篮球
	Title     string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`          // 标题
	Content   string `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`      // 内容
}

func (x *PushMsgMatchRequestParam) Reset() {
	*x = PushMsgMatchRequestParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_match_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsgMatchRequestParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgMatchRequestParam) ProtoMessage() {}

func (x *PushMsgMatchRequestParam) ProtoReflect() protoreflect.Message {
	mi := &file_collect_match_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgMatchRequestParam.ProtoReflect.Descriptor instead.
func (*PushMsgMatchRequestParam) Descriptor() ([]byte, []int) {
	return file_collect_match_proto_rawDescGZIP(), []int{1}
}

func (x *PushMsgMatchRequestParam) GetMatchId() int64 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

func (x *PushMsgMatchRequestParam) GetSportType() int64 {
	if x != nil {
		return x.SportType
	}
	return 0
}

func (x *PushMsgMatchRequestParam) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PushMsgMatchRequestParam) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_collect_match_proto protoreflect.FileDescriptor

var file_collect_match_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x18, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x08, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1e, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x82, 0x01, 0x0a, 0x18, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x18, 0x0a, 0x07,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x53, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_collect_match_proto_rawDescOnce sync.Once
	file_collect_match_proto_rawDescData = file_collect_match_proto_rawDesc
)

func file_collect_match_proto_rawDescGZIP() []byte {
	file_collect_match_proto_rawDescOnce.Do(func() {
		file_collect_match_proto_rawDescData = protoimpl.X.CompressGZIP(file_collect_match_proto_rawDescData)
	})
	return file_collect_match_proto_rawDescData
}

var file_collect_match_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_collect_match_proto_goTypes = []interface{}{
	(*CollectMatchRequestParam)(nil), // 0: CollectMatchRequestParam
	(*PushMsgMatchRequestParam)(nil), // 1: PushMsgMatchRequestParam
}
var file_collect_match_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_collect_match_proto_init() }
func file_collect_match_proto_init() {
	if File_collect_match_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_collect_match_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectMatchRequestParam); i {
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
		file_collect_match_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsgMatchRequestParam); i {
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
			RawDescriptor: file_collect_match_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_collect_match_proto_goTypes,
		DependencyIndexes: file_collect_match_proto_depIdxs,
		MessageInfos:      file_collect_match_proto_msgTypes,
	}.Build()
	File_collect_match_proto = out.File
	file_collect_match_proto_rawDesc = nil
	file_collect_match_proto_goTypes = nil
	file_collect_match_proto_depIdxs = nil
}