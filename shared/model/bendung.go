package model

import "gorm.io/datatypes"

type Bendung struct {
	Geometry                datatypes.JSON `gorm:"type:geometry(POINT)"`
	DASID                   string         `json:"das_id"`
	DASName                 string         `json:"das_name"`
	DebitIntakeMusimHujan   *float64       `json:"debit_intake_musim_hujan"`
	DebitIntakeMusimKemarau *float64       `json:"debit_intake_musim_kemarau"`
	ID                      string         `json:"id"`
	IsPUPR                  bool           `json:"is_pupr"`
	JenisBendung            *string        `json:"jenis_bendung"`
	Kabupaten               string         `json:"kabupaten"`
	Kecamatan               string         `json:"kecamatan"`
	Kelurahan               string         `json:"kelurahan"`
	Kewenangan              string         `json:"kewenangan"`
	Kondisi                 *string        `json:"kondisi"`
	LebarBendung            *float64       `json:"lebar_bendung"`
	Name                    string         `json:"name"`
	PengelolaID             int            `json:"pengelola_id"`
	PengelolaName           string         `json:"pengelola_name"`
	Provinsi                string         `json:"provinsi"`
	RefID                   *string        `json:"ref_id"`
	TinggiBendung           *float64       `json:"tinggi_bendung"`
	TypeID                  int            `json:"type_id"`
	TypeName                string         `json:"type_name"`
	WSID                    string         `json:"ws_id"`
	WSName                  string         `json:"ws_name"`
}
