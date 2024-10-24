package service

import (
	"context"

	pb "test_kratos/api/helloworld/v1"
)

type ActivityService struct {
	pb.UnimplementedActivityServer
}

func NewActivityService() *ActivityService {
	return &ActivityService{}
}

func (s *ActivityService) GetActivity(ctx context.Context, req *pb.GetActivityRequest) (*pb.GetActivityReply, error) {
	return &pb.GetActivityReply{}, nil
}
