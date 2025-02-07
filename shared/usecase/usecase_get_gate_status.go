package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
	"shared/model"
	"sort"
)

type GetGateStatusReq struct {
	WaterChannelDoorID int
}

type GetGateStatusRes struct {
	Gates []Gate `json:"gates"`
}

type Gate struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`           // get from WaterChannelDevice.Name where Category = "controller" and WaterChannelDevice.WaterChannelDoorID = x
	GateLevel     float32 `json:"gate_level"`     // get from WaterGate.GateLevel where WaterGate.WaterChannelDoorID = x and WaterGate.GroupRelay is same as WaterChannelDevice.GroupRelay
	GroupRelay    int     `json:"group_relay"`    // get from WaterChannelDevice.GroupRelay where Category = "controller" and WaterChannelDevice.WaterChannelDoorID = x
	SecurityRelay bool    `json:"security_relay"` // get from WaterGate.SecurityRelay where WaterGate.WaterChannelDoorID = x and WaterGate.GroupRelay is same as WaterChannelDevice.GroupRelay
	Sensor        bool    `json:"sensor"`         // get from WaterGate.Sensor where WaterGate.WaterChannelDoorID = x and WaterGate.GroupRelay is same as WaterChannelDevice.GroupRelay
	Status        bool    `json:"status"`         // get from WaterGate.Status where WaterGate.WaterChannelDoorID = x and WaterGate.GroupRelay is same as WaterChannelDevice.GroupRelay
	Run           bool    `json:"run"`            // get from WaterGate.Run where WaterGate.WaterChannelDoorID = x and WaterGate.GroupRelay is same as WaterChannelDevice.GroupRelay
	UpperLimit    int     `json:"upper_limit"`    // get from WaterChannelDevice.UpperLimit where Category = "controller" and WaterChannelDevice.WaterChannelDoorID = x
	LowerLimit    int     `json:"lower_limit"`    // get from WaterChannelDevice.LowerLimit where Category = "controller" and WaterChannelDevice.WaterChannelDoorID = x
}

type GetGateStatusUseCase = core.ActionHandler[GetGateStatusReq, GetGateStatusRes]

func ImplGetGateStatus(
	getWaterChannelDevices gateway.GetWaterChannelDevicesByDoorID,
	getWaterGate gateway.GetLatestWaterGatesByDoorID,

) GetGateStatusUseCase {
	return func(ctx context.Context, req GetGateStatusReq) (*GetGateStatusRes, error) {

		devicesRes, err := getWaterChannelDevices(ctx, gateway.GetWaterChannelDevicesByDoorIDReq{WaterChannelDoorID: req.WaterChannelDoorID})
		if err != nil {
			return nil, err
		}

		var deviceIDs []int
		for _, x := range devicesRes.Devices {
			deviceIDs = append(deviceIDs, x.ExternalID)
		}

		waterGatesRes, err := getWaterGate(ctx, gateway.GetLatestWaterGatesByDoorIDReq{
			WaterChannelDoorID: req.WaterChannelDoorID,
			DeviceIDs:          deviceIDs,
		})
		if err != nil {
			return nil, err
		}

		return &GetGateStatusRes{
			Gates: mapGates(devicesRes.Devices, waterGatesRes.WaterGates),
		}, nil
	}
}

func mapGates(devices []model.WaterChannelDevice, waterGates []model.WaterGateData) []Gate {
	gateMap := make(map[int]Gate)
	for _, d := range devices {
		if d.Category != "controller" {
			continue
		}

		upperLimit := 100
		if d.UpperLimit != nil {
			upperLimit = *d.UpperLimit
		}

		lowerLimit := 0
		if d.LowerLimit != nil {
			lowerLimit = *d.LowerLimit
		}

		gateMap[d.ExternalID] = Gate{
			GroupRelay: d.GroupRelay,
			Name:       d.Name,
			ID:         d.ExternalID,
			UpperLimit: upperLimit,
			LowerLimit: lowerLimit,
		}
	}

	for _, wg := range waterGates {
		if gate, ok := gateMap[wg.DeviceID]; ok {
			gate.GateLevel = wg.GateLevel
			gate.SecurityRelay = wg.SecurityRelay
			gate.Status = wg.Status
			gate.Sensor = wg.Sensor
			gate.Run = wg.Run
			gateMap[wg.DeviceID] = gate
		}
	}

	var gates []Gate
	for _, gate := range gateMap {
		gates = append(gates, gate)
	}

	sort.Slice(gates, func(i, j int) bool {
		return gates[i].Name < gates[j].Name
	})

	return gates
}
