package model

import (
	"fmt"
	"time"
)

type IdentifikasiRisikoStrategisOpd struct {
	ID                     *string    `json:"id"`
	NamaPemda              *string    `json:"nama_pemda"`
	NamaOPD                *string    `json:"nama_opd"` // originally a reference to opd.id
	TahunPenilaian         *time.Time `json:"tahun_penilaian"`
	Periode                *string    `json:"periode"`
	UrusanPemerintahan     *string    `json:"urusan_pemerintahan"`
	TujuanSasaranStrategis *string    `json:"tujuan_sasaran_strategis"`
	IndikatorKinerja       *string    `json:"indikator_kinerja"`
	KategoriRisiko         *string    `json:"kategori_risiko"` // references kategori_resiko.nama
	NomorUraianRisiko      *int       `json:"nomor_uraian_risiko"`
	UraianRisiko           *string    `json:"uraian_risiko"`
	KodeRisiko             *string    `json:"kode_risiko"`
	PemilikRisiko          *string    `json:"pemilik_risiko"`
	UraianSebab            *string    `json:"uraian_sebab"` // references rca.akar_penyebab
	SumberSebab            *string    `json:"sumber_sebab"` // references rca.jenis_penyebab
	Controllable           *string    `json:"controllable"`
	UraianDampak           *string    `json:"uraian_dampak"`
	PihakDampak            *string    `json:"pihak_dampak"`
}

func (irsopd *IdentifikasiRisikoStrategisOpd) GenerateKodeRisiko(kodeKategoriRisiko string, kodeOpd string) error {
	if irsopd.TahunPenilaian == nil {
		return fmt.Errorf("TahunPenilaian is nil")
	}

	if irsopd.NamaOPD == nil {
		return fmt.Errorf("NamaOPD is nil")
	}

	if irsopd.NomorUraianRisiko == nil {
		return fmt.Errorf("NomorUraianRisiko is nil")
	}

	if irsopd.KategoriRisiko == nil {
		return fmt.Errorf("KategoriRisiko is nil")
	}

	yearSuffix := fmt.Sprintf("%02d", irsopd.TahunPenilaian.Year()%100)
	iterStr := fmt.Sprintf("%03d", *irsopd.NomorUraianRisiko)

	kodeRisiko := fmt.Sprintf("RSO.%s.%s.%s.%s", yearSuffix, kodeOpd, kodeKategoriRisiko, iterStr)
	irsopd.KodeRisiko = &kodeRisiko
	return nil
}
