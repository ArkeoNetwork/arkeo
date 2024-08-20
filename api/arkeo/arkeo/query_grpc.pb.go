// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package arkeo

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

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	FetchProvider(ctx context.Context, in *QueryFetchProviderRequest, opts ...grpc.CallOption) (*QueryFetchProviderResponse, error)
	ProviderAll(ctx context.Context, in *QueryAllProviderRequest, opts ...grpc.CallOption) (*QueryAllProviderResponse, error)
	FetchContract(ctx context.Context, in *QueryFetchContractRequest, opts ...grpc.CallOption) (*QueryFetchContractResponse, error)
	ContractAll(ctx context.Context, in *QueryAllContractRequest, opts ...grpc.CallOption) (*QueryAllContractResponse, error)
	// Queries an active contract by spender, provider and service.
	ActiveContract(ctx context.Context, in *QueryActiveContractRequest, opts ...grpc.CallOption) (*QueryActiveContractResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/arkeo.arkeo.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) FetchProvider(ctx context.Context, in *QueryFetchProviderRequest, opts ...grpc.CallOption) (*QueryFetchProviderResponse, error) {
	out := new(QueryFetchProviderResponse)
	err := c.cc.Invoke(ctx, "/arkeo.arkeo.Query/FetchProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ProviderAll(ctx context.Context, in *QueryAllProviderRequest, opts ...grpc.CallOption) (*QueryAllProviderResponse, error) {
	out := new(QueryAllProviderResponse)
	err := c.cc.Invoke(ctx, "/arkeo.arkeo.Query/ProviderAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) FetchContract(ctx context.Context, in *QueryFetchContractRequest, opts ...grpc.CallOption) (*QueryFetchContractResponse, error) {
	out := new(QueryFetchContractResponse)
	err := c.cc.Invoke(ctx, "/arkeo.arkeo.Query/FetchContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ContractAll(ctx context.Context, in *QueryAllContractRequest, opts ...grpc.CallOption) (*QueryAllContractResponse, error) {
	out := new(QueryAllContractResponse)
	err := c.cc.Invoke(ctx, "/arkeo.arkeo.Query/ContractAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ActiveContract(ctx context.Context, in *QueryActiveContractRequest, opts ...grpc.CallOption) (*QueryActiveContractResponse, error) {
	out := new(QueryActiveContractResponse)
	err := c.cc.Invoke(ctx, "/arkeo.arkeo.Query/ActiveContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	FetchProvider(context.Context, *QueryFetchProviderRequest) (*QueryFetchProviderResponse, error)
	ProviderAll(context.Context, *QueryAllProviderRequest) (*QueryAllProviderResponse, error)
	FetchContract(context.Context, *QueryFetchContractRequest) (*QueryFetchContractResponse, error)
	ContractAll(context.Context, *QueryAllContractRequest) (*QueryAllContractResponse, error)
	// Queries an active contract by spender, provider and service.
	ActiveContract(context.Context, *QueryActiveContractRequest) (*QueryActiveContractResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) FetchProvider(context.Context, *QueryFetchProviderRequest) (*QueryFetchProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchProvider not implemented")
}
func (UnimplementedQueryServer) ProviderAll(context.Context, *QueryAllProviderRequest) (*QueryAllProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProviderAll not implemented")
}
func (UnimplementedQueryServer) FetchContract(context.Context, *QueryFetchContractRequest) (*QueryFetchContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchContract not implemented")
}
func (UnimplementedQueryServer) ContractAll(context.Context, *QueryAllContractRequest) (*QueryAllContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContractAll not implemented")
}
func (UnimplementedQueryServer) ActiveContract(context.Context, *QueryActiveContractRequest) (*QueryActiveContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActiveContract not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arkeo.arkeo.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_FetchProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFetchProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FetchProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arkeo.arkeo.Query/FetchProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FetchProvider(ctx, req.(*QueryFetchProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ProviderAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ProviderAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arkeo.arkeo.Query/ProviderAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ProviderAll(ctx, req.(*QueryAllProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_FetchContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFetchContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FetchContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arkeo.arkeo.Query/FetchContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FetchContract(ctx, req.(*QueryFetchContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ContractAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ContractAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arkeo.arkeo.Query/ContractAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ContractAll(ctx, req.(*QueryAllContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ActiveContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryActiveContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ActiveContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arkeo.arkeo.Query/ActiveContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ActiveContract(ctx, req.(*QueryActiveContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "arkeo.arkeo.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "FetchProvider",
			Handler:    _Query_FetchProvider_Handler,
		},
		{
			MethodName: "ProviderAll",
			Handler:    _Query_ProviderAll_Handler,
		},
		{
			MethodName: "FetchContract",
			Handler:    _Query_FetchContract_Handler,
		},
		{
			MethodName: "ContractAll",
			Handler:    _Query_ContractAll_Handler,
		},
		{
			MethodName: "ActiveContract",
			Handler:    _Query_ActiveContract_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "arkeo/arkeo/query.proto",
}
