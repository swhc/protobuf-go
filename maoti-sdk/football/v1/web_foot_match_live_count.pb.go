// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: web_foot_match_live_count.proto

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
type WebFootMatchCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchId  int64  `protobuf:"varint,1,opt,name=MatchId,proto3" json:"MatchId,omitempty"`  //比赛id
	Language string `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"` //请求语言
}

func (x *WebFootMatchCountRequest) Reset() {
	*x = WebFootMatchCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_match_live_count_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootMatchCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootMatchCountRequest) ProtoMessage() {}

func (x *WebFootMatchCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_match_live_count_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootMatchCountRequest.ProtoReflect.Descriptor instead.
func (*WebFootMatchCountRequest) Descriptor() ([]byte, []int) {
	return file_web_foot_match_live_count_proto_rawDescGZIP(), []int{0}
}

func (x *WebFootMatchCountRequest) GetMatchId() int64 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

func (x *WebFootMatchCountRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type WebFootMatchCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MainTeamId   int64              `protobuf:"varint,5,opt,name=MainTeamId,proto3" json:"MainTeamId,omitempty"`     //主队ID
	MainTeamName string             `protobuf:"bytes,6,opt,name=MainTeamName,proto3" json:"MainTeamName,omitempty"`  //主队名称
	MainTeamLogo string             `protobuf:"bytes,7,opt,name=MainTeamLogo,proto3" json:"MainTeamLogo,omitempty"`  //主队Logo
	CustTeamId   int64              `protobuf:"varint,9,opt,name=CustTeamId,proto3" json:"CustTeamId,omitempty"`     //客队ID
	CustTeamName string             `protobuf:"bytes,10,opt,name=CustTeamName,proto3" json:"CustTeamName,omitempty"` //客队名称
	CustTeamLogo string             `protobuf:"bytes,11,opt,name=CustTeamLogo,proto3" json:"CustTeamLogo,omitempty"` //客队Logo
	Home         *WebFootMatchCount `protobuf:"bytes,1,opt,name=Home,proto3" json:"Home,omitempty"`
	Away         *WebFootMatchCount `protobuf:"bytes,2,opt,name=Away,proto3" json:"Away,omitempty"`
	LivingUrl    string             `protobuf:"bytes,3,opt,name=livingUrl,proto3" json:"livingUrl,omitempty"`
}

func (x *WebFootMatchCountResponse) Reset() {
	*x = WebFootMatchCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_match_live_count_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootMatchCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootMatchCountResponse) ProtoMessage() {}

func (x *WebFootMatchCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_match_live_count_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootMatchCountResponse.ProtoReflect.Descriptor instead.
func (*WebFootMatchCountResponse) Descriptor() ([]byte, []int) {
	return file_web_foot_match_live_count_proto_rawDescGZIP(), []int{1}
}

func (x *WebFootMatchCountResponse) GetMainTeamId() int64 {
	if x != nil {
		return x.MainTeamId
	}
	return 0
}

func (x *WebFootMatchCountResponse) GetMainTeamName() string {
	if x != nil {
		return x.MainTeamName
	}
	return ""
}

func (x *WebFootMatchCountResponse) GetMainTeamLogo() string {
	if x != nil {
		return x.MainTeamLogo
	}
	return ""
}

func (x *WebFootMatchCountResponse) GetCustTeamId() int64 {
	if x != nil {
		return x.CustTeamId
	}
	return 0
}

func (x *WebFootMatchCountResponse) GetCustTeamName() string {
	if x != nil {
		return x.CustTeamName
	}
	return ""
}

func (x *WebFootMatchCountResponse) GetCustTeamLogo() string {
	if x != nil {
		return x.CustTeamLogo
	}
	return ""
}

func (x *WebFootMatchCountResponse) GetHome() *WebFootMatchCount {
	if x != nil {
		return x.Home
	}
	return nil
}

func (x *WebFootMatchCountResponse) GetAway() *WebFootMatchCount {
	if x != nil {
		return x.Away
	}
	return nil
}

func (x *WebFootMatchCountResponse) GetLivingUrl() string {
	if x != nil {
		return x.LivingUrl
	}
	return ""
}

type WebFootMatchCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Possession              int64  `protobuf:"varint,1,opt,name=Possession,proto3" json:"Possession,omitempty"`                          //控球率
	Shots                   int64  `protobuf:"varint,2,opt,name=Shots,proto3" json:"Shots,omitempty"`                                    //射门
	ShotsOnTarget           int64  `protobuf:"varint,3,opt,name=ShotsOnTarget,proto3" json:"ShotsOnTarget,omitempty"`                    //射正
	ShotsOnTargetProportion string `protobuf:"bytes,4,opt,name=ShotsOnTargetProportion,proto3" json:"ShotsOnTargetProportion,omitempty"` //进攻
	Attacks                 int64  `protobuf:"varint,5,opt,name=Attacks,proto3" json:"Attacks,omitempty"`                                //危险进攻
	DangerousAttacks        int64  `protobuf:"varint,6,opt,name=DangerousAttacks,proto3" json:"DangerousAttacks,omitempty"`              //角球
	YellowCards             int64  `protobuf:"varint,7,opt,name=YellowCards,proto3" json:"YellowCards,omitempty"`                        //扑救
	RedCards                int64  `protobuf:"varint,8,opt,name=RedCards,proto3" json:"RedCards,omitempty"`                              //黄牌
	FreeKicks               int64  `protobuf:"varint,9,opt,name=FreeKicks,proto3" json:"FreeKicks,omitempty"`                            //红牌
	Corners                 int64  `protobuf:"varint,10,opt,name=Corners,proto3" json:"Corners,omitempty"`                               //任意球
}

func (x *WebFootMatchCount) Reset() {
	*x = WebFootMatchCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_match_live_count_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootMatchCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootMatchCount) ProtoMessage() {}

func (x *WebFootMatchCount) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_match_live_count_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootMatchCount.ProtoReflect.Descriptor instead.
func (*WebFootMatchCount) Descriptor() ([]byte, []int) {
	return file_web_foot_match_live_count_proto_rawDescGZIP(), []int{2}
}

func (x *WebFootMatchCount) GetPossession() int64 {
	if x != nil {
		return x.Possession
	}
	return 0
}

func (x *WebFootMatchCount) GetShots() int64 {
	if x != nil {
		return x.Shots
	}
	return 0
}

func (x *WebFootMatchCount) GetShotsOnTarget() int64 {
	if x != nil {
		return x.ShotsOnTarget
	}
	return 0
}

func (x *WebFootMatchCount) GetShotsOnTargetProportion() string {
	if x != nil {
		return x.ShotsOnTargetProportion
	}
	return ""
}

func (x *WebFootMatchCount) GetAttacks() int64 {
	if x != nil {
		return x.Attacks
	}
	return 0
}

func (x *WebFootMatchCount) GetDangerousAttacks() int64 {
	if x != nil {
		return x.DangerousAttacks
	}
	return 0
}

func (x *WebFootMatchCount) GetYellowCards() int64 {
	if x != nil {
		return x.YellowCards
	}
	return 0
}

func (x *WebFootMatchCount) GetRedCards() int64 {
	if x != nil {
		return x.RedCards
	}
	return 0
}

func (x *WebFootMatchCount) GetFreeKicks() int64 {
	if x != nil {
		return x.FreeKicks
	}
	return 0
}

func (x *WebFootMatchCount) GetCorners() int64 {
	if x != nil {
		return x.Corners
	}
	return 0
}

var File_web_foot_match_live_count_proto protoreflect.FileDescriptor

var file_web_foot_match_live_count_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x77, 0x65, 0x62, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x5f, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x50, 0x0a, 0x18, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x22, 0xd9, 0x02, 0x0a, 0x19, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61,
	0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61,
	0x6d, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4d, 0x61, 0x69,
	0x6e, 0x54, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73,
	0x74, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43,
	0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x75, 0x73,
	0x74, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x43, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x43, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67,
	0x6f, 0x12, 0x26, 0x0a, 0x04, 0x48, 0x6f, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x04, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x41, 0x77, 0x61,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f,
	0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x04, 0x41, 0x77, 0x61,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x22,
	0xe5, 0x02, 0x0a, 0x11, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x50, 0x6f, 0x73, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x68, 0x6f, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x53, 0x68, 0x6f, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x53,
	0x68, 0x6f, 0x74, 0x73, 0x4f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x53, 0x68, 0x6f, 0x74, 0x73, 0x4f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x12, 0x38, 0x0a, 0x17, 0x53, 0x68, 0x6f, 0x74, 0x73, 0x4f, 0x6e, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x17, 0x53, 0x68, 0x6f, 0x74, 0x73, 0x4f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x44, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x6f,
	0x75, 0x73, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x10, 0x44, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x6f, 0x75, 0x73, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x59, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x61, 0x72, 0x64, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x59, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x64, 0x43, 0x61, 0x72, 0x64, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x52, 0x65, 0x64, 0x43, 0x61, 0x72, 0x64, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x46, 0x72, 0x65, 0x65, 0x4b, 0x69, 0x63, 0x6b, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x46, 0x72, 0x65, 0x65, 0x4b, 0x69, 0x63, 0x6b, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x43, 0x6f, 0x72, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x43, 0x6f, 0x72, 0x6e, 0x65, 0x72, 0x73, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web_foot_match_live_count_proto_rawDescOnce sync.Once
	file_web_foot_match_live_count_proto_rawDescData = file_web_foot_match_live_count_proto_rawDesc
)

