// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: collect_v2.proto

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

//游客收藏数量不能超过10个，还有取消收藏
type CollectRequestParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids            []int64 `protobuf:"varint,1,rep,packed,name=Ids,proto3" json:"Ids,omitempty"`               //数据列表
	UserId         int64   `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`                //用户id
	ActionType     int64   `protobuf:"varint,3,opt,name=ActionType,proto3" json:"ActionType,omitempty"`        // 1-收藏 2-取消收藏
	Type           int64   `protobuf:"varint,4,opt,name=Type,proto3" json:"Type,omitempty"`                    // 数据类型 1：赛事，2：比赛，3：队伍，4：专家
	IsGuest        int64   `protobuf:"varint,5,opt,name=IsGuest,proto3" json:"IsGuest,omitempty"`              //是否游客  0：否，1：是
	SportId        int64   `protobuf:"varint,6,opt,name=SportId,proto3" json:"SportId,omitempty"`              //运动类型 1:足球，2：篮球
	TournamentName string  `protobuf:"bytes,7,opt,name=TournamentName,proto3" json:"TournamentName,omitempty"` //赛事名称
	MatchTime      int64   `protobuf:"varint,8,opt,name=MatchTime,proto3" json:"MatchTime,omitempty"`          //比赛时间
}

func (x *CollectRequestParam) Reset() {
	*x = CollectRequestParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_v2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectRequestParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectRequestParam) ProtoMessage() {}

func (x *CollectRequestParam) ProtoReflect() protoreflect.Message {
	mi := &file_collect_v2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectRequestParam.ProtoReflect.Descriptor instead.
func (*CollectRequestParam) Descriptor() ([]byte, []int) {
	return file_collect_v2_proto_rawDescGZIP(), []int{0}
}

func (x *CollectRequestParam) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *CollectRequestParam) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CollectRequestParam) GetActionType() int64 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

func (x *CollectRequestParam) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *CollectRequestParam) GetIsGuest() int64 {
	if x != nil {
		return x.IsGuest
	}
	return 0
}

func (x *CollectRequestParam) GetSportId() int64 {
	if x != nil {
		return x.SportId
	}
	return 0
}

func (x *CollectRequestParam) GetTournamentName() string {
	if x != nil {
		return x.TournamentName
	}
	return ""
}

func (x *CollectRequestParam) GetMatchTime() int64 {
	if x != nil {
		return x.MatchTime
	}
	return 0
}

type CollectListRequestParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`     //用户id
	Type     int64 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`         // 数据类型 1：赛事，2：比赛，3：队伍，4：专家
	IsGuest  int64 `protobuf:"varint,5,opt,name=IsGuest,proto3" json:"IsGuest,omitempty"`   //是否游客  0：否，1：是
	SportId  int64 `protobuf:"varint,6,opt,name=SportId,proto3" json:"SportId,omitempty"`   //运动类型 1:足球，2：篮球
	Page     int64 `protobuf:"varint,7,opt,name=Page,proto3" json:"Page,omitempty"`         //页码
	PageSize int64 `protobuf:"varint,8,opt,name=PageSize,proto3" json:"PageSize,omitempty"` //页数
}

func (x *CollectListRequestParam) Reset() {
	*x = CollectListRequestParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_v2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectListRequestParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectListRequestParam) ProtoMessage() {}

func (x *CollectListRequestParam) ProtoReflect() protoreflect.Message {
	mi := &file_collect_v2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectListRequestParam.ProtoReflect.Descriptor instead.
func (*CollectListRequestParam) Descriptor() ([]byte, []int) {
	return file_collect_v2_proto_rawDescGZIP(), []int{1}
}

func (x *CollectListRequestParam) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CollectListRequestParam) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *CollectListRequestParam) GetIsGuest() int64 {
	if x != nil {
		return x.IsGuest
	}
	return 0
}

func (x *CollectListRequestParam) GetSportId() int64 {
	if x != nil {
		return x.SportId
	}
	return 0
}

