// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: get_foot_data.proto

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
type CheckFootDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckType int64   `protobuf:"varint,2,opt,name=CheckType,proto3" json:"CheckType,omitempty"` //1：联赛 2：比赛 3：队伍 4：赛季
	Ids       []int64 `protobuf:"varint,3,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	IsHot     int64   `protobuf:"varint,1,opt,name=isHot,proto3" json:"isHot,omitempty"`
	Language  string  `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`  //请求语言  1:zh  2:en
	KeyWord   string  `protobuf:"bytes,5,opt,name=keyWord,proto3" json:"keyWord,omitempty"`    //球队筛选
	Page      int64   `protobuf:"varint,6,opt,name=page,proto3" json:"page,omitempty"`         //页
	PageSize  int64   `protobuf:"varint,7,opt,name=pageSize,proto3" json:"pageSize,omitempty"` //每页数量
}

func (x *CheckFootDataRequest) Reset() {
	*x = CheckFootDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_foot_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckFootDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckFootDataRequest) ProtoMessage() {}

func (x *CheckFootDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_get_foot_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckFootDataRequest.ProtoReflect.Descriptor instead.
func (*CheckFootDataRequest) Descriptor() ([]byte, []int) {
	return file_get_foot_data_proto_rawDescGZIP(), []int{0}
}

func (x *CheckFootDataRequest) GetCheckType() int64 {
	if x != nil {
		return x.CheckType
	}
	return 0
}

func (x *CheckFootDataRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *CheckFootDataRequest) GetIsHot() int64 {
	if x != nil {
		return x.IsHot
	}
	return 0
}

func (x *CheckFootDataRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *CheckFootDataRequest) GetKeyWord() string {
	if x != nil {
		return x.KeyWord
	}
	return ""
}

func (x *CheckFootDataRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *CheckFootDataRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type CheckFootDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComData   []*FootCheckLeagueData `protobuf:"bytes,1,rep,name=com_data,json=comData,proto3" json:"com_data,omitempty"`
	MatchData []*FootCheckMatchData  `protobuf:"bytes,2,rep,name=match_data,json=matchData,proto3" json:"match_data,omitempty"`
	TeamData  []*FootCheckTeamData   `protobuf:"bytes,3,rep,name=team_data,json=teamData,proto3" json:"team_data,omitempty"`
}

func (x *CheckFootDataResponse) Reset() {
	*x = CheckFootDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_foot_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckFootDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckFootDataResponse) ProtoMessage() {}

func (x *CheckFootDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_get_foot_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckFootDataResponse.ProtoReflect.Descriptor instead.
func (*CheckFootDataResponse) Descriptor() ([]byte, []int) {
	return file_get_foot_data_proto_rawDescGZIP(), []int{1}
}

func (x *CheckFootDataResponse) GetComData() []*FootCheckLeagueData {
	if x != nil {
		return x.ComData
	}
	return nil
}

func (x *CheckFootDataResponse) GetMatchData() []*FootCheckMatchData {
	if x != nil {
		return x.MatchData
	}
	return nil
}

func (x *CheckFootDataResponse) GetTeamData() []*FootCheckTeamData {
	if x != nil {
		return x.TeamData
	}
	return nil
}

type FootCheckLeagueData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	NameCh        string `protobuf:"bytes,2,opt,name=nameCh,proto3" json:"nameCh,omitempty"`
	Kind          int64  `protobuf:"varint,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Gender        int64  `protobuf:"varint,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Logo          string `protobuf:"bytes,5,opt,name=logo,proto3" json:"logo,omitempty"`
	Info          string `protobuf:"bytes,6,opt,name=info,proto3" json:"info,omitempty"`
	CountryName   string `protobuf:"bytes,7,opt,name=countryName,proto3" json:"countryName,omitempty"`
	CountryId     int64  `protobuf:"varint,13,opt,name=countryId,proto3" json:"countryId,omitempty"`
	CountryImage  string `protobuf:"bytes,16,opt,name=countryImage,proto3" json:"countryImage,omitempty"`
	ContinentName string `protobuf:"bytes,14,opt,name=continentName,proto3" json:"continentName,omitempty"`
	ContinentId   int64  `protobuf:"varint,15,opt,name=continentId,proto3" json:"continentId,omitempty"`
	Level         int64  `protobuf:"varint,8,opt,name=level,proto3" json:"level,omitempty"`
	PeriodType    int64  `protobuf:"varint,9,opt,name=periodType,proto3" json:"periodType,omitempty"`
	PeriodTime    int64  `protobuf:"varint,10,opt,name=periodTime,proto3" json:"periodTime,omitempty"`
	Reverse       int64  `protobuf:"varint,11,opt,name=reverse,proto3" json:"reverse,omitempty"`
	Grade         int64  `protobuf:"varint,12,opt,name=grade,proto3" json:"grade,omitempty"`
}

