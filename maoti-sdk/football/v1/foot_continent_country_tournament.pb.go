// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: foot_continent_country_tournament.proto

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
type FootContinentCountryTournamentInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"` //请求语言
	SportId  string `protobuf:"bytes,1,opt,name=sportId,proto3" json:"sportId,omitempty"`   //1足球 2篮球
}

func (x *FootContinentCountryTournamentInfoRequest) Reset() {
	*x = FootContinentCountryTournamentInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_continent_country_tournament_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootContinentCountryTournamentInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootContinentCountryTournamentInfoRequest) ProtoMessage() {}

func (x *FootContinentCountryTournamentInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foot_continent_country_tournament_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootContinentCountryTournamentInfoRequest.ProtoReflect.Descriptor instead.
func (*FootContinentCountryTournamentInfoRequest) Descriptor() ([]byte, []int) {
	return file_foot_continent_country_tournament_proto_rawDescGZIP(), []int{0}
}

func (x *FootContinentCountryTournamentInfoRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *FootContinentCountryTournamentInfoRequest) GetSportId() string {
	if x != nil {
		return x.SportId
	}
	return ""
}

type FootContinentCountryTournamentInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*FootContinentCountryTournament `protobuf:"bytes,25,rep,name=Data,proto3" json:"Data,omitempty"` //数据列表
}

func (x *FootContinentCountryTournamentInfoResponse) Reset() {
	*x = FootContinentCountryTournamentInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_continent_country_tournament_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootContinentCountryTournamentInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootContinentCountryTournamentInfoResponse) ProtoMessage() {}

func (x *FootContinentCountryTournamentInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_foot_continent_country_tournament_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootContinentCountryTournamentInfoResponse.ProtoReflect.Descriptor instead.
func (*FootContinentCountryTournamentInfoResponse) Descriptor() ([]byte, []int) {
	return file_foot_continent_country_tournament_proto_rawDescGZIP(), []int{1}
}

func (x *FootContinentCountryTournamentInfoResponse) GetData() []*FootContinentCountryTournament {
	if x != nil {
		return x.Data
	}
	return nil
}

type FootContinentCountryTournament struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64           `protobuf:"varint,3,opt,name=Id,proto3" json:"Id,omitempty"`              //地区ID
	TabName   string          `protobuf:"bytes,4,opt,name=TabName,proto3" json:"TabName,omitempty"`     //选项卡名称
	AreaImage string          `protobuf:"bytes,5,opt,name=AreaImage,proto3" json:"AreaImage,omitempty"` //地区的图片
	Sort      int64           `protobuf:"varint,6,opt,name=Sort,proto3" json:"Sort,omitempty"`          //排序(忽略掉，返回数据就是已经排好了的)
	AreaList  []*FootAreaList `protobuf:"bytes,7,rep,name=AreaList,proto3" json:"AreaList,omitempty"`   //国家列表
}

func (x *FootContinentCountryTournament) Reset() {
	*x = FootContinentCountryTournament{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_continent_country_tournament_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootContinentCountryTournament) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootContinentCountryTournament) ProtoMessage() {}

func (x *FootContinentCountryTournament) ProtoReflect() protoreflect.Message {
	mi := &file_foot_continent_country_tournament_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootContinentCountryTournament.ProtoReflect.Descriptor instead.
func (*FootContinentCountryTournament) Descriptor() ([]byte, []int) {
	return file_foot_continent_country_tournament_proto_rawDescGZIP(), []int{2}
}

func (x *FootContinentCountryTournament) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FootContinentCountryTournament) GetTabName() string {
	if x != nil {
		return x.TabName
	}
	return ""
}

func (x *FootContinentCountryTournament) GetAreaImage() string {
	if x != nil {
		return x.AreaImage
	}
	return ""
}

func (x *FootContinentCountryTournament) GetSort() int64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *FootContinentCountryTournament) GetAreaList() []*FootAreaList {
	if x != nil {
		return x.AreaList
	}
	return nil
}

type FootAreaList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AreaName  string           `protobuf:"bytes,8,opt,name=AreaName,proto3" json:"AreaName,omitempty"`    //地区名称
	Count     int64            `protobuf:"varint,9,opt,name=Count,proto3" json:"Count,omitempty"`         //赛事数量
	Id        int64            `protobuf:"varint,10,opt,name=Id,proto3" json:"Id,omitempty"`              //国家id
	PId       int64            `protobuf:"varint,11,opt,name=PId,proto3" json:"PId,omitempty"`            //地区父级ID
	AreaImage string           `protobuf:"bytes,12,opt,name=AreaImage,proto3" json:"AreaImage,omitempty"` //地区的图片
	AreaSort  int64            `protobuf:"varint,7,opt,name=AreaSort,proto3" json:"AreaSort,omitempty"`   //设置的国家排序
	EventList []*FootEventList `protobuf:"bytes,13,rep,name=EventList,proto3" json:"EventList,omitempty"` //地区下的赛事数据
}

