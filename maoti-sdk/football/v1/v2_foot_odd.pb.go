// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: v2_foot_odd.proto

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

//v2 即时比赛列表 - 比赛详情 - 指数列表 - 欧赔（1x2）列表
type V2FootMatchEuropeanOddsListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId  int64  `protobuf:"varint,1,opt,name=eventId,proto3" json:"eventId,omitempty"`   //比赛id
	Language string `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`  //请求语言
	TimeZone int64  `protobuf:"varint,3,opt,name=timeZone,proto3" json:"timeZone,omitempty"` //时区
	Odds     int64  `protobuf:"varint,4,opt,name=odds,proto3" json:"odds,omitempty"`         //是否是赔率页面使用；0.否；1.是
}

func (x *V2FootMatchEuropeanOddsListRequest) Reset() {
	*x = V2FootMatchEuropeanOddsListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_odd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchEuropeanOddsListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchEuropeanOddsListRequest) ProtoMessage() {}

func (x *V2FootMatchEuropeanOddsListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_odd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchEuropeanOddsListRequest.ProtoReflect.Descriptor instead.
func (*V2FootMatchEuropeanOddsListRequest) Descriptor() ([]byte, []int) {
	return file_v2_foot_odd_proto_rawDescGZIP(), []int{0}
}

func (x *V2FootMatchEuropeanOddsListRequest) GetEventId() int64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *V2FootMatchEuropeanOddsListRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *V2FootMatchEuropeanOddsListRequest) GetTimeZone() int64 {
	if x != nil {
		return x.TimeZone
	}
	return 0
}

func (x *V2FootMatchEuropeanOddsListRequest) GetOdds() int64 {
	if x != nil {
		return x.Odds
	}
	return 0
}

type V2FootMatchEuropeanOddsListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Maximum *V2OddsDetails     `protobuf:"bytes,1,opt,name=maximum,proto3" json:"maximum,omitempty"`
	Minimum *V2OddsDetails     `protobuf:"bytes,2,opt,name=minimum,proto3" json:"minimum,omitempty"`
	Average *V2OddsDetails     `protobuf:"bytes,4,opt,name=Average,proto3" json:"Average,omitempty"`
	List    []*V2OddDetailList `protobuf:"bytes,6,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *V2FootMatchEuropeanOddsListResponse) Reset() {
	*x = V2FootMatchEuropeanOddsListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_odd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchEuropeanOddsListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchEuropeanOddsListResponse) ProtoMessage() {}

func (x *V2FootMatchEuropeanOddsListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_odd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchEuropeanOddsListResponse.ProtoReflect.Descriptor instead.
func (*V2FootMatchEuropeanOddsListResponse) Descriptor() ([]byte, []int) {
	return file_v2_foot_odd_proto_rawDescGZIP(), []int{1}
}

func (x *V2FootMatchEuropeanOddsListResponse) GetMaximum() *V2OddsDetails {
	if x != nil {
		return x.Maximum
	}
	return nil
}

func (x *V2FootMatchEuropeanOddsListResponse) GetMinimum() *V2OddsDetails {
	if x != nil {
		return x.Minimum
	}
	return nil
}

func (x *V2FootMatchEuropeanOddsListResponse) GetAverage() *V2OddsDetails {
	if x != nil {
		return x.Average
	}
	return nil
}

func (x *V2FootMatchEuropeanOddsListResponse) GetList() []*V2OddDetailList {
	if x != nil {
		return x.List
	}
	return nil
}

type V2OddDetailList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Init        *V2OddsDetails `protobuf:"bytes,1,opt,name=init,proto3" json:"init,omitempty"`
	Curr        *V2OddsDetails `protobuf:"bytes,2,opt,name=curr,proto3" json:"curr,omitempty"`
	CompanyName string         `protobuf:"bytes,3,opt,name=companyName,proto3" json:"companyName,omitempty"` //公司名称
	CompanyId   int64          `protobuf:"varint,4,opt,name=companyId,proto3" json:"companyId,omitempty"`    //公司id
	Mark        int64          `protobuf:"varint,5,opt,name=mark,proto3" json:"mark,omitempty"`              //最早和最晚：1.最早；2.最晚；0.无。指数列表使用
}

func (x *V2OddDetailList) Reset() {
	*x = V2OddDetailList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_odd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2OddDetailList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2OddDetailList) ProtoMessage() {}

func (x *V2OddDetailList) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_odd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2OddDetailList.ProtoReflect.Descriptor instead.
func (*V2OddDetailList) Descriptor() ([]byte, []int) {
	return file_v2_foot_odd_proto_rawDescGZIP(), []int{2}
}

func (x *V2OddDetailList) GetInit() *V2OddsDetails {
	if x != nil {
		return x.Init
	}
	return nil
}

func (x *V2OddDetailList) GetCurr() *V2OddsDetails {
	if x != nil {
		return x.Curr
	}
	return nil
}

func (x *V2OddDetailList) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *V2OddDetailList) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *V2OddDetailList) GetMark() int64 {
	if x != nil {
		return x.Mark
	}
	return 0
}

type V2OddsDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Odds       []string `protobuf:"bytes,3,rep,name=odds,proto3" json:"odds,omitempty"`
	ReturnProb float32  `protobuf:"fixed32,4,opt,name=returnProb,proto3" json:"returnProb,omitempty"` //返还率
	Variety    int64    `protobuf:"varint,5,opt,name=variety,proto3" json:"variety,omitempty"`        //时间
}

func (x *V2OddsDetails) Reset() {
	*x = V2OddsDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_odd_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2OddsDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2OddsDetails) ProtoMessage() {}

func (x *V2OddsDetails) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_odd_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2OddsDetails.ProtoReflect.Descriptor instead.
func (*V2OddsDetails) Descriptor() ([]byte, []int) {
	return file_v2_foot_odd_proto_rawDescGZIP(), []int{3}
}

func (x *V2OddsDetails) GetOdds() []string {
	if x != nil {
		return x.Odds
	}
	return nil
}

func (x *V2OddsDetails) GetReturnProb() float32 {
	if x != nil {
		return x.ReturnProb
	}
	return 0
}

func (x *V2OddsDetails) GetVariety() int64 {
	if x != nil {
		return x.Variety
	}
	return 0
}

//v2 即时比赛列表 - 比赛详情 - 指数列表 - 亚赔（HDP）和大小球（O/U）列表
type V2FootMatchAsiaOddsAndBigBallListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OddType  int64  `protobuf:"varint,1,opt,name=oddType,proto3" json:"oddType,omitempty"`   //odd  1压盘 2大小球
	EventId  int64  `protobuf:"varint,2,opt,name=eventId,proto3" json:"eventId,omitempty"`   //比赛id
	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`  //请求语言
	TimeZone int64  `protobuf:"varint,6,opt,name=timeZone,proto3" json:"timeZone,omitempty"` //时区
	Odds     int64  `protobuf:"varint,7,opt,name=odds,proto3" json:"odds,omitempty"`         //是否是赔率页面使用；0.否；1.是
}