func (x *FootCheckLeagueData) Reset() {
	*x = FootCheckLeagueData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_foot_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootCheckLeagueData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootCheckLeagueData) ProtoMessage() {}

func (x *FootCheckLeagueData) ProtoReflect() protoreflect.Message {
	mi := &file_get_foot_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootCheckLeagueData.ProtoReflect.Descriptor instead.
func (*FootCheckLeagueData) Descriptor() ([]byte, []int) {
	return file_get_foot_data_proto_rawDescGZIP(), []int{2}
}

func (x *FootCheckLeagueData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FootCheckLeagueData) GetNameCh() string {
	if x != nil {
		return x.NameCh
	}
	return ""
}

func (x *FootCheckLeagueData) GetKind() int64 {
	if x != nil {
		return x.Kind
	}
	return 0
}

func (x *FootCheckLeagueData) GetGender() int64 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *FootCheckLeagueData) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *FootCheckLeagueData) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *FootCheckLeagueData) GetCountryName() string {
	if x != nil {
		return x.CountryName
	}
	return ""
}

func (x *FootCheckLeagueData) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *FootCheckLeagueData) GetCountryImage() string {
	if x != nil {
		return x.CountryImage
	}
	return ""
}

func (x *FootCheckLeagueData) GetContinentName() string {
	if x != nil {
		return x.ContinentName
	}
	return ""
}

func (x *FootCheckLeagueData) GetContinentId() int64 {
	if x != nil {
		return x.ContinentId
	}
	return 0
}

func (x *FootCheckLeagueData) GetLevel() int64 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *FootCheckLeagueData) GetPeriodType() int64 {
	if x != nil {
		return x.PeriodType
	}
	return 0
}

func (x *FootCheckLeagueData) GetPeriodTime() int64 {
	if x != nil {
		return x.PeriodTime
	}
	return 0
}

func (x *FootCheckLeagueData) GetReverse() int64 {
	if x != nil {
		return x.Reverse
	}
	return 0
}

func (x *FootCheckLeagueData) GetGrade() int64 {
	if x != nil {
		return x.Grade
	}
	return 0
}

