package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/SAP-samples/cf-http2/example"
)

type server struct {
	pb.ExampleServer
}

func (s *server) Run(c context.Context, r *pb.Request) (*pb.Response, error) {
	return &pb.Response{Message: "Hello! This Go application is speaking gRPC"}, nil
}

func main() {
	port := os.Getenv("PORT")
	address := fmt.Sprintf("0.0.0.0:%s", port)
	fmt.Printf("Listening [%s]...\n", address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error TCP listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterExampleServer(s, &server{})

	// Register reflection service on gRPC server to make debugging easier
	reflection.Register(s)

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("Error server serve: %v", err)
	}
}
