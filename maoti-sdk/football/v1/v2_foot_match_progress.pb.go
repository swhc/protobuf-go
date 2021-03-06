// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: v2_foot_match_progress.proto

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

type V2FootMatchProgressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchId int64 `protobuf:"varint,1,opt,name=matchId,proto3" json:"matchId,omitempty"`
}

func (x *V2FootMatchProgressRequest) Reset() {
	*x = V2FootMatchProgressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_match_progress_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchProgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchProgressRequest) ProtoMessage() {}

func (x *V2FootMatchProgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_match_progress_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchProgressRequest.ProtoReflect.Descriptor instead.
func (*V2FootMatchProgressRequest) Descriptor() ([]byte, []int) {
	return file_v2_foot_match_progress_proto_rawDescGZIP(), []int{0}
}

func (x *V2FootMatchProgressRequest) GetMatchId() int64 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

type V2FootMatchProgressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HomeId          int64                      `protobuf:"varint,1,opt,name=homeId,proto3" json:"homeId,omitempty"`                   //??????ID
	AwayId          int64                      `protobuf:"varint,3,opt,name=awayId,proto3" json:"awayId,omitempty"`                   //??????ID
	Minute          int64                      `protobuf:"varint,5,opt,name=minute,proto3" json:"minute,omitempty"`                   //????????????????????????
	CurrMatchStatus int64                      `protobuf:"varint,6,opt,name=currMatchStatus,proto3" json:"currMatchStatus,omitempty"` //?????????????????????1.????????????2.????????????3.????????????4.????????????5.?????????6.??????????????????7.??????????????????8.?????????9.?????????10.?????????11.??????????????????12.??????????????????13.???????????????14.?????????????????????15.??????????????????16.?????????17.?????????18.?????????19.?????????20.?????????21.?????????22.??????
	Progress        []*V2FootMatchProgressInfo `protobuf:"bytes,7,rep,name=progress,proto3" json:"progress,omitempty"`                //????????????
}

func (x *V2FootMatchProgressResponse) Reset() {
	*x = V2FootMatchProgressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_match_progress_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchProgressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchProgressResponse) ProtoMessage() {}

func (x *V2FootMatchProgressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_match_progress_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchProgressResponse.ProtoReflect.Descriptor instead.
func (*V2FootMatchProgressResponse) Descriptor() ([]byte, []int) {
	return file_v2_foot_match_progress_proto_rawDescGZIP(), []int{1}
}

func (x *V2FootMatchProgressResponse) GetHomeId() int64 {
	if x != nil {
		return x.HomeId
	}
	return 0
}

func (x *V2FootMatchProgressResponse) GetAwayId() int64 {
	if x != nil {
		return x.AwayId
	}
	return 0
}

func (x *V2FootMatchProgressResponse) GetMinute() int64 {
	if x != nil {
		return x.Minute
	}
	return 0
}

func (x *V2FootMatchProgressResponse) GetCurrMatchStatus() int64 {
	if x != nil {
		return x.CurrMatchStatus
	}
	return 0
}

func (x *V2FootMatchProgressResponse) GetProgress() []*V2FootMatchProgressInfo {
	if x != nil {
		return x.Progress
	}
	return nil
}

type V2FootMatchProgressInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Minute       int64 `protobuf:"varint,1,opt,name=minute,proto3" json:"minute,omitempty"`             //???????????????????????????????????????
	ProgressType int64 `protobuf:"varint,2,opt,name=progressType,proto3" json:"progressType,omitempty"` //???????????????1.?????????2.?????????3.????????????
	TeamType     int64 `protobuf:"varint,3,opt,name=teamType,proto3" json:"teamType,omitempty"`         //????????????????????????1.?????????2.??????
	MatchStatus  int64 `protobuf:"varint,4,opt,name=matchStatus,proto3" json:"matchStatus,omitempty"`   //?????????????????????????????????0.????????????1.????????????2.????????????10.????????????21.??????????????????22.??????????????????25.?????????30.?????????31.?????????32.??????????????????33.??????????????????34.???????????????35.?????????????????????36.??????????????????40.?????????41.?????????42.?????????43.?????????44.?????????45.?????????100.??????
	CoordinateX  int64 `protobuf:"varint,5,opt,name=coordinateX,proto3" json:"coordinateX,omitempty"`   //??????X????????????????????????????????????????????????
}

