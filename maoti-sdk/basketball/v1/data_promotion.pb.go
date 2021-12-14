// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: data_promotion.proto

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

//球员榜
type DataPromotionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language string `protobuf:"bytes,1,opt,name=Language,proto3" json:"Language,omitempty"`  //请求语言  1:zh  2:en
	SeasonId int64  `protobuf:"varint,6,opt,name=SeasonId,proto3" json:"SeasonId,omitempty"` //年份对应id
}

func (x *DataPromotionRequest) Reset() {
	*x = DataPromotionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_promotion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataPromotionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPromotionRequest) ProtoMessage() {}

func (x *DataPromotionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_promotion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPromotionRequest.ProtoReflect.Descriptor instead.
func (*DataPromotionRequest) Descriptor() ([]byte, []int) {
	return file_data_promotion_proto_rawDescGZIP(), []int{0}
}

func (x *DataPromotionRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *DataPromotionRequest) GetSeasonId() int64 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

type DataPromotionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EighthFinals  *EighthFinals  `protobuf:"bytes,1,opt,name=EighthFinals,proto3" json:"EighthFinals,omitempty"`
	QuarterFinals *QuarterFinals `protobuf:"bytes,2,opt,name=QuarterFinals,proto3" json:"QuarterFinals,omitempty"`
	HalfFinals    *HalfFinals    `protobuf:"bytes,3,opt,name=HalfFinals,proto3" json:"HalfFinals,omitempty"`
	Finals        *Finals        `protobuf:"bytes,4,opt,name=Finals,proto3" json:"Finals,omitempty"`
}

func (x *DataPromotionResponse) Reset() {
	*x = DataPromotionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_promotion_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataPromotionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPromotionResponse) ProtoMessage() {}

func (x *DataPromotionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_data_promotion_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPromotionResponse.ProtoReflect.Descriptor instead.
func (*DataPromotionResponse) Descriptor() ([]byte, []int) {
	return file_data_promotion_proto_rawDescGZIP(), []int{1}
}

func (x *DataPromotionResponse) GetEighthFinals() *EighthFinals {
	if x != nil {
		return x.EighthFinals
	}
	return nil
}

func (x *DataPromotionResponse) GetQuarterFinals() *QuarterFinals {
	if x != nil {
		return x.QuarterFinals
	}
	return nil
}

func (x *DataPromotionResponse) GetHalfFinals() *HalfFinals {
	if x != nil {
		return x.HalfFinals
	}
	return nil
}

func (x *DataPromotionResponse) GetFinals() *Finals {
	if x != nil {
		return x.Finals
	}
	return nil
}

type EighthFinals struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopRows    []*TeamRows `protobuf:"bytes,1,rep,name=top_rows,json=topRows,proto3" json:"top_rows,omitempty"`
	BottomRows []*TeamRows `protobuf:"bytes,2,rep,name=bottom_rows,json=bottomRows,proto3" json:"bottom_rows,omitempty"`
	IsBye      bool        `protobuf:"varint,3,opt,name=is_bye,json=isBye,proto3" json:"is_bye,omitempty"`
}

func (x *EighthFinals) Reset() {
	*x = EighthFinals{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_promotion_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EighthFinals) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EighthFinals) ProtoMessage() {}

func (x *EighthFinals) ProtoReflect() protoreflect.Message {
	mi := &file_data_promotion_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EighthFinals.ProtoReflect.Descriptor instead.
func (*EighthFinals) Descriptor() ([]byte, []int) {
	return file_data_promotion_proto_rawDescGZIP(), []int{2}
}

func (x *EighthFinals) GetTopRows() []*TeamRows {
	if x != nil {
		return x.TopRows
	}
	return nil
}

func (x *EighthFinals) GetBottomRows() []*TeamRows {
	if x != nil {
		return x.BottomRows
	}
	return nil
}

func (x *EighthFinals) GetIsBye() bool {
	if x != nil {
		return x.IsBye
	}
	return false
}

type QuarterFinals struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopRows    []*TeamRows `protobuf:"bytes,1,rep,name=top_rows,json=topRows,proto3" json:"top_rows,omitempty"`
	BottomRows []*TeamRows `protobuf:"bytes,2,rep,name=bottom_rows,json=bottomRows,proto3" json:"bottom_rows,omitempty"`
}

func (x *QuarterFinals) Reset() {
	*x = QuarterFinals{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_promotion_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuarterFinals) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuarterFinals) ProtoMessage() {}

