package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
	"shared/model"
	"time"
)

type AlarmHistoryGetAllReq struct {
	Page               int       `json:"page"`
	Size               int       `json:"size"`
	WaterChannelDoorID int       `json:"water_channel_door_id"` // pintu
	DeviceID           int       `json:"device_id"`             // perangkat
	MinTime            time.Time `json:"min_time"`
	MaxTime            time.Time `json:"max_time"`
	Priority           string
	Metric             string
	SortOrder          string
	SortBy             string
}

type AlarmHistoryGetAllRes struct {
	DoorControlHistories []model.AlarmHistory `json:"alarm_histories"`
	Metadata             *Metadata            `json:"metadata"`
}

type AlarmHistoryGetAll = core.ActionHandler[AlarmHistoryGetAllReq, AlarmHistoryGetAllRes]

func ImplAlarmHistoryGetAll(
	getAll gateway.AlarmHistoryGetAll,
) AlarmHistoryGetAll {
	return func(ctx context.Context, req AlarmHistoryGetAllReq) (*AlarmHistoryGetAllRes, error) {

		res, err := getAll(ctx, gateway.AlarmHistoryGetAllReq{
			Page:               req.Page,
			Size:               req.Size,
			WaterChannelDoorID: req.WaterChannelDoorID,
			DeviceID:           req.DeviceID,
			MinTime:            req.MinTime,
			MaxTime:            req.MaxTime,
			Priority:           string(req.Priority),
			Metric:             string(req.Metric),
			SortOrder:          req.SortOrder,
			SortBy:             req.SortBy,
		})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &AlarmHistoryGetAllRes{
			DoorControlHistories: res.Items,
			Metadata: &Metadata{
				Pagination: Pagination{
					Page:       req.Page,
					Limit:      req.Size,
					TotalPages: totalPages,
					TotalItems: totalItems,
				},
			},
		}, nil
	}
}
