// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: foot_match_list.proto

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
type FootMatchListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CmpId    int64  `protobuf:"varint,1,opt,name=cmp_id,json=cmpId,proto3" json:"cmp_id,omitempty"`         //  联赛id：46
	SuppId   int64  `protobuf:"varint,2,opt,name=supp_id,json=suppId,proto3" json:"supp_id,omitempty"`      //  年份对应id
	CptTurns string `protobuf:"bytes,3,opt,name=cpt_turns,json=cptTurns,proto3" json:"cpt_turns,omitempty"` //  轮次：第17轮
	Language string `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`                 //请求语言  1:zh  2:en
}

func (x *FootMatchListRequest) Reset() {
	*x = FootMatchListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootMatchListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootMatchListRequest) ProtoMessage() {}

func (x *FootMatchListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootMatchListRequest.ProtoReflect.Descriptor instead.
func (*FootMatchListRequest) Descriptor() ([]byte, []int) {
	return file_foot_match_list_proto_rawDescGZIP(), []int{0}
}

func (x *FootMatchListRequest) GetCmpId() int64 {
	if x != nil {
		return x.CmpId
	}
	return 0
}

func (x *FootMatchListRequest) GetSuppId() int64 {
	if x != nil {
		return x.SuppId
	}
	return 0
}

func (x *FootMatchListRequest) GetCptTurns() string {
	if x != nil {
		return x.CptTurns
	}
	return ""
}

func (x *FootMatchListRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type FootMatchListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64            `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	List  []*FootMatchList `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *FootMatchListResponse) Reset() {
	*x = FootMatchListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootMatchListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootMatchListResponse) ProtoMessage() {}

func (x *FootMatchListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootMatchListResponse.ProtoReflect.Descriptor instead.
func (*FootMatchListResponse) Descriptor() ([]byte, []int) {
	return file_foot_match_list_proto_rawDescGZIP(), []int{1}
}

func (x *FootMatchListResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *FootMatchListResponse) GetList() []*FootMatchList {
	if x != nil {
		return x.List
	}
	return nil
}

type FootMatchList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                              //   id
	CptDate       string `protobuf:"bytes,2,opt,name=cpt_date,json=cptDate,proto3" json:"cpt_date,omitempty"`                      //   日期
	MainTeamName  string `protobuf:"bytes,3,opt,name=main_team_name,json=mainTeamName,proto3" json:"main_team_name,omitempty"`     //   主队名称
	CustTeamName  string `protobuf:"bytes,4,opt,name=cust_team_name,json=custTeamName,proto3" json:"cust_team_name,omitempty"`     //   客队名称
	TeamMainScore int64  `protobuf:"varint,5,opt,name=team_main_score,json=teamMainScore,proto3" json:"team_main_score,omitempty"` //   主队得分分数
	TeamCustScore int64  `protobuf:"varint,6,opt,name=team_cust_score,json=teamCustScore,proto3" json:"team_cust_score,omitempty"` //   客队得分分数
	Score         string `protobuf:"bytes,7,opt,name=score,proto3" json:"score,omitempty"`                                         //   比分
	MainTeamId    int64  `protobuf:"varint,8,opt,name=main_team_id,json=mainTeamId,proto3" json:"main_team_id,omitempty"`          //   主队id
	CustTeamId    int64  `protobuf:"varint,9,opt,name=cust_team_id,json=custTeamId,proto3" json:"cust_team_id,omitempty"`          //   客队id
	MainImage     string `protobuf:"bytes,10,opt,name=main_image,json=mainImage,proto3" json:"main_image,omitempty"`               //   主队图片
	CustImage     string `protobuf:"bytes,11,opt,name=cust_image,json=custImage,proto3" json:"cust_image,omitempty"`               //   客队图片
}

func (x *FootMatchList) Reset() {
	*x = FootMatchList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootMatchList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootMatchList) ProtoMessage() {}

func (x *FootMatchList) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootMatchList.ProtoReflect.Descriptor instead.
func (*FootMatchList) Descriptor() ([]byte, []int) {
	return file_foot_match_list_proto_rawDescGZIP(), []int{2}
}

func (x *FootMatchList) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FootMatchList) GetCptDate() string {
	if x != nil {
		return x.CptDate
	}
	return ""
}

