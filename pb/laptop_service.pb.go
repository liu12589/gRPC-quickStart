// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: laptop_service.proto

package pb

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

type CreateLaptopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Laptop *Laptop `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
}

func (x *CreateLaptopRequest) Reset() {
	*x = CreateLaptopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLaptopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLaptopRequest) ProtoMessage() {}

func (x *CreateLaptopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLaptopRequest.ProtoReflect.Descriptor instead.
func (*CreateLaptopRequest) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLaptopRequest) GetLaptop() *Laptop {
	if x != nil {
		return x.Laptop
	}
	return nil
}

type CreateLaptopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateLaptopResponse) Reset() {
	*x = CreateLaptopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLaptopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLaptopResponse) ProtoMessage() {}

func (x *CreateLaptopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLaptopResponse.ProtoReflect.Descriptor instead.
func (*CreateLaptopResponse) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLaptopResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_laptop_service_proto protoreflect.FileDescriptor

var file_laptop_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x52, 0x50, 0x43, 0x2e, 0x71, 0x75, 0x69,
	0x63, 0x6b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x43, 0x70, 0x75, 0x1a, 0x14,
	0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61,
	0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x6c,
	0x61, 0x70, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x52,
	0x50, 0x43, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x2e, 0x75, 0x73,
	0x65, 0x43, 0x70, 0x75, 0x2e, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x06, 0x6c, 0x61, 0x70,
	0x74, 0x6f, 0x70, 0x22, 0x26, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70,
	0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x7c, 0x0a, 0x0d, 0x4c,
	0x61, 0x70, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b, 0x0a, 0x0c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x12, 0x2b, 0x2e, 0x67,
	0x52, 0x50, 0x43, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x2e, 0x75,
	0x73, 0x65, 0x43, 0x70, 0x75, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70, 0x74,
	0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x67, 0x52, 0x50, 0x43,
	0x2e, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x43,
	0x70, 0x75, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_laptop_service_proto_rawDescOnce sync.Once
	file_laptop_service_proto_rawDescData = file_laptop_service_proto_rawDesc
)

func file_laptop_service_proto_rawDescGZIP() []byte {
	file_laptop_service_proto_rawDescOnce.Do(func() {
		file_laptop_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_laptop_service_proto_rawDescData)
	})
	return file_laptop_service_proto_rawDescData
}

var file_laptop_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_laptop_service_proto_goTypes = []interface{}{
	(*CreateLaptopRequest)(nil),  // 0: gRPC.quickStart.useCpu.CreateLaptopRequest
	(*CreateLaptopResponse)(nil), // 1: gRPC.quickStart.useCpu.CreateLaptopResponse
	(*Laptop)(nil),               // 2: gRPC.quickStart.useCpu.Laptop
}
var file_laptop_service_proto_depIdxs = []int32{
	2, // 0: gRPC.quickStart.useCpu.CreateLaptopRequest.laptop:type_name -> gRPC.quickStart.useCpu.Laptop
	0, // 1: gRPC.quickStart.useCpu.LaptopService.CreateLaptop:input_type -> gRPC.quickStart.useCpu.CreateLaptopRequest
	1, // 2: gRPC.quickStart.useCpu.LaptopService.CreateLaptop:output_type -> gRPC.quickStart.useCpu.CreateLaptopResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_laptop_service_proto_init() }
func file_laptop_service_proto_init() {
	if File_laptop_service_proto != nil {
		return
	}
	file_laptop_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_laptop_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLaptopRequest); i {
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
		file_laptop_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLaptopResponse); i {
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
			RawDescriptor: file_laptop_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_laptop_service_proto_goTypes,
		DependencyIndexes: file_laptop_service_proto_depIdxs,
		MessageInfos:      file_laptop_service_proto_msgTypes,
	}.Build()
	File_laptop_service_proto = out.File
	file_laptop_service_proto_rawDesc = nil
	file_laptop_service_proto_goTypes = nil
	file_laptop_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LaptopServiceClient is the client API for LaptopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LaptopServiceClient interface {
	CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLaptopResponse, error)
}

type laptopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLaptopServiceClient(cc grpc.ClientConnInterface) LaptopServiceClient {
	return &laptopServiceClient{cc}
}

func (c *laptopServiceClient) CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLaptopResponse, error) {
	out := new(CreateLaptopResponse)
	err := c.cc.Invoke(ctx, "/gRPC.quickStart.useCpu.LaptopService/CreateLaptop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LaptopServiceServer is the server API for LaptopService service.
type LaptopServiceServer interface {
	CreateLaptop(context.Context, *CreateLaptopRequest) (*CreateLaptopResponse, error)
}

// UnimplementedLaptopServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLaptopServiceServer struct {
}

func (*UnimplementedLaptopServiceServer) CreateLaptop(context.Context, *CreateLaptopRequest) (*CreateLaptopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLaptop not implemented")
}

func RegisterLaptopServiceServer(s *grpc.Server, srv LaptopServiceServer) {
	s.RegisterService(&_LaptopService_serviceDesc, srv)
}

func _LaptopService_CreateLaptop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLaptopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaptopServiceServer).CreateLaptop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gRPC.quickStart.useCpu.LaptopService/CreateLaptop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaptopServiceServer).CreateLaptop(ctx, req.(*CreateLaptopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LaptopService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gRPC.quickStart.useCpu.LaptopService",
	HandlerType: (*LaptopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLaptop",
			Handler:    _LaptopService_CreateLaptop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "laptop_service.proto",
}
