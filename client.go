package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "protoc/protoc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// サーバーへの接続を確立
	time.Sleep(1 * time.Second) // サーバーが起動するまで待機
	conn, err := grpc.NewClient(
		"localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// クライアントを作成
	client := pb.NewStringServiceClient(conn)

	// コンテキストを作成
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// メッセージを作成
	msg := &pb.StringMessage{
		Content: "Hello from client!",
		Id:      42,
	}

	fmt.Printf("Client sent - ID: %d, Content: %s\n", msg.Id, msg.Content)

	// サーバーにメッセージを送信して応答を受け取る
	response, err := client.Echo(ctx, msg)
	if err != nil {
		log.Fatalf("Failed to call Echo: %v", err)
	}

	fmt.Printf("Client received - ID: %d, Content: %s\n", response.Id, response.Content)
}
