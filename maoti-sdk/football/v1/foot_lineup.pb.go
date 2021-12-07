// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: foot_lineup.proto

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

//指数-列表
type FootLineupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SportId  string `protobuf:"bytes,1,opt,name=sportId,proto3" json:"sportId,omitempty"`   //1足球 2篮球
	TeamId   int64  `protobuf:"varint,2,opt,name=teamId,proto3" json:"teamId,omitempty"`    //球队id
	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"` //请求语言  1:zh  2:en
}

func (x *FootLineupRequest) Reset() {
	*x = FootLineupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_lineup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootLineupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootLineupRequest) ProtoMessage() {}

func (x *FootLineupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foot_lineup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootLineupRequest.ProtoReflect.Descriptor instead.
func (*FootLineupRequest) Descriptor() ([]byte, []int) {
	return file_foot_lineup_proto_rawDescGZIP(), []int{0}
}

func (x *FootLineupRequest) GetSportId() string {
	if x != nil {
		return x.SportId
	}
	return ""
}

func (x *FootLineupRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *FootLineupRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type FootLineupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Back       []*FootLineupBack `protobuf:"bytes,1,rep,name=back,proto3" json:"back,omitempty"`             //后卫
	Coach      *Coach            `protobuf:"bytes,2,opt,name=coach,proto3" json:"coach,omitempty"`           //教练
	Goalkeeper []*FootLineupBack `protobuf:"bytes,3,rep,name=goalkeeper,proto3" json:"goalkeeper,omitempty"` //守门员
	Midfield   []*FootLineupBack `protobuf:"bytes,4,rep,name=midfield,proto3" json:"midfield,omitempty"`     //中场
	Vanguard   []*FootLineupBack `protobuf:"bytes,5,rep,name=vanguard,proto3" json:"vanguard,omitempty"`     //前锋
}

func (x *FootLineupResponse) Reset() {
	*x = FootLineupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_lineup_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootLineupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootLineupResponse) ProtoMessage() {}

func (x *FootLineupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_foot_lineup_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootLineupResponse.ProtoReflect.Descriptor instead.
func (*FootLineupResponse) Descriptor() ([]byte, []int) {
	return file_foot_lineup_proto_rawDescGZIP(), []int{1}
}

func (x *FootLineupResponse) GetBack() []*FootLineupBack {
	if x != nil {
		return x.Back
	}
	return nil
}

func (x *FootLineupResponse) GetCoach() *Coach {
	if x != nil {
		return x.Coach
	}
	return nil
}

func (x *FootLineupResponse) GetGoalkeeper() []*FootLineupBack {
	if x != nil {
		return x.Goalkeeper
	}
	return nil
}

func (x *FootLineupResponse) GetMidfield() []*FootLineupBack {
	if x != nil {
		return x.Midfield
	}
	return nil
}

func (x *FootLineupResponse) GetVanguard() []*FootLineupBack {
	if x != nil {
		return x.Vanguard
	}
	return nil
}

type FootLineupBack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonId    int64  `protobuf:"varint,1,opt,name=personId,proto3" json:"personId,omitempty"`
	PersonName  string `protobuf:"bytes,2,opt,name=personName,proto3" json:"personName,omitempty"`
	Age         int64  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	ClothesNum  int64  `protobuf:"varint,4,opt,name=clothesNum,proto3" json:"clothesNum,omitempty"`
	Country     string `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	Image       string `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	Position    string `protobuf:"bytes,7,opt,name=position,proto3" json:"position,omitempty"`
	GamesPlayed int64  `protobuf:"varint,8,opt,name=gamesPlayed,proto3" json:"gamesPlayed,omitempty"`
	Goals       int64  `protobuf:"varint,9,opt,name=goals,proto3" json:"goals,omitempty"`
	Assists     int64  `protobuf:"varint,10,opt,name=assists,proto3" json:"assists,omitempty"`
}

func (x *FootLineupBack) Reset() {
	*x = FootLineupBack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_lineup_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootLineupBack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootLineupBack) ProtoMessage() {}

func (x *FootLineupBack) ProtoReflect() protoreflect.Message {
	mi := &file_foot_lineup_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootLineupBack.ProtoReflect.Descriptor instead.
func (*FootLineupBack) Descriptor() ([]byte, []int) {
	return file_foot_lineup_proto_rawDescGZIP(), []int{2}
}

func (x *FootLineupBack) GetPersonId() int64 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

func (x *FootLineupBack) GetPersonName() string {
	if x != nil {
		return x.PersonName
	}
	return ""
}

func (x *FootLineupBack) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *FootLineupBack) GetClothesNum() int64 {
	if x != nil {
		return x.ClothesNum
	}
	return 0
}

func (x *FootLineupBack) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *FootLineupBack) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *FootLineupBack) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *FootLineupBack) GetGamesPlayed() int64 {
	if x != nil {
		return x.GamesPlayed
	}
	return 0
}

