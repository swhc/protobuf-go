// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: v2_foot_player_overview.proto

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

// 基本信息 (接口)
type FootPlayerInfoBasicsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonId   int64  `protobuf:"varint,1,opt,name=PersonId,proto3" json:"PersonId,omitempty"`    //球员id
	Image      string `protobuf:"bytes,2,opt,name=Image,proto3" json:"Image,omitempty"`           //头像
	PersonName string `protobuf:"bytes,3,opt,name=PersonName,proto3" json:"PersonName,omitempty"` //球员名称 根据language
	TeamId     int64  `protobuf:"varint,5,opt,name=TeamId,proto3" json:"TeamId,omitempty"`        //球队id
	CountryId  int64  `protobuf:"varint,6,opt,name=CountryId,proto3" json:"CountryId,omitempty"`  //球员国家id
	Age        int64  `protobuf:"varint,7,opt,name=Age,proto3" json:"Age,omitempty"`              //球员年龄
	Height     string `protobuf:"bytes,8,opt,name=Height,proto3" json:"Height,omitempty"`         //身高
	Weight     string `protobuf:"bytes,9,opt,name=Weight,proto3" json:"Weight,omitempty"`         //体重
	Value      string `protobuf:"bytes,10,opt,name=Value,proto3" json:"Value,omitempty"`          //身价
}

func (x *FootPlayerInfoBasicsResp) Reset() {
	*x = FootPlayerInfoBasicsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_player_overview_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootPlayerInfoBasicsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootPlayerInfoBasicsResp) ProtoMessage() {}

func (x *FootPlayerInfoBasicsResp) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_player_overview_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootPlayerInfoBasicsResp.ProtoReflect.Descriptor instead.
func (*FootPlayerInfoBasicsResp) Descriptor() ([]byte, []int) {
	return file_v2_foot_player_overview_proto_rawDescGZIP(), []int{0}
}

func (x *FootPlayerInfoBasicsResp) GetPersonId() int64 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

func (x *FootPlayerInfoBasicsResp) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *FootPlayerInfoBasicsResp) GetPersonName() string {
	if x != nil {
		return x.PersonName
	}
	return ""
}

func (x *FootPlayerInfoBasicsResp) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *FootPlayerInfoBasicsResp) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *FootPlayerInfoBasicsResp) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *FootPlayerInfoBasicsResp) GetHeight() string {
	if x != nil {
		return x.Height
	}
	return ""
}

func (x *FootPlayerInfoBasicsResp) GetWeight() string {
	if x != nil {
		return x.Weight
	}
	return ""
}

func (x *FootPlayerInfoBasicsResp) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// 球员特征、位置、首发胜率 (接口)
type FootPlayerInfoFeaturesAndPositionOnThePitchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeaturesScore string `protobuf:"bytes,1,opt,name=FeaturesScore,proto3" json:"FeaturesScore,omitempty"` // 特征评分 进攻力,创造力,防守力,战术力,技术力
	Strengths     string `protobuf:"bytes,2,opt,name=Strengths,proto3" json:"Strengths,omitempty"`         // 优点
	Weaknesses    string `protobuf:"bytes,3,opt,name=Weaknesses,proto3" json:"Weaknesses,omitempty"`       // 缺点
	Score         string `protobuf:"bytes,4,opt,name=Score,proto3" json:"Score,omitempty"`                 // 总评分
	StartedWins   string `protobuf:"bytes,5,opt,name=StartedWins,proto3" json:"StartedWins,omitempty"`     // 首发胜率
	Positions     string `protobuf:"bytes,6,opt,name=positions,proto3" json:"positions,omitempty"`         // 位置
}

func (x *FootPlayerInfoFeaturesAndPositionOnThePitchResp) Reset() {
	*x = FootPlayerInfoFeaturesAndPositionOnThePitchResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_player_overview_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootPlayerInfoFeaturesAndPositionOnThePitchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootPlayerInfoFeaturesAndPositionOnThePitchResp) ProtoMessage() {}

func (x *FootPlayerInfoFeaturesAndPositionOnThePitchResp) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_player_overview_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootPlayerInfoFeaturesAndPositionOnThePitchResp.ProtoReflect.Descriptor instead.
func (*FootPlayerInfoFeaturesAndPositionOnThePitchResp) Descriptor() ([]byte, []int) {
	return file_v2_foot_player_overview_proto_rawDescGZIP(), []int{1}
}

func (x *FootPlayerInfoFeaturesAndPositionOnThePitchResp) GetFeaturesScore() string {
	if x != nil {
		return x.FeaturesScore
	}
	return ""
}

