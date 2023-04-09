package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/blck-snwmn/hello-buf-grpc/gen/buf/proto/greet/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	address := "tunnel.bsnwmnbox.net"
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Connection failed.: %+v", err)
	}
	defer conn.Close()

	client := greet.NewGreeterServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	resp, err := client.Hello(ctx, &greet.HelloRequest{Name: "aaaa"})
	if err != nil {
		log.Fatalf("failed send: %+v", err)
	}
	fmt.Println(resp.Message)
}
