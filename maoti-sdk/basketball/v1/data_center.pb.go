// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: data_center.proto

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

type DataCenterCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment []*Comment `protobuf:"bytes,1,rep,name=Comment,proto3" json:"Comment,omitempty"`
}

func (x *DataCenterCommentRequest) Reset() {
	*x = DataCenterCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_center_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataCenterCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataCenterCommentRequest) ProtoMessage() {}

func (x *DataCenterCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_center_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataCenterCommentRequest.ProtoReflect.Descriptor instead.
func (*DataCenterCommentRequest) Descriptor() ([]byte, []int) {
	return file_data_center_proto_rawDescGZIP(), []int{0}
}

func (x *DataCenterCommentRequest) GetComment() []*Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchId     int64  `protobuf:"varint,1,opt,name=MatchId,proto3" json:"MatchId,omitempty"`
	Team        int64  `protobuf:"varint,2,opt,name=Team,proto3" json:"Team,omitempty"`
	ElapsedPlus int64  `protobuf:"varint,3,opt,name=ElapsedPlus,proto3" json:"ElapsedPlus,omitempty"`
	Elapsed     string `protobuf:"bytes,4,opt,name=Elapsed,proto3" json:"Elapsed,omitempty"`
	Score       string `protobuf:"bytes,5,opt,name=Score,proto3" json:"Score,omitempty"`
	Sort        int64  `protobuf:"varint,6,opt,name=Sort,proto3" json:"Sort,omitempty"`
	Scope       int64  `protobuf:"varint,7,opt,name=Scope,proto3" json:"Scope,omitempty"`
	Content     string `protobuf:"bytes,8,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_center_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_data_center_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_data_center_proto_rawDescGZIP(), []int{1}
}

func (x *Comment) GetMatchId() int64 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

func (x *Comment) GetTeam() int64 {
	if x != nil {
		return x.Team
	}
	return 0
}

func (x *Comment) GetElapsedPlus() int64 {
	if x != nil {
		return x.ElapsedPlus
	}
	return 0
}

func (x *Comment) GetElapsed() string {
	if x != nil {
		return x.Elapsed
	}
	return ""
}

func (x *Comment) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

func (x *Comment) GetSort() int64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *Comment) GetScope() int64 {
	if x != nil {
		return x.Scope
	}
	return 0
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type DataCenterLiveScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LiveScore []*LiveScore `protobuf:"bytes,1,rep,name=LiveScore,proto3" json:"LiveScore,omitempty"`
}

func (x *DataCenterLiveScoreRequest) Reset() {
	*x = DataCenterLiveScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_center_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataCenterLiveScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataCenterLiveScoreRequest) ProtoMessage() {}

func (x *DataCenterLiveScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_center_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataCenterLiveScoreRequest.ProtoReflect.Descriptor instead.
func (*DataCenterLiveScoreRequest) Descriptor() ([]byte, []int) {
	return file_data_center_proto_rawDescGZIP(), []int{2}
}

func (x *DataCenterLiveScoreRequest) GetLiveScore() []*LiveScore {
	if x != nil {
		return x.LiveScore
	}
	return nil
}

type LiveScore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SportId          int64            `protobuf:"varint,1,opt,name=SportId,proto3" json:"SportId,omitempty"`
	MatchId          int64            `protobuf:"varint,2,opt,name=MatchId,proto3" json:"MatchId,omitempty"`
	HomeId           int64            `protobuf:"varint,3,opt,name=HomeId,proto3" json:"HomeId,omitempty"`
	HomeName         string           `protobuf:"bytes,4,opt,name=HomeName,proto3" json:"HomeName,omitempty"`
	AwayId           int64            `protobuf:"varint,5,opt,name=AwayId,proto3" json:"AwayId,omitempty"`
	AwayName         string           `protobuf:"bytes,6,opt,name=AwayName,proto3" json:"AwayName,omitempty"`
	MatchTime        int64            `protobuf:"varint,7,opt,name=MatchTime,proto3" json:"MatchTime,omitempty"`
	StartTime        int64            `protobuf:"varint,8,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	HomeScore        string           `protobuf:"bytes,9,opt,name=HomeScore,proto3" json:"HomeScore,omitempty"`
	AwayScore        string           `protobuf:"bytes,10,opt,name=AwayScore,proto3" json:"AwayScore,omitempty"`
	HomeHalfScore    string           `protobuf:"bytes,11,opt,name=HomeHalfScore,proto3" json:"HomeHalfScore,omitempty"`
	AwayHalfScore    string           `protobuf:"bytes,12,opt,name=AwayHalfScore,proto3" json:"AwayHalfScore,omitempty"`
	MatchStatus      int64            `protobuf:"varint,13,opt,name=MatchStatus,proto3" json:"MatchStatus,omitempty"`
	MatchChildStatus int64            `protobuf:"varint,14,opt,name=MatchChildStatus,proto3" json:"MatchChildStatus,omitempty"`
	TimePlayed       int64            `protobuf:"varint,15,opt,name=TimePlayed,proto3" json:"TimePlayed,omitempty"`
	TimeRemaining    int64            `protobuf:"varint,16,opt,name=TimeRemaining,proto3" json:"TimeRemaining,omitempty"`
	TimeRunning      int64            `protobuf:"varint,17,opt,name=TimeRunning,proto3" json:"TimeRunning,omitempty"`
	TimeUpdate       int64            `protobuf:"varint,18,opt,name=TimeUpdate,proto3" json:"TimeUpdate,omitempty"`
	MatchIncident    []*MatchIncident `protobuf:"bytes,19,rep,name=MatchIncident,proto3" json:"MatchIncident,omitempty"`
	MatchTeamCount   *MatchTeamCount  `protobuf:"bytes,20,opt,name=MatchTeamCount,proto3" json:"MatchTeamCount,omitempty"`
}

