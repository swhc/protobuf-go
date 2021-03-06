// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: v2_foot_match_odds_detail.proto

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

//v2 即时比赛列表 - 比赛详情 - 指数列表 - 指数详情
type V2FootMatchOddsDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId   int64  `protobuf:"varint,2,opt,name=eventId,proto3" json:"eventId,omitempty"`     //比赛ID
	Id        int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`               //请求的前一条ID
	PageSize  int64  `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize,omitempty"`   //页大小
	CompanyId int64  `protobuf:"varint,5,opt,name=companyId,proto3" json:"companyId,omitempty"` //公司ID
	PanId     int64  `protobuf:"varint,6,opt,name=panId,proto3" json:"panId,omitempty"`         //亚盘/欧盘/大小球ID
	Language  string `protobuf:"bytes,7,opt,name=language,proto3" json:"language,omitempty"`    //请求语言
}

func (x *V2FootMatchOddsDetailRequest) Reset() {
	*x = V2FootMatchOddsDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_match_odds_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchOddsDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchOddsDetailRequest) ProtoMessage() {}

func (x *V2FootMatchOddsDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_match_odds_detail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchOddsDetailRequest.ProtoReflect.Descriptor instead.
func (*V2FootMatchOddsDetailRequest) Descriptor() ([]byte, []int) {
	return file_v2_foot_match_odds_detail_proto_rawDescGZIP(), []int{0}
}

func (x *V2FootMatchOddsDetailRequest) GetEventId() int64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *V2FootMatchOddsDetailRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *V2FootMatchOddsDetailRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *V2FootMatchOddsDetailRequest) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *V2FootMatchOddsDetailRequest) GetPanId() int64 {
	if x != nil {
		return x.PanId
	}
	return 0
}

func (x *V2FootMatchOddsDetailRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type V2FootOddDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GunId  int64                  `protobuf:"varint,1,opt,name=gunId,proto3" json:"gunId,omitempty"`  //滚盘ID
	GunStr string                 `protobuf:"bytes,2,opt,name=gunStr,proto3" json:"gunStr,omitempty"` //滚盘字符串
	ChuId  int64                  `protobuf:"varint,3,opt,name=chuId,proto3" json:"chuId,omitempty"`  //初盘ID
	ChuStr string                 `protobuf:"bytes,4,opt,name=chuStr,proto3" json:"chuStr,omitempty"` //初盘字符串
	Count  int64                  `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`  //总数
	Info   []*V2FootOddDetailInfo `protobuf:"bytes,6,rep,name=info,proto3" json:"info,omitempty"`
}

func (x *V2FootOddDetailResponse) Reset() {
	*x = V2FootOddDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_match_odds_detail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootOddDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootOddDetailResponse) ProtoMessage() {}

func (x *V2FootOddDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_match_odds_detail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootOddDetailResponse.ProtoReflect.Descriptor instead.
func (*V2FootOddDetailResponse) Descriptor() ([]byte, []int) {
	return file_v2_foot_match_odds_detail_proto_rawDescGZIP(), []int{1}
}

func (x *V2FootOddDetailResponse) GetGunId() int64 {
	if x != nil {
		return x.GunId
	}
	return 0
}

func (x *V2FootOddDetailResponse) GetGunStr() string {
	if x != nil {
		return x.GunStr
	}
	return ""
}

func (x *V2FootOddDetailResponse) GetChuId() int64 {
	if x != nil {
		return x.ChuId
	}
	return 0
}

func (x *V2FootOddDetailResponse) GetChuStr() string {
	if x != nil {
		return x.ChuStr
	}
	return ""
}

func (x *V2FootOddDetailResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *V2FootOddDetailResponse) GetInfo() []*V2FootOddDetailInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type V2FootOddDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Column1     string `protobuf:"bytes,1,opt,name=column1,proto3" json:"column1,omitempty"`         //第1列数据 亚盘：主胜 欧盘：主胜 大小球：大
	Column2     string `protobuf:"bytes,2,opt,name=column2,proto3" json:"column2,omitempty"`         //第2列数据 亚盘：盘口 欧盘：平 大小球：盘口
	Column3     string `protobuf:"bytes,3,opt,name=column3,proto3" json:"column3,omitempty"`         //第3列数据 亚盘：客胜 欧盘：客胜 大小球：小
	UpdateTime  int64  `protobuf:"varint,4,opt,name=updateTime,proto3" json:"updateTime,omitempty"`  //变化时间
	Id          int64  `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`                  //数据ID
	CompanyName string `protobuf:"bytes,6,opt,name=companyName,proto3" json:"companyName,omitempty"` //公司名称
}

