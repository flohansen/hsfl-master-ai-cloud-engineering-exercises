package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/flohansen/grpc-go-example/internal/proto"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	name := request.GetName()
	response := &proto.HelloResponse{
		Greeting: fmt.Sprintf("Hello %s!", name),
	}
	return response, nil
}

func main() {
	listener, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalf("could not listen: %v", err)
	}

	srv := grpc.NewServer()
	proto.RegisterGreeterServer(srv, &server{})

	log.Println("listening on port 3001")
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("could not serve: %v", err)
	}
}
