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

// GameClient is the client API for Game service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameClient interface {
	GetStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StatusReply, error)
	Run(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*StatusReply, error)
}

type gameClient struct {
	cc grpc.ClientConnInterface
}

func NewGameClient(cc grpc.ClientConnInterface) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) GetStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/tictactoe.Game/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) Run(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/tictactoe.Game/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServer is the server API for Game service.
// All implementations must embed UnimplementedGameServer
// for forward compatibility
type GameServer interface {
	GetStatus(context.Context, *Empty) (*StatusReply, error)
	Run(context.Context, *CommandRequest) (*StatusReply, error)
	mustEmbedUnimplementedGameServer()
}

// UnimplementedGameServer must be embedded to have forward compatible implementations.
type UnimplementedGameServer struct {
}

func (UnimplementedGameServer) GetStatus(context.Context, *Empty) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedGameServer) Run(context.Context, *CommandRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (UnimplementedGameServer) mustEmbedUnimplementedGameServer() {}

// UnsafeGameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServer will
// result in compilation errors.
type UnsafeGameServer interface {
	mustEmbedUnimplementedGameServer()
}

func RegisterGameServer(s grpc.ServiceRegistrar, srv GameServer) {
	s.RegisterService(&Game_ServiceDesc, srv)
}

func _Game_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tictactoe.Game/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).GetStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tictactoe.Game/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).Run(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Game_ServiceDesc is the grpc.ServiceDesc for Game service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Game_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tictactoe.Game",
	HandlerType: (*GameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _Game_GetStatus_Handler,
		},
		{
			MethodName: "Run",
			Handler:    _Game_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/tictactoe.proto",
}
