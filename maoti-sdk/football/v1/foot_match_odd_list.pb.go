// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: foot_match_odd_list.proto

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

//指数-列表
type FootOddsMatchListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SportId      string `protobuf:"bytes,1,opt,name=sportId,proto3" json:"sportId,omitempty"`            //1足球 2篮球
	Page         int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`                 //页数
	PageSize     int64  `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`         //条数
	DateTime     string `protobuf:"bytes,4,opt,name=dateTime,proto3" json:"dateTime,omitempty"`          //比赛对应的时间，默认是今天格式：2021-06-08
	TournamentId int64  `protobuf:"varint,5,opt,name=tournamentId,proto3" json:"tournamentId,omitempty"` //赛事搜索对应的赛事id 格式：1,2,3
	Language     string `protobuf:"bytes,6,opt,name=language,proto3" json:"language,omitempty"`          //请求语言
}

func (x *FootOddsMatchListRequest) Reset() {
	*x = FootOddsMatchListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_odd_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootOddsMatchListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootOddsMatchListRequest) ProtoMessage() {}

func (x *FootOddsMatchListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_odd_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootOddsMatchListRequest.ProtoReflect.Descriptor instead.
func (*FootOddsMatchListRequest) Descriptor() ([]byte, []int) {
	return file_foot_match_odd_list_proto_rawDescGZIP(), []int{0}
}

func (x *FootOddsMatchListRequest) GetSportId() string {
	if x != nil {
		return x.SportId
	}
	return ""
}

func (x *FootOddsMatchListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FootOddsMatchListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *FootOddsMatchListRequest) GetDateTime() string {
	if x != nil {
		return x.DateTime
	}
	return ""
}

func (x *FootOddsMatchListRequest) GetTournamentId() int64 {
	if x != nil {
		return x.TournamentId
	}
	return 0
}

func (x *FootOddsMatchListRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type FootOddsMatchListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64       `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"` //总条数
	List  []*FootList `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`    //列表
}

func (x *FootOddsMatchListResponse) Reset() {
	*x = FootOddsMatchListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_odd_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootOddsMatchListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootOddsMatchListResponse) ProtoMessage() {}

