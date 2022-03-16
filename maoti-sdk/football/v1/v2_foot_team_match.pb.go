// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: v2_foot_team_match.proto

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

//球队赛程
type V2FootTeamMatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId     int64  `protobuf:"varint,1,opt,name=teamId,proto3" json:"teamId,omitempty"`        //球队ID
	MatchMonth string `protobuf:"bytes,2,opt,name=matchMonth,proto3" json:"matchMonth,omitempty"` //比赛月份，比如：2022-03
	Language   string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`     //请求语言
}

func (x *V2FootTeamMatchRequest) Reset() {
	*x = V2FootTeamMatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_team_match_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootTeamMatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootTeamMatchRequest) ProtoMessage() {}

func (x *V2FootTeamMatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_team_match_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootTeamMatchRequest.ProtoReflect.Descriptor instead.
func (*V2FootTeamMatchRequest) Descriptor() ([]byte, []int) {
	return file_v2_foot_team_match_proto_rawDescGZIP(), []int{0}
}

func (x *V2FootTeamMatchRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *V2FootTeamMatchRequest) GetMatchMonth() string {
	if x != nil {
		return x.MatchMonth
	}
	return ""
}

func (x *V2FootTeamMatchRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type V2FootTeamMatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info []*V2FootTeamMatchInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
}

func (x *V2FootTeamMatchResponse) Reset() {
	*x = V2FootTeamMatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_team_match_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootTeamMatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootTeamMatchResponse) ProtoMessage() {}

func (x *V2FootTeamMatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_team_match_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootTeamMatchResponse.ProtoReflect.Descriptor instead.
func (*V2FootTeamMatchResponse) Descriptor() ([]byte, []int) {
	return file_v2_foot_team_match_proto_rawDescGZIP(), []int{1}
}

func (x *V2FootTeamMatchResponse) GetInfo() []*V2FootTeamMatchInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type V2FootTeamMatchInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchId       int64  `protobuf:"varint,1,opt,name=matchId,proto3" json:"matchId,omitempty"`              //比赛ID
	MatchTime     int64  `protobuf:"varint,2,opt,name=matchTime,proto3" json:"matchTime,omitempty"`          //比赛时间
	TourId        int64  `protobuf:"varint,3,opt,name=tourId,proto3" json:"tourId,omitempty"`                //联赛ID
	TourName      string `protobuf:"bytes,4,opt,name=tourName,proto3" json:"tourName,omitempty"`             //联赛名称
	HomeId        int64  `protobuf:"varint,5,opt,name=homeId,proto3" json:"homeId,omitempty"`                //主队ID
	HomeName      string `protobuf:"bytes,6,opt,name=homeName,proto3" json:"homeName,omitempty"`             //主队名称
	HomeLogo      string `protobuf:"bytes,7,opt,name=homeLogo,proto3" json:"homeLogo,omitempty"`             //主队图片
	HomeScore     int64  `protobuf:"varint,8,opt,name=homeScore,proto3" json:"homeScore,omitempty"`          //主队得分
	HomeHalfScore int64  `protobuf:"varint,9,opt,name=homeHalfScore,proto3" json:"homeHalfScore,omitempty"`  //主队半场得分
	AwayId        int64  `protobuf:"varint,10,opt,name=awayId,proto3" json:"awayId,omitempty"`               //客队ID
	AwayName      string `protobuf:"bytes,11,opt,name=awayName,proto3" json:"awayName,omitempty"`            //客队名称
	AwayLogo      string `protobuf:"bytes,12,opt,name=awayLogo,proto3" json:"awayLogo,omitempty"`            //客队图片
	AwayScore     int64  `protobuf:"varint,13,opt,name=awayScore,proto3" json:"awayScore,omitempty"`         //客队得分
	AwayHalfScore int64  `protobuf:"varint,14,opt,name=awayHalfScore,proto3" json:"awayHalfScore,omitempty"` //客队半场得分
	Winner        string `protobuf:"bytes,15,opt,name=winner,proto3" json:"winner,omitempty"`                //比赛结果
	RoundId       int64  `protobuf:"varint,16,opt,name=roundId,proto3" json:"roundId,omitempty"`             //比赛轮次
	RoundName     string `protobuf:"bytes,17,opt,name=roundName,proto3" json:"roundName,omitempty"`          //比赛轮次名次
}

func (x *V2FootTeamMatchInfo) Reset() {
	*x = V2FootTeamMatchInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_team_match_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootTeamMatchInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootTeamMatchInfo) ProtoMessage() {}

func (x *V2FootTeamMatchInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_team_match_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootTeamMatchInfo.ProtoReflect.Descriptor instead.
func (*V2FootTeamMatchInfo) Descriptor() ([]byte, []int) {
	return file_v2_foot_team_match_proto_rawDescGZIP(), []int{2}
}

func (x *V2FootTeamMatchInfo) GetMatchId() int64 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

func (x *V2FootTeamMatchInfo) GetMatchTime() int64 {
	if x != nil {
		return x.MatchTime
	}
	return 0
}

func (x *V2FootTeamMatchInfo) GetTourId() int64 {
	if x != nil {
		return x.TourId
	}
	return 0
}

func (x *V2FootTeamMatchInfo) GetTourName() string {
	if x != nil {
		return x.TourName
	}
	return ""
}

func (x *V2FootTeamMatchInfo) GetHomeId() int64 {
	if x != nil {
		return x.HomeId
	}
	return 0
}

func (x *V2FootTeamMatchInfo) GetHomeName() string {
	if x != nil {
		return x.HomeName
	}
	return ""
}

func (x *V2FootTeamMatchInfo) GetHomeLogo() string {
	if x != nil {
		return x.HomeLogo
	}
	return ""
}

func (x *V2FootTeamMatchInfo) GetHomeScore() int64 {
	if x != nil {
		return x.HomeScore
	}
	return 0
}

func (x *V2FootTeamMatchInfo) GetHomeHalfScore() int64 {
	if x != nil {
		return x.HomeHalfScore
	}
	return 0
}

func (x *V2FootTeamMatchInfo) GetAwayId() int64 {
	if x != nil {
		return x.AwayId
	}
	return 0
}

func (x *V2FootTeamMatchInfo) GetAwayName() string {
	if x != nil {
		return x.AwayName
	}
	return ""
}

func (x *V2FootTeamMatchInfo) GetAwayLogo() string {
	if x != nil {
		return x.AwayLogo
	}
	return ""
}

func (x *V2FootTeamMatchInfo) GetAwayScore() int64 {
	if x != nil {
		return x.AwayScore
	}
	return 0
}

func (x *V2FootTeamMatchInfo) GetAwayHalfScore() int64 {
	if x != nil {
		return x.AwayHalfScore
	}
	return 0
}

func (x *V2FootTeamMatchInfo) GetWinner() string {
	if x != nil {
		return x.Winner
	}
	return ""
}

func (x *V2FootTeamMatchInfo) GetRoundId() int64 {
	if x != nil {
		return x.RoundId
	}
	return 0
}

func (x *V2FootTeamMatchInfo) GetRoundName() string {
	if x != nil {
		return x.RoundName
	}
	return ""
}

var File_v2_foot_team_match_proto protoreflect.FileDescriptor

var file_v2_foot_team_match_proto_rawDesc = []byte{
	0x0a, 0x18, 0x76, 0x32, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x16, 0x56, 0x32,
	0x46, 0x6f, 0x6f, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x43, 0x0a, 0x17, 0x56, 0x32, 0x46, 0x6f,
	0x6f, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xf9, 0x03,
	0x0a, 0x13, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x6f, 0x75, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74,
	0x6f, 0x75, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x6f, 0x75, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x75, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x68, 0x6f, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x6d,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x6d,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x4c, 0x6f, 0x67,
	0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x4c, 0x6f, 0x67,
	0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x68, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x68, 0x6f, 0x6d, 0x65, 0x48, 0x61, 0x6c, 0x66, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x68, 0x6f, 0x6d, 0x65, 0x48, 0x61, 0x6c, 0x66,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x77, 0x61, 0x79, 0x49, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x77, 0x61, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x77, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x77, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x77, 0x61,
	0x79, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x77, 0x61,
	0x79, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x77, 0x61, 0x79, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x77, 0x61, 0x79, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x77, 0x61, 0x79, 0x48, 0x61, 0x6c, 0x66, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x77, 0x61, 0x79,
	0x48, 0x61, 0x6c, 0x66, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_foot_team_match_proto_rawDescOnce sync.Once
	file_v2_foot_team_match_proto_rawDescData = file_v2_foot_team_match_proto_rawDesc
)

func file_v2_foot_team_match_proto_rawDescGZIP() []byte {
	file_v2_foot_team_match_proto_rawDescOnce.Do(func() {
		file_v2_foot_team_match_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_foot_team_match_proto_rawDescData)
	})
	return file_v2_foot_team_match_proto_rawDescData
}

var file_v2_foot_team_match_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v2_foot_team_match_proto_goTypes = []interface{}{
	(*V2FootTeamMatchRequest)(nil),  // 0: V2FootTeamMatchRequest
	(*V2FootTeamMatchResponse)(nil), // 1: V2FootTeamMatchResponse
	(*V2FootTeamMatchInfo)(nil),     // 2: V2FootTeamMatchInfo
}
var file_v2_foot_team_match_proto_depIdxs = []int32{
	2, // 0: V2FootTeamMatchResponse.info:type_name -> V2FootTeamMatchInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v2_foot_team_match_proto_init() }
func file_v2_foot_team_match_proto_init() {
	if File_v2_foot_team_match_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_foot_team_match_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootTeamMatchRequest); i {
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
		file_v2_foot_team_match_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootTeamMatchResponse); i {
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
		file_v2_foot_team_match_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootTeamMatchInfo); i {
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
			RawDescriptor: file_v2_foot_team_match_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_foot_team_match_proto_goTypes,
		DependencyIndexes: file_v2_foot_team_match_proto_depIdxs,
		MessageInfos:      file_v2_foot_team_match_proto_msgTypes,
	}.Build()
	File_v2_foot_team_match_proto = out.File
	file_v2_foot_team_match_proto_rawDesc = nil
	file_v2_foot_team_match_proto_goTypes = nil
	file_v2_foot_team_match_proto_depIdxs = nil
}
