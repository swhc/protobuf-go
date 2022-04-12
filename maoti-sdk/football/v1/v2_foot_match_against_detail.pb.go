// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: v2_foot_match_against_detail.proto

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
type V2FootMatchAgainstDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"` //请求语言
	SportId  int64  `protobuf:"varint,1,opt,name=sportId,proto3" json:"sportId,omitempty"`  //1足球 2篮球
	EventId  int64  `protobuf:"varint,2,opt,name=eventId,proto3" json:"eventId,omitempty"`  //比赛id
}

func (x *V2FootMatchAgainstDetailRequest) Reset() {
	*x = V2FootMatchAgainstDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_match_against_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchAgainstDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchAgainstDetailRequest) ProtoMessage() {}

func (x *V2FootMatchAgainstDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_match_against_detail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchAgainstDetailRequest.ProtoReflect.Descriptor instead.
func (*V2FootMatchAgainstDetailRequest) Descriptor() ([]byte, []int) {
	return file_v2_foot_match_against_detail_proto_rawDescGZIP(), []int{0}
}

func (x *V2FootMatchAgainstDetailRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *V2FootMatchAgainstDetailRequest) GetSportId() int64 {
	if x != nil {
		return x.SportId
	}
	return 0
}

func (x *V2FootMatchAgainstDetailRequest) GetEventId() int64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

type V2FootMatchAgainstDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64                `protobuf:"varint,3,opt,name=Id,proto3" json:"Id,omitempty"`                                                   //比赛id
	ChName            string               `protobuf:"bytes,4,opt,name=ChName,proto3" json:"ChName,omitempty"`                                            //比赛名称
	Round             string               `protobuf:"bytes,5,opt,name=Round,proto3" json:"Round,omitempty"`                                              //场次
	StartTime         int64                `protobuf:"varint,6,opt,name=StartTime,proto3" json:"StartTime,omitempty"`                                     //开始时间
	MainTeamId        int64                `protobuf:"varint,7,opt,name=MainTeamId,proto3" json:"MainTeamId,omitempty"`                                   //主队id
	MainTeamName      string               `protobuf:"bytes,8,opt,name=MainTeamName,proto3" json:"MainTeamName,omitempty"`                                //主队名称
	MainTeamImages    string               `protobuf:"bytes,9,opt,name=MainTeamImages,proto3" json:"MainTeamImages,omitempty"`                            //主队头像
	CustTeamId        int64                `protobuf:"varint,10,opt,name=CustTeamId,proto3" json:"CustTeamId,omitempty"`                                  //客队id
	CustTeamName      string               `protobuf:"bytes,11,opt,name=CustTeamName,proto3" json:"CustTeamName,omitempty"`                               //客队名称
	CustTeamImages    string               `protobuf:"bytes,12,opt,name=CustTeamImages,proto3" json:"CustTeamImages,omitempty"`                           //客队头像
	MainTeamRank      int64                `protobuf:"varint,13,opt,name=MainTeamRank,proto3" json:"MainTeamRank,omitempty"`                              //主队排名
	CustTeamRank      int64                `protobuf:"varint,14,opt,name=CustTeamRank,proto3" json:"CustTeamRank,omitempty"`                              //客队排名
	EventStatusResult int64                `protobuf:"varint,15,opt,name=EventStatusResult,proto3" json:"EventStatusResult,omitempty"`                    //状态描述：1:未开始，2:开始，3:结束,4:中断，6:取消
	IsCollect         bool                 `protobuf:"varint,16,opt,name=IsCollect,proto3" json:"IsCollect,omitempty"`                                    //是否收藏
	TeamMainScore     int64                `protobuf:"varint,17,opt,name=TeamMainScore,proto3" json:"TeamMainScore,omitempty"`                            //主队得分
	TeamCustScore     int64                `protobuf:"varint,18,opt,name=TeamCustScore,proto3" json:"TeamCustScore,omitempty"`                            //客队得分
	TeamMainRed       int64                `protobuf:"varint,41,opt,name=TeamMainRed,proto3" json:"TeamMainRed,omitempty"`                                //主队红牌
	TeamCustRed       int64                `protobuf:"varint,42,opt,name=TeamCustRed,proto3" json:"TeamCustRed,omitempty"`                                //客队红牌
	TeamMainYellow    int64                `protobuf:"varint,43,opt,name=TeamMainYellow,proto3" json:"TeamMainYellow,omitempty"`                          //主队黄牌
	TeamCustYellow    int64                `protobuf:"varint,44,opt,name=TeamCustYellow,proto3" json:"TeamCustYellow,omitempty"`                          //客队黄牌
	TeamMainCorner    int64                `protobuf:"varint,45,opt,name=TeamMainCorner,proto3" json:"TeamMainCorner,omitempty"`                          //主队角球
	TeamCustCorner    int64                `protobuf:"varint,46,opt,name=TeamCustCorner,proto3" json:"TeamCustCorner,omitempty"`                          //客队角球
	EventStatusId     int64                `protobuf:"varint,19,opt,name=EventStatusId,proto3" json:"EventStatusId,omitempty"`                            //2：上半场，10：中场休息，3：下半场，14：等待加时，8：加时赛上半场，9：加时赛下半场，20：等待点球，4：点球大战，13：点球大战完场
	TimePlayed        int64                `protobuf:"varint,20,opt,name=TimePlayed,proto3" json:"TimePlayed,omitempty"`                                  //足球比赛已进行时间（秒）；示例：5656
	TimeRunning       int64                `protobuf:"varint,21,opt,name=TimeRunning,proto3" json:"TimeRunning,omitempty"`                                //比赛时钟是否计时中：1计时中 0暂停计时；示例：0
	TimeUpdate        int64                `protobuf:"varint,32,opt,name=TimeUpdate,proto3" json:"TimeUpdate,omitempty"`                                  //比赛进行时间更新时间；示例：2020-10-25 00:46:22.000
	UpperTime         int64                `protobuf:"varint,35,opt,name=upper_time,json=upperTime,proto3" json:"upper_time,omitempty"`                   //上场开始时间（或加时赛）
	LowerTime         int64                `protobuf:"varint,36,opt,name=lower_time,json=lowerTime,proto3" json:"lower_time,omitempty"`                   //下场开始时间（或加时赛）
	MainScore         int64                `protobuf:"varint,22,opt,name=MainScore,proto3" json:"MainScore,omitempty"`                                    //（加时赛）主队得分
	MainPointSphere   int64                `protobuf:"varint,23,opt,name=MainPoint_sphere,json=MainPointSphere,proto3" json:"MainPoint_sphere,omitempty"` //主队点球数
	CustScore         int64                `protobuf:"varint,24,opt,name=CustScore,proto3" json:"CustScore,omitempty"`                                    //（加时赛）客队得分
	CustPointSphere   int64                `protobuf:"varint,25,opt,name=CustPoint_sphere,json=CustPointSphere,proto3" json:"CustPoint_sphere,omitempty"` //客队点球数
	IsExtraTime       bool                 `protobuf:"varint,31,opt,name=IsExtraTime,proto3" json:"IsExtraTime,omitempty"`                                //是否加时赛
	IsPenaltyKick     bool                 `protobuf:"varint,26,opt,name=IsPenaltyKick,proto3" json:"IsPenaltyKick,omitempty"`                            //是否点球
	Elapsed           int64                `protobuf:"varint,27,opt,name=Elapsed,proto3" json:"Elapsed,omitempty"`                                        //时长
	CurrentTime       int64                `protobuf:"varint,28,opt,name=CurrentTime,proto3" json:"CurrentTime,omitempty"`                                //服务器当前时间
	PlaybackType      int64                `protobuf:"varint,33,opt,name=PlaybackType,proto3" json:"PlaybackType,omitempty"`                              //1:直播，2：3D ,3:都存在
	AgainstLive       []*V2FootAgainstLive `protobuf:"bytes,34,rep,name=againstLive,proto3" json:"againstLive,omitempty"`                                 //直播地址
	Animation         *V2FootAnimation     `protobuf:"bytes,30,opt,name=animation,proto3" json:"animation,omitempty"`                                     //动画
	StreamName        string               `protobuf:"bytes,37,opt,name=StreamName,proto3" json:"StreamName,omitempty"`
	TournamentId      int64                `protobuf:"varint,38,opt,name=tournamentId,proto3" json:"tournamentId,omitempty"`    // 联赛id
	TournamentLogo    string               `protobuf:"bytes,39,opt,name=TournamentLogo,proto3" json:"TournamentLogo,omitempty"` //联赛logo
	MatchCount        int64                `protobuf:"varint,40,opt,name=MatchCount,proto3" json:"MatchCount,omitempty"`        //近30天比赛数量
}

