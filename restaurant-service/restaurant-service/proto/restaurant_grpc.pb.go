// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: proto/restaurant.proto

package proto

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
	RestaurantService_AddRestaurant_FullMethodName          = "/RestaurantService/AddRestaurant"
	RestaurantService_GetAllRestaurants_FullMethodName      = "/RestaurantService/GetAllRestaurants"
	RestaurantService_GetRestaurantByName_FullMethodName    = "/RestaurantService/GetRestaurantByName"
	RestaurantService_GetMenuItemsByCategory_FullMethodName = "/RestaurantService/GetMenuItemsByCategory"
)

// RestaurantServiceClient is the client API for RestaurantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RestaurantServiceClient interface {
	AddRestaurant(ctx context.Context, in *Restaurant, opts ...grpc.CallOption) (*Response, error)
	GetAllRestaurants(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RestaurantList, error)
	GetRestaurantByName(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*Restaurant, error)
	GetMenuItemsByCategory(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*MenuItemList, error)
}

type restaurantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRestaurantServiceClient(cc grpc.ClientConnInterface) RestaurantServiceClient {
	return &restaurantServiceClient{cc}
}

func (c *restaurantServiceClient) AddRestaurant(ctx context.Context, in *Restaurant, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, RestaurantService_AddRestaurant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) GetAllRestaurants(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RestaurantList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RestaurantList)
	err := c.cc.Invoke(ctx, RestaurantService_GetAllRestaurants_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) GetRestaurantByName(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*Restaurant, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Restaurant)
	err := c.cc.Invoke(ctx, RestaurantService_GetRestaurantByName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) GetMenuItemsByCategory(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*MenuItemList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MenuItemList)
	err := c.cc.Invoke(ctx, RestaurantService_GetMenuItemsByCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RestaurantServiceServer is the server API for RestaurantService service.
// All implementations must embed UnimplementedRestaurantServiceServer
// for forward compatibility.
type RestaurantServiceServer interface {
	AddRestaurant(context.Context, *Restaurant) (*Response, error)
	GetAllRestaurants(context.Context, *Empty) (*RestaurantList, error)
	GetRestaurantByName(context.Context, *NameRequest) (*Restaurant, error)
	GetMenuItemsByCategory(context.Context, *CategoryRequest) (*MenuItemList, error)
	mustEmbedUnimplementedRestaurantServiceServer()
}

// UnimplementedRestaurantServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRestaurantServiceServer struct{}

func (UnimplementedRestaurantServiceServer) AddRestaurant(context.Context, *Restaurant) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRestaurant not implemented")
}
func (UnimplementedRestaurantServiceServer) GetAllRestaurants(context.Context, *Empty) (*RestaurantList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRestaurants not implemented")
}
func (UnimplementedRestaurantServiceServer) GetRestaurantByName(context.Context, *NameRequest) (*Restaurant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRestaurantByName not implemented")
}
func (UnimplementedRestaurantServiceServer) GetMenuItemsByCategory(context.Context, *CategoryRequest) (*MenuItemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMenuItemsByCategory not implemented")
}
func (UnimplementedRestaurantServiceServer) mustEmbedUnimplementedRestaurantServiceServer() {}
func (UnimplementedRestaurantServiceServer) testEmbeddedByValue()                           {}

// UnsafeRestaurantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RestaurantServiceServer will
// result in compilation errors.
type UnsafeRestaurantServiceServer interface {
	mustEmbedUnimplementedRestaurantServiceServer()
}

func RegisterRestaurantServiceServer(s grpc.ServiceRegistrar, srv RestaurantServiceServer) {
	// If the following call pancis, it indicates UnimplementedRestaurantServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RestaurantService_ServiceDesc, srv)
}

func _RestaurantService_AddRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Restaurant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).AddRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RestaurantService_AddRestaurant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).AddRestaurant(ctx, req.(*Restaurant))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_GetAllRestaurants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).GetAllRestaurants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RestaurantService_GetAllRestaurants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).GetAllRestaurants(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_GetRestaurantByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).GetRestaurantByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RestaurantService_GetRestaurantByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).GetRestaurantByName(ctx, req.(*NameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_GetMenuItemsByCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).GetMenuItemsByCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RestaurantService_GetMenuItemsByCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).GetMenuItemsByCategory(ctx, req.(*CategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RestaurantService_ServiceDesc is the grpc.ServiceDesc for RestaurantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RestaurantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RestaurantService",
	HandlerType: (*RestaurantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRestaurant",
			Handler:    _RestaurantService_AddRestaurant_Handler,
		},
		{
			MethodName: "GetAllRestaurants",
			Handler:    _RestaurantService_GetAllRestaurants_Handler,
		},
		{
			MethodName: "GetRestaurantByName",
			Handler:    _RestaurantService_GetRestaurantByName_Handler,
		},
		{
			MethodName: "GetMenuItemsByCategory",
			Handler:    _RestaurantService_GetMenuItemsByCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/restaurant.proto",
}
