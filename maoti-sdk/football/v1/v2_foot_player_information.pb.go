// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: v2_foot_player_information.proto

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

// 详细信息  （接口）
type FootPlayerInfoPersonalDataResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DateOfBirth       string `protobuf:"bytes,1,opt,name=DateOfBirth,proto3" json:"DateOfBirth,omitempty"`             //出生年月日
	CountryId         int64  `protobuf:"varint,2,opt,name=CountryId,proto3" json:"CountryId,omitempty"`                //球员国家id
	Positions         string `protobuf:"bytes,3,opt,name=positions,proto3" json:"positions,omitempty"`                 // 位置
	PreferredFoot     string `protobuf:"bytes,4,opt,name=PreferredFoot,proto3" json:"PreferredFoot,omitempty"`         // 惯用脚
	CurrentTeamId     int64  `protobuf:"varint,5,opt,name=CurrentTeamId,proto3" json:"CurrentTeamId,omitempty"`        // 现在球队id
	TournamentId      int64  `protobuf:"varint,6,opt,name=TournamentId,proto3" json:"TournamentId,omitempty"`          // 所属联赛id
	PreviousTeamId    int64  `protobuf:"varint,7,opt,name=PreviousTeamId,proto3" json:"PreviousTeamId,omitempty"`      // 上个球队id
	ClothesNum        int64  `protobuf:"varint,8,opt,name=ClothesNum,proto3" json:"ClothesNum,omitempty"`              //现在球衣号
	HistoryClothesNum string `protobuf:"bytes,9,opt,name=HistoryClothesNum,proto3" json:"HistoryClothesNum,omitempty"` // 历史球衣号
	CareerYears       int64  `protobuf:"varint,10,opt,name=CareerYears,proto3" json:"CareerYears,omitempty"`           // 职业生涯年份
}

func (x *FootPlayerInfoPersonalDataResp) Reset() {
	*x = FootPlayerInfoPersonalDataResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_player_information_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootPlayerInfoPersonalDataResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootPlayerInfoPersonalDataResp) ProtoMessage() {}

func (x *FootPlayerInfoPersonalDataResp) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_player_information_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootPlayerInfoPersonalDataResp.ProtoReflect.Descriptor instead.
func (*FootPlayerInfoPersonalDataResp) Descriptor() ([]byte, []int) {
	return file_v2_foot_player_information_proto_rawDescGZIP(), []int{0}
}

func (x *FootPlayerInfoPersonalDataResp) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *FootPlayerInfoPersonalDataResp) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *FootPlayerInfoPersonalDataResp) GetPositions() string {
	if x != nil {
		return x.Positions
	}
	return ""
}

func (x *FootPlayerInfoPersonalDataResp) GetPreferredFoot() string {
	if x != nil {
		return x.PreferredFoot
	}
	return ""
}

func (x *FootPlayerInfoPersonalDataResp) GetCurrentTeamId() int64 {
	if x != nil {
		return x.CurrentTeamId
	}
	return 0
}

func (x *FootPlayerInfoPersonalDataResp) GetTournamentId() int64 {
	if x != nil {
		return x.TournamentId
	}
	return 0
}

func (x *FootPlayerInfoPersonalDataResp) GetPreviousTeamId() int64 {
	if x != nil {
		return x.PreviousTeamId
	}
	return 0
}

func (x *FootPlayerInfoPersonalDataResp) GetClothesNum() int64 {
	if x != nil {
		return x.ClothesNum
	}
	return 0
}

func (x *FootPlayerInfoPersonalDataResp) GetHistoryClothesNum() string {
	if x != nil {
		return x.HistoryClothesNum
	}
	return ""
}

func (x *FootPlayerInfoPersonalDataResp) GetCareerYears() int64 {
	if x != nil {
		return x.CareerYears
	}
	return 0
}

// 转会记录和冠军数据 （接口）
type FootPlayerInfoHistoricalTransfersAndChampionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HistoricalTransfers []*HistoricalTransfersItem `protobuf:"bytes,1,rep,name=HistoricalTransfers,proto3" json:"HistoricalTransfers,omitempty"`
	Champion            []*ChampionItem            `protobuf:"bytes,2,rep,name=Champion,proto3" json:"Champion,omitempty"`
}

