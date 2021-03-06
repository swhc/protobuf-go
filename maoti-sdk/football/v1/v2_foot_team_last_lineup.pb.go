// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: v2_foot_team_last_lineup.proto

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

//阵容
type V2FootMatchLastLineupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`  //请求语言
	TeamId   int64  `protobuf:"varint,6,opt,name=teamId,proto3" json:"teamId,omitempty"`     //球队ID
	TimeZone int64  `protobuf:"varint,1,opt,name=TimeZone,proto3" json:"TimeZone,omitempty"` //时区
}

func (x *V2FootMatchLastLineupRequest) Reset() {
	*x = V2FootMatchLastLineupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_team_last_lineup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchLastLineupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchLastLineupRequest) ProtoMessage() {}

func (x *V2FootMatchLastLineupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_team_last_lineup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchLastLineupRequest.ProtoReflect.Descriptor instead.
func (*V2FootMatchLastLineupRequest) Descriptor() ([]byte, []int) {
	return file_v2_foot_team_last_lineup_proto_rawDescGZIP(), []int{0}
}

func (x *V2FootMatchLastLineupRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *V2FootMatchLastLineupRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *V2FootMatchLastLineupRequest) GetTimeZone() int64 {
	if x != nil {
		return x.TimeZone
	}
	return 0
}

type V2FootMatchLastLineupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64                    `protobuf:"varint,4,opt,name=Id,proto3" json:"Id,omitempty"`                  //比赛id
	MatchTime     int64                    `protobuf:"varint,6,opt,name=matchTime,proto3" json:"matchTime,omitempty"`    //比赛时间
	HomeTeamId    int64                    `protobuf:"varint,13,opt,name=homeTeamId,proto3" json:"homeTeamId,omitempty"` //主队Id
	HomeTeam      string                   `protobuf:"bytes,7,opt,name=homeTeam,proto3" json:"homeTeam,omitempty"`       //主队
	HomeScore     int64                    `protobuf:"varint,9,opt,name=homeScore,proto3" json:"homeScore,omitempty"`    //主队比分
	AwayTeamId    int64                    `protobuf:"varint,14,opt,name=awayTeamId,proto3" json:"awayTeamId,omitempty"` //客队Id
	AwayTeam      string                   `protobuf:"bytes,10,opt,name=awayTeam,proto3" json:"awayTeam,omitempty"`      //客队
	AwayScore     int64                    `protobuf:"varint,12,opt,name=awayScore,proto3" json:"awayScore,omitempty"`   //客队比分
	HomeImage     string                   `protobuf:"bytes,8,opt,name=homeImage,proto3" json:"homeImage,omitempty"`
	AwayImage     string                   `protobuf:"bytes,5,opt,name=awayImage,proto3" json:"awayImage,omitempty"`
	HomeFormation string                   `protobuf:"bytes,1,opt,name=homeFormation,proto3" json:"homeFormation,omitempty"` //主队队形
	AwayFormation string                   `protobuf:"bytes,2,opt,name=awayFormation,proto3" json:"awayFormation,omitempty"` //客队队形
	Lineup        []*V2FootMatchLastLineup `protobuf:"bytes,3,rep,name=lineup,proto3" json:"lineup,omitempty"`
}

func (x *V2FootMatchLastLineupResponse) Reset() {
	*x = V2FootMatchLastLineupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_team_last_lineup_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchLastLineupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchLastLineupResponse) ProtoMessage() {}

func (x *V2FootMatchLastLineupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_team_last_lineup_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchLastLineupResponse.ProtoReflect.Descriptor instead.
func (*V2FootMatchLastLineupResponse) Descriptor() ([]byte, []int) {
	return file_v2_foot_team_last_lineup_proto_rawDescGZIP(), []int{1}
}

func (x *V2FootMatchLastLineupResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *V2FootMatchLastLineupResponse) GetMatchTime() int64 {
	if x != nil {
		return x.MatchTime
	}
	return 0
}

func (x *V2FootMatchLastLineupResponse) GetHomeTeamId() int64 {
	if x != nil {
		return x.HomeTeamId
	}
	return 0
}

func (x *V2FootMatchLastLineupResponse) GetHomeTeam() string {
	if x != nil {
		return x.HomeTeam
	}
	return ""
}

func (x *V2FootMatchLastLineupResponse) GetHomeScore() int64 {
	if x != nil {
		return x.HomeScore
	}
	return 0
}

func (x *V2FootMatchLastLineupResponse) GetAwayTeamId() int64 {
	if x != nil {
		return x.AwayTeamId
	}
	return 0
}

func (x *V2FootMatchLastLineupResponse) GetAwayTeam() string {
	if x != nil {
		return x.AwayTeam
	}
	return ""
}

func (x *V2FootMatchLastLineupResponse) GetAwayScore() int64 {
	if x != nil {
		return x.AwayScore
	}
	return 0
}

func (x *V2FootMatchLastLineupResponse) GetHomeImage() string {
	if x != nil {
		return x.HomeImage
	}
	return ""
}

func (x *V2FootMatchLastLineupResponse) GetAwayImage() string {
	if x != nil {
		return x.AwayImage
	}
	return ""
}

func (x *V2FootMatchLastLineupResponse) GetHomeFormation() string {
	if x != nil {
		return x.HomeFormation
	}
	return ""
}

func (x *V2FootMatchLastLineupResponse) GetAwayFormation() string {
	if x != nil {
		return x.AwayFormation
	}
	return ""
}

func (x *V2FootMatchLastLineupResponse) GetLineup() []*V2FootMatchLastLineup {
	if x != nil {
		return x.Lineup
	}
	return nil
}

type V2FootMatchLastLineup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pos      string `protobuf:"bytes,1,opt,name=pos,proto3" json:"pos,omitempty"`            //位置
	Position int64  `protobuf:"varint,2,opt,name=position,proto3" json:"position,omitempty"` //位置坐标
	Number   int64  `protobuf:"varint,3,opt,name=Number,proto3" json:"Number,omitempty"`     //球衣号
	Score    string `protobuf:"bytes,4,opt,name=Score,proto3" json:"Score,omitempty"`        //评分
	Name     string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`          //球员名字
	Image    string `protobuf:"bytes,7,opt,name=image,proto3" json:"image,omitempty"`        //球员图片
	Id       int64  `protobuf:"varint,6,opt,name=id,proto3" json:"id,omitempty"`             //球员id
}

