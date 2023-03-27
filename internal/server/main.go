package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "go_Playlist_gRPC/internal/proto/music_player"
)

type server struct {
	pb.UnimplementedMusicPlayerServer
}

func (s *server) Play(ctx context.Context, e *pb.Empty) (*pb.Empty, error) {
	fmt.Println("Play")
	return &pb.Empty{}, nil
}

func (s *server) Pause(ctx context.Context, e *pb.Empty) (*pb.Empty, error) {
	fmt.Println("Pause")
	return &pb.Empty{}, nil
}

func (s *server) Next(ctx context.Context, e *pb.Empty) (*pb.Empty, error) {
	fmt.Println("Next")
	return &pb.Empty{}, nil
}

func (s *server) Prev(ctx context.Context, e *pb.Empty) (*pb.Empty, error) {
	fmt.Println("Prev")
	return &pb.Empty{}, nil
}

func (s *server) AddSong(ctx context.Context, req *pb.SongRequest) (*pb.Empty, error) {
	fmt.Printf("AddSong: name=%s, duration=%d\n", req.GetName(), req.GetDuration())
	return &pb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterMusicPlayerServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