type FootCheckMatchData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	HomeId                 int64  `protobuf:"varint,2,opt,name=homeId,proto3" json:"homeId,omitempty"`
	HomeName               string `protobuf:"bytes,3,opt,name=homeName,proto3" json:"homeName,omitempty"`
	AwayId                 int64  `protobuf:"varint,4,opt,name=awayId,proto3" json:"awayId,omitempty"`
	AwayName               string `protobuf:"bytes,5,opt,name=awayName,proto3" json:"awayName,omitempty"`
	TournamentId           int64  `protobuf:"varint,6,opt,name=tournamentId,proto3" json:"tournamentId,omitempty"`
	SeasonId               int64  `protobuf:"varint,7,opt,name=seasonId,proto3" json:"seasonId,omitempty"`
	MatchTime              int64  `protobuf:"varint,8,opt,name=matchTime,proto3" json:"matchTime,omitempty"`
	StartTime              int64  `protobuf:"varint,9,opt,name=startTime,proto3" json:"startTime,omitempty"`
	RoundType              string `protobuf:"bytes,23,opt,name=RoundType,proto3" json:"RoundType,omitempty"`
	RoundTypeId            int64  `protobuf:"varint,34,opt,name=RoundTypeId,proto3" json:"RoundTypeId,omitempty"`
	HomeScore              int64  `protobuf:"varint,10,opt,name=homeScore,proto3" json:"homeScore,omitempty"`
	AwayScore              int64  `protobuf:"varint,11,opt,name=awayScore,proto3" json:"awayScore,omitempty"`
	RedCards               int64  `protobuf:"varint,12,opt,name=redCards,proto3" json:"redCards,omitempty"`
	YellowCards            int64  `protobuf:"varint,13,opt,name=yellowCards,proto3" json:"yellowCards,omitempty"`
	CurrentPeriodStartTime int64  `protobuf:"varint,14,opt,name=currentPeriodStartTime,proto3" json:"currentPeriodStartTime,omitempty"`
	TimePlayed             int64  `protobuf:"varint,15,opt,name=timePlayed,proto3" json:"timePlayed,omitempty"`
	TimeRunning            int64  `protobuf:"varint,16,opt,name=timeRunning,proto3" json:"timeRunning,omitempty"`
	MatchStatus            int64  `protobuf:"varint,17,opt,name=matchStatus,proto3" json:"matchStatus,omitempty"`
	MatchStatusId          int64  `protobuf:"varint,18,opt,name=matchStatusId,proto3" json:"matchStatusId,omitempty"`
	MatchStatusResult      int64  `protobuf:"varint,19,opt,name=matchStatusResult,proto3" json:"matchStatusResult,omitempty"`
	HomeFormation          string `protobuf:"bytes,20,opt,name=homeFormation,proto3" json:"homeFormation,omitempty"`
	AwayFormation          string `protobuf:"bytes,21,opt,name=awayFormation,proto3" json:"awayFormation,omitempty"`
	Winner                 int64  `protobuf:"varint,22,opt,name=winner,proto3" json:"winner,omitempty"`
	UpTime                 int64  `protobuf:"varint,24,opt,name=upTime,proto3" json:"upTime,omitempty"`
	HomeLogo               string `protobuf:"bytes,25,opt,name=homeLogo,proto3" json:"homeLogo,omitempty"`
	AwayLogo               string `protobuf:"bytes,26,opt,name=awayLogo,proto3" json:"awayLogo,omitempty"`
}

func (x *FootCheckMatchData) Reset() {
	*x = FootCheckMatchData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_foot_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootCheckMatchData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootCheckMatchData) ProtoMessage() {}

func (x *FootCheckMatchData) ProtoReflect() protoreflect.Message {
	mi := &file_get_foot_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootCheckMatchData.ProtoReflect.Descriptor instead.
func (*FootCheckMatchData) Descriptor() ([]byte, []int) {
	return file_get_foot_data_proto_rawDescGZIP(), []int{3}
}

func (x *FootCheckMatchData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FootCheckMatchData) GetHomeId() int64 {
	if x != nil {
		return x.HomeId
	}
	return 0
}

func (x *FootCheckMatchData) GetHomeName() string {
	if x != nil {
		return x.HomeName
	}
	return ""
}

func (x *FootCheckMatchData) GetAwayId() int64 {
	if x != nil {
		return x.AwayId
	}
	return 0
}

func (x *FootCheckMatchData) GetAwayName() string {
	if x != nil {
		return x.AwayName
	}
	return ""
}

func (x *FootCheckMatchData) GetTournamentId() int64 {
	if x != nil {
		return x.TournamentId
	}
	return 0
}

func (x *FootCheckMatchData) GetSeasonId() int64 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

func (x *FootCheckMatchData) GetMatchTime() int64 {
	if x != nil {
		return x.MatchTime
	}
	return 0
}

func (x *FootCheckMatchData) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *FootCheckMatchData) GetRoundType() string {
	if x != nil {
		return x.RoundType
	}
	return ""
}

