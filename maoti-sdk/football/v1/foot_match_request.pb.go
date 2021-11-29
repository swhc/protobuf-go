// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: foot_match_request.proto

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

//赛事对阵详情
type MatchInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"` //请求语言  1:zh  2:en
	SportId  string `protobuf:"bytes,1,opt,name=sportId,proto3" json:"sportId,omitempty"`   //1足球 2篮球
	EventId  string `protobuf:"bytes,2,opt,name=eventId,proto3" json:"eventId,omitempty"`   //比赛id
}

func (x *MatchInfoRequest) Reset() {
	*x = MatchInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchInfoRequest) ProtoMessage() {}

func (x *MatchInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchInfoRequest.ProtoReflect.Descriptor instead.
func (*MatchInfoRequest) Descriptor() ([]byte, []int) {
	return file_foot_match_request_proto_rawDescGZIP(), []int{0}
}

func (x *MatchInfoRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *MatchInfoRequest) GetSportId() string {
	if x != nil {
		return x.SportId
	}
	return ""
}

func (x *MatchInfoRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

type MatchInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64      `protobuf:"varint,3,opt,name=Id,proto3" json:"Id,omitempty"`                                                   //比赛id
	ChName            string     `protobuf:"bytes,4,opt,name=ChName,proto3" json:"ChName,omitempty"`                                            //比赛名称
	Round             string     `protobuf:"bytes,5,opt,name=Round,proto3" json:"Round,omitempty"`                                              //场次
	StartTime         int64      `protobuf:"varint,6,opt,name=StartTime,proto3" json:"StartTime,omitempty"`                                     //开始时间
	MainTeamId        int64      `protobuf:"varint,7,opt,name=MainTeamId,proto3" json:"MainTeamId,omitempty"`                                   //主队id
	MainTeamName      string     `protobuf:"bytes,8,opt,name=MainTeamName,proto3" json:"MainTeamName,omitempty"`                                //主队名称
	MainTeamImages    string     `protobuf:"bytes,9,opt,name=MainTeamImages,proto3" json:"MainTeamImages,omitempty"`                            //主队头像
	CustTeamId        int64      `protobuf:"varint,10,opt,name=CustTeamId,proto3" json:"CustTeamId,omitempty"`                                  //客队id
	CustTeamName      string     `protobuf:"bytes,11,opt,name=CustTeamName,proto3" json:"CustTeamName,omitempty"`                               //客队名称
	CustTeamImages    string     `protobuf:"bytes,12,opt,name=CustTeamImages,proto3" json:"CustTeamImages,omitempty"`                           //客队头像
	MainTeamRank      int64      `protobuf:"varint,13,opt,name=MainTeamRank,proto3" json:"MainTeamRank,omitempty"`                              //主队排名
	CustTeamRank      int64      `protobuf:"varint,14,opt,name=CustTeamRank,proto3" json:"CustTeamRank,omitempty"`                              //客队排名
	EventStatusResult int64      `protobuf:"varint,15,opt,name=EventStatusResult,proto3" json:"EventStatusResult,omitempty"`                    //状态描述：1:未开始，2:开始，3:结束,4:中断，6:取消
	IsCollect         bool       `protobuf:"varint,16,opt,name=IsCollect,proto3" json:"IsCollect,omitempty"`                                    //是否收藏
	TeamMainScore     int64      `protobuf:"varint,17,opt,name=TeamMainScore,proto3" json:"TeamMainScore,omitempty"`                            //主队得分
	TeamCustScore     int64      `protobuf:"varint,18,opt,name=TeamCustScore,proto3" json:"TeamCustScore,omitempty"`                            //客队得分
	EventStatusId     int64      `protobuf:"varint,19,opt,name=EventStatusId,proto3" json:"EventStatusId,omitempty"`                            //2：上半场，10：中场休息，3：下半场，14：等待加时，8：加时赛上半场，9：加时赛下半场，20：等待点球，4：点球大战，13：点球大战完场
	UpperTime         int64      `protobuf:"varint,20,opt,name=UpperTime,proto3" json:"UpperTime,omitempty"`                                    //上场开始时间（或加时赛）
	LowerTime         int64      `protobuf:"varint,21,opt,name=LowerTime,proto3" json:"LowerTime,omitempty"`                                    //下场开始时间（或加时赛）
	MainScore         int64      `protobuf:"varint,22,opt,name=MainScore,proto3" json:"MainScore,omitempty"`                                    //（加时赛）主队得分
	MainPointSphere   int64      `protobuf:"varint,23,opt,name=MainPoint_sphere,json=MainPointSphere,proto3" json:"MainPoint_sphere,omitempty"` //主队点球数
	CustScore         int64      `protobuf:"varint,24,opt,name=CustScore,proto3" json:"CustScore,omitempty"`                                    //（加时赛）客队得分
	CustPointSphere   int64      `protobuf:"varint,25,opt,name=CustPoint_sphere,json=CustPointSphere,proto3" json:"CustPoint_sphere,omitempty"` //客队点球数
	IsExtraTime       bool       `protobuf:"varint,31,opt,name=IsExtraTime,proto3" json:"IsExtraTime,omitempty"`                                //是否加时赛
	IsPenaltyKick     bool       `protobuf:"varint,26,opt,name=IsPenaltyKick,proto3" json:"IsPenaltyKick,omitempty"`                            //是否点球
	Elapsed           int64      `protobuf:"varint,27,opt,name=Elapsed,proto3" json:"Elapsed,omitempty"`                                        //时长
	CurrentTime       int64      `protobuf:"varint,28,opt,name=CurrentTime,proto3" json:"CurrentTime,omitempty"`                                //服务器当前时间
	Animation         *Animation `protobuf:"bytes,30,opt,name=animation,proto3" json:"animation,omitempty"`                                     //动画
}

func (x *MatchInfoResponse) Reset() {
	*x = MatchInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchInfoResponse) ProtoMessage() {}

func (x *MatchInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchInfoResponse.ProtoReflect.Descriptor instead.
func (*MatchInfoResponse) Descriptor() ([]byte, []int) {
	return file_foot_match_request_proto_rawDescGZIP(), []int{1}
}

func (x *MatchInfoResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MatchInfoResponse) GetChName() string {
	if x != nil {
		return x.ChName
	}
	return ""
}

func (x *MatchInfoResponse) GetRound() string {
	if x != nil {
		return x.Round
	}
	return ""
}

func (x *MatchInfoResponse) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *MatchInfoResponse) GetMainTeamId() int64 {
	if x != nil {
		return x.MainTeamId
	}
	return 0
}

func (x *MatchInfoResponse) GetMainTeamName() string {
	if x != nil {
		return x.MainTeamName
	}
	return ""
}

func (x *MatchInfoResponse) GetMainTeamImages() string {
	if x != nil {
		return x.MainTeamImages
	}
	return ""
}

func (x *MatchInfoResponse) GetCustTeamId() int64 {
	if x != nil {
		return x.CustTeamId
	}
	return 0
}

func (x *MatchInfoResponse) GetCustTeamName() string {
	if x != nil {
		return x.CustTeamName
	}
	return ""
}

func (x *MatchInfoResponse) GetCustTeamImages() string {
	if x != nil {
		return x.CustTeamImages
	}
	return ""
}

func (x *MatchInfoResponse) GetMainTeamRank() int64 {
	if x != nil {
		return x.MainTeamRank
	}
	return 0
}

func (x *MatchInfoResponse) GetCustTeamRank() int64 {
	if x != nil {
		return x.CustTeamRank
	}
	return 0
}

func (x *MatchInfoResponse) GetEventStatusResult() int64 {
	if x != nil {
		return x.EventStatusResult
	}
	return 0
}

func (x *MatchInfoResponse) GetIsCollect() bool {
	if x != nil {
		return x.IsCollect
	}
	return false
}

func (x *MatchInfoResponse) GetTeamMainScore() int64 {
	if x != nil {
		return x.TeamMainScore
	}
	return 0
}

func (x *MatchInfoResponse) GetTeamCustScore() int64 {
	if x != nil {
		return x.TeamCustScore
	}
	return 0
}

func (x *MatchInfoResponse) GetEventStatusId() int64 {
	if x != nil {
		return x.EventStatusId
	}
	return 0
}

func (x *MatchInfoResponse) GetUpperTime() int64 {
	if x != nil {
		return x.UpperTime
	}
	return 0
}

func (x *MatchInfoResponse) GetLowerTime() int64 {
	if x != nil {
		return x.LowerTime
	}
	return 0
}

func (x *MatchInfoResponse) GetMainScore() int64 {
	if x != nil {
		return x.MainScore
	}
	return 0
}

func (x *MatchInfoResponse) GetMainPointSphere() int64 {
	if x != nil {
		return x.MainPointSphere
	}
	return 0
}

func (x *MatchInfoResponse) GetCustScore() int64 {
	if x != nil {
		return x.CustScore
	}
	return 0
}

func (x *MatchInfoResponse) GetCustPointSphere() int64 {
	if x != nil {
		return x.CustPointSphere
	}
	return 0
}

func (x *MatchInfoResponse) GetIsExtraTime() bool {
	if x != nil {
		return x.IsExtraTime
	}
	return false
}

func (x *MatchInfoResponse) GetIsPenaltyKick() bool {
	if x != nil {
		return x.IsPenaltyKick
	}
	return false
}

func (x *MatchInfoResponse) GetElapsed() int64 {
	if x != nil {
		return x.Elapsed
	}
	return 0
}

func (x *MatchInfoResponse) GetCurrentTime() int64 {
	if x != nil {
		return x.CurrentTime
	}
	return 0
}

func (x *MatchInfoResponse) GetAnimation() *Animation {
	if x != nil {
		return x.Animation
	}
	return nil
}

type Animation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url    string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`        //url
	Width  int64  `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`   //宽度
	Height int64  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"` //高度
}