func (x *FootOddsMatchListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_odd_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootOddsMatchListResponse.ProtoReflect.Descriptor instead.
func (*FootOddsMatchListResponse) Descriptor() ([]byte, []int) {
	return file_foot_match_odd_list_proto_rawDescGZIP(), []int{1}
}

func (x *FootOddsMatchListResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *FootOddsMatchListResponse) GetList() []*FootList {
	if x != nil {
		return x.List
	}
	return nil
}

type FootList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                  //当前记录ID
	TournamentName     string      `protobuf:"bytes,2,opt,name=tournamentName,proto3" json:"tournamentName,omitempty"`           //联赛名称
	ChName             string      `protobuf:"bytes,3,opt,name=chName,proto3" json:"chName,omitempty"`                           //当前比赛的名称
	TeamMainId         int64       `protobuf:"varint,4,opt,name=teamMainId,proto3" json:"teamMainId,omitempty"`                  //主队ID
	TeamCustId         int64       `protobuf:"varint,5,opt,name=teamCustId,proto3" json:"teamCustId,omitempty"`                  //客队ID
	TeamMainScore      string      `protobuf:"bytes,6,opt,name=teamMainScore,proto3" json:"teamMainScore,omitempty"`             //主队进球总分
	TeamCustScore      string      `protobuf:"bytes,7,opt,name=teamCustScore,proto3" json:"teamCustScore,omitempty"`             //客队进球总分
	MainCustScore      string      `protobuf:"bytes,8,opt,name=mainCustScore,proto3" json:"mainCustScore,omitempty"`             //主客总比分
	Result             string      `protobuf:"bytes,9,opt,name=result,proto3" json:"result,omitempty"`                           //结果
	StartTime          int64       `protobuf:"varint,10,opt,name=startTime,proto3" json:"startTime,omitempty"`                   //当前比赛开始时间戳
	EndTime            int64       `protobuf:"varint,11,opt,name=endTime,proto3" json:"endTime,omitempty"`                       //结束时间
	FirstHalfTime      int64       `protobuf:"varint,12,opt,name=firstHalfTime,proto3" json:"firstHalfTime,omitempty"`           //上半场开始时间戳
	SecondHalfTime     int64       `protobuf:"varint,13,opt,name=secondHalfTime,proto3" json:"secondHalfTime,omitempty"`         //下半场开始时间戳
	CurrentTime        int64       `protobuf:"varint,14,opt,name=currentTime,proto3" json:"currentTime,omitempty"`               //当前时间戳
	Elapsed            int64       `protobuf:"varint,15,opt,name=elapsed,proto3" json:"elapsed,omitempty"`                       //事件发生时的比赛时长
	EventStatusId      int64       `protobuf:"varint,16,opt,name=eventStatusId,proto3" json:"eventStatusId,omitempty"`           //数据中心的比赛状态，具体的对应关系请参照详情页面
	IntermissionStatus int64       `protobuf:"varint,17,opt,name=intermissionStatus,proto3" json:"intermissionStatus,omitempty"` //比赛状态：【1：上半场，2：中场休息，3、下半场，4、其它状态，5、推迟比赛，8：上半场加时赛，9：下半场加时赛】
	EventStatusResult  int64       `protobuf:"varint,18,opt,name=eventStatusResult,proto3" json:"eventStatusResult,omitempty"`   //当前比赛状态 1：未开始，2为开始，3为结束 4特殊状况临时中断比赛  6、取消
	Reason             string      `protobuf:"bytes,19,opt,name=reason,proto3" json:"reason,omitempty"`                          //中断或者临时暂停原因
	HalfCourt          *HalfCourt  `protobuf:"bytes,20,opt,name=halfCourt,proto3" json:"halfCourt,omitempty"`                    //当前半场得分状况
	CornerKick         *CornerKick `protobuf:"bytes,21,opt,name=cornerKick,proto3" json:"cornerKick,omitempty"`                  //角球
	IsCollect          bool        `protobuf:"varint,22,opt,name=isCollect,proto3" json:"isCollect,omitempty"`                   //是否收藏 true:是,false:否
	PlaybackType       int64       `protobuf:"varint,23,opt,name=playbackType,proto3" json:"playbackType,omitempty"`             //播放类型（1:直播，2：3D）
	PlayUrl            string      `protobuf:"bytes,24,opt,name=playUrl,proto3" json:"playUrl,omitempty"`                        //跳转地址
	TeamMain           *TeamMain   `protobuf:"bytes,25,opt,name=teamMain,proto3" json:"teamMain,omitempty"`                      //主队数据信息
	TeamCust           *TeamMain   `protobuf:"bytes,26,opt,name=teamCust,proto3" json:"teamCust,omitempty"`                      //客队数据信息
	PlanNum            int64       `protobuf:"varint,27,opt,name=plan_num,json=planNum,proto3" json:"plan_num,omitempty"`        //方案数量（list_type = 4 时才有值，其他为0）
	Color              string      `protobuf:"bytes,28,opt,name=color,proto3" json:"color,omitempty"`                            //颜色
	Round              string      `protobuf:"bytes,29,opt,name=round,proto3" json:"round,omitempty"`                            //竞彩编号
	BeidanNo           string      `protobuf:"bytes,30,opt,name=beidanNo,proto3" json:"beidanNo,omitempty"`                      //北单编号
	ZucaiNo            string      `protobuf:"bytes,31,opt,name=zucaiNo,proto3" json:"zucaiNo,omitempty"`                        //足彩编号
	Odds               *Odd        `protobuf:"bytes,32,opt,name=odds,proto3" json:"odds,omitempty"`                              //指数
}

func (x *FootList) Reset() {
	*x = FootList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_odd_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootList) ProtoMessage() {}

func (x *FootList) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_odd_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootList.ProtoReflect.Descriptor instead.
func (*FootList) Descriptor() ([]byte, []int) {
	return file_foot_match_odd_list_proto_rawDescGZIP(), []int{2}
}

func (x *FootList) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FootList) GetTournamentName() string {
	if x != nil {
		return x.TournamentName
	}
	return ""
}

func (x *FootList) GetChName() string {
	if x != nil {
		return x.ChName
	}
	return ""
}

func (x *FootList) GetTeamMainId() int64 {
	if x != nil {
		return x.TeamMainId
	}
	return 0
}

func (x *FootList) GetTeamCustId() int64 {
	if x != nil {
		return x.TeamCustId
	}
	return 0
}

func (x *FootList) GetTeamMainScore() string {
	if x != nil {
		return x.TeamMainScore
	}
	return ""
}

func (x *FootList) GetTeamCustScore() string {
	if x != nil {
		return x.TeamCustScore
	}
	return ""
}

func (x *FootList) GetMainCustScore() string {
	if x != nil {
		return x.MainCustScore
	}
	return ""
}

func (x *FootList) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *FootList) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *FootList) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *FootList) GetFirstHalfTime() int64 {
	if x != nil {
		return x.FirstHalfTime
	}
	return 0
}

