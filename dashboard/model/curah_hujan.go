package model

import "gorm.io/datatypes"

// TODO: remove this because it's not used (using rainfall post instead)
type CurahHujan struct {
	Geometry         datatypes.JSON `gorm:"type:geometry(POINT)"`
	ID               string         `json:"id"`
	Name             string         `json:"name"`
	InfrastructureID int            `json:"infrastructure_id"`
	Provinsi         string         `json:"provinsi"`
	Kabupaten        string         `json:"kabupaten"`
	Kecamatan        string         `json:"kecamatan"`
	Kelurahan        string         `json:"kelurahan"`
	IsPUPR           bool           `json:"is_pupr"`
	WSID             string         `json:"ws_id"`
	WSName           string         `json:"ws_name"`
	DASID            string         `json:"das_id"`
	DASName          string         `json:"das_name"`
	PengelolaID      int            `json:"pengelola_id"`
	PengelolaName    string         `json:"pengelola_name"`
	Type             string         `json:"type"`
	AssetKey         string         `json:"asset_key"`
	SensorType       datatypes.JSON `json:"sensor_type" `
	// Coordinate string `json:"coordinate"`
}
