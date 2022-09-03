package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/blck-snwmn/hello-buf-grpc/gen/grpc/proto/greet/v1"
	"google.golang.org/grpc"
)

type greetServer struct {
	greet.UnimplementedGreeterServiceServer
}

func (s *greetServer) Hello(ctx context.Context, in *greet.HelloRequest) (*greet.HelloResponse, error) {
	if len(in.Name) < 2 {
		return nil, errors.New("text string")
	}
	return &greet.HelloResponse{
		Message: fmt.Sprintf("hello %s", in.Name),
	}, nil
}

func main() {
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
}