func (x *FootPlayerInfoFeaturesAndPositionOnThePitchResp) GetStrengths() string {
	if x != nil {
		return x.Strengths
	}
	return ""
}

func (x *FootPlayerInfoFeaturesAndPositionOnThePitchResp) GetWeaknesses() string {
	if x != nil {
		return x.Weaknesses
	}
	return ""
}

func (x *FootPlayerInfoFeaturesAndPositionOnThePitchResp) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

func (x *FootPlayerInfoFeaturesAndPositionOnThePitchResp) GetStartedWins() string {
	if x != nil {
		return x.StartedWins
	}
	return ""
}

func (x *FootPlayerInfoFeaturesAndPositionOnThePitchResp) GetPositions() string {
	if x != nil {
		return x.Positions
	}
	return ""
}

// 球员历史效力俱乐部表现列表 (接口)
type FootPlayerInfoHistoricalPerformanceInClubsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HistoricalPerformanceInClubs []*HistoricalPerformanceInClubsItem `protobuf:"bytes,1,rep,name=HistoricalPerformanceInClubs,proto3" json:"HistoricalPerformanceInClubs,omitempty"`
}

func (x *FootPlayerInfoHistoricalPerformanceInClubsResp) Reset() {
	*x = FootPlayerInfoHistoricalPerformanceInClubsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_player_overview_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootPlayerInfoHistoricalPerformanceInClubsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootPlayerInfoHistoricalPerformanceInClubsResp) ProtoMessage() {}

func (x *FootPlayerInfoHistoricalPerformanceInClubsResp) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_player_overview_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootPlayerInfoHistoricalPerformanceInClubsResp.ProtoReflect.Descriptor instead.
func (*FootPlayerInfoHistoricalPerformanceInClubsResp) Descriptor() ([]byte, []int) {
	return file_v2_foot_player_overview_proto_rawDescGZIP(), []int{2}
}

func (x *FootPlayerInfoHistoricalPerformanceInClubsResp) GetHistoricalPerformanceInClubs() []*HistoricalPerformanceInClubsItem {
	if x != nil {
		return x.HistoricalPerformanceInClubs
	}
	return nil
}

type HistoricalPerformanceInClubsItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId      int64 `protobuf:"varint,1,opt,name=TeamId,proto3" json:"TeamId,omitempty"`           // 球队id
	SeasonCount int64 `protobuf:"varint,2,opt,name=SeasonCount,proto3" json:"SeasonCount,omitempty"` // 参与赛季数
	P           int64 `protobuf:"varint,3,opt,name=P,proto3" json:"P,omitempty"`                     // 出场次数
	Goals       int64 `protobuf:"varint,4,opt,name=Goals,proto3" json:"Goals,omitempty"`             // 进球数
	Assists     int64 `protobuf:"varint,5,opt,name=Assists,proto3" json:"Assists,omitempty"`         // 助攻数
	YellowCard  int64 `protobuf:"varint,6,opt,name=YellowCard,proto3" json:"YellowCard,omitempty"`   // 黄牌数
	RedCard     int64 `protobuf:"varint,7,opt,name=RedCard,proto3" json:"RedCard,omitempty"`         // 红牌数
}

func (x *HistoricalPerformanceInClubsItem) Reset() {
	*x = HistoricalPerformanceInClubsItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_player_overview_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoricalPerformanceInClubsItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoricalPerformanceInClubsItem) ProtoMessage() {}

func (x *HistoricalPerformanceInClubsItem) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_player_overview_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoricalPerformanceInClubsItem.ProtoReflect.Descriptor instead.
func (*HistoricalPerformanceInClubsItem) Descriptor() ([]byte, []int) {
	return file_v2_foot_player_overview_proto_rawDescGZIP(), []int{3}
}

func (x *HistoricalPerformanceInClubsItem) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *HistoricalPerformanceInClubsItem) GetSeasonCount() int64 {
	if x != nil {
		return x.SeasonCount
	}
	return 0
}

func (x *HistoricalPerformanceInClubsItem) GetP() int64 {
	if x != nil {
		return x.P
	}
	return 0
}

func (x *HistoricalPerformanceInClubsItem) GetGoals() int64 {
	if x != nil {
		return x.Goals
	}
	return 0
}

func (x *HistoricalPerformanceInClubsItem) GetAssists() int64 {
	if x != nil {
		return x.Assists
	}
	return 0
}

