// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: foot_lineup_match_data.proto

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
type FootLineupMatchDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchId  int64  `protobuf:"varint,1,opt,name=MatchId,proto3" json:"MatchId,omitempty"`  // 比赛id
	Language string `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"` //请求语言
}

func (x *FootLineupMatchDataRequest) Reset() {
	*x = FootLineupMatchDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_lineup_match_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootLineupMatchDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootLineupMatchDataRequest) ProtoMessage() {}

func (x *FootLineupMatchDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foot_lineup_match_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootLineupMatchDataRequest.ProtoReflect.Descriptor instead.
func (*FootLineupMatchDataRequest) Descriptor() ([]byte, []int) {
	return file_foot_lineup_match_data_proto_rawDescGZIP(), []int{0}
}

func (x *FootLineupMatchDataRequest) GetMatchId() int64 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

func (x *FootLineupMatchDataRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type FootLineupMatchDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*DataMatchList `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *FootLineupMatchDataResponse) Reset() {
	*x = FootLineupMatchDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_lineup_match_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootLineupMatchDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootLineupMatchDataResponse) ProtoMessage() {}

func (x *FootLineupMatchDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_foot_lineup_match_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootLineupMatchDataResponse.ProtoReflect.Descriptor instead.
func (*FootLineupMatchDataResponse) Descriptor() ([]byte, []int) {
	return file_foot_lineup_match_data_proto_rawDescGZIP(), []int{1}
}

func (x *FootLineupMatchDataResponse) GetData() []*DataMatchList {
	if x != nil {
		return x.Data
	}
	return nil
}

type DataMatchList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchId        int64  `protobuf:"varint,1,opt,name=matchId,proto3" json:"matchId,omitempty"`                //比赛id
	HomeTeamId     int64  `protobuf:"varint,2,opt,name=homeTeamId,proto3" json:"homeTeamId,omitempty"`          //主队id
	HomeTeamLogo   string `protobuf:"bytes,18,opt,name=homeTeamLogo,proto3" json:"homeTeamLogo,omitempty"`      // 主队图片
	HomeName       string `protobuf:"bytes,3,opt,name=homeName,proto3" json:"homeName,omitempty"`               //主队名称
	HomeScore      string `protobuf:"bytes,4,opt,name=homeScore,proto3" json:"homeScore,omitempty"`             //主队比分
	AwayTeamId     int64  `protobuf:"varint,5,opt,name=awayTeamId,proto3" json:"awayTeamId,omitempty"`          // 客队id
	AwayTeamLogo   string `protobuf:"bytes,19,opt,name=awayTeamLogo,proto3" json:"awayTeamLogo,omitempty"`      // 客队图片
	AwayName       string `protobuf:"bytes,6,opt,name=awayName,proto3" json:"awayName,omitempty"`               // 客队名称
	AwayScore      string `protobuf:"bytes,7,opt,name=awayScore,proto3" json:"awayScore,omitempty"`             // 客队得分
	MatchTime      int64  `protobuf:"varint,8,opt,name=matchTime,proto3" json:"matchTime,omitempty"`            // 比赛时间
	TournamentId   int64  `protobuf:"varint,9,opt,name=tournamentId,proto3" json:"tournamentId,omitempty"`      // 联赛id
	TournamentName string `protobuf:"bytes,10,opt,name=tournamentName,proto3" json:"tournamentName,omitempty"`  // 联赛名称
	IsTournament   bool   `protobuf:"varint,11,opt,name=isTournament,proto3" json:"isTournament,omitempty"`     // 是否本赛事
	IsHomeTeam     bool   `protobuf:"varint,12,opt,name=isHomeTeam,proto3" json:"isHomeTeam,omitempty"`         // 是否同主
	IsAwayTeam     bool   `protobuf:"varint,13,opt,name=isAwayTeam,proto3" json:"isAwayTeam,omitempty"`         // 是否同客
	IsIdenticalLet bool   `protobuf:"varint,14,opt,name=isIdenticalLet,proto3" json:"isIdenticalLet,omitempty"` // 是否相同让球
	HomeHalfScore  string `protobuf:"bytes,15,opt,name=homeHalfScore,proto3" json:"homeHalfScore,omitempty"`    //主队半场比分
	AwayHalfScore  string `protobuf:"bytes,16,opt,name=awayHalfScore,proto3" json:"awayHalfScore,omitempty"`    //客队半场比分
	Winner         int64  `protobuf:"varint,17,opt,name=winner,proto3" json:"winner,omitempty"`                 // 0未知 1主 2客3平
}

func (x *DataMatchList) Reset() {
	*x = DataMatchList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_lineup_match_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataMatchList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataMatchList) ProtoMessage() {}

func (x *DataMatchList) ProtoReflect() protoreflect.Message {
	mi := &file_foot_lineup_match_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataMatchList.ProtoReflect.Descriptor instead.
func (*DataMatchList) Descriptor() ([]byte, []int) {
	return file_foot_lineup_match_data_proto_rawDescGZIP(), []int{2}
}

func (x *DataMatchList) GetMatchId() int64 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

func (x *DataMatchList) GetHomeTeamId() int64 {
	if x != nil {
		return x.HomeTeamId
	}
	return 0
}

func (x *DataMatchList) GetHomeTeamLogo() string {
	if x != nil {
		return x.HomeTeamLogo
	}
	return ""
}

func (x *DataMatchList) GetHomeName() string {
	if x != nil {
		return x.HomeName
	}
	return ""
}

func (x *DataMatchList) GetHomeScore() string {
	if x != nil {
		return x.HomeScore
	}
	return ""
}

func (x *DataMatchList) GetAwayTeamId() int64 {
	if x != nil {
		return x.AwayTeamId
	}
	return 0
}

func (x *DataMatchList) GetAwayTeamLogo() string {
	if x != nil {
		return x.AwayTeamLogo
	}
	return ""
}

func (x *DataMatchList) GetAwayName() string {
	if x != nil {
		return x.AwayName
	}
	return ""
}

func (x *DataMatchList) GetAwayScore() string {
	if x != nil {
		return x.AwayScore
	}
	return ""
}

func (x *DataMatchList) GetMatchTime() int64 {
	if x != nil {
		return x.MatchTime
	}
	return 0
}

func (x *DataMatchList) GetTournamentId() int64 {
	if x != nil {
		return x.TournamentId
	}
	return 0
}

func (x *DataMatchList) GetTournamentName() string {
	if x != nil {
		return x.TournamentName
	}
	return ""
}

func (x *DataMatchList) GetIsTournament() bool {
	if x != nil {
		return x.IsTournament
	}
	return false
}

func (x *DataMatchList) GetIsHomeTeam() bool {
	if x != nil {
		return x.IsHomeTeam
	}
	return false
}

func (x *DataMatchList) GetIsAwayTeam() bool {
	if x != nil {
		return x.IsAwayTeam
	}
	return false
}

func (x *DataMatchList) GetIsIdenticalLet() bool {
	if x != nil {
		return x.IsIdenticalLet
	}
	return false
}

func (x *DataMatchList) GetHomeHalfScore() string {
	if x != nil {
		return x.HomeHalfScore
	}
	return ""
}

func (x *DataMatchList) GetAwayHalfScore() string {
	if x != nil {
		return x.AwayHalfScore
	}
	return ""
}

func (x *DataMatchList) GetWinner() int64 {
	if x != nil {
		return x.Winner
	}
	return 0
}

var File_foot_lineup_match_data_proto protoreflect.FileDescriptor

var file_foot_lineup_match_data_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x5f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52,
	0x0a, 0x1a, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x22, 0x41, 0x0a, 0x1b, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x22, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0xff, 0x04, 0x0a, 0x0d, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x6f, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67,
	0x6f, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x6f, 0x6d, 0x65, 0x54, 0x65, 0x61,
	0x6d, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x61, 0x77, 0x61, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x77, 0x61, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x0c, 0x61, 0x77, 0x61, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x6f, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x77, 0x61, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x4c,
	0x6f, 0x67, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x77, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x77, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x77, 0x61, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x77, 0x61, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x26, 0x0a, 0x0e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x54, 0x6f, 0x75,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69,
	0x73, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69,
	0x73, 0x48, 0x6f, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x69, 0x73, 0x48, 0x6f, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x69,
	0x73, 0x41, 0x77, 0x61, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x69, 0x73, 0x41, 0x77, 0x61, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x26, 0x0a, 0x0e, 0x69,
	0x73, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x4c, 0x65, 0x74, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x6c,
	0x4c, 0x65, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x6f, 0x6d, 0x65, 0x48, 0x61, 0x6c, 0x66, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x68, 0x6f, 0x6d, 0x65,
	0x48, 0x61, 0x6c, 0x66, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x77, 0x61,
	0x79, 0x48, 0x61, 0x6c, 0x66, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x61, 0x77, 0x61, 0x79, 0x48, 0x61, 0x6c, 0x66, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foot_lineup_match_data_proto_rawDescOnce sync.Once
	file_foot_lineup_match_data_proto_rawDescData = file_foot_lineup_match_data_proto_rawDesc
)

func file_foot_lineup_match_data_proto_rawDescGZIP() []byte {
	file_foot_lineup_match_data_proto_rawDescOnce.Do(func() {
		file_foot_lineup_match_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_foot_lineup_match_data_proto_rawDescData)
	})
	return file_foot_lineup_match_data_proto_rawDescData
}

var file_foot_lineup_match_data_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_foot_lineup_match_data_proto_goTypes = []interface{}{
	(*FootLineupMatchDataRequest)(nil),  // 0: FootLineupMatchDataRequest
	(*FootLineupMatchDataResponse)(nil), // 1: FootLineupMatchDataResponse
	(*DataMatchList)(nil),               // 2: DataMatchList
}
var file_foot_lineup_match_data_proto_depIdxs = []int32{
	2, // 0: FootLineupMatchDataResponse.Data:type_name -> DataMatchList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_foot_lineup_match_data_proto_init() }
func file_foot_lineup_match_data_proto_init() {
	if File_foot_lineup_match_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foot_lineup_match_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootLineupMatchDataRequest); i {
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
		file_foot_lineup_match_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootLineupMatchDataResponse); i {
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
		file_foot_lineup_match_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataMatchList); i {
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
			RawDescriptor: file_foot_lineup_match_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foot_lineup_match_data_proto_goTypes,
		DependencyIndexes: file_foot_lineup_match_data_proto_depIdxs,
		MessageInfos:      file_foot_lineup_match_data_proto_msgTypes,
	}.Build()
	File_foot_lineup_match_data_proto = out.File
	file_foot_lineup_match_data_proto_rawDesc = nil
	file_foot_lineup_match_data_proto_goTypes = nil
	file_foot_lineup_match_data_proto_depIdxs = nil
}
