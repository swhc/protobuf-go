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

type CollectRequestParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CId        string         `protobuf:"bytes,1,opt,name=CId,proto3" json:"CId,omitempty"` //设备标识Id
	Collection []*CollectItem `protobuf:"bytes,2,rep,name=Collection,proto3" json:"Collection,omitempty"`
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

func (x *CollectRequestParam) GetCId() string {
	if x != nil {
		return x.CId
	}
	return ""
}

func (x *CollectRequestParam) GetCollection() []*CollectItem {
	if x != nil {
		return x.Collection
	}
	return nil
}

type CollectItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string  `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`       // 数据类型 match;fb_tournament;bb_tournament;fb_team;bb_team;fb_player;bb_player;expert;plans
	Ids  []int64 `protobuf:"varint,2,rep,packed,name=Ids,proto3" json:"Ids,omitempty"` //  收藏的id 正数为收藏 负数为取消
}

func (x *CollectItem) Reset() {
	*x = CollectItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_v2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectItem) ProtoMessage() {}

func (x *CollectItem) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CollectItem.ProtoReflect.Descriptor instead.
func (*CollectItem) Descriptor() ([]byte, []int) {
	return file_collect_v2_proto_rawDescGZIP(), []int{1}
}

func (x *CollectItem) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CollectItem) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type CollectListRequestParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CId  string `protobuf:"bytes,1,opt,name=CId,proto3" json:"CId,omitempty"`   //设备标识Id
	Type string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"` // 筛选项 空值全量返回
}

func (x *CollectListRequestParam) Reset() {
	*x = CollectListRequestParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_v2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectListRequestParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectListRequestParam) ProtoMessage() {}

func (x *CollectListRequestParam) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CollectListRequestParam.ProtoReflect.Descriptor instead.
func (*CollectListRequestParam) Descriptor() ([]byte, []int) {
	return file_collect_v2_proto_rawDescGZIP(), []int{2}
}

func (x *CollectListRequestParam) GetCId() string {
	if x != nil {
		return x.CId
	}
	return ""
}

func (x *CollectListRequestParam) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type CollectListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collection []*CollectItem `protobuf:"bytes,1,rep,name=Collection,proto3" json:"Collection,omitempty"` //
}

func (x *CollectListResp) Reset() {
	*x = CollectListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_v2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectListResp) ProtoMessage() {}

func (x *CollectListResp) ProtoReflect() protoreflect.Message {
	mi := &file_collect_v2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectListResp.ProtoReflect.Descriptor instead.
func (*CollectListResp) Descriptor() ([]byte, []int) {
	return file_collect_v2_proto_rawDescGZIP(), []int{3}
}

func (x *CollectListResp) GetCollection() []*CollectItem {
	if x != nil {
		return x.Collection
	}
	return nil
}

type CollectCidUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid  string `protobuf:"bytes,1,opt,name=Cid,proto3" json:"Cid,omitempty"`   //设备标识Id
	Uuid string `protobuf:"bytes,2,opt,name=Uuid,proto3" json:"Uuid,omitempty"` // 登录后的标识id
}

func (x *CollectCidUpdateRequest) Reset() {
	*x = CollectCidUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_v2_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectCidUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectCidUpdateRequest) ProtoMessage() {}

func (x *CollectCidUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_collect_v2_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectCidUpdateRequest.ProtoReflect.Descriptor instead.
func (*CollectCidUpdateRequest) Descriptor() ([]byte, []int) {
	return file_collect_v2_proto_rawDescGZIP(), []int{4}
}

func (x *CollectCidUpdateRequest) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *CollectCidUpdateRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
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
		mi := &file_collect_v2_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushSettingRequestParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushSettingRequestParam) ProtoMessage() {}

func (x *PushSettingRequestParam) ProtoReflect() protoreflect.Message {
	mi := &file_collect_v2_proto_msgTypes[5]
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
	return file_collect_v2_proto_rawDescGZIP(), []int{5}
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
	0x74, 0x6f, 0x22, 0x55, 0x0a, 0x13, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x43, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x0a, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x33, 0x0a, 0x0b, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x49, 0x64, 0x73, 0x22, 0x3f,
	0x0a, 0x17, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x43, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x3f, 0x0a, 0x0f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x2c, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x3f, 0x0a, 0x17, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x69, 0x64, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x43,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x43, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x55, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x75, 0x69,
	0x64, 0x22, 0xeb, 0x02, 0x0a, 0x17, 0x50, 0x75, 0x73, 0x68, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x3f, 0x0a,
	0x07, 0x55, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x50, 0x75, 0x73, 0x68, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x55, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x55, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x3f,
	0x0a, 0x07, 0x54, 0x6f, 0x75, 0x72, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x69, 0x73,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x54, 0x6f, 0x75, 0x72, 0x69, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x49, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x49, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x3a, 0x0a, 0x0c, 0x55, 0x70, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x54, 0x6f, 0x75, 0x72, 0x69, 0x73, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_collect_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_collect_v2_proto_goTypes = []interface{}{
	(*CollectRequestParam)(nil),     // 0: CollectRequestParam
	(*CollectItem)(nil),             // 1: CollectItem
	(*CollectListRequestParam)(nil), // 2: CollectListRequestParam
	(*CollectListResp)(nil),         // 3: CollectListResp
	(*CollectCidUpdateRequest)(nil), // 4: CollectCidUpdateRequest
	(*PushSettingRequestParam)(nil), // 5: PushSettingRequestParam
	nil,                             // 6: PushSettingRequestParam.UpFieldEntry
	nil,                             // 7: PushSettingRequestParam.TouristEntry
}
var file_collect_v2_proto_depIdxs = []int32{
	1, // 0: CollectRequestParam.Collection:type_name -> CollectItem
	1, // 1: CollectListResp.Collection:type_name -> CollectItem
	6, // 2: PushSettingRequestParam.UpField:type_name -> PushSettingRequestParam.UpFieldEntry
	7, // 3: PushSettingRequestParam.Tourist:type_name -> PushSettingRequestParam.TouristEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
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
			switch v := v.(*CollectItem); i {
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
		file_collect_v2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectListResp); i {
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
		file_collect_v2_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectCidUpdateRequest); i {
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
		file_collect_v2_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   8,
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
