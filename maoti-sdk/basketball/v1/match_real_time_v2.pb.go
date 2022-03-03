// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: match_real_time_v2.proto

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

type MatchRealTimeRequestV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   int64   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`            //即时的状态 0全部,1：未开始，2为开始，3为结束
	DateTime string  `protobuf:"bytes,2,opt,name=dateTime,proto3" json:"dateTime,omitempty"`         //日期 - 全部比赛 2021-11-09
	EventIds []int64 `protobuf:"varint,3,rep,packed,name=eventIds,proto3" json:"eventIds,omitempty"` //比赛筛选
}

func (x *MatchRealTimeRequestV2) Reset() {
	*x = MatchRealTimeRequestV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_real_time_v2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchRealTimeRequestV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchRealTimeRequestV2) ProtoMessage() {}

func (x *MatchRealTimeRequestV2) ProtoReflect() protoreflect.Message {
	mi := &file_match_real_time_v2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchRealTimeRequestV2.ProtoReflect.Descriptor instead.
func (*MatchRealTimeRequestV2) Descriptor() ([]byte, []int) {
	return file_match_real_time_v2_proto_rawDescGZIP(), []int{0}
}

func (x *MatchRealTimeRequestV2) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MatchRealTimeRequestV2) GetDateTime() string {
	if x != nil {
		return x.DateTime
	}
	return ""
}

func (x *MatchRealTimeRequestV2) GetEventIds() []int64 {
	if x != nil {
		return x.EventIds
	}
	return nil
}

type MatchRealTimeResponseV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64                  `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	List  []*MatchRealTimeInfoV2 `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *MatchRealTimeResponseV2) Reset() {
	*x = MatchRealTimeResponseV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_real_time_v2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchRealTimeResponseV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchRealTimeResponseV2) ProtoMessage() {}

func (x *MatchRealTimeResponseV2) ProtoReflect() protoreflect.Message {
	mi := &file_match_real_time_v2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchRealTimeResponseV2.ProtoReflect.Descriptor instead.
func (*MatchRealTimeResponseV2) Descriptor() ([]byte, []int) {
	return file_match_real_time_v2_proto_rawDescGZIP(), []int{1}
}

func (x *MatchRealTimeResponseV2) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *MatchRealTimeResponseV2) GetList() []*MatchRealTimeInfoV2 {
	if x != nil {
		return x.List
	}
	return nil
}

type MatchRealTimeInfoV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                               //当前记录ID 赛事iD用于收藏传递
	ComName            string      `protobuf:"bytes,3,opt,name=comName,proto3" json:"comName,omitempty"`                      //联赛名称
	ComIcon            string      `protobuf:"bytes,4,opt,name=comIcon,proto3" json:"comIcon,omitempty"`                      //联赛图标
	ComId              int64       `protobuf:"varint,5,opt,name=comId,proto3" json:"comId,omitempty"`                         //联赛id
	TeamMainId         int64       `protobuf:"varint,6,opt,name=teamMainId,proto3" json:"teamMainId,omitempty"`               //主队ID[文件：a.proto]
	TeamCustId         int64       `protobuf:"varint,7,opt,name=teamCustId,proto3" json:"teamCustId,omitempty"`               //客队ID
	TeamMainScore      int64       `protobuf:"varint,9,opt,name=teamMainScore,proto3" json:"teamMainScore,omitempty"`         //主队进球总分
	TeamCustScore      int64       `protobuf:"varint,10,opt,name=teamCustScore,proto3" json:"teamCustScore,omitempty"`        //客队进球总分
	MainCustScore      string      `protobuf:"bytes,15,opt,name=mainCustScore,proto3" json:"mainCustScore,omitempty"`         //主客总比分
	Result             string      `protobuf:"bytes,16,opt,name=result,proto3" json:"result,omitempty"`                       //结果[文件：a.proto]
	StartTime          int64       `protobuf:"varint,17,opt,name=startTime,proto3" json:"startTime,omitempty"`                //当前比赛开始时间戳
	EndTime            int64       `protobuf:"varint,8,opt,name=endTime,proto3" json:"endTime,omitempty"`                     //结束时间
	CurrentTime        int64       `protobuf:"varint,11,opt,name=currentTime,proto3" json:"currentTime,omitempty"`            //当前时间戳
	Elapsed            int64       `protobuf:"varint,12,opt,name=elapsed,proto3" json:"elapsed,omitempty"`                    //比赛时长
	EventStatusId      int64       `protobuf:"varint,13,opt,name=eventStatusId,proto3" json:"eventStatusId,omitempty"`        //数据中心的比赛状态，具体的对应关系请参照详情页面
	EventStatusResult  int64       `protobuf:"varint,2,opt,name=eventStatusResult,proto3" json:"eventStatusResult,omitempty"` //当前比赛状态 1：未开始，2为开始，3为结束 4特殊状况临时中断比赛  5、延迟  6、取消
	TimeRemaining      int64       `protobuf:"varint,23,opt,name=timeRemaining,proto3" json:"timeRemaining,omitempty"`        //比赛节数剩余时间
	Reason             string      `protobuf:"bytes,18,opt,name=reason,proto3" json:"reason,omitempty"`                       //中断或者临时暂停原因
	MainTeam           *TeamInfoV2 `protobuf:"bytes,19,opt,name=mainTeam,proto3" json:"mainTeam,omitempty"`
	AwayTeam           *TeamInfoV2 `protobuf:"bytes,20,opt,name=awayTeam,proto3" json:"awayTeam,omitempty"`
	LiveMatches        []*LiveV2   `protobuf:"bytes,21,rep,name=live_matches,json=liveMatches,proto3" json:"live_matches,omitempty"`
	Odds               []*OddsV2   `protobuf:"bytes,22,rep,name=odds,proto3" json:"odds,omitempty"`                             //实时指数
	LotteryCompetition string      `protobuf:"bytes,27,opt,name=lotteryCompetition,proto3" json:"lotteryCompetition,omitempty"` //竞彩赛事编号例：周二110
	Color              string      `protobuf:"bytes,25,opt,name=color,proto3" json:"color,omitempty"`                           //颜色
	MatchStatus        string      `protobuf:"bytes,26,opt,name=matchStatus,proto3" json:"matchStatus,omitempty"`               //加时节数
	PlaybackType       int64       `protobuf:"varint,29,opt,name=playbackType,proto3" json:"playbackType,omitempty"`            //播放类型（1:直播，2：2D动画，3 全部都有）
	ClassifyTime       string      `protobuf:"bytes,30,opt,name=classifyTime,proto3" json:"classifyTime,omitempty"`             //分类用的时间
}

func (x *MatchRealTimeInfoV2) Reset() {
	*x = MatchRealTimeInfoV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_real_time_v2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchRealTimeInfoV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchRealTimeInfoV2) ProtoMessage() {}

func (x *MatchRealTimeInfoV2) ProtoReflect() protoreflect.Message {
	mi := &file_match_real_time_v2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchRealTimeInfoV2.ProtoReflect.Descriptor instead.
func (*MatchRealTimeInfoV2) Descriptor() ([]byte, []int) {
	return file_match_real_time_v2_proto_rawDescGZIP(), []int{2}
}

func (x *MatchRealTimeInfoV2) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MatchRealTimeInfoV2) GetComName() string {
	if x != nil {
		return x.ComName
	}
	return ""
}

func (x *MatchRealTimeInfoV2) GetComIcon() string {
	if x != nil {
		return x.ComIcon
	}
	return ""
}

func (x *MatchRealTimeInfoV2) GetComId() int64 {
	if x != nil {
		return x.ComId
	}
	return 0
}

func (x *MatchRealTimeInfoV2) GetTeamMainId() int64 {
	if x != nil {
		return x.TeamMainId
	}
	return 0
}

func (x *MatchRealTimeInfoV2) GetTeamCustId() int64 {
	if x != nil {
		return x.TeamCustId
	}
	return 0
}

func (x *MatchRealTimeInfoV2) GetTeamMainScore() int64 {
	if x != nil {
		return x.TeamMainScore
	}
	return 0
}

func (x *MatchRealTimeInfoV2) GetTeamCustScore() int64 {
	if x != nil {
		return x.TeamCustScore
	}
	return 0
}

func (x *MatchRealTimeInfoV2) GetMainCustScore() string {
	if x != nil {
		return x.MainCustScore
	}
	return ""
}

func (x *MatchRealTimeInfoV2) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *MatchRealTimeInfoV2) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *MatchRealTimeInfoV2) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *MatchRealTimeInfoV2) GetCurrentTime() int64 {
	if x != nil {
		return x.CurrentTime
	}
	return 0
}

func (x *MatchRealTimeInfoV2) GetElapsed() int64 {
	if x != nil {
		return x.Elapsed
	}
	return 0
}

func (x *MatchRealTimeInfoV2) GetEventStatusId() int64 {
	if x != nil {
		return x.EventStatusId
	}
	return 0
}

func (x *MatchRealTimeInfoV2) GetEventStatusResult() int64 {
	if x != nil {
		return x.EventStatusResult
	}
	return 0
}

func (x *MatchRealTimeInfoV2) GetTimeRemaining() int64 {
	if x != nil {
		return x.TimeRemaining
	}
	return 0
}

func (x *MatchRealTimeInfoV2) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *MatchRealTimeInfoV2) GetMainTeam() *TeamInfoV2 {
	if x != nil {
		return x.MainTeam
	}
	return nil
}

func (x *MatchRealTimeInfoV2) GetAwayTeam() *TeamInfoV2 {
	if x != nil {
		return x.AwayTeam
	}
	return nil
}

func (x *MatchRealTimeInfoV2) GetLiveMatches() []*LiveV2 {
	if x != nil {
		return x.LiveMatches
	}
	return nil
}

func (x *MatchRealTimeInfoV2) GetOdds() []*OddsV2 {
	if x != nil {
		return x.Odds
	}
	return nil
}

func (x *MatchRealTimeInfoV2) GetLotteryCompetition() string {
	if x != nil {
		return x.LotteryCompetition
	}
	return ""
}

func (x *MatchRealTimeInfoV2) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *MatchRealTimeInfoV2) GetMatchStatus() string {
	if x != nil {
		return x.MatchStatus
	}
	return ""
}

func (x *MatchRealTimeInfoV2) GetPlaybackType() int64 {
	if x != nil {
		return x.PlaybackType
	}
	return 0
}

func (x *MatchRealTimeInfoV2) GetClassifyTime() string {
	if x != nil {
		return x.ClassifyTime
	}
	return ""
}

type TeamInfoV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`         //主队名称
	TeamId   int64          `protobuf:"varint,3,opt,name=teamId,proto3" json:"teamId,omitempty"`    //队伍id
	TeamIcon string         `protobuf:"bytes,6,opt,name=teamIcon,proto3" json:"teamIcon,omitempty"` //图标
	ComName  string         `protobuf:"bytes,4,opt,name=comName,proto3" json:"comName,omitempty"`   //当前队伍的联赛名称
	Rank     string         `protobuf:"bytes,9,opt,name=rank,proto3" json:"rank,omitempty"`         //排名
	Scores   []*RealscoreV2 `protobuf:"bytes,5,rep,name=scores,proto3" json:"scores,omitempty"`     //得分数组
}