func (x *HistoricalPerformanceInClubsItem) GetYellowCard() int64 {
	if x != nil {
		return x.YellowCard
	}
	return 0
}

func (x *HistoricalPerformanceInClubsItem) GetRedCard() int64 {
	if x != nil {
		return x.RedCard
	}
	return 0
}

var File_v2_foot_player_overview_proto protoreflect.FileDescriptor

var file_v2_foot_player_overview_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x76, 0x32, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfa, 0x01, 0x0a, 0x18, 0x46, 0x6f, 0x6f, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x41, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xeb, 0x01, 0x0a,
	0x2f, 0x46, 0x6f, 0x6f, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x41, 0x6e, 0x64, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x6e, 0x54, 0x68, 0x65, 0x50, 0x69, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x24, 0x0a, 0x0d, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x72, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x72, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x57, 0x65, 0x61, 0x6b, 0x6e, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x57, 0x65, 0x61, 0x6b, 0x6e, 0x65,
	0x73, 0x73, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x57, 0x69, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x57, 0x69, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x97, 0x01, 0x0a, 0x2e, 0x46,
	0x6f, 0x6f, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x6e, 0x43, 0x6c, 0x75, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x65, 0x0a,
	0x1c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x43, 0x6c, 0x75, 0x62, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c,
	0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x43, 0x6c, 0x75,
	0x62, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x1c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63,
	0x61, 0x6c, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x43,
	0x6c, 0x75, 0x62, 0x73, 0x22, 0xd4, 0x01, 0x0a, 0x20, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69,
	0x63, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e,
	0x43, 0x6c, 0x75, 0x62, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x65, 0x61,
	0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x54, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x50, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01,
	0x50, 0x12, 0x14, 0x0a, 0x05, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x73, 0x73, 0x69, 0x73,
	0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x59, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x61, 0x72, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x59, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x61, 0x72,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x64, 0x43, 0x61, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x52, 0x65, 0x64, 0x43, 0x61, 0x72, 0x64, 0x42, 0x07, 0x5a, 0x05, 0x2e,
	0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_foot_player_overview_proto_rawDescOnce sync.Once
	file_v2_foot_player_overview_proto_rawDescData = file_v2_foot_player_overview_proto_rawDesc
)

func file_v2_foot_player_overview_proto_rawDescGZIP() []byte {
	file_v2_foot_player_overview_proto_rawDescOnce.Do(func() {
		file_v2_foot_player_overview_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_foot_player_overview_proto_rawDescData)
	})
	return file_v2_foot_player_overview_proto_rawDescData
}

var file_v2_foot_player_overview_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v2_foot_player_overview_proto_goTypes = []interface{}{
	(*FootPlayerInfoBasicsResp)(nil),                        // 0: FootPlayerInfoBasicsResp
	(*FootPlayerInfoFeaturesAndPositionOnThePitchResp)(nil), // 1: FootPlayerInfoFeaturesAndPositionOnThePitchResp
	(*FootPlayerInfoHistoricalPerformanceInClubsResp)(nil),  // 2: FootPlayerInfoHistoricalPerformanceInClubsResp
	(*HistoricalPerformanceInClubsItem)(nil),                // 3: HistoricalPerformanceInClubsItem
}
var file_v2_foot_player_overview_proto_depIdxs = []int32{
	3, // 0: FootPlayerInfoHistoricalPerformanceInClubsResp.HistoricalPerformanceInClubs:type_name -> HistoricalPerformanceInClubsItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v2_foot_player_overview_proto_init() }
func file_v2_foot_player_overview_proto_init() {
	if File_v2_foot_player_overview_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_foot_player_overview_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootPlayerInfoBasicsResp); i {
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
		file_v2_foot_player_overview_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootPlayerInfoFeaturesAndPositionOnThePitchResp); i {
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
		file_v2_foot_player_overview_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootPlayerInfoHistoricalPerformanceInClubsResp); i {
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
		file_v2_foot_player_overview_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoricalPerformanceInClubsItem); i {
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
			RawDescriptor: file_v2_foot_player_overview_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_foot_player_overview_proto_goTypes,
		DependencyIndexes: file_v2_foot_player_overview_proto_depIdxs,
		MessageInfos:      file_v2_foot_player_overview_proto_msgTypes,
	}.Build()
	File_v2_foot_player_overview_proto = out.File
	file_v2_foot_player_overview_proto_rawDesc = nil
	file_v2_foot_player_overview_proto_goTypes = nil
	file_v2_foot_player_overview_proto_depIdxs = nil
}
