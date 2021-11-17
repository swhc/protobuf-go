// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: league_stand.proto

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
type LeagueStandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeasonId int64 `protobuf:"varint,3,opt,name=SeasonId,proto3" json:"SeasonId,omitempty"` //年份对应id
}

func (x *LeagueStandRequest) Reset() {
	*x = LeagueStandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_league_stand_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeagueStandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeagueStandRequest) ProtoMessage() {}

func (x *LeagueStandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_league_stand_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeagueStandRequest.ProtoReflect.Descriptor instead.
func (*LeagueStandRequest) Descriptor() ([]byte, []int) {
	return file_league_stand_proto_rawDescGZIP(), []int{0}
}

func (x *LeagueStandRequest) GetSeasonId() int64 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

type LeagueStandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*GroupCompetition `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *LeagueStandResponse) Reset() {
	*x = LeagueStandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_league_stand_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeagueStandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeagueStandResponse) ProtoMessage() {}

func (x *LeagueStandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_league_stand_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeagueStandResponse.ProtoReflect.Descriptor instead.
func (*LeagueStandResponse) Descriptor() ([]byte, []int) {
	return file_league_stand_proto_rawDescGZIP(), []int{1}
}

func (x *LeagueStandResponse) GetList() []*GroupCompetition {
	if x != nil {
		return x.List
	}
	return nil
}

type GroupCompetition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoundName string             `protobuf:"bytes,1,opt,name=roundName,proto3" json:"roundName,omitempty"` //分组名称
	List      []*LeagueStandList `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *GroupCompetition) Reset() {
	*x = GroupCompetition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_league_stand_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupCompetition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupCompetition) ProtoMessage() {}

func (x *GroupCompetition) ProtoReflect() protoreflect.Message {
	mi := &file_league_stand_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupCompetition.ProtoReflect.Descriptor instead.
func (*GroupCompetition) Descriptor() ([]byte, []int) {
	return file_league_stand_proto_rawDescGZIP(), []int{2}
}

func (x *GroupCompetition) GetRoundName() string {
	if x != nil {
		return x.RoundName
	}
	return ""
}

func (x *GroupCompetition) GetList() []*LeagueStandList {
	if x != nil {
		return x.List
	}
	return nil
}

type LeagueStandList struct {
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

func (x *LeagueStandList) Reset() {
	*x = LeagueStandList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_league_stand_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeagueStandList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeagueStandList) ProtoMessage() {}

func (x *LeagueStandList) ProtoReflect() protoreflect.Message {
	mi := &file_league_stand_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeagueStandList.ProtoReflect.Descriptor instead.
func (*LeagueStandList) Descriptor() ([]byte, []int) {
	return file_league_stand_proto_rawDescGZIP(), []int{3}
}

func (x *LeagueStandList) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LeagueStandList) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

func (x *LeagueStandList) GetTeamImage() string {
	if x != nil {
		return x.TeamImage
	}
	return ""
}

func (x *LeagueStandList) GetGamesplayed() int64 {
	if x != nil {
		return x.Gamesplayed
	}
	return 0
}

func (x *LeagueStandList) GetWin() int64 {
	if x != nil {
		return x.Win
	}
	return 0
}

func (x *LeagueStandList) GetLosses() int64 {
	if x != nil {
		return x.Losses
	}
	return 0
}

func (x *LeagueStandList) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *LeagueStandList) GetWinPctTotal() string {
	if x != nil {
		return x.WinPctTotal
	}
	return ""
}

func (x *LeagueStandList) GetGameBehind() string {
	if x != nil {
		return x.GameBehind
	}
	return ""
}

func (x *LeagueStandList) GetStreak() int64 {
	if x != nil {
		return x.Streak
	}
	return 0
}

var File_league_stand_proto protoreflect.FileDescriptor

var file_league_stand_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x12, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x74,
	0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x13, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65,
	0x53, 0x74, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0x56, 0x0a, 0x10, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6d,
	0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x74, 0x61,
	0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x95, 0x02, 0x0a,
	0x0f, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x54, 0x65, 0x61, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x57, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x57, 0x69, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x4c, 0x6f, 0x73, 0x73, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x4c, 0x6f, 0x73, 0x73, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x61, 0x6e, 0x6b, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x69,
	0x6e, 0x50, 0x63, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x77, 0x69, 0x6e, 0x50, 0x63, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x0a,
	0x67, 0x61, 0x6d, 0x65, 0x42, 0x65, 0x68, 0x69, 0x6e, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x42, 0x65, 0x68, 0x69, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6b, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_league_stand_proto_rawDescOnce sync.Once
	file_league_stand_proto_rawDescData = file_league_stand_proto_rawDesc
)

func file_league_stand_proto_rawDescGZIP() []byte {
	file_league_stand_proto_rawDescOnce.Do(func() {
		file_league_stand_proto_rawDescData = protoimpl.X.CompressGZIP(file_league_stand_proto_rawDescData)
	})
	return file_league_stand_proto_rawDescData
}

var file_league_stand_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_league_stand_proto_goTypes = []interface{}{
	(*LeagueStandRequest)(nil),  // 0: LeagueStandRequest
	(*LeagueStandResponse)(nil), // 1: LeagueStandResponse
	(*GroupCompetition)(nil),    // 2: GroupCompetition
	(*LeagueStandList)(nil),     // 3: LeagueStandList
}
var file_league_stand_proto_depIdxs = []int32{
	2, // 0: LeagueStandResponse.list:type_name -> GroupCompetition
	3, // 1: GroupCompetition.list:type_name -> LeagueStandList
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_league_stand_proto_init() }
func file_league_stand_proto_init() {
	if File_league_stand_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_league_stand_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeagueStandRequest); i {
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
		file_league_stand_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeagueStandResponse); i {
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
		file_league_stand_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupCompetition); i {
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
		file_league_stand_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeagueStandList); i {
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
			RawDescriptor: file_league_stand_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_league_stand_proto_goTypes,
		DependencyIndexes: file_league_stand_proto_depIdxs,
		MessageInfos:      file_league_stand_proto_msgTypes,
	}.Build()
	File_league_stand_proto = out.File
	file_league_stand_proto_rawDesc = nil
	file_league_stand_proto_goTypes = nil
	file_league_stand_proto_depIdxs = nil
}