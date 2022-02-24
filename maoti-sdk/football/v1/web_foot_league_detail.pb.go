// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: web_foot_league_detail.proto

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

//web联赛详情
type WebFootLeagueDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeagueId int64  `protobuf:"varint,1,opt,name=LeagueId,proto3" json:"LeagueId,omitempty"` //联赛id
	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`  //请求语言
}

func (x *WebFootLeagueDetailRequest) Reset() {
	*x = WebFootLeagueDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_league_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootLeagueDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootLeagueDetailRequest) ProtoMessage() {}

func (x *WebFootLeagueDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_league_detail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootLeagueDetailRequest.ProtoReflect.Descriptor instead.
func (*WebFootLeagueDetailRequest) Descriptor() ([]byte, []int) {
	return file_web_foot_league_detail_proto_rawDescGZIP(), []int{0}
}

func (x *WebFootLeagueDetailRequest) GetLeagueId() int64 {
	if x != nil {
		return x.LeagueId
	}
	return 0
}

func (x *WebFootLeagueDetailRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type WebFootLeagueDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeagueId    int64                      `protobuf:"varint,1,opt,name=LeagueId,proto3" json:"LeagueId,omitempty"`      //联赛id
	LeagueName  string                     `protobuf:"bytes,3,opt,name=LeagueName,proto3" json:"LeagueName,omitempty"`   //联赛名称
	LeagueType  int64                      `protobuf:"varint,9,opt,name=LeagueType,proto3" json:"LeagueType,omitempty"`  //联赛类型：1.联赛；2.杯赛
	Logo        string                     `protobuf:"bytes,4,opt,name=Logo,proto3" json:"Logo,omitempty"`               //联赛图片
	CountryId   int64                      `protobuf:"varint,5,opt,name=CountryId,proto3" json:"CountryId,omitempty"`    //联赛国家ID
	CountryName string                     `protobuf:"bytes,8,opt,name=CountryName,proto3" json:"CountryName,omitempty"` //联赛国家
	CountryLogo string                     `protobuf:"bytes,6,opt,name=CountryLogo,proto3" json:"CountryLogo,omitempty"` //联赛国家图片
	Season      []*WebFootLeagueSeasonList `protobuf:"bytes,7,rep,name=season,proto3" json:"season,omitempty"`           //联赛下的赛季列表
}

func (x *WebFootLeagueDetailResponse) Reset() {
	*x = WebFootLeagueDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_league_detail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootLeagueDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootLeagueDetailResponse) ProtoMessage() {}

func (x *WebFootLeagueDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_league_detail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootLeagueDetailResponse.ProtoReflect.Descriptor instead.
func (*WebFootLeagueDetailResponse) Descriptor() ([]byte, []int) {
	return file_web_foot_league_detail_proto_rawDescGZIP(), []int{1}
}

func (x *WebFootLeagueDetailResponse) GetLeagueId() int64 {
	if x != nil {
		return x.LeagueId
	}
	return 0
}

func (x *WebFootLeagueDetailResponse) GetLeagueName() string {
	if x != nil {
		return x.LeagueName
	}
	return ""
}

func (x *WebFootLeagueDetailResponse) GetLeagueType() int64 {
	if x != nil {
		return x.LeagueType
	}
	return 0
}

func (x *WebFootLeagueDetailResponse) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *WebFootLeagueDetailResponse) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *WebFootLeagueDetailResponse) GetCountryName() string {
	if x != nil {
		return x.CountryName
	}
	return ""
}

func (x *WebFootLeagueDetailResponse) GetCountryLogo() string {
	if x != nil {
		return x.CountryLogo
	}
	return ""
}

func (x *WebFootLeagueDetailResponse) GetSeason() []*WebFootLeagueSeasonList {
	if x != nil {
		return x.Season
	}
	return nil
}

type WebFootLeagueSeasonList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeasonId   int64  `protobuf:"varint,1,opt,name=SeasonId,proto3" json:"SeasonId,omitempty"`    //赛季id
	SeasonName string `protobuf:"bytes,2,opt,name=SeasonName,proto3" json:"SeasonName,omitempty"` //赛季名称
}

func (x *WebFootLeagueSeasonList) Reset() {
	*x = WebFootLeagueSeasonList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_league_detail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootLeagueSeasonList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootLeagueSeasonList) ProtoMessage() {}

func (x *WebFootLeagueSeasonList) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_league_detail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootLeagueSeasonList.ProtoReflect.Descriptor instead.
func (*WebFootLeagueSeasonList) Descriptor() ([]byte, []int) {
	return file_web_foot_league_detail_proto_rawDescGZIP(), []int{2}
}

func (x *WebFootLeagueSeasonList) GetSeasonId() int64 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

func (x *WebFootLeagueSeasonList) GetSeasonName() string {
	if x != nil {
		return x.SeasonName
	}
	return ""
}

//web联赛详情 - 统计数据侧边栏
type WebFootLeagueStatsMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"` //请求语言
	MenuType string `protobuf:"bytes,4,opt,name=MenuType,proto3" json:"MenuType,omitempty"` //侧边栏类型  team:查询球队的侧边栏数据，player查询球员的侧边栏数据
}

