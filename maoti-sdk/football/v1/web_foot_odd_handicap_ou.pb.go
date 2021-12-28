// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: web_foot_odd_handicap_ou.proto

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
type WebOddHandicapOURequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OddType  int64  `protobuf:"varint,1,opt,name=oddType,proto3" json:"oddType,omitempty"`  //odd  1压盘 2大小球
	EventId  int64  `protobuf:"varint,2,opt,name=eventId,proto3" json:"eventId,omitempty"`  //比赛id
	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"` //请求语言
}

func (x *WebOddHandicapOURequest) Reset() {
	*x = WebOddHandicapOURequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_odd_handicap_ou_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebOddHandicapOURequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebOddHandicapOURequest) ProtoMessage() {}

func (x *WebOddHandicapOURequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_odd_handicap_ou_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebOddHandicapOURequest.ProtoReflect.Descriptor instead.
func (*WebOddHandicapOURequest) Descriptor() ([]byte, []int) {
	return file_web_foot_odd_handicap_ou_proto_rawDescGZIP(), []int{0}
}

func (x *WebOddHandicapOURequest) GetOddType() int64 {
	if x != nil {
		return x.OddType
	}
	return 0
}

func (x *WebOddHandicapOURequest) GetEventId() int64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *WebOddHandicapOURequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type WebOddHandicapOUResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*WebOddHandicapOUList `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *WebOddHandicapOUResponse) Reset() {
	*x = WebOddHandicapOUResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_odd_handicap_ou_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebOddHandicapOUResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebOddHandicapOUResponse) ProtoMessage() {}

func (x *WebOddHandicapOUResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_odd_handicap_ou_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebOddHandicapOUResponse.ProtoReflect.Descriptor instead.
func (*WebOddHandicapOUResponse) Descriptor() ([]byte, []int) {
	return file_web_foot_odd_handicap_ou_proto_rawDescGZIP(), []int{1}
}

func (x *WebOddHandicapOUResponse) GetList() []*WebOddHandicapOUList {
	if x != nil {
		return x.List
	}
	return nil
}

type WebOddHandicapOUList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyName string   `protobuf:"bytes,5,opt,name=companyName,proto3" json:"companyName,omitempty"` //公司名称
	InitOdds    *OddInfo `protobuf:"bytes,1,opt,name=InitOdds,proto3" json:"InitOdds,omitempty"`
	CurrOdds    *OddInfo `protobuf:"bytes,3,opt,name=CurrOdds,proto3" json:"CurrOdds,omitempty"`
	Times       int64    `protobuf:"varint,6,opt,name=times,proto3" json:"times,omitempty"`
}

func (x *WebOddHandicapOUList) Reset() {
	*x = WebOddHandicapOUList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_odd_handicap_ou_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebOddHandicapOUList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebOddHandicapOUList) ProtoMessage() {}

func (x *WebOddHandicapOUList) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_odd_handicap_ou_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebOddHandicapOUList.ProtoReflect.Descriptor instead.
func (*WebOddHandicapOUList) Descriptor() ([]byte, []int) {
	return file_web_foot_odd_handicap_ou_proto_rawDescGZIP(), []int{2}
}

func (x *WebOddHandicapOUList) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *WebOddHandicapOUList) GetInitOdds() *OddInfo {
	if x != nil {
		return x.InitOdds
	}
	return nil
}

func (x *WebOddHandicapOUList) GetCurrOdds() *OddInfo {
	if x != nil {
		return x.CurrOdds
	}
	return nil
}

func (x *WebOddHandicapOUList) GetTimes() int64 {
	if x != nil {
		return x.Times
	}
	return 0
}

type OddInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Win  float32 `protobuf:"fixed32,1,opt,name=win,proto3" json:"win,omitempty"`   //胜
	Draw float32 `protobuf:"fixed32,2,opt,name=draw,proto3" json:"draw,omitempty"` //平
	Lose float32 `protobuf:"fixed32,3,opt,name=lose,proto3" json:"lose,omitempty"` //负
}

func (x *OddInfo) Reset() {
	*x = OddInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_odd_handicap_ou_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OddInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OddInfo) ProtoMessage() {}

func (x *OddInfo) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_odd_handicap_ou_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OddInfo.ProtoReflect.Descriptor instead.
func (*OddInfo) Descriptor() ([]byte, []int) {
	return file_web_foot_odd_handicap_ou_proto_rawDescGZIP(), []int{3}
}

func (x *OddInfo) GetWin() float32 {
	if x != nil {
		return x.Win
	}
	return 0
}

func (x *OddInfo) GetDraw() float32 {
	if x != nil {
		return x.Draw
	}
	return 0
}

func (x *OddInfo) GetLose() float32 {
	if x != nil {
		return x.Lose
	}
	return 0
}

var File_web_foot_odd_handicap_ou_proto protoreflect.FileDescriptor

var file_web_foot_odd_handicap_ou_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x77, 0x65, 0x62, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6f, 0x64, 0x64, 0x5f, 0x68,
	0x61, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x70, 0x5f, 0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x69, 0x0a, 0x17, 0x57, 0x65, 0x62, 0x4f, 0x64, 0x64, 0x48, 0x61, 0x6e, 0x64, 0x69, 0x63,
	0x61, 0x70, 0x4f, 0x55, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x64, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x64,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x45, 0x0a, 0x18, 0x57,
	0x65, 0x62, 0x4f, 0x64, 0x64, 0x48, 0x61, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x70, 0x4f, 0x55, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x57, 0x65, 0x62, 0x4f, 0x64, 0x64, 0x48, 0x61,
	0x6e, 0x64, 0x69, 0x63, 0x61, 0x70, 0x4f, 0x55, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0x9a, 0x01, 0x0a, 0x14, 0x57, 0x65, 0x62, 0x4f, 0x64, 0x64, 0x48, 0x61, 0x6e,
	0x64, 0x69, 0x63, 0x61, 0x70, 0x4f, 0x55, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a,
	0x08, 0x49, 0x6e, 0x69, 0x74, 0x4f, 0x64, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x4f, 0x64, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x49, 0x6e, 0x69, 0x74, 0x4f,
	0x64, 0x64, 0x73, 0x12, 0x24, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x4f, 0x64, 0x64, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4f, 0x64, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x08, 0x43, 0x75, 0x72, 0x72, 0x4f, 0x64, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x22,
	0x43, 0x0a, 0x07, 0x4f, 0x64, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x77, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x72, 0x61, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x64, 0x72, 0x61, 0x77,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04,
	0x6c, 0x6f, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web_foot_odd_handicap_ou_proto_rawDescOnce sync.Once
	file_web_foot_odd_handicap_ou_proto_rawDescData = file_web_foot_odd_handicap_ou_proto_rawDesc
)

func file_web_foot_odd_handicap_ou_proto_rawDescGZIP() []byte {
	file_web_foot_odd_handicap_ou_proto_rawDescOnce.Do(func() {
		file_web_foot_odd_handicap_ou_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_foot_odd_handicap_ou_proto_rawDescData)
	})
	return file_web_foot_odd_handicap_ou_proto_rawDescData
}

var file_web_foot_odd_handicap_ou_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_web_foot_odd_handicap_ou_proto_goTypes = []interface{}{
	(*WebOddHandicapOURequest)(nil),  // 0: WebOddHandicapOURequest
	(*WebOddHandicapOUResponse)(nil), // 1: WebOddHandicapOUResponse
	(*WebOddHandicapOUList)(nil),     // 2: WebOddHandicapOUList
	(*OddInfo)(nil),                  // 3: OddInfo
}
var file_web_foot_odd_handicap_ou_proto_depIdxs = []int32{
	2, // 0: WebOddHandicapOUResponse.list:type_name -> WebOddHandicapOUList
	3, // 1: WebOddHandicapOUList.InitOdds:type_name -> OddInfo
	3, // 2: WebOddHandicapOUList.CurrOdds:type_name -> OddInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_web_foot_odd_handicap_ou_proto_init() }
func file_web_foot_odd_handicap_ou_proto_init() {
	if File_web_foot_odd_handicap_ou_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_foot_odd_handicap_ou_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebOddHandicapOURequest); i {
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
		file_web_foot_odd_handicap_ou_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebOddHandicapOUResponse); i {
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
		file_web_foot_odd_handicap_ou_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebOddHandicapOUList); i {
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
		file_web_foot_odd_handicap_ou_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OddInfo); i {
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
			RawDescriptor: file_web_foot_odd_handicap_ou_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_foot_odd_handicap_ou_proto_goTypes,
		DependencyIndexes: file_web_foot_odd_handicap_ou_proto_depIdxs,
		MessageInfos:      file_web_foot_odd_handicap_ou_proto_msgTypes,
	}.Build()
	File_web_foot_odd_handicap_ou_proto = out.File
	file_web_foot_odd_handicap_ou_proto_rawDesc = nil
	file_web_foot_odd_handicap_ou_proto_goTypes = nil
	file_web_foot_odd_handicap_ou_proto_depIdxs = nil
}
