package gateway

import (
	"context"
	"fmt"
	"shared/core"
	"shared/helper"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type AlarmConfigGetAllReq struct {
	Page               int
	Size               int
	WaterChannelDoorID int
	Priority           model.AlarmConfigPriority
	Metric             model.AlarmMetric
	SortOrder          string
	SortBy             string
}

type AlarmConfigGetAllRes struct {
	Items []model.AlarmConfig
	Count int64
}

type AlarmConfigGetAll = core.ActionHandler[AlarmConfigGetAllReq, AlarmConfigGetAllRes]

func ImplAlarmConfigGetAll(db *gorm.DB) AlarmConfigGetAll {
	return func(ctx context.Context, req AlarmConfigGetAllReq) (*AlarmConfigGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.WaterChannelDoorID != 0 {
			query = query.Where("channel_id = ?", req.WaterChannelDoorID)
		}

		if req.Priority != "" {
			query = query.Where("priority = ?", req.Priority)
		}

		if req.Metric != "" {
			query = query.Where("metric = ?", req.Metric)
		}

		var count int64

		if err := query.
			Model(&model.AlarmConfig{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		// Validate sortBy
		allowedSortBy := map[string]bool{
			"channel_name": true,
			"door_name":    true,
			"priority":     true,
			"metric":       true,
		}

		// Validate and get sorting parameters
		sortBy, sortOrder, err := helper.ValidateSortParams(allowedSortBy, req.SortBy, req.SortOrder, "channel_name")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := helper.ValidatePageSize(req.Page, req.Size)

		var objs []model.AlarmConfig

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Order(orderClause).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &AlarmConfigGetAllRes{Items: objs, Count: count}, nil
	}
}
