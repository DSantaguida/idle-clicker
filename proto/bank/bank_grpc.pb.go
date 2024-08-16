// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.2
// source: bank.proto

package bank

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

// BankServiceClient is the client API for BankService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankServiceClient interface {
	CreateBank(ctx context.Context, in *BankRequest, opts ...grpc.CallOption) (*BankResponse, error)
	GetBankData(ctx context.Context, in *GetBankDataRequest, opts ...grpc.CallOption) (*BankResponse, error)
	SetBankData(ctx context.Context, in *SetBankDataRequest, opts ...grpc.CallOption) (*BankResponse, error)
}

type bankServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBankServiceClient(cc grpc.ClientConnInterface) BankServiceClient {
	return &bankServiceClient{cc}
}

func (c *bankServiceClient) CreateBank(ctx context.Context, in *BankRequest, opts ...grpc.CallOption) (*BankResponse, error) {
	out := new(BankResponse)
	err := c.cc.Invoke(ctx, "/bank.BankService/CreateBank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) GetBankData(ctx context.Context, in *GetBankDataRequest, opts ...grpc.CallOption) (*BankResponse, error) {
	out := new(BankResponse)
	err := c.cc.Invoke(ctx, "/bank.BankService/GetBankData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) SetBankData(ctx context.Context, in *SetBankDataRequest, opts ...grpc.CallOption) (*BankResponse, error) {
	out := new(BankResponse)
	err := c.cc.Invoke(ctx, "/bank.BankService/SetBankData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankServiceServer is the server API for BankService service.
// All implementations must embed UnimplementedBankServiceServer
// for forward compatibility
type BankServiceServer interface {
	CreateBank(context.Context, *BankRequest) (*BankResponse, error)
	GetBankData(context.Context, *GetBankDataRequest) (*BankResponse, error)
	SetBankData(context.Context, *SetBankDataRequest) (*BankResponse, error)
	mustEmbedUnimplementedBankServiceServer()
}

// UnimplementedBankServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBankServiceServer struct {
}

func (UnimplementedBankServiceServer) CreateBank(context.Context, *BankRequest) (*BankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBank not implemented")
}
func (UnimplementedBankServiceServer) GetBankData(context.Context, *GetBankDataRequest) (*BankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBankData not implemented")
}
func (UnimplementedBankServiceServer) SetBankData(context.Context, *SetBankDataRequest) (*BankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBankData not implemented")
}
func (UnimplementedBankServiceServer) mustEmbedUnimplementedBankServiceServer() {}

// UnsafeBankServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankServiceServer will
// result in compilation errors.
type UnsafeBankServiceServer interface {
	mustEmbedUnimplementedBankServiceServer()
}

func RegisterBankServiceServer(s grpc.ServiceRegistrar, srv BankServiceServer) {
	s.RegisterService(&BankService_ServiceDesc, srv)
}

func _BankService_CreateBank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).CreateBank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bank.BankService/CreateBank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).CreateBank(ctx, req.(*BankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_GetBankData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBankDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).GetBankData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bank.BankService/GetBankData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).GetBankData(ctx, req.(*GetBankDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_SetBankData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetBankDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).SetBankData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bank.BankService/SetBankData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).SetBankData(ctx, req.(*SetBankDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BankService_ServiceDesc is the grpc.ServiceDesc for BankService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BankService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bank.BankService",
	HandlerType: (*BankServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBank",
			Handler:    _BankService_CreateBank_Handler,
		},
		{
			MethodName: "GetBankData",
			Handler:    _BankService_GetBankData_Handler,
		},
		{
			MethodName: "SetBankData",
			Handler:    _BankService_SetBankData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bank.proto",
}
