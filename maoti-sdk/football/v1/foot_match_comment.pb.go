// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: foot_match_comment.proto

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

//比赛文字直播
type FootMatchCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId   string `protobuf:"bytes,1,opt,name=eventId,proto3" json:"eventId,omitempty"`
	CommentId string `protobuf:"bytes,2,opt,name=commentId,proto3" json:"commentId,omitempty"`
	Language  string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"` //请求语言  1:zh  2:en
}

func (x *FootMatchCommentRequest) Reset() {
	*x = FootMatchCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootMatchCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootMatchCommentRequest) ProtoMessage() {}

func (x *FootMatchCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootMatchCommentRequest.ProtoReflect.Descriptor instead.
func (*FootMatchCommentRequest) Descriptor() ([]byte, []int) {
	return file_foot_match_comment_proto_rawDescGZIP(), []int{0}
}

func (x *FootMatchCommentRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *FootMatchCommentRequest) GetCommentId() string {
	if x != nil {
		return x.CommentId
	}
	return ""
}

func (x *FootMatchCommentRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type FootMatchCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*FootCommentList `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *FootMatchCommentResponse) Reset() {
	*x = FootMatchCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootMatchCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootMatchCommentResponse) ProtoMessage() {}

func (x *FootMatchCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootMatchCommentResponse.ProtoReflect.Descriptor instead.
func (*FootMatchCommentResponse) Descriptor() ([]byte, []int) {
	return file_foot_match_comment_proto_rawDescGZIP(), []int{1}
}

func (x *FootMatchCommentResponse) GetList() []*FootCommentList {
	if x != nil {
		return x.List
	}
	return nil
}

type FootCommentList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CommentId    int64  `protobuf:"varint,2,opt,name=commentId,proto3" json:"commentId,omitempty"`
	TypeId       int64  `protobuf:"varint,3,opt,name=typeId,proto3" json:"typeId,omitempty"`
	Image        string `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Elapsed      string `protobuf:"bytes,5,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	ElapsedPlus  string `protobuf:"bytes,6,opt,name=elapsedPlus,proto3" json:"elapsedPlus,omitempty"`
	Comment      string `protobuf:"bytes,7,opt,name=comment,proto3" json:"comment,omitempty"`
	IncidentText string `protobuf:"bytes,8,opt,name=incidentText,proto3" json:"incidentText,omitempty"`
	AddTime      int64  `protobuf:"varint,9,opt,name=addTime,proto3" json:"addTime,omitempty"`
}

func (x *FootCommentList) Reset() {
	*x = FootCommentList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootCommentList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootCommentList) ProtoMessage() {}

func (x *FootCommentList) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootCommentList.ProtoReflect.Descriptor instead.
func (*FootCommentList) Descriptor() ([]byte, []int) {
	return file_foot_match_comment_proto_rawDescGZIP(), []int{2}
}

func (x *FootCommentList) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FootCommentList) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *FootCommentList) GetTypeId() int64 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *FootCommentList) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *FootCommentList) GetElapsed() string {
	if x != nil {
		return x.Elapsed
	}
	return ""
}

func (x *FootCommentList) GetElapsedPlus() string {
	if x != nil {
		return x.ElapsedPlus
	}
	return ""
}

func (x *FootCommentList) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *FootCommentList) GetIncidentText() string {
	if x != nil {
		return x.IncidentText
	}
	return ""
}

func (x *FootCommentList) GetAddTime() int64 {
	if x != nil {
		return x.AddTime
	}
	return 0
}

var File_foot_match_comment_proto protoreflect.FileDescriptor

var file_foot_match_comment_proto_rawDesc = []byte{
	0x0a, 0x18, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x17, 0x46, 0x6f,
	0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x40, 0x0a, 0x18, 0x46, 0x6f, 0x6f,
	0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x85, 0x02, 0x0a, 0x0f,
	0x46, 0x6f, 0x6f, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6c, 0x61,
	0x70, 0x73, 0x65, 0x64, 0x50, 0x6c, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x50, 0x6c, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x54, 0x65, 0x78, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x63,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x64, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foot_match_comment_proto_rawDescOnce sync.Once
	file_foot_match_comment_proto_rawDescData = file_foot_match_comment_proto_rawDesc
)

func file_foot_match_comment_proto_rawDescGZIP() []byte {
	file_foot_match_comment_proto_rawDescOnce.Do(func() {
		file_foot_match_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_foot_match_comment_proto_rawDescData)
	})
	return file_foot_match_comment_proto_rawDescData
}

var file_foot_match_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_foot_match_comment_proto_goTypes = []interface{}{
	(*FootMatchCommentRequest)(nil),  // 0: FootMatchCommentRequest
	(*FootMatchCommentResponse)(nil), // 1: FootMatchCommentResponse
	(*FootCommentList)(nil),          // 2: FootCommentList
}
var file_foot_match_comment_proto_depIdxs = []int32{
	2, // 0: FootMatchCommentResponse.list:type_name -> FootCommentList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_foot_match_comment_proto_init() }
func file_foot_match_comment_proto_init() {
	if File_foot_match_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foot_match_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootMatchCommentRequest); i {
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
		file_foot_match_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootMatchCommentResponse); i {
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
		file_foot_match_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootCommentList); i {
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
			RawDescriptor: file_foot_match_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foot_match_comment_proto_goTypes,
		DependencyIndexes: file_foot_match_comment_proto_depIdxs,
		MessageInfos:      file_foot_match_comment_proto_msgTypes,
	}.Build()
	File_foot_match_comment_proto = out.File
	file_foot_match_comment_proto_rawDesc = nil
	file_foot_match_comment_proto_goTypes = nil
	file_foot_match_comment_proto_depIdxs = nil
}