func (x *FootList) GetSecondHalfTime() int64 {
	if x != nil {
		return x.SecondHalfTime
	}
	return 0
}

func (x *FootList) GetCurrentTime() int64 {
	if x != nil {
		return x.CurrentTime
	}
	return 0
}

func (x *FootList) GetElapsed() int64 {
	if x != nil {
		return x.Elapsed
	}
	return 0
}

func (x *FootList) GetEventStatusId() int64 {
	if x != nil {
		return x.EventStatusId
	}
	return 0
}

func (x *FootList) GetIntermissionStatus() int64 {
	if x != nil {
		return x.IntermissionStatus
	}
	return 0
}

func (x *FootList) GetEventStatusResult() int64 {
	if x != nil {
		return x.EventStatusResult
	}
	return 0
}

func (x *FootList) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *FootList) GetHalfCourt() *HalfCourt {
	if x != nil {
		return x.HalfCourt
	}
	return nil
}

func (x *FootList) GetCornerKick() *CornerKick {
	if x != nil {
		return x.CornerKick
	}
	return nil
}

func (x *FootList) GetIsCollect() bool {
	if x != nil {
		return x.IsCollect
	}
	return false
}

func (x *FootList) GetPlaybackType() int64 {
	if x != nil {
		return x.PlaybackType
	}
	return 0
}

func (x *FootList) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *FootList) GetTeamMain() *TeamMain {
	if x != nil {
		return x.TeamMain
	}
	return nil
}

func (x *FootList) GetTeamCust() *TeamMain {
	if x != nil {
		return x.TeamCust
	}
	return nil
}

func (x *FootList) GetPlanNum() int64 {
	if x != nil {
		return x.PlanNum
	}
	return 0
}

func (x *FootList) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *FootList) GetRound() string {
	if x != nil {
		return x.Round
	}
	return ""
}

func (x *FootList) GetBeidanNo() string {
	if x != nil {
		return x.BeidanNo
	}
	return ""
}

func (x *FootList) GetZucaiNo() string {
	if x != nil {
		return x.ZucaiNo
	}
	return ""
}

func (x *FootList) GetOdds() *Odd {
	if x != nil {
		return x.Odds
	}
	return nil
}

type HalfCourt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamMain int64 `protobuf:"varint,1,opt,name=teamMain,proto3" json:"teamMain,omitempty"`
	TeamCust int64 `protobuf:"varint,2,opt,name=teamCust,proto3" json:"teamCust,omitempty"`
}

func (x *HalfCourt) Reset() {
	*x = HalfCourt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_odd_list_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HalfCourt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HalfCourt) ProtoMessage() {}

func (x *HalfCourt) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_odd_list_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HalfCourt.ProtoReflect.Descriptor instead.
func (*HalfCourt) Descriptor() ([]byte, []int) {
	return file_foot_match_odd_list_proto_rawDescGZIP(), []int{3}
}

func (x *HalfCourt) GetTeamMain() int64 {
	if x != nil {
		return x.TeamMain
	}
	return 0
}

func (x *HalfCourt) GetTeamCust() int64 {
	if x != nil {
		return x.TeamCust
	}
	return 0
}

type CornerKick struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamMain int64 `protobuf:"varint,1,opt,name=teamMain,proto3" json:"teamMain,omitempty"`
	TeamCust int64 `protobuf:"varint,2,opt,name=teamCust,proto3" json:"teamCust,omitempty"`
}

func (x *CornerKick) Reset() {
	*x = CornerKick{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_odd_list_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CornerKick) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CornerKick) ProtoMessage() {}

func (x *CornerKick) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_odd_list_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CornerKick.ProtoReflect.Descriptor instead.
func (*CornerKick) Descriptor() ([]byte, []int) {
	return file_foot_match_odd_list_proto_rawDescGZIP(), []int{4}
}

func (x *CornerKick) GetTeamMain() int64 {
	if x != nil {
		return x.TeamMain
	}
	return 0
}

func (x *CornerKick) GetTeamCust() int64 {
	if x != nil {
		return x.TeamCust
	}
	return 0
}

type TeamMain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LeagueMatchName string `protobuf:"bytes,2,opt,name=leagueMatchName,proto3" json:"leagueMatchName,omitempty"`
	Rank            int64  `protobuf:"varint,3,opt,name=rank,proto3" json:"rank,omitempty"`
	YellowCard      int64  `protobuf:"varint,4,opt,name=yellowCard,proto3" json:"yellowCard,omitempty"`
	RedCard         int64  `protobuf:"varint,5,opt,name=redCard,proto3" json:"redCard,omitempty"`
}

