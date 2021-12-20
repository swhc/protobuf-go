// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: web_foot_team_detail.proto

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

//球员榜
type WebFootTeamDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"` //请求语言
	TeamId   int64  `protobuf:"varint,5,opt,name=teamId,proto3" json:"teamId,omitempty"`    //球队ID
}

func (x *WebFootTeamDetailRequest) Reset() {
	*x = WebFootTeamDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_team_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootTeamDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootTeamDetailRequest) ProtoMessage() {}

func (x *WebFootTeamDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_team_detail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootTeamDetailRequest.ProtoReflect.Descriptor instead.
func (*WebFootTeamDetailRequest) Descriptor() ([]byte, []int) {
	return file_web_foot_team_detail_proto_rawDescGZIP(), []int{0}
}

func (x *WebFootTeamDetailRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *WebFootTeamDetailRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

type WebFootTeamDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List         []*WebFootInjuryOrSuspension `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`                 //受伤或停赛
	TransferList []*WebTeamTransfer           `protobuf:"bytes,2,rep,name=TransferList,proto3" json:"TransferList,omitempty"` //转会记录
}

func (x *WebFootTeamDetailResponse) Reset() {
	*x = WebFootTeamDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_team_detail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootTeamDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootTeamDetailResponse) ProtoMessage() {}

func (x *WebFootTeamDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_team_detail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootTeamDetailResponse.ProtoReflect.Descriptor instead.
func (*WebFootTeamDetailResponse) Descriptor() ([]byte, []int) {
	return file_web_foot_team_detail_proto_rawDescGZIP(), []int{1}
}

func (x *WebFootTeamDetailResponse) GetList() []*WebFootInjuryOrSuspension {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *WebFootTeamDetailResponse) GetTransferList() []*WebTeamTransfer {
	if x != nil {
		return x.TransferList
	}
	return nil
}

type WebTeamTransfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                    int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`                                      //球员id
	Name                  string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`                                   //球员名称
	Image                 string `protobuf:"bytes,3,opt,name=Image,proto3" json:"Image,omitempty"`                                 //球员头像
	Age                   string `protobuf:"bytes,4,opt,name=age,proto3" json:"age,omitempty"`                                     //球员年龄
	Reason                string `protobuf:"bytes,5,opt,name=Reason,proto3" json:"Reason,omitempty"`                               //原因
	Price                 string `protobuf:"bytes,6,opt,name=price,proto3" json:"price,omitempty"`                                 //转会金额
	OriginalTeam          string `protobuf:"bytes,7,opt,name=OriginalTeam,proto3" json:"OriginalTeam,omitempty"`                   //原球队
	OriginalTeamStartTime string `protobuf:"bytes,8,opt,name=OriginalTeamStartTime,proto3" json:"OriginalTeamStartTime,omitempty"` //原球队开始时间
}

func (x *WebTeamTransfer) Reset() {
	*x = WebTeamTransfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_team_detail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebTeamTransfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebTeamTransfer) ProtoMessage() {}

func (x *WebTeamTransfer) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_team_detail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebTeamTransfer.ProtoReflect.Descriptor instead.
func (*WebTeamTransfer) Descriptor() ([]byte, []int) {
	return file_web_foot_team_detail_proto_rawDescGZIP(), []int{2}
}

func (x *WebTeamTransfer) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WebTeamTransfer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WebTeamTransfer) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *WebTeamTransfer) GetAge() string {
	if x != nil {
		return x.Age
	}
	return ""
}

func (x *WebTeamTransfer) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *WebTeamTransfer) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *WebTeamTransfer) GetOriginalTeam() string {
	if x != nil {
		return x.OriginalTeam
	}
	return ""
}

func (x *WebTeamTransfer) GetOriginalTeamStartTime() string {
	if x != nil {
		return x.OriginalTeamStartTime
	}
	return ""
}

type WebFootInjuryOrSuspension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`              //球员id
	Name      string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`           //球员名称
	Image     string `protobuf:"bytes,3,opt,name=Image,proto3" json:"Image,omitempty"`         //球员头像
	Age       string `protobuf:"bytes,7,opt,name=age,proto3" json:"age,omitempty"`             //球员年龄
	Reason    string `protobuf:"bytes,4,opt,name=Reason,proto3" json:"Reason,omitempty"`       //原因
	Position  string `protobuf:"bytes,5,opt,name=position,proto3" json:"position,omitempty"`   //位置
	StartTime string `protobuf:"bytes,6,opt,name=startTime,proto3" json:"startTime,omitempty"` //受伤或停赛开始时间
}

