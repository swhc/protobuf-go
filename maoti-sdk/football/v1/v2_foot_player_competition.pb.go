// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: v2_foot_player_competition.proto

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

type V2PlayerInfoCompetitionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SportId  string `protobuf:"bytes,1,opt,name=sportId,proto3" json:"sportId,omitempty"`    // 1足球 2篮球
	PersonId int64  `protobuf:"varint,2,opt,name=personId,proto3" json:"personId,omitempty"` // 球员id
	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`  // 请求语言
	Year     string `protobuf:"bytes,4,opt,name=Year,proto3" json:"Year,omitempty"`          // 年份
}

func (x *V2PlayerInfoCompetitionRequest) Reset() {
	*x = V2PlayerInfoCompetitionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_player_competition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2PlayerInfoCompetitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2PlayerInfoCompetitionRequest) ProtoMessage() {}

func (x *V2PlayerInfoCompetitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_player_competition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2PlayerInfoCompetitionRequest.ProtoReflect.Descriptor instead.
func (*V2PlayerInfoCompetitionRequest) Descriptor() ([]byte, []int) {
	return file_v2_foot_player_competition_proto_rawDescGZIP(), []int{0}
}

func (x *V2PlayerInfoCompetitionRequest) GetSportId() string {
	if x != nil {
		return x.SportId
	}
	return ""
}

func (x *V2PlayerInfoCompetitionRequest) GetPersonId() int64 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

func (x *V2PlayerInfoCompetitionRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *V2PlayerInfoCompetitionRequest) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

type V2FootPlayerInfoCompetitionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year        string                             `protobuf:"bytes,1,opt,name=Year,proto3" json:"Year,omitempty"`               // 年份
	Competition []*V2FootPlayerInfoCompetitionItem `protobuf:"bytes,2,rep,name=Competition,proto3" json:"Competition,omitempty"` // 比赛数据
}

func (x *V2FootPlayerInfoCompetitionResp) Reset() {
	*x = V2FootPlayerInfoCompetitionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_player_competition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootPlayerInfoCompetitionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootPlayerInfoCompetitionResp) ProtoMessage() {}

func (x *V2FootPlayerInfoCompetitionResp) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_player_competition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootPlayerInfoCompetitionResp.ProtoReflect.Descriptor instead.
func (*V2FootPlayerInfoCompetitionResp) Descriptor() ([]byte, []int) {
	return file_v2_foot_player_competition_proto_rawDescGZIP(), []int{1}
}

func (x *V2FootPlayerInfoCompetitionResp) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *V2FootPlayerInfoCompetitionResp) GetCompetition() []*V2FootPlayerInfoCompetitionItem {
	if x != nil {
		return x.Competition
	}
	return nil
}

type V2FootPlayerInfoCompetitionItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchTime         string `protobuf:"bytes,1,opt,name=MatchTime,proto3" json:"MatchTime,omitempty"`                  //比赛时间
	TournamentId      int64  `protobuf:"varint,2,opt,name=TournamentId,proto3" json:"TournamentId,omitempty"`           // 联赛id
	HomeTeamId        int64  `protobuf:"varint,3,opt,name=HomeTeamId,proto3" json:"HomeTeamId,omitempty"`               // 主队id
	AwayTeamId        int64  `protobuf:"varint,4,opt,name=AwayTeamId,proto3" json:"AwayTeamId,omitempty"`               // 客队id
	HomeTeamScore     int64  `protobuf:"varint,5,opt,name=HomeTeamScore,proto3" json:"HomeTeamScore,omitempty"`         // 主队比分
	AwayTeamScore     int64  `protobuf:"varint,6,opt,name=AwayTeamScore,proto3" json:"AwayTeamScore,omitempty"`         // 客队比分
	HomeTeamHalfScore int64  `protobuf:"varint,7,opt,name=HomeTeamHalfScore,proto3" json:"HomeTeamHalfScore,omitempty"` // 主队半场比分
	AwayTeamHalfScore int64  `protobuf:"varint,8,opt,name=AwayTeamHalfScore,proto3" json:"AwayTeamHalfScore,omitempty"` // 客队半场比分
	PlayerScore       string `protobuf:"bytes,9,opt,name=PlayerScore,proto3" json:"PlayerScore,omitempty"`              // 球员评分
	MatchId           int64  `protobuf:"varint,10,opt,name=MatchId,proto3" json:"MatchId,omitempty"`                    // 比赛自增id
}

func (x *V2FootPlayerInfoCompetitionItem) Reset() {
	*x = V2FootPlayerInfoCompetitionItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_player_competition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootPlayerInfoCompetitionItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootPlayerInfoCompetitionItem) ProtoMessage() {}

func (x *V2FootPlayerInfoCompetitionItem) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_player_competition_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootPlayerInfoCompetitionItem.ProtoReflect.Descriptor instead.
func (*V2FootPlayerInfoCompetitionItem) Descriptor() ([]byte, []int) {
	return file_v2_foot_player_competition_proto_rawDescGZIP(), []int{2}
}

func (x *V2FootPlayerInfoCompetitionItem) GetMatchTime() string {
	if x != nil {
		return x.MatchTime
	}
	return ""
}

func (x *V2FootPlayerInfoCompetitionItem) GetTournamentId() int64 {
	if x != nil {
		return x.TournamentId
	}
	return 0
}

func (x *V2FootPlayerInfoCompetitionItem) GetHomeTeamId() int64 {
	if x != nil {
		return x.HomeTeamId
	}
	return 0
}

func (x *V2FootPlayerInfoCompetitionItem) GetAwayTeamId() int64 {
	if x != nil {
		return x.AwayTeamId
	}
	return 0
}

func (x *V2FootPlayerInfoCompetitionItem) GetHomeTeamScore() int64 {
	if x != nil {
		return x.HomeTeamScore
	}
	return 0
}

func (x *V2FootPlayerInfoCompetitionItem) GetAwayTeamScore() int64 {
	if x != nil {
		return x.AwayTeamScore
	}
	return 0
}

func (x *V2FootPlayerInfoCompetitionItem) GetHomeTeamHalfScore() int64 {
	if x != nil {
		return x.HomeTeamHalfScore
	}
	return 0
}

func (x *V2FootPlayerInfoCompetitionItem) GetAwayTeamHalfScore() int64 {
	if x != nil {
		return x.AwayTeamHalfScore
	}
	return 0
}

func (x *V2FootPlayerInfoCompetitionItem) GetPlayerScore() string {
	if x != nil {
		return x.PlayerScore
	}
	return ""
}

func (x *V2FootPlayerInfoCompetitionItem) GetMatchId() int64 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

var File_v2_foot_player_competition_proto protoreflect.FileDescriptor

var file_v2_foot_player_competition_proto_rawDesc = []byte{
	0x0a, 0x20, 0x76, 0x32, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x1e, 0x56, 0x32, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x22, 0x79, 0x0a, 0x1f, 0x56,
	0x32, 0x46, 0x6f, 0x6f, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x43,
	0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x59, 0x65,
	0x61, 0x72, 0x12, 0x42, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x65,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x87, 0x03, 0x0a, 0x1f, 0x56, 0x32, 0x46, 0x6f, 0x6f,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x65,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x48, 0x6f, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x48, 0x6f, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x41, 0x77, 0x61, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x41, 0x77, 0x61, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d,
	0x48, 0x6f, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x48, 0x6f, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x41, 0x77, 0x61, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x41, 0x77, 0x61, 0x79, 0x54,
	0x65, 0x61, 0x6d, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x48, 0x6f, 0x6d, 0x65,
	0x54, 0x65, 0x61, 0x6d, 0x48, 0x61, 0x6c, 0x66, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x11, 0x48, 0x6f, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x48, 0x61, 0x6c,
	0x66, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x41, 0x77, 0x61, 0x79, 0x54, 0x65,
	0x61, 0x6d, 0x48, 0x61, 0x6c, 0x66, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x11, 0x41, 0x77, 0x61, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x48, 0x61, 0x6c, 0x66, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64,
	0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_v2_foot_player_competition_proto_rawDescOnce sync.Once
	file_v2_foot_player_competition_proto_rawDescData = file_v2_foot_player_competition_proto_rawDesc
)

func file_v2_foot_player_competition_proto_rawDescGZIP() []byte {
	file_v2_foot_player_competition_proto_rawDescOnce.Do(func() {
		file_v2_foot_player_competition_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_foot_player_competition_proto_rawDescData)
	})
	return file_v2_foot_player_competition_proto_rawDescData
}

var file_v2_foot_player_competition_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v2_foot_player_competition_proto_goTypes = []interface{}{
	(*V2PlayerInfoCompetitionRequest)(nil),  // 0: V2PlayerInfoCompetitionRequest
	(*V2FootPlayerInfoCompetitionResp)(nil), // 1: V2FootPlayerInfoCompetitionResp
	(*V2FootPlayerInfoCompetitionItem)(nil), // 2: V2FootPlayerInfoCompetitionItem
}
var file_v2_foot_player_competition_proto_depIdxs = []int32{
	2, // 0: V2FootPlayerInfoCompetitionResp.Competition:type_name -> V2FootPlayerInfoCompetitionItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v2_foot_player_competition_proto_init() }
func file_v2_foot_player_competition_proto_init() {
	if File_v2_foot_player_competition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_foot_player_competition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2PlayerInfoCompetitionRequest); i {
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
		file_v2_foot_player_competition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootPlayerInfoCompetitionResp); i {
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
		file_v2_foot_player_competition_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootPlayerInfoCompetitionItem); i {
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
			RawDescriptor: file_v2_foot_player_competition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_foot_player_competition_proto_goTypes,
		DependencyIndexes: file_v2_foot_player_competition_proto_depIdxs,
		MessageInfos:      file_v2_foot_player_competition_proto_msgTypes,
	}.Build()
	File_v2_foot_player_competition_proto = out.File
	file_v2_foot_player_competition_proto_rawDesc = nil
	file_v2_foot_player_competition_proto_goTypes = nil
	file_v2_foot_player_competition_proto_depIdxs = nil
}
