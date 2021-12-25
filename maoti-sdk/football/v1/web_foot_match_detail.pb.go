// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: web_foot_match_detail.proto

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
type WebFootMatchDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchId  int64  `protobuf:"varint,1,opt,name=MatchId,proto3" json:"MatchId,omitempty"`  //比赛id
	Language string `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"` //请求语言  1:zh  2:en
}

func (x *WebFootMatchDetailRequest) Reset() {
	*x = WebFootMatchDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_match_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootMatchDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootMatchDetailRequest) ProtoMessage() {}

func (x *WebFootMatchDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_match_detail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootMatchDetailRequest.ProtoReflect.Descriptor instead.
func (*WebFootMatchDetailRequest) Descriptor() ([]byte, []int) {
	return file_web_foot_match_detail_proto_rawDescGZIP(), []int{0}
}

func (x *WebFootMatchDetailRequest) GetMatchId() int64 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

func (x *WebFootMatchDetailRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type WebFootMatchDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeagueId     int64  `protobuf:"varint,1,opt,name=LeagueId,proto3" json:"LeagueId,omitempty"`         //联赛ID
	LeagueName   string `protobuf:"bytes,2,opt,name=LeagueName,proto3" json:"LeagueName,omitempty"`      //联赛名称
	RoundName    string `protobuf:"bytes,3,opt,name=roundName,proto3" json:"roundName,omitempty"`        //比赛轮次
	MatchTime    string `protobuf:"bytes,4,opt,name=matchTime,proto3" json:"matchTime,omitempty"`        //比赛时间
	MainTeamId   int64  `protobuf:"varint,5,opt,name=mainTeamId,proto3" json:"mainTeamId,omitempty"`     //主队ID
	MainTeamName string `protobuf:"bytes,6,opt,name=mainTeamName,proto3" json:"mainTeamName,omitempty"`  //主队名称
	MainTeamLogo string `protobuf:"bytes,7,opt,name=mainTeamLogo,proto3" json:"mainTeamLogo,omitempty"`  //主队Logo
	MainScore    string `protobuf:"bytes,8,opt,name=mainScore,proto3" json:"mainScore,omitempty"`        //主队得分
	CustTeamId   int64  `protobuf:"varint,9,opt,name=custTeamId,proto3" json:"custTeamId,omitempty"`     //客队ID
	CustTeamName string `protobuf:"bytes,10,opt,name=custTeamName,proto3" json:"custTeamName,omitempty"` //客队名称
	CustTeamLogo string `protobuf:"bytes,11,opt,name=custTeamLogo,proto3" json:"custTeamLogo,omitempty"` //客队Logo
	CustScore    string `protobuf:"bytes,12,opt,name=custScore,proto3" json:"custScore,omitempty"`       //客队得分
	CourtId      string `protobuf:"bytes,13,opt,name=courtId,proto3" json:"courtId,omitempty"`           //球场ID
	CourtName    string `protobuf:"bytes,14,opt,name=courtName,proto3" json:"courtName,omitempty"`       //球场名称
	Weather      string `protobuf:"bytes,15,opt,name=weather,proto3" json:"weather,omitempty"`           //天气
}

func (x *WebFootMatchDetailResponse) Reset() {
	*x = WebFootMatchDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_match_detail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootMatchDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootMatchDetailResponse) ProtoMessage() {}

func (x *WebFootMatchDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_match_detail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootMatchDetailResponse.ProtoReflect.Descriptor instead.
func (*WebFootMatchDetailResponse) Descriptor() ([]byte, []int) {
	return file_web_foot_match_detail_proto_rawDescGZIP(), []int{1}
}

func (x *WebFootMatchDetailResponse) GetLeagueId() int64 {
	if x != nil {
		return x.LeagueId
	}
	return 0
}

func (x *WebFootMatchDetailResponse) GetLeagueName() string {
	if x != nil {
		return x.LeagueName
	}
	return ""
}

func (x *WebFootMatchDetailResponse) GetRoundName() string {
	if x != nil {
		return x.RoundName
	}
	return ""
}

func (x *WebFootMatchDetailResponse) GetMatchTime() string {
	if x != nil {
		return x.MatchTime
	}
	return ""
}

func (x *WebFootMatchDetailResponse) GetMainTeamId() int64 {
	if x != nil {
		return x.MainTeamId
	}
	return 0
}

func (x *WebFootMatchDetailResponse) GetMainTeamName() string {
	if x != nil {
		return x.MainTeamName
	}
	return ""
}

func (x *WebFootMatchDetailResponse) GetMainTeamLogo() string {
	if x != nil {
		return x.MainTeamLogo
	}
	return ""
}

func (x *WebFootMatchDetailResponse) GetMainScore() string {
	if x != nil {
		return x.MainScore
	}
	return ""
}

func (x *WebFootMatchDetailResponse) GetCustTeamId() int64 {
	if x != nil {
		return x.CustTeamId
	}
	return 0
}

func (x *WebFootMatchDetailResponse) GetCustTeamName() string {
	if x != nil {
		return x.CustTeamName
	}
	return ""
}

func (x *WebFootMatchDetailResponse) GetCustTeamLogo() string {
	if x != nil {
		return x.CustTeamLogo
	}
	return ""
}

func (x *WebFootMatchDetailResponse) GetCustScore() string {
	if x != nil {
		return x.CustScore
	}
	return ""
}

func (x *WebFootMatchDetailResponse) GetCourtId() string {
	if x != nil {
		return x.CourtId
	}
	return ""
}

func (x *WebFootMatchDetailResponse) GetCourtName() string {
	if x != nil {
		return x.CourtName
	}
	return ""
}

func (x *WebFootMatchDetailResponse) GetWeather() string {
	if x != nil {
		return x.Weather
	}
	return ""
}

var File_web_foot_match_detail_proto protoreflect.FileDescriptor

var file_web_foot_match_detail_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x77, 0x65, 0x62, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a,
	0x19, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x22, 0xf2, 0x03, 0x0a, 0x1a, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x4c,
	0x65, 0x61, 0x67, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x54,
	0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x61, 0x69,
	0x6e, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x69, 0x6e, 0x54,
	0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d,
	0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d,
	0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x6f, 0x12,
	0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x63, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67,
	0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61,
	0x6d, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x75, 0x73, 0x74, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x75, 0x73, 0x74, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x74, 0x49, 0x64, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x75, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x77,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web_foot_match_detail_proto_rawDescOnce sync.Once
	file_web_foot_match_detail_proto_rawDescData = file_web_foot_match_detail_proto_rawDesc
)

func file_web_foot_match_detail_proto_rawDescGZIP() []byte {
	file_web_foot_match_detail_proto_rawDescOnce.Do(func() {
		file_web_foot_match_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_foot_match_detail_proto_rawDescData)
	})
	return file_web_foot_match_detail_proto_rawDescData
}

var file_web_foot_match_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_web_foot_match_detail_proto_goTypes = []interface{}{
	(*WebFootMatchDetailRequest)(nil),  // 0: WebFootMatchDetailRequest
	(*WebFootMatchDetailResponse)(nil), // 1: WebFootMatchDetailResponse
}
var file_web_foot_match_detail_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_web_foot_match_detail_proto_init() }
func file_web_foot_match_detail_proto_init() {
	if File_web_foot_match_detail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_foot_match_detail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootMatchDetailRequest); i {
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
		file_web_foot_match_detail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootMatchDetailResponse); i {
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
			RawDescriptor: file_web_foot_match_detail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_foot_match_detail_proto_goTypes,
		DependencyIndexes: file_web_foot_match_detail_proto_depIdxs,
		MessageInfos:      file_web_foot_match_detail_proto_msgTypes,
	}.Build()
	File_web_foot_match_detail_proto = out.File
	file_web_foot_match_detail_proto_rawDesc = nil
	file_web_foot_match_detail_proto_goTypes = nil
	file_web_foot_match_detail_proto_depIdxs = nil
}