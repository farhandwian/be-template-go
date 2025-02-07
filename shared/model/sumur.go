package model

import "gorm.io/datatypes"

type Sumur struct {
	Geometry      datatypes.JSON `gorm:"type:geometry(POINT)"`
	DASID         string         `json:"das_id"`
	DASName       string         `json:"das_name"`
	ID            string         `json:"id"`
	IsPUPR        bool           `json:"is_pupr"`
	Kabupaten     string         `json:"kabupaten"`
	Kecamatan     string         `json:"kecamatan"`
	Kelurahan     string         `json:"kelurahan"`
	Kewenangan    string         `json:"kewenangan"`
	Name          string         `json:"name"`
	PengelolaID   int            `json:"pengelola_id"`
	PengelolaName string         `json:"pengelola_name"`
	Provinsi      string         `json:"provinsi"`
	RefID         string         `json:"ref_id"`
	TypeID        int            `json:"type_id"`
	TypeName      string         `json:"type_name"`
	WSID          string         `json:"ws_id"`
	WSName        string         `json:"ws_name"`
}