func (x *LiveScore) Reset() {
	*x = LiveScore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_center_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveScore) ProtoMessage() {}

func (x *LiveScore) ProtoReflect() protoreflect.Message {
	mi := &file_data_center_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveScore.ProtoReflect.Descriptor instead.
func (*LiveScore) Descriptor() ([]byte, []int) {
	return file_data_center_proto_rawDescGZIP(), []int{3}
}

func (x *LiveScore) GetSportId() int64 {
	if x != nil {
		return x.SportId
	}
	return 0
}

func (x *LiveScore) GetMatchId() int64 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

func (x *LiveScore) GetHomeId() int64 {
	if x != nil {
		return x.HomeId
	}
	return 0
}

func (x *LiveScore) GetHomeName() string {
	if x != nil {
		return x.HomeName
	}
	return ""
}

func (x *LiveScore) GetAwayId() int64 {
	if x != nil {
		return x.AwayId
	}
	return 0
}

func (x *LiveScore) GetAwayName() string {
	if x != nil {
		return x.AwayName
	}
	return ""
}

func (x *LiveScore) GetMatchTime() int64 {
	if x != nil {
		return x.MatchTime
	}
	return 0
}

func (x *LiveScore) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *LiveScore) GetHomeScore() string {
	if x != nil {
		return x.HomeScore
	}
	return ""
}

func (x *LiveScore) GetAwayScore() string {
	if x != nil {
		return x.AwayScore
	}
	return ""
}

func (x *LiveScore) GetHomeHalfScore() string {
	if x != nil {
		return x.HomeHalfScore
	}
	return ""
}

func (x *LiveScore) GetAwayHalfScore() string {
	if x != nil {
		return x.AwayHalfScore
	}
	return ""
}

func (x *LiveScore) GetMatchStatus() int64 {
	if x != nil {
		return x.MatchStatus
	}
	return 0
}

func (x *LiveScore) GetMatchChildStatus() int64 {
	if x != nil {
		return x.MatchChildStatus
	}
	return 0
}

func (x *LiveScore) GetTimePlayed() int64 {
	if x != nil {
		return x.TimePlayed
	}
	return 0
}

func (x *LiveScore) GetTimeRemaining() int64 {
	if x != nil {
		return x.TimeRemaining
	}
	return 0
}

func (x *LiveScore) GetTimeRunning() int64 {
	if x != nil {
		return x.TimeRunning
	}
	return 0
}

