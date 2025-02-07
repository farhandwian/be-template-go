package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
	"shared/model"
	"sort"
	"strings"
	"time"
)

type ChartPintuAirReq struct {
	WaterChannelDoorID int       `json:"water_channel_door_id"`
	MinTime            time.Time `json:"min_time"`
	MaxTime            time.Time `json:"max_time"`
}

type ResponseChart map[string]interface{}

type ChartPintuAirRes struct {
	Charts []ResponseChart `json:"charts"`
}

type ChartPintuAirUseCase = core.ActionHandler[ChartPintuAirReq, ChartPintuAirRes]

func ImplChartPintuAirUseCase(
	getDevice gateway.GetDevice,
	getGate gateway.GetGate,
) ChartPintuAirUseCase {
	return func(ctx context.Context, req ChartPintuAirReq) (*ChartPintuAirRes, error) {

		duration := req.MaxTime.Sub(req.MinTime)
		var bucketDuration time.Duration
		if duration > 7*24*time.Hour {
			bucketDuration = 24 * time.Hour
		} else if duration > 24*time.Hour {
			bucketDuration = time.Hour
		} else {
			bucketDuration = 5 * time.Minute
		}

		// Fetch device data
		devicesObj, err := getDevice(ctx, gateway.GetDeviceReq{
			WaterChannelDoorID: req.WaterChannelDoorID,
		})
		if err != nil {
			return nil, err
		}

		// Prepare device names and mappings
		deviceNames := make([]string, len(devicesObj.Devices))
		deviceNameMap := make(map[int]string)
		for i, device := range devicesObj.Devices {
			gateName := strings.ToLower(strings.ReplaceAll(device.Name, " ", "_"))
			deviceNames[i] = gateName
			deviceNameMap[device.ExternalID] = gateName
		}
		sort.Strings(deviceNames)

		// Collect gate data asynchronously
		gatesChan := make(chan gateResult, len(devicesObj.Devices))
		for _, device := range devicesObj.Devices {
			go func(dev model.WaterChannelDevice) {
				gateObj, err := getGate(ctx, gateway.GetGateReq{
					WaterChannelDoorID: req.WaterChannelDoorID,
					DeviceID:           dev.ExternalID,
					MinTime:            req.MinTime,
					MaxTime:            req.MaxTime,
				})
				gatesChan <- gateResult{
					deviceID: dev.ExternalID,
					gates:    gateObj,
					err:      err,
				}
			}(device)
		}

		// Aggregate gate data into a time map
		timeMap := make(map[time.Time]ResponseChart)
		for i := 0; i < len(devicesObj.Devices); i++ {
			select {
			case result := <-gatesChan:
				if result.err != nil {
					return nil, result.err
				}

				gateName := deviceNameMap[result.deviceID]
				for _, gate := range result.gates.Gates {
					// Adjust timestamp to the nearest bucket duration
					bucketedTime := gate.Time.Truncate(bucketDuration)

					if _, exists := timeMap[bucketedTime]; !exists {
						// Initialize time entry with zero values for all gates
						timeMap[bucketedTime] = ResponseChart{
							"time": bucketedTime,
						}
						for _, name := range deviceNames {
							timeMap[bucketedTime][name] = float64(0)
						}
					}
					// Assign gate value to the respective time entry
					timeMap[bucketedTime][gateName] = gate.GateLevel
				}

			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}

		// Prepare charts from timeMap
		charts := make([]ResponseChart, 0, len(timeMap))
		for _, chart := range timeMap {
			// Include chart only if it contains data
			hasData := false
			for _, name := range deviceNames {
				if value, ok := chart[name].(float64); ok && value != 0 {
					hasData = true
					break
				}
			}
			if hasData {
				charts = append(charts, chart)
			}
		}

		// Sort charts by timestamp in descending order
		sort.Slice(charts, func(i, j int) bool {
			return charts[i]["time"].(time.Time).Before(charts[j]["time"].(time.Time))
		})

		return &ChartPintuAirRes{
			Charts: charts,
		}, nil
	}
}

type gateResult struct {
	deviceID int
	gates    *gateway.GetGateRes
	err      error
}
