package model

type DeviceStatus struct {
	Up    int `json:"up"`
	Down  int `json:"down"`
	Total int `json:"total"`
}
