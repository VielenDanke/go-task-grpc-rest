package handler

import (
	"context"

	"github.com/vielendanke/task/internal/app/taskservice/service"
	pb "github.com/vielendanke/task/proto"
)

type TaskHandler struct {
	ts service.TaskService
}

func NewTaskHandler(ts service.TaskService) *TaskHandler {
	return &TaskHandler{ts: ts}
}

func (th *TaskHandler) CompanyByIIN(ctx context.Context, req *pb.CompanyByIINRequest) (*pb.CompanyByIINResponse, error) {
	return th.ts.CompanyByIIN(ctx, req)
}