func (x *V2FootMatchAsiaOddsAndBigBallListRequest) Reset() {
	*x = V2FootMatchAsiaOddsAndBigBallListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_odd_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchAsiaOddsAndBigBallListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchAsiaOddsAndBigBallListRequest) ProtoMessage() {}

func (x *V2FootMatchAsiaOddsAndBigBallListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_odd_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchAsiaOddsAndBigBallListRequest.ProtoReflect.Descriptor instead.
func (*V2FootMatchAsiaOddsAndBigBallListRequest) Descriptor() ([]byte, []int) {
	return file_v2_foot_odd_proto_rawDescGZIP(), []int{4}
}

func (x *V2FootMatchAsiaOddsAndBigBallListRequest) GetOddType() int64 {
	if x != nil {
		return x.OddType
	}
	return 0
}

func (x *V2FootMatchAsiaOddsAndBigBallListRequest) GetEventId() int64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *V2FootMatchAsiaOddsAndBigBallListRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *V2FootMatchAsiaOddsAndBigBallListRequest) GetTimeZone() int64 {
	if x != nil {
		return x.TimeZone
	}
	return 0
}

func (x *V2FootMatchAsiaOddsAndBigBallListRequest) GetOdds() int64 {
	if x != nil {
		return x.Odds
	}
	return 0
}

type V2FootMatchAsiaOddsAndBigBallListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info []*V2FootMatchAsiaOddsAndBigBallInfo `protobuf:"bytes,3,rep,name=info,proto3" json:"info,omitempty"`
}

func (x *V2FootMatchAsiaOddsAndBigBallListResponse) Reset() {
	*x = V2FootMatchAsiaOddsAndBigBallListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_odd_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchAsiaOddsAndBigBallListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchAsiaOddsAndBigBallListResponse) ProtoMessage() {}