func (x *TeamInfoV2) Reset() {
	*x = TeamInfoV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_real_time_v2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamInfoV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamInfoV2) ProtoMessage() {}

func (x *TeamInfoV2) ProtoReflect() protoreflect.Message {
	mi := &file_match_real_time_v2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamInfoV2.ProtoReflect.Descriptor instead.
func (*TeamInfoV2) Descriptor() ([]byte, []int) {
	return file_match_real_time_v2_proto_rawDescGZIP(), []int{3}
}

func (x *TeamInfoV2) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TeamInfoV2) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *TeamInfoV2) GetTeamIcon() string {
	if x != nil {
		return x.TeamIcon
	}
	return ""
}

func (x *TeamInfoV2) GetComName() string {
	if x != nil {
		return x.ComName
	}
	return ""
}

func (x *TeamInfoV2) GetRank() string {
	if x != nil {
		return x.Rank
	}
	return ""
}

func (x *TeamInfoV2) GetScores() []*RealscoreV2 {
	if x != nil {
		return x.Scores
	}
	return nil
}

type RealscoreV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag   string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`      //节数，默认顺序为：一	二	三	四 节
	Score int64  `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"` //得分
}

func (x *RealscoreV2) Reset() {
	*x = RealscoreV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_real_time_v2_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RealscoreV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RealscoreV2) ProtoMessage() {}

