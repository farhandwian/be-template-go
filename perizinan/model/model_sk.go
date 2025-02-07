package model

import (
	"fmt"
	"time"
)

type SKPerizinanID string

type NomorSK string

func (p NomorSK) Validate() error {

	if p == "" {
		return fmt.Errorf("nomor sk tidak boleh kosong")
	}

	return nil
}

type SKPerizinan struct {
	ID                         SKPerizinanID     `json:"id"`
	Status                     SKPerizinanStatus `json:"status"`
	NoSK                       NomorSK           `json:"no_sk" gorm:"unique"`
	MasaBerlakuSK              string            `json:"masa_berlaku_sk"`
	KoordinatDiDalamSK         Geometry          `json:"koordinat_di_dalam_sk"`
	CaraPengambilanSDADalamSK  string            `json:"cara_pengambilan_sda_dalam_sk"`
	JenisTipeKonstruksiDalamSK string            `json:"jenis_tipe_konstruksi_dalam_sk"`
	KuotaAirDalamSK            float64           `json:"kuota_air_dalam_sk"`
	JadwalPengambilanDalamSK   string            `json:"jadwal_pengambilan_dalam_sk"`
	JadwalPembangunanDalamSK   string            `json:"jadwal_pembangunan_dalam_sk"`
	KetentuanTeknisLainnya     string            `json:"ketentuan_teknis_lainnya"`

	Perpanjangan      string `json:"perpanjangan"`
	TanggalSK         string `json:"tanggal_sk"`
	JenisUsaha        string `json:"jenis_usaha"`
	KabKota           string `json:"kab_kota"`
	Kecamatan         string `json:"kecamatan"`
	Desa              string `json:"desa"`
	SumberAir         string `json:"sumber_air"`
	AlamatPemohon     string `json:"alamat_pemohon"`
	PerusahaanPemohon string `json:"perusahaan_pemohon"`
	Pemohon           string `json:"pemohon"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SKPerizinanStatus string

const (
	SKPerizinanStatusBerlaku     SKPerizinanStatus = "Berlaku"
	SKPerizinanStatusKadaluwarsa SKPerizinanStatus = "Tidak Berlaku"
)
