package usecase

import (
	bigboardGateaway "bigboard/gateway"
	bigboardModel "bigboard/model"
	"context"
	"dashboard/gateway"
	"encoding/json"
	"fmt"
	iamgw "iam/gateway"
	iammodel "iam/model"
	"sync"

	"shared/constant"
	"shared/core"
	sharedGateway "shared/gateway"
	sharedModel "shared/model"
	"shared/usecase"
	"time"

	"github.com/google/uuid"
)

type AIDoorControlRunDirectlyReq struct {
	IDSensitveJob string          `json:"id_sensitive_job"`
	Pin           string          `json:"pin"`
	OfficerId     iammodel.UserID `json:"-"`
}

type DoorControlRunDirectlyRes struct {
	ExternalID            int               `json:"external_id"`
	Name                  string            `json:"name"`
	DebitActual           string            `json:"actual_debit"`      // get from WaterChannelDoor.DebitActual where WaterChannelDoor.ExternalID = x
	DebitRequirement      string            `json:"debit_requirement"` // get from WaterChannelDoor.DebitRequirement where WaterChannelDoor.ExternalID = x
	WaterSurfaceElevation string            `json:"water_surface_elevation"`
	Officers              []usecase.Officer `json:"officers"` // see `Officer` struct
	NCCTV                 int               `json:"number_of_cctv"`
	Latitude              string            `json:"latitude"`
	Longitude             string            `json:"longitude"`
	Status                string            `json:"status"`
}

type DoorControlRunDirectly struct {
	Now                time.Time `json:"timestamp"`
	WaterChannelDoorID int       `json:"water_channel_door_id"`
	DeviceID           int       `json:"device_id"`
	OpenTarget         float32   `json:"open_target"`
	Reason             string    `json:"reason"`
}

type AIDoorControlRunDirectlyUseCase = core.ActionHandler[AIDoorControlRunDirectlyReq, DoorControlRunDirectlyRes]

