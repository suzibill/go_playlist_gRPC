package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"go_Playlist_gRPC/internal/playlist"
	pb "go_Playlist_gRPC/internal/proto/music_player"
)

type server struct {
	pb.UnimplementedMusicPlayerServer
	Cmd chan playlist.Commands
}

func (s server) Play(ctx context.Context, e *pb.Empty) (*pb.PlaylistResponse, error) {
	fmt.Println("Play")
	s.Cmd <- playlist.Commands{Command: playlist.Play}
	return &pb.PlaylistResponse{Code: 200, Message: "Play"}, nil
}

func (s server) Pause(ctx context.Context, e *pb.Empty) (*pb.PlaylistResponse, error) {
	fmt.Println("Pause")
	s.Cmd <- playlist.Commands{Command: playlist.Pause}
	return &pb.PlaylistResponse{Code: 200, Message: "Pause"}, nil
}

func (s server) Next(ctx context.Context, e *pb.Empty) (*pb.PlaylistResponse, error) {
	fmt.Println("Next")
	s.Cmd <- playlist.Commands{Command: playlist.Next}
	return &pb.PlaylistResponse{Code: 200, Message: "Next"}, nil
}

func (s server) Prev(ctx context.Context, e *pb.Empty) (*pb.PlaylistResponse, error) {
	fmt.Println("Prev")
	s.Cmd <- playlist.Commands{Command: playlist.Prev}
	return &pb.PlaylistResponse{Code: 200, Message: "Prev"}, nil
}

func (s server) AddSong(ctx context.Context, req *pb.SongRequest) (*pb.PlaylistResponse, error) {
	fmt.Printf("AddSong: name=%s, duration=%d\n", req.GetName(), req.GetDuration())
	s.Cmd <- playlist.Commands{Command: playlist.AddSong, Song: playlist.Song{Name: req.GetName(), Duration: time.Duration(req.GetDuration())}}
	return &pb.PlaylistResponse{Code: 200, Message: "AddSong"}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	srv := server{}
	pb.RegisterMusicPlayerServer(s, &srv)
	reflection.Register(s)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	srv.Cmd = make(chan playlist.Commands)
	go playlist.DoWork(ctx, srv.Cmd)

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-stopCh
		cancel()
		s.GracefulStop()
		close(stopCh)
		fmt.Println("Server stopped")
	}()

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