func (x *FootLineupBack) GetGoals() int64 {
	if x != nil {
		return x.Goals
	}
	return 0
}

func (x *FootLineupBack) GetAssists() int64 {
	if x != nil {
		return x.Assists
	}
	return 0
}

type Coach struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image      string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	PersonName string `protobuf:"bytes,2,opt,name=personName,proto3" json:"personName,omitempty"`
}

func (x *Coach) Reset() {
	*x = Coach{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_lineup_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coach) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coach) ProtoMessage() {}

func (x *Coach) ProtoReflect() protoreflect.Message {
	mi := &file_foot_lineup_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coach.ProtoReflect.Descriptor instead.
func (*Coach) Descriptor() ([]byte, []int) {
	return file_foot_lineup_proto_rawDescGZIP(), []int{3}
}

func (x *Coach) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Coach) GetPersonName() string {
	if x != nil {
		return x.PersonName
	}
	return ""
}

var File_foot_lineup_proto protoreflect.FileDescriptor

var file_foot_lineup_proto_rawDesc = []byte{
	0x0a, 0x11, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x11, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0xe2, 0x01, 0x0a, 0x12, 0x46, 0x6f, 0x6f, 0x74, 0x4c,
	0x69, 0x6e, 0x65, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a,
	0x04, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x46, 0x6f,
	0x6f, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x42, 0x61, 0x63, 0x6b, 0x52, 0x04, 0x62, 0x61,
	0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x05, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x06, 0x2e, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x52, 0x05, 0x63, 0x6f, 0x61, 0x63, 0x68,
	0x12, 0x2f, 0x0a, 0x0a, 0x67, 0x6f, 0x61, 0x6c, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x75,
	0x70, 0x42, 0x61, 0x63, 0x6b, 0x52, 0x0a, 0x67, 0x6f, 0x61, 0x6c, 0x6b, 0x65, 0x65, 0x70, 0x65,
	0x72, 0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x69, 0x64, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70,
	0x42, 0x61, 0x63, 0x6b, 0x52, 0x08, 0x6d, 0x69, 0x64, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2b,
	0x0a, 0x08, 0x76, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x42, 0x61, 0x63,
	0x6b, 0x52, 0x08, 0x76, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x22, 0x9c, 0x02, 0x0a, 0x0e,
	0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x42, 0x61, 0x63, 0x6b, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x6c, 0x6f, 0x74, 0x68, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x63, 0x6c, 0x6f, 0x74, 0x68, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x67, 0x61, 0x6d, 0x65,
	0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x67,
	0x61, 0x6d, 0x65, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x6f,
	0x61, 0x6c, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x67, 0x6f, 0x61, 0x6c, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x73, 0x22, 0x3d, 0x0a, 0x05, 0x43, 0x6f,
	0x61, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foot_lineup_proto_rawDescOnce sync.Once
	file_foot_lineup_proto_rawDescData = file_foot_lineup_proto_rawDesc
)

func file_foot_lineup_proto_rawDescGZIP() []byte {
	file_foot_lineup_proto_rawDescOnce.Do(func() {
		file_foot_lineup_proto_rawDescData = protoimpl.X.CompressGZIP(file_foot_lineup_proto_rawDescData)
	})
	return file_foot_lineup_proto_rawDescData
}

var file_foot_lineup_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_foot_lineup_proto_goTypes = []interface{}{
	(*FootLineupRequest)(nil),  // 0: FootLineupRequest
	(*FootLineupResponse)(nil), // 1: FootLineupResponse
	(*FootLineupBack)(nil),     // 2: FootLineupBack
	(*Coach)(nil),              // 3: Coach
}
var file_foot_lineup_proto_depIdxs = []int32{
	2, // 0: FootLineupResponse.back:type_name -> FootLineupBack
	3, // 1: FootLineupResponse.coach:type_name -> Coach
	2, // 2: FootLineupResponse.goalkeeper:type_name -> FootLineupBack
	2, // 3: FootLineupResponse.midfield:type_name -> FootLineupBack
	2, // 4: FootLineupResponse.vanguard:type_name -> FootLineupBack
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_foot_lineup_proto_init() }
func file_foot_lineup_proto_init() {
	if File_foot_lineup_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foot_lineup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootLineupRequest); i {
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
		file_foot_lineup_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootLineupResponse); i {
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
		file_foot_lineup_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootLineupBack); i {
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
		file_foot_lineup_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coach); i {
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
			RawDescriptor: file_foot_lineup_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foot_lineup_proto_goTypes,
		DependencyIndexes: file_foot_lineup_proto_depIdxs,
		MessageInfos:      file_foot_lineup_proto_msgTypes,
	}.Build()
	File_foot_lineup_proto = out.File
	file_foot_lineup_proto_rawDesc = nil
	file_foot_lineup_proto_goTypes = nil
	file_foot_lineup_proto_depIdxs = nil
}
