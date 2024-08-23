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

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	BondProvider(ctx context.Context, in *MsgBondProvider, opts ...grpc.CallOption) (*MsgBondProviderResponse, error)
	ModProvider(ctx context.Context, in *MsgModProvider, opts ...grpc.CallOption) (*MsgModProviderResponse, error)
	OpenContract(ctx context.Context, in *MsgOpenContract, opts ...grpc.CallOption) (*MsgOpenContractResponse, error)
	CloseContract(ctx context.Context, in *MsgCloseContract, opts ...grpc.CallOption) (*MsgCloseContractResponse, error)
	ClaimContractIncome(ctx context.Context, in *MsgClaimContractIncome, opts ...grpc.CallOption) (*MsgClaimContractIncomeResponse, error)
	// this line is used by starport scaffolding # proto/tx/rpc
	SetVersion(ctx context.Context, in *MsgSetVersion, opts ...grpc.CallOption) (*MsgSetVersionResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) BondProvider(ctx context.Context, in *MsgBondProvider, opts ...grpc.CallOption) (*MsgBondProviderResponse, error) {
	out := new(MsgBondProviderResponse)
	err := c.cc.Invoke(ctx, "/arkeo.arkeo.Msg/BondProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ModProvider(ctx context.Context, in *MsgModProvider, opts ...grpc.CallOption) (*MsgModProviderResponse, error) {
	out := new(MsgModProviderResponse)
	err := c.cc.Invoke(ctx, "/arkeo.arkeo.Msg/ModProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) OpenContract(ctx context.Context, in *MsgOpenContract, opts ...grpc.CallOption) (*MsgOpenContractResponse, error) {
	out := new(MsgOpenContractResponse)
	err := c.cc.Invoke(ctx, "/arkeo.arkeo.Msg/OpenContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CloseContract(ctx context.Context, in *MsgCloseContract, opts ...grpc.CallOption) (*MsgCloseContractResponse, error) {
	out := new(MsgCloseContractResponse)
	err := c.cc.Invoke(ctx, "/arkeo.arkeo.Msg/CloseContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClaimContractIncome(ctx context.Context, in *MsgClaimContractIncome, opts ...grpc.CallOption) (*MsgClaimContractIncomeResponse, error) {
	out := new(MsgClaimContractIncomeResponse)
	err := c.cc.Invoke(ctx, "/arkeo.arkeo.Msg/ClaimContractIncome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetVersion(ctx context.Context, in *MsgSetVersion, opts ...grpc.CallOption) (*MsgSetVersionResponse, error) {
	out := new(MsgSetVersionResponse)
	err := c.cc.Invoke(ctx, "/arkeo.arkeo.Msg/SetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	BondProvider(context.Context, *MsgBondProvider) (*MsgBondProviderResponse, error)
	ModProvider(context.Context, *MsgModProvider) (*MsgModProviderResponse, error)
	OpenContract(context.Context, *MsgOpenContract) (*MsgOpenContractResponse, error)
	CloseContract(context.Context, *MsgCloseContract) (*MsgCloseContractResponse, error)
	ClaimContractIncome(context.Context, *MsgClaimContractIncome) (*MsgClaimContractIncomeResponse, error)
	// this line is used by starport scaffolding # proto/tx/rpc
	SetVersion(context.Context, *MsgSetVersion) (*MsgSetVersionResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) BondProvider(context.Context, *MsgBondProvider) (*MsgBondProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BondProvider not implemented")
}
func (UnimplementedMsgServer) ModProvider(context.Context, *MsgModProvider) (*MsgModProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModProvider not implemented")
}
func (UnimplementedMsgServer) OpenContract(context.Context, *MsgOpenContract) (*MsgOpenContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenContract not implemented")
}
func (UnimplementedMsgServer) CloseContract(context.Context, *MsgCloseContract) (*MsgCloseContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseContract not implemented")
}
func (UnimplementedMsgServer) ClaimContractIncome(context.Context, *MsgClaimContractIncome) (*MsgClaimContractIncomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimContractIncome not implemented")
}
func (UnimplementedMsgServer) SetVersion(context.Context, *MsgSetVersion) (*MsgSetVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetVersion not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_BondProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBondProvider)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BondProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arkeo.arkeo.Msg/BondProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BondProvider(ctx, req.(*MsgBondProvider))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ModProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgModProvider)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ModProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arkeo.arkeo.Msg/ModProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ModProvider(ctx, req.(*MsgModProvider))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_OpenContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgOpenContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).OpenContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arkeo.arkeo.Msg/OpenContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).OpenContract(ctx, req.(*MsgOpenContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CloseContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCloseContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CloseContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arkeo.arkeo.Msg/CloseContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CloseContract(ctx, req.(*MsgCloseContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClaimContractIncome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaimContractIncome)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimContractIncome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arkeo.arkeo.Msg/ClaimContractIncome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimContractIncome(ctx, req.(*MsgClaimContractIncome))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetVersion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arkeo.arkeo.Msg/SetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetVersion(ctx, req.(*MsgSetVersion))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "arkeo.arkeo.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BondProvider",
			Handler:    _Msg_BondProvider_Handler,
		},
		{
			MethodName: "ModProvider",
			Handler:    _Msg_ModProvider_Handler,
		},
		{
			MethodName: "OpenContract",
			Handler:    _Msg_OpenContract_Handler,
		},
		{
			MethodName: "CloseContract",
			Handler:    _Msg_CloseContract_Handler,
		},
		{
			MethodName: "ClaimContractIncome",
			Handler:    _Msg_ClaimContractIncome_Handler,
		},
		{
			MethodName: "SetVersion",
			Handler:    _Msg_SetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "arkeo/arkeo/tx.proto",
}
