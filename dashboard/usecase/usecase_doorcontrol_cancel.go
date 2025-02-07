package usecase

import (
	"context"
	"dashboard/gateway"
	"shared/core"
	"shared/model"
	"time"
)

type DoorControlCancelReq struct {
	ID  model.DoorControlID `json:"id"`
	Now time.Time
}

type DoorControlCancelRes struct{}

type DoorControlCancel = core.ActionHandler[DoorControlCancelReq, DoorControlCancelRes]

func ImplDoorControlCancel(
	getOne gateway.DoorControlGetOne,
	save gateway.DoorControlSave,
	cronjobCancel gateway.CronjobCancelDoorControlGateway,
	generateId gateway.GenerateId,
	getOneWaterGatesByDoorID gateway.GetOneWaterGatesByDoorID,
	saveHistory gateway.DoorControlHistorySave,
) DoorControlCancel {
	return func(ctx context.Context, req DoorControlCancelReq) (*DoorControlCancelRes, error) {

		dc, err := getOne(ctx, gateway.DoorControlGetOneReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		dc.Item.Status = model.StatusDibatalkan

		if _, err = save(ctx, gateway.DoorControlSaveReq{DoorControl: dc.Item}); err != nil {
			return nil, err
		}

		if _, err := cronjobCancel(ctx, gateway.CronjobCancelDoorControlReq{
			DoorControlID: req.ID,
		}); err != nil {
			return nil, err
		}

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		waterGatesRes, err := getOneWaterGatesByDoorID(ctx, gateway.GetOneWaterGatesByDoorIDReq{
			WaterChannelDoorID: dc.Item.WaterChannelDoorID,
			DeviceID:           dc.Item.DeviceID,
		})
		if err != nil {
			return nil, err
		}

		dch := model.DoorControlHistory{
			ID:                 model.DoorControlHistoryID(genObj.RandomId),
			Date:               req.Now,
			WaterChannelDoorID: dc.Item.WaterChannelDoorID,
			DeviceID:           dc.Item.DeviceID,
			DoorName:           dc.Item.DoorName,
			OpenTarget:         dc.Item.OpenTarget,
			OpenCurrent:        waterGatesRes.WaterGates.GateLevel,
			Type:               model.TypeTerjadwal,
			Reason:             dc.Item.Reason,
			Status:             dc.Item.Status,
			OfficerId:          dc.Item.OfficerId,
			OfficerName:        dc.Item.OfficerName,
		}

		if _, err = saveHistory(ctx, gateway.DoorControlHistorySaveReq{DoorControlHistory: &dch}); err != nil {
			return nil, err
		}

		return &DoorControlCancelRes{}, nil
	}
}
