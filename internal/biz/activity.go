package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Activity struct {
	ID        string
	Name      string
	Desc      string
	BeginTime time.Time
	EndTime   time.Time
}

type ActivityRepo interface {
	GetActivity(ctx context.Context, id string) (*Activity, error)
}
type ActivityUsecase struct {
	repo ActivityRepo
}

func NewActivityUsecase(repo ActivityRepo, logger log.Logger) *ActivityUsecase {
	return &ActivityUsecase{repo: repo}
}

func (uc *ActivityUsecase) GetActivity(ctx context.Context, id string) (*Activity, error) {
	return uc.repo.GetActivity(ctx, id)
}