func (x *WebFootInjuryOrSuspension) Reset() {
	*x = WebFootInjuryOrSuspension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_foot_team_detail_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebFootInjuryOrSuspension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebFootInjuryOrSuspension) ProtoMessage() {}

func (x *WebFootInjuryOrSuspension) ProtoReflect() protoreflect.Message {
	mi := &file_web_foot_team_detail_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebFootInjuryOrSuspension.ProtoReflect.Descriptor instead.
func (*WebFootInjuryOrSuspension) Descriptor() ([]byte, []int) {
	return file_web_foot_team_detail_proto_rawDescGZIP(), []int{3}
}

func (x *WebFootInjuryOrSuspension) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WebFootInjuryOrSuspension) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WebFootInjuryOrSuspension) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *WebFootInjuryOrSuspension) GetAge() string {
	if x != nil {
		return x.Age
	}
	return ""
}

func (x *WebFootInjuryOrSuspension) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *WebFootInjuryOrSuspension) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *WebFootInjuryOrSuspension) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

var File_web_foot_team_detail_proto protoreflect.FileDescriptor

var file_web_foot_team_detail_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x77, 0x65, 0x62, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x18,
	0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x81, 0x01, 0x0a,
	0x19, 0x57, 0x65, 0x62, 0x46, 0x6f, 0x6f, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x57, 0x65, 0x62, 0x46, 0x6f,
	0x6f, 0x74, 0x49, 0x6e, 0x6a, 0x75, 0x72, 0x79, 0x4f, 0x72, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0c, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x57, 0x65, 0x62, 0x54, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x52, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0xe5, 0x01, 0x0a, 0x0f, 0x57, 0x65, 0x62, 0x54, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x65,
	0x61, 0x6d, 0x12, 0x34, 0x0a, 0x15, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x65,
	0x61, 0x6d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x15, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xb9, 0x01, 0x0a, 0x19, 0x57, 0x65, 0x62,
	0x46, 0x6f, 0x6f, 0x74, 0x49, 0x6e, 0x6a, 0x75, 0x72, 0x79, 0x4f, 0x72, 0x53, 0x75, 0x73, 0x70,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web_foot_team_detail_proto_rawDescOnce sync.Once
	file_web_foot_team_detail_proto_rawDescData = file_web_foot_team_detail_proto_rawDesc
)

func file_web_foot_team_detail_proto_rawDescGZIP() []byte {
	file_web_foot_team_detail_proto_rawDescOnce.Do(func() {
		file_web_foot_team_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_foot_team_detail_proto_rawDescData)
	})
	return file_web_foot_team_detail_proto_rawDescData
}

var file_web_foot_team_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_web_foot_team_detail_proto_goTypes = []interface{}{
	(*WebFootTeamDetailRequest)(nil),  // 0: WebFootTeamDetailRequest
	(*WebFootTeamDetailResponse)(nil), // 1: WebFootTeamDetailResponse
	(*WebTeamTransfer)(nil),           // 2: WebTeamTransfer
	(*WebFootInjuryOrSuspension)(nil), // 3: WebFootInjuryOrSuspension
}
var file_web_foot_team_detail_proto_depIdxs = []int32{
	3, // 0: WebFootTeamDetailResponse.List:type_name -> WebFootInjuryOrSuspension
	2, // 1: WebFootTeamDetailResponse.TransferList:type_name -> WebTeamTransfer
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_web_foot_team_detail_proto_init() }
func file_web_foot_team_detail_proto_init() {
	if File_web_foot_team_detail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_foot_team_detail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootTeamDetailRequest); i {
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
		file_web_foot_team_detail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootTeamDetailResponse); i {
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
		file_web_foot_team_detail_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebTeamTransfer); i {
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
		file_web_foot_team_detail_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebFootInjuryOrSuspension); i {
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
			RawDescriptor: file_web_foot_team_detail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_foot_team_detail_proto_goTypes,
		DependencyIndexes: file_web_foot_team_detail_proto_depIdxs,
		MessageInfos:      file_web_foot_team_detail_proto_msgTypes,
	}.Build()
	File_web_foot_team_detail_proto = out.File
	file_web_foot_team_detail_proto_rawDesc = nil
	file_web_foot_team_detail_proto_goTypes = nil
	file_web_foot_team_detail_proto_depIdxs = nil
}
