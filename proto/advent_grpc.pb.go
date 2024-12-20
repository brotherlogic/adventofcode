// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: advent.proto

package proto

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

const (
	AdventOfCodeService_Solve_FullMethodName = "/adventofcode.AdventOfCodeService/Solve"
)

// AdventOfCodeServiceClient is the client API for AdventOfCodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdventOfCodeServiceClient interface {
	Solve(ctx context.Context, in *SolveRequest, opts ...grpc.CallOption) (*SolveResponse, error)
}

type adventOfCodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdventOfCodeServiceClient(cc grpc.ClientConnInterface) AdventOfCodeServiceClient {
	return &adventOfCodeServiceClient{cc}
}

func (c *adventOfCodeServiceClient) Solve(ctx context.Context, in *SolveRequest, opts ...grpc.CallOption) (*SolveResponse, error) {
	out := new(SolveResponse)
	err := c.cc.Invoke(ctx, AdventOfCodeService_Solve_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdventOfCodeServiceServer is the server API for AdventOfCodeService service.
// All implementations should embed UnimplementedAdventOfCodeServiceServer
// for forward compatibility
type AdventOfCodeServiceServer interface {
	Solve(context.Context, *SolveRequest) (*SolveResponse, error)
}

// UnimplementedAdventOfCodeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAdventOfCodeServiceServer struct {
}

func (UnimplementedAdventOfCodeServiceServer) Solve(context.Context, *SolveRequest) (*SolveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Solve not implemented")
}

// UnsafeAdventOfCodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdventOfCodeServiceServer will
// result in compilation errors.
type UnsafeAdventOfCodeServiceServer interface {
	mustEmbedUnimplementedAdventOfCodeServiceServer()
}

func RegisterAdventOfCodeServiceServer(s grpc.ServiceRegistrar, srv AdventOfCodeServiceServer) {
	s.RegisterService(&AdventOfCodeService_ServiceDesc, srv)
}

func _AdventOfCodeService_Solve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventOfCodeServiceServer).Solve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdventOfCodeService_Solve_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventOfCodeServiceServer).Solve(ctx, req.(*SolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdventOfCodeService_ServiceDesc is the grpc.ServiceDesc for AdventOfCodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdventOfCodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "adventofcode.AdventOfCodeService",
	HandlerType: (*AdventOfCodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Solve",
			Handler:    _AdventOfCodeService_Solve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "advent.proto",
}

const (
	AdventOfCodeInternalService_Upload_FullMethodName      = "/adventofcode.AdventOfCodeInternalService/Upload"
	AdventOfCodeInternalService_Register_FullMethodName    = "/adventofcode.AdventOfCodeInternalService/Register"
	AdventOfCodeInternalService_AddSolution_FullMethodName = "/adventofcode.AdventOfCodeInternalService/AddSolution"
	AdventOfCodeInternalService_GetSolution_FullMethodName = "/adventofcode.AdventOfCodeInternalService/GetSolution"
	AdventOfCodeInternalService_SetCookie_FullMethodName   = "/adventofcode.AdventOfCodeInternalService/SetCookie"
)

// AdventOfCodeInternalServiceClient is the client API for AdventOfCodeInternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdventOfCodeInternalServiceClient interface {
	Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	AddSolution(ctx context.Context, in *AddSolutionRequest, opts ...grpc.CallOption) (*AddSolutionResponse, error)
	GetSolution(ctx context.Context, in *GetSolutionRequest, opts ...grpc.CallOption) (*GetSolutionResponse, error)
	SetCookie(ctx context.Context, in *SetCookieRequest, opts ...grpc.CallOption) (*SetCookieResponse, error)
}

type adventOfCodeInternalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdventOfCodeInternalServiceClient(cc grpc.ClientConnInterface) AdventOfCodeInternalServiceClient {
	return &adventOfCodeInternalServiceClient{cc}
}

func (c *adventOfCodeInternalServiceClient) Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error) {
	out := new(UploadResponse)
	err := c.cc.Invoke(ctx, AdventOfCodeInternalService_Upload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventOfCodeInternalServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, AdventOfCodeInternalService_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventOfCodeInternalServiceClient) AddSolution(ctx context.Context, in *AddSolutionRequest, opts ...grpc.CallOption) (*AddSolutionResponse, error) {
	out := new(AddSolutionResponse)
	err := c.cc.Invoke(ctx, AdventOfCodeInternalService_AddSolution_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventOfCodeInternalServiceClient) GetSolution(ctx context.Context, in *GetSolutionRequest, opts ...grpc.CallOption) (*GetSolutionResponse, error) {
	out := new(GetSolutionResponse)
	err := c.cc.Invoke(ctx, AdventOfCodeInternalService_GetSolution_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventOfCodeInternalServiceClient) SetCookie(ctx context.Context, in *SetCookieRequest, opts ...grpc.CallOption) (*SetCookieResponse, error) {
	out := new(SetCookieResponse)
	err := c.cc.Invoke(ctx, AdventOfCodeInternalService_SetCookie_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdventOfCodeInternalServiceServer is the server API for AdventOfCodeInternalService service.
// All implementations should embed UnimplementedAdventOfCodeInternalServiceServer
// for forward compatibility
type AdventOfCodeInternalServiceServer interface {
	Upload(context.Context, *UploadRequest) (*UploadResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	AddSolution(context.Context, *AddSolutionRequest) (*AddSolutionResponse, error)
	GetSolution(context.Context, *GetSolutionRequest) (*GetSolutionResponse, error)
	SetCookie(context.Context, *SetCookieRequest) (*SetCookieResponse, error)
}

// UnimplementedAdventOfCodeInternalServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAdventOfCodeInternalServiceServer struct {
}

func (UnimplementedAdventOfCodeInternalServiceServer) Upload(context.Context, *UploadRequest) (*UploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedAdventOfCodeInternalServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAdventOfCodeInternalServiceServer) AddSolution(context.Context, *AddSolutionRequest) (*AddSolutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSolution not implemented")
}
func (UnimplementedAdventOfCodeInternalServiceServer) GetSolution(context.Context, *GetSolutionRequest) (*GetSolutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSolution not implemented")
}
func (UnimplementedAdventOfCodeInternalServiceServer) SetCookie(context.Context, *SetCookieRequest) (*SetCookieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCookie not implemented")
}

// UnsafeAdventOfCodeInternalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdventOfCodeInternalServiceServer will
// result in compilation errors.
type UnsafeAdventOfCodeInternalServiceServer interface {
	mustEmbedUnimplementedAdventOfCodeInternalServiceServer()
}

func RegisterAdventOfCodeInternalServiceServer(s grpc.ServiceRegistrar, srv AdventOfCodeInternalServiceServer) {
	s.RegisterService(&AdventOfCodeInternalService_ServiceDesc, srv)
}

func _AdventOfCodeInternalService_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventOfCodeInternalServiceServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdventOfCodeInternalService_Upload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventOfCodeInternalServiceServer).Upload(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdventOfCodeInternalService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventOfCodeInternalServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdventOfCodeInternalService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventOfCodeInternalServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdventOfCodeInternalService_AddSolution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSolutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventOfCodeInternalServiceServer).AddSolution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdventOfCodeInternalService_AddSolution_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventOfCodeInternalServiceServer).AddSolution(ctx, req.(*AddSolutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdventOfCodeInternalService_GetSolution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSolutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventOfCodeInternalServiceServer).GetSolution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdventOfCodeInternalService_GetSolution_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventOfCodeInternalServiceServer).GetSolution(ctx, req.(*GetSolutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdventOfCodeInternalService_SetCookie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCookieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventOfCodeInternalServiceServer).SetCookie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdventOfCodeInternalService_SetCookie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventOfCodeInternalServiceServer).SetCookie(ctx, req.(*SetCookieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdventOfCodeInternalService_ServiceDesc is the grpc.ServiceDesc for AdventOfCodeInternalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdventOfCodeInternalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "adventofcode.AdventOfCodeInternalService",
	HandlerType: (*AdventOfCodeInternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upload",
			Handler:    _AdventOfCodeInternalService_Upload_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _AdventOfCodeInternalService_Register_Handler,
		},
		{
			MethodName: "AddSolution",
			Handler:    _AdventOfCodeInternalService_AddSolution_Handler,
		},
		{
			MethodName: "GetSolution",
			Handler:    _AdventOfCodeInternalService_GetSolution_Handler,
		},
		{
			MethodName: "SetCookie",
			Handler:    _AdventOfCodeInternalService_SetCookie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "advent.proto",
}

const (
	SolverService_Solve_FullMethodName = "/adventofcode.SolverService/Solve"
)

// SolverServiceClient is the client API for SolverService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SolverServiceClient interface {
	Solve(ctx context.Context, in *SolveRequest, opts ...grpc.CallOption) (*SolveResponse, error)
}

type solverServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSolverServiceClient(cc grpc.ClientConnInterface) SolverServiceClient {
	return &solverServiceClient{cc}
}

func (c *solverServiceClient) Solve(ctx context.Context, in *SolveRequest, opts ...grpc.CallOption) (*SolveResponse, error) {
	out := new(SolveResponse)
	err := c.cc.Invoke(ctx, SolverService_Solve_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SolverServiceServer is the server API for SolverService service.
// All implementations should embed UnimplementedSolverServiceServer
// for forward compatibility
type SolverServiceServer interface {
	Solve(context.Context, *SolveRequest) (*SolveResponse, error)
}

// UnimplementedSolverServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSolverServiceServer struct {
}

func (UnimplementedSolverServiceServer) Solve(context.Context, *SolveRequest) (*SolveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Solve not implemented")
}

// UnsafeSolverServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SolverServiceServer will
// result in compilation errors.
type UnsafeSolverServiceServer interface {
	mustEmbedUnimplementedSolverServiceServer()
}

func RegisterSolverServiceServer(s grpc.ServiceRegistrar, srv SolverServiceServer) {
	s.RegisterService(&SolverService_ServiceDesc, srv)
}

func _SolverService_Solve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolverServiceServer).Solve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolverService_Solve_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolverServiceServer).Solve(ctx, req.(*SolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SolverService_ServiceDesc is the grpc.ServiceDesc for SolverService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SolverService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "adventofcode.SolverService",
	HandlerType: (*SolverServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Solve",
			Handler:    _SolverService_Solve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "advent.proto",
}
