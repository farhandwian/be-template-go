package gateway

import (
	"context"
	"fmt"
	"shared/core"
	"shared/helper"
	"shared/middleware"
	"shared/model"
	"time"

	"gorm.io/gorm"
)

type AlarmHistoryGetAllReq struct {
	Page               int
	Size               int
	WaterChannelDoorID int
	DeviceID           int
	MinTime            time.Time
	MaxTime            time.Time
	Priority           string
	Metric             string
	SortOrder          string
	SortBy             string
}

type AlarmHistoryGetAllRes struct {
	Items []model.AlarmHistory
	Count int64
}

type AlarmHistoryGetAll = core.ActionHandler[AlarmHistoryGetAllReq, AlarmHistoryGetAllRes]

func ImplAlarmHistoryGetAll(db *gorm.DB) AlarmHistoryGetAll {
	return func(ctx context.Context, req AlarmHistoryGetAllReq) (*AlarmHistoryGetAllRes, error) {

		query := middleware.
			GetDBFromContext(ctx, db).
			Model(&model.AlarmHistory{})

		var count int64

		if req.WaterChannelDoorID != 0 {
			query = query.Where("channel_id = ?", req.WaterChannelDoorID)
		}

		if req.DeviceID != 0 {
			query = query.Where("door_id = ?", req.DeviceID)
		}

		if !req.MaxTime.IsZero() {
			query = query.Where("created_at <= ?", req.MaxTime.Format("2006-01-02 15:04:05.999"))
		}

		if !req.MinTime.IsZero() {
			query = query.Where("created_at >= ?", req.MinTime.Format("2006-01-02 15:04:05.999"))
		}

		if req.Metric != "" {
			query = query.Where("metric = ?", req.Metric)
		}

		if req.Priority != "" {
			query = query.Where("priority = ?", req.Priority)
		}

		var objs []model.AlarmHistory

		if err := query.
			Count(&count).
			Error; err != nil {

			if err == gorm.ErrRecordNotFound {
				return &AlarmHistoryGetAllRes{
					Items: objs,
					Count: count,
				}, nil
			}

			return nil, core.NewInternalServerError(err)
		}

		// Validate sortBy
		allowedSortBy := map[string]bool{
			"channel_name": true,
			"door_name":    true,
			"priority":     true,
			"metric":       true,
			"created_at":   true,
		}

		// Validate and get sorting parameters
		sortBy, sortOrder, err := helper.ValidateSortParams(allowedSortBy, req.SortBy, req.SortOrder, "created_at")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := helper.ValidatePageSize(req.Page, req.Size)

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Order(orderClause).
			Find(&objs).
			Error; err != nil {

			if err == gorm.ErrRecordNotFound {
				return &AlarmHistoryGetAllRes{
					Items: objs,
					Count: count,
				}, nil
			}

			return nil, core.NewInternalServerError(err)
		}

		return &AlarmHistoryGetAllRes{
			Items: objs,
			Count: count,
		}, nil
	}
}
