// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: web_foot_match_list.proto

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

//联赛的回合赛程
type WebFootMatchListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TournamentId int64  `protobuf:"varint,1,opt,name=tournamentId,proto3" json:"tournamentId,omitempty"` //联赛id
	SeasonId     int64  `protobuf:"varint,2,opt,name=seasonId,proto3" json:"seasonId,omitempty"`         //赛季id
	Language     string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`          //请求语言
	CptTurns     string `protobuf:"bytes,4,opt,name=CptTurns,proto3" json:"CptTurns,omitempty"`          //指定round
	TimeZone     int64  `protobuf:"varint,6,opt,name=TimeZone,proto3" json:"TimeZone,omitempty"`         //时区
}

func (x *WebFootMatchListRequest) Reset() {
	*x = WebFootMatchListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_match_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootMatchListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootMatchListRequest) ProtoMessage() {}

func (x *WebFootMatchListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_match_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootMatchListRequest.ProtoReflect.Descriptor instead.
func (*WebFootMatchListRequest) Descriptor() ([]byte, []int) {
	return file_web_foot_match_list_proto_rawDescGZIP(), []int{0}
}

func (x *WebFootMatchListRequest) GetTournamentId() int64 {
	if x != nil {
		return x.TournamentId
	}
	return 0
}

func (x *WebFootMatchListRequest) GetSeasonId() int64 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

func (x *WebFootMatchListRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *WebFootMatchListRequest) GetCptTurns() string {
	if x != nil {
		return x.CptTurns
	}
	return ""
}

func (x *WebFootMatchListRequest) GetTimeZone() int64 {
	if x != nil {
		return x.TimeZone
	}
	return 0
}

type WebFootMatchListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*WebMatchList `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *WebFootMatchListResponse) Reset() {
	*x = WebFootMatchListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_match_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootMatchListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootMatchListResponse) ProtoMessage() {}

func (x *WebFootMatchListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_match_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootMatchListResponse.ProtoReflect.Descriptor instead.
func (*WebFootMatchListResponse) Descriptor() ([]byte, []int) {
	return file_web_foot_match_list_proto_rawDescGZIP(), []int{1}
}

func (x *WebFootMatchListResponse) GetList() []*WebMatchList {
	if x != nil {
		return x.List
	}
	return nil
}

type WebMatchList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                        //   比赛id
	MatchTime     int64  `protobuf:"varint,2,opt,name=matchTime,proto3" json:"matchTime,omitempty"`          //   日期
	MainTeamName  string `protobuf:"bytes,3,opt,name=mainTeamName,proto3" json:"mainTeamName,omitempty"`     //   主队名称
	CustTeamName  string `protobuf:"bytes,4,opt,name=custTeamName,proto3" json:"custTeamName,omitempty"`     //   客队名称
	TeamMainScore int64  `protobuf:"varint,5,opt,name=teamMainScore,proto3" json:"teamMainScore,omitempty"`  //   主队得分分数
	TeamCustScore int64  `protobuf:"varint,6,opt,name=teamCustScore,proto3" json:"teamCustScore,omitempty"`  //   客队得分分数
	MainTeamId    int64  `protobuf:"varint,8,opt,name=mainTeamId,proto3" json:"mainTeamId,omitempty"`        //   主队id
	CustTeamId    int64  `protobuf:"varint,9,opt,name=custTeamId,proto3" json:"custTeamId,omitempty"`        //   客队id
	MainImage     string `protobuf:"bytes,10,opt,name=mainImage,proto3" json:"mainImage,omitempty"`          //   主队图片
	CustImage     string `protobuf:"bytes,11,opt,name=custImage,proto3" json:"custImage,omitempty"`          //   客队图片
	RoundType     string `protobuf:"bytes,13,opt,name=roundType,proto3" json:"roundType,omitempty"`          //   供应商比赛轮次类型
	MainHalfScore int64  `protobuf:"varint,14,opt,name=mainHalfScore,proto3" json:"mainHalfScore,omitempty"` //主队半场比分
	CustHalfScore int64  `protobuf:"varint,15,opt,name=custHalfScore,proto3" json:"custHalfScore,omitempty"` //客队半场比分
	MatchStatus   string `protobuf:"bytes,16,opt,name=MatchStatus,proto3" json:"MatchStatus,omitempty"`      //比赛状态
}

func (x *WebMatchList) Reset() {
	*x = WebMatchList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_match_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebMatchList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebMatchList) ProtoMessage() {}

func (x *WebMatchList) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_match_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebMatchList.ProtoReflect.Descriptor instead.
func (*WebMatchList) Descriptor() ([]byte, []int) {
	return file_web_foot_match_list_proto_rawDescGZIP(), []int{2}
}

