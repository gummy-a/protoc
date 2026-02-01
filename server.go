package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "protoc/protoc/pb"

	"google.golang.org/grpc"
)

// StringServiceServerの実装
type server struct {
	pb.UnimplementedStringServiceServer
}

// Echo メソッドの実装
func (s *server) Echo(ctx context.Context, msg *pb.StringMessage) (*pb.StringMessage, error) {
	fmt.Printf("Server received - ID: %d, Content: %s\n", msg.Id, msg.Content)

	// レスポンスを作成
	response := &pb.StringMessage{
		Content: "Echo: " + msg.Content,
		Id:      msg.Id + 1,
	}

	fmt.Printf("Server sent - ID: %d, Content: %s\n", response.Id, response.Content)
	return response, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Server listening on :8080")

	// gRPC サーバーを作成
	s := grpc.NewServer()
	pb.RegisterStringServiceServer(s, &server{})

	// サーバーを開始
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
