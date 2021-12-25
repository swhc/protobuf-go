// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: web_foot_stand.proto

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

//积分榜
type WebLeagueStandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeasonId int64  `protobuf:"varint,2,opt,name=seasonId,proto3" json:"seasonId,omitempty"` //联赛id
	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`  //请求语言  1:zh  2:en
}

func (x *WebLeagueStandRequest) Reset() {
	*x = WebLeagueStandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_stand_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebLeagueStandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebLeagueStandRequest) ProtoMessage() {}

func (x *WebLeagueStandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_stand_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebLeagueStandRequest.ProtoReflect.Descriptor instead.
func (*WebLeagueStandRequest) Descriptor() ([]byte, []int) {
	return file_web_foot_stand_proto_rawDescGZIP(), []int{0}
}

func (x *WebLeagueStandRequest) GetSeasonId() int64 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

func (x *WebLeagueStandRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type WebLeagueStandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*LeagueStand `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Count int64          `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *WebLeagueStandResponse) Reset() {
	*x = WebLeagueStandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_stand_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebLeagueStandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebLeagueStandResponse) ProtoMessage() {}

func (x *WebLeagueStandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_stand_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebLeagueStandResponse.ProtoReflect.Descriptor instead.
func (*WebLeagueStandResponse) Descriptor() ([]byte, []int) {
	return file_web_foot_stand_proto_rawDescGZIP(), []int{1}
}

func (x *WebLeagueStandResponse) GetList() []*LeagueStand {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *WebLeagueStandResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type LeagueStand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64              `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`    //积分榜id
	Name string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` //积分榜名称
	Info []*LeagueStandInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"` //积分榜
}

func (x *LeagueStand) Reset() {
	*x = LeagueStand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_stand_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeagueStand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeagueStand) ProtoMessage() {}

func (x *LeagueStand) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_stand_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeagueStand.ProtoReflect.Descriptor instead.
func (*LeagueStand) Descriptor() ([]byte, []int) {
	return file_web_foot_stand_proto_rawDescGZIP(), []int{2}
}

func (x *LeagueStand) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LeagueStand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LeagueStand) GetInfo() []*LeagueStandInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type LeagueStandInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId    int64  `protobuf:"varint,2,opt,name=teamId,proto3" json:"teamId,omitempty"`       //队伍id
	TeamName  string `protobuf:"bytes,1,opt,name=teamName,proto3" json:"teamName,omitempty"`    //队伍名字
	TeamImage string `protobuf:"bytes,4,opt,name=teamImage,proto3" json:"teamImage,omitempty"`  //队伍logo
	PlayCount int64  `protobuf:"varint,3,opt,name=playCount,proto3" json:"playCount,omitempty"` //已赛场次
	Win       int64  `protobuf:"varint,5,opt,name=win,proto3" json:"win,omitempty"`             //赢场
	Draw      int64  `protobuf:"varint,6,opt,name=draw,proto3" json:"draw,omitempty"`           //平场次
	Lose      int64  `protobuf:"varint,7,opt,name=lose,proto3" json:"lose,omitempty"`           //输场次
	Goals     int64  `protobuf:"varint,8,opt,name=goals,proto3" json:"goals,omitempty"`         //进球
	Turnovers int64  `protobuf:"varint,9,opt,name=turnovers,proto3" json:"turnovers,omitempty"` //失球
	Score     int64  `protobuf:"varint,10,opt,name=score,proto3" json:"score,omitempty"`        //积分
	Trend     string `protobuf:"bytes,11,opt,name=trend,proto3" json:"trend,omitempty"`         //近六场走势
}

func (x *LeagueStandInfo) Reset() {
	*x = LeagueStandInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_stand_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeagueStandInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeagueStandInfo) ProtoMessage() {}

func (x *LeagueStandInfo) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_stand_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeagueStandInfo.ProtoReflect.Descriptor instead.
func (*LeagueStandInfo) Descriptor() ([]byte, []int) {
	return file_web_foot_stand_proto_rawDescGZIP(), []int{3}
}

func (x *LeagueStandInfo) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *LeagueStandInfo) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

