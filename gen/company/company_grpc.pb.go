// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: company/company.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CompanyService_GetCompanyBotMessages_FullMethodName = "/company.v1.CompanyService/GetCompanyBotMessages"
	CompanyService_GetCompanyExtensions_FullMethodName  = "/company.v1.CompanyService/GetCompanyExtensions"
)

// CompanyServiceClient is the client API for CompanyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompanyServiceClient interface {
	GetCompanyBotMessages(ctx context.Context, in *GetCompanyBotMessagesRequest, opts ...grpc.CallOption) (*GetCompanyBotMessagesResponse, error)
	GetCompanyExtensions(ctx context.Context, in *GetCompanyExtensionsRequest, opts ...grpc.CallOption) (*GetCompanyExtensionsResponse, error)
}

type companyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyServiceClient(cc grpc.ClientConnInterface) CompanyServiceClient {
	return &companyServiceClient{cc}
}

func (c *companyServiceClient) GetCompanyBotMessages(ctx context.Context, in *GetCompanyBotMessagesRequest, opts ...grpc.CallOption) (*GetCompanyBotMessagesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCompanyBotMessagesResponse)
	err := c.cc.Invoke(ctx, CompanyService_GetCompanyBotMessages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) GetCompanyExtensions(ctx context.Context, in *GetCompanyExtensionsRequest, opts ...grpc.CallOption) (*GetCompanyExtensionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCompanyExtensionsResponse)
	err := c.cc.Invoke(ctx, CompanyService_GetCompanyExtensions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyServiceServer is the server API for CompanyService service.
// All implementations should embed UnimplementedCompanyServiceServer
// for forward compatibility.
type CompanyServiceServer interface {
	GetCompanyBotMessages(context.Context, *GetCompanyBotMessagesRequest) (*GetCompanyBotMessagesResponse, error)
	GetCompanyExtensions(context.Context, *GetCompanyExtensionsRequest) (*GetCompanyExtensionsResponse, error)
}

// UnimplementedCompanyServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCompanyServiceServer struct{}

func (UnimplementedCompanyServiceServer) GetCompanyBotMessages(context.Context, *GetCompanyBotMessagesRequest) (*GetCompanyBotMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyBotMessages not implemented")
}
func (UnimplementedCompanyServiceServer) GetCompanyExtensions(context.Context, *GetCompanyExtensionsRequest) (*GetCompanyExtensionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyExtensions not implemented")
}
func (UnimplementedCompanyServiceServer) testEmbeddedByValue() {}

// UnsafeCompanyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompanyServiceServer will
// result in compilation errors.
type UnsafeCompanyServiceServer interface {
	mustEmbedUnimplementedCompanyServiceServer()
}

func RegisterCompanyServiceServer(s grpc.ServiceRegistrar, srv CompanyServiceServer) {
	// If the following call pancis, it indicates UnimplementedCompanyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CompanyService_ServiceDesc, srv)
}

func _CompanyService_GetCompanyBotMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyBotMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).GetCompanyBotMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyService_GetCompanyBotMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).GetCompanyBotMessages(ctx, req.(*GetCompanyBotMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_GetCompanyExtensions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyExtensionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).GetCompanyExtensions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyService_GetCompanyExtensions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).GetCompanyExtensions(ctx, req.(*GetCompanyExtensionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CompanyService_ServiceDesc is the grpc.ServiceDesc for CompanyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompanyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "company.v1.CompanyService",
	HandlerType: (*CompanyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCompanyBotMessages",
			Handler:    _CompanyService_GetCompanyBotMessages_Handler,
		},
		{
			MethodName: "GetCompanyExtensions",
			Handler:    _CompanyService_GetCompanyExtensions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "company/company.proto",
}
