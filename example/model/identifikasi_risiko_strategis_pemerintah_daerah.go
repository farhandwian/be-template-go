package model

import (
	"fmt"
	"time"
)

type IdentifikasiRisikoStrategisPemerintahDaerah struct {
	ID                 *string    `json:"id"`
	NamaPemda          *string    `json:"nama_pemda"`
	TahunPenilaian     *time.Time `json:"tahun_penilaian"`
	Periode            *string    `json:"periode"`
	UrusanPemerintahan *string    `json:"urusan_pemerintahan"`
	TujuanStrategis    *string    `json:"tujuan_strategis"`
	IndikatorKinerja   *string    `json:"indikator_kinerja"`
	KategoriRisiko     *string    `json:"kategori_resiko"` // references kategori_resiko.id
	UraianRisiko       *string    `json:"uraian_resiko"`
	NomorUraian        *int       `json:"nomor_risiko"`
	KodeRisiko         *string    `json:"kode_resiko"`
	PemilikRisiko      *string    `json:"pemilik_resiko"`
	UraianSebab        *string    `json:"uraian_sebab"` // references rca.akar_penyebab
	SumberSebab        *string    `json:"sumber_sebab"` // references rca.jenis_penyebab
	Controllable       *string    `json:"controllable"` // could be boolean if desired
	UraianDampak       *string    `json:"uraian_dampak"`
	PihakDampak        *string    `json:"pihak_dampak"`
}

func (irspd *IdentifikasiRisikoStrategisPemerintahDaerah) GenerateKodeRisiko(kategori_risiko string) error {
	if irspd.TahunPenilaian == nil {
		return fmt.Errorf("TahunPenilaian is nil")
	}

	if irspd.NomorUraian == nil {
		return fmt.Errorf("NomorUraian is nil")
	}
	yearSuffix := fmt.Sprintf("%02d", irspd.TahunPenilaian.Year()%100)
	iterStr := fmt.Sprintf("%03d", *irspd.NomorUraian)

	kodeRisiko := fmt.Sprintf("RSP.%s.%s.%s", yearSuffix, kategori_risiko, iterStr)
	irspd.KodeRisiko = &kodeRisiko

	return nil
}
