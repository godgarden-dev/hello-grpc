package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/kancers/hello-grpc"
)

func main() {
	addr := "localhost:50051"
	creds, err := credentials.NewClientTLSFromFile("server.crt", "")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(creds))
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	name := os.Args[1]
	ctx := context.Background()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
