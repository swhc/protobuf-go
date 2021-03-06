// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: foot_tournament_search.proto

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

type TournamentSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fuzzy    string `protobuf:"bytes,1,opt,name=fuzzy,proto3" json:"fuzzy,omitempty"`
	Language string `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *TournamentSearchRequest) Reset() {
	*x = TournamentSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_tournament_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TournamentSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TournamentSearchRequest) ProtoMessage() {}

func (x *TournamentSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foot_tournament_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TournamentSearchRequest.ProtoReflect.Descriptor instead.
func (*TournamentSearchRequest) Descriptor() ([]byte, []int) {
	return file_foot_tournament_search_proto_rawDescGZIP(), []int{0}
}

func (x *TournamentSearchRequest) GetFuzzy() string {
	if x != nil {
		return x.Fuzzy
	}
	return ""
}

func (x *TournamentSearchRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type TournamentSearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*SearchList `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *TournamentSearchResponse) Reset() {
	*x = TournamentSearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_tournament_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TournamentSearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TournamentSearchResponse) ProtoMessage() {}

func (x *TournamentSearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_foot_tournament_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TournamentSearchResponse.ProtoReflect.Descriptor instead.
func (*TournamentSearchResponse) Descriptor() ([]byte, []int) {
	return file_foot_tournament_search_proto_rawDescGZIP(), []int{1}
}

func (x *TournamentSearchResponse) GetList() []*SearchList {
	if x != nil {
		return x.List
	}
	return nil
}

type SearchList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"` //1?????????2?????????3??????
	Id   int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`    //ID
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"` //??????
}

func (x *SearchList) Reset() {
	*x = SearchList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_tournament_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchList) ProtoMessage() {}

func (x *SearchList) ProtoReflect() protoreflect.Message {
	mi := &file_foot_tournament_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchList.ProtoReflect.Descriptor instead.
func (*SearchList) Descriptor() ([]byte, []int) {
	return file_foot_tournament_search_proto_rawDescGZIP(), []int{2}
}

func (x *SearchList) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SearchList) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SearchList) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_foot_tournament_search_proto protoreflect.FileDescriptor

var file_foot_tournament_search_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b,
	0x0a, 0x17, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x75, 0x7a,
	0x7a, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x75, 0x7a, 0x7a, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x3b, 0x0a, 0x18, 0x54,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x44, 0x0a, 0x0a, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x07,
	0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foot_tournament_search_proto_rawDescOnce sync.Once
	file_foot_tournament_search_proto_rawDescData = file_foot_tournament_search_proto_rawDesc
)

func file_foot_tournament_search_proto_rawDescGZIP() []byte {
	file_foot_tournament_search_proto_rawDescOnce.Do(func() {
		file_foot_tournament_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_foot_tournament_search_proto_rawDescData)
	})
	return file_foot_tournament_search_proto_rawDescData
}

var file_foot_tournament_search_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_foot_tournament_search_proto_goTypes = []interface{}{
	(*TournamentSearchRequest)(nil),  // 0: TournamentSearchRequest
	(*TournamentSearchResponse)(nil), // 1: TournamentSearchResponse
	(*SearchList)(nil),               // 2: searchList
}
var file_foot_tournament_search_proto_depIdxs = []int32{
	2, // 0: TournamentSearchResponse.list:type_name -> searchList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_foot_tournament_search_proto_init() }
func file_foot_tournament_search_proto_init() {
	if File_foot_tournament_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foot_tournament_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TournamentSearchRequest); i {
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
		file_foot_tournament_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TournamentSearchResponse); i {
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
		file_foot_tournament_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchList); i {
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
			RawDescriptor: file_foot_tournament_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foot_tournament_search_proto_goTypes,
		DependencyIndexes: file_foot_tournament_search_proto_depIdxs,
		MessageInfos:      file_foot_tournament_search_proto_msgTypes,
	}.Build()
	File_foot_tournament_search_proto = out.File
	file_foot_tournament_search_proto_rawDesc = nil
	file_foot_tournament_search_proto_goTypes = nil
	file_foot_tournament_search_proto_depIdxs = nil
}
