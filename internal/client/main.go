package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

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
	//_, err = c.Play(context.Background(), &pb.Empty{})
	//if err != nil {
	//	log.Fatalf("Failed to play: %v", err)
	//}
	// вызываем метод AddSong
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//defer cancel()
	ans, err := c.AddSong(context.Background(), &pb.SongRequest{
		Name:     "Song Name One",
		Duration: int64(3 * time.Second), // в секундах
	})
	if err != nil {
		log.Fatalf("Failed to add song: %v", err)
	}
	fmt.Println(ans)
	ans, err = c.AddSong(context.Background(), &pb.SongRequest{
		Name:     "Another Song Name",
		Duration: int64(3 * time.Second), // в секундах
	})
	if err != nil {
		log.Fatalf("Failed to add song: %v", err)
	}
	fmt.Println(ans)

	ans, err = c.AddSong(context.Background(), &pb.SongRequest{
		Name:     "Simple Song Name",
		Duration: int64(3 * time.Second), // в секундах
	})
	if err != nil {
		log.Fatalf("Failed to add song: %v", err)
	}
	fmt.Println(ans)

	// вызываем метод Play
	ans, err = c.Play(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Failed to play: %v", err)
	}
	fmt.Println(ans)
	time.Sleep(1 * time.Second)
	ans, err = c.Next(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Failed to play: %v", err)
	}
	fmt.Println(ans)
	time.Sleep(1 * time.Second)
	ans, err = c.Prev(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Failed to play: %v", err)
	}
	fmt.Println(ans)
	time.Sleep(1 * time.Second)
	ans, err = c.Pause(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Failed to play: %v", err)
	}
	fmt.Println(ans)
	time.Sleep(1 * time.Second)
	ans, err = c.Play(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Failed to play: %v", err)
	}
	fmt.Println(ans)
}
