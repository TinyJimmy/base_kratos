package service

import (
	"context"

	pb "test_kratos/api/helloworld/v1"
	"test_kratos/internal/biz"
	"test_kratos/internal/zlog"

	"github.com/go-kratos/kratos/v2/errors"
	"go.uber.org/zap"
)

type ActivityService struct {
	pb.UnimplementedActivityServer
	uc *biz.ActivityUsecase
}

func NewActivityService(uc *biz.ActivityUsecase) *ActivityService {
	return &ActivityService{uc: uc}
}

func (s *ActivityService) GetActivity(ctx context.Context, req *pb.GetActivityRequest) (*pb.GetActivityReply, error) {
	zlog.RouterLogger.Info("request", zap.String("method", "GetActivity"), zap.String("activity_id", req.ActivityId), zap.String("url", "/api/v1/activity/get"))
	activity, err := s.uc.GetActivity(ctx, req.ActivityId)
	if err != nil {
		return nil, err
	}
	if activity == nil {
		return nil, errors.NotFound("activity", "activity not found")
	}
	return &pb.GetActivityReply{Body: &pb.ActivityDetail{
		ActivityId:        activity.ID,
		ActivityName:      activity.Name,
		ActivityDesc:      activity.Desc,
		ActivityBeginTime: activity.BeginTime,
		ActivityEndTime:   activity.EndTime,
	}}, nil
}
