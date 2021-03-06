// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: v2_foot_player_information.proto

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

// 详细信息和转会记录  （接口）
type V2FootPlayerInfoPersonalDataResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DateOfBirth       string `protobuf:"bytes,1,opt,name=DateOfBirth,proto3" json:"DateOfBirth,omitempty"`             //出生年月日
	CountryId         int64  `protobuf:"varint,2,opt,name=CountryId,proto3" json:"CountryId,omitempty"`                //球员国家id
	Positions         string `protobuf:"bytes,3,opt,name=positions,proto3" json:"positions,omitempty"`                 // 位置
	PreferredFoot     string `protobuf:"bytes,4,opt,name=PreferredFoot,proto3" json:"PreferredFoot,omitempty"`         // 惯用脚
	CurrentTeamId     int64  `protobuf:"varint,5,opt,name=CurrentTeamId,proto3" json:"CurrentTeamId,omitempty"`        // 现在球队id
	TournamentId      int64  `protobuf:"varint,6,opt,name=TournamentId,proto3" json:"TournamentId,omitempty"`          // 所属联赛id
	PreviousTeamId    int64  `protobuf:"varint,7,opt,name=PreviousTeamId,proto3" json:"PreviousTeamId,omitempty"`      // 上个球队id
	ClothesNum        int64  `protobuf:"varint,8,opt,name=ClothesNum,proto3" json:"ClothesNum,omitempty"`              //现在球衣号
	HistoryClothesNum string `protobuf:"bytes,9,opt,name=HistoryClothesNum,proto3" json:"HistoryClothesNum,omitempty"` // 历史球衣号
	CareerYears       int64  `protobuf:"varint,10,opt,name=CareerYears,proto3" json:"CareerYears,omitempty"`           // 职业生涯年份
}

func (x *V2FootPlayerInfoPersonalDataResp) Reset() {
	*x = V2FootPlayerInfoPersonalDataResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_player_information_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootPlayerInfoPersonalDataResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootPlayerInfoPersonalDataResp) ProtoMessage() {}

func (x *V2FootPlayerInfoPersonalDataResp) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_player_information_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootPlayerInfoPersonalDataResp.ProtoReflect.Descriptor instead.
func (*V2FootPlayerInfoPersonalDataResp) Descriptor() ([]byte, []int) {
	return file_v2_foot_player_information_proto_rawDescGZIP(), []int{0}
}

func (x *V2FootPlayerInfoPersonalDataResp) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *V2FootPlayerInfoPersonalDataResp) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *V2FootPlayerInfoPersonalDataResp) GetPositions() string {
	if x != nil {
		return x.Positions
	}
	return ""
}

func (x *V2FootPlayerInfoPersonalDataResp) GetPreferredFoot() string {
	if x != nil {
		return x.PreferredFoot
	}
	return ""
}

func (x *V2FootPlayerInfoPersonalDataResp) GetCurrentTeamId() int64 {
	if x != nil {
		return x.CurrentTeamId
	}
	return 0
}

func (x *V2FootPlayerInfoPersonalDataResp) GetTournamentId() int64 {
	if x != nil {
		return x.TournamentId
	}
	return 0
}

func (x *V2FootPlayerInfoPersonalDataResp) GetPreviousTeamId() int64 {
	if x != nil {
		return x.PreviousTeamId
	}
	return 0
}

func (x *V2FootPlayerInfoPersonalDataResp) GetClothesNum() int64 {
	if x != nil {
		return x.ClothesNum
	}
	return 0
}

func (x *V2FootPlayerInfoPersonalDataResp) GetHistoryClothesNum() string {
	if x != nil {
		return x.HistoryClothesNum
	}
	return ""
}

func (x *V2FootPlayerInfoPersonalDataResp) GetCareerYears() int64 {
	if x != nil {
		return x.CareerYears
	}
	return 0
}

// 冠军数据 （接口）
type V2FootPlayerInfoChampionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Champion []*V2ChampionItem `protobuf:"bytes,1,rep,name=Champion,proto3" json:"Champion,omitempty"`
}

func (x *V2FootPlayerInfoChampionResp) Reset() {
	*x = V2FootPlayerInfoChampionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_player_information_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootPlayerInfoChampionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootPlayerInfoChampionResp) ProtoMessage() {}

func (x *V2FootPlayerInfoChampionResp) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_player_information_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootPlayerInfoChampionResp.ProtoReflect.Descriptor instead.
func (*V2FootPlayerInfoChampionResp) Descriptor() ([]byte, []int) {
	return file_v2_foot_player_information_proto_rawDescGZIP(), []int{1}
}

func (x *V2FootPlayerInfoChampionResp) GetChampion() []*V2ChampionItem {
	if x != nil {
		return x.Champion
	}
	return nil
}

