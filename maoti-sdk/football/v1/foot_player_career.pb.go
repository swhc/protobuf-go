// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: foot_player_career.proto

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
type FootPlayerCareerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"` //请求语言  1:zh  2:en
	PersonId string `protobuf:"bytes,2,opt,name=personId,proto3" json:"personId,omitempty"` //比赛id
}

func (x *FootPlayerCareerRequest) Reset() {
	*x = FootPlayerCareerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_player_career_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootPlayerCareerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootPlayerCareerRequest) ProtoMessage() {}

func (x *FootPlayerCareerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foot_player_career_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootPlayerCareerRequest.ProtoReflect.Descriptor instead.
func (*FootPlayerCareerRequest) Descriptor() ([]byte, []int) {
	return file_foot_player_career_proto_rawDescGZIP(), []int{0}
}

func (x *FootPlayerCareerRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *FootPlayerCareerRequest) GetPersonId() string {
	if x != nil {
		return x.PersonId
	}
	return ""
}

type FootPlayerCareerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*FootPlayerRecords `protobuf:"bytes,20,rep,name=Records,proto3" json:"Records,omitempty"` //转会记录
	Career  []*FootPlayerCareer  `protobuf:"bytes,21,rep,name=Career,proto3" json:"Career,omitempty"`   //职业生涯
}

func (x *FootPlayerCareerResponse) Reset() {
	*x = FootPlayerCareerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_player_career_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootPlayerCareerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootPlayerCareerResponse) ProtoMessage() {}

func (x *FootPlayerCareerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_foot_player_career_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootPlayerCareerResponse.ProtoReflect.Descriptor instead.
func (*FootPlayerCareerResponse) Descriptor() ([]byte, []int) {
	return file_foot_player_career_proto_rawDescGZIP(), []int{1}
}

func (x *FootPlayerCareerResponse) GetRecords() []*FootPlayerRecords {
	if x != nil {
		return x.Records
	}
	return nil
}

func (x *FootPlayerCareerResponse) GetCareer() []*FootPlayerCareer {
	if x != nil {
		return x.Career
	}
	return nil
}

type FootPlayerRecords struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddTime            int64  `protobuf:"varint,4,opt,name=AddTime,proto3" json:"AddTime,omitempty"`                      //转会时间
	EndTime            int64  `protobuf:"varint,5,opt,name=EndTime,proto3" json:"EndTime,omitempty"`                      //结束时间
	ContractCost       string `protobuf:"bytes,6,opt,name=ContractCost,proto3" json:"ContractCost,omitempty"`             //签订费用方式
	OriginalTeamId     string `protobuf:"bytes,7,opt,name=OriginalTeamId,proto3" json:"OriginalTeamId,omitempty"`         //原球队id
	OriginalTeam       string `protobuf:"bytes,8,opt,name=OriginalTeam,proto3" json:"OriginalTeam,omitempty"`             //原球队名称
	OriginalTeamImages string `protobuf:"bytes,9,opt,name=OriginalTeamImages,proto3" json:"OriginalTeamImages,omitempty"` //原队伍图片
	NewTeamId          string `protobuf:"bytes,10,opt,name=NewTeamId,proto3" json:"NewTeamId,omitempty"`                  //新球队id
	NewTeam            string `protobuf:"bytes,11,opt,name=NewTeam,proto3" json:"NewTeam,omitempty"`                      //新球队名称
	NewTeamImages      string `protobuf:"bytes,12,opt,name=NewTeamImages,proto3" json:"NewTeamImages,omitempty"`          //新球队图片
}

func (x *FootPlayerRecords) Reset() {
	*x = FootPlayerRecords{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_player_career_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootPlayerRecords) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootPlayerRecords) ProtoMessage() {}

func (x *FootPlayerRecords) ProtoReflect() protoreflect.Message {
	mi := &file_foot_player_career_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootPlayerRecords.ProtoReflect.Descriptor instead.
func (*FootPlayerRecords) Descriptor() ([]byte, []int) {
	return file_foot_player_career_proto_rawDescGZIP(), []int{2}
}

func (x *FootPlayerRecords) GetAddTime() int64 {
	if x != nil {
		return x.AddTime
	}
	return 0
}

func (x *FootPlayerRecords) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *FootPlayerRecords) GetContractCost() string {
	if x != nil {
		return x.ContractCost
	}
	return ""
}

func (x *FootPlayerRecords) GetOriginalTeamId() string {
	if x != nil {
		return x.OriginalTeamId
	}
	return ""
}

func (x *FootPlayerRecords) GetOriginalTeam() string {
	if x != nil {
		return x.OriginalTeam
	}
	return ""
}

func (x *FootPlayerRecords) GetOriginalTeamImages() string {
	if x != nil {
		return x.OriginalTeamImages
	}
	return ""
}

func (x *FootPlayerRecords) GetNewTeamId() string {
	if x != nil {
		return x.NewTeamId
	}
	return ""
}

func (x *FootPlayerRecords) GetNewTeam() string {
	if x != nil {
		return x.NewTeam
	}
	return ""
}

func (x *FootPlayerRecords) GetNewTeamImages() string {
	if x != nil {
		return x.NewTeamImages
	}
	return ""
}

