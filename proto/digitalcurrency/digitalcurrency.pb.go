// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        (unknown)
// source: proto/digitalcurrency/digitalcurrency.proto

package digitalcurrency

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ConvertToSpecificCurrencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount        int64  `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	BaseCurrency  string `protobuf:"bytes,2,opt,name=baseCurrency,proto3" json:"baseCurrency,omitempty"`
	QuoteCurrency string `protobuf:"bytes,3,opt,name=quoteCurrency,proto3" json:"quoteCurrency,omitempty"`
}

func (x *ConvertToSpecificCurrencyRequest) Reset() {
	*x = ConvertToSpecificCurrencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_digitalcurrency_digitalcurrency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertToSpecificCurrencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertToSpecificCurrencyRequest) ProtoMessage() {}

func (x *ConvertToSpecificCurrencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_digitalcurrency_digitalcurrency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertToSpecificCurrencyRequest.ProtoReflect.Descriptor instead.
func (*ConvertToSpecificCurrencyRequest) Descriptor() ([]byte, []int) {
	return file_proto_digitalcurrency_digitalcurrency_proto_rawDescGZIP(), []int{0}
}

func (x *ConvertToSpecificCurrencyRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ConvertToSpecificCurrencyRequest) GetBaseCurrency() string {
	if x != nil {
		return x.BaseCurrency
	}
	return ""
}

func (x *ConvertToSpecificCurrencyRequest) GetQuoteCurrency() string {
	if x != nil {
		return x.QuoteCurrency
	}
	return ""
}

type ConvertToSpecificCurrencyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount        float32 `protobuf:"fixed32,1,opt,name=amount,proto3" json:"amount,omitempty"`
	BaseCurrency  string  `protobuf:"bytes,2,opt,name=baseCurrency,proto3" json:"baseCurrency,omitempty"`
	QuoteCurrency string  `protobuf:"bytes,3,opt,name=quoteCurrency,proto3" json:"quoteCurrency,omitempty"`
}

func (x *ConvertToSpecificCurrencyResponse) Reset() {
	*x = ConvertToSpecificCurrencyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_digitalcurrency_digitalcurrency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertToSpecificCurrencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertToSpecificCurrencyResponse) ProtoMessage() {}

func (x *ConvertToSpecificCurrencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_digitalcurrency_digitalcurrency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertToSpecificCurrencyResponse.ProtoReflect.Descriptor instead.
func (*ConvertToSpecificCurrencyResponse) Descriptor() ([]byte, []int) {
	return file_proto_digitalcurrency_digitalcurrency_proto_rawDescGZIP(), []int{1}
}

func (x *ConvertToSpecificCurrencyResponse) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ConvertToSpecificCurrencyResponse) GetBaseCurrency() string {
	if x != nil {
		return x.BaseCurrency
	}
	return ""
}

func (x *ConvertToSpecificCurrencyResponse) GetQuoteCurrency() string {
	if x != nil {
		return x.QuoteCurrency
	}
	return ""
}

var File_proto_digitalcurrency_digitalcurrency_proto protoreflect.FileDescriptor

var file_proto_digitalcurrency_digitalcurrency_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01,
	0x0a, 0x20, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x54, 0x6f, 0x53, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x61,
	0x73, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x24,
	0x0a, 0x0d, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x22, 0x85, 0x01, 0x0a, 0x21, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x54, 0x6f, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x71,
	0x75, 0x6f, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x32, 0x77, 0x0a, 0x0f,
	0x44, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x64, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x54, 0x6f, 0x53, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x63, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x21, 0x2e, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x54, 0x6f, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x54, 0x6f, 0x53, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64,
	0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_digitalcurrency_digitalcurrency_proto_rawDescOnce sync.Once
	file_proto_digitalcurrency_digitalcurrency_proto_rawDescData = file_proto_digitalcurrency_digitalcurrency_proto_rawDesc
)

func file_proto_digitalcurrency_digitalcurrency_proto_rawDescGZIP() []byte {
	file_proto_digitalcurrency_digitalcurrency_proto_rawDescOnce.Do(func() {
		file_proto_digitalcurrency_digitalcurrency_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_digitalcurrency_digitalcurrency_proto_rawDescData)
	})
	return file_proto_digitalcurrency_digitalcurrency_proto_rawDescData
}

var file_proto_digitalcurrency_digitalcurrency_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_digitalcurrency_digitalcurrency_proto_goTypes = []interface{}{
	(*ConvertToSpecificCurrencyRequest)(nil),  // 0: ConvertToSpecificCurrencyRequest
	(*ConvertToSpecificCurrencyResponse)(nil), // 1: ConvertToSpecificCurrencyResponse
}
var file_proto_digitalcurrency_digitalcurrency_proto_depIdxs = []int32{
	0, // 0: DigitalCurrency.ConvertToSpecificCurrency:input_type -> ConvertToSpecificCurrencyRequest
	1, // 1: DigitalCurrency.ConvertToSpecificCurrency:output_type -> ConvertToSpecificCurrencyResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_digitalcurrency_digitalcurrency_proto_init() }
func file_proto_digitalcurrency_digitalcurrency_proto_init() {
	if File_proto_digitalcurrency_digitalcurrency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_digitalcurrency_digitalcurrency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertToSpecificCurrencyRequest); i {
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
		file_proto_digitalcurrency_digitalcurrency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertToSpecificCurrencyResponse); i {
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
			RawDescriptor: file_proto_digitalcurrency_digitalcurrency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_digitalcurrency_digitalcurrency_proto_goTypes,
		DependencyIndexes: file_proto_digitalcurrency_digitalcurrency_proto_depIdxs,
		MessageInfos:      file_proto_digitalcurrency_digitalcurrency_proto_msgTypes,
	}.Build()
	File_proto_digitalcurrency_digitalcurrency_proto = out.File
	file_proto_digitalcurrency_digitalcurrency_proto_rawDesc = nil
	file_proto_digitalcurrency_digitalcurrency_proto_goTypes = nil
	file_proto_digitalcurrency_digitalcurrency_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DigitalCurrencyClient is the client API for DigitalCurrency service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DigitalCurrencyClient interface {
	ConvertToSpecificCurrency(ctx context.Context, in *ConvertToSpecificCurrencyRequest, opts ...grpc.CallOption) (*ConvertToSpecificCurrencyResponse, error)
}

type digitalCurrencyClient struct {
	cc grpc.ClientConnInterface
}

func NewDigitalCurrencyClient(cc grpc.ClientConnInterface) DigitalCurrencyClient {
	return &digitalCurrencyClient{cc}
}

func (c *digitalCurrencyClient) ConvertToSpecificCurrency(ctx context.Context, in *ConvertToSpecificCurrencyRequest, opts ...grpc.CallOption) (*ConvertToSpecificCurrencyResponse, error) {
	out := new(ConvertToSpecificCurrencyResponse)
	err := c.cc.Invoke(ctx, "/DigitalCurrency/ConvertToSpecificCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DigitalCurrencyServer is the server API for DigitalCurrency service.
type DigitalCurrencyServer interface {
	ConvertToSpecificCurrency(context.Context, *ConvertToSpecificCurrencyRequest) (*ConvertToSpecificCurrencyResponse, error)
}

// UnimplementedDigitalCurrencyServer can be embedded to have forward compatible implementations.
type UnimplementedDigitalCurrencyServer struct {
}

func (*UnimplementedDigitalCurrencyServer) ConvertToSpecificCurrency(context.Context, *ConvertToSpecificCurrencyRequest) (*ConvertToSpecificCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertToSpecificCurrency not implemented")
}

func RegisterDigitalCurrencyServer(s *grpc.Server, srv DigitalCurrencyServer) {
	s.RegisterService(&_DigitalCurrency_serviceDesc, srv)
}

func _DigitalCurrency_ConvertToSpecificCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvertToSpecificCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DigitalCurrencyServer).ConvertToSpecificCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DigitalCurrency/ConvertToSpecificCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DigitalCurrencyServer).ConvertToSpecificCurrency(ctx, req.(*ConvertToSpecificCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DigitalCurrency_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DigitalCurrency",
	HandlerType: (*DigitalCurrencyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConvertToSpecificCurrency",
			Handler:    _DigitalCurrency_ConvertToSpecificCurrency_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/digitalcurrency/digitalcurrency.proto",
}