func (x *RealscoreV2) ProtoReflect() protoreflect.Message {
	mi := &file_match_real_time_v2_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RealscoreV2.ProtoReflect.Descriptor instead.
func (*RealscoreV2) Descriptor() ([]byte, []int) {
	return file_match_real_time_v2_proto_rawDescGZIP(), []int{4}
}

func (x *RealscoreV2) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *RealscoreV2) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type LiveV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Desc        string `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty"`               //描述说明
	AddTime     int64  `protobuf:"varint,2,opt,name=addTime,proto3" json:"addTime,omitempty"`        //添加时间
	CommentTime string `protobuf:"bytes,3,opt,name=commentTime,proto3" json:"commentTime,omitempty"` //时间
}

func (x *LiveV2) Reset() {
	*x = LiveV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_real_time_v2_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveV2) ProtoMessage() {}

func (x *LiveV2) ProtoReflect() protoreflect.Message {
	mi := &file_match_real_time_v2_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveV2.ProtoReflect.Descriptor instead.
func (*LiveV2) Descriptor() ([]byte, []int) {
	return file_match_real_time_v2_proto_rawDescGZIP(), []int{5}
}

func (x *LiveV2) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *LiveV2) GetAddTime() int64 {
	if x != nil {
		return x.AddTime
	}
	return 0
}

