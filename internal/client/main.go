package main

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	pb "go_Playlist_gRPC/internal/proto/music_player"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewMusicPlayerClient(conn)

	// вызываем метод Play
	_, err = c.Play(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Failed to play: %v", err)
	}

	// вызываем метод AddSong
	_, err = c.AddSong(context.Background(), &pb.SongRequest{
		Name:     "Song Name",
		Duration: 180, // в секундах
	})
	if err != nil {
		log.Fatalf("Failed to add song: %v", err)
	}
}