func (x *V2FootMatchLastLineup) Reset() {
	*x = V2FootMatchLastLineup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_team_last_lineup_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchLastLineup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchLastLineup) ProtoMessage() {}

func (x *V2FootMatchLastLineup) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_team_last_lineup_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchLastLineup.ProtoReflect.Descriptor instead.
func (*V2FootMatchLastLineup) Descriptor() ([]byte, []int) {
	return file_v2_foot_team_last_lineup_proto_rawDescGZIP(), []int{2}
}

func (x *V2FootMatchLastLineup) GetPos() string {
	if x != nil {
		return x.Pos
	}
	return ""
}

func (x *V2FootMatchLastLineup) GetPosition() int64 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *V2FootMatchLastLineup) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *V2FootMatchLastLineup) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

func (x *V2FootMatchLastLineup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *V2FootMatchLastLineup) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *V2FootMatchLastLineup) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_v2_foot_team_last_lineup_proto protoreflect.FileDescriptor

var file_v2_foot_team_last_lineup_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x76, 0x32, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6e, 0x0a, 0x1c, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c,
	0x61, 0x73, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65,
	0x61, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65,
	0x22, 0xb9, 0x03, 0x0a, 0x1d, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x4c, 0x61, 0x73, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x1c, 0x0a, 0x09,
	0x68, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x68, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x77,
	0x61, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x61, 0x77, 0x61, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x77,
	0x61, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x77,
	0x61, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x77, 0x61, 0x79, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x77, 0x61, 0x79, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x6f, 0x6d, 0x65, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x6f, 0x6d, 0x65, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x77, 0x61, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x77, 0x61, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x24, 0x0a, 0x0d, 0x68, 0x6f, 0x6d, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x68, 0x6f, 0x6d, 0x65, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x77, 0x61, 0x79, 0x46, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61,
	0x77, 0x61, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x06,
	0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x56,
	0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x61, 0x73, 0x74, 0x4c, 0x69,
	0x6e, 0x65, 0x75, 0x70, 0x52, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x22, 0xad, 0x01, 0x0a,
	0x15, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x61, 0x73, 0x74,
	0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_foot_team_last_lineup_proto_rawDescOnce sync.Once
	file_v2_foot_team_last_lineup_proto_rawDescData = file_v2_foot_team_last_lineup_proto_rawDesc
)

func file_v2_foot_team_last_lineup_proto_rawDescGZIP() []byte {
	file_v2_foot_team_last_lineup_proto_rawDescOnce.Do(func() {
		file_v2_foot_team_last_lineup_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_foot_team_last_lineup_proto_rawDescData)
	})
	return file_v2_foot_team_last_lineup_proto_rawDescData
}

var file_v2_foot_team_last_lineup_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v2_foot_team_last_lineup_proto_goTypes = []interface{}{
	(*V2FootMatchLastLineupRequest)(nil),  // 0: V2FootMatchLastLineupRequest
	(*V2FootMatchLastLineupResponse)(nil), // 1: V2FootMatchLastLineupResponse
	(*V2FootMatchLastLineup)(nil),         // 2: V2FootMatchLastLineup
}
var file_v2_foot_team_last_lineup_proto_depIdxs = []int32{
	2, // 0: V2FootMatchLastLineupResponse.lineup:type_name -> V2FootMatchLastLineup
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v2_foot_team_last_lineup_proto_init() }
func file_v2_foot_team_last_lineup_proto_init() {
	if File_v2_foot_team_last_lineup_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_foot_team_last_lineup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchLastLineupRequest); i {
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
		file_v2_foot_team_last_lineup_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchLastLineupResponse); i {
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
		file_v2_foot_team_last_lineup_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchLastLineup); i {
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
			RawDescriptor: file_v2_foot_team_last_lineup_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_foot_team_last_lineup_proto_goTypes,
		DependencyIndexes: file_v2_foot_team_last_lineup_proto_depIdxs,
		MessageInfos:      file_v2_foot_team_last_lineup_proto_msgTypes,
	}.Build()
	File_v2_foot_team_last_lineup_proto = out.File
	file_v2_foot_team_last_lineup_proto_rawDesc = nil
	file_v2_foot_team_last_lineup_proto_goTypes = nil
	file_v2_foot_team_last_lineup_proto_depIdxs = nil
}
