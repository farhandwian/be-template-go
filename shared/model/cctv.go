package model

type CCTV struct {
	ExternalID      int    `json:"external_id"`
	Name            string `json:"name"`       // get from WaterChannelDevice.Name where Category = "cctv" and WaterChannelDevice.WaterChannelDoorID = x
	IPAddress       string `json:"ip_address"` // get from WaterChannelDevice.IPAddress where Category = "cctv" and WaterChannelDevice.WaterChannelDoorID = x
	GarbageDetected bool   `json:"garbage_detected"`
	HumanDetected   bool   `json:"human_detected"`
}

func MapCCTVs(devices []WaterChannelDevice) []CCTV {
	var cctvs []CCTV
	for _, d := range devices {
		if d.Category == "cctv" {

			humanDetected := false
			garbageDetected := false
			if d.DetectedObject == "human" {
				humanDetected = true
			} else if d.DetectedObject == "garbage" {
				garbageDetected = true
			}

			cctvs = append(cctvs, CCTV{
				ExternalID:      d.ExternalID,
				Name:            d.Name,
				IPAddress:       d.IPAddress,
				HumanDetected:   humanDetected,
				GarbageDetected: garbageDetected,
			})
		}
	}
	return cctvs
}