// 冠军数据子项
type V2ChampionItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TournamentId        int64    `protobuf:"varint,1,opt,name=TournamentId,proto3" json:"TournamentId,omitempty"`              // 联赛id
	ChampionSeasonDates []string `protobuf:"bytes,2,rep,name=ChampionSeasonDates,proto3" json:"ChampionSeasonDates,omitempty"` // 冠军赛季日期集
	ChampionCount       int64    `protobuf:"varint,3,opt,name=ChampionCount,proto3" json:"ChampionCount,omitempty"`            // 次数
}

func (x *V2ChampionItem) Reset() {
	*x = V2ChampionItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_player_information_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2ChampionItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2ChampionItem) ProtoMessage() {}

func (x *V2ChampionItem) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_player_information_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2ChampionItem.ProtoReflect.Descriptor instead.
func (*V2ChampionItem) Descriptor() ([]byte, []int) {
	return file_v2_foot_player_information_proto_rawDescGZIP(), []int{2}
}

func (x *V2ChampionItem) GetTournamentId() int64 {
	if x != nil {
		return x.TournamentId
	}
	return 0
}

func (x *V2ChampionItem) GetChampionSeasonDates() []string {
	if x != nil {
		return x.ChampionSeasonDates
	}
	return nil
}

func (x *V2ChampionItem) GetChampionCount() int64 {
	if x != nil {
		return x.ChampionCount
	}
	return 0
}

var File_v2_foot_player_information_proto protoreflect.FileDescriptor

var file_v2_foot_player_information_proto_rawDesc = []byte{
	0x0a, 0x20, 0x76, 0x32, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x88, 0x03, 0x0a, 0x20, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x65, 0x4f,
	0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x61,
	0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x65, 0x64, 0x46, 0x6f, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x50, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x46, 0x6f, 0x6f, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75,
	0x73, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x50,
	0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x43, 0x6c, 0x6f, 0x74, 0x68, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x43, 0x6c, 0x6f, 0x74, 0x68, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x2c, 0x0a,
	0x11, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6c, 0x6f, 0x74, 0x68, 0x65, 0x73, 0x4e,
	0x75, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x43, 0x6c, 0x6f, 0x74, 0x68, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x43,
	0x61, 0x72, 0x65, 0x65, 0x72, 0x59, 0x65, 0x61, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x43, 0x61, 0x72, 0x65, 0x65, 0x72, 0x59, 0x65, 0x61, 0x72, 0x73, 0x22, 0x4b, 0x0a,
	0x1c, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2b, 0x0a,
	0x08, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x56, 0x32, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x08, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x22, 0x8c, 0x01, 0x0a, 0x0e, 0x56,
	0x32, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x22, 0x0a,
	0x0c, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x30, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13,
	0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x43, 0x68, 0x61, 0x6d,
	0x70, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_foot_player_information_proto_rawDescOnce sync.Once
	file_v2_foot_player_information_proto_rawDescData = file_v2_foot_player_information_proto_rawDesc
)

func file_v2_foot_player_information_proto_rawDescGZIP() []byte {
	file_v2_foot_player_information_proto_rawDescOnce.Do(func() {
		file_v2_foot_player_information_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_foot_player_information_proto_rawDescData)
	})
	return file_v2_foot_player_information_proto_rawDescData
}

var file_v2_foot_player_information_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v2_foot_player_information_proto_goTypes = []interface{}{
	(*V2FootPlayerInfoPersonalDataResp)(nil), // 0: V2FootPlayerInfoPersonalDataResp
	(*V2FootPlayerInfoChampionResp)(nil),     // 1: V2FootPlayerInfoChampionResp
	(*V2ChampionItem)(nil),                   // 2: V2ChampionItem
}
var file_v2_foot_player_information_proto_depIdxs = []int32{
	2, // 0: V2FootPlayerInfoChampionResp.Champion:type_name -> V2ChampionItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v2_foot_player_information_proto_init() }
func file_v2_foot_player_information_proto_init() {
	if File_v2_foot_player_information_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_foot_player_information_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootPlayerInfoPersonalDataResp); i {
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
		file_v2_foot_player_information_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootPlayerInfoChampionResp); i {
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
		file_v2_foot_player_information_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2ChampionItem); i {
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
			RawDescriptor: file_v2_foot_player_information_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_foot_player_information_proto_goTypes,
		DependencyIndexes: file_v2_foot_player_information_proto_depIdxs,
		MessageInfos:      file_v2_foot_player_information_proto_msgTypes,
	}.Build()
	File_v2_foot_player_information_proto = out.File
	file_v2_foot_player_information_proto_rawDesc = nil
	file_v2_foot_player_information_proto_goTypes = nil
	file_v2_foot_player_information_proto_depIdxs = nil
}