func (x *Animation) Reset() {
	*x = Animation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Animation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Animation) ProtoMessage() {}

func (x *Animation) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Animation.ProtoReflect.Descriptor instead.
func (*Animation) Descriptor() ([]byte, []int) {
	return file_foot_match_request_proto_rawDescGZIP(), []int{2}
}

func (x *Animation) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Animation) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Animation) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

var File_foot_match_request_proto protoreflect.FileDescriptor

var file_foot_match_request_proto_rawDesc = []byte{
	0x0a, 0x18, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x10, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xc9,
	0x07, 0x0a, 0x11, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x52, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x6f, 0x75,
	0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64,
	0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x4d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x4d, 0x61,
	0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x43, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c,
	0x43, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x43, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x43, 0x75, 0x73, 0x74, 0x54, 0x65,
	0x61, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x61, 0x69, 0x6e,
	0x54, 0x65, 0x61, 0x6d, 0x52, 0x61, 0x6e, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x4d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x22, 0x0a, 0x0c,
	0x43, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x61, 0x6e, 0x6b, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x61, 0x6e, 0x6b,
	0x12, 0x2c, 0x0a, 0x11, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x49, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x49, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x24, 0x0a, 0x0d,
	0x54, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x54, 0x65, 0x61, 0x6d, 0x43,
	0x75, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x55, 0x70, 0x70, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x55, 0x70, 0x70, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x4c, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x61,
	0x69, 0x6e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x4d,
	0x61, 0x69, 0x6e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x4d, 0x61, 0x69, 0x6e,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0f, 0x4d, 0x61, 0x69, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x70, 0x68,
	0x65, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x75, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x75, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x29, 0x0a, 0x10, 0x43, 0x75, 0x73, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x73,
	0x70, 0x68, 0x65, 0x72, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x43, 0x75, 0x73,
	0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x70, 0x68, 0x65, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x49, 0x73, 0x45, 0x78, 0x74, 0x72, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x49, 0x73, 0x45, 0x78, 0x74, 0x72, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24,
	0x0a, 0x0d, 0x49, 0x73, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x4b, 0x69, 0x63, 0x6b, 0x18,
	0x1a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x49, 0x73, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79,
	0x4b, 0x69, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x18,
	0x1b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x45, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x1c, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x09, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4b, 0x0a, 0x09, 0x41, 0x6e,
	0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foot_match_request_proto_rawDescOnce sync.Once
	file_foot_match_request_proto_rawDescData = file_foot_match_request_proto_rawDesc
)

