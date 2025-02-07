package usecase

import (
	"context"
	"dashboard/gateway"
	"fmt"
	"github.com/google/uuid"
	iamgw "iam/gateway"
	iammodel "iam/model"
	"shared/constant"
	"shared/core"
	sharedGateway "shared/gateway"
	sharedModel "shared/model"
	"time"
)

type DoorControlRunSecurityRelayReq struct {
	Now                time.Time
	WaterChannelDoorID int
	DeviceID           int
	OfficerId          iammodel.UserID `json:"-"`
	Pin                iammodel.PIN
	Status             int
}

type DoorControlRunSecurityRelayRes struct {
}

type DoorControlRunSecurityRelay = core.ActionHandler[DoorControlRunSecurityRelayReq, DoorControlRunSecurityRelayRes]

func ImplDoorControlRunSecurityRelay(

	pinValidate iamgw.PasswordValidate,
	callApi gateway.DoorControlSecurityRelay,
	generateId gateway.GenerateId,
	saveHistory gateway.DoorControlHistorySave,
	getWaterGate gateway.GetOneWaterGatesByDoorID,
	getOfficer iamgw.UserGetOneByID,
	getWaterChannelDevice gateway.GetOneWaterChannelDeviceByDoorID,
	createActivityMonitoring sharedGateway.CreateActivityMonitoringGateway,
	getWaterChannelDoorByID sharedGateway.GetWaterChannelDoorByID,
) DoorControlRunSecurityRelay {
	return func(ctx context.Context, req DoorControlRunSecurityRelayReq) (*DoorControlRunSecurityRelayRes, error) {

		officer, err := getOfficer(ctx, iamgw.UserGetOneByIDReq{UserID: req.OfficerId})
		if err != nil {
			return nil, err
		}

		if officer == nil {
			return nil, fmt.Errorf("officer with id %v not found", req.OfficerId)
		}

		if _, err := pinValidate(ctx, iamgw.PasswordValidateReq{
			PasswordPlain:     string(req.Pin),
			PasswordEncrypted: officer.User.Pin,
		}); err != nil {
			return nil, err
		}

		wcdObj, err := getWaterChannelDevice(ctx, gateway.GetOneWaterChannelDeviceByDoorIDReq{
			WaterChannelDoorID: req.WaterChannelDoorID,
			DeviceID:           req.DeviceID,
		})
		if err != nil {
			return nil, err
		}

		if wcdObj == nil {
			return nil, fmt.Errorf("door with gate_id %v and device_id %v not found", req.WaterChannelDoorID, req.DeviceID)
		}

		if req.Status != 1 && req.Status != 0 {
			return nil, fmt.Errorf("status must 0 or 1")
		}

		if _, err := callApi(ctx, gateway.DoorControlSecurityRelayReq{
			IPAddress: wcdObj.WaterChannelDevice.IPAddress,
			Status:    req.Status,
			Now:       req.Now,
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
				Description:  fmt.Sprintf("%s mengubah security relay perangkat \"%s\" di pintu air %s menjadi %d", officer.User.Name, wcdObj.WaterChannelDevice.Name, waterChannelDoor.WaterChannelDoor.Name, req.Status),
			},
		})
		if err != nil {
			return nil, err
		}

		return &DoorControlRunSecurityRelayRes{}, nil
	}
}
