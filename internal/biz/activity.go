package biz

import (
	"context"
	"test_kratos/utils"

	"github.com/go-kratos/kratos/v2/log"
)

type Activity struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	BeginTime string `json:"begin_time"`
	EndTime   string `json:"end_time"`
}

type Activitys struct {
	Items   []*Activity `json:"items"`
	Version string      `json:"version"`
}

var (
	activityDataKey = "activity"
)

type ActivityRepo interface {
	GetActivity(ctx context.Context, id string) (*Activity, error)
}
type ActivityUsecase struct {
	appConfig GlobalAppConfigInterface
}

func NewActivityUsecase(appConfig GlobalAppConfigInterface, logger log.Logger) *ActivityUsecase {
	return &ActivityUsecase{appConfig: appConfig}
}

func (uc *ActivityUsecase) GetActivity(ctx context.Context, id string) (*Activity, error) {
	allActivitiesMap, err := uc.appConfig.GetConfig(activityDataKey)
	if err != nil {
		return nil, err
	}
	activitysStruct := Activitys{}
	err = utils.MapToStruct(allActivitiesMap, &activitysStruct)
	if err != nil {
		return nil, err
	}
	for _, activity := range activitysStruct.Items {
		if activity.ID == id {
			return activity, nil
		}
	}
	return nil, nil
}