func (x *V2FootOddDetailInfo) Reset() {
	*x = V2FootOddDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_foot_match_odds_detail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootOddDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootOddDetailInfo) ProtoMessage() {}

func (x *V2FootOddDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v2_foot_match_odds_detail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootOddDetailInfo.ProtoReflect.Descriptor instead.
func (*V2FootOddDetailInfo) Descriptor() ([]byte, []int) {
	return file_v2_foot_match_odds_detail_proto_rawDescGZIP(), []int{2}
}

func (x *V2FootOddDetailInfo) GetColumn1() string {
	if x != nil {
		return x.Column1
	}
	return ""
}

func (x *V2FootOddDetailInfo) GetColumn2() string {
	if x != nil {
		return x.Column2
	}
	return ""
}

func (x *V2FootOddDetailInfo) GetColumn3() string {
	if x != nil {
		return x.Column3
	}
	return ""
}

func (x *V2FootOddDetailInfo) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *V2FootOddDetailInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *V2FootOddDetailInfo) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

var File_v2_foot_match_odds_detail_proto protoreflect.FileDescriptor

var file_v2_foot_match_odds_detail_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x76, 0x32, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x6f, 0x64, 0x64, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb4, 0x01, 0x0a, 0x1c, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x4f, 0x64, 0x64, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x6e, 0x49, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0xb5, 0x01, 0x0a, 0x17, 0x56, 0x32, 0x46,
	0x6f, 0x6f, 0x74, 0x4f, 0x64, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x75, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x67, 0x75, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x75,
	0x6e, 0x53, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x75, 0x6e, 0x53,
	0x74, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x63, 0x68, 0x75, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x75, 0x53,
	0x74, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x75, 0x53, 0x74, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4f, 0x64, 0x64,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x22, 0xb5, 0x01, 0x0a, 0x13, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4f, 0x64, 0x64, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x32, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x32, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x33, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_foot_match_odds_detail_proto_rawDescOnce sync.Once
	file_v2_foot_match_odds_detail_proto_rawDescData = file_v2_foot_match_odds_detail_proto_rawDesc
)

func file_v2_foot_match_odds_detail_proto_rawDescGZIP() []byte {
	file_v2_foot_match_odds_detail_proto_rawDescOnce.Do(func() {
		file_v2_foot_match_odds_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_foot_match_odds_detail_proto_rawDescData)
	})
	return file_v2_foot_match_odds_detail_proto_rawDescData
}

var file_v2_foot_match_odds_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v2_foot_match_odds_detail_proto_goTypes = []interface{}{
	(*V2FootMatchOddsDetailRequest)(nil), // 0: V2FootMatchOddsDetailRequest
	(*V2FootOddDetailResponse)(nil),      // 1: V2FootOddDetailResponse
	(*V2FootOddDetailInfo)(nil),          // 2: V2FootOddDetailInfo
}
var file_v2_foot_match_odds_detail_proto_depIdxs = []int32{
	2, // 0: V2FootOddDetailResponse.info:type_name -> V2FootOddDetailInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v2_foot_match_odds_detail_proto_init() }
func file_v2_foot_match_odds_detail_proto_init() {
	if File_v2_foot_match_odds_detail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_foot_match_odds_detail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchOddsDetailRequest); i {
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
		file_v2_foot_match_odds_detail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootOddDetailResponse); i {
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
		file_v2_foot_match_odds_detail_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootOddDetailInfo); i {
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
			RawDescriptor: file_v2_foot_match_odds_detail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_foot_match_odds_detail_proto_goTypes,
		DependencyIndexes: file_v2_foot_match_odds_detail_proto_depIdxs,
		MessageInfos:      file_v2_foot_match_odds_detail_proto_msgTypes,
	}.Build()
	File_v2_foot_match_odds_detail_proto = out.File
	file_v2_foot_match_odds_detail_proto_rawDesc = nil
	file_v2_foot_match_odds_detail_proto_goTypes = nil
	file_v2_foot_match_odds_detail_proto_depIdxs = nil
}
