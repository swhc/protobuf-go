// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: foot_team_transfer.proto

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

type FootTeamTransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId   int64  `protobuf:"varint,1,opt,name=teamId,proto3" json:"teamId,omitempty"`
	Year     int64  `protobuf:"varint,2,opt,name=year,proto3" json:"year,omitempty"`
	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"` //请求语言  1:zh  2:en
}

func (x *FootTeamTransferRequest) Reset() {
	*x = FootTeamTransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_team_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootTeamTransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootTeamTransferRequest) ProtoMessage() {}

func (x *FootTeamTransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foot_team_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootTeamTransferRequest.ProtoReflect.Descriptor instead.
func (*FootTeamTransferRequest) Descriptor() ([]byte, []int) {
	return file_foot_team_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *FootTeamTransferRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *FootTeamTransferRequest) GetYear() int64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *FootTeamTransferRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type FootTeamTransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*FootTeamTransfer `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *FootTeamTransferResponse) Reset() {
	*x = FootTeamTransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_team_transfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootTeamTransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootTeamTransferResponse) ProtoMessage() {}

func (x *FootTeamTransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_foot_team_transfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootTeamTransferResponse.ProtoReflect.Descriptor instead.
func (*FootTeamTransferResponse) Descriptor() ([]byte, []int) {
	return file_foot_team_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *FootTeamTransferResponse) GetData() []*FootTeamTransfer {
	if x != nil {
		return x.Data
	}
	return nil
}

type FootTeamTransfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TitleType int64                   `protobuf:"varint,1,opt,name=TitleType,proto3" json:"TitleType,omitempty"`
	List      []*FootTeamTransferList `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *FootTeamTransfer) Reset() {
	*x = FootTeamTransfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_team_transfer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootTeamTransfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootTeamTransfer) ProtoMessage() {}

func (x *FootTeamTransfer) ProtoReflect() protoreflect.Message {
	mi := &file_foot_team_transfer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootTeamTransfer.ProtoReflect.Descriptor instead.
func (*FootTeamTransfer) Descriptor() ([]byte, []int) {
	return file_foot_team_transfer_proto_rawDescGZIP(), []int{2}
}

func (x *FootTeamTransfer) GetTitleType() int64 {
	if x != nil {
		return x.TitleType
	}
	return 0
}

func (x *FootTeamTransfer) GetList() []*FootTeamTransferList {
	if x != nil {
		return x.List
	}
	return nil
}

type FootTeamTransferList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonId    int64  `protobuf:"varint,1,opt,name=PersonId,proto3" json:"PersonId,omitempty"`      //球员id
	Type        int64  `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`              //转会类型
	TranTime    int64  `protobuf:"varint,3,opt,name=TranTime,proto3" json:"TranTime,omitempty"`      // 时间
	PersonName  string `protobuf:"bytes,4,opt,name=PersonName,proto3" json:"PersonName,omitempty"`   //球员中文名
	Age         int64  `protobuf:"varint,5,opt,name=Age,proto3" json:"Age,omitempty"`                //球员年龄
	Position    string `protobuf:"bytes,6,opt,name=Position,proto3" json:"Position,omitempty"`       //位置【前锋、后卫...】
	CountryName string `protobuf:"bytes,7,opt,name=CountryName,proto3" json:"CountryName,omitempty"` //球员国家
	Image       string `protobuf:"bytes,8,opt,name=Image,proto3" json:"Image,omitempty"`             //头像
}

func (x *FootTeamTransferList) Reset() {
	*x = FootTeamTransferList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_team_transfer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootTeamTransferList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootTeamTransferList) ProtoMessage() {}

func (x *FootTeamTransferList) ProtoReflect() protoreflect.Message {
	mi := &file_foot_team_transfer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootTeamTransferList.ProtoReflect.Descriptor instead.
func (*FootTeamTransferList) Descriptor() ([]byte, []int) {
	return file_foot_team_transfer_proto_rawDescGZIP(), []int{3}
}

func (x *FootTeamTransferList) GetPersonId() int64 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

func (x *FootTeamTransferList) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *FootTeamTransferList) GetTranTime() int64 {
	if x != nil {
		return x.TranTime
	}
	return 0
}

func (x *FootTeamTransferList) GetPersonName() string {
	if x != nil {
		return x.PersonName
	}
	return ""
}

func (x *FootTeamTransferList) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *FootTeamTransferList) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *FootTeamTransferList) GetCountryName() string {
	if x != nil {
		return x.CountryName
	}
	return ""
}

func (x *FootTeamTransferList) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

var File_foot_team_transfer_proto protoreflect.FileDescriptor

var file_foot_team_transfer_proto_rawDesc = []byte{
	0x0a, 0x18, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x17, 0x46, 0x6f,
	0x6f, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x41, 0x0a,
	0x18, 0x46, 0x6f, 0x6f, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x54, 0x65,
	0x61, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x5b, 0x0a, 0x10, 0x46, 0x6f, 0x6f, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xe8, 0x01,
	0x0a, 0x14, 0x46, 0x6f, 0x6f, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x41, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foot_team_transfer_proto_rawDescOnce sync.Once
	file_foot_team_transfer_proto_rawDescData = file_foot_team_transfer_proto_rawDesc
)

func file_foot_team_transfer_proto_rawDescGZIP() []byte {
	file_foot_team_transfer_proto_rawDescOnce.Do(func() {
		file_foot_team_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_foot_team_transfer_proto_rawDescData)
	})
	return file_foot_team_transfer_proto_rawDescData
}

var file_foot_team_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_foot_team_transfer_proto_goTypes = []interface{}{
	(*FootTeamTransferRequest)(nil),  // 0: FootTeamTransferRequest
	(*FootTeamTransferResponse)(nil), // 1: FootTeamTransferResponse
	(*FootTeamTransfer)(nil),         // 2: FootTeamTransfer
	(*FootTeamTransferList)(nil),     // 3: FootTeamTransferList
}
var file_foot_team_transfer_proto_depIdxs = []int32{
	2, // 0: FootTeamTransferResponse.Data:type_name -> FootTeamTransfer
	3, // 1: FootTeamTransfer.list:type_name -> FootTeamTransferList
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_foot_team_transfer_proto_init() }
func file_foot_team_transfer_proto_init() {
	if File_foot_team_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foot_team_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootTeamTransferRequest); i {
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
		file_foot_team_transfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootTeamTransferResponse); i {
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
		file_foot_team_transfer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootTeamTransfer); i {
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
		file_foot_team_transfer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootTeamTransferList); i {
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
			RawDescriptor: file_foot_team_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foot_team_transfer_proto_goTypes,
		DependencyIndexes: file_foot_team_transfer_proto_depIdxs,
		MessageInfos:      file_foot_team_transfer_proto_msgTypes,
	}.Build()
	File_foot_team_transfer_proto = out.File
	file_foot_team_transfer_proto_rawDesc = nil
	file_foot_team_transfer_proto_goTypes = nil
	file_foot_team_transfer_proto_depIdxs = nil
}
