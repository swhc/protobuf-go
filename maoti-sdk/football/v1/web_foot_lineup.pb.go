// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: web_foot_lineup.proto

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

//阵容
type WebFootMatchLineupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`  //请求语言
	MatchId  int64  `protobuf:"varint,6,opt,name=matchId,proto3" json:"matchId,omitempty"`   //比赛ID
	TimeZone int64  `protobuf:"varint,1,opt,name=TimeZone,proto3" json:"TimeZone,omitempty"` //时区
}

func (x *WebFootMatchLineupRequest) Reset() {
	*x = WebFootMatchLineupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_lineup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootMatchLineupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootMatchLineupRequest) ProtoMessage() {}

func (x *WebFootMatchLineupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_lineup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootMatchLineupRequest.ProtoReflect.Descriptor instead.
func (*WebFootMatchLineupRequest) Descriptor() ([]byte, []int) {
	return file_web_foot_lineup_proto_rawDescGZIP(), []int{0}
}

func (x *WebFootMatchLineupRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *WebFootMatchLineupRequest) GetMatchId() int64 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

func (x *WebFootMatchLineupRequest) GetTimeZone() int64 {
	if x != nil {
		return x.TimeZone
	}
	return 0
}

type WebFootMatchLineupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64               `protobuf:"varint,4,opt,name=Id,proto3" json:"Id,omitempty"`               //比赛id
	MatchTime    int64               `protobuf:"varint,6,opt,name=matchTime,proto3" json:"matchTime,omitempty"` //比赛时间
	StartingMain *WebFootMatchLineup `protobuf:"bytes,3,opt,name=starting_main,json=startingMain,proto3" json:"starting_main,omitempty"`
	StartingCust *WebFootMatchLineup `protobuf:"bytes,13,opt,name=starting_cust,json=startingCust,proto3" json:"starting_cust,omitempty"`
	ForecastMain *WebFootMatchLineup `protobuf:"bytes,14,opt,name=forecast_main,json=forecastMain,proto3" json:"forecast_main,omitempty"`
	ForecastCust *WebFootMatchLineup `protobuf:"bytes,15,opt,name=forecast_cust,json=forecastCust,proto3" json:"forecast_cust,omitempty"`
}

func (x *WebFootMatchLineupResponse) Reset() {
	*x = WebFootMatchLineupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_lineup_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootMatchLineupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootMatchLineupResponse) ProtoMessage() {}

func (x *WebFootMatchLineupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_lineup_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootMatchLineupResponse.ProtoReflect.Descriptor instead.
func (*WebFootMatchLineupResponse) Descriptor() ([]byte, []int) {
	return file_web_foot_lineup_proto_rawDescGZIP(), []int{1}
}

func (x *WebFootMatchLineupResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WebFootMatchLineupResponse) GetMatchTime() int64 {
	if x != nil {
		return x.MatchTime
	}
	return 0
}

func (x *WebFootMatchLineupResponse) GetStartingMain() *WebFootMatchLineup {
	if x != nil {
		return x.StartingMain
	}
	return nil
}

func (x *WebFootMatchLineupResponse) GetStartingCust() *WebFootMatchLineup {
	if x != nil {
		return x.StartingCust
	}
	return nil
}

func (x *WebFootMatchLineupResponse) GetForecastMain() *WebFootMatchLineup {
	if x != nil {
		return x.ForecastMain
	}
	return nil
}

func (x *WebFootMatchLineupResponse) GetForecastCust() *WebFootMatchLineup {
	if x != nil {
		return x.ForecastCust
	}
	return nil
}

type WebFootMatchLineup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId         int64                      `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`                           //球队id
	TeamImage      string                     `protobuf:"bytes,2,opt,name=team_image,json=teamImage,proto3" json:"team_image,omitempty"`                   //球队队徽
	TeamName       string                     `protobuf:"bytes,3,opt,name=team_name,json=teamName,proto3" json:"team_name,omitempty"`                      //球队名称
	Formation      string                     `protobuf:"bytes,4,opt,name=formation,proto3" json:"formation,omitempty"`                                    //阵容
	WinRate        int64                      `protobuf:"varint,5,opt,name=win_rate,json=winRate,proto3" json:"win_rate,omitempty"`                        //胜率
	BigBallRate    int64                      `protobuf:"varint,6,opt,name=big_ball_rate,json=bigBallRate,proto3" json:"big_ball_rate,omitempty"`          //大球率
	CornerKickRate int64                      `protobuf:"varint,7,opt,name=corner_kick_rate,json=cornerKickRate,proto3" json:"corner_kick_rate,omitempty"` //大角球率
	Score          string                     `protobuf:"bytes,8,opt,name=score,proto3" json:"score,omitempty"`                                            //比分 8.5
	Member         []*WebFootLiveLineupMember `protobuf:"bytes,9,rep,name=member,proto3" json:"member,omitempty"`                                          //球队成员详情
}

func (x *WebFootMatchLineup) Reset() {
	*x = WebFootMatchLineup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_lineup_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootMatchLineup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootMatchLineup) ProtoMessage() {}

func (x *WebFootMatchLineup) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_lineup_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootMatchLineup.ProtoReflect.Descriptor instead.
func (*WebFootMatchLineup) Descriptor() ([]byte, []int) {
	return file_web_foot_lineup_proto_rawDescGZIP(), []int{2}
}

func (x *WebFootMatchLineup) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *WebFootMatchLineup) GetTeamImage() string {
	if x != nil {
		return x.TeamImage
	}
	return ""
}

func (x *WebFootMatchLineup) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

func (x *WebFootMatchLineup) GetFormation() string {
	if x != nil {
		return x.Formation
	}
	return ""
}

func (x *WebFootMatchLineup) GetWinRate() int64 {
	if x != nil {
		return x.WinRate
	}
	return 0
}

func (x *WebFootMatchLineup) GetBigBallRate() int64 {
	if x != nil {
		return x.BigBallRate
	}
	return 0
}

func (x *WebFootMatchLineup) GetCornerKickRate() int64 {
	if x != nil {
		return x.CornerKickRate
	}
	return 0
}

func (x *WebFootMatchLineup) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

func (x *WebFootMatchLineup) GetMember() []*WebFootLiveLineupMember {
	if x != nil {
		return x.Member
	}
	return nil
}

//球队成员详情
type WebFootLiveLineupMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonId    int64  `protobuf:"varint,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`          //球员id
	Image       string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`                                 //球员图片
	PersonName  string `protobuf:"bytes,3,opt,name=person_name,json=personName,proto3" json:"person_name,omitempty"`     //球员名称
	ElevateType int64  `protobuf:"varint,4,opt,name=elevate_type,json=elevateType,proto3" json:"elevate_type,omitempty"` //提升类型 （1：在场时胜率提升，2：在场时总进球提升，3：在场时角球数提升）
	ElevateRate int64  `protobuf:"varint,5,opt,name=elevate_rate,json=elevateRate,proto3" json:"elevate_rate,omitempty"` //提升百分比
	ClothesNum  int64  `protobuf:"varint,6,opt,name=clothes_num,json=clothesNum,proto3" json:"clothes_num,omitempty"`    //球衣号码
	Position    int64  `protobuf:"varint,7,opt,name=position,proto3" json:"position,omitempty"`                          //x轴
}

func (x *WebFootLiveLineupMember) Reset() {
	*x = WebFootLiveLineupMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_lineup_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootLiveLineupMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootLiveLineupMember) ProtoMessage() {}

func (x *WebFootLiveLineupMember) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_lineup_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootLiveLineupMember.ProtoReflect.Descriptor instead.
func (*WebFootLiveLineupMember) Descriptor() ([]byte, []int) {
	return file_web_foot_lineup_proto_rawDescGZIP(), []int{3}
}

func (x *WebFootLiveLineupMember) GetPersonId() int64 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

func (x *WebFootLiveLineupMember) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *WebFootLiveLineupMember) GetPersonName() string {
	if x != nil {
		return x.PersonName
	}
	return ""
}

func (x *WebFootLiveLineupMember) GetElevateType() int64 {
	if x != nil {
		return x.ElevateType
	}
	return 0
}

func (x *WebFootLiveLineupMember) GetElevateRate() int64 {
	if x != nil {
		return x.ElevateRate
	}
	return 0
}

func (x *WebFootLiveLineupMember) GetClothesNum() int64 {
	if x != nil {
		return x.ClothesNum
	}
	return 0
}

func (x *WebFootLiveLineupMember) GetPosition() int64 {
	if x != nil {
		return x.Position
	}
	return 0
}

var File_web_foot_lineup_proto protoreflect.FileDescriptor

var file_web_foot_lineup_proto_rawDesc = []byte{
	0x0a, 0x15, 0x77, 0x65, 0x62, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x75,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x19, 0x57, 0x65, 0x62, 0x46, 0x6f,
	0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x69,
	0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x54, 0x69,
	0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0xb2, 0x02, 0x0a, 0x1a, 0x57, 0x65, 0x62, 0x46, 0x6f,
	0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x57, 0x65, 0x62,
	0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x52,
	0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x69, 0x6e, 0x12, 0x38, 0x0a,
	0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x43, 0x75, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0d, 0x66, 0x6f, 0x72, 0x65, 0x63,
	0x61, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x6e,
	0x65, 0x75, 0x70, 0x52, 0x0c, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x61, 0x69,
	0x6e, 0x12, 0x38, 0x0a, 0x0d, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x75,
	0x73, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x57, 0x65, 0x62, 0x46, 0x6f,
	0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x52, 0x0c, 0x66,
	0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x43, 0x75, 0x73, 0x74, 0x22, 0xb8, 0x02, 0x0a, 0x12,
	0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x6e, 0x65,
	0x75, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x65, 0x61, 0x6d, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65,
	0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x77, 0x69, 0x6e, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x77, 0x69, 0x6e, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x22, 0x0a, 0x0d, 0x62, 0x69, 0x67, 0x5f, 0x62, 0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x62, 0x69, 0x67, 0x42, 0x61, 0x6c, 0x6c,
	0x52, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6f, 0x72, 0x6e, 0x65, 0x72, 0x5f, 0x6b,
	0x69, 0x63, 0x6b, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x63, 0x6f, 0x72, 0x6e, 0x65, 0x72, 0x4b, 0x69, 0x63, 0x6b, 0x52, 0x61, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69,
	0x76, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xf0, 0x01, 0x0a, 0x17, 0x57, 0x65, 0x62, 0x46, 0x6f,
	0x6f, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x65, 0x6c,
	0x65, 0x76, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6c, 0x65,
	0x76, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x6c, 0x6f, 0x74, 0x68, 0x65, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x63, 0x6c, 0x6f, 0x74, 0x68, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web_foot_lineup_proto_rawDescOnce sync.Once
	file_web_foot_lineup_proto_rawDescData = file_web_foot_lineup_proto_rawDesc
)

func file_web_foot_lineup_proto_rawDescGZIP() []byte {
	file_web_foot_lineup_proto_rawDescOnce.Do(func() {
		file_web_foot_lineup_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_foot_lineup_proto_rawDescData)
	})
	return file_web_foot_lineup_proto_rawDescData
}

var file_web_foot_lineup_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_web_foot_lineup_proto_goTypes = []interface{}{
	(*WebFootMatchLineupRequest)(nil),  // 0: WebFootMatchLineupRequest
	(*WebFootMatchLineupResponse)(nil), // 1: WebFootMatchLineupResponse
	(*WebFootMatchLineup)(nil),         // 2: WebFootMatchLineup
	(*WebFootLiveLineupMember)(nil),    // 3: WebFootLiveLineupMember
}
var file_web_foot_lineup_proto_depIdxs = []int32{
	2, // 0: WebFootMatchLineupResponse.starting_main:type_name -> WebFootMatchLineup
	2, // 1: WebFootMatchLineupResponse.starting_cust:type_name -> WebFootMatchLineup
	2, // 2: WebFootMatchLineupResponse.forecast_main:type_name -> WebFootMatchLineup
	2, // 3: WebFootMatchLineupResponse.forecast_cust:type_name -> WebFootMatchLineup
	3, // 4: WebFootMatchLineup.member:type_name -> WebFootLiveLineupMember
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_web_foot_lineup_proto_init() }
func file_web_foot_lineup_proto_init() {
	if File_web_foot_lineup_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_foot_lineup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootMatchLineupRequest); i {
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
		file_web_foot_lineup_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootMatchLineupResponse); i {
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
		file_web_foot_lineup_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootMatchLineup); i {
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
		file_web_foot_lineup_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootLiveLineupMember); i {
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
			RawDescriptor: file_web_foot_lineup_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_foot_lineup_proto_goTypes,
		DependencyIndexes: file_web_foot_lineup_proto_depIdxs,
		MessageInfos:      file_web_foot_lineup_proto_msgTypes,
	}.Build()
	File_web_foot_lineup_proto = out.File
	file_web_foot_lineup_proto_rawDesc = nil
	file_web_foot_lineup_proto_goTypes = nil
	file_web_foot_lineup_proto_depIdxs = nil
}
