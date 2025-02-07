package usecase

import (
	"bigboard/gateway"
	"bigboard/model"
	"context"
	"encoding/json"
	"fmt"
	"shared/core"
	sharedGateway "shared/gateway"
	firebaseHelper "shared/helper/firebase"
	"time"

	"firebase.google.com/go/messaging"
)

type AiDoorControlRunStoreReq struct {
	Now                time.Time
	WaterChannelDoorID *int
	ControllerIndex    *int     `json:"controller_index"`
	UpBy               *float32 `json:"up_by"`
	DownBy             *float32 `json:"down_by"`
	SetTo              *float32 `json:"set_to"`
	Access             string
}

type AiDoorControlRunStoreRes struct {
	Message    string `json:"message"`
	Name       string `json:"name"`
	ExternalID int    `json:"external_id"`
	Status     string `json:"status"`
	GateLevel  int    `json:"gate_level"`
	OpenTarget int    `json:"open_target"`
}

type AiDoorControlRunStoreUseCase = core.ActionHandler[AiDoorControlRunStoreReq, AiDoorControlRunStoreRes]

func ImplAiDoorControlRunDirectly(
	generateId gateway.GenerateId,
	save gateway.SensitiveJobsSave,
	sendSSEGateway sharedGateway.SendSSEMessage,
	getWaterChannelDoor sharedGateway.GetWaterChannelDoorByID,
	getWaterChannelDevices sharedGateway.GetWaterChannelDevicesByDoorID,
	getWaterGate sharedGateway.GetLatestWaterGatesByDoorID,
	getOneFCMToken gateway.GetOneFCMToken,
) AiDoorControlRunStoreUseCase {
	return func(ctx context.Context, req AiDoorControlRunStoreReq) (*AiDoorControlRunStoreRes, error) {
		firebaseClient, err := firebaseHelper.InitFirebaseApp()
		if err != nil {
			return nil, fmt.Errorf("failed to initialize Firebase app: %w", err)
		}

		client, err := firebaseClient.Messaging(ctx)
		if err != nil {
			return nil, err
		}

		if req.WaterChannelDoorID == nil {
			return nil, fmt.Errorf("water channel door id is required")
		}

		waterChannelDoorObj, err := getWaterChannelDoor(ctx, sharedGateway.GetWaterChannelDoorByIDReq{WaterChannelDoorID: *req.WaterChannelDoorID})
		if err != nil {
			return nil, err
		}

		// Generate transaction id
		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		devicesRes, err := getWaterChannelDevices(ctx, sharedGateway.GetWaterChannelDevicesByDoorIDReq{WaterChannelDoorID: *req.WaterChannelDoorID})
		if err != nil {
			return nil, err
		}

		var deviceIDs []int
		for _, x := range devicesRes.Devices {
			deviceIDs = append(deviceIDs, x.ExternalID)
		}

		waterGates, err := getWaterGate(ctx, sharedGateway.GetLatestWaterGatesByDoorIDReq{
			WaterChannelDoorID: *req.WaterChannelDoorID,
			DeviceIDs:          deviceIDs,
		})
		if err != nil {
			return nil, err
		}

		if req.ControllerIndex == nil {
			return nil, fmt.Errorf("controller index is required")
		}

		if *req.ControllerIndex >= len(waterGates.WaterGates) {
			return nil, fmt.Errorf("controller index %d is invalid", *req.ControllerIndex)
		}

		waterGate := waterGates.WaterGates[*req.ControllerIndex]
		if !waterGate.SecurityRelay {
			return nil, fmt.Errorf("security relay tidak aktif")
		}

		openTarget := waterGate.GateLevel

		if req.SetTo != nil && *req.SetTo >= 0 {
			openTarget = *req.SetTo
		} else if req.DownBy != nil && *req.DownBy > 0 {
			openTarget -= *req.DownBy
		} else if req.UpBy != nil && *req.UpBy > 0 {
			openTarget += *req.UpBy
		}

		if openTarget < 0 || openTarget > 100 {
			return nil, fmt.Errorf("open target %f is invalid", openTarget)
		}

		payload := map[string]interface{}{
			"water_channel_door_id": *req.WaterChannelDoorID,
			"device_id":             waterGate.DeviceID,
			"open_target":           openTarget,
			"reason":                "AI Door Control Run Directly",
			"timestamp":             req.Now.Format(time.RFC3339),
		}
		payloadJson, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}

		obj := model.SensitiveJobs{
			ID:       model.SensitiveJobsID(genObj.RandomId),
			FuncName: model.FuncTypeDoorControl,
			Status:   model.StatusCreated,
			Payload:  payloadJson,
		}

		if _, err := save(ctx, gateway.SensitiveJobsSaveReq{SensitiveJobs: obj}); err != nil {
			return nil, err
		}

		notificationBody := fmt.Sprintf("Anda akan menjalankan pintu air %s dengan tinggi bukaan air sebesar %s?", (waterChannelDoorObj.WaterChannelDoor.Name), fmt.Sprintf("%.2f", openTarget))
		payloadNotification := map[string]interface{}{
			"id_sensitive_job": obj.ID,
			"endpoint_target":  "/bigboard/ai/doorcontrol-run",
			"message":          notificationBody,
		}
		payloadNotificationJson, err := json.Marshal(payloadNotification)
		if err != nil {
			return nil, err
		}

		res, err := getOneFCMToken(ctx, gateway.GetOneFCMTokenReq{ID: 1})

		if err != nil {
			return nil, err
		}

		message := &messaging.Message{
			Notification: &messaging.Notification{
				Title: "Door Control",
				Body:  notificationBody,
			},
			Token: res.Token,
			Data: map[string]string{
				"payload": string(payloadNotificationJson),
			},
		}
		if _, err := client.Send(ctx, message); err != nil {
			return nil, err
		}

		// Send SSE message for waiting authorization
		_, _ = sendSSEGateway(ctx, sharedGateway.SendSSEMessageReq{
			Subject:      "door-control-run-request",
			FunctionName: "doorControlRun",
			Data: &AiDoorControlRunStoreRes{
				Message:    "Menunggu otorisasi",
				GateLevel:  int(waterGate.GateLevel),
				OpenTarget: int(openTarget),
				Name:       waterChannelDoorObj.WaterChannelDoor.Name,
				ExternalID: waterChannelDoorObj.WaterChannelDoor.ExternalID,
				Status:     string(obj.Status),
			},
		})

		messageBody := "Permintaan pengontrolan sudah diterima. Silakan masukan pin otorisasi pada perangkat autentikator untuk memulai."

		return &AiDoorControlRunStoreRes{
			Message:    messageBody,
			GateLevel:  int(waterGate.GateLevel),
			OpenTarget: int(openTarget),
			Name:       waterChannelDoorObj.WaterChannelDoor.Name,
			ExternalID: waterChannelDoorObj.WaterChannelDoor.ExternalID,
			Status:     string(obj.Status),
		}, nil
	}
}