func file_foot_match_request_proto_rawDescGZIP() []byte {
	file_foot_match_request_proto_rawDescOnce.Do(func() {
		file_foot_match_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_foot_match_request_proto_rawDescData)
	})
	return file_foot_match_request_proto_rawDescData
}

var file_foot_match_request_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_foot_match_request_proto_goTypes = []interface{}{
	(*MatchInfoRequest)(nil),  // 0: MatchInfoRequest
	(*MatchInfoResponse)(nil), // 1: MatchInfoResponse
	(*Animation)(nil),         // 2: Animation
}
var file_foot_match_request_proto_depIdxs = []int32{
	2, // 0: MatchInfoResponse.animation:type_name -> Animation
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_foot_match_request_proto_init() }
func file_foot_match_request_proto_init() {
	if File_foot_match_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foot_match_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchInfoRequest); i {
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
		file_foot_match_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchInfoResponse); i {
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
		file_foot_match_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Animation); i {
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
			RawDescriptor: file_foot_match_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foot_match_request_proto_goTypes,
		DependencyIndexes: file_foot_match_request_proto_depIdxs,
		MessageInfos:      file_foot_match_request_proto_msgTypes,
	}.Build()
	File_foot_match_request_proto = out.File
	file_foot_match_request_proto_rawDesc = nil
	file_foot_match_request_proto_goTypes = nil
	file_foot_match_request_proto_depIdxs = nil
}
