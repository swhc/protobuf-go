// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: foot_league_stand.proto

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

//资料库 - 排名（积分榜）
type FootLeagueStandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`  //请求语言  1:zh  2:en
	SportId  string `protobuf:"bytes,1,opt,name=sportId,proto3" json:"sportId,omitempty"`    //1足球 2篮球
	SeasonId int64  `protobuf:"varint,2,opt,name=SeasonId,proto3" json:"SeasonId,omitempty"` //年份对应id
}

func (x *FootLeagueStandRequest) Reset() {
	*x = FootLeagueStandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_league_stand_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootLeagueStandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootLeagueStandRequest) ProtoMessage() {}

func (x *FootLeagueStandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foot_league_stand_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootLeagueStandRequest.ProtoReflect.Descriptor instead.
func (*FootLeagueStandRequest) Descriptor() ([]byte, []int) {
	return file_foot_league_stand_proto_rawDescGZIP(), []int{0}
}

func (x *FootLeagueStandRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *FootLeagueStandRequest) GetSportId() string {
	if x != nil {
		return x.SportId
	}
	return ""
}

func (x *FootLeagueStandRequest) GetSeasonId() int64 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

type FootLeagueStandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*FootGroupCompetition `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *FootLeagueStandResponse) Reset() {
	*x = FootLeagueStandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_league_stand_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootLeagueStandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootLeagueStandResponse) ProtoMessage() {}

func (x *FootLeagueStandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_foot_league_stand_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootLeagueStandResponse.ProtoReflect.Descriptor instead.
func (*FootLeagueStandResponse) Descriptor() ([]byte, []int) {
	return file_foot_league_stand_proto_rawDescGZIP(), []int{1}
}

func (x *FootLeagueStandResponse) GetList() []*FootGroupCompetition {
	if x != nil {
		return x.List
	}
	return nil
}

type FootGroupCompetition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoundName string                 `protobuf:"bytes,1,opt,name=roundName,proto3" json:"roundName,omitempty"` //分组名称
	List      []*FootLeagueStandList `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *FootGroupCompetition) Reset() {
	*x = FootGroupCompetition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_league_stand_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootGroupCompetition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootGroupCompetition) ProtoMessage() {}

func (x *FootGroupCompetition) ProtoReflect() protoreflect.Message {
	mi := &file_foot_league_stand_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootGroupCompetition.ProtoReflect.Descriptor instead.
func (*FootGroupCompetition) Descriptor() ([]byte, []int) {
	return file_foot_league_stand_proto_rawDescGZIP(), []int{2}
}

func (x *FootGroupCompetition) GetRoundName() string {
	if x != nil {
		return x.RoundName
	}
	return ""
}

func (x *FootGroupCompetition) GetList() []*FootLeagueStandList {
	if x != nil {
		return x.List
	}
	return nil
}

type FootLeagueStandList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`                   //id
	TeamName    string `protobuf:"bytes,2,opt,name=TeamName,proto3" json:"TeamName,omitempty"`        //球队名称
	TeamImage   string `protobuf:"bytes,3,opt,name=TeamImage,proto3" json:"TeamImage,omitempty"`      //球队图片
	Gamesplayed int64  `protobuf:"varint,4,opt,name=Gamesplayed,proto3" json:"Gamesplayed,omitempty"` //比赛次数
	Win         int64  `protobuf:"varint,5,opt,name=Win,proto3" json:"Win,omitempty"`                 //胜
	Losses      int64  `protobuf:"varint,7,opt,name=Losses,proto3" json:"Losses,omitempty"`           //负
	Rank        int64  `protobuf:"varint,10,opt,name=Rank,proto3" json:"Rank,omitempty"`              //排名
	WinPctTotal string `protobuf:"bytes,6,opt,name=winPctTotal,proto3" json:"winPctTotal,omitempty"`  //胜率 40
	GameBehind  string `protobuf:"bytes,9,opt,name=gameBehind,proto3" json:"gameBehind,omitempty"`    //胜差 43
	Streak      int64  `protobuf:"varint,8,opt,name=streak,proto3" json:"streak,omitempty"`           //连续战绩，+连胜，-连败 37
}

func (x *FootLeagueStandList) Reset() {
	*x = FootLeagueStandList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_league_stand_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootLeagueStandList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootLeagueStandList) ProtoMessage() {}