func (x *QuarterFinals) ProtoReflect() protoreflect.Message {
	mi := &file_data_promotion_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuarterFinals.ProtoReflect.Descriptor instead.
func (*QuarterFinals) Descriptor() ([]byte, []int) {
	return file_data_promotion_proto_rawDescGZIP(), []int{3}
}

func (x *QuarterFinals) GetTopRows() []*TeamRows {
	if x != nil {
		return x.TopRows
	}
	return nil
}

func (x *QuarterFinals) GetBottomRows() []*TeamRows {
	if x != nil {
		return x.BottomRows
	}
	return nil
}

type HalfFinals struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopRows    []*TeamRows `protobuf:"bytes,1,rep,name=top_rows,json=topRows,proto3" json:"top_rows,omitempty"`
	BottomRows []*TeamRows `protobuf:"bytes,2,rep,name=bottom_rows,json=bottomRows,proto3" json:"bottom_rows,omitempty"`
}

func (x *HalfFinals) Reset() {
	*x = HalfFinals{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_promotion_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HalfFinals) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HalfFinals) ProtoMessage() {}

func (x *HalfFinals) ProtoReflect() protoreflect.Message {
	mi := &file_data_promotion_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HalfFinals.ProtoReflect.Descriptor instead.
func (*HalfFinals) Descriptor() ([]byte, []int) {
	return file_data_promotion_proto_rawDescGZIP(), []int{4}
}

func (x *HalfFinals) GetTopRows() []*TeamRows {
	if x != nil {
		return x.TopRows
	}
	return nil
}

func (x *HalfFinals) GetBottomRows() []*TeamRows {
	if x != nil {
		return x.BottomRows
	}
	return nil
}

type Finals struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopRows           *TeamData `protobuf:"bytes,1,opt,name=top_rows,json=topRows,proto3" json:"top_rows,omitempty"`
	BottomRows        *TeamData `protobuf:"bytes,2,opt,name=bottom_rows,json=bottomRows,proto3" json:"bottom_rows,omitempty"`
	Rounds            int64     `protobuf:"varint,3,opt,name=rounds,proto3" json:"rounds,omitempty"`
	Score             string    `protobuf:"bytes,4,opt,name=score,proto3" json:"score,omitempty"`
	PromotionTeamId   int64     `protobuf:"varint,5,opt,name=promotion_team_id,json=promotionTeamId,proto3" json:"promotion_team_id,omitempty"`
	EventStatusResult int64     `protobuf:"varint,11,opt,name=event_status_result,json=eventStatusResult,proto3" json:"event_status_result,omitempty"`
}

func (x *Finals) Reset() {
	*x = Finals{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_promotion_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Finals) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Finals) ProtoMessage() {}

func (x *Finals) ProtoReflect() protoreflect.Message {
	mi := &file_data_promotion_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Finals.ProtoReflect.Descriptor instead.
func (*Finals) Descriptor() ([]byte, []int) {
	return file_data_promotion_proto_rawDescGZIP(), []int{5}
}

func (x *Finals) GetTopRows() *TeamData {
	if x != nil {
		return x.TopRows
	}
	return nil
}

func (x *Finals) GetBottomRows() *TeamData {
	if x != nil {
		return x.BottomRows
	}
	return nil
}

func (x *Finals) GetRounds() int64 {
	if x != nil {
		return x.Rounds
	}
	return 0
}

func (x *Finals) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

func (x *Finals) GetPromotionTeamId() int64 {
	if x != nil {
		return x.PromotionTeamId
	}
	return 0
}

func (x *Finals) GetEventStatusResult() int64 {
	if x != nil {
		return x.EventStatusResult
	}
	return 0
}

type TeamRows struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FrontTeam         *TeamData `protobuf:"bytes,1,opt,name=front_team,json=frontTeam,proto3" json:"front_team,omitempty"`
	LateTeam          *TeamData `protobuf:"bytes,2,opt,name=late_team,json=lateTeam,proto3" json:"late_team,omitempty"`
	Rounds            int64     `protobuf:"varint,3,opt,name=rounds,proto3" json:"rounds,omitempty"`
	Score             string    `protobuf:"bytes,4,opt,name=score,proto3" json:"score,omitempty"`
	PromotionTeamId   int64     `protobuf:"varint,5,opt,name=promotion_team_id,json=promotionTeamId,proto3" json:"promotion_team_id,omitempty"`
	EventStatusResult int64     `protobuf:"varint,11,opt,name=event_status_result,json=eventStatusResult,proto3" json:"event_status_result,omitempty"`
}

func (x *TeamRows) Reset() {
	*x = TeamRows{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_promotion_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamRows) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamRows) ProtoMessage() {}

