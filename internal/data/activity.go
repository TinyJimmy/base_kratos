package data

import (
	"context"
	"test_kratos/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type activityRepo struct {
	data *Data
	log  *log.Helper
}

func NewActivityRepo(data *Data, logger log.Logger) biz.ActivityRepo {
	return &activityRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ar *activityRepo) GetActivity(ctx context.Context, id string) (*biz.Activity, error) {
	return nil, nil
}
