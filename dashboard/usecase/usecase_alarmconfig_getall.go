package usecase

import (
	"context"
	"shared/core"
	sg "shared/gateway"
	"shared/model"
	"shared/usecase"
)

type AlarmConfigGetAllReq struct {
	Page               int `json:"page"`
	Size               int `json:"size"`
	WaterChannelDoorID int
	Priority           model.AlarmConfigPriority
	Metric             model.AlarmMetric
	SortOrder          string
	SortBy             string
}

type AlarmConfigGetAllRes struct {
	AlarmConfigs []model.AlarmConfig `json:"alarm_configs"`
	Metadata     *usecase.Metadata   `json:"metadata"`
}

type AlarmConfigGetAll = core.ActionHandler[AlarmConfigGetAllReq, AlarmConfigGetAllRes]

func ImplAlarmConfigGetAll(
	getAll sg.AlarmConfigGetAll,
) AlarmConfigGetAll {
	return func(ctx context.Context, req AlarmConfigGetAllReq) (*AlarmConfigGetAllRes, error) {

		res, err := getAll(ctx, sg.AlarmConfigGetAllReq{
			Page: req.Page, Size: req.Size,
			WaterChannelDoorID: req.WaterChannelDoorID,
			Priority:           req.Priority,
			Metric:             req.Metric,
			SortOrder:          req.SortOrder,
			SortBy:             req.SortBy,
		})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &AlarmConfigGetAllRes{
			AlarmConfigs: res.Items,
			Metadata: &usecase.Metadata{
				Pagination: usecase.Pagination{
					Page:       req.Page,
					Limit:      req.Size,
					TotalPages: totalPages,
					TotalItems: totalItems,
				},
			},
		}, nil
	}
}
