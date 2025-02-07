package model

import "gorm.io/datatypes"

type Embung struct {
	Geometry                 datatypes.JSON `gorm:"type:geometry(POINT)"`
	DASID                    string         `json:"das_id"`
	DASName                  string         `json:"das_name"`
	ElevasiMukaAirMax        float64        `json:"elevasi_muka_air_max"`
	ElevasiMukaAirMin        float64        `json:"elevasi_muka_air_min"`
	ElevasiMukaAirNormal     float64        `json:"elevasi_muka_air_normal"`
	ElevasiMukaAirPuncak     float64        `json:"elevasi_muka_air_puncak"`
	ID                       string         `json:"id"`
	IsPUPR                   bool           `json:"is_pupr"`
	JenisEmbung              string         `json:"jenis_embung"`
	Kabupaten                string         `json:"kabupaten"`
	Kecamatan                string         `json:"kecamatan"`
	Kelurahan                string         `json:"kelurahan"`
	Kewenangan               string         `json:"kewenangan"`
	LebarTubuhEmbung         float64        `json:"lebar_tubuh_embung"`
	Name                     string         `json:"name"`
	PanjangTubuhEmbung       float64        `json:"panjang_tubuh_embung"`
	PengelolaID              int            `json:"pengelola_id"`
	PengelolaName            string         `json:"pengelola_name"`
	Provinsi                 string         `json:"provinsi"`
	RefID                    string         `json:"ref_id"`
	TinggiTubuhEmbungPondasi float64        `json:"tinggi_tubuh_embung_pondasi"`
	TinggiTubuhEmbungSungai  float64        `json:"tinggi_tubuh_embung_sungai"`
	TipeTubuhEmbung          string         `json:"tipe_tubuh_embung"`
	TMA                      float64        `json:"tma"`
	TypeID                   int            `json:"type_id"`
	TypeName                 string         `json:"type_name"`
	VolumeTubuhEmbung        float64        `json:"volume_tubuh_embung"`
	WSID                     string         `json:"ws_id"`
	WSName                   string         `json:"ws_name"`
}