func (x *V2FootMatchAsiaOddsAndBigBallListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_odd_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchAsiaOddsAndBigBallListResponse.ProtoReflect.Descriptor instead.
func (*V2FootMatchAsiaOddsAndBigBallListResponse) Descriptor() ([]byte, []int) {
	return file_v2_foot_odd_proto_rawDescGZIP(), []int{5}
}

func (x *V2FootMatchAsiaOddsAndBigBallListResponse) GetInfo() []*V2FootMatchAsiaOddsAndBigBallInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type V2FootMatchAsiaOddsAndBigBallInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId   int64    `protobuf:"varint,1,opt,name=companyId,proto3" json:"companyId,omitempty"`    //公司id
	CompanyName string   `protobuf:"bytes,2,opt,name=companyName,proto3" json:"companyName,omitempty"` //公司名称
	InitOdds    []string `protobuf:"bytes,3,rep,name=initOdds,proto3" json:"initOdds,omitempty"`
	CurrOdds    []string `protobuf:"bytes,4,rep,name=currOdds,proto3" json:"currOdds,omitempty"`
	Times       int64    `protobuf:"varint,5,opt,name=times,proto3" json:"times,omitempty"`
	Mark        int64    `protobuf:"varint,6,opt,name=mark,proto3" json:"mark,omitempty"` //最早和最晚：1.最早；2.最晚；0.无。指数列表使用
}

func (x *V2FootMatchAsiaOddsAndBigBallInfo) Reset() {
	*x = V2FootMatchAsiaOddsAndBigBallInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_odd_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchAsiaOddsAndBigBallInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchAsiaOddsAndBigBallInfo) ProtoMessage() {}

func (x *V2FootMatchAsiaOddsAndBigBallInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_odd_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchAsiaOddsAndBigBallInfo.ProtoReflect.Descriptor instead.
func (*V2FootMatchAsiaOddsAndBigBallInfo) Descriptor() ([]byte, []int) {
	return file_v2_foot_odd_proto_rawDescGZIP(), []int{6}
}

func (x *V2FootMatchAsiaOddsAndBigBallInfo) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *V2FootMatchAsiaOddsAndBigBallInfo) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *V2FootMatchAsiaOddsAndBigBallInfo) GetInitOdds() []string {
	if x != nil {
		return x.InitOdds
	}
	return nil
}

func (x *V2FootMatchAsiaOddsAndBigBallInfo) GetCurrOdds() []string {
	if x != nil {
		return x.CurrOdds
	}
	return nil
}

func (x *V2FootMatchAsiaOddsAndBigBallInfo) GetTimes() int64 {
	if x != nil {
		return x.Times
	}
	return 0
}

func (x *V2FootMatchAsiaOddsAndBigBallInfo) GetMark() int64 {
	if x != nil {
		return x.Mark
	}
	return 0
}

var File_v2_foot_odd_proto protoreflect.FileDescriptor

var file_v2_foot_odd_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x32, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6f, 0x64, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x22, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x61, 0x6e, 0x4f, 0x64, 0x64, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6f, 0x64, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6f, 0x64, 0x64, 0x73,
	0x22, 0xc9, 0x01, 0x0a, 0x23, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x61, 0x6e, 0x4f, 0x64, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x69,
	0x6d, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x56, 0x32, 0x4f, 0x64,
	0x64, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d,
	0x75, 0x6d, 0x12, 0x28, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x56, 0x32, 0x4f, 0x64, 0x64, 0x73, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x28, 0x0a, 0x07,
	0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x56, 0x32, 0x4f, 0x64, 0x64, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x41,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x56, 0x32, 0x4f, 0x64, 0x64, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xad, 0x01, 0x0a,
	0x0f, 0x56, 0x32, 0x4f, 0x64, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x04, 0x69, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x56, 0x32, 0x4f, 0x64, 0x64, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x04,
	0x69, 0x6e, 0x69, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x63, 0x75, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x56, 0x32, 0x4f, 0x64, 0x64, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x04, 0x63, 0x75, 0x72, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x5d, 0x0a, 0x0d,
	0x56, 0x32, 0x4f, 0x64, 0x64, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6f, 0x64, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x64, 0x64,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x50, 0x72, 0x6f, 0x62, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x50, 0x72, 0x6f,
	0x62, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x22, 0xaa, 0x01, 0x0a, 0x28,
	0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x41, 0x73, 0x69, 0x61, 0x4f,
	0x64, 0x64, 0x73, 0x41, 0x6e, 0x64, 0x42, 0x69, 0x67, 0x42, 0x61, 0x6c, 0x6c, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x64, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x64, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65,
	0x5a, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65,
	0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x64, 0x64, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x6f, 0x64, 0x64, 0x73, 0x22, 0x63, 0x0a, 0x29, 0x56, 0x32, 0x46, 0x6f,
	0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x41, 0x73, 0x69, 0x61, 0x4f, 0x64, 0x64, 0x73, 0x41,
	0x6e, 0x64, 0x42, 0x69, 0x67, 0x42, 0x61, 0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x41, 0x73, 0x69, 0x61, 0x4f, 0x64, 0x64, 0x73, 0x41, 0x6e, 0x64, 0x42, 0x69, 0x67, 0x42,
	0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xc5, 0x01,
	0x0a, 0x21, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x41, 0x73, 0x69,
	0x61, 0x4f, 0x64, 0x64, 0x73, 0x41, 0x6e, 0x64, 0x42, 0x69, 0x67, 0x42, 0x61, 0x6c, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x69, 0x74, 0x4f, 0x64, 0x64, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x69, 0x74, 0x4f, 0x64, 0x64, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x4f, 0x64, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x4f, 0x64, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x6d, 0x61, 0x72, 0x6b, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_foot_odd_proto_rawDescOnce sync.Once
	file_v2_foot_odd_proto_rawDescData = file_v2_foot_odd_proto_rawDesc
)

