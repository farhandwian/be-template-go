package usecase

import (
	"context"
	"dashboard/gateway"
	"dashboard/model"
	"errors"
	"shared/core"
	sharedGateway "shared/gateway"
	"time"
)

type CctvImageProcessingReq struct {
	WaterChannelDoorId int    `json:"water_channel_door_id"`
	IP                 string `json:"ip"`
	DetectedObject     string `json:"detected_object"`
	ImagePath          string `json:"image_path"`
}

type CctvImageProcessingRes struct {
	Detection *model.CctvImageProcessing `json:"detection"`
}

type CctvImageProcessing = core.ActionHandler[CctvImageProcessingReq, CctvImageProcessingRes]

func ImplCctvImageProcessing(
	generateId gateway.GenerateId,
	save gateway.CctvImageProcessingSave,
	waterChannelDoor sharedGateway.GetWaterChannelDoorByID,
) CctvImageProcessing {
	return func(ctx context.Context, req CctvImageProcessingReq) (*CctvImageProcessingRes, error) {
		if err := req.Validate(); err != nil {
			return nil, err
		}

		obj := model.CctvImageProcessing{
			Timestamp:          time.Now(),
			WaterChannelDoorId: req.WaterChannelDoorId,
			IP:                 req.IP,
			DetectedObject:     req.DetectedObject,
			ImagePath:          req.ImagePath,
		}

		if _, err := save(ctx, gateway.CctvImageProcessingSaveReq{CctvImageProcessing: obj}); err != nil {
			return nil, err
		}
		return &CctvImageProcessingRes{Detection: &obj}, nil
	}
}

func (r CctvImageProcessingReq) Validate() error {
	if r.WaterChannelDoorId == 0 {
		return errors.New("WaterChannelDoorId is required")
	}

	if r.IP == "" {
		return errors.New("IP is required")
	}

	return nil
}
