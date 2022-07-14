package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc/pb"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:5051", grpc.WithInsecure()) //TLS disabled
	if err != nil {
		log.Fatal("Failed to connect to grpc server: %v", err)
	}
	defer conn.Close()

	client := pb.NewStudentServiceClient(conn)

	req := &pb.CreateStudentRequest{
		Name: "Student1",
	}

	res, err := client.CreateStudent(context.Background(), req)
	if err != nil {
		log.Fatal("Failed to connect to grpc server: %v", err)
	}

	log.Println(res)
}
