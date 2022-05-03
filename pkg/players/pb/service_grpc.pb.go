// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: pkg/players/proto/service.proto

package player_pb

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

// PlayerServiceClient is the client API for PlayerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlayerServiceClient interface {
	GetPlayers(ctx context.Context, in *TeamPlayerRequest, opts ...grpc.CallOption) (*TeamPlayerResponse, error)
	UploadPlayers(ctx context.Context, in *UploadPlayersRequest, opts ...grpc.CallOption) (*UploadPlayersResponse, error)
}

type playerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlayerServiceClient(cc grpc.ClientConnInterface) PlayerServiceClient {
	return &playerServiceClient{cc}
}

func (c *playerServiceClient) GetPlayers(ctx context.Context, in *TeamPlayerRequest, opts ...grpc.CallOption) (*TeamPlayerResponse, error) {
	out := new(TeamPlayerResponse)
	err := c.cc.Invoke(ctx, "/player_pb.PlayerService/GetPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerServiceClient) UploadPlayers(ctx context.Context, in *UploadPlayersRequest, opts ...grpc.CallOption) (*UploadPlayersResponse, error) {
	out := new(UploadPlayersResponse)
	err := c.cc.Invoke(ctx, "/player_pb.PlayerService/UploadPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayerServiceServer is the server API for PlayerService service.
// All implementations must embed UnimplementedPlayerServiceServer
// for forward compatibility
type PlayerServiceServer interface {
	GetPlayers(context.Context, *TeamPlayerRequest) (*TeamPlayerResponse, error)
	UploadPlayers(context.Context, *UploadPlayersRequest) (*UploadPlayersResponse, error)
	mustEmbedUnimplementedPlayerServiceServer()
}

// UnimplementedPlayerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPlayerServiceServer struct {
}

func (UnimplementedPlayerServiceServer) GetPlayers(context.Context, *TeamPlayerRequest) (*TeamPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayers not implemented")
}
func (UnimplementedPlayerServiceServer) UploadPlayers(context.Context, *UploadPlayersRequest) (*UploadPlayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadPlayers not implemented")
}
func (UnimplementedPlayerServiceServer) mustEmbedUnimplementedPlayerServiceServer() {}

// UnsafePlayerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlayerServiceServer will
// result in compilation errors.
type UnsafePlayerServiceServer interface {
	mustEmbedUnimplementedPlayerServiceServer()
}

func RegisterPlayerServiceServer(s grpc.ServiceRegistrar, srv PlayerServiceServer) {
	s.RegisterService(&PlayerService_ServiceDesc, srv)
}

func _PlayerService_GetPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).GetPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/player_pb.PlayerService/GetPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).GetPlayers(ctx, req.(*TeamPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerService_UploadPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadPlayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).UploadPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/player_pb.PlayerService/UploadPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).UploadPlayers(ctx, req.(*UploadPlayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlayerService_ServiceDesc is the grpc.ServiceDesc for PlayerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlayerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "player_pb.PlayerService",
	HandlerType: (*PlayerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlayers",
			Handler:    _PlayerService_GetPlayers_Handler,
		},
		{
			MethodName: "UploadPlayers",
			Handler:    _PlayerService_UploadPlayers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/players/proto/service.proto",
}