func (x *TeamMain) Reset() {
	*x = TeamMain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_odd_list_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamMain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamMain) ProtoMessage() {}

func (x *TeamMain) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_odd_list_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamMain.ProtoReflect.Descriptor instead.
func (*TeamMain) Descriptor() ([]byte, []int) {
	return file_foot_match_odd_list_proto_rawDescGZIP(), []int{5}
}

func (x *TeamMain) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TeamMain) GetLeagueMatchName() string {
	if x != nil {
		return x.LeagueMatchName
	}
	return ""
}

func (x *TeamMain) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *TeamMain) GetYellowCard() int64 {
	if x != nil {
		return x.YellowCard
	}
	return 0
}

func (x *TeamMain) GetRedCard() int64 {
	if x != nil {
		return x.RedCard
	}
	return 0
}

type Odd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OddsOne   string `protobuf:"bytes,1,opt,name=oddsOne,proto3" json:"oddsOne,omitempty"`     //欧赔的即时指数（格式：1.70 +0/0.5 2.00）
	OddsTwo   string `protobuf:"bytes,2,opt,name=oddsTwo,proto3" json:"oddsTwo,omitempty"`     //亚赔的即时指数【格式：0.5|u,+2.5,0.25|d】
	OddsThree string `protobuf:"bytes,3,opt,name=oddsThree,proto3" json:"oddsThree,omitempty"` //大小球的即时指数【格式：0.5|u,+2.5,0.25|d】
}

func (x *Odd) Reset() {
	*x = Odd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_odd_list_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Odd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Odd) ProtoMessage() {}

func (x *Odd) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_odd_list_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Odd.ProtoReflect.Descriptor instead.
func (*Odd) Descriptor() ([]byte, []int) {
	return file_foot_match_odd_list_proto_rawDescGZIP(), []int{6}
}

func (x *Odd) GetOddsOne() string {
	if x != nil {
		return x.OddsOne
	}
	return ""
}

func (x *Odd) GetOddsTwo() string {
	if x != nil {
		return x.OddsTwo
	}
	return ""
}

func (x *Odd) GetOddsThree() string {
	if x != nil {
		return x.OddsThree
	}
	return ""
}

var File_foot_match_odd_list_proto protoreflect.FileDescriptor

var file_foot_match_odd_list_proto_rawDesc = []byte{
	0x0a, 0x19, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6f, 0x64, 0x64,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x18,
	0x46, 0x6f, 0x6f, 0x74, 0x4f, 0x64, 0x64, 0x73, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x50,
	0x0a, 0x19, 0x46, 0x6f, 0x6f, 0x74, 0x4f, 0x64, 0x64, 0x73, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1d, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x22, 0x9a, 0x08, 0x0a, 0x08, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a,
	0x0e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x74, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x74, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x0d, 0x74, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x65, 0x61, 0x6d,
	0x43, 0x75, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x61, 0x69,
	0x6e, 0x43, 0x75, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x75, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x66, 0x69, 0x72, 0x73, 0x74, 0x48, 0x61, 0x6c, 0x66, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x69, 0x72, 0x73, 0x74, 0x48, 0x61, 0x6c,
	0x66, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x48,
	0x61, 0x6c, 0x66, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x48, 0x61, 0x6c, 0x66, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x12,
	0x2e, 0x0a, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x2c, 0x0a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x09, 0x68, 0x61, 0x6c, 0x66, 0x43, 0x6f, 0x75,
	0x72, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x68, 0x61, 0x6c, 0x66, 0x43,
	0x6f, 0x75, 0x72, 0x74, 0x52, 0x09, 0x68, 0x61, 0x6c, 0x66, 0x43, 0x6f, 0x75, 0x72, 0x74, 0x12,
	0x2b, 0x0a, 0x0a, 0x63, 0x6f, 0x72, 0x6e, 0x65, 0x72, 0x4b, 0x69, 0x63, 0x6b, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x6f, 0x72, 0x6e, 0x65, 0x72, 0x4b, 0x69, 0x63, 0x6b,
	0x52, 0x0a, 0x63, 0x6f, 0x72, 0x6e, 0x65, 0x72, 0x4b, 0x69, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x6c,
	0x61, 0x79, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a, 0x08, 0x74, 0x65, 0x61, 0x6d,
	0x4d, 0x61, 0x69, 0x6e, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x74, 0x65, 0x61,
	0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x12,
	0x25, 0x0a, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x18, 0x1a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x52, 0x08, 0x74, 0x65,
	0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x4e, 0x75,
	0x6d, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x62, 0x65, 0x69, 0x64, 0x61, 0x6e, 0x4e, 0x6f, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x65, 0x69, 0x64, 0x61, 0x6e, 0x4e, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x7a, 0x75, 0x63,
	0x61, 0x69, 0x4e, 0x6f, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a, 0x75, 0x63, 0x61,
	0x69, 0x4e, 0x6f, 0x12, 0x18, 0x0a, 0x04, 0x6f, 0x64, 0x64, 0x73, 0x18, 0x20, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x04, 0x2e, 0x6f, 0x64, 0x64, 0x52, 0x04, 0x6f, 0x64, 0x64, 0x73, 0x22, 0x43, 0x0a,
	0x09, 0x68, 0x61, 0x6c, 0x66, 0x43, 0x6f, 0x75, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65,
	0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x65,
	0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x43, 0x75,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x43, 0x75,
	0x73, 0x74, 0x22, 0x44, 0x0a, 0x0a, 0x63, 0x6f, 0x72, 0x6e, 0x65, 0x72, 0x4b, 0x69, 0x63, 0x6b,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x74, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x22, 0x96, 0x01, 0x0a, 0x08, 0x74, 0x65, 0x61,
	0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6c, 0x65, 0x61,
	0x67, 0x75, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x79, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x43, 0x61, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x79, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x43, 0x61, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x64, 0x43, 0x61,
	0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x64, 0x43, 0x61, 0x72,
	0x64, 0x22, 0x57, 0x0a, 0x03, 0x6f, 0x64, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x64, 0x64, 0x73,
	0x4f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x64, 0x64, 0x73, 0x4f,
	0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x64, 0x64, 0x73, 0x54, 0x77, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x64, 0x64, 0x73, 0x54, 0x77, 0x6f, 0x12, 0x1c, 0x0a, 0x09,
	0x6f, 0x64, 0x64, 0x73, 0x54, 0x68, 0x72, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6f, 0x64, 0x64, 0x73, 0x54, 0x68, 0x72, 0x65, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foot_match_odd_list_proto_rawDescOnce sync.Once
	file_foot_match_odd_list_proto_rawDescData = file_foot_match_odd_list_proto_rawDesc
)