func (x *FootPlayerInfoHistoricalTransfersAndChampionResp) Reset() {
	*x = FootPlayerInfoHistoricalTransfersAndChampionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_player_information_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootPlayerInfoHistoricalTransfersAndChampionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootPlayerInfoHistoricalTransfersAndChampionResp) ProtoMessage() {}

func (x *FootPlayerInfoHistoricalTransfersAndChampionResp) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_player_information_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootPlayerInfoHistoricalTransfersAndChampionResp.ProtoReflect.Descriptor instead.
func (*FootPlayerInfoHistoricalTransfersAndChampionResp) Descriptor() ([]byte, []int) {
	return file_v2_foot_player_information_proto_rawDescGZIP(), []int{1}
}

func (x *FootPlayerInfoHistoricalTransfersAndChampionResp) GetHistoricalTransfers() []*HistoricalTransfersItem {
	if x != nil {
		return x.HistoricalTransfers
	}
	return nil
}

func (x *FootPlayerInfoHistoricalTransfersAndChampionResp) GetChampion() []*ChampionItem {
	if x != nil {
		return x.Champion
	}
	return nil
}

// 转会记录子项
type HistoricalTransfersItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date          string `protobuf:"bytes,1,opt,name=Date,proto3" json:"Date,omitempty"`                    // 日期
	OriginID      int64  `protobuf:"varint,2,opt,name=OriginID,proto3" json:"OriginID,omitempty"`           //转出球队ID
	DestinationId int64  `protobuf:"varint,3,opt,name=DestinationId,proto3" json:"DestinationId,omitempty"` //转入球队ID
	Cost          string `protobuf:"bytes,4,opt,name=Cost,proto3" json:"Cost,omitempty"`                    //转会价格
}

func (x *HistoricalTransfersItem) Reset() {
	*x = HistoricalTransfersItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_player_information_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoricalTransfersItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoricalTransfersItem) ProtoMessage() {}

func (x *HistoricalTransfersItem) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_player_information_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoricalTransfersItem.ProtoReflect.Descriptor instead.
func (*HistoricalTransfersItem) Descriptor() ([]byte, []int) {
	return file_v2_foot_player_information_proto_rawDescGZIP(), []int{2}
}

func (x *HistoricalTransfersItem) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *HistoricalTransfersItem) GetOriginID() int64 {
	if x != nil {
		return x.OriginID
	}
	return 0
}

func (x *HistoricalTransfersItem) GetDestinationId() int64 {
	if x != nil {
		return x.DestinationId
	}
	return 0
}

func (x *HistoricalTransfersItem) GetCost() string {
	if x != nil {
		return x.Cost
	}
	return ""
}

// 冠军数据子项
type ChampionItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TournamentId        int64    `protobuf:"varint,1,opt,name=TournamentId,proto3" json:"TournamentId,omitempty"`              // 联赛id
	ChampionSeasonDates []string `protobuf:"bytes,2,rep,name=ChampionSeasonDates,proto3" json:"ChampionSeasonDates,omitempty"` // 冠军赛季日期集
	ChampionCount       int64    `protobuf:"varint,3,opt,name=ChampionCount,proto3" json:"ChampionCount,omitempty"`            // 次数
}

func (x *ChampionItem) Reset() {
	*x = ChampionItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_player_information_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChampionItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChampionItem) ProtoMessage() {}

func (x *ChampionItem) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_player_information_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChampionItem.ProtoReflect.Descriptor instead.
func (*ChampionItem) Descriptor() ([]byte, []int) {
	return file_v2_foot_player_information_proto_rawDescGZIP(), []int{3}
}

func (x *ChampionItem) GetTournamentId() int64 {
	if x != nil {
		return x.TournamentId
	}
	return 0
}

func (x *ChampionItem) GetChampionSeasonDates() []string {
	if x != nil {
		return x.ChampionSeasonDates
	}
	return nil
}

