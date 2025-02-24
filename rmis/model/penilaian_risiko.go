package model

import "time"

type PenilaianRisiko struct {
	ID                        string    `db:"id"`
	NamaPemda                 string    `db:"nama_pemda"`
	TahunPenilaian            time.Time `db:"tahun_penilaian"`
	TujuanStrategis           string    `db:"tujuan_strategis"`
	UrusanPemerintahan        string    `db:"urusan_pemerintahan"`
	RisikoPrioritas           string    `db:"risiko_prioritas"`
	KodeRisiko                string    `db:"kode_resiko"`
	UraianPengendalian        string    `db:"uraian_pengendalian"`
	CelahPengendalian         string    `db:"celah_pengendalian"`
	RencanaTindakPengendalian string    `db:"rencana_tindak_pengendalian"`
	PemilikPenanggungJawab    string    `db:"pemilik_penanggung_jawab"`
	TargetWaktuPenyelesaian   string    `db:"target_waktu_penyelesaian"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at"`
}
