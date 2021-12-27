// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: web_foot_match_incident.proto

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

//事件
type WebFootMatchIncidentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchId  int64  `protobuf:"varint,1,opt,name=matchId,proto3" json:"matchId,omitempty"`  //比赛id
	Language string `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"` //请求语言
}

func (x *WebFootMatchIncidentRequest) Reset() {
	*x = WebFootMatchIncidentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_match_incident_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootMatchIncidentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootMatchIncidentRequest) ProtoMessage() {}

func (x *WebFootMatchIncidentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_match_incident_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootMatchIncidentRequest.ProtoReflect.Descriptor instead.
func (*WebFootMatchIncidentRequest) Descriptor() ([]byte, []int) {
	return file_web_foot_match_incident_proto_rawDescGZIP(), []int{0}
}

func (x *WebFootMatchIncidentRequest) GetMatchId() int64 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

func (x *WebFootMatchIncidentRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type WebFootMatchIncidentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                    //   比赛id
	MatchTime    int64               `protobuf:"varint,2,opt,name=matchTime,proto3" json:"matchTime,omitempty"`      //   日期
	MainTeamName string              `protobuf:"bytes,3,opt,name=mainTeamName,proto3" json:"mainTeamName,omitempty"` //   主队名称
	CustTeamName string              `protobuf:"bytes,4,opt,name=custTeamName,proto3" json:"custTeamName,omitempty"` //   客队名称
	MainTeamId   int64               `protobuf:"varint,8,opt,name=mainTeamId,proto3" json:"mainTeamId,omitempty"`    //   主队id
	CustTeamId   int64               `protobuf:"varint,9,opt,name=custTeamId,proto3" json:"custTeamId,omitempty"`    //   客队id
	MainImage    string              `protobuf:"bytes,10,opt,name=mainImage,proto3" json:"mainImage,omitempty"`      //   主队图片
	CustImage    string              `protobuf:"bytes,11,opt,name=custImage,proto3" json:"custImage,omitempty"`      //   客队图片
	List         []*WebMatchIncident `protobuf:"bytes,5,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *WebFootMatchIncidentResponse) Reset() {
	*x = WebFootMatchIncidentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_match_incident_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootMatchIncidentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootMatchIncidentResponse) ProtoMessage() {}

func (x *WebFootMatchIncidentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_match_incident_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootMatchIncidentResponse.ProtoReflect.Descriptor instead.
func (*WebFootMatchIncidentResponse) Descriptor() ([]byte, []int) {
	return file_web_foot_match_incident_proto_rawDescGZIP(), []int{1}
}

func (x *WebFootMatchIncidentResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WebFootMatchIncidentResponse) GetMatchTime() int64 {
	if x != nil {
		return x.MatchTime
	}
	return 0
}

func (x *WebFootMatchIncidentResponse) GetMainTeamName() string {
	if x != nil {
		return x.MainTeamName
	}
	return ""
}

func (x *WebFootMatchIncidentResponse) GetCustTeamName() string {
	if x != nil {
		return x.CustTeamName
	}
	return ""
}

func (x *WebFootMatchIncidentResponse) GetMainTeamId() int64 {
	if x != nil {
		return x.MainTeamId
	}
	return 0
}

func (x *WebFootMatchIncidentResponse) GetCustTeamId() int64 {
	if x != nil {
		return x.CustTeamId
	}
	return 0
}

func (x *WebFootMatchIncidentResponse) GetMainImage() string {
	if x != nil {
		return x.MainImage
	}
	return ""
}

func (x *WebFootMatchIncidentResponse) GetCustImage() string {
	if x != nil {
		return x.CustImage
	}
	return ""
}

func (x *WebFootMatchIncidentResponse) GetList() []*WebMatchIncident {
	if x != nil {
		return x.List
	}
	return nil
}

type WebMatchIncident struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamMainScore     int64  `protobuf:"varint,5,opt,name=teamMainScore,proto3" json:"teamMainScore,omitempty"`        //   主队得分分数
	TeamCustScore     int64  `protobuf:"varint,6,opt,name=teamCustScore,proto3" json:"teamCustScore,omitempty"`        //   客队得分分数
	Team              int64  `protobuf:"varint,1,opt,name=team,proto3" json:"team,omitempty"`                          //0主队 1客队
	PlayerId          int64  `protobuf:"varint,7,opt,name=playerId,proto3" json:"playerId,omitempty"`                  //球员名称
	RelatePlayerId    int64  `protobuf:"varint,8,opt,name=RelatePlayerId,proto3" json:"RelatePlayerId,omitempty"`      //关联球员，默认为空
	PlayerName        string `protobuf:"bytes,2,opt,name=playerName,proto3" json:"playerName,omitempty"`               //球员名称
	PlayerImage       string `protobuf:"bytes,10,opt,name=playerImage,proto3" json:"playerImage,omitempty"`            //球员图片
	RelatePlayerName  string `protobuf:"bytes,4,opt,name=RelatePlayerName,proto3" json:"RelatePlayerName,omitempty"`   //关联球员，默认为空
	RelatePlayerImage string `protobuf:"bytes,9,opt,name=RelatePlayerImage,proto3" json:"RelatePlayerImage,omitempty"` //关联球员图片，默认为空
	Record            int64  `protobuf:"varint,3,opt,name=record,proto3" json:"record,omitempty"`                      //记录（类型）
}

func (x *WebMatchIncident) Reset() {
	*x = WebMatchIncident{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_match_incident_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebMatchIncident) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebMatchIncident) ProtoMessage() {}

func (x *WebMatchIncident) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_match_incident_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebMatchIncident.ProtoReflect.Descriptor instead.
func (*WebMatchIncident) Descriptor() ([]byte, []int) {
	return file_web_foot_match_incident_proto_rawDescGZIP(), []int{2}
}

func (x *WebMatchIncident) GetTeamMainScore() int64 {
	if x != nil {
		return x.TeamMainScore
	}
	return 0
}

func (x *WebMatchIncident) GetTeamCustScore() int64 {
	if x != nil {
		return x.TeamCustScore
	}
	return 0
}

func (x *WebMatchIncident) GetTeam() int64 {
	if x != nil {
		return x.Team
	}
	return 0
}

func (x *WebMatchIncident) GetPlayerId() int64 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *WebMatchIncident) GetRelatePlayerId() int64 {
	if x != nil {
		return x.RelatePlayerId
	}
	return 0
}

func (x *WebMatchIncident) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *WebMatchIncident) GetPlayerImage() string {
	if x != nil {
		return x.PlayerImage
	}
	return ""
}

func (x *WebMatchIncident) GetRelatePlayerName() string {
	if x != nil {
		return x.RelatePlayerName
	}
	return ""
}

func (x *WebMatchIncident) GetRelatePlayerImage() string {
	if x != nil {
		return x.RelatePlayerImage
	}
	return ""
}

func (x *WebMatchIncident) GetRecord() int64 {
	if x != nil {
		return x.Record
	}
	return 0
}

var File_web_foot_match_incident_proto protoreflect.FileDescriptor

var file_web_foot_match_incident_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x77, 0x65, 0x62, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x5f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x53, 0x0a, 0x1b, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49,
	0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x22, 0xb7, 0x02, 0x0a, 0x1c, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x69, 0x6e, 0x54,
	0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x54,
	0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63,
	0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d,
	0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6d,
	0x61, 0x69, 0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x61, 0x69, 0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x75, 0x73,
	0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x75,
	0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x57, 0x65, 0x62, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xea,
	0x02, 0x0a, 0x10, 0x57, 0x65, 0x62, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x63, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x65, 0x61, 0x6d,
	0x4d, 0x61, 0x69, 0x6e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x65, 0x61,
	0x6d, 0x43, 0x75, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x74, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74,
	0x65, 0x61, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x26, 0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x07, 0x5a, 0x05, 0x2e,
	0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web_foot_match_incident_proto_rawDescOnce sync.Once
	file_web_foot_match_incident_proto_rawDescData = file_web_foot_match_incident_proto_rawDesc
)

func file_web_foot_match_incident_proto_rawDescGZIP() []byte {
	file_web_foot_match_incident_proto_rawDescOnce.Do(func() {
		file_web_foot_match_incident_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_foot_match_incident_proto_rawDescData)
	})
	return file_web_foot_match_incident_proto_rawDescData
}

var file_web_foot_match_incident_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_web_foot_match_incident_proto_goTypes = []interface{}{
	(*WebFootMatchIncidentRequest)(nil),  // 0: WebFootMatchIncidentRequest
	(*WebFootMatchIncidentResponse)(nil), // 1: WebFootMatchIncidentResponse
	(*WebMatchIncident)(nil),             // 2: WebMatchIncident
}
var file_web_foot_match_incident_proto_depIdxs = []int32{
	2, // 0: WebFootMatchIncidentResponse.list:type_name -> WebMatchIncident
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_web_foot_match_incident_proto_init() }
func file_web_foot_match_incident_proto_init() {
	if File_web_foot_match_incident_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_foot_match_incident_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootMatchIncidentRequest); i {
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
		file_web_foot_match_incident_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootMatchIncidentResponse); i {
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
		file_web_foot_match_incident_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebMatchIncident); i {
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
			RawDescriptor: file_web_foot_match_incident_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_foot_match_incident_proto_goTypes,
		DependencyIndexes: file_web_foot_match_incident_proto_depIdxs,
		MessageInfos:      file_web_foot_match_incident_proto_msgTypes,
	}.Build()
	File_web_foot_match_incident_proto = out.File
	file_web_foot_match_incident_proto_rawDesc = nil
	file_web_foot_match_incident_proto_goTypes = nil
	file_web_foot_match_incident_proto_depIdxs = nil
}
