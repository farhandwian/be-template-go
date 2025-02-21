package model

import (
	"fmt"
	"time"
)

// Form 3C

type IdentifikasiRisikoOperasionalOpd struct {
	ID                         *string    `json:"id"`
	IsiForm3B                  *string    `json:"isi_form3b"`
	NamaPemda                  *string    `json:"nama_pemda"`
	NamaOPD                    *string    `json:"nama_opd"` // references opd.nama
	TahunPenilaian             *time.Time `json:"tahun_penilaian"`
	Periode                    *string    `json:"periode"`
	UrusanPemerintahan         *string    `json:"urusan_pemerintahan"`
	ProgramKegiatanSubkegiatan *string    `json:"program_kegiatan_subkegiatan"`
	IndikatorPengeluaran       *string    `json:"indikator_pengeluaran"`
	TahapRisiko                *string    `json:"tahap_risiko"`
	KategoriRisiko             *string    `json:"kategori_risiko"` // references kategori_resiko.nama
	NomorUraianRisiko          *int       `json:"nomor_uraian_risiko"`
	UraianRisiko               *string    `json:"uraian_risiko"`
	KodeRisiko                 *string    `json:"kode_risiko"`
	Pemilik                    *string    `json:"pemilik"`
	UraianSebab                *string    `json:"uraian_sebab"` // references rca.akar_penyebab
	SumberSebab                *string    `json:"sumber_sebab"` // references rca.jenis_penyebab
	Controllable               *string    `json:"controllable"`
	UraianDampak               *string    `json:"uraian_dampak"`
	PihakDampak                *string    `json:"pihak_dampak"`
}

func (iroopd *IdentifikasiRisikoOperasionalOpd) GenerateKodeRisiko(kodeKategoriRisiko string, kodeOpd string) error {
	if iroopd.TahunPenilaian == nil {
		return fmt.Errorf("TahunPenilaian is nil")
	}

	if iroopd.NamaOPD == nil {
		return fmt.Errorf("NamaOPD is nil")
	}

	if iroopd.NomorUraianRisiko == nil {
		return fmt.Errorf("NomorUraianRisiko is nil")
	}

	if iroopd.KategoriRisiko == nil {
		return fmt.Errorf("KategoriRisiko is nil")
	}

	yearSuffix := fmt.Sprintf("%02d", iroopd.TahunPenilaian.Year()%100)
	iterStr := fmt.Sprintf("%03d", *iroopd.NomorUraianRisiko)

	kodeRisiko := fmt.Sprintf("ROO.%s.%s.%s.%s", yearSuffix, kodeOpd, kodeKategoriRisiko, iterStr)
	iroopd.KodeRisiko = &kodeRisiko
	return nil
}
