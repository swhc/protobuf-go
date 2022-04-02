// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: v2_foot_team_historical_match.proto

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

type V2FootTeamHistoricalMatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language string `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`  //请求语言
	MatchId  int64  `protobuf:"varint,2,opt,name=matchId,proto3" json:"matchId,omitempty"`   //比赛ID
	TimeZone int64  `protobuf:"varint,3,opt,name=TimeZone,proto3" json:"TimeZone,omitempty"` //时区
}

func (x *V2FootTeamHistoricalMatchRequest) Reset() {
	*x = V2FootTeamHistoricalMatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_team_historical_match_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootTeamHistoricalMatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootTeamHistoricalMatchRequest) ProtoMessage() {}

func (x *V2FootTeamHistoricalMatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_team_historical_match_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootTeamHistoricalMatchRequest.ProtoReflect.Descriptor instead.
func (*V2FootTeamHistoricalMatchRequest) Descriptor() ([]byte, []int) {
	return file_v2_foot_team_historical_match_proto_rawDescGZIP(), []int{0}
}

func (x *V2FootTeamHistoricalMatchRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *V2FootTeamHistoricalMatchRequest) GetMatchId() int64 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

func (x *V2FootTeamHistoricalMatchRequest) GetTimeZone() int64 {
	if x != nil {
		return x.TimeZone
	}
	return 0
}

type V2FootTeamHistoricalMatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*V2FootTeamHistoricalMatchInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *V2FootTeamHistoricalMatchResponse) Reset() {
	*x = V2FootTeamHistoricalMatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_team_historical_match_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootTeamHistoricalMatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootTeamHistoricalMatchResponse) ProtoMessage() {}

func (x *V2FootTeamHistoricalMatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_team_historical_match_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootTeamHistoricalMatchResponse.ProtoReflect.Descriptor instead.
func (*V2FootTeamHistoricalMatchResponse) Descriptor() ([]byte, []int) {
	return file_v2_foot_team_historical_match_proto_rawDescGZIP(), []int{1}
}

func (x *V2FootTeamHistoricalMatchResponse) GetList() []*V2FootTeamHistoricalMatchInfo {
	if x != nil {
		return x.List
	}
	return nil
}

type V2FootTeamHistoricalMatchInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchId       int64  `protobuf:"varint,1,opt,name=matchId,proto3" json:"matchId,omitempty"`              //比赛ID
	MatchTime     int64  `protobuf:"varint,2,opt,name=matchTime,proto3" json:"matchTime,omitempty"`          //比赛时间
	LeagueId      int64  `protobuf:"varint,3,opt,name=LeagueId,proto3" json:"LeagueId,omitempty"`            //联赛ID
	LeagueLogo    string `protobuf:"bytes,4,opt,name=LeagueLogo,proto3" json:"LeagueLogo,omitempty"`         //联赛logo
	LeagueName    string `protobuf:"bytes,5,opt,name=LeagueName,proto3" json:"LeagueName,omitempty"`         //联赛名称
	HomeId        int64  `protobuf:"varint,6,opt,name=homeId,proto3" json:"homeId,omitempty"`                //主队ID
	HomeLogo      string `protobuf:"bytes,7,opt,name=homeLogo,proto3" json:"homeLogo,omitempty"`             //主队logo
	Home          string `protobuf:"bytes,8,opt,name=home,proto3" json:"home,omitempty"`                     //主队名称
	HomeScore     int64  `protobuf:"varint,9,opt,name=homeScore,proto3" json:"homeScore,omitempty"`          //主队得分
	HomeHalfScore int64  `protobuf:"varint,10,opt,name=homeHalfScore,proto3" json:"homeHalfScore,omitempty"` //主队半场得分
	AwayId        int64  `protobuf:"varint,11,opt,name=awayId,proto3" json:"awayId,omitempty"`               //客队ID
	AwayLogo      string `protobuf:"bytes,12,opt,name=awayLogo,proto3" json:"awayLogo,omitempty"`            //客队logo
	Away          string `protobuf:"bytes,13,opt,name=away,proto3" json:"away,omitempty"`                    //客队名称
	AwayScore     int64  `protobuf:"varint,14,opt,name=awayScore,proto3" json:"awayScore,omitempty"`         //客队得分
	AwayHalfScore int64  `protobuf:"varint,15,opt,name=awayHalfScore,proto3" json:"awayHalfScore,omitempty"` //客队半场得分
	Result        int64  `protobuf:"varint,16,opt,name=result,proto3" json:"result,omitempty"`               //胜平负
	SameLeague    bool   `protobuf:"varint,17,opt,name=sameLeague,proto3" json:"sameLeague,omitempty"`       //相同联赛
	SameHome      bool   `protobuf:"varint,18,opt,name=sameHome,proto3" json:"sameHome,omitempty"`           //相同主队
}