func (x *WebMatchList) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WebMatchList) GetMatchTime() int64 {
	if x != nil {
		return x.MatchTime
	}
	return 0
}

func (x *WebMatchList) GetMainTeamName() string {
	if x != nil {
		return x.MainTeamName
	}
	return ""
}

func (x *WebMatchList) GetCustTeamName() string {
	if x != nil {
		return x.CustTeamName
	}
	return ""
}

func (x *WebMatchList) GetTeamMainScore() int64 {
	if x != nil {
		return x.TeamMainScore
	}
	return 0
}

func (x *WebMatchList) GetTeamCustScore() int64 {
	if x != nil {
		return x.TeamCustScore
	}
	return 0
}

func (x *WebMatchList) GetMainTeamId() int64 {
	if x != nil {
		return x.MainTeamId
	}
	return 0
}

func (x *WebMatchList) GetCustTeamId() int64 {
	if x != nil {
		return x.CustTeamId
	}
	return 0
}

func (x *WebMatchList) GetMainImage() string {
	if x != nil {
		return x.MainImage
	}
	return ""
}

func (x *WebMatchList) GetCustImage() string {
	if x != nil {
		return x.CustImage
	}
	return ""
}

func (x *WebMatchList) GetRoundType() string {
	if x != nil {
		return x.RoundType
	}
	return ""
}

func (x *WebMatchList) GetMainHalfScore() int64 {
	if x != nil {
		return x.MainHalfScore
	}
	return 0
}

func (x *WebMatchList) GetCustHalfScore() int64 {
	if x != nil {
		return x.CustHalfScore
	}
	return 0
}

func (x *WebMatchList) GetMatchStatus() string {
	if x != nil {
		return x.MatchStatus
	}
	return ""
}

var File_web_foot_match_list_proto protoreflect.FileDescriptor

var file_web_foot_match_list_proto_rawDesc = []byte{
	0x0a, 0x19, 0x77, 0x65, 0x62, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x17,
	0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x70, 0x74, 0x54, 0x75, 0x72, 0x6e, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x70, 0x74, 0x54, 0x75, 0x72, 0x6e, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x3d, 0x0a, 0x18, 0x57,
	0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x57, 0x65, 0x62, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xd8, 0x03, 0x0a, 0x0c, 0x57,
	0x65, 0x62, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x69,
	0x6e, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x63, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x65, 0x61, 0x6d, 0x4d, 0x61,
	0x69, 0x6e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x65, 0x61, 0x6d, 0x43,
	0x75, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x74, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x75, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x75, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x61, 0x69, 0x6e, 0x48,
	0x61, 0x6c, 0x66, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x6d, 0x61, 0x69, 0x6e, 0x48, 0x61, 0x6c, 0x66, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x63, 0x75, 0x73, 0x74, 0x48, 0x61, 0x6c, 0x66, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x48, 0x61, 0x6c, 0x66, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web_foot_match_list_proto_rawDescOnce sync.Once
	file_web_foot_match_list_proto_rawDescData = file_web_foot_match_list_proto_rawDesc
)

func file_web_foot_match_list_proto_rawDescGZIP() []byte {
	file_web_foot_match_list_proto_rawDescOnce.Do(func() {
		file_web_foot_match_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_foot_match_list_proto_rawDescData)
	})
	return file_web_foot_match_list_proto_rawDescData
}

var file_web_foot_match_list_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_web_foot_match_list_proto_goTypes = []interface{}{
	(*WebFootMatchListRequest)(nil),  // 0: WebFootMatchListRequest
	(*WebFootMatchListResponse)(nil), // 1: WebFootMatchListResponse
	(*WebMatchList)(nil),             // 2: WebMatchList
}
var file_web_foot_match_list_proto_depIdxs = []int32{
	2, // 0: WebFootMatchListResponse.list:type_name -> WebMatchList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_web_foot_match_list_proto_init() }
func file_web_foot_match_list_proto_init() {
	if File_web_foot_match_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_foot_match_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootMatchListRequest); i {
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
		file_web_foot_match_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootMatchListResponse); i {
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
		file_web_foot_match_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebMatchList); i {
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
			RawDescriptor: file_web_foot_match_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_foot_match_list_proto_goTypes,
		DependencyIndexes: file_web_foot_match_list_proto_depIdxs,
		MessageInfos:      file_web_foot_match_list_proto_msgTypes,
	}.Build()
	File_web_foot_match_list_proto = out.File
	file_web_foot_match_list_proto_rawDesc = nil
	file_web_foot_match_list_proto_goTypes = nil
	file_web_foot_match_list_proto_depIdxs = nil
}