func (x *LeagueStandInfo) GetTeamImage() string {
	if x != nil {
		return x.TeamImage
	}
	return ""
}

func (x *LeagueStandInfo) GetPlayCount() int64 {
	if x != nil {
		return x.PlayCount
	}
	return 0
}

func (x *LeagueStandInfo) GetWin() int64 {
	if x != nil {
		return x.Win
	}
	return 0
}

func (x *LeagueStandInfo) GetDraw() int64 {
	if x != nil {
		return x.Draw
	}
	return 0
}

func (x *LeagueStandInfo) GetLose() int64 {
	if x != nil {
		return x.Lose
	}
	return 0
}

func (x *LeagueStandInfo) GetGoals() int64 {
	if x != nil {
		return x.Goals
	}
	return 0
}

func (x *LeagueStandInfo) GetTurnovers() int64 {
	if x != nil {
		return x.Turnovers
	}
	return 0
}

func (x *LeagueStandInfo) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *LeagueStandInfo) GetTrend() string {
	if x != nil {
		return x.Trend
	}
	return ""
}

var File_web_foot_stand_proto protoreflect.FileDescriptor

var file_web_foot_stand_proto_rawDesc = []byte{
	0x0a, 0x14, 0x77, 0x65, 0x62, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x15, 0x57, 0x65, 0x62, 0x4c, 0x65, 0x61,
	0x67, 0x75, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x50, 0x0a, 0x16, 0x57, 0x65, 0x62, 0x4c, 0x65,
	0x61, 0x67, 0x75, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x57, 0x0a, 0x0b, 0x6c, 0x65, 0x61,
	0x67, 0x75, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6c, 0x65, 0x61,
	0x67, 0x75, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x22, 0x9b, 0x02, 0x0a, 0x0f, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x74, 0x61,
	0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65,
	0x61, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x65, 0x61, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x6c, 0x61,
	0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x69, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x77, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x72, 0x61, 0x77,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x72, 0x61, 0x77, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x6f, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6c, 0x6f, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x6f, 0x61, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x67, 0x6f, 0x61, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x75, 0x72, 0x6e, 0x6f, 0x76,
	0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x75, 0x72, 0x6e, 0x6f,
	0x76, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x72,
	0x65, 0x6e, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x72, 0x65, 0x6e, 0x64,
	0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_web_foot_stand_proto_rawDescOnce sync.Once
	file_web_foot_stand_proto_rawDescData = file_web_foot_stand_proto_rawDesc
)

func file_web_foot_stand_proto_rawDescGZIP() []byte {
	file_web_foot_stand_proto_rawDescOnce.Do(func() {
		file_web_foot_stand_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_foot_stand_proto_rawDescData)
	})
	return file_web_foot_stand_proto_rawDescData
}

var file_web_foot_stand_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_web_foot_stand_proto_goTypes = []interface{}{
	(*WebLeagueStandRequest)(nil),  // 0: WebLeagueStandRequest
	(*WebLeagueStandResponse)(nil), // 1: WebLeagueStandResponse
	(*LeagueStand)(nil),            // 2: leagueStand
	(*LeagueStandInfo)(nil),        // 3: leagueStandInfo
}
var file_web_foot_stand_proto_depIdxs = []int32{
	2, // 0: WebLeagueStandResponse.list:type_name -> leagueStand
	3, // 1: leagueStand.info:type_name -> leagueStandInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_web_foot_stand_proto_init() }
func file_web_foot_stand_proto_init() {
	if File_web_foot_stand_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_foot_stand_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebLeagueStandRequest); i {
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
		file_web_foot_stand_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebLeagueStandResponse); i {
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
		file_web_foot_stand_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeagueStand); i {
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
		file_web_foot_stand_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeagueStandInfo); i {
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
			RawDescriptor: file_web_foot_stand_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_foot_stand_proto_goTypes,
		DependencyIndexes: file_web_foot_stand_proto_depIdxs,
		MessageInfos:      file_web_foot_stand_proto_msgTypes,
	}.Build()
	File_web_foot_stand_proto = out.File
	file_web_foot_stand_proto_rawDesc = nil
	file_web_foot_stand_proto_goTypes = nil
	file_web_foot_stand_proto_depIdxs = nil
}
