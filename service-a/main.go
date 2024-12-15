package main

import (
	"context"
	"log"
	"net"

	pb "github.com/khouini/microservices-hello-world/service-a/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedServiceAHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.ServiceAHelloRequest) (*pb.ServiceAHelloResponse, error) {
	log.Println("SayHello called with name: " + req.Name)
	return &pb.ServiceAHelloResponse{
		Message: "Hello from Service A: " + req.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterServiceAHelloServer(s, &server{})

	log.Println("Service A listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
