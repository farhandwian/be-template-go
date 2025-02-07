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

type DoorControlRunDirectlyReq struct {
	Now                time.Time
	WaterChannelDoorID int
	DeviceID           int
	OpenTarget         float32
	Reason             string
	OfficerId          iammodel.UserID `json:"-"`
	Pin                iammodel.PIN
}

type DoorControlRunDirectlyRes struct {
}

type DoorControlRunDirectly = core.ActionHandler[DoorControlRunDirectlyReq, DoorControlRunDirectlyRes]

func ImplDoorControlRunDirectly(

	pinValidate iamgw.PasswordValidate,
	callApi gateway.DoorControlAPI,
	generateId gateway.GenerateId,
	saveHistory gateway.DoorControlHistorySave,
	getWaterGate gateway.GetOneWaterGatesByDoorID,
	getOfficer iamgw.UserGetOneByID,
	getWaterChannelDevice gateway.GetOneWaterChannelDeviceByDoorID,
	createActivityMonitoring sharedGateway.CreateActivityMonitoringGateway,
	getWaterChannelDoorByID sharedGateway.GetWaterChannelDoorByID,
) DoorControlRunDirectly {
	return func(ctx context.Context, req DoorControlRunDirectlyReq) (*DoorControlRunDirectlyRes, error) {

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

		var status sharedModel.DoorControlStatus = sharedModel.StatusDieksekusi // positif thinking dong
		var errorMessage string = ""

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

		waterGatesRes, err := getWaterGate(ctx, gateway.GetOneWaterGatesByDoorIDReq{
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

		// {
		//   "perangkat_id": 2686,
		//   "ip_address": "10.0.17.226",
		//   "gate": 1,
		//   "instruction": "up",
		//   "position": 43,
		//   "value": 12,
		//   "full_time": 60000,
		//   "max_heigh": 100,
		//   "limit": true
		// }

		if _, err := callApi(ctx, gateway.DoorControlAPIReq{
			PerangkatID: req.DeviceID,
			Value:       req.OpenTarget,
			Instruction: instruction,
			GateID:      wcdObj.WaterChannelDevice.GroupRelay,
			IPAddress:   wcdObj.WaterChannelDevice.IPAddress,
			Position:    wcdObj.WaterChannelDevice.MaxHeightSensor - req.OpenTarget,
			FullTime:    wcdObj.WaterChannelDevice.FullTime,
			MaxHeight:   100,
			Limit:       true,
			Now:         req.Now,
		}); err != nil {
			status = sharedModel.StatusGagal
			errorMessage = err.Error()
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
			DoorName:           wcdObj.WaterChannelDevice.Name,
			OpenTarget:         req.OpenTarget,
			OpenCurrent:        waterGatesRes.WaterGates.GateLevel,
			Type:               sharedModel.TypeLangsung,
			Reason:             req.Reason,
			Status:             status,
			OfficerId:          req.OfficerId,
			OfficerName:        officer.User.Name,
			ErrorMessage:       errorMessage,
		}

		if _, err = saveHistory(ctx, gateway.DoorControlHistorySaveReq{DoorControlHistory: &dch}); err != nil {
			return nil, err
		}

		if status == sharedModel.StatusGagal {
			return nil, fmt.Errorf(errorMessage)
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
				Description:  fmt.Sprintf("%s membuat pengontrolan langsung untuk perangkat \"%s\" di pintu air %s ke %.2f", officer.User.Name, wcdObj.WaterChannelDevice.Name, waterChannelDoor.WaterChannelDoor.Name, req.OpenTarget),
			},
		})
		if err != nil {
			return nil, err
		}

		return &DoorControlRunDirectlyRes{}, nil
	}
}

func checkSecurityRelay(device sharedModel.WaterChannelDevice, waterGates sharedModel.WaterGateData) bool {

	return false
}
