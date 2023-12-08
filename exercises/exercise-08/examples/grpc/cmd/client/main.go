package main

import (
	"context"
	"fmt"
	"log"

	"github.com/flohansen/grpc-go-example/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewGreeterClient(conn)
	request := &proto.HelloRequest{Name: "Student"}

	res, err := client.SayHello(context.Background(), request)
	if err != nil {
		log.Fatalf("could not call rpc: %v", err)
	}

	fmt.Println(res.GetGreeting())
}