func (x *LiveV2) GetCommentTime() string {
	if x != nil {
		return x.CommentTime
	}
	return ""
}

type OddsV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OddsTitle string `protobuf:"bytes,1,opt,name=oddsTitle,proto3" json:"oddsTitle,omitempty"`
	OddsMain  string `protobuf:"bytes,2,opt,name=oddsMain,proto3" json:"oddsMain,omitempty"`
	OddsCust  string `protobuf:"bytes,3,opt,name=oddsCust,proto3" json:"oddsCust,omitempty"`
	OddsType  string `protobuf:"bytes,4,opt,name=oddsType,proto3" json:"oddsType,omitempty"`
}

func (x *OddsV2) Reset() {
	*x = OddsV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_real_time_v2_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OddsV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OddsV2) ProtoMessage() {}

func (x *OddsV2) ProtoReflect() protoreflect.Message {
	mi := &file_match_real_time_v2_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OddsV2.ProtoReflect.Descriptor instead.
func (*OddsV2) Descriptor() ([]byte, []int) {
	return file_match_real_time_v2_proto_rawDescGZIP(), []int{6}
}

func (x *OddsV2) GetOddsTitle() string {
	if x != nil {
		return x.OddsTitle
	}
	return ""
}

func (x *OddsV2) GetOddsMain() string {
	if x != nil {
		return x.OddsMain
	}
	return ""
}

func (x *OddsV2) GetOddsCust() string {
	if x != nil {
		return x.OddsCust
	}
	return ""
}

func (x *OddsV2) GetOddsType() string {
	if x != nil {
		return x.OddsType
	}
	return ""
}

var File_match_real_time_v2_proto protoreflect.FileDescriptor

var file_match_real_time_v2_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x16, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x56, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x73, 0x22, 0x59, 0x0a, 0x17, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x61,
	0x6c, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x32, 0x12,
	0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x61, 0x6c, 0x54,
	0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x56, 0x32, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22,
	0x8a, 0x07, 0x0a, 0x13, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x61, 0x6c, 0x54, 0x69, 0x6d,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x56, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x6d, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x6d, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x49, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x65, 0x61, 0x6d, 0x4d, 0x61,
	0x69, 0x6e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x65, 0x61, 0x6d, 0x43,
	0x75, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x74, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x75, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x75, 0x73, 0x74, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x12,
	0x24, 0x0a, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x6d, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x18, 0x17, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65,
	0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x27, 0x0a, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x56, 0x32,
	0x52, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x27, 0x0a, 0x08, 0x61, 0x77,
	0x61, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74,
	0x65, 0x61, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x56, 0x32, 0x52, 0x08, 0x61, 0x77, 0x61, 0x79, 0x54,
	0x65, 0x61, 0x6d, 0x12, 0x2a, 0x0a, 0x0c, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x6c, 0x69, 0x76, 0x65,
	0x56, 0x32, 0x52, 0x0b, 0x6c, 0x69, 0x76, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12,
	0x1b, 0x0a, 0x04, 0x6f, 0x64, 0x64, 0x73, 0x18, 0x16, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x6f, 0x64, 0x64, 0x73, 0x56, 0x32, 0x52, 0x04, 0x6f, 0x64, 0x64, 0x73, 0x12, 0x2e, 0x0a, 0x12,
	0x6c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6c, 0x6f, 0x74, 0x74, 0x65, 0x72,
	0x79, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x79,
	0x62, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x66, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xa8, 0x01, 0x0a,
	0x0a, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x56, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x49,
	0x63, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x49,
	0x63, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x61, 0x6e,
	0x6b, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x56, 0x32, 0x52,
	0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x22, 0x35, 0x0a, 0x0b, 0x72, 0x65, 0x61, 0x6c, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x56, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x58,
	0x0a, 0x06, 0x6c, 0x69, 0x76, 0x65, 0x56, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x7a, 0x0a, 0x06, 0x6f, 0x64, 0x64, 0x73,
	0x56, 0x32, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x64, 0x64, 0x73, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x64, 0x64, 0x73, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x64, 0x64, 0x73, 0x4d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6f, 0x64, 0x64, 0x73, 0x4d, 0x61, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x6f, 0x64, 0x64, 0x73, 0x43, 0x75, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6f, 0x64, 0x64, 0x73, 0x43, 0x75, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x64, 0x64, 0x73,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x64, 0x64, 0x73,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_match_real_time_v2_proto_rawDescOnce sync.Once
	file_match_real_time_v2_proto_rawDescData = file_match_real_time_v2_proto_rawDesc
)

