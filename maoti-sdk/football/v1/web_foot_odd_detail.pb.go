// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: web_foot_odd_detail.proto

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

//指数-详情
type WebOddDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId   int64  `protobuf:"varint,2,opt,name=eventId,proto3" json:"eventId,omitempty"`     //比赛id
	Language  string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`    //请求语言  1:zh  2:en
	CompanyID int64  `protobuf:"varint,4,opt,name=companyID,proto3" json:"companyID,omitempty"` //指数公司ID
	OddType   int64  `protobuf:"varint,5,opt,name=oddType,proto3" json:"oddType,omitempty"`     //指数类型 2.亚盘；1.欧赔；3.大小球
	TimeZone  int64  `protobuf:"varint,6,opt,name=TimeZone,proto3" json:"TimeZone,omitempty"`   //时区
}

func (x *WebOddDetailRequest) Reset() {
	*x = WebOddDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_odd_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebOddDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebOddDetailRequest) ProtoMessage() {}

func (x *WebOddDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_odd_detail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebOddDetailRequest.ProtoReflect.Descriptor instead.
func (*WebOddDetailRequest) Descriptor() ([]byte, []int) {
	return file_web_foot_odd_detail_proto_rawDescGZIP(), []int{0}
}

func (x *WebOddDetailRequest) GetEventId() int64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *WebOddDetailRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *WebOddDetailRequest) GetCompanyID() int64 {
	if x != nil {
		return x.CompanyID
	}
	return 0
}

func (x *WebOddDetailRequest) GetOddType() int64 {
	if x != nil {
		return x.OddType
	}
	return 0
}

func (x *WebOddDetailRequest) GetTimeZone() int64 {
	if x != nil {
		return x.TimeZone
	}
	return 0
}

type WebOddDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Home   string       `protobuf:"bytes,1,opt,name=home,proto3" json:"home,omitempty"` //主队名称；当指数类型为亚盘时该字段才有值，作为表头的主队名称列
	Away   string       `protobuf:"bytes,2,opt,name=away,proto3" json:"away,omitempty"` //客队名称；当指数类型为亚盘时该字段才有值，作为表头的客队名称列
	Detail []*OddDetail `protobuf:"bytes,3,rep,name=detail,proto3" json:"detail,omitempty"`
}

func (x *WebOddDetailResponse) Reset() {
	*x = WebOddDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_odd_detail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebOddDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebOddDetailResponse) ProtoMessage() {}

func (x *WebOddDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_odd_detail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebOddDetailResponse.ProtoReflect.Descriptor instead.
func (*WebOddDetailResponse) Descriptor() ([]byte, []int) {
	return file_web_foot_odd_detail_proto_rawDescGZIP(), []int{1}
}

func (x *WebOddDetailResponse) GetHome() string {
	if x != nil {
		return x.Home
	}
	return ""
}

func (x *WebOddDetailResponse) GetAway() string {
	if x != nil {
		return x.Away
	}
	return ""
}

func (x *WebOddDetailResponse) GetDetail() []*OddDetail {
	if x != nil {
		return x.Detail
	}
	return nil
}

type OddDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time        int64  `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`              //指数更新时间，也叫比赛时间
	Minute      int64  `protobuf:"varint,8,opt,name=minute,proto3" json:"minute,omitempty"`          //指数更新的分钟数，也叫比赛进行的分钟数
	HomeScore   int64  `protobuf:"varint,2,opt,name=homeScore,proto3" json:"homeScore,omitempty"`    //主队比分
	AwayScore   int64  `protobuf:"varint,3,opt,name=awayScore,proto3" json:"awayScore,omitempty"`    //客队比分
	MainBall    string `protobuf:"bytes,4,opt,name=MainBall,proto3" json:"MainBall,omitempty"`       //主胜/大球；当指数类型为亚盘或欧赔时，该字段为主胜；当指数类型为大小球时，该字段为大球
	PlateFlat   string `protobuf:"bytes,5,opt,name=PlateFlat,proto3" json:"PlateFlat,omitempty"`     //亚盘/平/盘口；当指数类型为亚盘时，该字段为亚盘；当指数类型为欧赔时，该字段为平；当指数类型为大小球时，该字段为盘口；
	GuestBall   string `protobuf:"bytes,6,opt,name=GuestBall,proto3" json:"GuestBall,omitempty"`     //客胜/小球；当指数类型为亚盘或欧赔时，该字段为客胜；当指数类型为大小球时，该字段为小球
	MatchStatus string `protobuf:"bytes,7,opt,name=matchStatus,proto3" json:"matchStatus,omitempty"` //比赛进行时状态；ongoing：进行中；front：赛前；lnitial：初始
}

func (x *OddDetail) Reset() {
	*x = OddDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_odd_detail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OddDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OddDetail) ProtoMessage() {}

func (x *OddDetail) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_odd_detail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OddDetail.ProtoReflect.Descriptor instead.
func (*OddDetail) Descriptor() ([]byte, []int) {
	return file_web_foot_odd_detail_proto_rawDescGZIP(), []int{2}
}

func (x *OddDetail) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *OddDetail) GetMinute() int64 {
	if x != nil {
		return x.Minute
	}
	return 0
}

func (x *OddDetail) GetHomeScore() int64 {
	if x != nil {
		return x.HomeScore
	}
	return 0
}

func (x *OddDetail) GetAwayScore() int64 {
	if x != nil {
		return x.AwayScore
	}
	return 0
}

func (x *OddDetail) GetMainBall() string {
	if x != nil {
		return x.MainBall
	}
	return ""
}

func (x *OddDetail) GetPlateFlat() string {
	if x != nil {
		return x.PlateFlat
	}
	return ""
}

func (x *OddDetail) GetGuestBall() string {
	if x != nil {
		return x.GuestBall
	}
	return ""
}

func (x *OddDetail) GetMatchStatus() string {
	if x != nil {
		return x.MatchStatus
	}
	return ""
}

var File_web_foot_odd_detail_proto protoreflect.FileDescriptor

var file_web_foot_odd_detail_proto_rawDesc = []byte{
	0x0a, 0x19, 0x77, 0x65, 0x62, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6f, 0x64, 0x64, 0x5f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x13,
	0x57, 0x65, 0x62, 0x4f, 0x64, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x64, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x64, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x62, 0x0a,
	0x14, 0x57, 0x65, 0x62, 0x4f, 0x64, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x77, 0x61,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x77, 0x61, 0x79, 0x12, 0x22, 0x0a,
	0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x4f, 0x64, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x22, 0xed, 0x01, 0x0a, 0x09, 0x4f, 0x64, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x68,
	0x6f, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x68, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x77, 0x61,
	0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x77,
	0x61, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x61, 0x69, 0x6e, 0x42,
	0x61, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x61, 0x69, 0x6e, 0x42,
	0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x6c, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x61,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x75, 0x65, 0x73, 0x74, 0x42, 0x61, 0x6c, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x75, 0x65, 0x73, 0x74, 0x42, 0x61, 0x6c, 0x6c, 0x12,
	0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_web_foot_odd_detail_proto_rawDescOnce sync.Once
	file_web_foot_odd_detail_proto_rawDescData = file_web_foot_odd_detail_proto_rawDesc
)

func file_web_foot_odd_detail_proto_rawDescGZIP() []byte {
	file_web_foot_odd_detail_proto_rawDescOnce.Do(func() {
		file_web_foot_odd_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_foot_odd_detail_proto_rawDescData)
	})
	return file_web_foot_odd_detail_proto_rawDescData
}

var file_web_foot_odd_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_web_foot_odd_detail_proto_goTypes = []interface{}{
	(*WebOddDetailRequest)(nil),  // 0: WebOddDetailRequest
	(*WebOddDetailResponse)(nil), // 1: WebOddDetailResponse
	(*OddDetail)(nil),            // 2: OddDetail
}
var file_web_foot_odd_detail_proto_depIdxs = []int32{
	2, // 0: WebOddDetailResponse.detail:type_name -> OddDetail
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_web_foot_odd_detail_proto_init() }
func file_web_foot_odd_detail_proto_init() {
	if File_web_foot_odd_detail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_foot_odd_detail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebOddDetailRequest); i {
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
		file_web_foot_odd_detail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebOddDetailResponse); i {
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
		file_web_foot_odd_detail_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OddDetail); i {
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
			RawDescriptor: file_web_foot_odd_detail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_foot_odd_detail_proto_goTypes,
		DependencyIndexes: file_web_foot_odd_detail_proto_depIdxs,
		MessageInfos:      file_web_foot_odd_detail_proto_msgTypes,
	}.Build()
	File_web_foot_odd_detail_proto = out.File
	file_web_foot_odd_detail_proto_rawDesc = nil
	file_web_foot_odd_detail_proto_goTypes = nil
	file_web_foot_odd_detail_proto_depIdxs = nil
}