func (x *FootCheckMatchData) GetRoundTypeId() int64 {
	if x != nil {
		return x.RoundTypeId
	}
	return 0
}

func (x *FootCheckMatchData) GetHomeScore() int64 {
	if x != nil {
		return x.HomeScore
	}
	return 0
}

func (x *FootCheckMatchData) GetAwayScore() int64 {
	if x != nil {
		return x.AwayScore
	}
	return 0
}

func (x *FootCheckMatchData) GetRedCards() int64 {
	if x != nil {
		return x.RedCards
	}
	return 0
}

func (x *FootCheckMatchData) GetYellowCards() int64 {
	if x != nil {
		return x.YellowCards
	}
	return 0
}

func (x *FootCheckMatchData) GetCurrentPeriodStartTime() int64 {
	if x != nil {
		return x.CurrentPeriodStartTime
	}
	return 0
}

func (x *FootCheckMatchData) GetTimePlayed() int64 {
	if x != nil {
		return x.TimePlayed
	}
	return 0
}

func (x *FootCheckMatchData) GetTimeRunning() int64 {
	if x != nil {
		return x.TimeRunning
	}
	return 0
}

func (x *FootCheckMatchData) GetMatchStatus() int64 {
	if x != nil {
		return x.MatchStatus
	}
	return 0
}

func (x *FootCheckMatchData) GetMatchStatusId() int64 {
	if x != nil {
		return x.MatchStatusId
	}
	return 0
}

func (x *FootCheckMatchData) GetMatchStatusResult() int64 {
	if x != nil {
		return x.MatchStatusResult
	}
	return 0
}

func (x *FootCheckMatchData) GetHomeFormation() string {
	if x != nil {
		return x.HomeFormation
	}
	return ""
}

func (x *FootCheckMatchData) GetAwayFormation() string {
	if x != nil {
		return x.AwayFormation
	}
	return ""
}

func (x *FootCheckMatchData) GetWinner() int64 {
	if x != nil {
		return x.Winner
	}
	return 0
}

func (x *FootCheckMatchData) GetUpTime() int64 {
	if x != nil {
		return x.UpTime
	}
	return 0
}

func (x *FootCheckMatchData) GetHomeLogo() string {
	if x != nil {
		return x.HomeLogo
	}
	return ""
}

func (x *FootCheckMatchData) GetAwayLogo() string {
	if x != nil {
		return x.AwayLogo
	}
	return ""
}