func file_match_real_time_v2_proto_rawDescGZIP() []byte {
	file_match_real_time_v2_proto_rawDescOnce.Do(func() {
		file_match_real_time_v2_proto_rawDescData = protoimpl.X.CompressGZIP(file_match_real_time_v2_proto_rawDescData)
	})
	return file_match_real_time_v2_proto_rawDescData
}

var file_match_real_time_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_match_real_time_v2_proto_goTypes = []interface{}{
	(*MatchRealTimeRequestV2)(nil),  // 0: MatchRealTimeRequestV2
	(*MatchRealTimeResponseV2)(nil), // 1: MatchRealTimeResponseV2
	(*MatchRealTimeInfoV2)(nil),     // 2: MatchRealTimeInfoV2
	(*TeamInfoV2)(nil),              // 3: teamInfoV2
	(*RealscoreV2)(nil),             // 4: realscoreV2
	(*LiveV2)(nil),                  // 5: liveV2
	(*OddsV2)(nil),                  // 6: oddsV2
}
var file_match_real_time_v2_proto_depIdxs = []int32{
	2, // 0: MatchRealTimeResponseV2.list:type_name -> MatchRealTimeInfoV2
	3, // 1: MatchRealTimeInfoV2.mainTeam:type_name -> teamInfoV2
	3, // 2: MatchRealTimeInfoV2.awayTeam:type_name -> teamInfoV2
	5, // 3: MatchRealTimeInfoV2.live_matches:type_name -> liveV2
	6, // 4: MatchRealTimeInfoV2.odds:type_name -> oddsV2
	4, // 5: teamInfoV2.scores:type_name -> realscoreV2
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_match_real_time_v2_proto_init() }
func file_match_real_time_v2_proto_init() {
	if File_match_real_time_v2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_match_real_time_v2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchRealTimeRequestV2); i {
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
		file_match_real_time_v2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchRealTimeResponseV2); i {
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
		file_match_real_time_v2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchRealTimeInfoV2); i {
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
		file_match_real_time_v2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamInfoV2); i {
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
		file_match_real_time_v2_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RealscoreV2); i {
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
		file_match_real_time_v2_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveV2); i {
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
		file_match_real_time_v2_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OddsV2); i {
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
			RawDescriptor: file_match_real_time_v2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_match_real_time_v2_proto_goTypes,
		DependencyIndexes: file_match_real_time_v2_proto_depIdxs,
		MessageInfos:      file_match_real_time_v2_proto_msgTypes,
	}.Build()
	File_match_real_time_v2_proto = out.File
	file_match_real_time_v2_proto_rawDesc = nil
	file_match_real_time_v2_proto_goTypes = nil
	file_match_real_time_v2_proto_depIdxs = nil
}
