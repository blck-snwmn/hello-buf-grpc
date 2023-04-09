package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/blck-snwmn/hello-buf-grpc/gen/grpc/proto/greet/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestXxx(t *testing.T) {
	go func() {
		listenOn := "127.0.0.1:8080"
		listener, err := net.Listen("tcp", listenOn)
		if err != nil {
			fmt.Printf("failed to listen on %s: %+v", listenOn, err)
			return
		}

		server := grpc.NewServer()
		greet.RegisterGreeterServiceServer(server, &greetServer{})
		log.Println("Listening on", listenOn)
		if err := server.Serve(listener); err != nil {
			fmt.Printf("failed to serve gRPC server: %+v", err)
			return
		}
	}()
	{
		address := "127.0.0.1:8080"
		conn, err := grpc.Dial(
			address,
			// grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
		)
		if err != nil {
			log.Fatalf("Connection failed.: %+v", err)
		}
		defer conn.Close()
		// 3. gRPCクライアントを生成
		client := greet.NewGreeterServiceClient(conn)
		ctx := context.Background()
		// ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		// defer cancel()
		resp, err := client.Hello(ctx, &greet.HelloRequest{Name: "aaaa"})
		if err != nil {
			log.Fatalf("failed send: %+v", err)
		}
		fmt.Println(resp.Message)
	}
}
