package service

import (
	"context"
	"net/http"

	pb "github.com/vielendanke/task/proto"
)

type TaskServiceImpl struct {
	cli *http.Client
}

func NewTaskService(cli *http.Client) TaskService {
	return &TaskServiceImpl{cli: cli}
}

func (ts *TaskServiceImpl) CompanyByIIN(ctx context.Context, req *pb.CompanyByIINRequest) (*pb.CompanyByIINResponse, error) {
	return &pb.CompanyByIINResponse{Inn: req.Inn}, nil
}