func (x *FootAreaList) Reset() {
	*x = FootAreaList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_continent_country_tournament_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootAreaList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootAreaList) ProtoMessage() {}

func (x *FootAreaList) ProtoReflect() protoreflect.Message {
	mi := &file_foot_continent_country_tournament_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootAreaList.ProtoReflect.Descriptor instead.
func (*FootAreaList) Descriptor() ([]byte, []int) {
	return file_foot_continent_country_tournament_proto_rawDescGZIP(), []int{3}
}

func (x *FootAreaList) GetAreaName() string {
	if x != nil {
		return x.AreaName
	}
	return ""
}

func (x *FootAreaList) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *FootAreaList) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FootAreaList) GetPId() int64 {
	if x != nil {
		return x.PId
	}
	return 0
}

func (x *FootAreaList) GetAreaImage() string {
	if x != nil {
		return x.AreaImage
	}
	return ""
}

func (x *FootAreaList) GetAreaSort() int64 {
	if x != nil {
		return x.AreaSort
	}
	return 0
}

func (x *FootAreaList) GetEventList() []*FootEventList {
	if x != nil {
		return x.EventList
	}
	return nil
}

type FootEventList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComName           string `protobuf:"bytes,14,opt,name=ComName,proto3" json:"ComName,omitempty"`                      //联赛名称
	TournamentStageId int64  `protobuf:"varint,15,opt,name=TournamentStageId,proto3" json:"TournamentStageId,omitempty"` //联赛ID
	TournamentIcon    string `protobuf:"bytes,16,opt,name=TournamentIcon,proto3" json:"TournamentIcon,omitempty"`        //联赛图标
	IsHot             bool   `protobuf:"varint,17,opt,name=IsHot,proto3" json:"IsHot,omitempty"`                         //是否热门赛事 （true:是，false:否）目前只有显示当前设置的热门和用户登录之后添加的赛事
	Sort              int64  `protobuf:"varint,18,opt,name=Sort,proto3" json:"Sort,omitempty"`                           //排序(忽略掉，返回数据就是已经排好了的)
	IsSpecial         int64  `protobuf:"varint,19,opt,name=IsSpecial,proto3" json:"IsSpecial,omitempty"`                 //0：正常；1：淘汰赛；2：特殊赛；
	IsChecked         bool   `protobuf:"varint,20,opt,name=IsChecked,proto3" json:"IsChecked,omitempty"`                 //是否选中 true:是，false:否
	IsKind            int64  `protobuf:"varint,25,opt,name=IsKind,proto3" json:"IsKind,omitempty"`                       //1:联赛，2：杯赛
	AreaId            int64  `protobuf:"varint,21,opt,name=AreaId,proto3" json:"AreaId,omitempty"`                       //国家id
	AreaPId           int64  `protobuf:"varint,22,opt,name=AreaPId,proto3" json:"AreaPId,omitempty"`                     //国家的父ai
	Type              int64  `protobuf:"varint,23,opt,name=Type,proto3" json:"Type,omitempty"`                           //运动类型
	LeagueType        int64  `protobuf:"varint,24,opt,name=LeagueType,proto3" json:"LeagueType,omitempty"`               //语言类型
}

func (x *FootEventList) Reset() {
	*x = FootEventList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_continent_country_tournament_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootEventList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootEventList) ProtoMessage() {}

func (x *FootEventList) ProtoReflect() protoreflect.Message {
	mi := &file_foot_continent_country_tournament_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootEventList.ProtoReflect.Descriptor instead.
func (*FootEventList) Descriptor() ([]byte, []int) {
	return file_foot_continent_country_tournament_proto_rawDescGZIP(), []int{4}
}

func (x *FootEventList) GetComName() string {
	if x != nil {
		return x.ComName
	}
	return ""
}

func (x *FootEventList) GetTournamentStageId() int64 {
	if x != nil {
		return x.TournamentStageId
	}
	return 0
}

func (x *FootEventList) GetTournamentIcon() string {
	if x != nil {
		return x.TournamentIcon
	}
	return ""
}

func (x *FootEventList) GetIsHot() bool {
	if x != nil {
		return x.IsHot
	}
	return false
}

func (x *FootEventList) GetSort() int64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *FootEventList) GetIsSpecial() int64 {
	if x != nil {
		return x.IsSpecial
	}
	return 0
}

func (x *FootEventList) GetIsChecked() bool {
	if x != nil {
		return x.IsChecked
	}
	return false
}

func (x *FootEventList) GetIsKind() int64 {
	if x != nil {
		return x.IsKind
	}
	return 0
}