func (x *V2FootTeamHistoricalMatchInfo) Reset() {
	*x = V2FootTeamHistoricalMatchInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_team_historical_match_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootTeamHistoricalMatchInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootTeamHistoricalMatchInfo) ProtoMessage() {}

func (x *V2FootTeamHistoricalMatchInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_team_historical_match_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootTeamHistoricalMatchInfo.ProtoReflect.Descriptor instead.
func (*V2FootTeamHistoricalMatchInfo) Descriptor() ([]byte, []int) {
	return file_v2_foot_team_historical_match_proto_rawDescGZIP(), []int{2}
}

func (x *V2FootTeamHistoricalMatchInfo) GetMatchId() int64 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

func (x *V2FootTeamHistoricalMatchInfo) GetMatchTime() int64 {
	if x != nil {
		return x.MatchTime
	}
	return 0
}

func (x *V2FootTeamHistoricalMatchInfo) GetLeagueId() int64 {
	if x != nil {
		return x.LeagueId
	}
	return 0
}

func (x *V2FootTeamHistoricalMatchInfo) GetLeagueLogo() string {
	if x != nil {
		return x.LeagueLogo
	}
	return ""
}

func (x *V2FootTeamHistoricalMatchInfo) GetLeagueName() string {
	if x != nil {
		return x.LeagueName
	}
	return ""
}

func (x *V2FootTeamHistoricalMatchInfo) GetHomeId() int64 {
	if x != nil {
		return x.HomeId
	}
	return 0
}

func (x *V2FootTeamHistoricalMatchInfo) GetHomeLogo() string {
	if x != nil {
		return x.HomeLogo
	}
	return ""
}

func (x *V2FootTeamHistoricalMatchInfo) GetHome() string {
	if x != nil {
		return x.Home
	}
	return ""
}

func (x *V2FootTeamHistoricalMatchInfo) GetHomeScore() int64 {
	if x != nil {
		return x.HomeScore
	}
	return 0
}

func (x *V2FootTeamHistoricalMatchInfo) GetHomeHalfScore() int64 {
	if x != nil {
		return x.HomeHalfScore
	}
	return 0
}

func (x *V2FootTeamHistoricalMatchInfo) GetAwayId() int64 {
	if x != nil {
		return x.AwayId
	}
	return 0
}

func (x *V2FootTeamHistoricalMatchInfo) GetAwayLogo() string {
	if x != nil {
		return x.AwayLogo
	}
	return ""
}

func (x *V2FootTeamHistoricalMatchInfo) GetAway() string {
	if x != nil {
		return x.Away
	}
	return ""
}

func (x *V2FootTeamHistoricalMatchInfo) GetAwayScore() int64 {
	if x != nil {
		return x.AwayScore
	}
	return 0
}

func (x *V2FootTeamHistoricalMatchInfo) GetAwayHalfScore() int64 {
	if x != nil {
		return x.AwayHalfScore
	}
	return 0
}

func (x *V2FootTeamHistoricalMatchInfo) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *V2FootTeamHistoricalMatchInfo) GetSameLeague() bool {
	if x != nil {
		return x.SameLeague
	}
	return false
}