func (x *TeamRows) ProtoReflect() protoreflect.Message {
	mi := &file_data_promotion_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamRows.ProtoReflect.Descriptor instead.
func (*TeamRows) Descriptor() ([]byte, []int) {
	return file_data_promotion_proto_rawDescGZIP(), []int{6}
}

func (x *TeamRows) GetFrontTeam() *TeamData {
	if x != nil {
		return x.FrontTeam
	}
	return nil
}

func (x *TeamRows) GetLateTeam() *TeamData {
	if x != nil {
		return x.LateTeam
	}
	return nil
}

func (x *TeamRows) GetRounds() int64 {
	if x != nil {
		return x.Rounds
	}
	return 0
}

func (x *TeamRows) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

func (x *TeamRows) GetPromotionTeamId() int64 {
	if x != nil {
		return x.PromotionTeamId
	}
	return 0
}

func (x *TeamRows) GetEventStatusResult() int64 {
	if x != nil {
		return x.EventStatusResult
	}
	return 0
}

type TeamData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegularSeasonRank int64  `protobuf:"varint,1,opt,name=regular_season_rank,json=regularSeasonRank,proto3" json:"regular_season_rank,omitempty"`
	TeamId            int64  `protobuf:"varint,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	TeamName          string `protobuf:"bytes,3,opt,name=team_name,json=teamName,proto3" json:"team_name,omitempty"`
	TeamImage         string `protobuf:"bytes,4,opt,name=team_image,json=teamImage,proto3" json:"team_image,omitempty"`
}

func (x *TeamData) Reset() {
	*x = TeamData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_promotion_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamData) ProtoMessage() {}

func (x *TeamData) ProtoReflect() protoreflect.Message {
	mi := &file_data_promotion_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamData.ProtoReflect.Descriptor instead.
func (*TeamData) Descriptor() ([]byte, []int) {
	return file_data_promotion_proto_rawDescGZIP(), []int{7}
}

func (x *TeamData) GetRegularSeasonRank() int64 {
	if x != nil {
		return x.RegularSeasonRank
	}
	return 0
}

func (x *TeamData) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *TeamData) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

func (x *TeamData) GetTeamImage() string {
	if x != nil {
		return x.TeamImage
	}
	return ""
}

var File_data_promotion_proto protoreflect.FileDescriptor

var file_data_promotion_proto_rawDesc = []byte{
	0x0a, 0x14, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x14, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xce, 0x01, 0x0a, 0x15, 0x44, 0x61, 0x74, 0x61, 0x50,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x0c, 0x45, 0x69, 0x67, 0x68, 0x74, 0x68, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x45, 0x69, 0x67, 0x68, 0x74, 0x68, 0x46,
	0x69, 0x6e, 0x61, 0x6c, 0x73, 0x52, 0x0c, 0x45, 0x69, 0x67, 0x68, 0x74, 0x68, 0x46, 0x69, 0x6e,
	0x61, 0x6c, 0x73, 0x12, 0x34, 0x0a, 0x0d, 0x51, 0x75, 0x61, 0x72, 0x74, 0x65, 0x72, 0x46, 0x69,
	0x6e, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x51, 0x75, 0x61,
	0x72, 0x74, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x73, 0x52, 0x0d, 0x51, 0x75, 0x61, 0x72,
	0x74, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x73, 0x12, 0x2b, 0x0a, 0x0a, 0x48, 0x61, 0x6c,
	0x66, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x48, 0x61, 0x6c, 0x66, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x73, 0x52, 0x0a, 0x48, 0x61, 0x6c, 0x66,
	0x46, 0x69, 0x6e, 0x61, 0x6c, 0x73, 0x12, 0x1f, 0x0a, 0x06, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x73, 0x52,
	0x06, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x73, 0x22, 0x77, 0x0a, 0x0c, 0x45, 0x69, 0x67, 0x68, 0x74,
	0x68, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x73, 0x12, 0x24, 0x0a, 0x08, 0x74, 0x6f, 0x70, 0x5f, 0x72,
	0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x54, 0x65, 0x61, 0x6d,
	0x52, 0x6f, 0x77, 0x73, 0x52, 0x07, 0x74, 0x6f, 0x70, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x2a, 0x0a,
	0x0b, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x6f, 0x77, 0x73, 0x52, 0x0a, 0x62,
	0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f,
	0x62, 0x79, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x42, 0x79, 0x65,
	0x22, 0x61, 0x0a, 0x0d, 0x51, 0x75, 0x61, 0x72, 0x74, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x61, 0x6c,
	0x73, 0x12, 0x24, 0x0a, 0x08, 0x74, 0x6f, 0x70, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x6f, 0x77, 0x73, 0x52, 0x07,
	0x74, 0x6f, 0x70, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x2a, 0x0a, 0x0b, 0x62, 0x6f, 0x74, 0x74, 0x6f,
	0x6d, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x54,
	0x65, 0x61, 0x6d, 0x52, 0x6f, 0x77, 0x73, 0x52, 0x0a, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x52,
	0x6f, 0x77, 0x73, 0x22, 0x5e, 0x0a, 0x0a, 0x48, 0x61, 0x6c, 0x66, 0x46, 0x69, 0x6e, 0x61, 0x6c,
	0x73, 0x12, 0x24, 0x0a, 0x08, 0x74, 0x6f, 0x70, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x6f, 0x77, 0x73, 0x52, 0x07,
	0x74, 0x6f, 0x70, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x2a, 0x0a, 0x0b, 0x62, 0x6f, 0x74, 0x74, 0x6f,
	0x6d, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x54,
	0x65, 0x61, 0x6d, 0x52, 0x6f, 0x77, 0x73, 0x52, 0x0a, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x52,
	0x6f, 0x77, 0x73, 0x22, 0xe4, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x73, 0x12, 0x24,
	0x0a, 0x08, 0x74, 0x6f, 0x70, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x74, 0x6f, 0x70,
	0x52, 0x6f, 0x77, 0x73, 0x12, 0x2a, 0x0a, 0x0b, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x5f, 0x72,
	0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x54, 0x65, 0x61, 0x6d,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x52, 0x6f, 0x77, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x2a,
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x61, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xe6, 0x01, 0x0a, 0x08, 0x54,
	0x65, 0x61, 0x6d, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x28, 0x0a, 0x0a, 0x66, 0x72, 0x6f, 0x6e, 0x74,
	0x5f, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x54, 0x65,
	0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x54, 0x65, 0x61,
	0x6d, 0x12, 0x26, 0x0a, 0x09, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x08, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x61,
	0x6d, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x8f, 0x01, 0x0a, 0x08, 0x54, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x73, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x72,
	0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x6b,
	0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x61,
	0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65,
	0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x6d,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_promotion_proto_rawDescOnce sync.Once
	file_data_promotion_proto_rawDescData = file_data_promotion_proto_rawDesc
)

func file_data_promotion_proto_rawDescGZIP() []byte {
	file_data_promotion_proto_rawDescOnce.Do(func() {
		file_data_promotion_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_promotion_proto_rawDescData)
	})
	return file_data_promotion_proto_rawDescData
}

var file_data_promotion_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_data_promotion_proto_goTypes = []interface{}{
	(*DataPromotionRequest)(nil),  // 0: DataPromotionRequest
	(*DataPromotionResponse)(nil), // 1: DataPromotionResponse
	(*EighthFinals)(nil),          // 2: EighthFinals
	(*QuarterFinals)(nil),         // 3: QuarterFinals
	(*HalfFinals)(nil),            // 4: HalfFinals
	(*Finals)(nil),                // 5: Finals
	(*TeamRows)(nil),              // 6: TeamRows
	(*TeamData)(nil),              // 7: TeamData
}
var file_data_promotion_proto_depIdxs = []int32{
	2,  // 0: DataPromotionResponse.EighthFinals:type_name -> EighthFinals
	3,  // 1: DataPromotionResponse.QuarterFinals:type_name -> QuarterFinals
	4,  // 2: DataPromotionResponse.HalfFinals:type_name -> HalfFinals
	5,  // 3: DataPromotionResponse.Finals:type_name -> Finals
	6,  // 4: EighthFinals.top_rows:type_name -> TeamRows
	6,  // 5: EighthFinals.bottom_rows:type_name -> TeamRows
	6,  // 6: QuarterFinals.top_rows:type_name -> TeamRows
	6,  // 7: QuarterFinals.bottom_rows:type_name -> TeamRows
	6,  // 8: HalfFinals.top_rows:type_name -> TeamRows
	6,  // 9: HalfFinals.bottom_rows:type_name -> TeamRows
	7,  // 10: Finals.top_rows:type_name -> TeamData
	7,  // 11: Finals.bottom_rows:type_name -> TeamData
	7,  // 12: TeamRows.front_team:type_name -> TeamData
	7,  // 13: TeamRows.late_team:type_name -> TeamData
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_data_promotion_proto_init() }
func file_data_promotion_proto_init() {
	if File_data_promotion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_promotion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataPromotionRequest); i {
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
		file_data_promotion_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataPromotionResponse); i {
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
		file_data_promotion_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EighthFinals); i {
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
		file_data_promotion_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuarterFinals); i {
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
		file_data_promotion_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HalfFinals); i {
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
		file_data_promotion_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Finals); i {
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
		file_data_promotion_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamRows); i {
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
		file_data_promotion_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamData); i {
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
			RawDescriptor: file_data_promotion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_promotion_proto_goTypes,
		DependencyIndexes: file_data_promotion_proto_depIdxs,
		MessageInfos:      file_data_promotion_proto_msgTypes,
	}.Build()
	File_data_promotion_proto = out.File
	file_data_promotion_proto_rawDesc = nil
	file_data_promotion_proto_goTypes = nil
	file_data_promotion_proto_depIdxs = nil
}
