// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: foot_match_filtrate.proto

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

type V2FootMatchFiltrateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilterType int64  `protobuf:"varint,1,opt,name=FilterType,proto3" json:"FilterType,omitempty"` //1=》即时数据，2=》热门，3=》国家
	ListType   int64  `protobuf:"varint,2,opt,name=listType,proto3" json:"listType,omitempty"`     //1.全部，2.即时，3.赛果，4.赛程
	SelectDate string `protobuf:"bytes,3,opt,name=selectDate,proto3" json:"selectDate,omitempty"`  //选择的日期，为空时默认今天
	Language   string `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`      //请求语言
}

func (x *V2FootMatchFiltrateRequest) Reset() {
	*x = V2FootMatchFiltrateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_filtrate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchFiltrateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchFiltrateRequest) ProtoMessage() {}

func (x *V2FootMatchFiltrateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_filtrate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchFiltrateRequest.ProtoReflect.Descriptor instead.
func (*V2FootMatchFiltrateRequest) Descriptor() ([]byte, []int) {
	return file_foot_match_filtrate_proto_rawDescGZIP(), []int{0}
}

func (x *V2FootMatchFiltrateRequest) GetFilterType() int64 {
	if x != nil {
		return x.FilterType
	}
	return 0
}

func (x *V2FootMatchFiltrateRequest) GetListType() int64 {
	if x != nil {
		return x.ListType
	}
	return 0
}

func (x *V2FootMatchFiltrateRequest) GetSelectDate() string {
	if x != nil {
		return x.SelectDate
	}
	return ""
}

func (x *V2FootMatchFiltrateRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type V2FootMatchFiltrateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*V2FootMatchFiltrate `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *V2FootMatchFiltrateResponse) Reset() {
	*x = V2FootMatchFiltrateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_filtrate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchFiltrateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchFiltrateResponse) ProtoMessage() {}

func (x *V2FootMatchFiltrateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_filtrate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchFiltrateResponse.ProtoReflect.Descriptor instead.
func (*V2FootMatchFiltrateResponse) Descriptor() ([]byte, []int) {
	return file_foot_match_filtrate_proto_rawDescGZIP(), []int{1}
}

func (x *V2FootMatchFiltrateResponse) GetData() []*V2FootMatchFiltrate {
	if x != nil {
		return x.Data
	}
	return nil
}

type V2FootMatchFiltrate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`       //大洲ID
	Title string                  `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"` //左侧标题 热门赛事/大洲名称
	List  []*V2FootTournamentInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *V2FootMatchFiltrate) Reset() {
	*x = V2FootMatchFiltrate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_filtrate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootMatchFiltrate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootMatchFiltrate) ProtoMessage() {}

func (x *V2FootMatchFiltrate) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_filtrate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootMatchFiltrate.ProtoReflect.Descriptor instead.
func (*V2FootMatchFiltrate) Descriptor() ([]byte, []int) {
	return file_foot_match_filtrate_proto_rawDescGZIP(), []int{2}
}

func (x *V2FootMatchFiltrate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *V2FootMatchFiltrate) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *V2FootMatchFiltrate) GetList() []*V2FootTournamentInfo {
	if x != nil {
		return x.List
	}
	return nil
}

type V2FootTournamentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TournamentId   int64  `protobuf:"varint,1,opt,name=tournamentId,proto3" json:"tournamentId,omitempty"`    //赛事ID/国家ID
	Count          int64  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`                  //赛事/国家下的比赛场次
	TournamentName string `protobuf:"bytes,2,opt,name=tournamentName,proto3" json:"tournamentName,omitempty"` //赛事名称/国家名称
	ChInitials     string `protobuf:"bytes,3,opt,name=chInitials,proto3" json:"chInitials,omitempty"`         //赛事/国家首拼
}

func (x *V2FootTournamentInfo) Reset() {
	*x = V2FootTournamentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_filtrate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2FootTournamentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2FootTournamentInfo) ProtoMessage() {}

func (x *V2FootTournamentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_filtrate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2FootTournamentInfo.ProtoReflect.Descriptor instead.
func (*V2FootTournamentInfo) Descriptor() ([]byte, []int) {
	return file_foot_match_filtrate_proto_rawDescGZIP(), []int{3}
}

func (x *V2FootTournamentInfo) GetTournamentId() int64 {
	if x != nil {
		return x.TournamentId
	}
	return 0
}

func (x *V2FootTournamentInfo) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *V2FootTournamentInfo) GetTournamentName() string {
	if x != nil {
		return x.TournamentName
	}
	return ""
}

func (x *V2FootTournamentInfo) GetChInitials() string {
	if x != nil {
		return x.ChInitials
	}
	return ""
}

var File_foot_match_filtrate_proto protoreflect.FileDescriptor

var file_foot_match_filtrate_proto_rawDesc = []byte{
	0x0a, 0x19, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x66, 0x69, 0x6c,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x1a,
	0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x69,
	0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x69,
	0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x22, 0x47, 0x0a, 0x1b, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x46, 0x69, 0x6c, 0x74, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x46, 0x69, 0x6c,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x66, 0x0a, 0x13, 0x56,
	0x32, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x54,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x22, 0x98, 0x01, 0x0a, 0x14, 0x56, 0x32, 0x46, 0x6f, 0x6f, 0x74, 0x54, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x0c,
	0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x68, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x68, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x42, 0x07,
	0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foot_match_filtrate_proto_rawDescOnce sync.Once
	file_foot_match_filtrate_proto_rawDescData = file_foot_match_filtrate_proto_rawDesc
)

func file_foot_match_filtrate_proto_rawDescGZIP() []byte {
	file_foot_match_filtrate_proto_rawDescOnce.Do(func() {
		file_foot_match_filtrate_proto_rawDescData = protoimpl.X.CompressGZIP(file_foot_match_filtrate_proto_rawDescData)
	})
	return file_foot_match_filtrate_proto_rawDescData
}

var file_foot_match_filtrate_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_foot_match_filtrate_proto_goTypes = []interface{}{
	(*V2FootMatchFiltrateRequest)(nil),  // 0: V2FootMatchFiltrateRequest
	(*V2FootMatchFiltrateResponse)(nil), // 1: V2FootMatchFiltrateResponse
	(*V2FootMatchFiltrate)(nil),         // 2: V2FootMatchFiltrate
	(*V2FootTournamentInfo)(nil),        // 3: V2FootTournamentInfo
}
var file_foot_match_filtrate_proto_depIdxs = []int32{
	2, // 0: V2FootMatchFiltrateResponse.data:type_name -> V2FootMatchFiltrate
	3, // 1: V2FootMatchFiltrate.list:type_name -> V2FootTournamentInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_foot_match_filtrate_proto_init() }
func file_foot_match_filtrate_proto_init() {
	if File_foot_match_filtrate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foot_match_filtrate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchFiltrateRequest); i {
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
		file_foot_match_filtrate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchFiltrateResponse); i {
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
		file_foot_match_filtrate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootMatchFiltrate); i {
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
		file_foot_match_filtrate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2FootTournamentInfo); i {
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
			RawDescriptor: file_foot_match_filtrate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foot_match_filtrate_proto_goTypes,
		DependencyIndexes: file_foot_match_filtrate_proto_depIdxs,
		MessageInfos:      file_foot_match_filtrate_proto_msgTypes,
	}.Build()
	File_foot_match_filtrate_proto = out.File
	file_foot_match_filtrate_proto_rawDesc = nil
	file_foot_match_filtrate_proto_goTypes = nil
	file_foot_match_filtrate_proto_depIdxs = nil
}
