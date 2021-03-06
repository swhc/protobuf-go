// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: trans.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 单条翻译
type TransOneRequestParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Lang    string `protobuf:"bytes,3,opt,name=lang,proto3" json:"lang,omitempty"`
}

func (x *TransOneRequestParam) Reset() {
	*x = TransOneRequestParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trans_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransOneRequestParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransOneRequestParam) ProtoMessage() {}

func (x *TransOneRequestParam) ProtoReflect() protoreflect.Message {
	mi := &file_trans_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransOneRequestParam.ProtoReflect.Descriptor instead.
func (*TransOneRequestParam) Descriptor() ([]byte, []int) {
	return file_trans_proto_rawDescGZIP(), []int{0}
}

func (x *TransOneRequestParam) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TransOneRequestParam) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *TransOneRequestParam) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

type TransOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *TransOneResponse) Reset() {
	*x = TransOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trans_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransOneResponse) ProtoMessage() {}

func (x *TransOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trans_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransOneResponse.ProtoReflect.Descriptor instead.
func (*TransOneResponse) Descriptor() ([]byte, []int) {
	return file_trans_proto_rawDescGZIP(), []int{1}
}

func (x *TransOneResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// 批量翻译
type TransRequestParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Content string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Lang    []string `protobuf:"bytes,3,rep,name=lang,proto3" json:"lang,omitempty"`
}

func (x *TransRequestParam) Reset() {
	*x = TransRequestParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trans_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransRequestParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransRequestParam) ProtoMessage() {}

func (x *TransRequestParam) ProtoReflect() protoreflect.Message {
	mi := &file_trans_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransRequestParam.ProtoReflect.Descriptor instead.
func (*TransRequestParam) Descriptor() ([]byte, []int) {
	return file_trans_proto_rawDescGZIP(), []int{2}
}

func (x *TransRequestParam) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TransRequestParam) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *TransRequestParam) GetLang() []string {
	if x != nil {
		return x.Lang
	}
	return nil
}

type TransResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret     int64             `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Msg     string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Success map[string]string `protobuf:"bytes,3,rep,name=success,proto3" json:"success,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Err     map[string]string `protobuf:"bytes,4,rep,name=err,proto3" json:"err,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TransResponse) Reset() {
	*x = TransResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trans_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransResponse) ProtoMessage() {}

func (x *TransResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trans_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransResponse.ProtoReflect.Descriptor instead.
func (*TransResponse) Descriptor() ([]byte, []int) {
	return file_trans_proto_rawDescGZIP(), []int{3}
}

func (x *TransResponse) GetRet() int64 {
	if x != nil {
		return x.Ret
	}
	return 0
}

func (x *TransResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *TransResponse) GetSuccess() map[string]string {
	if x != nil {
		return x.Success
	}
	return nil
}

func (x *TransResponse) GetErr() map[string]string {
	if x != nil {
		return x.Err
	}
	return nil
}

var File_trans_proto protoreflect.FileDescriptor

var file_trans_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a,
	0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x22, 0x2c, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x55, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x22, 0x89, 0x02, 0x0a,
	0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x72, 0x65, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x35, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x03, 0x65, 0x72, 0x72,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x03, 0x65, 0x72, 0x72, 0x1a, 0x3a, 0x0a, 0x0c, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x36, 0x0a, 0x08, 0x45, 0x72, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x7c, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x0e, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0c,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x65, 0x12, 0x15, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x1a, 0x11, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4f, 0x6e, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trans_proto_rawDescOnce sync.Once
	file_trans_proto_rawDescData = file_trans_proto_rawDesc
)

func file_trans_proto_rawDescGZIP() []byte {
	file_trans_proto_rawDescOnce.Do(func() {
		file_trans_proto_rawDescData = protoimpl.X.CompressGZIP(file_trans_proto_rawDescData)
	})
	return file_trans_proto_rawDescData
}

var file_trans_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_trans_proto_goTypes = []interface{}{
	(*TransOneRequestParam)(nil), // 0: TransOneRequestParam
	(*TransOneResponse)(nil),     // 1: TransOneResponse
	(*TransRequestParam)(nil),    // 2: TransRequestParam
	(*TransResponse)(nil),        // 3: TransResponse
	nil,                          // 4: TransResponse.SuccessEntry
	nil,                          // 5: TransResponse.ErrEntry
}
var file_trans_proto_depIdxs = []int32{
	4, // 0: TransResponse.success:type_name -> TransResponse.SuccessEntry
	5, // 1: TransResponse.err:type_name -> TransResponse.ErrEntry
	2, // 2: Service.TranslateMultiple:input_type -> TransRequestParam
	0, // 3: Service.TranslateOne:input_type -> TransOneRequestParam
	3, // 4: Service.TranslateMultiple:output_type -> TransResponse
	1, // 5: Service.TranslateOne:output_type -> TransOneResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_trans_proto_init() }
func file_trans_proto_init() {
	if File_trans_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trans_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransOneRequestParam); i {
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
		file_trans_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransOneResponse); i {
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
		file_trans_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransRequestParam); i {
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
		file_trans_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransResponse); i {
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
			RawDescriptor: file_trans_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trans_proto_goTypes,
		DependencyIndexes: file_trans_proto_depIdxs,
		MessageInfos:      file_trans_proto_msgTypes,
	}.Build()
	File_trans_proto = out.File
	file_trans_proto_rawDesc = nil
	file_trans_proto_goTypes = nil
	file_trans_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	TranslateMultiple(ctx context.Context, in *TransRequestParam, opts ...grpc.CallOption) (*TransResponse, error)
	TranslateOne(ctx context.Context, in *TransOneRequestParam, opts ...grpc.CallOption) (*TransOneResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) TranslateMultiple(ctx context.Context, in *TransRequestParam, opts ...grpc.CallOption) (*TransResponse, error) {
	out := new(TransResponse)
	err := c.cc.Invoke(ctx, "/Service/TranslateMultiple", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) TranslateOne(ctx context.Context, in *TransOneRequestParam, opts ...grpc.CallOption) (*TransOneResponse, error) {
	out := new(TransOneResponse)
	err := c.cc.Invoke(ctx, "/Service/TranslateOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	TranslateMultiple(context.Context, *TransRequestParam) (*TransResponse, error)
	TranslateOne(context.Context, *TransOneRequestParam) (*TransOneResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) TranslateMultiple(context.Context, *TransRequestParam) (*TransResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TranslateMultiple not implemented")
}
func (*UnimplementedServiceServer) TranslateOne(context.Context, *TransOneRequestParam) (*TransOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TranslateOne not implemented")
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_TranslateMultiple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransRequestParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).TranslateMultiple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Service/TranslateMultiple",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).TranslateMultiple(ctx, req.(*TransRequestParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_TranslateOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransOneRequestParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).TranslateOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Service/TranslateOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).TranslateOne(ctx, req.(*TransOneRequestParam))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TranslateMultiple",
			Handler:    _Service_TranslateMultiple_Handler,
		},
		{
			MethodName: "TranslateOne",
			Handler:    _Service_TranslateOne_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trans.proto",
}
