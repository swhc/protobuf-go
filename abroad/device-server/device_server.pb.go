// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: device_server.proto

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_device_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_device_server_proto_rawDescGZIP(), []int{0}
}

var File_device_server_proto protoreflect.FileDescriptor

var file_device_server_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x32, 0x6e, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a,
	0x0c, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x19, 0x2e,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x31, 0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x19, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x06, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_device_server_proto_rawDescOnce sync.Once
	file_device_server_proto_rawDescData = file_device_server_proto_rawDesc
)

func file_device_server_proto_rawDescGZIP() []byte {
	file_device_server_proto_rawDescOnce.Do(func() {
		file_device_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_device_server_proto_rawDescData)
	})
	return file_device_server_proto_rawDescData
}

var file_device_server_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_device_server_proto_goTypes = []interface{}{
	(*Empty)(nil),                    // 0: Empty
	(*CollectMatchRequestParam)(nil), // 1: CollectMatchRequestParam
	(*PushMsgMatchRequestParam)(nil), // 2: PushMsgMatchRequestParam
}
var file_device_server_proto_depIdxs = []int32{
	1, // 0: Device.CollectMatch:input_type -> CollectMatchRequestParam
	2, // 1: Device.PushMsgMatch:input_type -> PushMsgMatchRequestParam
	0, // 2: Device.CollectMatch:output_type -> Empty
	0, // 3: Device.PushMsgMatch:output_type -> Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_device_server_proto_init() }
func file_device_server_proto_init() {
	if File_device_server_proto != nil {
		return
	}
	file_collect_match_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_device_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_device_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_device_server_proto_goTypes,
		DependencyIndexes: file_device_server_proto_depIdxs,
		MessageInfos:      file_device_server_proto_msgTypes,
	}.Build()
	File_device_server_proto = out.File
	file_device_server_proto_rawDesc = nil
	file_device_server_proto_goTypes = nil
	file_device_server_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DeviceClient is the client API for Device service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceClient interface {
	// 收藏比赛 (取消也在)
	CollectMatch(ctx context.Context, in *CollectMatchRequestParam, opts ...grpc.CallOption) (*Empty, error)
	// 收藏比赛事件推送
	PushMsgMatch(ctx context.Context, in *PushMsgMatchRequestParam, opts ...grpc.CallOption) (*Empty, error)
}

type deviceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceClient(cc grpc.ClientConnInterface) DeviceClient {
	return &deviceClient{cc}
}

func (c *deviceClient) CollectMatch(ctx context.Context, in *CollectMatchRequestParam, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Device/CollectMatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) PushMsgMatch(ctx context.Context, in *PushMsgMatchRequestParam, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Device/PushMsgMatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceServer is the server API for Device service.
type DeviceServer interface {
	// 收藏比赛 (取消也在)
	CollectMatch(context.Context, *CollectMatchRequestParam) (*Empty, error)
	// 收藏比赛事件推送
	PushMsgMatch(context.Context, *PushMsgMatchRequestParam) (*Empty, error)
}

// UnimplementedDeviceServer can be embedded to have forward compatible implementations.
type UnimplementedDeviceServer struct {
}

func (*UnimplementedDeviceServer) CollectMatch(context.Context, *CollectMatchRequestParam) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectMatch not implemented")
}
func (*UnimplementedDeviceServer) PushMsgMatch(context.Context, *PushMsgMatchRequestParam) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushMsgMatch not implemented")
}

func RegisterDeviceServer(s *grpc.Server, srv DeviceServer) {
	s.RegisterService(&_Device_serviceDesc, srv)
}

func _Device_CollectMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectMatchRequestParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).CollectMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Device/CollectMatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).CollectMatch(ctx, req.(*CollectMatchRequestParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_PushMsgMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMsgMatchRequestParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).PushMsgMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Device/PushMsgMatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).PushMsgMatch(ctx, req.(*PushMsgMatchRequestParam))
	}
	return interceptor(ctx, in, info, handler)
}

var _Device_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Device",
	HandlerType: (*DeviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CollectMatch",
			Handler:    _Device_CollectMatch_Handler,
		},
		{
			MethodName: "PushMsgMatch",
			Handler:    _Device_PushMsgMatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "device_server.proto",
}