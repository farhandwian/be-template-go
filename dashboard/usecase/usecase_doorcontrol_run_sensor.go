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

type DoorControlRunSensorReq struct {
	Now                time.Time
	WaterChannelDoorID int
	DeviceID           int
	OfficerId          iammodel.UserID `json:"-"`
	Pin                iammodel.PIN
	Status             int
	// OpenTarget         float32
	// Reason             string
}

type DoorControlRunSensorRes struct {
}

type DoorControlRunSensor = core.ActionHandler[DoorControlRunSensorReq, DoorControlRunSensorRes]

func ImplDoorControlRunSensor(

	pinValidate iamgw.PasswordValidate,
	callApi gateway.DoorControlSensor,
	generateId gateway.GenerateId,
	saveHistory gateway.DoorControlHistorySave,
	getWaterGate gateway.GetOneWaterGatesByDoorID,
	getOfficer iamgw.UserGetOneByID,
	getWaterChannelDevice gateway.GetOneWaterChannelDeviceByDoorID,
	createActivityMonitoring sharedGateway.CreateActivityMonitoringGateway,
	getWaterChannelDoorByID sharedGateway.GetWaterChannelDoorByID,

) DoorControlRunSensor {
	return func(ctx context.Context, req DoorControlRunSensorReq) (*DoorControlRunSensorRes, error) {

		officer, err := getOfficer(ctx, iamgw.UserGetOneByIDReq{UserID: req.OfficerId})
		if err != nil {
			return nil, err
		}

		if officer == nil {
			return nil, fmt.Errorf("officer with id %v not found", req.OfficerId)
		}

		// check pin
		if _, err := pinValidate(ctx, iamgw.PasswordValidateReq{
			PasswordPlain:     string(req.Pin),
			PasswordEncrypted: officer.User.Pin,
		}); err != nil {
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

		if wcdObj == nil {
			return nil, fmt.Errorf("door with gate_id %v and device_id %v not found", req.WaterChannelDoorID, req.DeviceID)
		}

		if req.Status != 1 && req.Status != 0 {
			return nil, fmt.Errorf("status must 0 or 1")
		}

		if _, err := callApi(ctx, gateway.DoorControlSensorReq{
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
				Description:  fmt.Sprintf("%s mengubah sensor perangkat \"%s\" di pintu air %s menjadi %d", officer.User.Name, wcdObj.WaterChannelDevice.Name, waterChannelDoor.WaterChannelDoor.Name, req.Status),
			},
		})
		if err != nil {
			return nil, err
		}

		return &DoorControlRunSensorRes{}, nil
	}
}
