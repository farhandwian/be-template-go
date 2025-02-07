package usecase

import (
	"context"
	"dashboard/gateway"
	"fmt"
	iamgw "iam/gateway"
	iammodel "iam/model"
	"shared/constant"
	"shared/core"
	sharedGateway "shared/gateway"
	sharedModel "shared/model"
	"time"

	"github.com/google/uuid"
)

type DoorControlCreateReq struct {
	DateTime           string          `json:"date"`
	WaterChannelDoorID int             `json:"water_channel_door_id"` // pintu
	DeviceID           int             `json:"device_id"`             // perangkat
	OpenTarget         float32         `json:"open_target"`
	Reason             string          `json:"reason"`
	OfficerId          iammodel.UserID `json:"-"`
	Now                time.Time       `json:"-"`
	Tz                 *time.Location  `json:"-"`
}

type DoorControlCreateRes struct {
	ID sharedModel.DoorControlID `json:"id"`
}

type DoorControlCreate = core.ActionHandler[DoorControlCreateReq, DoorControlCreateRes]

func ImplDoorControlCreate(
	generateId gateway.GenerateId,
	save gateway.DoorControlSave,
	getOfficer iamgw.UserGetOneByID,
	callCronjob gateway.CronjobSetDoorControlGateway,
	getWaterChannelDevice gateway.GetOneWaterChannelDeviceByDoorID,
	createActivityMonitoring sharedGateway.CreateActivityMonitoringGateway,
	getWaterChannelDoorByID sharedGateway.GetWaterChannelDoorByID,
) DoorControlCreate {
	return func(ctx context.Context, req DoorControlCreateReq) (*DoorControlCreateRes, error) {

		date, err := time.ParseInLocation("2006-01-02 15:04:05", req.DateTime, req.Tz)
		if err != nil {
			return nil, fmt.Errorf("invalid format. follow '2006-01-02 15:04:05' format")
		}

		if date.Before(req.Now) {
			return nil, fmt.Errorf("date must be in future")
		}
		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		officer, err := getOfficer(ctx, iamgw.UserGetOneByIDReq{UserID: req.OfficerId})
		if err != nil {
			return nil, err
		}

		if officer == nil {
			return nil, fmt.Errorf("officer with id %v not found", req.OfficerId)
		}

		wcdObj, err := getWaterChannelDevice(ctx, gateway.GetOneWaterChannelDeviceByDoorIDReq{
			WaterChannelDoorID: req.WaterChannelDoorID,
			DeviceID:           req.DeviceID,
		})
		if err != nil {
			return nil, err
		}

		obj := sharedModel.DoorControl{
			ID:                 sharedModel.DoorControlID(genObj.RandomId),
			Date:               date,
			OpenTarget:         req.OpenTarget,
			Reason:             req.Reason,
			WaterChannelDoorID: req.WaterChannelDoorID,
			DeviceID:           req.DeviceID,
			DoorName:           wcdObj.WaterChannelDevice.Name,
			OfficerId:          req.OfficerId,
			OfficerName:        officer.User.Name,
			Status:             sharedModel.StatusMenunggu,
		}

		if _, err = save(ctx, gateway.DoorControlSaveReq{DoorControl: &obj}); err != nil {
			return nil, err
		}

		if _, err = callCronjob(ctx, gateway.CronjobSetDoorControlReq{
			TargetDate:         date,
			DoorControlID:      sharedModel.DoorControlID(genObj.RandomId),
			OpenTarget:         req.OpenTarget,
			WaterChannelDoorID: req.WaterChannelDoorID,
			DeviceID:           wcdObj.WaterChannelDevice.ExternalID,
			IPAddress:          wcdObj.WaterChannelDevice.IPAddress,
		}); err != nil {
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
				UserName:     officer.User.Name,
				Category:     constant.MONITORING_TYPE_DOOR_CONTROL,
				ActivityTime: time.Now(),
				Description:  fmt.Sprintf("%s membuat pengontrolan perangkat \"%s\" di pintu air %s ke %.2f", officer.User.Name, wcdObj.WaterChannelDevice.Name, waterChannelDoor.WaterChannelDoor.Name, req.OpenTarget),
			},
		})
		if err != nil {
			return nil, err
		}

		return &DoorControlCreateRes{ID: sharedModel.DoorControlID(genObj.RandomId)}, nil
	}
}