func (x *V2FootMatchAgainstDetailResponse) Reset() {
	*x = V2FootMatchAgainstDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_match_against_detail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchAgainstDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchAgainstDetailResponse) ProtoMessage() {}

func (x *V2FootMatchAgainstDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_match_against_detail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchAgainstDetailResponse.ProtoReflect.Descriptor instead.
func (*V2FootMatchAgainstDetailResponse) Descriptor() ([]byte, []int) {
	return file_v2_foot_match_against_detail_proto_rawDescGZIP(), []int{1}
}

func (x *V2FootMatchAgainstDetailResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetChName() string {
	if x != nil {
		return x.ChName
	}
	return ""
}

func (x *V2FootMatchAgainstDetailResponse) GetRound() string {
	if x != nil {
		return x.Round
	}
	return ""
}

func (x *V2FootMatchAgainstDetailResponse) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetMainTeamId() int64 {
	if x != nil {
		return x.MainTeamId
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetMainTeamName() string {
	if x != nil {
		return x.MainTeamName
	}
	return ""
}

func (x *V2FootMatchAgainstDetailResponse) GetMainTeamImages() string {
	if x != nil {
		return x.MainTeamImages
	}
	return ""
}

func (x *V2FootMatchAgainstDetailResponse) GetCustTeamId() int64 {
	if x != nil {
		return x.CustTeamId
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetCustTeamName() string {
	if x != nil {
		return x.CustTeamName
	}
	return ""
}

func (x *V2FootMatchAgainstDetailResponse) GetCustTeamImages() string {
	if x != nil {
		return x.CustTeamImages
	}
	return ""
}

func (x *V2FootMatchAgainstDetailResponse) GetMainTeamRank() int64 {
	if x != nil {
		return x.MainTeamRank
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetCustTeamRank() int64 {
	if x != nil {
		return x.CustTeamRank
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetEventStatusResult() int64 {
	if x != nil {
		return x.EventStatusResult
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetIsCollect() bool {
	if x != nil {
		return x.IsCollect
	}
	return false
}

func (x *V2FootMatchAgainstDetailResponse) GetTeamMainScore() int64 {
	if x != nil {
		return x.TeamMainScore
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetTeamCustScore() int64 {
	if x != nil {
		return x.TeamCustScore
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetTeamMainRed() int64 {
	if x != nil {
		return x.TeamMainRed
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetTeamCustRed() int64 {
	if x != nil {
		return x.TeamCustRed
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetTeamMainYellow() int64 {
	if x != nil {
		return x.TeamMainYellow
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetTeamCustYellow() int64 {
	if x != nil {
		return x.TeamCustYellow
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetTeamMainCorner() int64 {
	if x != nil {
		return x.TeamMainCorner
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetTeamCustCorner() int64 {
	if x != nil {
		return x.TeamCustCorner
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetEventStatusId() int64 {
	if x != nil {
		return x.EventStatusId
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetTimePlayed() int64 {
	if x != nil {
		return x.TimePlayed
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetTimeRunning() int64 {
	if x != nil {
		return x.TimeRunning
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetTimeUpdate() int64 {
	if x != nil {
		return x.TimeUpdate
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetUpperTime() int64 {
	if x != nil {
		return x.UpperTime
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetLowerTime() int64 {
	if x != nil {
		return x.LowerTime
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetMainScore() int64 {
	if x != nil {
		return x.MainScore
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetMainPointSphere() int64 {
	if x != nil {
		return x.MainPointSphere
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetCustScore() int64 {
	if x != nil {
		return x.CustScore
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetCustPointSphere() int64 {
	if x != nil {
		return x.CustPointSphere
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetIsExtraTime() bool {
	if x != nil {
		return x.IsExtraTime
	}
	return false
}

func (x *V2FootMatchAgainstDetailResponse) GetIsPenaltyKick() bool {
	if x != nil {
		return x.IsPenaltyKick
	}
	return false
}

func (x *V2FootMatchAgainstDetailResponse) GetElapsed() int64 {
	if x != nil {
		return x.Elapsed
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetCurrentTime() int64 {
	if x != nil {
		return x.CurrentTime
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetPlaybackType() int64 {
	if x != nil {
		return x.PlaybackType
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetAgainstLive() []*V2FootAgainstLive {
	if x != nil {
		return x.AgainstLive
	}
	return nil
}

func (x *V2FootMatchAgainstDetailResponse) GetAnimation() *V2FootAnimation {
	if x != nil {
		return x.Animation
	}
	return nil
}

func (x *V2FootMatchAgainstDetailResponse) GetStreamName() string {
	if x != nil {
		return x.StreamName
	}
	return ""
}

func (x *V2FootMatchAgainstDetailResponse) GetTournamentId() int64 {
	if x != nil {
		return x.TournamentId
	}
	return 0
}

func (x *V2FootMatchAgainstDetailResponse) GetTournamentLogo() string {
	if x != nil {
		return x.TournamentLogo
	}
	return ""
}

func (x *V2FootMatchAgainstDetailResponse) GetMatchCount() int64 {
	if x != nil {
		return x.MatchCount
	}
	return 0
}

type V2FootAgainstLive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"` //名称
	Url   string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`     //地址
}

func (x *V2FootAgainstLive) Reset() {
	*x = V2FootAgainstLive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_match_against_detail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootAgainstLive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootAgainstLive) ProtoMessage() {}

func (x *V2FootAgainstLive) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_match_against_detail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootAgainstLive.ProtoReflect.Descriptor instead.
func (*V2FootAgainstLive) Descriptor() ([]byte, []int) {
	return file_v2_foot_match_against_detail_proto_rawDescGZIP(), []int{2}
}

func (x *V2FootAgainstLive) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *V2FootAgainstLive) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type V2FootAnimation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url    string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`        //链接
	Width  int64  `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`   //宽度
	Height int64  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"` //高度
}

func (x *V2FootAnimation) Reset() {
	*x = V2FootAnimation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_match_against_detail_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootAnimation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootAnimation) ProtoMessage() {}

func (x *V2FootAnimation) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_match_against_detail_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootAnimation.ProtoReflect.Descriptor instead.
func (*V2FootAnimation) Descriptor() ([]byte, []int) {
	return file_v2_foot_match_against_detail_proto_rawDescGZIP(), []int{3}
}

func (x *V2FootAnimation) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *V2FootAnimation) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *V2FootAnimation) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

//v2 即时比赛列表 - 比赛详情 - 预测投票
type V2FootMatchPredictiveVotingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId int64 `protobuf:"varint,1,opt,name=eventId,proto3" json:"eventId,omitempty"` //比赛ID
}

func (x *V2FootMatchPredictiveVotingRequest) Reset() {
	*x = V2FootMatchPredictiveVotingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_match_against_detail_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchPredictiveVotingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchPredictiveVotingRequest) ProtoMessage() {}

func (x *V2FootMatchPredictiveVotingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_match_against_detail_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchPredictiveVotingRequest.ProtoReflect.Descriptor instead.
func (*V2FootMatchPredictiveVotingRequest) Descriptor() ([]byte, []int) {
	return file_v2_foot_match_against_detail_proto_rawDescGZIP(), []int{4}
}

func (x *V2FootMatchPredictiveVotingRequest) GetEventId() int64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

type V2FootMatchPredictiveVotingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Handicap string `protobuf:"bytes,1,opt,name=handicap,proto3" json:"handicap,omitempty"` //让球
	BigBall  string `protobuf:"bytes,2,opt,name=bigBall,proto3" json:"bigBall,omitempty"`   //大小球
}

func (x *V2FootMatchPredictiveVotingResponse) Reset() {
	*x = V2FootMatchPredictiveVotingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_match_against_detail_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchPredictiveVotingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchPredictiveVotingResponse) ProtoMessage() {}

func (x *V2FootMatchPredictiveVotingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_match_against_detail_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchPredictiveVotingResponse.ProtoReflect.Descriptor instead.
func (*V2FootMatchPredictiveVotingResponse) Descriptor() ([]byte, []int) {
	return file_v2_foot_match_against_detail_proto_rawDescGZIP(), []int{5}
}

func (x *V2FootMatchPredictiveVotingResponse) GetHandicap() string {
	if x != nil {
		return x.Handicap
	}
	return ""
}

func (x *V2FootMatchPredictiveVotingResponse) GetBigBall() string {
	if x != nil {
		return x.BigBall
	}
	return ""
}

var File_v2_foot_match_against_detail_proto protoreflect.FileDescriptor

var file_v2_foot_match_against_detail_proto_rawDesc = []byte{
	0x0a, 0x22, 0x76, 0x32, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x61, 0x67, 0x61, 0x69, 0x6e, 0x73, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x1f, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x41, 0x67, 0x61, 0x69, 0x6e, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x8c, 0x0c, 0x0a, 0x20, 0x56, 0x32, 0x46, 0x6f,
	0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x41, 0x67, 0x61, 0x69, 0x6e, 0x73, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x43, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x68,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x61, 0x69, 0x6e,
	0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4d, 0x61,
	0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x61, 0x69, 0x6e,
	0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x4d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x4d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x4d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d,
	0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x54, 0x65,
	0x61, 0x6d, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x75, 0x73, 0x74,
	0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x43, 0x75, 0x73, 0x74,
	0x54, 0x65, 0x61, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x43, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x61, 0x6e, 0x6b,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x4d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d,
	0x52, 0x61, 0x6e, 0x6b, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d,
	0x52, 0x61, 0x6e, 0x6b, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x43, 0x75, 0x73, 0x74,
	0x54, 0x65, 0x61, 0x6d, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x2c, 0x0a, 0x11, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x11, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x54, 0x65, 0x61,
	0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x65,
	0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x64, 0x18,
	0x29, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x52,
	0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x52, 0x65,
	0x64, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73,
	0x74, 0x52, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e,
	0x59, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x54, 0x65,
	0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x59, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x26, 0x0a, 0x0e,
	0x54, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x59, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x2c,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x59, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x26, 0x0a, 0x0e, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e,
	0x43, 0x6f, 0x72, 0x6e, 0x65, 0x72, 0x18, 0x2d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x54, 0x65,
	0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x72, 0x6e, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e,
	0x54, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x43, 0x6f, 0x72, 0x6e, 0x65, 0x72, 0x18, 0x2e,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x43, 0x6f,
	0x72, 0x6e, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x49, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x69,
	0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x54, 0x69, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a,
	0x54, 0x69, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x23, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x75, 0x70, 0x70, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x24, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x61,
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
	0x12, 0x22, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x21, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x61, 0x67, 0x61, 0x69, 0x6e, 0x73, 0x74, 0x4c,
	0x69, 0x76, 0x65, 0x18, 0x22, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x56, 0x32, 0x46, 0x6f,
	0x6f, 0x74, 0x41, 0x67, 0x61, 0x69, 0x6e, 0x73, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x0b, 0x61,
	0x67, 0x61, 0x69, 0x6e, 0x73, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x61, 0x6e,
	0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x25, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x26, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x0e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x6f,
	0x18, 0x27, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3b, 0x0a, 0x11, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74,
	0x41, 0x67, 0x61, 0x69, 0x6e, 0x73, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x22, 0x51, 0x0a, 0x0f, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x41, 0x6e, 0x69,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x3e, 0x0a, 0x22, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x76, 0x65, 0x56,
	0x6f, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x23, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x76, 0x65, 0x56,
	0x6f, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x68, 0x61, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x68, 0x61, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x69, 0x67,
	0x42, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x69, 0x67, 0x42,
	0x61, 0x6c, 0x6c, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_foot_match_against_detail_proto_rawDescOnce sync.Once
	file_v2_foot_match_against_detail_proto_rawDescData = file_v2_foot_match_against_detail_proto_rawDesc
)

func file_v2_foot_match_against_detail_proto_rawDescGZIP() []byte {
	file_v2_foot_match_against_detail_proto_rawDescOnce.Do(func() {
		file_v2_foot_match_against_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_foot_match_against_detail_proto_rawDescData)
	})
	return file_v2_foot_match_against_detail_proto_rawDescData
}

var file_v2_foot_match_against_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v2_foot_match_against_detail_proto_goTypes = []interface{}{
	(*V2FootMatchAgainstDetailRequest)(nil),     // 0: V2FootMatchAgainstDetailRequest
	(*V2FootMatchAgainstDetailResponse)(nil),    // 1: V2FootMatchAgainstDetailResponse
	(*V2FootAgainstLive)(nil),                   // 2: V2FootAgainstLive
	(*V2FootAnimation)(nil),                     // 3: V2FootAnimation
	(*V2FootMatchPredictiveVotingRequest)(nil),  // 4: V2FootMatchPredictiveVotingRequest
	(*V2FootMatchPredictiveVotingResponse)(nil), // 5: V2FootMatchPredictiveVotingResponse
}
var file_v2_foot_match_against_detail_proto_depIdxs = []int32{
	2, // 0: V2FootMatchAgainstDetailResponse.againstLive:type_name -> V2FootAgainstLive
	3, // 1: V2FootMatchAgainstDetailResponse.animation:type_name -> V2FootAnimation
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v2_foot_match_against_detail_proto_init() }
func file_v2_foot_match_against_detail_proto_init() {
	if File_v2_foot_match_against_detail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_foot_match_against_detail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchAgainstDetailRequest); i {
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
		file_v2_foot_match_against_detail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchAgainstDetailResponse); i {
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
		file_v2_foot_match_against_detail_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootAgainstLive); i {
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
		file_v2_foot_match_against_detail_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootAnimation); i {
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
		file_v2_foot_match_against_detail_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchPredictiveVotingRequest); i {
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
		file_v2_foot_match_against_detail_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchPredictiveVotingResponse); i {
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
			RawDescriptor: file_v2_foot_match_against_detail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_foot_match_against_detail_proto_goTypes,
		DependencyIndexes: file_v2_foot_match_against_detail_proto_depIdxs,
		MessageInfos:      file_v2_foot_match_against_detail_proto_msgTypes,
	}.Build()
	File_v2_foot_match_against_detail_proto = out.File
	file_v2_foot_match_against_detail_proto_rawDesc = nil
	file_v2_foot_match_against_detail_proto_goTypes = nil
	file_v2_foot_match_against_detail_proto_depIdxs = nil
}
