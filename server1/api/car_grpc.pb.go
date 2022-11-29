// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: api/car.proto

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

// CarServiceClient is the client API for CarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CarServiceClient interface {
	BuyCar(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetUserCars(ctx context.Context, in *GetCarsRequest, opts ...grpc.CallOption) (*GetCarsResponse, error)
	SellUserCar(ctx context.Context, in *GetCarsRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type carServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCarServiceClient(cc grpc.ClientConnInterface) CarServiceClient {
	return &carServiceClient{cc}
}

func (c *carServiceClient) BuyCar(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/models.CarService/BuyCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carServiceClient) GetUserCars(ctx context.Context, in *GetCarsRequest, opts ...grpc.CallOption) (*GetCarsResponse, error) {
	out := new(GetCarsResponse)
	err := c.cc.Invoke(ctx, "/models.CarService/GetUserCars", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carServiceClient) SellUserCar(ctx context.Context, in *GetCarsRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/models.CarService/SellUserCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarServiceServer is the server API for CarService service.
// All implementations must embed UnimplementedCarServiceServer
// for forward compatibility
type CarServiceServer interface {
	BuyCar(context.Context, *Request) (*Response, error)
	GetUserCars(context.Context, *GetCarsRequest) (*GetCarsResponse, error)
	SellUserCar(context.Context, *GetCarsRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedCarServiceServer()
}

// UnimplementedCarServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCarServiceServer struct {
}

func (UnimplementedCarServiceServer) BuyCar(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyCar not implemented")
}
func (UnimplementedCarServiceServer) GetUserCars(context.Context, *GetCarsRequest) (*GetCarsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCars not implemented")
}
func (UnimplementedCarServiceServer) SellUserCar(context.Context, *GetCarsRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellUserCar not implemented")
}
func (UnimplementedCarServiceServer) mustEmbedUnimplementedCarServiceServer() {}

// UnsafeCarServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CarServiceServer will
// result in compilation errors.
type UnsafeCarServiceServer interface {
	mustEmbedUnimplementedCarServiceServer()
}

func RegisterCarServiceServer(s grpc.ServiceRegistrar, srv CarServiceServer) {
	s.RegisterService(&CarService_ServiceDesc, srv)
}

func _CarService_BuyCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServiceServer).BuyCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.CarService/BuyCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServiceServer).BuyCar(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarService_GetUserCars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCarsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServiceServer).GetUserCars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.CarService/GetUserCars",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServiceServer).GetUserCars(ctx, req.(*GetCarsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarService_SellUserCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCarsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServiceServer).SellUserCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.CarService/SellUserCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServiceServer).SellUserCar(ctx, req.(*GetCarsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CarService_ServiceDesc is the grpc.ServiceDesc for CarService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CarService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "models.CarService",
	HandlerType: (*CarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuyCar",
			Handler:    _CarService_BuyCar_Handler,
		},
		{
			MethodName: "GetUserCars",
			Handler:    _CarService_GetUserCars_Handler,
		},
		{
			MethodName: "SellUserCar",
			Handler:    _CarService_SellUserCar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/car.proto",
}