func (x *WebFootLeagueStatsMenuRequest) Reset() {
	*x = WebFootLeagueStatsMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_league_detail_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootLeagueStatsMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootLeagueStatsMenuRequest) ProtoMessage() {}

func (x *WebFootLeagueStatsMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_league_detail_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootLeagueStatsMenuRequest.ProtoReflect.Descriptor instead.
func (*WebFootLeagueStatsMenuRequest) Descriptor() ([]byte, []int) {
	return file_web_foot_league_detail_proto_rawDescGZIP(), []int{3}
}

func (x *WebFootLeagueStatsMenuRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *WebFootLeagueStatsMenuRequest) GetMenuType() string {
	if x != nil {
		return x.MenuType
	}
	return ""
}

type WebFootLeagueStatsMenuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*WebFootLeagueStatsMenu `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *WebFootLeagueStatsMenuResponse) Reset() {
	*x = WebFootLeagueStatsMenuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_league_detail_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootLeagueStatsMenuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootLeagueStatsMenuResponse) ProtoMessage() {}

func (x *WebFootLeagueStatsMenuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_league_detail_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootLeagueStatsMenuResponse.ProtoReflect.Descriptor instead.
func (*WebFootLeagueStatsMenuResponse) Descriptor() ([]byte, []int) {
	return file_web_foot_league_detail_proto_rawDescGZIP(), []int{4}
}

func (x *WebFootLeagueStatsMenuResponse) GetData() []*WebFootLeagueStatsMenu {
	if x != nil {
		return x.Data
	}
	return nil
}

type WebFootLeagueStatsMenu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mark     string `protobuf:"bytes,2,opt,name=mark,proto3" json:"mark,omitempty"`                         //标记
	MarkName string `protobuf:"bytes,3,opt,name=mark_name,json=markName,proto3" json:"mark_name,omitempty"` //标记名称
}

func (x *WebFootLeagueStatsMenu) Reset() {
	*x = WebFootLeagueStatsMenu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_league_detail_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootLeagueStatsMenu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootLeagueStatsMenu) ProtoMessage() {}

func (x *WebFootLeagueStatsMenu) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_league_detail_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootLeagueStatsMenu.ProtoReflect.Descriptor instead.
func (*WebFootLeagueStatsMenu) Descriptor() ([]byte, []int) {
	return file_web_foot_league_detail_proto_rawDescGZIP(), []int{5}
}

func (x *WebFootLeagueStatsMenu) GetMark() string {
	if x != nil {
		return x.Mark
	}
	return ""
}

func (x *WebFootLeagueStatsMenu) GetMarkName() string {
	if x != nil {
		return x.MarkName
	}
	return ""
}

//web联赛详情 - 联赛赛季下的球队列表
type WebFootLeagueTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeagueId int64  `protobuf:"varint,1,opt,name=LeagueId,proto3" json:"LeagueId,omitempty"` //联赛id
	SeasonId int64  `protobuf:"varint,2,opt,name=SeasonId,proto3" json:"SeasonId,omitempty"` //赛季id
	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`  //请求语言
}

