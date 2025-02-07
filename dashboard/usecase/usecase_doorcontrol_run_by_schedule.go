package usecase

import (
	"context"
	"dashboard/gateway"
	"fmt"
	"shared/constant"
	"shared/core"
	sharedGateway "shared/gateway"
	sharedModel "shared/model"
	"time"

	"github.com/google/uuid"
)

type DoorControlRunScheduledReq struct {
	Now                time.Time
	DoorControlID      sharedModel.DoorControlID
	WaterChannelDoorID int
	DeviceID           int
	OpenTarget         float32
	IPAddress          string
}

type DoorControlRunScheduledRes struct {
}

type DoorControlRunScheduled = core.ActionHandler[DoorControlRunScheduledReq, DoorControlRunScheduledRes]

func ImplDoorControlRunScheduled(

	getOne gateway.DoorControlGetOne,
	callApi gateway.DoorControlAPI,
	saveDoorControl gateway.DoorControlSave,
	generateId gateway.GenerateId,
	saveHistory gateway.DoorControlHistorySave,
	getOneWaterGatesByDoorID gateway.GetOneWaterGatesByDoorID,
	getWaterChannelDevice gateway.GetOneWaterChannelDeviceByDoorID,
	createActivityMonitoring sharedGateway.CreateActivityMonitoringGateway,
	getWaterChannelDoorByID sharedGateway.GetWaterChannelDoorByID,
) DoorControlRunScheduled {
	return func(ctx context.Context, req DoorControlRunScheduledReq) (*DoorControlRunScheduledRes, error) {

		obj, err := getOne(ctx, gateway.DoorControlGetOneReq{ID: req.DoorControlID})
		if err != nil {
			return nil, err
		}

		waterGatesRes, err := getOneWaterGatesByDoorID(ctx, gateway.GetOneWaterGatesByDoorIDReq{
			WaterChannelDoorID: req.WaterChannelDoorID,
			DeviceID:           req.DeviceID,
		})
		if err != nil {
			return nil, err
		}

		// ambil ip_address dan perangkat_id
		wcdObj, err := getWaterChannelDevice(ctx, gateway.GetOneWaterChannelDeviceByDoorIDReq{
			WaterChannelDoorID: req.WaterChannelDoorID,
			DeviceID:           req.DeviceID,
		})
		if err != nil {
			return nil, err
		}

		if !waterGatesRes.WaterGates.SecurityRelay {
			return nil, fmt.Errorf("security relay tidak aktif")
		}

		instruction := "up"
		if waterGatesRes.WaterGates.GateLevel > req.OpenTarget {
			instruction = "down"
		}

		if _, err = callApi(ctx, gateway.DoorControlAPIReq{
			Value:       req.OpenTarget,
			PerangkatID: req.DeviceID,
			Instruction: instruction,
			GateID:      wcdObj.WaterChannelDevice.GroupRelay,
			IPAddress:   wcdObj.WaterChannelDevice.IPAddress,
			Position:    wcdObj.WaterChannelDevice.MaxHeightSensor - req.OpenTarget,
			FullTime:    wcdObj.WaterChannelDevice.FullTime,
			MaxHeight:   100,
			Limit:       true,
			Now:         req.Now,
		}); err != nil {
			obj.Item.Status = sharedModel.StatusGagal
			obj.Item.ErrorMessage = err.Error()
		} else {
			obj.Item.Status = sharedModel.StatusDieksekusi
		}

		if _, err = saveDoorControl(ctx, gateway.DoorControlSaveReq{DoorControl: obj.Item}); err != nil {
			return nil, err
		}

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		dch := sharedModel.DoorControlHistory{
			ID:                 sharedModel.DoorControlHistoryID(genObj.RandomId),
			Date:               req.Now,
			WaterChannelDoorID: req.WaterChannelDoorID,
			DeviceID:           req.DeviceID,
			DoorName:           obj.Item.DoorName,
			OpenTarget:         obj.Item.OpenTarget,
			OpenCurrent:        waterGatesRes.WaterGates.GateLevel,
			Type:               sharedModel.TypeTerjadwal,
			Reason:             obj.Item.Reason,
			Status:             obj.Item.Status,
			OfficerId:          obj.Item.OfficerId,
			OfficerName:        obj.Item.OfficerName,
			ErrorMessage:       obj.Item.ErrorMessage,
		}

		if _, err = saveHistory(ctx, gateway.DoorControlHistorySaveReq{DoorControlHistory: &dch}); err != nil {
			return nil, err
		}

		waterChannelDoor, err := getWaterChannelDoorByID(ctx, sharedGateway.GetWaterChannelDoorByIDReq{WaterChannelDoorID: req.WaterChannelDoorID})
		if err != nil {
			return nil, err
		}

		//store logging
		_, err = createActivityMonitoring(ctx, sharedGateway.CreateActivityMonitoringReq{
			ActivityMonitor: sharedModel.ActivityMonitor{
				ID:           uuid.NewString(),
				UserName:     obj.Item.OfficerName,
				Category:     constant.MONITORING_TYPE_DOOR_CONTROL,
				ActivityTime: time.Now(),
				Description:  fmt.Sprintf("%s membuat pengontrolan terjadwal untuk perangkat \"%s\" di pintu air %s ke %.2f", obj.Item.OfficerName, wcdObj.WaterChannelDevice.Name, waterChannelDoor.WaterChannelDoor.Name, req.OpenTarget),
			},
		})
		if err != nil {
			return nil, err
		}

		return &DoorControlRunScheduledRes{}, nil
	}
}
