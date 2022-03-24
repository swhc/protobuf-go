// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: v2_foot_player_collects.proto

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

type V2FootPlayerCollectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language  string  `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"` //请求语言
	PlayerIds []int64 `protobuf:"varint,2,rep,packed,name=PlayerIds,proto3" json:"PlayerIds,omitempty"`
}

func (x *V2FootPlayerCollectsRequest) Reset() {
	*x = V2FootPlayerCollectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_player_collects_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootPlayerCollectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootPlayerCollectsRequest) ProtoMessage() {}

func (x *V2FootPlayerCollectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_player_collects_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootPlayerCollectsRequest.ProtoReflect.Descriptor instead.
func (*V2FootPlayerCollectsRequest) Descriptor() ([]byte, []int) {
	return file_v2_foot_player_collects_proto_rawDescGZIP(), []int{0}
}

func (x *V2FootPlayerCollectsRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *V2FootPlayerCollectsRequest) GetPlayerIds() []int64 {
	if x != nil {
		return x.PlayerIds
	}
	return nil
}

type V2FootPlayerCollectsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*V2FootPlayerCollectsItem `protobuf:"bytes,1,rep,name=Players,proto3" json:"Players,omitempty"`
}

func (x *V2FootPlayerCollectsResp) Reset() {
	*x = V2FootPlayerCollectsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_player_collects_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootPlayerCollectsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootPlayerCollectsResp) ProtoMessage() {}

func (x *V2FootPlayerCollectsResp) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_player_collects_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootPlayerCollectsResp.ProtoReflect.Descriptor instead.
func (*V2FootPlayerCollectsResp) Descriptor() ([]byte, []int) {
	return file_v2_foot_player_collects_proto_rawDescGZIP(), []int{1}
}

func (x *V2FootPlayerCollectsResp) GetPlayers() []*V2FootPlayerCollectsItem {
	if x != nil {
		return x.Players
	}
	return nil
}

type V2FootPlayerCollectsItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SportType    int64  `protobuf:"varint,1,opt,name=SportType,proto3" json:"SportType,omitempty"`
	PersonName   string `protobuf:"bytes,2,opt,name=PersonName,proto3" json:"PersonName,omitempty"`
	PersonImages string `protobuf:"bytes,3,opt,name=PersonImages,proto3" json:"PersonImages,omitempty"`
	Matches      int64  `protobuf:"varint,4,opt,name=Matches,proto3" json:"Matches,omitempty"`        // 比赛次数
	Goals        int64  `protobuf:"varint,5,opt,name=Goals,proto3" json:"Goals,omitempty"`            // 进球/扑救
	Assist       int64  `protobuf:"varint,6,opt,name=Assist,proto3" json:"Assist,omitempty"`          // 助攻/失球
	PlayerScore  string `protobuf:"bytes,7,opt,name=PlayerScore,proto3" json:"PlayerScore,omitempty"` // 球员评分
	Position     int64  `protobuf:"varint,8,opt,name=Position,proto3" json:"Position,omitempty"`      // 位置(1:门将 2:球员)
}

func (x *V2FootPlayerCollectsItem) Reset() {
	*x = V2FootPlayerCollectsItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_player_collects_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootPlayerCollectsItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootPlayerCollectsItem) ProtoMessage() {}

func (x *V2FootPlayerCollectsItem) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_player_collects_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootPlayerCollectsItem.ProtoReflect.Descriptor instead.
func (*V2FootPlayerCollectsItem) Descriptor() ([]byte, []int) {
	return file_v2_foot_player_collects_proto_rawDescGZIP(), []int{2}
}

func (x *V2FootPlayerCollectsItem) GetSportType() int64 {
	if x != nil {
		return x.SportType
	}
	return 0
}

func (x *V2FootPlayerCollectsItem) GetPersonName() string {
	if x != nil {
		return x.PersonName
	}
	return ""
}

func (x *V2FootPlayerCollectsItem) GetPersonImages() string {
	if x != nil {
		return x.PersonImages
	}
	return ""
}

func (x *V2FootPlayerCollectsItem) GetMatches() int64 {
	if x != nil {
		return x.Matches
	}
	return 0
}

func (x *V2FootPlayerCollectsItem) GetGoals() int64 {
	if x != nil {
		return x.Goals
	}
	return 0
}

func (x *V2FootPlayerCollectsItem) GetAssist() int64 {
	if x != nil {
		return x.Assist
	}
	return 0
}

func (x *V2FootPlayerCollectsItem) GetPlayerScore() string {
	if x != nil {
		return x.PlayerScore
	}
	return ""
}

func (x *V2FootPlayerCollectsItem) GetPosition() int64 {
	if x != nil {
		return x.Position
	}
	return 0
}

var File_v2_foot_player_collects_proto protoreflect.FileDescriptor

var file_v2_foot_player_collects_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x76, 0x32, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x57, 0x0a, 0x1b, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x09, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x4f, 0x0a, 0x18, 0x56, 0x32, 0x46, 0x6f,
	0x6f, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x33, 0x0a, 0x07, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x73, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x07, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x82, 0x02, 0x0a, 0x18, 0x56, 0x32,
	0x46, 0x6f, 0x6f, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x53, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x73, 0x73, 0x69,
	0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x07,
	0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_foot_player_collects_proto_rawDescOnce sync.Once
	file_v2_foot_player_collects_proto_rawDescData = file_v2_foot_player_collects_proto_rawDesc
)

func file_v2_foot_player_collects_proto_rawDescGZIP() []byte {
	file_v2_foot_player_collects_proto_rawDescOnce.Do(func() {
		file_v2_foot_player_collects_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_foot_player_collects_proto_rawDescData)
	})
	return file_v2_foot_player_collects_proto_rawDescData
}

var file_v2_foot_player_collects_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v2_foot_player_collects_proto_goTypes = []interface{}{
	(*V2FootPlayerCollectsRequest)(nil), // 0: V2FootPlayerCollectsRequest
	(*V2FootPlayerCollectsResp)(nil),    // 1: V2FootPlayerCollectsResp
	(*V2FootPlayerCollectsItem)(nil),    // 2: V2FootPlayerCollectsItem
}
var file_v2_foot_player_collects_proto_depIdxs = []int32{
	2, // 0: V2FootPlayerCollectsResp.Players:type_name -> V2FootPlayerCollectsItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v2_foot_player_collects_proto_init() }
func file_v2_foot_player_collects_proto_init() {
	if File_v2_foot_player_collects_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_foot_player_collects_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootPlayerCollectsRequest); i {
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
		file_v2_foot_player_collects_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootPlayerCollectsResp); i {
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
		file_v2_foot_player_collects_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootPlayerCollectsItem); i {
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
			RawDescriptor: file_v2_foot_player_collects_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_foot_player_collects_proto_goTypes,
		DependencyIndexes: file_v2_foot_player_collects_proto_depIdxs,
		MessageInfos:      file_v2_foot_player_collects_proto_msgTypes,
	}.Build()
	File_v2_foot_player_collects_proto = out.File
	file_v2_foot_player_collects_proto_rawDesc = nil
	file_v2_foot_player_collects_proto_goTypes = nil
	file_v2_foot_player_collects_proto_depIdxs = nil
}