func (x *LiveScore) GetTimeUpdate() int64 {
	if x != nil {
		return x.TimeUpdate
	}
	return 0
}

func (x *LiveScore) GetMatchIncident() []*MatchIncident {
	if x != nil {
		return x.MatchIncident
	}
	return nil
}

func (x *LiveScore) GetMatchTeamCount() *MatchTeamCount {
	if x != nil {
		return x.MatchTeamCount
	}
	return nil
}

type MatchIncident struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchId          int64  `protobuf:"varint,1,opt,name=MatchId,proto3" json:"MatchId,omitempty"`
	TeamId           int64  `protobuf:"varint,2,opt,name=TeamId,proto3" json:"TeamId,omitempty"`
	TeamName         string `protobuf:"bytes,3,opt,name=TeamName,proto3" json:"TeamName,omitempty"`
	IncidentTypeId   int64  `protobuf:"varint,4,opt,name=IncidentTypeId,proto3" json:"IncidentTypeId,omitempty"`
	IncidentType     string `protobuf:"bytes,5,opt,name=IncidentType,proto3" json:"IncidentType,omitempty"`
	Index            int64  `protobuf:"varint,6,opt,name=Index,proto3" json:"Index,omitempty"`
	Elapsed          int64  `protobuf:"varint,7,opt,name=Elapsed,proto3" json:"Elapsed,omitempty"`
	ElapsedPlus      int64  `protobuf:"varint,8,opt,name=ElapsedPlus,proto3" json:"ElapsedPlus,omitempty"`
	PlayerId         int64  `protobuf:"varint,9,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
	PlayerName       string `protobuf:"bytes,10,opt,name=PlayerName,proto3" json:"PlayerName,omitempty"`
	RelatePlayerId   int64  `protobuf:"varint,11,opt,name=RelatePlayerId,proto3" json:"RelatePlayerId,omitempty"`
	RelatePlayerName string `protobuf:"bytes,12,opt,name=RelatePlayerName,proto3" json:"RelatePlayerName,omitempty"`
	X                int64  `protobuf:"varint,13,opt,name=X,proto3" json:"X,omitempty"`
	Y                int64  `protobuf:"varint,14,opt,name=Y,proto3" json:"Y,omitempty"`
	GoalType         int64  `protobuf:"varint,15,opt,name=GoalType,proto3" json:"GoalType,omitempty"`
	Scores           string `protobuf:"bytes,16,opt,name=Scores,proto3" json:"Scores,omitempty"`
	IsHome           bool   `protobuf:"varint,17,opt,name=IsHome,proto3" json:"IsHome,omitempty"`
}

func (x *MatchIncident) Reset() {
	*x = MatchIncident{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_center_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchIncident) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchIncident) ProtoMessage() {}

func (x *MatchIncident) ProtoReflect() protoreflect.Message {
	mi := &file_data_center_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchIncident.ProtoReflect.Descriptor instead.
func (*MatchIncident) Descriptor() ([]byte, []int) {
	return file_data_center_proto_rawDescGZIP(), []int{4}
}

func (x *MatchIncident) GetMatchId() int64 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

func (x *MatchIncident) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *MatchIncident) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

func (x *MatchIncident) GetIncidentTypeId() int64 {
	if x != nil {
		return x.IncidentTypeId
	}
	return 0
}

func (x *MatchIncident) GetIncidentType() string {
	if x != nil {
		return x.IncidentType
	}
	return ""
}

func (x *MatchIncident) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *MatchIncident) GetElapsed() int64 {
	if x != nil {
		return x.Elapsed
	}
	return 0
}

func (x *MatchIncident) GetElapsedPlus() int64 {
	if x != nil {
		return x.ElapsedPlus
	}
	return 0
}

func (x *MatchIncident) GetPlayerId() int64 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *MatchIncident) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *MatchIncident) GetRelatePlayerId() int64 {
	if x != nil {
		return x.RelatePlayerId
	}
	return 0
}

func (x *MatchIncident) GetRelatePlayerName() string {
	if x != nil {
		return x.RelatePlayerName
	}
	return ""
}

func (x *MatchIncident) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *MatchIncident) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *MatchIncident) GetGoalType() int64 {
	if x != nil {
		return x.GoalType
	}
	return 0
}

func (x *MatchIncident) GetScores() string {
	if x != nil {
		return x.Scores
	}
	return ""
}

func (x *MatchIncident) GetIsHome() bool {
	if x != nil {
		return x.IsHome
	}
	return false
}

//比赛球队统计
type MatchTeamCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchId  int64          `protobuf:"varint,1,opt,name=MatchId,proto3" json:"MatchId,omitempty"`
	HomeId   int64          `protobuf:"varint,2,opt,name=HomeId,proto3" json:"HomeId,omitempty"`
	HomeName string         `protobuf:"bytes,3,opt,name=HomeName,proto3" json:"HomeName,omitempty"`
	AwayId   int64          `protobuf:"varint,4,opt,name=AwayId,proto3" json:"AwayId,omitempty"`
	AwayName string         `protobuf:"bytes,5,opt,name=AwayName,proto3" json:"AwayName,omitempty"`
	Stats    []*CountParams `protobuf:"bytes,6,rep,name=Stats,proto3" json:"Stats,omitempty"`
}

func (x *MatchTeamCount) Reset() {
	*x = MatchTeamCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_center_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchTeamCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchTeamCount) ProtoMessage() {}

func (x *MatchTeamCount) ProtoReflect() protoreflect.Message {
	mi := &file_data_center_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchTeamCount.ProtoReflect.Descriptor instead.
func (*MatchTeamCount) Descriptor() ([]byte, []int) {
	return file_data_center_proto_rawDescGZIP(), []int{5}
}

func (x *MatchTeamCount) GetMatchId() int64 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

func (x *MatchTeamCount) GetHomeId() int64 {
	if x != nil {
		return x.HomeId
	}
	return 0
}

func (x *MatchTeamCount) GetHomeName() string {
	if x != nil {
		return x.HomeName
	}
	return ""
}

func (x *MatchTeamCount) GetAwayId() int64 {
	if x != nil {
		return x.AwayId
	}
	return 0
}

func (x *MatchTeamCount) GetAwayName() string {
	if x != nil {
		return x.AwayName
	}
	return ""
}

func (x *MatchTeamCount) GetStats() []*CountParams {
	if x != nil {
		return x.Stats
	}
	return nil
}

type CountParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeId int64  `protobuf:"varint,1,opt,name=TypeId,proto3" json:"TypeId,omitempty"`
	Type   string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	Home   string `protobuf:"bytes,3,opt,name=Home,proto3" json:"Home,omitempty"`
	Away   string `protobuf:"bytes,4,opt,name=Away,proto3" json:"Away,omitempty"`
}

func (x *CountParams) Reset() {
	*x = CountParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_center_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountParams) ProtoMessage() {}

func (x *CountParams) ProtoReflect() protoreflect.Message {
	mi := &file_data_center_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountParams.ProtoReflect.Descriptor instead.
func (*CountParams) Descriptor() ([]byte, []int) {
	return file_data_center_proto_rawDescGZIP(), []int{6}
}

func (x *CountParams) GetTypeId() int64 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *CountParams) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CountParams) GetHome() string {
	if x != nil {
		return x.Home
	}
	return ""
}

func (x *CountParams) GetAway() string {
	if x != nil {
		return x.Away
	}
	return ""
}

var File_data_center_proto protoreflect.FileDescriptor

var file_data_center_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x18, 0x44, 0x61, 0x74, 0x61, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0xcd, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x61,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x20, 0x0a,
	0x0b, 0x45, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x50, 0x6c, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x45, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x50, 0x6c, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x45, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x45, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x53,
	0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x46, 0x0a, 0x1a, 0x44, 0x61, 0x74, 0x61, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x09, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x52, 0x09, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x22, 0xb0, 0x05, 0x0a, 0x09,
	0x4c, 0x69, 0x76, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x70, 0x6f,
	0x72, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x70, 0x6f, 0x72,
	0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x48, 0x6f, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x48,
	0x6f, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x6f, 0x6d, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x48, 0x6f, 0x6d, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x77, 0x61, 0x79, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x41, 0x77, 0x61, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x77, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x77, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x41, 0x77, 0x61, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x41, 0x77, 0x61, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x48, 0x6f, 0x6d, 0x65, 0x48, 0x61, 0x6c, 0x66, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x48, 0x6f, 0x6d, 0x65, 0x48, 0x61, 0x6c, 0x66, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x41, 0x77, 0x61, 0x79, 0x48, 0x61, 0x6c, 0x66, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x41, 0x77, 0x61, 0x79,
	0x48, 0x61, 0x6c, 0x66, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x68, 0x69, 0x6c,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x54, 0x69, 0x6d,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x0a,
	0x0b, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12,
	0x1e, 0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x34, 0x0a, 0x0d, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x63,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x0e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x65,
	0x61, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0e,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xf3,
	0x03, 0x0a, 0x0d, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x65,
	0x61, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x54, 0x65, 0x61, 0x6d,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26,
	0x0a, 0x0e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x49, 0x6e,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x18, 0x0a, 0x07, 0x45, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x45, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x6c,
	0x61, 0x70, 0x73, 0x65, 0x64, 0x50, 0x6c, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x45, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x50, 0x6c, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x2a, 0x0a, 0x10, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0c, 0x0a, 0x01,
	0x58, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x59, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x6f, 0x61, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x47, 0x6f, 0x61, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x49, 0x73, 0x48, 0x6f, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x49, 0x73,
	0x48, 0x6f, 0x6d, 0x65, 0x22, 0xb6, 0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x65,
	0x61, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x6f, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x48, 0x6f, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x6f, 0x6d,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x48, 0x6f, 0x6d,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x77, 0x61, 0x79, 0x49, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x41, 0x77, 0x61, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x41, 0x77, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x41, 0x77, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x22, 0x61, 0x0a,
	0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x6f, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x41, 0x77, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x77, 0x61, 0x79,
	0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_data_center_proto_rawDescOnce sync.Once
	file_data_center_proto_rawDescData = file_data_center_proto_rawDesc
)

func file_data_center_proto_rawDescGZIP() []byte {
	file_data_center_proto_rawDescOnce.Do(func() {
		file_data_center_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_center_proto_rawDescData)
	})
	return file_data_center_proto_rawDescData
}

var file_data_center_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_data_center_proto_goTypes = []interface{}{
	(*DataCenterCommentRequest)(nil),   // 0: DataCenterCommentRequest
	(*Comment)(nil),                    // 1: Comment
	(*DataCenterLiveScoreRequest)(nil), // 2: DataCenterLiveScoreRequest
	(*LiveScore)(nil),                  // 3: LiveScore
	(*MatchIncident)(nil),              // 4: MatchIncident
	(*MatchTeamCount)(nil),             // 5: MatchTeamCount
	(*CountParams)(nil),                // 6: CountParams
}
var file_data_center_proto_depIdxs = []int32{
	1, // 0: DataCenterCommentRequest.Comment:type_name -> Comment
	3, // 1: DataCenterLiveScoreRequest.LiveScore:type_name -> LiveScore
	4, // 2: LiveScore.MatchIncident:type_name -> MatchIncident
	5, // 3: LiveScore.MatchTeamCount:type_name -> MatchTeamCount
	6, // 4: MatchTeamCount.Stats:type_name -> CountParams
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_data_center_proto_init() }
func file_data_center_proto_init() {
	if File_data_center_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_center_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataCenterCommentRequest); i {
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
		file_data_center_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_data_center_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataCenterLiveScoreRequest); i {
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
		file_data_center_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveScore); i {
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
		file_data_center_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchIncident); i {
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
		file_data_center_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchTeamCount); i {
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
		file_data_center_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountParams); i {
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
			RawDescriptor: file_data_center_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_center_proto_goTypes,
		DependencyIndexes: file_data_center_proto_depIdxs,
		MessageInfos:      file_data_center_proto_msgTypes,
	}.Build()
	File_data_center_proto = out.File
	file_data_center_proto_rawDesc = nil
	file_data_center_proto_goTypes = nil
	file_data_center_proto_depIdxs = nil
}
