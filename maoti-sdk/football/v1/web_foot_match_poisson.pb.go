// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: web_foot_match_poisson.proto

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

//web比赛详情
type WebFootMatchPoissonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchId  int64  `protobuf:"varint,1,opt,name=MatchId,proto3" json:"MatchId,omitempty"`  //比赛id
	Language string `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"` //请求语言
}

func (x *WebFootMatchPoissonRequest) Reset() {
	*x = WebFootMatchPoissonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_match_poisson_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootMatchPoissonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootMatchPoissonRequest) ProtoMessage() {}

func (x *WebFootMatchPoissonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_match_poisson_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootMatchPoissonRequest.ProtoReflect.Descriptor instead.
func (*WebFootMatchPoissonRequest) Descriptor() ([]byte, []int) {
	return file_web_foot_match_poisson_proto_rawDescGZIP(), []int{0}
}

func (x *WebFootMatchPoissonRequest) GetMatchId() int64 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

func (x *WebFootMatchPoissonRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type WebFootMatchPoissonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RateForecast   *RateForecastResponse     `protobuf:"bytes,1,opt,name=RateForecast,proto3" json:"RateForecast,omitempty"`     //  概率预测
	ResultForecast []*ResultForecastResponse `protobuf:"bytes,2,rep,name=ResultForecast,proto3" json:"ResultForecast,omitempty"` //  赛果预测
	ExpectTrend    *ExpectTrendResponse      `protobuf:"bytes,3,opt,name=ExpectTrend,proto3" json:"ExpectTrend,omitempty"`       //  期望值
}

func (x *WebFootMatchPoissonResponse) Reset() {
	*x = WebFootMatchPoissonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_match_poisson_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootMatchPoissonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootMatchPoissonResponse) ProtoMessage() {}

func (x *WebFootMatchPoissonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_match_poisson_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootMatchPoissonResponse.ProtoReflect.Descriptor instead.
func (*WebFootMatchPoissonResponse) Descriptor() ([]byte, []int) {
	return file_web_foot_match_poisson_proto_rawDescGZIP(), []int{1}
}

func (x *WebFootMatchPoissonResponse) GetRateForecast() *RateForecastResponse {
	if x != nil {
		return x.RateForecast
	}
	return nil
}

func (x *WebFootMatchPoissonResponse) GetResultForecast() []*ResultForecastResponse {
	if x != nil {
		return x.ResultForecast
	}
	return nil
}

func (x *WebFootMatchPoissonResponse) GetExpectTrend() *ExpectTrendResponse {
	if x != nil {
		return x.ExpectTrend
	}
	return nil
}

// 概率预测
type RateForecastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HomeScore         string `protobuf:"bytes,1,opt,name=HomeScore,proto3" json:"HomeScore,omitempty"` // 预期进球
	AwayScore         string `protobuf:"bytes,2,opt,name=AwayScore,proto3" json:"AwayScore,omitempty"`
	HomeDataPlus      string `protobuf:"bytes,3,opt,name=HomeDataPlus,proto3" json:"HomeDataPlus,omitempty"` // 数据增强
	AwayDataPlus      string `protobuf:"bytes,4,opt,name=AwayDataPlus,proto3" json:"AwayDataPlus,omitempty"`
	HomeFormationRate string `protobuf:"bytes,5,opt,name=HomeFormationRate,proto3" json:"HomeFormationRate,omitempty"` // 预期阵容
	AwayFormationRate string `protobuf:"bytes,6,opt,name=AwayFormationRate,proto3" json:"AwayFormationRate,omitempty"`
	WinRate           string `protobuf:"bytes,7,opt,name=WinRate,proto3" json:"WinRate,omitempty"`   // 胜概率
	FlatRate          string `protobuf:"bytes,8,opt,name=FlatRate,proto3" json:"FlatRate,omitempty"` // 平概率
	LoseRate          string `protobuf:"bytes,9,opt,name=LoseRate,proto3" json:"LoseRate,omitempty"` // 负概率
}

func (x *RateForecastResponse) Reset() {
	*x = RateForecastResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_match_poisson_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateForecastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateForecastResponse) ProtoMessage() {}

func (x *RateForecastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_match_poisson_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateForecastResponse.ProtoReflect.Descriptor instead.
func (*RateForecastResponse) Descriptor() ([]byte, []int) {
	return file_web_foot_match_poisson_proto_rawDescGZIP(), []int{2}
}

func (x *RateForecastResponse) GetHomeScore() string {
	if x != nil {
		return x.HomeScore
	}
	return ""
}

func (x *RateForecastResponse) GetAwayScore() string {
	if x != nil {
		return x.AwayScore
	}
	return ""
}

func (x *RateForecastResponse) GetHomeDataPlus() string {
	if x != nil {
		return x.HomeDataPlus
	}
	return ""
}

func (x *RateForecastResponse) GetAwayDataPlus() string {
	if x != nil {
		return x.AwayDataPlus
	}
	return ""
}

func (x *RateForecastResponse) GetHomeFormationRate() string {
	if x != nil {
		return x.HomeFormationRate
	}
	return ""
}

func (x *RateForecastResponse) GetAwayFormationRate() string {
	if x != nil {
		return x.AwayFormationRate
	}
	return ""
}

func (x *RateForecastResponse) GetWinRate() string {
	if x != nil {
		return x.WinRate
	}
	return ""
}

func (x *RateForecastResponse) GetFlatRate() string {
	if x != nil {
		return x.FlatRate
	}
	return ""
}

func (x *RateForecastResponse) GetLoseRate() string {
	if x != nil {
		return x.LoseRate
	}
	return ""
}

// 赛果预测
type ResultForecastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Winner int64              `protobuf:"varint,1,opt,name=winner,proto3" json:"winner,omitempty"` // 1胜  3平  2负 4-总进球数
	Data   []*ScoreByProbItem `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ResultForecastResponse) Reset() {
	*x = ResultForecastResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_match_poisson_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultForecastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultForecastResponse) ProtoMessage() {}

func (x *ResultForecastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_match_poisson_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultForecastResponse.ProtoReflect.Descriptor instead.
func (*ResultForecastResponse) Descriptor() ([]byte, []int) {
	return file_web_foot_match_poisson_proto_rawDescGZIP(), []int{3}
}

func (x *ResultForecastResponse) GetWinner() int64 {
	if x != nil {
		return x.Winner
	}
	return 0
}

func (x *ResultForecastResponse) GetData() []*ScoreByProbItem {
	if x != nil {
		return x.Data
	}
	return nil
}

type ScoreByProbItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score       string  `protobuf:"bytes,1,opt,name=score,proto3" json:"score,omitempty"`
	Probability float32 `protobuf:"fixed32,2,opt,name=probability,proto3" json:"probability,omitempty"`
}

func (x *ScoreByProbItem) Reset() {
	*x = ScoreByProbItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_match_poisson_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScoreByProbItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScoreByProbItem) ProtoMessage() {}

func (x *ScoreByProbItem) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_match_poisson_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScoreByProbItem.ProtoReflect.Descriptor instead.
func (*ScoreByProbItem) Descriptor() ([]byte, []int) {
	return file_web_foot_match_poisson_proto_rawDescGZIP(), []int{4}
}

func (x *ScoreByProbItem) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

func (x *ScoreByProbItem) GetProbability() float32 {
	if x != nil {
		return x.Probability
	}
	return 0
}

// 期望值走势
type ExpectTrendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Home *ExpectTrendItem `protobuf:"bytes,1,opt,name=Home,proto3" json:"Home,omitempty"`
	Away *ExpectTrendItem `protobuf:"bytes,2,opt,name=away,proto3" json:"away,omitempty"`
}

func (x *ExpectTrendResponse) Reset() {
	*x = ExpectTrendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_match_poisson_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpectTrendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpectTrendResponse) ProtoMessage() {}

func (x *ExpectTrendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_match_poisson_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpectTrendResponse.ProtoReflect.Descriptor instead.
func (*ExpectTrendResponse) Descriptor() ([]byte, []int) {
	return file_web_foot_match_poisson_proto_rawDescGZIP(), []int{5}
}

func (x *ExpectTrendResponse) GetHome() *ExpectTrendItem {
	if x != nil {
		return x.Home
	}
	return nil
}

func (x *ExpectTrendResponse) GetAway() *ExpectTrendItem {
	if x != nil {
		return x.Away
	}
	return nil
}

type ExpectTrendItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Optimum     int64   `protobuf:"varint,1,opt,name=Optimum,proto3" json:"Optimum,omitempty"`         // 最佳
	Worst       int64   `protobuf:"varint,2,opt,name=Worst,proto3" json:"Worst,omitempty"`             // 最差
	ExpectCount int64   `protobuf:"varint,3,opt,name=ExpectCount,proto3" json:"ExpectCount,omitempty"` //期望值
	IsHit       []int64 `protobuf:"varint,4,rep,packed,name=IsHit,proto3" json:"IsHit,omitempty"`      // 中-1;不中-0
}

func (x *ExpectTrendItem) Reset() {
	*x = ExpectTrendItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_match_poisson_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpectTrendItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpectTrendItem) ProtoMessage() {}

func (x *ExpectTrendItem) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_match_poisson_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpectTrendItem.ProtoReflect.Descriptor instead.
func (*ExpectTrendItem) Descriptor() ([]byte, []int) {
	return file_web_foot_match_poisson_proto_rawDescGZIP(), []int{6}
}

func (x *ExpectTrendItem) GetOptimum() int64 {
	if x != nil {
		return x.Optimum
	}
	return 0
}

func (x *ExpectTrendItem) GetWorst() int64 {
	if x != nil {
		return x.Worst
	}
	return 0
}

func (x *ExpectTrendItem) GetExpectCount() int64 {
	if x != nil {
		return x.ExpectCount
	}
	return 0
}

func (x *ExpectTrendItem) GetIsHit() []int64 {
	if x != nil {
		return x.IsHit
	}
	return nil
}

var File_web_foot_match_poisson_proto protoreflect.FileDescriptor

var file_web_foot_match_poisson_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x77, 0x65, 0x62, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x5f, 0x70, 0x6f, 0x69, 0x73, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52,
	0x0a, 0x1a, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x6f,
	0x69, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x22, 0xd1, 0x01, 0x0a, 0x1b, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x50, 0x6f, 0x69, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x52, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x46,
	0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x0c, 0x52, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x3f, 0x0a,
	0x0e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x46, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x36,
	0x0a, 0x0b, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x54, 0x72, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0b, 0x45, 0x78, 0x70, 0x65, 0x63,
	0x74, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x22, 0xc8, 0x02, 0x0a, 0x14, 0x52, 0x61, 0x74, 0x65, 0x46,
	0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x41, 0x77, 0x61, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x41, 0x77, 0x61, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x48,
	0x6f, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6c, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x48, 0x6f, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6c, 0x75, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x41, 0x77, 0x61, 0x79, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6c, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x41, 0x77, 0x61, 0x79, 0x44, 0x61, 0x74, 0x61, 0x50,
	0x6c, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x48, 0x6f, 0x6d, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x48, 0x6f, 0x6d, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x2c, 0x0a, 0x11, 0x41, 0x77, 0x61, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x41, 0x77,
	0x61, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x57, 0x69, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x57, 0x69, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x6c, 0x61,
	0x74, 0x52, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x6c, 0x61,
	0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x73, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x6f, 0x73, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x22, 0x56, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63,
	0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77,
	0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x77, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x62, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x49, 0x0a, 0x0f, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x62, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x22, 0x61, 0x0a, 0x13, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x54, 0x72,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x48,
	0x6f, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x45, 0x78, 0x70, 0x65,
	0x63, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x48, 0x6f, 0x6d,
	0x65, 0x12, 0x24, 0x0a, 0x04, 0x61, 0x77, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x04, 0x61, 0x77, 0x61, 0x79, 0x22, 0x79, 0x0a, 0x0f, 0x45, 0x78, 0x70, 0x65, 0x63,
	0x74, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x70,
	0x74, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4f, 0x70, 0x74,
	0x69, 0x6d, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x6f, 0x72, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x57, 0x6f, 0x72, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x78,
	0x70, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x49, 0x73, 0x48, 0x69, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x49, 0x73, 0x48,
	0x69, 0x74, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_web_foot_match_poisson_proto_rawDescOnce sync.Once
	file_web_foot_match_poisson_proto_rawDescData = file_web_foot_match_poisson_proto_rawDesc
)

func file_web_foot_match_poisson_proto_rawDescGZIP() []byte {
	file_web_foot_match_poisson_proto_rawDescOnce.Do(func() {
		file_web_foot_match_poisson_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_foot_match_poisson_proto_rawDescData)
	})
	return file_web_foot_match_poisson_proto_rawDescData
}

var file_web_foot_match_poisson_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_web_foot_match_poisson_proto_goTypes = []interface{}{
	(*WebFootMatchPoissonRequest)(nil),  // 0: WebFootMatchPoissonRequest
	(*WebFootMatchPoissonResponse)(nil), // 1: WebFootMatchPoissonResponse
	(*RateForecastResponse)(nil),        // 2: RateForecastResponse
	(*ResultForecastResponse)(nil),      // 3: ResultForecastResponse
	(*ScoreByProbItem)(nil),             // 4: ScoreByProbItem
	(*ExpectTrendResponse)(nil),         // 5: ExpectTrendResponse
	(*ExpectTrendItem)(nil),             // 6: ExpectTrendItem
}
var file_web_foot_match_poisson_proto_depIdxs = []int32{
	2, // 0: WebFootMatchPoissonResponse.RateForecast:type_name -> RateForecastResponse
	3, // 1: WebFootMatchPoissonResponse.ResultForecast:type_name -> ResultForecastResponse
	5, // 2: WebFootMatchPoissonResponse.ExpectTrend:type_name -> ExpectTrendResponse
	4, // 3: ResultForecastResponse.data:type_name -> ScoreByProbItem
	6, // 4: ExpectTrendResponse.Home:type_name -> ExpectTrendItem
	6, // 5: ExpectTrendResponse.away:type_name -> ExpectTrendItem
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_web_foot_match_poisson_proto_init() }
func file_web_foot_match_poisson_proto_init() {
	if File_web_foot_match_poisson_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_foot_match_poisson_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootMatchPoissonRequest); i {
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
		file_web_foot_match_poisson_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootMatchPoissonResponse); i {
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
		file_web_foot_match_poisson_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateForecastResponse); i {
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
		file_web_foot_match_poisson_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultForecastResponse); i {
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
		file_web_foot_match_poisson_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScoreByProbItem); i {
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
		file_web_foot_match_poisson_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpectTrendResponse); i {
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
		file_web_foot_match_poisson_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpectTrendItem); i {
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
			RawDescriptor: file_web_foot_match_poisson_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_foot_match_poisson_proto_goTypes,
		DependencyIndexes: file_web_foot_match_poisson_proto_depIdxs,
		MessageInfos:      file_web_foot_match_poisson_proto_msgTypes,
	}.Build()
	File_web_foot_match_poisson_proto = out.File
	file_web_foot_match_poisson_proto_rawDesc = nil
	file_web_foot_match_poisson_proto_goTypes = nil
	file_web_foot_match_poisson_proto_depIdxs = nil
}
