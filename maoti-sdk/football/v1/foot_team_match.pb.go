// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: foot_team_match.proto

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

type FootTeamMatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId     int64  `protobuf:"varint,1,opt,name=teamId,proto3" json:"teamId,omitempty"`
	SearchDate string `protobuf:"bytes,2,opt,name=searchDate,proto3" json:"searchDate,omitempty"`
	Type       int64  `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Page       int64  `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	PageSize   int64  `protobuf:"varint,5,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Language   string `protobuf:"bytes,6,opt,name=language,proto3" json:"language,omitempty"` //请求语言  1:zh  2:en
}

func (x *FootTeamMatchRequest) Reset() {
	*x = FootTeamMatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_team_match_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootTeamMatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootTeamMatchRequest) ProtoMessage() {}

func (x *FootTeamMatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foot_team_match_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootTeamMatchRequest.ProtoReflect.Descriptor instead.
func (*FootTeamMatchRequest) Descriptor() ([]byte, []int) {
	return file_foot_team_match_proto_rawDescGZIP(), []int{0}
}

func (x *FootTeamMatchRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *FootTeamMatchRequest) GetSearchDate() string {
	if x != nil {
		return x.SearchDate
	}
	return ""
}

func (x *FootTeamMatchRequest) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *FootTeamMatchRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FootTeamMatchRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *FootTeamMatchRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type FootTeamMatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year      string           `protobuf:"bytes,1,opt,name=year,proto3" json:"year,omitempty"`
	MatchList []*TeamMatchList `protobuf:"bytes,2,rep,name=matchList,proto3" json:"matchList,omitempty"`
}

func (x *FootTeamMatchResponse) Reset() {
	*x = FootTeamMatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_team_match_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootTeamMatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootTeamMatchResponse) ProtoMessage() {}

func (x *FootTeamMatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_foot_team_match_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootTeamMatchResponse.ProtoReflect.Descriptor instead.
func (*FootTeamMatchResponse) Descriptor() ([]byte, []int) {
	return file_foot_team_match_proto_rawDescGZIP(), []int{1}
}

func (x *FootTeamMatchResponse) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *FootTeamMatchResponse) GetMatchList() []*TeamMatchList {
	if x != nil {
		return x.MatchList
	}
	return nil
}

type TeamMatchList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	YearTime      string `protobuf:"bytes,2,opt,name=yearTime,proto3" json:"yearTime,omitempty"`           //年份时间
	ChName        string `protobuf:"bytes,3,opt,name=chName,proto3" json:"chName,omitempty"`               //比赛名称
	Round         string `protobuf:"bytes,4,opt,name=round,proto3" json:"round,omitempty"`                 //场次
	TeamMainName  string `protobuf:"bytes,5,opt,name=teamMainName,proto3" json:"teamMainName,omitempty"`   //主队名称
	TeamCustName  string `protobuf:"bytes,6,opt,name=teamCustName,proto3" json:"teamCustName,omitempty"`   //客队名称
	MainCustScore string `protobuf:"bytes,7,opt,name=mainCustScore,proto3" json:"mainCustScore,omitempty"` //主客队分数
	IsCollect     bool   `protobuf:"varint,8,opt,name=isCollect,proto3" json:"isCollect,omitempty"`
	StartTime     int64  `protobuf:"varint,9,opt,name=startTime,proto3" json:"startTime,omitempty"`          //开始时间
	MainImages    string `protobuf:"bytes,10,opt,name=mainImages,proto3" json:"mainImages,omitempty"`        //主队图片
	CustImages    string `protobuf:"bytes,11,opt,name=custImages,proto3" json:"custImages,omitempty"`        //客队图片
	TeamMainScore int64  `protobuf:"varint,12,opt,name=teamMainScore,proto3" json:"teamMainScore,omitempty"` //主队比分
	TeamCustScore int64  `protobuf:"varint,13,opt,name=teamCustScore,proto3" json:"teamCustScore,omitempty"` //客队比分
	TeamMainId    int64  `protobuf:"varint,14,opt,name=teamMainId,proto3" json:"teamMainId,omitempty"`       //主队ID
	TeamCustId    int64  `protobuf:"varint,15,opt,name=teamCustId,proto3" json:"teamCustId,omitempty"`       //客队ID
	IsStart       int64  `protobuf:"varint,16,opt,name=isStart,proto3" json:"isStart,omitempty"`             //比赛状态
}

func (x *TeamMatchList) Reset() {
	*x = TeamMatchList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_team_match_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamMatchList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamMatchList) ProtoMessage() {}

func (x *TeamMatchList) ProtoReflect() protoreflect.Message {
	mi := &file_foot_team_match_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamMatchList.ProtoReflect.Descriptor instead.
func (*TeamMatchList) Descriptor() ([]byte, []int) {
	return file_foot_team_match_proto_rawDescGZIP(), []int{2}
}

func (x *TeamMatchList) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TeamMatchList) GetYearTime() string {
	if x != nil {
		return x.YearTime
	}
	return ""
}

func (x *TeamMatchList) GetChName() string {
	if x != nil {
		return x.ChName
	}
	return ""
}

func (x *TeamMatchList) GetRound() string {
	if x != nil {
		return x.Round
	}
	return ""
}

func (x *TeamMatchList) GetTeamMainName() string {
	if x != nil {
		return x.TeamMainName
	}
	return ""
}

func (x *TeamMatchList) GetTeamCustName() string {
	if x != nil {
		return x.TeamCustName
	}
	return ""
}

func (x *TeamMatchList) GetMainCustScore() string {
	if x != nil {
		return x.MainCustScore
	}
	return ""
}

func (x *TeamMatchList) GetIsCollect() bool {
	if x != nil {
		return x.IsCollect
	}
	return false
}

func (x *TeamMatchList) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *TeamMatchList) GetMainImages() string {
	if x != nil {
		return x.MainImages
	}
	return ""
}

func (x *TeamMatchList) GetCustImages() string {
	if x != nil {
		return x.CustImages
	}
	return ""
}

func (x *TeamMatchList) GetTeamMainScore() int64 {
	if x != nil {
		return x.TeamMainScore
	}
	return 0
}

func (x *TeamMatchList) GetTeamCustScore() int64 {
	if x != nil {
		return x.TeamCustScore
	}
	return 0
}

func (x *TeamMatchList) GetTeamMainId() int64 {
	if x != nil {
		return x.TeamMainId
	}
	return 0
}

func (x *TeamMatchList) GetTeamCustId() int64 {
	if x != nil {
		return x.TeamCustId
	}
	return 0
}

func (x *TeamMatchList) GetIsStart() int64 {
	if x != nil {
		return x.IsStart
	}
	return 0
}

var File_foot_team_match_proto protoreflect.FileDescriptor

var file_foot_team_match_proto_rawDesc = []byte{
	0x0a, 0x15, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x14, 0x46, 0x6f, 0x6f, 0x74,
	0x54, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x59, 0x0a, 0x15, 0x46, 0x6f, 0x6f, 0x74,
	0x54, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x2c, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0xf9, 0x03, 0x0a, 0x0d, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x79, 0x65, 0x61, 0x72, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x79, 0x65, 0x61, 0x72, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12,
	0x22, 0x0a, 0x0c, 0x74, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x61, 0x6d, 0x43,
	0x75, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x61, 0x69, 0x6e, 0x43,
	0x75, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6d, 0x61, 0x69, 0x6e, 0x43, 0x75, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x69, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x69,
	0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d,
	0x61, 0x69, 0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73,
	0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x65, 0x61,
	0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x74, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x74, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69,
	0x6e, 0x49, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x4d,
	0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73,
	0x74, 0x49, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x43,
	0x75, 0x73, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x73, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42,
	0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foot_team_match_proto_rawDescOnce sync.Once
	file_foot_team_match_proto_rawDescData = file_foot_team_match_proto_rawDesc
)

func file_foot_team_match_proto_rawDescGZIP() []byte {
	file_foot_team_match_proto_rawDescOnce.Do(func() {
		file_foot_team_match_proto_rawDescData = protoimpl.X.CompressGZIP(file_foot_team_match_proto_rawDescData)
	})
	return file_foot_team_match_proto_rawDescData
}

var file_foot_team_match_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_foot_team_match_proto_goTypes = []interface{}{
	(*FootTeamMatchRequest)(nil),  // 0: FootTeamMatchRequest
	(*FootTeamMatchResponse)(nil), // 1: FootTeamMatchResponse
	(*TeamMatchList)(nil),         // 2: TeamMatchList
}
var file_foot_team_match_proto_depIdxs = []int32{
	2, // 0: FootTeamMatchResponse.matchList:type_name -> TeamMatchList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_foot_team_match_proto_init() }
func file_foot_team_match_proto_init() {
	if File_foot_team_match_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foot_team_match_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootTeamMatchRequest); i {
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
		file_foot_team_match_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootTeamMatchResponse); i {
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
		file_foot_team_match_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamMatchList); i {
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
			RawDescriptor: file_foot_team_match_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foot_team_match_proto_goTypes,
		DependencyIndexes: file_foot_team_match_proto_depIdxs,
		MessageInfos:      file_foot_team_match_proto_msgTypes,
	}.Build()
	File_foot_team_match_proto = out.File
	file_foot_team_match_proto_rawDesc = nil
	file_foot_team_match_proto_goTypes = nil
	file_foot_team_match_proto_depIdxs = nil
}
