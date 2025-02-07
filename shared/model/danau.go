package model

import "gorm.io/datatypes"

type Danau struct {
	Geometry      datatypes.JSON `gorm:"type:geometry(POINT)"`
	DasID         string         `json:"das_id"`
	DasName       string         `json:"das_name"`
	ID            string         `json:"id"`
	IsPupr        bool           `json:"is_pupr"`
	Jenis         *string        `json:"jenis"`
	Kabupaten     string         `json:"kabupaten"`
	Kecamatan     string         `json:"kecamatan"`
	Kelurahan     string         `json:"kelurahan"`
	Kewenangan    string         `json:"kewenangan"`
	Name          string         `json:"name"`
	PengelolaID   int            `json:"pengelola_id"`
	PengelolaName string         `json:"pengelola_name"`
	Provinsi      string         `json:"provinsi"`
	RefID         *string        `json:"ref_id"`
	TypeID        int            `json:"type_id"`
	TypeName      string         `json:"type_name"`
	WsID          string         `json:"ws_id"`
	WsName        string         `json:"ws_name"`
}
