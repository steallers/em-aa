package handler

import (
	"context"
	"log"

	pb "github.com/steallers/employee-management/connectors/employee.pb.go"
)

func (h *Handler) AddNewEmployee(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}
