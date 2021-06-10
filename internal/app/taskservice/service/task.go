package service

import (
	"context"

	pb "github.com/vielendanke/task/proto"
)

type TaskService interface {
	CompanyByIIN(context.Context, *pb.CompanyByIINRequest) (*pb.CompanyByIINResponse, error)
}
