// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: season_list.proto

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

//赛季列表
type SeasonListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComId    int64 `protobuf:"varint,1,opt,name=comId,proto3" json:"comId,omitempty"`       //联赛id
	SeasonId int64 `protobuf:"varint,2,opt,name=seasonId,proto3" json:"seasonId,omitempty"` //赛事id
}

func (x *SeasonListRequest) Reset() {
	*x = SeasonListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_season_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeasonListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeasonListRequest) ProtoMessage() {}

func (x *SeasonListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_season_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeasonListRequest.ProtoReflect.Descriptor instead.
func (*SeasonListRequest) Descriptor() ([]byte, []int) {
	return file_season_list_proto_rawDescGZIP(), []int{0}
}

func (x *SeasonListRequest) GetComId() int64 {
	if x != nil {
		return x.ComId
	}
	return 0
}

func (x *SeasonListRequest) GetSeasonId() int64 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

type SeasonListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*SeasonList `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SeasonListResponse) Reset() {
	*x = SeasonListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_season_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeasonListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeasonListResponse) ProtoMessage() {}

func (x *SeasonListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_season_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeasonListResponse.ProtoReflect.Descriptor instead.
func (*SeasonListResponse) Descriptor() ([]byte, []int) {
	return file_season_list_proto_rawDescGZIP(), []int{1}
}

func (x *SeasonListResponse) GetData() []*SeasonList {
	if x != nil {
		return x.Data
	}
	return nil
}

type SeasonList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year     string   `protobuf:"bytes,1,opt,name=year,proto3" json:"year,omitempty"`          // 赛季年份
	SupplyId int64    `protobuf:"varint,4,opt,name=supplyId,proto3" json:"supplyId,omitempty"` // 赛季年份
	Current  string   `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`    //当前分组赛
	Num      []string `protobuf:"bytes,3,rep,name=num,proto3" json:"num,omitempty"`            //分组赛列表
}

func (x *SeasonList) Reset() {
	*x = SeasonList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_season_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeasonList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeasonList) ProtoMessage() {}

func (x *SeasonList) ProtoReflect() protoreflect.Message {
	mi := &file_season_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeasonList.ProtoReflect.Descriptor instead.
func (*SeasonList) Descriptor() ([]byte, []int) {
	return file_season_list_proto_rawDescGZIP(), []int{2}
}

func (x *SeasonList) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *SeasonList) GetSupplyId() int64 {
	if x != nil {
		return x.SupplyId
	}
	return 0
}

func (x *SeasonList) GetCurrent() string {
	if x != nil {
		return x.Current
	}
	return ""
}

func (x *SeasonList) GetNum() []string {
	if x != nil {
		return x.Num
	}
	return nil
}

var File_season_list_proto protoreflect.FileDescriptor

var file_season_list_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x11, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6d, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x12, 0x53, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x68, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x42, 0x07, 0x5a, 0x05, 0x2e,
	0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_season_list_proto_rawDescOnce sync.Once
	file_season_list_proto_rawDescData = file_season_list_proto_rawDesc
)

func file_season_list_proto_rawDescGZIP() []byte {
	file_season_list_proto_rawDescOnce.Do(func() {
		file_season_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_season_list_proto_rawDescData)
	})
	return file_season_list_proto_rawDescData
}

var file_season_list_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_season_list_proto_goTypes = []interface{}{
	(*SeasonListRequest)(nil),  // 0: SeasonListRequest
	(*SeasonListResponse)(nil), // 1: SeasonListResponse
	(*SeasonList)(nil),         // 2: SeasonList
}
var file_season_list_proto_depIdxs = []int32{
	2, // 0: SeasonListResponse.data:type_name -> SeasonList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_season_list_proto_init() }
func file_season_list_proto_init() {
	if File_season_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_season_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeasonListRequest); i {
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
		file_season_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeasonListResponse); i {
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
		file_season_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeasonList); i {
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
			RawDescriptor: file_season_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_season_list_proto_goTypes,
		DependencyIndexes: file_season_list_proto_depIdxs,
		MessageInfos:      file_season_list_proto_msgTypes,
	}.Build()
	File_season_list_proto = out.File
	file_season_list_proto_rawDesc = nil
	file_season_list_proto_goTypes = nil
	file_season_list_proto_depIdxs = nil
}