func (x *V2FootTeamHistoricalMatchInfo) GetSameHome() bool {
	if x != nil {
		return x.SameHome
	}
	return false
}

var File_v2_foot_team_historical_match_proto protoreflect.FileDescriptor

var file_v2_foot_team_historical_match_proto_rawDesc = []byte{
	0x0a, 0x23, 0x76, 0x32, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x20, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x54,
	0x65, 0x61, 0x6d, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x57, 0x0a, 0x21, 0x56,
	0x32, 0x46, 0x6f, 0x6f, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69,
	0x63, 0x61, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x32, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x69, 0x63, 0x61, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0x9f, 0x04, 0x0a, 0x1d, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x54,
	0x65, 0x61, 0x6d, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x65,
	0x61, 0x67, 0x75, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x65,
	0x61, 0x67, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f,
	0x6d, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x6f, 0x6d, 0x65,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x6f, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x68, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x24, 0x0a, 0x0d, 0x68, 0x6f, 0x6d, 0x65, 0x48, 0x61, 0x6c, 0x66, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x68, 0x6f, 0x6d, 0x65, 0x48, 0x61, 0x6c,
	0x66, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x77, 0x61, 0x79, 0x49, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x77, 0x61, 0x79, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x77, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x77, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x77,
	0x61, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x77, 0x61, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x77, 0x61, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x61, 0x77, 0x61, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x61, 0x77, 0x61, 0x79, 0x48, 0x61, 0x6c, 0x66, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x77, 0x61, 0x79, 0x48, 0x61, 0x6c, 0x66, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x61,
	0x6d, 0x65, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x73, 0x61, 0x6d, 0x65, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x61,
	0x6d, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x61,
	0x6d, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_foot_team_historical_match_proto_rawDescOnce sync.Once
	file_v2_foot_team_historical_match_proto_rawDescData = file_v2_foot_team_historical_match_proto_rawDesc
)

func file_v2_foot_team_historical_match_proto_rawDescGZIP() []byte {
	file_v2_foot_team_historical_match_proto_rawDescOnce.Do(func() {
		file_v2_foot_team_historical_match_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_foot_team_historical_match_proto_rawDescData)
	})
	return file_v2_foot_team_historical_match_proto_rawDescData
}

var file_v2_foot_team_historical_match_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v2_foot_team_historical_match_proto_goTypes = []interface{}{
	(*V2FootTeamHistoricalMatchRequest)(nil),  // 0: V2FootTeamHistoricalMatchRequest
	(*V2FootTeamHistoricalMatchResponse)(nil), // 1: V2FootTeamHistoricalMatchResponse
	(*V2FootTeamHistoricalMatchInfo)(nil),     // 2: V2FootTeamHistoricalMatchInfo
}
var file_v2_foot_team_historical_match_proto_depIdxs = []int32{
	2, // 0: V2FootTeamHistoricalMatchResponse.list:type_name -> V2FootTeamHistoricalMatchInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v2_foot_team_historical_match_proto_init() }
func file_v2_foot_team_historical_match_proto_init() {
	if File_v2_foot_team_historical_match_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_foot_team_historical_match_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootTeamHistoricalMatchRequest); i {
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
		file_v2_foot_team_historical_match_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootTeamHistoricalMatchResponse); i {
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
		file_v2_foot_team_historical_match_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootTeamHistoricalMatchInfo); i {
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
			RawDescriptor: file_v2_foot_team_historical_match_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_foot_team_historical_match_proto_goTypes,
		DependencyIndexes: file_v2_foot_team_historical_match_proto_depIdxs,
		MessageInfos:      file_v2_foot_team_historical_match_proto_msgTypes,
	}.Build()
	File_v2_foot_team_historical_match_proto = out.File
	file_v2_foot_team_historical_match_proto_rawDesc = nil
	file_v2_foot_team_historical_match_proto_goTypes = nil
	file_v2_foot_team_historical_match_proto_depIdxs = nil
}
