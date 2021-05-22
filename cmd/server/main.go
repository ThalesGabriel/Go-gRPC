package main

import (
	"github/ThalesGabriel/go-grpc/pb/pb"
	"github/ThalesGabriel/go-grpc/services"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not initialize connection %v", err)
	}

	grpc_server := grpc.NewServer()
	pb.RegisterUserServiceServer(grpc_server, services.NewUserService())
	reflection.Register(grpc_server)

	if err := grpc_server.Serve(lis); err != nil {
		log.Fatalf("Could not initialize server %v", err)
	}

}