func (x *FootMatchList) GetMainTeamName() string {
	if x != nil {
		return x.MainTeamName
	}
	return ""
}

func (x *FootMatchList) GetCustTeamName() string {
	if x != nil {
		return x.CustTeamName
	}
	return ""
}

func (x *FootMatchList) GetTeamMainScore() int64 {
	if x != nil {
		return x.TeamMainScore
	}
	return 0
}

func (x *FootMatchList) GetTeamCustScore() int64 {
	if x != nil {
		return x.TeamCustScore
	}
	return 0
}

func (x *FootMatchList) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

func (x *FootMatchList) GetMainTeamId() int64 {
	if x != nil {
		return x.MainTeamId
	}
	return 0
}

func (x *FootMatchList) GetCustTeamId() int64 {
	if x != nil {
		return x.CustTeamId
	}
	return 0
}

func (x *FootMatchList) GetMainImage() string {
	if x != nil {
		return x.MainImage
	}
	return ""
}

func (x *FootMatchList) GetCustImage() string {
	if x != nil {
		return x.CustImage
	}
	return ""
}

var File_foot_match_list_proto protoreflect.FileDescriptor

var file_foot_match_list_proto_rawDesc = []byte{
	0x0a, 0x15, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x14, 0x46, 0x6f, 0x6f, 0x74, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x15, 0x0a, 0x06, 0x63, 0x6d, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x63, 0x6d, 0x70, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x75, 0x70, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x75, 0x70, 0x70, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x70, 0x74, 0x5f, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x70, 0x74, 0x54, 0x75, 0x72, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x51, 0x0a, 0x15, 0x46, 0x6f, 0x6f, 0x74,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xee, 0x02, 0x0a, 0x0d,
	0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x63, 0x70, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x70, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24,
	0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74,
	0x65, 0x61, 0x6d, 0x4d, 0x61, 0x69, 0x6e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x0f,
	0x74, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x65, 0x61, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c,
	0x63, 0x75, 0x73, 0x74, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x75, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foot_match_list_proto_rawDescOnce sync.Once
	file_foot_match_list_proto_rawDescData = file_foot_match_list_proto_rawDesc
)

func file_foot_match_list_proto_rawDescGZIP() []byte {
	file_foot_match_list_proto_rawDescOnce.Do(func() {
		file_foot_match_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_foot_match_list_proto_rawDescData)
	})
	return file_foot_match_list_proto_rawDescData
}

var file_foot_match_list_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_foot_match_list_proto_goTypes = []interface{}{
	(*FootMatchListRequest)(nil),  // 0: FootMatchListRequest
	(*FootMatchListResponse)(nil), // 1: FootMatchListResponse
	(*FootMatchList)(nil),         // 2: FootMatchList
}
var file_foot_match_list_proto_depIdxs = []int32{
	2, // 0: FootMatchListResponse.list:type_name -> FootMatchList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_foot_match_list_proto_init() }
func file_foot_match_list_proto_init() {
	if File_foot_match_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foot_match_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootMatchListRequest); i {
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
		file_foot_match_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootMatchListResponse); i {
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
		file_foot_match_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootMatchList); i {
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
			RawDescriptor: file_foot_match_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foot_match_list_proto_goTypes,
		DependencyIndexes: file_foot_match_list_proto_depIdxs,
		MessageInfos:      file_foot_match_list_proto_msgTypes,
	}.Build()
	File_foot_match_list_proto = out.File
	file_foot_match_list_proto_rawDesc = nil
	file_foot_match_list_proto_goTypes = nil
	file_foot_match_list_proto_depIdxs = nil
}
