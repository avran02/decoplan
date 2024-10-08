// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.1
// source: chat-storage.proto

package pb

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
	ChatStorageService_SaveMessage_FullMethodName       = "/chat_storage.ChatStorageService/SaveMessage"
	ChatStorageService_GetMessages_FullMethodName       = "/chat_storage.ChatStorageService/GetMessages"
	ChatStorageService_DeleteMessage_FullMethodName     = "/chat_storage.ChatStorageService/DeleteMessage"
	ChatStorageService_CacheLastMessages_FullMethodName = "/chat_storage.ChatStorageService/CacheLastMessages"
	ChatStorageService_CreateChat_FullMethodName        = "/chat_storage.ChatStorageService/CreateChat"
)

// ChatStorageServiceClient is the client API for ChatStorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatStorageServiceClient interface {
	SaveMessage(ctx context.Context, in *SaveMessageRequest, opts ...grpc.CallOption) (*SaveMessageResponse, error)
	GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error)
	DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...grpc.CallOption) (*DeleteMessageResponse, error)
	CacheLastMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error)
	CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*CreateChatResponse, error)
}

type chatStorageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatStorageServiceClient(cc grpc.ClientConnInterface) ChatStorageServiceClient {
	return &chatStorageServiceClient{cc}
}

func (c *chatStorageServiceClient) SaveMessage(ctx context.Context, in *SaveMessageRequest, opts ...grpc.CallOption) (*SaveMessageResponse, error) {
	out := new(SaveMessageResponse)
	err := c.cc.Invoke(ctx, ChatStorageService_SaveMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatStorageServiceClient) GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error) {
	out := new(GetMessagesResponse)
	err := c.cc.Invoke(ctx, ChatStorageService_GetMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatStorageServiceClient) DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...grpc.CallOption) (*DeleteMessageResponse, error) {
	out := new(DeleteMessageResponse)
	err := c.cc.Invoke(ctx, ChatStorageService_DeleteMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatStorageServiceClient) CacheLastMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error) {
	out := new(GetMessagesResponse)
	err := c.cc.Invoke(ctx, ChatStorageService_CacheLastMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatStorageServiceClient) CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*CreateChatResponse, error) {
	out := new(CreateChatResponse)
	err := c.cc.Invoke(ctx, ChatStorageService_CreateChat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatStorageServiceServer is the server API for ChatStorageService service.
// All implementations must embed UnimplementedChatStorageServiceServer
// for forward compatibility
type ChatStorageServiceServer interface {
	SaveMessage(context.Context, *SaveMessageRequest) (*SaveMessageResponse, error)
	GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error)
	DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error)
	CacheLastMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error)
	CreateChat(context.Context, *CreateChatRequest) (*CreateChatResponse, error)
	mustEmbedUnimplementedChatStorageServiceServer()
}

// UnimplementedChatStorageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatStorageServiceServer struct {
}

func (UnimplementedChatStorageServiceServer) SaveMessage(context.Context, *SaveMessageRequest) (*SaveMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveMessage not implemented")
}
func (UnimplementedChatStorageServiceServer) GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedChatStorageServiceServer) DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessage not implemented")
}
func (UnimplementedChatStorageServiceServer) CacheLastMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CacheLastMessages not implemented")
}
func (UnimplementedChatStorageServiceServer) CreateChat(context.Context, *CreateChatRequest) (*CreateChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedChatStorageServiceServer) mustEmbedUnimplementedChatStorageServiceServer() {}

// UnsafeChatStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatStorageServiceServer will
// result in compilation errors.
type UnsafeChatStorageServiceServer interface {
	mustEmbedUnimplementedChatStorageServiceServer()
}

func RegisterChatStorageServiceServer(s grpc.ServiceRegistrar, srv ChatStorageServiceServer) {
	s.RegisterService(&ChatStorageService_ServiceDesc, srv)
}

func _ChatStorageService_SaveMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatStorageServiceServer).SaveMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatStorageService_SaveMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatStorageServiceServer).SaveMessage(ctx, req.(*SaveMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatStorageService_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatStorageServiceServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatStorageService_GetMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatStorageServiceServer).GetMessages(ctx, req.(*GetMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatStorageService_DeleteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatStorageServiceServer).DeleteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatStorageService_DeleteMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatStorageServiceServer).DeleteMessage(ctx, req.(*DeleteMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatStorageService_CacheLastMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatStorageServiceServer).CacheLastMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatStorageService_CacheLastMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatStorageServiceServer).CacheLastMessages(ctx, req.(*GetMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatStorageService_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatStorageServiceServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatStorageService_CreateChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatStorageServiceServer).CreateChat(ctx, req.(*CreateChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatStorageService_ServiceDesc is the grpc.ServiceDesc for ChatStorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatStorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat_storage.ChatStorageService",
	HandlerType: (*ChatStorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveMessage",
			Handler:    _ChatStorageService_SaveMessage_Handler,
		},
		{
			MethodName: "GetMessages",
			Handler:    _ChatStorageService_GetMessages_Handler,
		},
		{
			MethodName: "DeleteMessage",
			Handler:    _ChatStorageService_DeleteMessage_Handler,
		},
		{
			MethodName: "CacheLastMessages",
			Handler:    _ChatStorageService_CacheLastMessages_Handler,
		},
		{
			MethodName: "CreateChat",
			Handler:    _ChatStorageService_CreateChat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat-storage.proto",
}