func file_v2_foot_odd_proto_rawDescGZIP() []byte {
	file_v2_foot_odd_proto_rawDescOnce.Do(func() {
		file_v2_foot_odd_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_foot_odd_proto_rawDescData)
	})
	return file_v2_foot_odd_proto_rawDescData
}

var file_v2_foot_odd_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v2_foot_odd_proto_goTypes = []interface{}{
	(*V2FootMatchEuropeanOddsListRequest)(nil),        // 0: V2FootMatchEuropeanOddsListRequest
	(*V2FootMatchEuropeanOddsListResponse)(nil),       // 1: V2FootMatchEuropeanOddsListResponse
	(*V2OddDetailList)(nil),                           // 2: V2OddDetailList
	(*V2OddsDetails)(nil),                             // 3: V2OddsDetails
	(*V2FootMatchAsiaOddsAndBigBallListRequest)(nil),  // 4: V2FootMatchAsiaOddsAndBigBallListRequest
	(*V2FootMatchAsiaOddsAndBigBallListResponse)(nil), // 5: V2FootMatchAsiaOddsAndBigBallListResponse
	(*V2FootMatchAsiaOddsAndBigBallInfo)(nil),         // 6: V2FootMatchAsiaOddsAndBigBallInfo
}
var file_v2_foot_odd_proto_depIdxs = []int32{
	3, // 0: V2FootMatchEuropeanOddsListResponse.maximum:type_name -> V2OddsDetails
	3, // 1: V2FootMatchEuropeanOddsListResponse.minimum:type_name -> V2OddsDetails
	3, // 2: V2FootMatchEuropeanOddsListResponse.Average:type_name -> V2OddsDetails
	2, // 3: V2FootMatchEuropeanOddsListResponse.list:type_name -> V2OddDetailList
	3, // 4: V2OddDetailList.init:type_name -> V2OddsDetails
	3, // 5: V2OddDetailList.curr:type_name -> V2OddsDetails
	6, // 6: V2FootMatchAsiaOddsAndBigBallListResponse.info:type_name -> V2FootMatchAsiaOddsAndBigBallInfo
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_v2_foot_odd_proto_init() }
func file_v2_foot_odd_proto_init() {
	if File_v2_foot_odd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_foot_odd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchEuropeanOddsListRequest); i {
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
		file_v2_foot_odd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchEuropeanOddsListResponse); i {
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
		file_v2_foot_odd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2OddDetailList); i {
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
		file_v2_foot_odd_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2OddsDetails); i {
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
		file_v2_foot_odd_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchAsiaOddsAndBigBallListRequest); i {
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
		file_v2_foot_odd_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchAsiaOddsAndBigBallListResponse); i {
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
		file_v2_foot_odd_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchAsiaOddsAndBigBallInfo); i {
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
			RawDescriptor: file_v2_foot_odd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_foot_odd_proto_goTypes,
		DependencyIndexes: file_v2_foot_odd_proto_depIdxs,
		MessageInfos:      file_v2_foot_odd_proto_msgTypes,
	}.Build()
	File_v2_foot_odd_proto = out.File
	file_v2_foot_odd_proto_rawDesc = nil
	file_v2_foot_odd_proto_goTypes = nil
	file_v2_foot_odd_proto_depIdxs = nil
}