func (x *CollectListRequestParam) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *CollectListRequestParam) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type PushSettingRequestParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64             `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`                                                                                                  //用户id
	UpField map[string]string `protobuf:"bytes,2,rep,name=UpField,proto3" json:"UpField,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //注册用户要更新的字段
	Tourist map[string]string `protobuf:"bytes,3,rep,name=Tourist,proto3" json:"Tourist,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //游客要更新的字段
	IsLogin int64             `protobuf:"varint,4,opt,name=IsLogin,proto3" json:"IsLogin,omitempty"`                                                                                        //是否游客  0：否，1：是
	Device  string            `protobuf:"bytes,5,opt,name=Device,proto3" json:"Device,omitempty"`                                                                                           //设备
	Token   string            `protobuf:"bytes,6,opt,name=Token,proto3" json:"Token,omitempty"`                                                                                             //游客token
}

func (x *PushSettingRequestParam) Reset() {
	*x = PushSettingRequestParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_v2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushSettingRequestParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushSettingRequestParam) ProtoMessage() {}

func (x *PushSettingRequestParam) ProtoReflect() protoreflect.Message {
	mi := &file_collect_v2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushSettingRequestParam.ProtoReflect.Descriptor instead.
func (*PushSettingRequestParam) Descriptor() ([]byte, []int) {
	return file_collect_v2_proto_rawDescGZIP(), []int{2}
}

func (x *PushSettingRequestParam) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PushSettingRequestParam) GetUpField() map[string]string {
	if x != nil {
		return x.UpField
	}
	return nil
}

func (x *PushSettingRequestParam) GetTourist() map[string]string {
	if x != nil {
		return x.Tourist
	}
	return nil
}

func (x *PushSettingRequestParam) GetIsLogin() int64 {
	if x != nil {
		return x.IsLogin
	}
	return 0
}

func (x *PushSettingRequestParam) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *PushSettingRequestParam) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_collect_v2_proto protoreflect.FileDescriptor

var file_collect_v2_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xed, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x49, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x73, 0x47, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x49, 0x73, 0x47, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e,
	0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x17, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x73,
	0x47, 0x75, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x49, 0x73, 0x47,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50, 0x61,
	0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xeb,
	0x02, 0x0a, 0x17, 0x50, 0x75, 0x73, 0x68, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x07, 0x55, 0x70,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x50, 0x75,
	0x73, 0x68, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x55, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x55, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x3f, 0x0a, 0x07, 0x54,
	0x6f, 0x75, 0x72, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x50,
	0x75, 0x73, 0x68, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x69, 0x73, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x54, 0x6f, 0x75, 0x72, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x49, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x49,
	0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x3a, 0x0a, 0x0c, 0x55, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x3a, 0x0a, 0x0c, 0x54, 0x6f, 0x75, 0x72, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_collect_v2_proto_rawDescOnce sync.Once
	file_collect_v2_proto_rawDescData = file_collect_v2_proto_rawDesc
)

func file_collect_v2_proto_rawDescGZIP() []byte {
	file_collect_v2_proto_rawDescOnce.Do(func() {
		file_collect_v2_proto_rawDescData = protoimpl.X.CompressGZIP(file_collect_v2_proto_rawDescData)
	})
	return file_collect_v2_proto_rawDescData
}

var file_collect_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_collect_v2_proto_goTypes = []interface{}{
	(*CollectRequestParam)(nil),     // 0: CollectRequestParam
	(*CollectListRequestParam)(nil), // 1: CollectListRequestParam
	(*PushSettingRequestParam)(nil), // 2: PushSettingRequestParam
	nil,                             // 3: PushSettingRequestParam.UpFieldEntry
	nil,                             // 4: PushSettingRequestParam.TouristEntry
}
var file_collect_v2_proto_depIdxs = []int32{
	3, // 0: PushSettingRequestParam.UpField:type_name -> PushSettingRequestParam.UpFieldEntry
	4, // 1: PushSettingRequestParam.Tourist:type_name -> PushSettingRequestParam.TouristEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_collect_v2_proto_init() }
func file_collect_v2_proto_init() {
	if File_collect_v2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_collect_v2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectRequestParam); i {
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
		file_collect_v2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectListRequestParam); i {
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
		file_collect_v2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushSettingRequestParam); i {
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
			RawDescriptor: file_collect_v2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_collect_v2_proto_goTypes,
		DependencyIndexes: file_collect_v2_proto_depIdxs,
		MessageInfos:      file_collect_v2_proto_msgTypes,
	}.Build()
	File_collect_v2_proto = out.File
	file_collect_v2_proto_rawDesc = nil
	file_collect_v2_proto_goTypes = nil
	file_collect_v2_proto_depIdxs = nil
}
