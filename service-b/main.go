package main

import (
	"context"
	"log"
	"net"

	pb "github.com/khouini/microservices-hello-world/service-b/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedServiceBHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.ServiceBHelloRequest) (*pb.ServiceBHelloResponse, error) {
	log.Println("SayHello called with name: " + req.Name)
	return &pb.ServiceBHelloResponse{
		Message: "Hello from Service B: " + req.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterServiceBHelloServer(s, &server{})

	log.Println("Service B listening on :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
