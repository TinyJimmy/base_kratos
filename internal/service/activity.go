package service

import (
	"context"

	pb "test_kratos/api/helloworld/v1"
	"test_kratos/internal/biz"
)

type ActivityService struct {
	pb.UnimplementedActivityServer
	uc *biz.ActivityUsecase
}

func NewActivityService(uc *biz.ActivityUsecase) *ActivityService {
	return &ActivityService{uc: uc}
}

func (s *ActivityService) GetActivity(ctx context.Context, req *pb.GetActivityRequest) (*pb.GetActivityReply, error) {
	return &pb.GetActivityReply{}, nil
}
