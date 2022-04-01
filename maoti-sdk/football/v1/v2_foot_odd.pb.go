// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: v2_foot_odd.proto

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

type V2FootMatchOddListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SportId  int64  `protobuf:"varint,1,opt,name=sportId,proto3" json:"sportId,omitempty"`  //1足球 2篮球
	EventId  int64  `protobuf:"varint,2,opt,name=eventId,proto3" json:"eventId,omitempty"`  //比赛id
	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"` //请求语言
}

func (x *V2FootMatchOddListRequest) Reset() {
	*x = V2FootMatchOddListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_odd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchOddListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchOddListRequest) ProtoMessage() {}

func (x *V2FootMatchOddListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_odd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchOddListRequest.ProtoReflect.Descriptor instead.
func (*V2FootMatchOddListRequest) Descriptor() ([]byte, []int) {
	return file_v2_foot_odd_proto_rawDescGZIP(), []int{0}
}

func (x *V2FootMatchOddListRequest) GetSportId() int64 {
	if x != nil {
		return x.SportId
	}
	return 0
}

func (x *V2FootMatchOddListRequest) GetEventId() int64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *V2FootMatchOddListRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type V2FootMatchOddListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*V2FootMatchOddData `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *V2FootMatchOddListResponse) Reset() {
	*x = V2FootMatchOddListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_odd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchOddListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchOddListResponse) ProtoMessage() {}

func (x *V2FootMatchOddListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_odd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchOddListResponse.ProtoReflect.Descriptor instead.
func (*V2FootMatchOddListResponse) Descriptor() ([]byte, []int) {
	return file_v2_foot_odd_proto_rawDescGZIP(), []int{1}
}

func (x *V2FootMatchOddListResponse) GetData() []*V2FootMatchOddData {
	if x != nil {
		return x.Data
	}
	return nil
}

type V2FootMatchOddData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` //tab名称
	Id     int64                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`    //tab id 1亚盘 2欧盘 3大小球
	Titles []string              `protobuf:"bytes,3,rep,name=titles,proto3" json:"titles,omitempty"`
	List   []*V2FootMatchOddList `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *V2FootMatchOddData) Reset() {
	*x = V2FootMatchOddData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_odd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchOddData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchOddData) ProtoMessage() {}

func (x *V2FootMatchOddData) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_odd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchOddData.ProtoReflect.Descriptor instead.
func (*V2FootMatchOddData) Descriptor() ([]byte, []int) {
	return file_v2_foot_odd_proto_rawDescGZIP(), []int{2}
}

func (x *V2FootMatchOddData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *V2FootMatchOddData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *V2FootMatchOddData) GetTitles() []string {
	if x != nil {
		return x.Titles
	}
	return nil
}

func (x *V2FootMatchOddData) GetList() []*V2FootMatchOddList {
	if x != nil {
		return x.List
	}
	return nil
}

type V2FootMatchOddList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`            //公司id
	Name     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`         //公司名称
	Mark     int64    `protobuf:"varint,5,opt,name=mark,proto3" json:"mark,omitempty"`        //初盘最晚和最早标记；1.最早；2.最晚；仅供初盘数据使用
	NowData  []string `protobuf:"bytes,3,rep,name=nowData,proto3" json:"nowData,omitempty"`   //当前数据
	InitData []string `protobuf:"bytes,4,rep,name=initData,proto3" json:"initData,omitempty"` //初始数据
}

func (x *V2FootMatchOddList) Reset() {
	*x = V2FootMatchOddList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_odd_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchOddList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchOddList) ProtoMessage() {}

func (x *V2FootMatchOddList) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_odd_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchOddList.ProtoReflect.Descriptor instead.
func (*V2FootMatchOddList) Descriptor() ([]byte, []int) {
	return file_v2_foot_odd_proto_rawDescGZIP(), []int{3}
}

func (x *V2FootMatchOddList) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *V2FootMatchOddList) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *V2FootMatchOddList) GetMark() int64 {
	if x != nil {
		return x.Mark
	}
	return 0
}

func (x *V2FootMatchOddList) GetNowData() []string {
	if x != nil {
		return x.NowData
	}
	return nil
}

func (x *V2FootMatchOddList) GetInitData() []string {
	if x != nil {
		return x.InitData
	}
	return nil
}

var File_v2_foot_odd_proto protoreflect.FileDescriptor

var file_v2_foot_odd_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x32, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6f, 0x64, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x19, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x4f, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x22, 0x45, 0x0a, 0x1a, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4f,
	0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x56,
	0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x64, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x79, 0x0a, 0x12, 0x56, 0x32, 0x46, 0x6f, 0x6f,
	0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x64, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x12, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x4f, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6d, 0x61, 0x72,
	0x6b, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x6e, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x6e, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x6e, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_foot_odd_proto_rawDescOnce sync.Once
	file_v2_foot_odd_proto_rawDescData = file_v2_foot_odd_proto_rawDesc
)

func file_v2_foot_odd_proto_rawDescGZIP() []byte {
	file_v2_foot_odd_proto_rawDescOnce.Do(func() {
		file_v2_foot_odd_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_foot_odd_proto_rawDescData)
	})
	return file_v2_foot_odd_proto_rawDescData
}

var file_v2_foot_odd_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v2_foot_odd_proto_goTypes = []interface{}{
	(*V2FootMatchOddListRequest)(nil),  // 0: V2FootMatchOddListRequest
	(*V2FootMatchOddListResponse)(nil), // 1: V2FootMatchOddListResponse
	(*V2FootMatchOddData)(nil),         // 2: V2FootMatchOddData
	(*V2FootMatchOddList)(nil),         // 3: V2FootMatchOddList
}
var file_v2_foot_odd_proto_depIdxs = []int32{
	2, // 0: V2FootMatchOddListResponse.data:type_name -> V2FootMatchOddData
	3, // 1: V2FootMatchOddData.list:type_name -> V2FootMatchOddList
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v2_foot_odd_proto_init() }
func file_v2_foot_odd_proto_init() {
	if File_v2_foot_odd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_foot_odd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchOddListRequest); i {
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
		file_v2_foot_odd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchOddListResponse); i {
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
		file_v2_foot_odd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchOddData); i {
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
		file_v2_foot_odd_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchOddList); i {
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
			RawDescriptor: file_v2_foot_odd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_foot_odd_proto_goTypes,
		DependencyIndexes: file_v2_foot_odd_proto_depIdxs,
		MessageInfos:      file_v2_foot_odd_proto_msgTypes,
	}.Build()
	File_v2_foot_odd_proto = out.File
	file_v2_foot_odd_proto_rawDesc = nil
	file_v2_foot_odd_proto_goTypes = nil
	file_v2_foot_odd_proto_depIdxs = nil
}