func file_foot_match_odd_list_proto_rawDescGZIP() []byte {
	file_foot_match_odd_list_proto_rawDescOnce.Do(func() {
		file_foot_match_odd_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_foot_match_odd_list_proto_rawDescData)
	})
	return file_foot_match_odd_list_proto_rawDescData
}

var file_foot_match_odd_list_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_foot_match_odd_list_proto_goTypes = []interface{}{
	(*FootOddsMatchListRequest)(nil),  // 0: FootOddsMatchListRequest
	(*FootOddsMatchListResponse)(nil), // 1: FootOddsMatchListResponse
	(*FootList)(nil),                  // 2: FootList
	(*HalfCourt)(nil),                 // 3: halfCourt
	(*CornerKick)(nil),                // 4: cornerKick
	(*TeamMain)(nil),                  // 5: teamMain
	(*Odd)(nil),                       // 6: odd
}
var file_foot_match_odd_list_proto_depIdxs = []int32{
	2, // 0: FootOddsMatchListResponse.list:type_name -> FootList
	3, // 1: FootList.halfCourt:type_name -> halfCourt
	4, // 2: FootList.cornerKick:type_name -> cornerKick
	5, // 3: FootList.teamMain:type_name -> teamMain
	5, // 4: FootList.teamCust:type_name -> teamMain
	6, // 5: FootList.odds:type_name -> odd
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_foot_match_odd_list_proto_init() }
func file_foot_match_odd_list_proto_init() {
	if File_foot_match_odd_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foot_match_odd_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootOddsMatchListRequest); i {
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
		file_foot_match_odd_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootOddsMatchListResponse); i {
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
		file_foot_match_odd_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootList); i {
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
		file_foot_match_odd_list_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HalfCourt); i {
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
		file_foot_match_odd_list_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CornerKick); i {
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
		file_foot_match_odd_list_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamMain); i {
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
		file_foot_match_odd_list_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Odd); i {
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
			RawDescriptor: file_foot_match_odd_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foot_match_odd_list_proto_goTypes,
		DependencyIndexes: file_foot_match_odd_list_proto_depIdxs,
		MessageInfos:      file_foot_match_odd_list_proto_msgTypes,
	}.Build()
	File_foot_match_odd_list_proto = out.File
	file_foot_match_odd_list_proto_rawDesc = nil
	file_foot_match_odd_list_proto_goTypes = nil
	file_foot_match_odd_list_proto_depIdxs = nil
}