func (x *FootEventList) GetAreaId() int64 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *FootEventList) GetAreaPId() int64 {
	if x != nil {
		return x.AreaPId
	}
	return 0
}

func (x *FootEventList) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *FootEventList) GetLeagueType() int64 {
	if x != nil {
		return x.LeagueType
	}
	return 0
}

var File_foot_continent_country_tournament_proto protoreflect.FileDescriptor

var file_foot_continent_country_tournament_proto_rawDesc = []byte{
	0x0a, 0x27, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x29, 0x46, 0x6f, 0x6f,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x22, 0x61, 0x0a, 0x2a,
	0x46, 0x6f, 0x6f, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x19, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x54,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22,
	0xa7, 0x01, 0x0a, 0x1e, 0x46, 0x6f, 0x6f, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x61, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x61, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x41, 0x72, 0x65, 0x61, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x41, 0x72, 0x65, 0x61, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x6f,
	0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x29,
	0x0a, 0x08, 0x41, 0x72, 0x65, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x41, 0x72, 0x65, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x08, 0x41, 0x72, 0x65, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xca, 0x01, 0x0a, 0x0c, 0x46, 0x6f,
	0x6f, 0x74, 0x41, 0x72, 0x65, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x72,
	0x65, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x72,
	0x65, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x50, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x50, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x41, 0x72, 0x65, 0x61, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x41, 0x72, 0x65, 0x61, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x41, 0x72, 0x65, 0x61, 0x53, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x41, 0x72, 0x65, 0x61, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x46, 0x6f,
	0x6f, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x09, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xe3, 0x02, 0x0a, 0x0d, 0x46, 0x6f, 0x6f, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x54,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x12, 0x26, 0x0a, 0x0e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x63,
	0x6f, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x73, 0x48, 0x6f,
	0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x49, 0x73, 0x48, 0x6f, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x53, 0x6f,
	0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x49, 0x73, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c,
	0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x49, 0x73, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x49, 0x73, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x72, 0x65, 0x61, 0x49, 0x64,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x41, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x41, 0x72, 0x65, 0x61, 0x50, 0x49, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x41, 0x72, 0x65, 0x61, 0x50, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foot_continent_country_tournament_proto_rawDescOnce sync.Once
	file_foot_continent_country_tournament_proto_rawDescData = file_foot_continent_country_tournament_proto_rawDesc
)

func file_foot_continent_country_tournament_proto_rawDescGZIP() []byte {
	file_foot_continent_country_tournament_proto_rawDescOnce.Do(func() {
		file_foot_continent_country_tournament_proto_rawDescData = protoimpl.X.CompressGZIP(file_foot_continent_country_tournament_proto_rawDescData)
	})
	return file_foot_continent_country_tournament_proto_rawDescData
}

var file_foot_continent_country_tournament_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_foot_continent_country_tournament_proto_goTypes = []interface{}{
	(*FootContinentCountryTournamentInfoRequest)(nil),  // 0: FootContinentCountryTournamentInfoRequest
	(*FootContinentCountryTournamentInfoResponse)(nil), // 1: FootContinentCountryTournamentInfoResponse
	(*FootContinentCountryTournament)(nil),             // 2: FootContinentCountryTournament
	(*FootAreaList)(nil),                               // 3: FootAreaList
	(*FootEventList)(nil),                              // 4: FootEventList
}
var file_foot_continent_country_tournament_proto_depIdxs = []int32{
	2, // 0: FootContinentCountryTournamentInfoResponse.Data:type_name -> FootContinentCountryTournament
	3, // 1: FootContinentCountryTournament.AreaList:type_name -> FootAreaList
	4, // 2: FootAreaList.EventList:type_name -> FootEventList
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_foot_continent_country_tournament_proto_init() }
func file_foot_continent_country_tournament_proto_init() {
	if File_foot_continent_country_tournament_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foot_continent_country_tournament_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootContinentCountryTournamentInfoRequest); i {
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
		file_foot_continent_country_tournament_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootContinentCountryTournamentInfoResponse); i {
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
		file_foot_continent_country_tournament_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootContinentCountryTournament); i {
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
		file_foot_continent_country_tournament_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootAreaList); i {
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
		file_foot_continent_country_tournament_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootEventList); i {
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
			RawDescriptor: file_foot_continent_country_tournament_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foot_continent_country_tournament_proto_goTypes,
		DependencyIndexes: file_foot_continent_country_tournament_proto_depIdxs,
		MessageInfos:      file_foot_continent_country_tournament_proto_msgTypes,
	}.Build()
	File_foot_continent_country_tournament_proto = out.File
	file_foot_continent_country_tournament_proto_rawDesc = nil
	file_foot_continent_country_tournament_proto_goTypes = nil
	file_foot_continent_country_tournament_proto_depIdxs = nil
}