func ImplAIDoorControlRunDirectly(
	pinValidate iamgw.PasswordValidate,
	callApi gateway.DoorControlAPI,
	generateId gateway.GenerateId,
	getSenstiveJob bigboardGateaway.GetListSensitiveJobGateway,
	updateSensitiveJob bigboardGateaway.SensitiveJobsSave,
	saveHistory gateway.DoorControlHistorySave,
	getWaterGate gateway.GetOneWaterGatesByDoorID,
	getOfficer iamgw.UserGetOneByID,
	getWaterChannelDevice gateway.GetOneWaterChannelDeviceByDoorID,
	getWaterChannelDoorByID sharedGateway.GetWaterChannelDoorByID,
	getWaterChannelDevicesByDoorID sharedGateway.GetWaterChannelDevicesByDoorID,
	getWaterChannelOfficersByDoorID sharedGateway.GetWaterChannelOfficersByDoorID,
	getLatestDebitGateway sharedGateway.GetLatestDebit,
	getWaterSurfaceElevationGateway sharedGateway.GetLatestWaterSurfaceElevationByDoorID,
	sendSSEGateway sharedGateway.SendSSEMessage,
	createActivityMonitoring sharedGateway.CreateActivityMonitoringGateway,
) AIDoorControlRunDirectlyUseCase {
	return func(ctx context.Context, req AIDoorControlRunDirectlyReq) (*DoorControlRunDirectlyRes, error) {
		var err error
		var status sharedModel.DoorControlStatus = sharedModel.StatusDieksekusi
		var errorMessage string

		var sensitiveJob *bigboardModel.SensitiveJobs
		defer func() {
			if err != nil {
				_, _ = sendSSEGateway(ctx, sharedGateway.SendSSEMessageReq{
					Subject:      "door-control-run",
					FunctionName: "doorControlRun",
					Data: map[string]interface{}{
						"message":    "Tinggi Pintu Air Gagal Diubah",
						"is_success": false,
					},
				})
				if sensitiveJob != nil && err.Error() == "invalid PIN" {
					sensitiveJob.Status = bigboardModel.StatusFailed
					if _, updateErr := updateSensitiveJob(ctx, bigboardGateaway.SensitiveJobsSaveReq{SensitiveJobs: *sensitiveJob}); updateErr != nil {
						err = updateErr
					}
				}
			}

		}()

		sensitiveJob, err = getSenstiveJob(ctx, bigboardGateaway.GetListSensitiveJobReq{ID: req.IDSensitveJob})
		if err != nil {
			return nil, err
		}

		if sensitiveJob.Status != bigboardModel.StatusCreated {
			err = fmt.Errorf("job with id %v has status other than created", req.IDSensitveJob)
			return nil, err
		}

		var doorControlRunDirectly DoorControlRunDirectly
		if err = json.Unmarshal(sensitiveJob.Payload, &doorControlRunDirectly); err != nil {
			return nil, err
		}

		officer, err := getOfficer(ctx, iamgw.UserGetOneByIDReq{UserID: req.OfficerId})
		if err != nil || officer == nil {
			err = fmt.Errorf("officer with id %v not found", req.OfficerId)
			return nil, err
		}

		if _, err = pinValidate(ctx, iamgw.PasswordValidateReq{
			PasswordPlain:     req.Pin,
			PasswordEncrypted: officer.User.Pin,
		}); err != nil {
			return nil, fmt.Errorf("invalid PIN")
		}

		sensitiveJob.Status = bigboardModel.StatusRunning
		if _, err = updateSensitiveJob(ctx, bigboardGateaway.SensitiveJobsSaveReq{SensitiveJobs: *sensitiveJob}); err != nil {
			return nil, err
		}

		wcdObj, err := getWaterChannelDevice(ctx, gateway.GetOneWaterChannelDeviceByDoorIDReq{
			WaterChannelDoorID: doorControlRunDirectly.WaterChannelDoorID,
			DeviceID:           doorControlRunDirectly.DeviceID,
		})
		if err != nil || wcdObj == nil {
			err = fmt.Errorf("door with gate_id %v and device_id %v not found",
				doorControlRunDirectly.WaterChannelDoorID, doorControlRunDirectly.DeviceID)
			return nil, err
		}

		waterGatesRes, err := getWaterGate(ctx, gateway.GetOneWaterGatesByDoorIDReq{
			WaterChannelDoorID: doorControlRunDirectly.WaterChannelDoorID,
			DeviceID:           doorControlRunDirectly.DeviceID,
		})
		if err != nil {
			return nil, err
		}

		if !waterGatesRes.WaterGates.SecurityRelay {
			err = fmt.Errorf("security relay tidak aktif")
			return nil, err
		}

		instruction := "up"
		if waterGatesRes.WaterGates.GateLevel > doorControlRunDirectly.OpenTarget {
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
		_, err = callApi(ctx, gateway.DoorControlAPIReq{
			PerangkatID: doorControlRunDirectly.DeviceID,
			Value:       doorControlRunDirectly.OpenTarget,
			Instruction: instruction,
			GateID:      wcdObj.WaterChannelDevice.GroupRelay,
			IPAddress:   wcdObj.WaterChannelDevice.IPAddress,
			Position:    wcdObj.WaterChannelDevice.MaxHeightSensor - doorControlRunDirectly.OpenTarget,
			FullTime:    wcdObj.WaterChannelDevice.FullTime,
			MaxHeight:   100,
			Limit:       true,
			Now:         doorControlRunDirectly.Now,
		})
		if err != nil {
			status = sharedModel.StatusGagal
			return nil, err
		}

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		dch := sharedModel.DoorControlHistory{
			ID:                 sharedModel.DoorControlHistoryID(genObj.RandomId),
			Date:               doorControlRunDirectly.Now,
			WaterChannelDoorID: doorControlRunDirectly.WaterChannelDoorID,
			DeviceID:           doorControlRunDirectly.DeviceID,
			DoorName:           wcdObj.WaterChannelDevice.Name,
			OpenTarget:         doorControlRunDirectly.OpenTarget,
			OpenCurrent:        waterGatesRes.WaterGates.GateLevel,
			Type:               sharedModel.TypeLangsung,
			Reason:             doorControlRunDirectly.Reason,
			Status:             status,
			OfficerId:          req.OfficerId,
			OfficerName:        officer.User.Name,
			ErrorMessage:       errorMessage,
		}

		if _, err = saveHistory(ctx, gateway.DoorControlHistorySaveReq{DoorControlHistory: &dch}); err != nil {
			return nil, err
		}

		sensitiveJob.Status = bigboardModel.StatusSuccess
		if _, err = updateSensitiveJob(ctx, bigboardGateaway.SensitiveJobsSaveReq{SensitiveJobs: *sensitiveJob}); err != nil {
			return nil, err
		}

		err = nil // Clear error to prevent deferred error handler
		_, _ = sendSSEGateway(ctx, sharedGateway.SendSSEMessageReq{
			Subject:      "door-control-run",
			FunctionName: "doorControlRun",
			Data:         map[string]interface{}{"message": "Tinggi Pintu Air Berhasil Diubah", "is_success": true},
		})

		// Get WaterChannelDoor
		doorRes, err := getWaterChannelDoorByID(ctx, sharedGateway.GetWaterChannelDoorByIDReq{WaterChannelDoorID: doorControlRunDirectly.WaterChannelDoorID})
		if err != nil {
			return nil, err
		}
		door := doorRes.WaterChannelDoor

		// Get WaterChannelDevices
		devicesRes, err := getWaterChannelDevicesByDoorID(ctx, sharedGateway.GetWaterChannelDevicesByDoorIDReq{WaterChannelDoorID: door.ExternalID})
		if err != nil {
			return nil, err
		}

		// Get WaterChannelOfficers
		officersRes, err := getWaterChannelOfficersByDoorID(ctx, sharedGateway.GetWaterChannelOfficersByDoorIDReq{WaterChannelDoorID: door.ExternalID})
		if err != nil {
			return nil, err
		}

		debitRes, err := getLatestDebitGateway(ctx, sharedGateway.GetLatestDebitReq{WaterChannelDoorID: door.ExternalID})
		if err != nil {
			return nil, err
		}

		waterSurfaceElevationsRes, err := getWaterSurfaceElevationGateway(ctx, sharedGateway.GetLatestWaterSurfaceElevationByDoorIDReq{
			WaterChannelDoorID: door.ExternalID,
		})
		if err != nil {
			return nil, err
		}

		_, err = createActivityMonitoring(ctx, sharedGateway.CreateActivityMonitoringReq{
			ActivityMonitor: sharedModel.ActivityMonitor{
				ID:           uuid.NewString(),
				UserName:     officer.User.Name,
				Category:     constant.MONITORING_TYPE_DOOR_CONTROL,
				ActivityTime: time.Now(),
				Description:  fmt.Sprintf("%s mengubah pintu air %s dari %2f menjadi %2f", officer.User.Name, doorRes.WaterChannelDoor.Name, waterGatesRes.WaterGates.GateLevel, doorControlRunDirectly.OpenTarget),
			},
		})
		if err != nil {
			return nil, err
		}

		response := &AiGetDetailWaterChannelResp{
			ExternalID:            door.ExternalID,
			Name:                  door.Name,
			DebitActual:           fmt.Sprintf("%.2f liter per detik", debitRes.Debit.ActualDebit),
			DebitRequirement:      door.DebitRequirement + " liter per detik",
			WaterSurfaceElevation: fmt.Sprintf("%.2f cm", waterSurfaceElevationsRes.WaterSurfaceElevation.WaterLevel),
			Officers:              usecase.MapOfficersForAI(officersRes.Officers),
			NCCTV:                 len(usecase.MapCCTVs(devicesRes.Devices)),
			Latitude:              door.Latitude,
			Longitude:             door.Longitude,
			Status:                getWaterChannelDoorStatus(debitRes.Debit.ActualDebit, usecase.ParseFloat64(door.DebitRequirement)),
		}
		var wg sync.WaitGroup
		wg.Add(1)

		time.AfterFunc(2*time.Second, func() {
			defer wg.Done()
			_, err := sendSSEGateway(ctx, sharedGateway.SendSSEMessageReq{
				Subject:      "get-detail-from-water-channel",
				FunctionName: "showDetailWaterChannel",
				Data:         response,
			})
			if err != nil {
				fmt.Printf("Error sending SSE: %v", err)
			}
		})
		wg.Wait()
		return (*DoorControlRunDirectlyRes)(response), nil
	}
}