type FootPlayerCareer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamName  string `protobuf:"bytes,13,opt,name=TeamName,proto3" json:"TeamName,omitempty"`    //球队名称
	TeamImage string `protobuf:"bytes,14,opt,name=TeamImage,proto3" json:"TeamImage,omitempty"`  //球队图片
	EndTime   int64  `protobuf:"varint,15,opt,name=EndTime,proto3" json:"EndTime,omitempty"`     //成立时间
	FoundTime int64  `protobuf:"varint,16,opt,name=FoundTime,proto3" json:"FoundTime,omitempty"` //（成立的）结束时间
	AllCount  int64  `protobuf:"varint,17,opt,name=AllCount,proto3" json:"AllCount,omitempty"`   //总场次
	GetScore  int64  `protobuf:"varint,18,opt,name=GetScore,proto3" json:"GetScore,omitempty"`   //得分
	HelpScore int64  `protobuf:"varint,19,opt,name=HelpScore,proto3" json:"HelpScore,omitempty"` //助攻
}

func (x *FootPlayerCareer) Reset() {
	*x = FootPlayerCareer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_player_career_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootPlayerCareer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootPlayerCareer) ProtoMessage() {}

func (x *FootPlayerCareer) ProtoReflect() protoreflect.Message {
	mi := &file_foot_player_career_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootPlayerCareer.ProtoReflect.Descriptor instead.
func (*FootPlayerCareer) Descriptor() ([]byte, []int) {
	return file_foot_player_career_proto_rawDescGZIP(), []int{3}
}

func (x *FootPlayerCareer) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

func (x *FootPlayerCareer) GetTeamImage() string {
	if x != nil {
		return x.TeamImage
	}
	return ""
}

func (x *FootPlayerCareer) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *FootPlayerCareer) GetFoundTime() int64 {
	if x != nil {
		return x.FoundTime
	}
	return 0
}

func (x *FootPlayerCareer) GetAllCount() int64 {
	if x != nil {
		return x.AllCount
	}
	return 0
}

func (x *FootPlayerCareer) GetGetScore() int64 {
	if x != nil {
		return x.GetScore
	}
	return 0
}

func (x *FootPlayerCareer) GetHelpScore() int64 {
	if x != nil {
		return x.HelpScore
	}
	return 0
}

var File_foot_player_career_proto protoreflect.FileDescriptor

var file_foot_player_career_proto_rawDesc = []byte{
	0x0a, 0x18, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x63, 0x61,
	0x72, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x17, 0x46, 0x6f,
	0x6f, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x61, 0x72, 0x65, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x73, 0x0a,
	0x18, 0x46, 0x6f, 0x6f, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x61, 0x72, 0x65, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x46, 0x6f, 0x6f,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x07,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x29, 0x0a, 0x06, 0x43, 0x61, 0x72, 0x65, 0x65,
	0x72, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x43, 0x61, 0x72, 0x65, 0x65, 0x72, 0x52, 0x06, 0x43, 0x61, 0x72, 0x65,
	0x65, 0x72, 0x22, 0xc5, 0x02, 0x0a, 0x11, 0x46, 0x6f, 0x6f, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x41, 0x64, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x6f, 0x73, 0x74,
	0x12, 0x26, 0x0a, 0x0e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x65, 0x61, 0x6d,
	0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x2e, 0x0a, 0x12,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x4e, 0x65, 0x77, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x4e, 0x65, 0x77, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x65,
	0x77, 0x54, 0x65, 0x61, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4e, 0x65, 0x77,
	0x54, 0x65, 0x61, 0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x4e, 0x65, 0x77, 0x54, 0x65, 0x61, 0x6d, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x4e, 0x65, 0x77,
	0x54, 0x65, 0x61, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0xda, 0x01, 0x0a, 0x10, 0x46,
	0x6f, 0x6f, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x61, 0x72, 0x65, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54,
	0x65, 0x61, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x54, 0x65, 0x61, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x47, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x48, 0x65, 0x6c,
	0x70, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x48, 0x65,
	0x6c, 0x70, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foot_player_career_proto_rawDescOnce sync.Once
	file_foot_player_career_proto_rawDescData = file_foot_player_career_proto_rawDesc
)

func file_foot_player_career_proto_rawDescGZIP() []byte {
	file_foot_player_career_proto_rawDescOnce.Do(func() {
		file_foot_player_career_proto_rawDescData = protoimpl.X.CompressGZIP(file_foot_player_career_proto_rawDescData)
	})
	return file_foot_player_career_proto_rawDescData
}

var file_foot_player_career_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_foot_player_career_proto_goTypes = []interface{}{
	(*FootPlayerCareerRequest)(nil),  // 0: FootPlayerCareerRequest
	(*FootPlayerCareerResponse)(nil), // 1: FootPlayerCareerResponse
	(*FootPlayerRecords)(nil),        // 2: FootPlayerRecords
	(*FootPlayerCareer)(nil),         // 3: FootPlayerCareer
}
var file_foot_player_career_proto_depIdxs = []int32{
	2, // 0: FootPlayerCareerResponse.Records:type_name -> FootPlayerRecords
	3, // 1: FootPlayerCareerResponse.Career:type_name -> FootPlayerCareer
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_foot_player_career_proto_init() }
func file_foot_player_career_proto_init() {
	if File_foot_player_career_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foot_player_career_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootPlayerCareerRequest); i {
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
		file_foot_player_career_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootPlayerCareerResponse); i {
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
		file_foot_player_career_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootPlayerRecords); i {
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
		file_foot_player_career_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootPlayerCareer); i {
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
			RawDescriptor: file_foot_player_career_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foot_player_career_proto_goTypes,
		DependencyIndexes: file_foot_player_career_proto_depIdxs,
		MessageInfos:      file_foot_player_career_proto_msgTypes,
	}.Build()
	File_foot_player_career_proto = out.File
	file_foot_player_career_proto_rawDesc = nil
	file_foot_player_career_proto_goTypes = nil
	file_foot_player_career_proto_depIdxs = nil
}
