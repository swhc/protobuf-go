// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: web_foot_odd.proto

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
type WebOddListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId  int64  `protobuf:"varint,2,opt,name=eventId,proto3" json:"eventId,omitempty"`  //比赛id
	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"` //请求语言  1:zh  2:en
}

func (x *WebOddListRequest) Reset() {
	*x = WebOddListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_odd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebOddListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebOddListRequest) ProtoMessage() {}

func (x *WebOddListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_odd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebOddListRequest.ProtoReflect.Descriptor instead.
func (*WebOddListRequest) Descriptor() ([]byte, []int) {
	return file_web_foot_odd_proto_rawDescGZIP(), []int{0}
}

func (x *WebOddListRequest) GetEventId() int64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *WebOddListRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type WebOddListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Maximum    *OddsDetails        `protobuf:"bytes,1,opt,name=Maximum,proto3" json:"Maximum,omitempty"`
	Minimum    *OddsDetails        `protobuf:"bytes,2,opt,name=Minimum,proto3" json:"Minimum,omitempty"`
	Average    *OddsDetails        `protobuf:"bytes,4,opt,name=Average,proto3" json:"Average,omitempty"`
	Dispersion []float32           `protobuf:"fixed32,11,rep,packed,name=dispersion,proto3" json:"dispersion,omitempty"` //凯利离散度
	List       []*WebOddDetailList `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *WebOddListResponse) Reset() {
	*x = WebOddListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_odd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebOddListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebOddListResponse) ProtoMessage() {}

func (x *WebOddListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_odd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebOddListResponse.ProtoReflect.Descriptor instead.
func (*WebOddListResponse) Descriptor() ([]byte, []int) {
	return file_web_foot_odd_proto_rawDescGZIP(), []int{1}
}

func (x *WebOddListResponse) GetMaximum() *OddsDetails {
	if x != nil {
		return x.Maximum
	}
	return nil
}

func (x *WebOddListResponse) GetMinimum() *OddsDetails {
	if x != nil {
		return x.Minimum
	}
	return nil
}

func (x *WebOddListResponse) GetAverage() *OddsDetails {
	if x != nil {
		return x.Average
	}
	return nil
}

func (x *WebOddListResponse) GetDispersion() []float32 {
	if x != nil {
		return x.Dispersion
	}
	return nil
}

func (x *WebOddListResponse) GetList() []*WebOddDetailList {
	if x != nil {
		return x.List
	}
	return nil
}

type WebOddDetailList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Init        *OddsDetails `protobuf:"bytes,1,opt,name=Init,proto3" json:"Init,omitempty"`
	Curr        *OddsDetails `protobuf:"bytes,2,opt,name=Curr,proto3" json:"Curr,omitempty"`
	CompanyName string       `protobuf:"bytes,6,opt,name=companyName,proto3" json:"companyName,omitempty"` //公司名称
}

func (x *WebOddDetailList) Reset() {
	*x = WebOddDetailList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_odd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebOddDetailList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebOddDetailList) ProtoMessage() {}

func (x *WebOddDetailList) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_odd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebOddDetailList.ProtoReflect.Descriptor instead.
func (*WebOddDetailList) Descriptor() ([]byte, []int) {
	return file_web_foot_odd_proto_rawDescGZIP(), []int{2}
}

func (x *WebOddDetailList) GetInit() *OddsDetails {
	if x != nil {
		return x.Init
	}
	return nil
}

func (x *WebOddDetailList) GetCurr() *OddsDetails {
	if x != nil {
		return x.Curr
	}
	return nil
}

func (x *WebOddDetailList) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

type OddsDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Win        float32   `protobuf:"fixed32,1,opt,name=win,proto3" json:"win,omitempty"`                      //胜
	Draw       float32   `protobuf:"fixed32,2,opt,name=draw,proto3" json:"draw,omitempty"`                    //平
	Lose       float32   `protobuf:"fixed32,3,opt,name=lose,proto3" json:"lose,omitempty"`                    //负
	WinProb    float32   `protobuf:"fixed32,4,opt,name=winProb,proto3" json:"winProb,omitempty"`              //胜率
	DrawProb   float32   `protobuf:"fixed32,5,opt,name=drawProb,proto3" json:"drawProb,omitempty"`            //平率
	LoseProb   float32   `protobuf:"fixed32,6,opt,name=loseProb,proto3" json:"loseProb,omitempty"`            //负率
	ReturnProb float32   `protobuf:"fixed32,7,opt,name=returnProb,proto3" json:"returnProb,omitempty"`        //返回率
	KellyIndex []float32 `protobuf:"fixed32,8,rep,packed,name=kellyIndex,proto3" json:"kellyIndex,omitempty"` //凯利指数
	Variety    int64     `protobuf:"varint,9,opt,name=variety,proto3" json:"variety,omitempty"`               //时间
	View       string    `protobuf:"bytes,10,opt,name=view,proto3" json:"view,omitempty"`                     //视图
}

func (x *OddsDetails) Reset() {
	*x = OddsDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_odd_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OddsDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OddsDetails) ProtoMessage() {}

func (x *OddsDetails) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_odd_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OddsDetails.ProtoReflect.Descriptor instead.
func (*OddsDetails) Descriptor() ([]byte, []int) {
	return file_web_foot_odd_proto_rawDescGZIP(), []int{3}
}

func (x *OddsDetails) GetWin() float32 {
	if x != nil {
		return x.Win
	}
	return 0
}

func (x *OddsDetails) GetDraw() float32 {
	if x != nil {
		return x.Draw
	}
	return 0
}

func (x *OddsDetails) GetLose() float32 {
	if x != nil {
		return x.Lose
	}
	return 0
}

func (x *OddsDetails) GetWinProb() float32 {
	if x != nil {
		return x.WinProb
	}
	return 0
}

func (x *OddsDetails) GetDrawProb() float32 {
	if x != nil {
		return x.DrawProb
	}
	return 0
}

func (x *OddsDetails) GetLoseProb() float32 {
	if x != nil {
		return x.LoseProb
	}
	return 0
}

func (x *OddsDetails) GetReturnProb() float32 {
	if x != nil {
		return x.ReturnProb
	}
	return 0
}

func (x *OddsDetails) GetKellyIndex() []float32 {
	if x != nil {
		return x.KellyIndex
	}
	return nil
}

func (x *OddsDetails) GetVariety() int64 {
	if x != nil {
		return x.Variety
	}
	return 0
}

func (x *OddsDetails) GetView() string {
	if x != nil {
		return x.View
	}
	return ""
}

var File_web_foot_odd_proto protoreflect.FileDescriptor

var file_web_foot_odd_proto_rawDesc = []byte{
	0x0a, 0x12, 0x77, 0x65, 0x62, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6f, 0x64, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x11, 0x57, 0x65, 0x62, 0x4f, 0x64, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22,
	0xd3, 0x01, 0x0a, 0x12, 0x57, 0x65, 0x62, 0x4f, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x07, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x64, 0x64, 0x73, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x26,
	0x0a, 0x07, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x4f, 0x64, 0x64, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x4d,
	0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x26, 0x0a, 0x07, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x64, 0x64, 0x73, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x64, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x02, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x57,
	0x65, 0x62, 0x4f, 0x64, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x78, 0x0a, 0x10, 0x57, 0x65, 0x62, 0x4f, 0x64, 0x64, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x49, 0x6e, 0x69,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x64, 0x64, 0x73, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x43,
	0x75, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x64, 0x64, 0x73,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x04, 0x43, 0x75, 0x72, 0x72, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x87, 0x02, 0x0a, 0x0b, 0x4f, 0x64, 0x64, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x77, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x77, 0x69,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x72, 0x61, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x04, 0x64, 0x72, 0x61, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x04, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x69, 0x6e,
	0x50, 0x72, 0x6f, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x77, 0x69, 0x6e, 0x50,
	0x72, 0x6f, 0x62, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x72, 0x61, 0x77, 0x50, 0x72, 0x6f, 0x62, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x64, 0x72, 0x61, 0x77, 0x50, 0x72, 0x6f, 0x62, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x12, 0x1e, 0x0a, 0x0a, 0x72,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x50, 0x72, 0x6f, 0x62, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x50, 0x72, 0x6f, 0x62, 0x12, 0x1e, 0x0a, 0x0a, 0x6b,
	0x65, 0x6c, 0x6c, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x08, 0x20, 0x03, 0x28, 0x02, 0x52,
	0x0a, 0x6b, 0x65, 0x6c, 0x6c, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x61,
	0x72, 0x69, 0x65, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web_foot_odd_proto_rawDescOnce sync.Once
	file_web_foot_odd_proto_rawDescData = file_web_foot_odd_proto_rawDesc
)

func file_web_foot_odd_proto_rawDescGZIP() []byte {
	file_web_foot_odd_proto_rawDescOnce.Do(func() {
		file_web_foot_odd_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_foot_odd_proto_rawDescData)
	})
	return file_web_foot_odd_proto_rawDescData
}

var file_web_foot_odd_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_web_foot_odd_proto_goTypes = []interface{}{
	(*WebOddListRequest)(nil),  // 0: WebOddListRequest
	(*WebOddListResponse)(nil), // 1: WebOddListResponse
	(*WebOddDetailList)(nil),   // 2: WebOddDetailList
	(*OddsDetails)(nil),        // 3: OddsDetails
}
var file_web_foot_odd_proto_depIdxs = []int32{
	3, // 0: WebOddListResponse.Maximum:type_name -> OddsDetails
	3, // 1: WebOddListResponse.Minimum:type_name -> OddsDetails
	3, // 2: WebOddListResponse.Average:type_name -> OddsDetails
	2, // 3: WebOddListResponse.list:type_name -> WebOddDetailList
	3, // 4: WebOddDetailList.Init:type_name -> OddsDetails
	3, // 5: WebOddDetailList.Curr:type_name -> OddsDetails
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_web_foot_odd_proto_init() }
func file_web_foot_odd_proto_init() {
	if File_web_foot_odd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_foot_odd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebOddListRequest); i {
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
		file_web_foot_odd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebOddListResponse); i {
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
		file_web_foot_odd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebOddDetailList); i {
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
		file_web_foot_odd_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OddsDetails); i {
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
			RawDescriptor: file_web_foot_odd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_foot_odd_proto_goTypes,
		DependencyIndexes: file_web_foot_odd_proto_depIdxs,
		MessageInfos:      file_web_foot_odd_proto_msgTypes,
	}.Build()
	File_web_foot_odd_proto = out.File
	file_web_foot_odd_proto_rawDesc = nil
	file_web_foot_odd_proto_goTypes = nil
	file_web_foot_odd_proto_depIdxs = nil
}