package server

import (
	"context"
	"github.com/google/uuid"
	"grpc/pb"
)

type Student struct {
	Id   string `json:"id"`
	Name string `json:"name,omitempty"`
}

func (s *Student) CreateStudent(ctx context.Context, in *pb.CreateStudentRequest) (*pb.CreateStudentResponse, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	s.Id = id.String()
	s.Name = in.Name

	//storage the student created

	return &pb.CreateStudentResponse{
		Id: s.Id,
		Name: s.Name,
	}, nil
}