type FootCheckTeamData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`              //当前记录ID
	TeamName  string `protobuf:"bytes,2,opt,name=TeamName,proto3" json:"TeamName,omitempty"`   //队伍名称
	TeamImage string `protobuf:"bytes,3,opt,name=TeamImage,proto3" json:"TeamImage,omitempty"` //队伍图片
	IsFollow  string `protobuf:"bytes,4,opt,name=IsFollow,proto3" json:"IsFollow,omitempty"`   //是否关注  0否 1 是
}

func (x *FootCheckTeamData) Reset() {
	*x = FootCheckTeamData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_foot_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootCheckTeamData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootCheckTeamData) ProtoMessage() {}

func (x *FootCheckTeamData) ProtoReflect() protoreflect.Message {
	mi := &file_get_foot_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootCheckTeamData.ProtoReflect.Descriptor instead.
func (*FootCheckTeamData) Descriptor() ([]byte, []int) {
	return file_get_foot_data_proto_rawDescGZIP(), []int{4}
}

func (x *FootCheckTeamData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FootCheckTeamData) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

func (x *FootCheckTeamData) GetTeamImage() string {
	if x != nil {
		return x.TeamImage
	}
	return ""
}

func (x *FootCheckTeamData) GetIsFollow() string {
	if x != nil {
		return x.IsFollow
	}
	return ""
}

var File_get_foot_data_proto protoreflect.FileDescriptor

var file_get_foot_data_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x65, 0x74, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x14, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46,
	0x6f, 0x6f, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x73, 0x48, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69,
	0x73, 0x48, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xad, 0x01, 0x0a, 0x15, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x46, 0x6f, 0x6f, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x46, 0x6f, 0x6f, 0x74,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x09, 0x74, 0x65, 0x61,
	0x6d, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x46,
	0x6f, 0x6f, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x22, 0xc3, 0x03, 0x0a, 0x13, 0x46,
	0x6f, 0x6f, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6e, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x69,
	0x6e, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x69, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x22, 0xe6, 0x06, 0x0a, 0x12, 0x46, 0x6f, 0x6f, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f, 0x6d, 0x65, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x6f, 0x6d, 0x65, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x77, 0x61, 0x79, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x77, 0x61,
	0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x77, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x77, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x52,
	0x6f, 0x75, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x52, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x6f, 0x75,
	0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x22, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x52, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x68,
	0x6f, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x68, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x77, 0x61,
	0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x77,
	0x61, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x64, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x64, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x79, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x61, 0x72,
	0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x79, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x43, 0x61, 0x72, 0x64, 0x73, 0x12, 0x36, 0x0a, 0x16, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x74, 0x69, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12,
	0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x49, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x11, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x6f, 0x6d, 0x65, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x68, 0x6f,
	0x6d, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x61,
	0x77, 0x61, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x61, 0x77, 0x61, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x70, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x19, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x77, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x77, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x6f, 0x22, 0x79, 0x0a, 0x11, 0x46, 0x6f, 0x6f,
	0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x65,
	0x61, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54,
	0x65, 0x61, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x73, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_get_foot_data_proto_rawDescOnce sync.Once
	file_get_foot_data_proto_rawDescData = file_get_foot_data_proto_rawDesc
)

func file_get_foot_data_proto_rawDescGZIP() []byte {
	file_get_foot_data_proto_rawDescOnce.Do(func() {
		file_get_foot_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_get_foot_data_proto_rawDescData)
	})
	return file_get_foot_data_proto_rawDescData
}

var file_get_foot_data_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_get_foot_data_proto_goTypes = []interface{}{
	(*CheckFootDataRequest)(nil),  // 0: CheckFootDataRequest
	(*CheckFootDataResponse)(nil), // 1: CheckFootDataResponse
	(*FootCheckLeagueData)(nil),   // 2: FootCheckLeagueData
	(*FootCheckMatchData)(nil),    // 3: FootCheckMatchData
	(*FootCheckTeamData)(nil),     // 4: FootCheckTeamData
}
var file_get_foot_data_proto_depIdxs = []int32{
	2, // 0: CheckFootDataResponse.com_data:type_name -> FootCheckLeagueData
	3, // 1: CheckFootDataResponse.match_data:type_name -> FootCheckMatchData
	4, // 2: CheckFootDataResponse.team_data:type_name -> FootCheckTeamData
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_get_foot_data_proto_init() }
func file_get_foot_data_proto_init() {
	if File_get_foot_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_get_foot_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckFootDataRequest); i {
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
		file_get_foot_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckFootDataResponse); i {
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
		file_get_foot_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootCheckLeagueData); i {
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
		file_get_foot_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootCheckMatchData); i {
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
		file_get_foot_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootCheckTeamData); i {
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
			RawDescriptor: file_get_foot_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_get_foot_data_proto_goTypes,
		DependencyIndexes: file_get_foot_data_proto_depIdxs,
		MessageInfos:      file_get_foot_data_proto_msgTypes,
	}.Build()
	File_get_foot_data_proto = out.File
	file_get_foot_data_proto_rawDesc = nil
	file_get_foot_data_proto_goTypes = nil
	file_get_foot_data_proto_depIdxs = nil
}