func file_web_foot_match_live_count_proto_rawDescGZIP() []byte {
	file_web_foot_match_live_count_proto_rawDescOnce.Do(func() {
		file_web_foot_match_live_count_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_foot_match_live_count_proto_rawDescData)
	})
	return file_web_foot_match_live_count_proto_rawDescData
}

var file_web_foot_match_live_count_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_web_foot_match_live_count_proto_goTypes = []interface{}{
	(*WebFootMatchCountRequest)(nil),  // 0: WebFootMatchCountRequest
	(*WebFootMatchCountResponse)(nil), // 1: WebFootMatchCountResponse
	(*WebFootMatchCount)(nil),         // 2: WebFootMatchCount
}
var file_web_foot_match_live_count_proto_depIdxs = []int32{
	2, // 0: WebFootMatchCountResponse.Home:type_name -> WebFootMatchCount
	2, // 1: WebFootMatchCountResponse.Away:type_name -> WebFootMatchCount
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_web_foot_match_live_count_proto_init() }
func file_web_foot_match_live_count_proto_init() {
	if File_web_foot_match_live_count_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_foot_match_live_count_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootMatchCountRequest); i {
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
		file_web_foot_match_live_count_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootMatchCountResponse); i {
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
		file_web_foot_match_live_count_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootMatchCount); i {
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
			RawDescriptor: file_web_foot_match_live_count_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_foot_match_live_count_proto_goTypes,
		DependencyIndexes: file_web_foot_match_live_count_proto_depIdxs,
		MessageInfos:      file_web_foot_match_live_count_proto_msgTypes,
	}.Build()
	File_web_foot_match_live_count_proto = out.File
	file_web_foot_match_live_count_proto_rawDesc = nil
	file_web_foot_match_live_count_proto_goTypes = nil
	file_web_foot_match_live_count_proto_depIdxs = nil
}
