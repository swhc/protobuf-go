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

type FootMatchFiltrateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilterType int64  `protobuf:"varint,6,opt,name=FilterType,proto3" json:"FilterType,omitempty"` //1=》即时数据，2=》热门
	Language   string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`      //请求语言
}

func (x *FootMatchFiltrateRequest) Reset() {
	*x = FootMatchFiltrateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_filtrate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootMatchFiltrateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootMatchFiltrateRequest) ProtoMessage() {}

func (x *FootMatchFiltrateRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FootMatchFiltrateRequest.ProtoReflect.Descriptor instead.
func (*FootMatchFiltrateRequest) Descriptor() ([]byte, []int) {
	return file_foot_match_filtrate_proto_rawDescGZIP(), []int{0}
}

func (x *FootMatchFiltrateRequest) GetFilterType() int64 {
	if x != nil {
		return x.FilterType
	}
	return 0
}

func (x *FootMatchFiltrateRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type FootMatchFiltrateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*FootMatchFiltrate `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *FootMatchFiltrateResponse) Reset() {
	*x = FootMatchFiltrateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_filtrate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootMatchFiltrateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootMatchFiltrateResponse) ProtoMessage() {}

func (x *FootMatchFiltrateResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FootMatchFiltrateResponse.ProtoReflect.Descriptor instead.
func (*FootMatchFiltrateResponse) Descriptor() ([]byte, []int) {
	return file_foot_match_filtrate_proto_rawDescGZIP(), []int{1}
}

func (x *FootMatchFiltrateResponse) GetData() []*FootMatchFiltrate {
	if x != nil {
		return x.Data
	}
	return nil
}

type FootMatchFiltrate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string                `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	List  []*FootTournamentInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *FootMatchFiltrate) Reset() {
	*x = FootMatchFiltrate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_filtrate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootMatchFiltrate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootMatchFiltrate) ProtoMessage() {}

func (x *FootMatchFiltrate) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FootMatchFiltrate.ProtoReflect.Descriptor instead.
func (*FootMatchFiltrate) Descriptor() ([]byte, []int) {
	return file_foot_match_filtrate_proto_rawDescGZIP(), []int{2}
}

func (x *FootMatchFiltrate) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FootMatchFiltrate) GetList() []*FootTournamentInfo {
	if x != nil {
		return x.List
	}
	return nil
}

type FootTournamentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TournamentId   int64  `protobuf:"varint,1,opt,name=tournamentId,proto3" json:"tournamentId,omitempty"`
	Count          int64  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	TournamentName string `protobuf:"bytes,2,opt,name=tournamentName,proto3" json:"tournamentName,omitempty"`
	ChInitials     string `protobuf:"bytes,3,opt,name=chInitials,proto3" json:"chInitials,omitempty"`
}

func (x *FootTournamentInfo) Reset() {
	*x = FootTournamentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_filtrate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootTournamentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootTournamentInfo) ProtoMessage() {}

func (x *FootTournamentInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FootTournamentInfo.ProtoReflect.Descriptor instead.
func (*FootTournamentInfo) Descriptor() ([]byte, []int) {
	return file_foot_match_filtrate_proto_rawDescGZIP(), []int{3}
}

func (x *FootTournamentInfo) GetTournamentId() int64 {
	if x != nil {
		return x.TournamentId
	}
	return 0
}

func (x *FootTournamentInfo) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *FootTournamentInfo) GetTournamentName() string {
	if x != nil {
		return x.TournamentName
	}
	return ""
}

func (x *FootTournamentInfo) GetChInitials() string {
	if x != nil {
		return x.ChInitials
	}
	return ""
}

var File_foot_match_filtrate_proto protoreflect.FileDescriptor

var file_foot_match_filtrate_proto_rawDesc = []byte{
	0x0a, 0x19, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x66, 0x69, 0x6c,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x18, 0x46,
	0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x22, 0x43, 0x0a, 0x19, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x46, 0x69, 0x6c, 0x74, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x52, 0x0a, 0x11, 0x46, 0x6f, 0x6f, 0x74,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x72, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x96, 0x01, 0x0a,
	0x12, 0x46, 0x6f, 0x6f, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a,
	0x0e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x68, 0x49, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*FootMatchFiltrateRequest)(nil),  // 0: FootMatchFiltrateRequest
	(*FootMatchFiltrateResponse)(nil), // 1: FootMatchFiltrateResponse
	(*FootMatchFiltrate)(nil),         // 2: FootMatchFiltrate
	(*FootTournamentInfo)(nil),        // 3: FootTournamentInfo
}
var file_foot_match_filtrate_proto_depIdxs = []int32{
	2, // 0: FootMatchFiltrateResponse.data:type_name -> FootMatchFiltrate
	3, // 1: FootMatchFiltrate.list:type_name -> FootTournamentInfo
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
			switch v := v.(*FootMatchFiltrateRequest); i {
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
			switch v := v.(*FootMatchFiltrateResponse); i {
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
			switch v := v.(*FootMatchFiltrate); i {
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
			switch v := v.(*FootTournamentInfo); i {
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
