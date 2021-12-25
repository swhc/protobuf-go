// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: web_foot_league_detail.proto

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

//web联赛详情
type WebFootLeagueDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeagueId int64  `protobuf:"varint,1,opt,name=LeagueId,proto3" json:"LeagueId,omitempty"` //联赛id
	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`  //请求语言  1:zh  2:en
}

func (x *WebFootLeagueDetailRequest) Reset() {
	*x = WebFootLeagueDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_league_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootLeagueDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootLeagueDetailRequest) ProtoMessage() {}

func (x *WebFootLeagueDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_league_detail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootLeagueDetailRequest.ProtoReflect.Descriptor instead.
func (*WebFootLeagueDetailRequest) Descriptor() ([]byte, []int) {
	return file_web_foot_league_detail_proto_rawDescGZIP(), []int{0}
}

func (x *WebFootLeagueDetailRequest) GetLeagueId() int64 {
	if x != nil {
		return x.LeagueId
	}
	return 0
}

func (x *WebFootLeagueDetailRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type WebFootLeagueDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeagueId    int64                      `protobuf:"varint,1,opt,name=LeagueId,proto3" json:"LeagueId,omitempty"`      //联赛id
	LeagueName  string                     `protobuf:"bytes,3,opt,name=LeagueName,proto3" json:"LeagueName,omitempty"`   //联赛名称
	Logo        string                     `protobuf:"bytes,4,opt,name=Logo,proto3" json:"Logo,omitempty"`               //联赛图片
	CountryId   int64                      `protobuf:"varint,5,opt,name=CountryId,proto3" json:"CountryId,omitempty"`    //联赛国家ID
	CountryName string                     `protobuf:"bytes,8,opt,name=CountryName,proto3" json:"CountryName,omitempty"` //联赛国家
	CountryLogo string                     `protobuf:"bytes,6,opt,name=CountryLogo,proto3" json:"CountryLogo,omitempty"` //联赛国家图片
	Season      []*WebFootLeagueSeasonList `protobuf:"bytes,7,rep,name=season,proto3" json:"season,omitempty"`           //联赛下的赛季列表
}

func (x *WebFootLeagueDetailResponse) Reset() {
	*x = WebFootLeagueDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_league_detail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootLeagueDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootLeagueDetailResponse) ProtoMessage() {}

func (x *WebFootLeagueDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_league_detail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootLeagueDetailResponse.ProtoReflect.Descriptor instead.
func (*WebFootLeagueDetailResponse) Descriptor() ([]byte, []int) {
	return file_web_foot_league_detail_proto_rawDescGZIP(), []int{1}
}

func (x *WebFootLeagueDetailResponse) GetLeagueId() int64 {
	if x != nil {
		return x.LeagueId
	}
	return 0
}

func (x *WebFootLeagueDetailResponse) GetLeagueName() string {
	if x != nil {
		return x.LeagueName
	}
	return ""
}

func (x *WebFootLeagueDetailResponse) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *WebFootLeagueDetailResponse) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *WebFootLeagueDetailResponse) GetCountryName() string {
	if x != nil {
		return x.CountryName
	}
	return ""
}

func (x *WebFootLeagueDetailResponse) GetCountryLogo() string {
	if x != nil {
		return x.CountryLogo
	}
	return ""
}

func (x *WebFootLeagueDetailResponse) GetSeason() []*WebFootLeagueSeasonList {
	if x != nil {
		return x.Season
	}
	return nil
}

type WebFootLeagueSeasonList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeasonId   int64  `protobuf:"varint,1,opt,name=SeasonId,proto3" json:"SeasonId,omitempty"`    //赛季id
	SeasonName string `protobuf:"bytes,2,opt,name=SeasonName,proto3" json:"SeasonName,omitempty"` //赛季名称
}

func (x *WebFootLeagueSeasonList) Reset() {
	*x = WebFootLeagueSeasonList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_league_detail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootLeagueSeasonList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootLeagueSeasonList) ProtoMessage() {}

func (x *WebFootLeagueSeasonList) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_league_detail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootLeagueSeasonList.ProtoReflect.Descriptor instead.
func (*WebFootLeagueSeasonList) Descriptor() ([]byte, []int) {
	return file_web_foot_league_detail_proto_rawDescGZIP(), []int{2}
}

func (x *WebFootLeagueSeasonList) GetSeasonId() int64 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

func (x *WebFootLeagueSeasonList) GetSeasonName() string {
	if x != nil {
		return x.SeasonName
	}
	return ""
}

var File_web_foot_league_detail_proto protoreflect.FileDescriptor

var file_web_foot_league_detail_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x77, 0x65, 0x62, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6c, 0x65, 0x61, 0x67, 0x75,
	0x65, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54,
	0x0a, 0x1a, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x22, 0x81, 0x02, 0x0a, 0x1b, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74,
	0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4c,
	0x6f, 0x67, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74,
	0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x06, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x55, 0x0a, 0x17, 0x57, 0x65, 0x62, 0x46,
	0x6f, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x42,
	0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web_foot_league_detail_proto_rawDescOnce sync.Once
	file_web_foot_league_detail_proto_rawDescData = file_web_foot_league_detail_proto_rawDesc
)

func file_web_foot_league_detail_proto_rawDescGZIP() []byte {
	file_web_foot_league_detail_proto_rawDescOnce.Do(func() {
		file_web_foot_league_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_foot_league_detail_proto_rawDescData)
	})
	return file_web_foot_league_detail_proto_rawDescData
}

var file_web_foot_league_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_web_foot_league_detail_proto_goTypes = []interface{}{
	(*WebFootLeagueDetailRequest)(nil),  // 0: WebFootLeagueDetailRequest
	(*WebFootLeagueDetailResponse)(nil), // 1: WebFootLeagueDetailResponse
	(*WebFootLeagueSeasonList)(nil),     // 2: WebFootLeagueSeasonList
}
var file_web_foot_league_detail_proto_depIdxs = []int32{
	2, // 0: WebFootLeagueDetailResponse.season:type_name -> WebFootLeagueSeasonList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_web_foot_league_detail_proto_init() }
func file_web_foot_league_detail_proto_init() {
	if File_web_foot_league_detail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_foot_league_detail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootLeagueDetailRequest); i {
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
		file_web_foot_league_detail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootLeagueDetailResponse); i {
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
		file_web_foot_league_detail_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootLeagueSeasonList); i {
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
			RawDescriptor: file_web_foot_league_detail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_foot_league_detail_proto_goTypes,
		DependencyIndexes: file_web_foot_league_detail_proto_depIdxs,
		MessageInfos:      file_web_foot_league_detail_proto_msgTypes,
	}.Build()
	File_web_foot_league_detail_proto = out.File
	file_web_foot_league_detail_proto_rawDesc = nil
	file_web_foot_league_detail_proto_goTypes = nil
	file_web_foot_league_detail_proto_depIdxs = nil
}
