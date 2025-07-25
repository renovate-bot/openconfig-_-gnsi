// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: github.com/openconfig/gnsi/certz/certz.proto

package cert

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
	Certz_Rotate_FullMethodName               = "/gnsi.certz.v1.Certz/Rotate"
	Certz_AddProfile_FullMethodName           = "/gnsi.certz.v1.Certz/AddProfile"
	Certz_DeleteProfile_FullMethodName        = "/gnsi.certz.v1.Certz/DeleteProfile"
	Certz_GetProfileList_FullMethodName       = "/gnsi.certz.v1.Certz/GetProfileList"
	Certz_CanGenerateCSR_FullMethodName       = "/gnsi.certz.v1.Certz/CanGenerateCSR"
	Certz_GetIntegrityManifest_FullMethodName = "/gnsi.certz.v1.Certz/GetIntegrityManifest"
)

// CertzClient is the client API for Certz service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CertzClient interface {
	Rotate(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[RotateCertificateRequest, RotateCertificateResponse], error)
	AddProfile(ctx context.Context, in *AddProfileRequest, opts ...grpc.CallOption) (*AddProfileResponse, error)
	DeleteProfile(ctx context.Context, in *DeleteProfileRequest, opts ...grpc.CallOption) (*DeleteProfileResponse, error)
	GetProfileList(ctx context.Context, in *GetProfileListRequest, opts ...grpc.CallOption) (*GetProfileListResponse, error)
	CanGenerateCSR(ctx context.Context, in *CanGenerateCSRRequest, opts ...grpc.CallOption) (*CanGenerateCSRResponse, error)
	GetIntegrityManifest(ctx context.Context, in *GetIntegrityManifestRequest, opts ...grpc.CallOption) (*GetIntegrityManifestResponse, error)
}

type certzClient struct {
	cc grpc.ClientConnInterface
}

func NewCertzClient(cc grpc.ClientConnInterface) CertzClient {
	return &certzClient{cc}
}

func (c *certzClient) Rotate(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[RotateCertificateRequest, RotateCertificateResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Certz_ServiceDesc.Streams[0], Certz_Rotate_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[RotateCertificateRequest, RotateCertificateResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Certz_RotateClient = grpc.BidiStreamingClient[RotateCertificateRequest, RotateCertificateResponse]

func (c *certzClient) AddProfile(ctx context.Context, in *AddProfileRequest, opts ...grpc.CallOption) (*AddProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddProfileResponse)
	err := c.cc.Invoke(ctx, Certz_AddProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certzClient) DeleteProfile(ctx context.Context, in *DeleteProfileRequest, opts ...grpc.CallOption) (*DeleteProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteProfileResponse)
	err := c.cc.Invoke(ctx, Certz_DeleteProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certzClient) GetProfileList(ctx context.Context, in *GetProfileListRequest, opts ...grpc.CallOption) (*GetProfileListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProfileListResponse)
	err := c.cc.Invoke(ctx, Certz_GetProfileList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certzClient) CanGenerateCSR(ctx context.Context, in *CanGenerateCSRRequest, opts ...grpc.CallOption) (*CanGenerateCSRResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CanGenerateCSRResponse)
	err := c.cc.Invoke(ctx, Certz_CanGenerateCSR_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certzClient) GetIntegrityManifest(ctx context.Context, in *GetIntegrityManifestRequest, opts ...grpc.CallOption) (*GetIntegrityManifestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetIntegrityManifestResponse)
	err := c.cc.Invoke(ctx, Certz_GetIntegrityManifest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertzServer is the server API for Certz service.
// All implementations should embed UnimplementedCertzServer
// for forward compatibility.
type CertzServer interface {
	Rotate(grpc.BidiStreamingServer[RotateCertificateRequest, RotateCertificateResponse]) error
	AddProfile(context.Context, *AddProfileRequest) (*AddProfileResponse, error)
	DeleteProfile(context.Context, *DeleteProfileRequest) (*DeleteProfileResponse, error)
	GetProfileList(context.Context, *GetProfileListRequest) (*GetProfileListResponse, error)
	CanGenerateCSR(context.Context, *CanGenerateCSRRequest) (*CanGenerateCSRResponse, error)
	GetIntegrityManifest(context.Context, *GetIntegrityManifestRequest) (*GetIntegrityManifestResponse, error)
}

// UnimplementedCertzServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCertzServer struct{}

func (UnimplementedCertzServer) Rotate(grpc.BidiStreamingServer[RotateCertificateRequest, RotateCertificateResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Rotate not implemented")
}
func (UnimplementedCertzServer) AddProfile(context.Context, *AddProfileRequest) (*AddProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProfile not implemented")
}
func (UnimplementedCertzServer) DeleteProfile(context.Context, *DeleteProfileRequest) (*DeleteProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProfile not implemented")
}
func (UnimplementedCertzServer) GetProfileList(context.Context, *GetProfileListRequest) (*GetProfileListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileList not implemented")
}
func (UnimplementedCertzServer) CanGenerateCSR(context.Context, *CanGenerateCSRRequest) (*CanGenerateCSRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CanGenerateCSR not implemented")
}
func (UnimplementedCertzServer) GetIntegrityManifest(context.Context, *GetIntegrityManifestRequest) (*GetIntegrityManifestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIntegrityManifest not implemented")
}
func (UnimplementedCertzServer) testEmbeddedByValue() {}

// UnsafeCertzServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertzServer will
// result in compilation errors.
type UnsafeCertzServer interface {
	mustEmbedUnimplementedCertzServer()
}

func RegisterCertzServer(s grpc.ServiceRegistrar, srv CertzServer) {
	// If the following call pancis, it indicates UnimplementedCertzServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Certz_ServiceDesc, srv)
}

func _Certz_Rotate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CertzServer).Rotate(&grpc.GenericServerStream[RotateCertificateRequest, RotateCertificateResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Certz_RotateServer = grpc.BidiStreamingServer[RotateCertificateRequest, RotateCertificateResponse]

func _Certz_AddProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertzServer).AddProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Certz_AddProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertzServer).AddProfile(ctx, req.(*AddProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certz_DeleteProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertzServer).DeleteProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Certz_DeleteProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertzServer).DeleteProfile(ctx, req.(*DeleteProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certz_GetProfileList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertzServer).GetProfileList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Certz_GetProfileList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertzServer).GetProfileList(ctx, req.(*GetProfileListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certz_CanGenerateCSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CanGenerateCSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertzServer).CanGenerateCSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Certz_CanGenerateCSR_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertzServer).CanGenerateCSR(ctx, req.(*CanGenerateCSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certz_GetIntegrityManifest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIntegrityManifestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertzServer).GetIntegrityManifest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Certz_GetIntegrityManifest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertzServer).GetIntegrityManifest(ctx, req.(*GetIntegrityManifestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Certz_ServiceDesc is the grpc.ServiceDesc for Certz service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Certz_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnsi.certz.v1.Certz",
	HandlerType: (*CertzServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProfile",
			Handler:    _Certz_AddProfile_Handler,
		},
		{
			MethodName: "DeleteProfile",
			Handler:    _Certz_DeleteProfile_Handler,
		},
		{
			MethodName: "GetProfileList",
			Handler:    _Certz_GetProfileList_Handler,
		},
		{
			MethodName: "CanGenerateCSR",
			Handler:    _Certz_CanGenerateCSR_Handler,
		},
		{
			MethodName: "GetIntegrityManifest",
			Handler:    _Certz_GetIntegrityManifest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Rotate",
			Handler:       _Certz_Rotate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/openconfig/gnsi/certz/certz.proto",
}
