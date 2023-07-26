// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: protobuf/todo.proto

package todoservice

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

// TodoClient is the client API for Todo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoClient interface {
	GetTodo(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoReply, error)
	GetUserTodos(ctx context.Context, in *UserTodosRequest, opts ...grpc.CallOption) (*UserTodosReply, error)
	GetUserTodosWithItems(ctx context.Context, in *UserTodosRequest, opts ...grpc.CallOption) (*UserTodosWithItemsReply, error)
	GetTodoWithItems(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoWithItemsReply, error)
	GetItem(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*ItemReply, error)
}

type todoClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoClient(cc grpc.ClientConnInterface) TodoClient {
	return &todoClient{cc}
}

func (c *todoClient) GetTodo(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoReply, error) {
	out := new(TodoReply)
	err := c.cc.Invoke(ctx, "/Todo/GetTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) GetUserTodos(ctx context.Context, in *UserTodosRequest, opts ...grpc.CallOption) (*UserTodosReply, error) {
	out := new(UserTodosReply)
	err := c.cc.Invoke(ctx, "/Todo/GetUserTodos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) GetUserTodosWithItems(ctx context.Context, in *UserTodosRequest, opts ...grpc.CallOption) (*UserTodosWithItemsReply, error) {
	out := new(UserTodosWithItemsReply)
	err := c.cc.Invoke(ctx, "/Todo/GetUserTodosWithItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) GetTodoWithItems(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoWithItemsReply, error) {
	out := new(TodoWithItemsReply)
	err := c.cc.Invoke(ctx, "/Todo/GetTodoWithItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) GetItem(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*ItemReply, error) {
	out := new(ItemReply)
	err := c.cc.Invoke(ctx, "/Todo/GetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServer is the server API for Todo service.
// All implementations must embed UnimplementedTodoServer
// for forward compatibility
type TodoServer interface {
	GetTodo(context.Context, *TodoRequest) (*TodoReply, error)
	GetUserTodos(context.Context, *UserTodosRequest) (*UserTodosReply, error)
	GetUserTodosWithItems(context.Context, *UserTodosRequest) (*UserTodosWithItemsReply, error)
	GetTodoWithItems(context.Context, *TodoRequest) (*TodoWithItemsReply, error)
	GetItem(context.Context, *ItemRequest) (*ItemReply, error)
	mustEmbedUnimplementedTodoServer()
}

// UnimplementedTodoServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServer struct {
}

func (UnimplementedTodoServer) GetTodo(context.Context, *TodoRequest) (*TodoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodo not implemented")
}
func (UnimplementedTodoServer) GetUserTodos(context.Context, *UserTodosRequest) (*UserTodosReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserTodos not implemented")
}
func (UnimplementedTodoServer) GetUserTodosWithItems(context.Context, *UserTodosRequest) (*UserTodosWithItemsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserTodosWithItems not implemented")
}
func (UnimplementedTodoServer) GetTodoWithItems(context.Context, *TodoRequest) (*TodoWithItemsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodoWithItems not implemented")
}
func (UnimplementedTodoServer) GetItem(context.Context, *ItemRequest) (*ItemReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
}
func (UnimplementedTodoServer) mustEmbedUnimplementedTodoServer() {}

// UnsafeTodoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServer will
// result in compilation errors.
type UnsafeTodoServer interface {
	mustEmbedUnimplementedTodoServer()
}

func RegisterTodoServer(s grpc.ServiceRegistrar, srv TodoServer) {
	s.RegisterService(&Todo_ServiceDesc, srv)
}

func _Todo_GetTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).GetTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Todo/GetTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).GetTodo(ctx, req.(*TodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_GetUserTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTodosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).GetUserTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Todo/GetUserTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).GetUserTodos(ctx, req.(*UserTodosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_GetUserTodosWithItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTodosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).GetUserTodosWithItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Todo/GetUserTodosWithItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).GetUserTodosWithItems(ctx, req.(*UserTodosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_GetTodoWithItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).GetTodoWithItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Todo/GetTodoWithItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).GetTodoWithItems(ctx, req.(*TodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Todo/GetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).GetItem(ctx, req.(*ItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Todo_ServiceDesc is the grpc.ServiceDesc for Todo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Todo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Todo",
	HandlerType: (*TodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTodo",
			Handler:    _Todo_GetTodo_Handler,
		},
		{
			MethodName: "GetUserTodos",
			Handler:    _Todo_GetUserTodos_Handler,
		},
		{
			MethodName: "GetUserTodosWithItems",
			Handler:    _Todo_GetUserTodosWithItems_Handler,
		},
		{
			MethodName: "GetTodoWithItems",
			Handler:    _Todo_GetTodoWithItems_Handler,
		},
		{
			MethodName: "GetItem",
			Handler:    _Todo_GetItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/todo.proto",
}
