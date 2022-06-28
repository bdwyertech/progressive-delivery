// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProgressiveDeliveryServiceClient is the client API for ProgressiveDeliveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProgressiveDeliveryServiceClient interface {
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error)
	//
	// ListCanaries returns with a list of Canary objects.
	ListCanaries(ctx context.Context, in *ListCanariesRequest, opts ...grpc.CallOption) (*ListCanariesResponse, error)
	//
	// GetCanary returns a Canary object.
	GetCanary(ctx context.Context, in *GetCanaryRequest, opts ...grpc.CallOption) (*GetCanaryResponse, error)
	//
	// IsFlaggerAvailable returns with a hashmap where the keys are the names of
	// the clusters, and the value is a boolean indicating whether Flagger is
	// installed or not on that cluster.
	IsFlaggerAvailable(ctx context.Context, in *IsFlaggerAvailableRequest, opts ...grpc.CallOption) (*IsFlaggerAvailableResponse, error)
	//
	// ListCanaries returns with a list of Canary objects.
	ListMetricTemplates(ctx context.Context, in *ListMetricTemplatesRequest, opts ...grpc.CallOption) (*ListMetricTemplatesResponse, error)
}

type progressiveDeliveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProgressiveDeliveryServiceClient(cc grpc.ClientConnInterface) ProgressiveDeliveryServiceClient {
	return &progressiveDeliveryServiceClient{cc}
}

func (c *progressiveDeliveryServiceClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error) {
	out := new(GetVersionResponse)
	err := c.cc.Invoke(ctx, "/ProgressiveDeliveryService/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressiveDeliveryServiceClient) ListCanaries(ctx context.Context, in *ListCanariesRequest, opts ...grpc.CallOption) (*ListCanariesResponse, error) {
	out := new(ListCanariesResponse)
	err := c.cc.Invoke(ctx, "/ProgressiveDeliveryService/ListCanaries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressiveDeliveryServiceClient) GetCanary(ctx context.Context, in *GetCanaryRequest, opts ...grpc.CallOption) (*GetCanaryResponse, error) {
	out := new(GetCanaryResponse)
	err := c.cc.Invoke(ctx, "/ProgressiveDeliveryService/GetCanary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressiveDeliveryServiceClient) IsFlaggerAvailable(ctx context.Context, in *IsFlaggerAvailableRequest, opts ...grpc.CallOption) (*IsFlaggerAvailableResponse, error) {
	out := new(IsFlaggerAvailableResponse)
	err := c.cc.Invoke(ctx, "/ProgressiveDeliveryService/IsFlaggerAvailable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressiveDeliveryServiceClient) ListMetricTemplates(ctx context.Context, in *ListMetricTemplatesRequest, opts ...grpc.CallOption) (*ListMetricTemplatesResponse, error) {
	out := new(ListMetricTemplatesResponse)
	err := c.cc.Invoke(ctx, "/ProgressiveDeliveryService/ListMetricTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProgressiveDeliveryServiceServer is the server API for ProgressiveDeliveryService service.
// All implementations must embed UnimplementedProgressiveDeliveryServiceServer
// for forward compatibility
type ProgressiveDeliveryServiceServer interface {
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)
	//
	// ListCanaries returns with a list of Canary objects.
	ListCanaries(context.Context, *ListCanariesRequest) (*ListCanariesResponse, error)
	//
	// GetCanary returns a Canary object.
	GetCanary(context.Context, *GetCanaryRequest) (*GetCanaryResponse, error)
	//
	// IsFlaggerAvailable returns with a hashmap where the keys are the names of
	// the clusters, and the value is a boolean indicating whether Flagger is
	// installed or not on that cluster.
	IsFlaggerAvailable(context.Context, *IsFlaggerAvailableRequest) (*IsFlaggerAvailableResponse, error)
	//
	// ListCanaries returns with a list of Canary objects.
	ListMetricTemplates(context.Context, *ListMetricTemplatesRequest) (*ListMetricTemplatesResponse, error)
	mustEmbedUnimplementedProgressiveDeliveryServiceServer()
}

// UnimplementedProgressiveDeliveryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProgressiveDeliveryServiceServer struct {
}

func (UnimplementedProgressiveDeliveryServiceServer) GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedProgressiveDeliveryServiceServer) ListCanaries(context.Context, *ListCanariesRequest) (*ListCanariesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCanaries not implemented")
}
func (UnimplementedProgressiveDeliveryServiceServer) GetCanary(context.Context, *GetCanaryRequest) (*GetCanaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCanary not implemented")
}
func (UnimplementedProgressiveDeliveryServiceServer) IsFlaggerAvailable(context.Context, *IsFlaggerAvailableRequest) (*IsFlaggerAvailableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFlaggerAvailable not implemented")
}
func (UnimplementedProgressiveDeliveryServiceServer) ListMetricTemplates(context.Context, *ListMetricTemplatesRequest) (*ListMetricTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMetricTemplates not implemented")
}
func (UnimplementedProgressiveDeliveryServiceServer) mustEmbedUnimplementedProgressiveDeliveryServiceServer() {
}

// UnsafeProgressiveDeliveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProgressiveDeliveryServiceServer will
// result in compilation errors.
type UnsafeProgressiveDeliveryServiceServer interface {
	mustEmbedUnimplementedProgressiveDeliveryServiceServer()
}

func RegisterProgressiveDeliveryServiceServer(s grpc.ServiceRegistrar, srv ProgressiveDeliveryServiceServer) {
	s.RegisterService(&ProgressiveDeliveryService_ServiceDesc, srv)
}

func _ProgressiveDeliveryService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressiveDeliveryServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProgressiveDeliveryService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressiveDeliveryServiceServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressiveDeliveryService_ListCanaries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCanariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressiveDeliveryServiceServer).ListCanaries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProgressiveDeliveryService/ListCanaries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressiveDeliveryServiceServer).ListCanaries(ctx, req.(*ListCanariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressiveDeliveryService_GetCanary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCanaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressiveDeliveryServiceServer).GetCanary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProgressiveDeliveryService/GetCanary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressiveDeliveryServiceServer).GetCanary(ctx, req.(*GetCanaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressiveDeliveryService_IsFlaggerAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsFlaggerAvailableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressiveDeliveryServiceServer).IsFlaggerAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProgressiveDeliveryService/IsFlaggerAvailable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressiveDeliveryServiceServer).IsFlaggerAvailable(ctx, req.(*IsFlaggerAvailableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressiveDeliveryService_ListMetricTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMetricTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressiveDeliveryServiceServer).ListMetricTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProgressiveDeliveryService/ListMetricTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressiveDeliveryServiceServer).ListMetricTemplates(ctx, req.(*ListMetricTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProgressiveDeliveryService_ServiceDesc is the grpc.ServiceDesc for ProgressiveDeliveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProgressiveDeliveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProgressiveDeliveryService",
	HandlerType: (*ProgressiveDeliveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _ProgressiveDeliveryService_GetVersion_Handler,
		},
		{
			MethodName: "ListCanaries",
			Handler:    _ProgressiveDeliveryService_ListCanaries_Handler,
		},
		{
			MethodName: "GetCanary",
			Handler:    _ProgressiveDeliveryService_GetCanary_Handler,
		},
		{
			MethodName: "IsFlaggerAvailable",
			Handler:    _ProgressiveDeliveryService_IsFlaggerAvailable_Handler,
		},
		{
			MethodName: "ListMetricTemplates",
			Handler:    _ProgressiveDeliveryService_ListMetricTemplates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/prog/prog.proto",
}
