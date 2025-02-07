package model

import "gorm.io/datatypes"

type Bendungan struct {
	Geometry                           datatypes.JSON `gorm:"type:geometry(POINT)"`
	ID                                 string         `json:"id"`
	DasID                              string         `json:"das_id"`
	DasName                            string         `json:"das_name"`
	Dmi                                *float64       `json:"dmi"`
	ElevasiPuncakBendungan             *float64       `json:"elevasi_puncak_bendungan"`
	Irigasi                            *float64       `json:"irigasi"`
	IsPupr                             bool           `json:"is_pupr"`
	Kabupaten                          string         `json:"kabupaten"`
	Kecamatan                          string         `json:"kecamatan"`
	Kelurahan                          string         `json:"kelurahan"`
	KeteranganTypeBendungan            *string        `json:"keterangan_type_bendungan"`
	Kewenangan                         string         `json:"kewenangan"`
	KondisiBangunan                    *string        `json:"kondisi_bangunan"`
	LebarBendungan                     *float64       `json:"lebar_bendungan"`
	LebarPuncakBendungan               *float64       `json:"lebar_puncak_bendungan"`
	LuasGenanganMinimal                *float64       `json:"luas_genangan_minimal"`
	LuasGenanganNormal                 *float64       `json:"luas_genangan_normal"`
	LuasGenanganTotal                  *float64       `json:"luas_genangan_total"`
	Name                               string         `json:"name"`
	PanjangBendungan                   *float64       `json:"panjang_bendungan"`
	PanjangPuncakBendungan             *float64       `json:"panjang_puncak_bendungan"`
	PelimpahDenganPintu                *string        `json:"pelimpah_dengan_pintu"`
	PelimpahElevasiMercuAmbang         *float64       `json:"pelimpah_elevasi_mercu_ambang"`
	PelimpahLebar                      *float64       `json:"pelimpah_lebar"`
	PelimpahLebarAmbang                *float64       `json:"pelimpah_lebar_ambang"`
	PelimpahPanjang                    *float64       `json:"pelimpah_panjang"`
	PelimpahPanjangSaluranTransisi     *float64       `json:"pelimpah_panjang_saluran_transisi"`
	PelimpahSaluran                    *string        `json:"pelimpah_saluran"`
	PelimpahType                       *string        `json:"pelimpah_type"`
	PelimpahTypeKeterangan             *string        `json:"pelimpah_type_keterangan"`
	PengelolaID                        int            `json:"pengelola_id"`
	PengelolaName                      string         `json:"pengelola_name"`
	PenguranganDebitBanjir             *float64       `json:"pengurangan_debit_banjir"`
	PenguranganLuasGenanganBanjirHilir *float64       `json:"pengurangan_luas_genangan_banjir_hilir"`
	Plta                               *float64       `json:"plta"`
	Provinsi                           string         `json:"provinsi"`
	RefID                              string         `json:"ref_id"`
	StatusInfrastructure               *string        `json:"status_infrastructure"`
	TahunPembangunan                   *string        `json:"tahun_pembangunan"`
	TinggiBendungan                    *float64       `json:"tinggi_bendungan"`
	TinggiDasarGalian                  *float64       `json:"tinggi_dasar_galian"`
	TinggiDasarSungai                  *float64       `json:"tinggi_dasar_sungai"`
	TypeBendungan                      *string        `json:"type_bendungan"`
	TypeID                             int            `json:"type_id"`
	TypeName                           string         `json:"type_name"`
	VolumeBendungan                    *float64       `json:"volume_bendungan"`
	VolumeTampungMinimal               *float64       `json:"volume_tampung_minimal"`
	VolumeTampungNormal                *float64       `json:"volume_tampung_normal"`
	VolumeTampungTotal                 *float64       `json:"volume_tampung_total"`
	WsID                               string         `json:"ws_id"`
	WsName                             string         `json:"ws_name"`
}
