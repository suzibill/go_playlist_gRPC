// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: music_player.proto

package music_player

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

// MusicPlayerClient is the client API for MusicPlayer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MusicPlayerClient interface {
	Play(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlaylistResponse, error)
	Pause(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlaylistResponse, error)
	Next(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlaylistResponse, error)
	Prev(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlaylistResponse, error)
	AddSong(ctx context.Context, in *SongRequest, opts ...grpc.CallOption) (*PlaylistResponse, error)
}

type musicPlayerClient struct {
	cc grpc.ClientConnInterface
}

func NewMusicPlayerClient(cc grpc.ClientConnInterface) MusicPlayerClient {
	return &musicPlayerClient{cc}
}

func (c *musicPlayerClient) Play(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlaylistResponse, error) {
	out := new(PlaylistResponse)
	err := c.cc.Invoke(ctx, "/music_player.MusicPlayer/play", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicPlayerClient) Pause(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlaylistResponse, error) {
	out := new(PlaylistResponse)
	err := c.cc.Invoke(ctx, "/music_player.MusicPlayer/pause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicPlayerClient) Next(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlaylistResponse, error) {
	out := new(PlaylistResponse)
	err := c.cc.Invoke(ctx, "/music_player.MusicPlayer/next", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicPlayerClient) Prev(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlaylistResponse, error) {
	out := new(PlaylistResponse)
	err := c.cc.Invoke(ctx, "/music_player.MusicPlayer/prev", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicPlayerClient) AddSong(ctx context.Context, in *SongRequest, opts ...grpc.CallOption) (*PlaylistResponse, error) {
	out := new(PlaylistResponse)
	err := c.cc.Invoke(ctx, "/music_player.MusicPlayer/addSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MusicPlayerServer is the server API for MusicPlayer service.
// All implementations must embed UnimplementedMusicPlayerServer
// for forward compatibility
type MusicPlayerServer interface {
	Play(context.Context, *Empty) (*PlaylistResponse, error)
	Pause(context.Context, *Empty) (*PlaylistResponse, error)
	Next(context.Context, *Empty) (*PlaylistResponse, error)
	Prev(context.Context, *Empty) (*PlaylistResponse, error)
	AddSong(context.Context, *SongRequest) (*PlaylistResponse, error)
	mustEmbedUnimplementedMusicPlayerServer()
}

// UnimplementedMusicPlayerServer must be embedded to have forward compatible implementations.
type UnimplementedMusicPlayerServer struct {
}

func (UnimplementedMusicPlayerServer) Play(context.Context, *Empty) (*PlaylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Play not implemented")
}
func (UnimplementedMusicPlayerServer) Pause(context.Context, *Empty) (*PlaylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedMusicPlayerServer) Next(context.Context, *Empty) (*PlaylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Next not implemented")
}
func (UnimplementedMusicPlayerServer) Prev(context.Context, *Empty) (*PlaylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prev not implemented")
}
func (UnimplementedMusicPlayerServer) AddSong(context.Context, *SongRequest) (*PlaylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSong not implemented")
}
func (UnimplementedMusicPlayerServer) mustEmbedUnimplementedMusicPlayerServer() {}

// UnsafeMusicPlayerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MusicPlayerServer will
// result in compilation errors.
type UnsafeMusicPlayerServer interface {
	mustEmbedUnimplementedMusicPlayerServer()
}

func RegisterMusicPlayerServer(s grpc.ServiceRegistrar, srv MusicPlayerServer) {
	s.RegisterService(&MusicPlayer_ServiceDesc, srv)
}

func _MusicPlayer_Play_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicPlayerServer).Play(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_player.MusicPlayer/play",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicPlayerServer).Play(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicPlayer_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicPlayerServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_player.MusicPlayer/pause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicPlayerServer).Pause(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicPlayer_Next_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicPlayerServer).Next(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_player.MusicPlayer/next",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicPlayerServer).Next(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicPlayer_Prev_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicPlayerServer).Prev(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_player.MusicPlayer/prev",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicPlayerServer).Prev(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicPlayer_AddSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicPlayerServer).AddSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_player.MusicPlayer/addSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicPlayerServer).AddSong(ctx, req.(*SongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MusicPlayer_ServiceDesc is the grpc.ServiceDesc for MusicPlayer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MusicPlayer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "music_player.MusicPlayer",
	HandlerType: (*MusicPlayerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "play",
			Handler:    _MusicPlayer_Play_Handler,
		},
		{
			MethodName: "pause",
			Handler:    _MusicPlayer_Pause_Handler,
		},
		{
			MethodName: "next",
			Handler:    _MusicPlayer_Next_Handler,
		},
		{
			MethodName: "prev",
			Handler:    _MusicPlayer_Prev_Handler,
		},
		{
			MethodName: "addSong",
			Handler:    _MusicPlayer_AddSong_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "music_player.proto",
}
