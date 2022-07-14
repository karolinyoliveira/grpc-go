package main

import (
	"grpc/pb"
	"grpc/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterStudentServiceServer(grpcServer, &server.Student{})
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":5051")
	if err != nil {
		log.Fatal("Failed to start server: %v", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Failed to start listen: %v", err)
	}
}
