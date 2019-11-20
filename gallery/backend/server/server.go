package main

import (
	"context"
	"github.com/kancers/hello-grpc/backend"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *backend.HelloRequest) (*backend.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &backend.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	addr := ":50051"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	cred, err := credentials.NewServerTLSFromFile("server.crt", "private.key")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer(grpc.Creds(cred))
	backend.RegisterGreeterServer(s, &server{})
	log.Printf("gRPC server listening on " + addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