func (x *WebFootLeagueTeamRequest) Reset() {
	*x = WebFootLeagueTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_league_detail_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootLeagueTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootLeagueTeamRequest) ProtoMessage() {}

func (x *WebFootLeagueTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_league_detail_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootLeagueTeamRequest.ProtoReflect.Descriptor instead.
func (*WebFootLeagueTeamRequest) Descriptor() ([]byte, []int) {
	return file_web_foot_league_detail_proto_rawDescGZIP(), []int{6}
}

func (x *WebFootLeagueTeamRequest) GetLeagueId() int64 {
	if x != nil {
		return x.LeagueId
	}
	return 0
}

func (x *WebFootLeagueTeamRequest) GetSeasonId() int64 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

func (x *WebFootLeagueTeamRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type WebFootLeagueTeamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*WebFootLeagueTeam `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *WebFootLeagueTeamResponse) Reset() {
	*x = WebFootLeagueTeamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_league_detail_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootLeagueTeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootLeagueTeamResponse) ProtoMessage() {}

func (x *WebFootLeagueTeamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_league_detail_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootLeagueTeamResponse.ProtoReflect.Descriptor instead.
func (*WebFootLeagueTeamResponse) Descriptor() ([]byte, []int) {
	return file_web_foot_league_detail_proto_rawDescGZIP(), []int{7}
}

func (x *WebFootLeagueTeamResponse) GetData() []*WebFootLeagueTeam {
	if x != nil {
		return x.Data
	}
	return nil
}

type WebFootLeagueTeam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId   int64  `protobuf:"varint,1,opt,name=TeamId,proto3" json:"TeamId,omitempty"`    //球队id
	TeamName string `protobuf:"bytes,2,opt,name=TeamName,proto3" json:"TeamName,omitempty"` //球队名称
}

func (x *WebFootLeagueTeam) Reset() {
	*x = WebFootLeagueTeam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_league_detail_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootLeagueTeam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootLeagueTeam) ProtoMessage() {}

func (x *WebFootLeagueTeam) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_league_detail_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootLeagueTeam.ProtoReflect.Descriptor instead.
func (*WebFootLeagueTeam) Descriptor() ([]byte, []int) {
	return file_web_foot_league_detail_proto_rawDescGZIP(), []int{8}
}

func (x *WebFootLeagueTeam) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *WebFootLeagueTeam) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

//web联赛详情 - 球队球员统计数据
type WebFootLeagueStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language  string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`   //请求语言
	LeagueId  int64  `protobuf:"varint,1,opt,name=LeagueId,proto3" json:"LeagueId,omitempty"`  //联赛id
	SeasonId  int64  `protobuf:"varint,7,opt,name=seasonId,proto3" json:"seasonId,omitempty"`  //赛季ID
	StatsType string `protobuf:"bytes,4,opt,name=statsType,proto3" json:"statsType,omitempty"` //统计类型  team:查询球队的侧边栏数据，player查询球员的侧边栏数据
	TeamId    int64  `protobuf:"varint,5,opt,name=teamId,proto3" json:"teamId,omitempty"`      //球队ID
	Mark      string `protobuf:"bytes,6,opt,name=mark,proto3" json:"mark,omitempty"`           //标记
}

func (x *WebFootLeagueStatsRequest) Reset() {
	*x = WebFootLeagueStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_league_detail_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootLeagueStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootLeagueStatsRequest) ProtoMessage() {}

func (x *WebFootLeagueStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_league_detail_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootLeagueStatsRequest.ProtoReflect.Descriptor instead.
func (*WebFootLeagueStatsRequest) Descriptor() ([]byte, []int) {
	return file_web_foot_league_detail_proto_rawDescGZIP(), []int{9}
}

func (x *WebFootLeagueStatsRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *WebFootLeagueStatsRequest) GetLeagueId() int64 {
	if x != nil {
		return x.LeagueId
	}
	return 0
}

func (x *WebFootLeagueStatsRequest) GetSeasonId() int64 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

func (x *WebFootLeagueStatsRequest) GetStatsType() string {
	if x != nil {
		return x.StatsType
	}
	return ""
}

func (x *WebFootLeagueStatsRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *WebFootLeagueStatsRequest) GetMark() string {
	if x != nil {
		return x.Mark
	}
	return ""
}

type WebFootLeagueStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*WebFootLeagueStats `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *WebFootLeagueStatsResponse) Reset() {
	*x = WebFootLeagueStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_league_detail_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootLeagueStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootLeagueStatsResponse) ProtoMessage() {}

func (x *WebFootLeagueStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_league_detail_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootLeagueStatsResponse.ProtoReflect.Descriptor instead.
func (*WebFootLeagueStatsResponse) Descriptor() ([]byte, []int) {
	return file_web_foot_league_detail_proto_rawDescGZIP(), []int{10}
}

func (x *WebFootLeagueStatsResponse) GetData() []*WebFootLeagueStats {
	if x != nil {
		return x.Data
	}
	return nil
}

type WebFootLeagueStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rank   int64   `protobuf:"varint,1,opt,name=Rank,proto3" json:"Rank,omitempty"`      //排名，后端已实现排序，前端根据数据顺序给出下标即可
	Id     int64   `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`          //球队id/球员id
	Name   string  `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`       //球队名称/球员名称
	Data   float32 `protobuf:"fixed32,4,opt,name=data,proto3" json:"data,omitempty"`     //数据
	Corner float32 `protobuf:"fixed32,5,opt,name=corner,proto3" json:"corner,omitempty"` //角球，当查询进球数据时，该字段才可能有值
	Logo   string  `protobuf:"bytes,6,opt,name=Logo,proto3" json:"Logo,omitempty"`       //球员球队图片
}

func (x *WebFootLeagueStats) Reset() {
	*x = WebFootLeagueStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_league_detail_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootLeagueStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootLeagueStats) ProtoMessage() {}

func (x *WebFootLeagueStats) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_league_detail_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootLeagueStats.ProtoReflect.Descriptor instead.
func (*WebFootLeagueStats) Descriptor() ([]byte, []int) {
	return file_web_foot_league_detail_proto_rawDescGZIP(), []int{11}
}

func (x *WebFootLeagueStats) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *WebFootLeagueStats) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WebFootLeagueStats) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WebFootLeagueStats) GetData() float32 {
	if x != nil {
		return x.Data
	}
	return 0
}

func (x *WebFootLeagueStats) GetCorner() float32 {
	if x != nil {
		return x.Corner
	}
	return 0
}

func (x *WebFootLeagueStats) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

var File_web_foot_league_detail_proto protoreflect.FileDescriptor

var file_web_foot_league_detail_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x77, 0x65, 0x62, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6c, 0x65, 0x61, 0x67, 0x75,
	0x65, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54,
	0x0a, 0x1a, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x22, 0xa1, 0x02, 0x0a, 0x1b, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74,
	0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4c,
	0x6f, 0x67, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74,
	0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x06, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x55, 0x0a, 0x17, 0x57, 0x65, 0x62, 0x46,
	0x6f, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x57, 0x0a, 0x1d, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x4d, 0x65, 0x6e, 0x75, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x4d, 0x65, 0x6e, 0x75, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4d, 0x0a, 0x1e, 0x57, 0x65, 0x62, 0x46,
	0x6f, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x4d, 0x65,
	0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x57, 0x65, 0x62, 0x46, 0x6f,
	0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x4d, 0x65, 0x6e,
	0x75, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x49, 0x0a, 0x16, 0x57, 0x65, 0x62, 0x46, 0x6f,
	0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x4d, 0x65, 0x6e,
	0x75, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x72, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x6e, 0x0a, 0x18, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x65, 0x61,
	0x67, 0x75, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x22, 0x43, 0x0a, 0x19, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x65, 0x61,
	0x67, 0x75, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x54, 0x65, 0x61,
	0x6d, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x47, 0x0a, 0x11, 0x57, 0x65, 0x62, 0x46, 0x6f,
	0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x06,
	0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x54, 0x65,
	0x61, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0xb9, 0x01, 0x0a, 0x19, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67,
	0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x65,
	0x61, 0x67, 0x75, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x4c, 0x65,
	0x61, 0x67, 0x75, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x45, 0x0a, 0x1a,
	0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x57, 0x65, 0x62, 0x46, 0x6f,
	0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x8c, 0x01, 0x0a, 0x12, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4c,
	0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x61,
	0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x72, 0x6e, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x63, 0x6f, 0x72, 0x6e, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x6f,
	0x67, 0x6f, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_web_foot_league_detail_proto_rawDescOnce sync.Once
	file_web_foot_league_detail_proto_rawDescData = file_web_foot_league_detail_proto_rawDesc
)

func file_web_foot_league_detail_proto_rawDescGZIP() []byte {
	file_web_foot_league_detail_proto_rawDescOnce.Do(func() {
		file_web_foot_league_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_foot_league_detail_proto_rawDescData)
	})
	return file_web_foot_league_detail_proto_rawDescData
}

var file_web_foot_league_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_web_foot_league_detail_proto_goTypes = []interface{}{
	(*WebFootLeagueDetailRequest)(nil),     // 0: WebFootLeagueDetailRequest
	(*WebFootLeagueDetailResponse)(nil),    // 1: WebFootLeagueDetailResponse
	(*WebFootLeagueSeasonList)(nil),        // 2: WebFootLeagueSeasonList
	(*WebFootLeagueStatsMenuRequest)(nil),  // 3: WebFootLeagueStatsMenuRequest
	(*WebFootLeagueStatsMenuResponse)(nil), // 4: WebFootLeagueStatsMenuResponse
	(*WebFootLeagueStatsMenu)(nil),         // 5: WebFootLeagueStatsMenu
	(*WebFootLeagueTeamRequest)(nil),       // 6: WebFootLeagueTeamRequest
	(*WebFootLeagueTeamResponse)(nil),      // 7: WebFootLeagueTeamResponse
	(*WebFootLeagueTeam)(nil),              // 8: WebFootLeagueTeam
	(*WebFootLeagueStatsRequest)(nil),      // 9: WebFootLeagueStatsRequest
	(*WebFootLeagueStatsResponse)(nil),     // 10: WebFootLeagueStatsResponse
	(*WebFootLeagueStats)(nil),             // 11: WebFootLeagueStats
}
var file_web_foot_league_detail_proto_depIdxs = []int32{
	2,  // 0: WebFootLeagueDetailResponse.season:type_name -> WebFootLeagueSeasonList
	5,  // 1: WebFootLeagueStatsMenuResponse.Data:type_name -> WebFootLeagueStatsMenu
	8,  // 2: WebFootLeagueTeamResponse.Data:type_name -> WebFootLeagueTeam
	11, // 3: WebFootLeagueStatsResponse.Data:type_name -> WebFootLeagueStats
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_web_foot_league_detail_proto_init() }
func file_web_foot_league_detail_proto_init() {
	if File_web_foot_league_detail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_foot_league_detail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootLeagueDetailRequest); i {
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
		file_web_foot_league_detail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootLeagueDetailResponse); i {
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
		file_web_foot_league_detail_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootLeagueSeasonList); i {
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
		file_web_foot_league_detail_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootLeagueStatsMenuRequest); i {
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
		file_web_foot_league_detail_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootLeagueStatsMenuResponse); i {
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
		file_web_foot_league_detail_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootLeagueStatsMenu); i {
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
		file_web_foot_league_detail_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootLeagueTeamRequest); i {
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
		file_web_foot_league_detail_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootLeagueTeamResponse); i {
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
		file_web_foot_league_detail_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootLeagueTeam); i {
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
		file_web_foot_league_detail_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootLeagueStatsRequest); i {
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
		file_web_foot_league_detail_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootLeagueStatsResponse); i {
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
		file_web_foot_league_detail_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootLeagueStats); i {
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
			RawDescriptor: file_web_foot_league_detail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_foot_league_detail_proto_goTypes,
		DependencyIndexes: file_web_foot_league_detail_proto_depIdxs,
		MessageInfos:      file_web_foot_league_detail_proto_msgTypes,
	}.Build()
	File_web_foot_league_detail_proto = out.File
	file_web_foot_league_detail_proto_rawDesc = nil
	file_web_foot_league_detail_proto_goTypes = nil
	file_web_foot_league_detail_proto_depIdxs = nil
}