func (x *V2FootMatchProgressInfo) Reset() {
	*x = V2FootMatchProgressInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_match_progress_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchProgressInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchProgressInfo) ProtoMessage() {}

func (x *V2FootMatchProgressInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_match_progress_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchProgressInfo.ProtoReflect.Descriptor instead.
func (*V2FootMatchProgressInfo) Descriptor() ([]byte, []int) {
	return file_v2_foot_match_progress_proto_rawDescGZIP(), []int{2}
}

func (x *V2FootMatchProgressInfo) GetMinute() int64 {
	if x != nil {
		return x.Minute
	}
	return 0
}

func (x *V2FootMatchProgressInfo) GetProgressType() int64 {
	if x != nil {
		return x.ProgressType
	}
	return 0
}

func (x *V2FootMatchProgressInfo) GetTeamType() int64 {
	if x != nil {
		return x.TeamType
	}
	return 0
}

func (x *V2FootMatchProgressInfo) GetMatchStatus() int64 {
	if x != nil {
		return x.MatchStatus
	}
	return 0
}

func (x *V2FootMatchProgressInfo) GetCoordinateX() int64 {
	if x != nil {
		return x.CoordinateX
	}
	return 0
}

var File_v2_foot_match_progress_proto protoreflect.FileDescriptor

var file_v2_foot_match_progress_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x76, 0x32, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36,
	0x0a, 0x1a, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x22, 0xc5, 0x01, 0x0a, 0x1b, 0x56, 0x32, 0x46, 0x6f, 0x6f,
	0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f, 0x6d, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x6f, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x77, 0x61, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x61, 0x77, 0x61, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x56, 0x32, 0x46,
	0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0xb5,
	0x01, 0x0a, 0x17, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x69,
	0x6e, 0x75, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x75,
	0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x58, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x58, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_foot_match_progress_proto_rawDescOnce sync.Once
	file_v2_foot_match_progress_proto_rawDescData = file_v2_foot_match_progress_proto_rawDesc
)

func file_v2_foot_match_progress_proto_rawDescGZIP() []byte {
	file_v2_foot_match_progress_proto_rawDescOnce.Do(func() {
		file_v2_foot_match_progress_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_foot_match_progress_proto_rawDescData)
	})
	return file_v2_foot_match_progress_proto_rawDescData
}

var file_v2_foot_match_progress_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v2_foot_match_progress_proto_goTypes = []interface{}{
	(*V2FootMatchProgressRequest)(nil),  // 0: V2FootMatchProgressRequest
	(*V2FootMatchProgressResponse)(nil), // 1: V2FootMatchProgressResponse
	(*V2FootMatchProgressInfo)(nil),     // 2: V2FootMatchProgressInfo
}
var file_v2_foot_match_progress_proto_depIdxs = []int32{
	2, // 0: V2FootMatchProgressResponse.progress:type_name -> V2FootMatchProgressInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v2_foot_match_progress_proto_init() }
func file_v2_foot_match_progress_proto_init() {
	if File_v2_foot_match_progress_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_foot_match_progress_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchProgressRequest); i {
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
		file_v2_foot_match_progress_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchProgressResponse); i {
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
		file_v2_foot_match_progress_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchProgressInfo); i {
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
			RawDescriptor: file_v2_foot_match_progress_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_foot_match_progress_proto_goTypes,
		DependencyIndexes: file_v2_foot_match_progress_proto_depIdxs,
		MessageInfos:      file_v2_foot_match_progress_proto_msgTypes,
	}.Build()
	File_v2_foot_match_progress_proto = out.File
	file_v2_foot_match_progress_proto_rawDesc = nil
	file_v2_foot_match_progress_proto_goTypes = nil
	file_v2_foot_match_progress_proto_depIdxs = nil
}