func (x *ChampionItem) GetChampionCount() int64 {
	if x != nil {
		return x.ChampionCount
	}
	return 0
}

var File_v2_foot_player_information_proto protoreflect.FileDescriptor

var file_v2_foot_player_information_proto_rawDesc = []byte{
	0x0a, 0x20, 0x76, 0x32, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x86, 0x03, 0x0a, 0x1e, 0x46, 0x6f, 0x6f, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42,
	0x69, 0x72, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x61, 0x74, 0x65,
	0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64,
	0x46, 0x6f, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x50, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x72, 0x65, 0x64, 0x46, 0x6f, 0x6f, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x0c, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x54,
	0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x50, 0x72, 0x65,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43,
	0x6c, 0x6f, 0x74, 0x68, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x43, 0x6c, 0x6f, 0x74, 0x68, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x2c, 0x0a, 0x11, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6c, 0x6f, 0x74, 0x68, 0x65, 0x73, 0x4e, 0x75, 0x6d,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x43,
	0x6c, 0x6f, 0x74, 0x68, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x61, 0x72,
	0x65, 0x65, 0x72, 0x59, 0x65, 0x61, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x43, 0x61, 0x72, 0x65, 0x65, 0x72, 0x59, 0x65, 0x61, 0x72, 0x73, 0x22, 0xa9, 0x01, 0x0a, 0x30,
	0x46, 0x6f, 0x6f, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x73, 0x41, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x4a, 0x0a, 0x13, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x13, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69,
	0x63, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x12, 0x29, 0x0a, 0x08,
	0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x43,
	0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x22, 0x83, 0x01, 0x0a, 0x17, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x73, 0x74, 0x22, 0x8a, 0x01,
	0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x22,
	0x0a, 0x0c, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x13, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x43, 0x68, 0x61,
	0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_foot_player_information_proto_rawDescOnce sync.Once
	file_v2_foot_player_information_proto_rawDescData = file_v2_foot_player_information_proto_rawDesc
)

func file_v2_foot_player_information_proto_rawDescGZIP() []byte {
	file_v2_foot_player_information_proto_rawDescOnce.Do(func() {
		file_v2_foot_player_information_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_foot_player_information_proto_rawDescData)
	})
	return file_v2_foot_player_information_proto_rawDescData
}

var file_v2_foot_player_information_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v2_foot_player_information_proto_goTypes = []interface{}{
	(*FootPlayerInfoPersonalDataResp)(nil),                   // 0: FootPlayerInfoPersonalDataResp
	(*FootPlayerInfoHistoricalTransfersAndChampionResp)(nil), // 1: FootPlayerInfoHistoricalTransfersAndChampionResp
	(*HistoricalTransfersItem)(nil),                          // 2: HistoricalTransfersItem
	(*ChampionItem)(nil),                                     // 3: ChampionItem
}
var file_v2_foot_player_information_proto_depIdxs = []int32{
	2, // 0: FootPlayerInfoHistoricalTransfersAndChampionResp.HistoricalTransfers:type_name -> HistoricalTransfersItem
	3, // 1: FootPlayerInfoHistoricalTransfersAndChampionResp.Champion:type_name -> ChampionItem
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v2_foot_player_information_proto_init() }
func file_v2_foot_player_information_proto_init() {
	if File_v2_foot_player_information_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_foot_player_information_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootPlayerInfoPersonalDataResp); i {
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
		file_v2_foot_player_information_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootPlayerInfoHistoricalTransfersAndChampionResp); i {
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
		file_v2_foot_player_information_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoricalTransfersItem); i {
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
		file_v2_foot_player_information_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChampionItem); i {
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
			RawDescriptor: file_v2_foot_player_information_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_foot_player_information_proto_goTypes,
		DependencyIndexes: file_v2_foot_player_information_proto_depIdxs,
		MessageInfos:      file_v2_foot_player_information_proto_msgTypes,
	}.Build()
	File_v2_foot_player_information_proto = out.File
	file_v2_foot_player_information_proto_rawDesc = nil
	file_v2_foot_player_information_proto_goTypes = nil
	file_v2_foot_player_information_proto_depIdxs = nil
}