func (x *FootLeagueStandList) ProtoReflect() protoreflect.Message {
	mi := &file_foot_league_stand_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootLeagueStandList.ProtoReflect.Descriptor instead.
func (*FootLeagueStandList) Descriptor() ([]byte, []int) {
	return file_foot_league_stand_proto_rawDescGZIP(), []int{3}
}

func (x *FootLeagueStandList) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FootLeagueStandList) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

func (x *FootLeagueStandList) GetTeamImage() string {
	if x != nil {
		return x.TeamImage
	}
	return ""
}

func (x *FootLeagueStandList) GetGamesplayed() int64 {
	if x != nil {
		return x.Gamesplayed
	}
	return 0
}

func (x *FootLeagueStandList) GetWin() int64 {
	if x != nil {
		return x.Win
	}
	return 0
}

func (x *FootLeagueStandList) GetLosses() int64 {
	if x != nil {
		return x.Losses
	}
	return 0
}

func (x *FootLeagueStandList) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *FootLeagueStandList) GetWinPctTotal() string {
	if x != nil {
		return x.WinPctTotal
	}
	return ""
}

func (x *FootLeagueStandList) GetGameBehind() string {
	if x != nil {
		return x.GameBehind
	}
	return ""
}

func (x *FootLeagueStandList) GetStreak() int64 {
	if x != nil {
		return x.Streak
	}
	return 0
}

var File_foot_league_stand_proto protoreflect.FileDescriptor

var file_foot_league_stand_proto_rawDesc = []byte{
	0x0a, 0x17, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x16, 0x46, 0x6f, 0x6f,
	0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x17, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x65, 0x61,
	0x67, 0x75, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x5e, 0x0a, 0x14, 0x46,
	0x6f, 0x6f, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x28, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x74, 0x61, 0x6e,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x99, 0x02, 0x0a, 0x13,
	0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x57, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x57, 0x69,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x6f, 0x73, 0x73, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x4c, 0x6f, 0x73, 0x73, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x61, 0x6e,
	0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x20, 0x0a,
	0x0b, 0x77, 0x69, 0x6e, 0x50, 0x63, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x77, 0x69, 0x6e, 0x50, 0x63, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x1e, 0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x42, 0x65, 0x68, 0x69, 0x6e, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x42, 0x65, 0x68, 0x69, 0x6e, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6b, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foot_league_stand_proto_rawDescOnce sync.Once
	file_foot_league_stand_proto_rawDescData = file_foot_league_stand_proto_rawDesc
)

func file_foot_league_stand_proto_rawDescGZIP() []byte {
	file_foot_league_stand_proto_rawDescOnce.Do(func() {
		file_foot_league_stand_proto_rawDescData = protoimpl.X.CompressGZIP(file_foot_league_stand_proto_rawDescData)
	})
	return file_foot_league_stand_proto_rawDescData
}

var file_foot_league_stand_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_foot_league_stand_proto_goTypes = []interface{}{
	(*FootLeagueStandRequest)(nil),  // 0: FootLeagueStandRequest
	(*FootLeagueStandResponse)(nil), // 1: FootLeagueStandResponse
	(*FootGroupCompetition)(nil),    // 2: FootGroupCompetition
	(*FootLeagueStandList)(nil),     // 3: FootLeagueStandList
}
var file_foot_league_stand_proto_depIdxs = []int32{
	2, // 0: FootLeagueStandResponse.list:type_name -> FootGroupCompetition
	3, // 1: FootGroupCompetition.list:type_name -> FootLeagueStandList
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_foot_league_stand_proto_init() }
func file_foot_league_stand_proto_init() {
	if File_foot_league_stand_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foot_league_stand_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootLeagueStandRequest); i {
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
		file_foot_league_stand_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootLeagueStandResponse); i {
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
		file_foot_league_stand_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootGroupCompetition); i {
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
		file_foot_league_stand_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootLeagueStandList); i {
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
			RawDescriptor: file_foot_league_stand_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foot_league_stand_proto_goTypes,
		DependencyIndexes: file_foot_league_stand_proto_depIdxs,
		MessageInfos:      file_foot_league_stand_proto_msgTypes,
	}.Build()
	File_foot_league_stand_proto = out.File
	file_foot_league_stand_proto_rawDesc = nil
	file_foot_league_stand_proto_goTypes = nil
	file_foot_league_stand_proto_depIdxs = nil
}