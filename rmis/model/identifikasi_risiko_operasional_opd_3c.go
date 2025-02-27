package model

import (
	"fmt"
	"time"
)

// Form 3C

type IdentifikasiRisikoOperasionalOPD struct {
	ID                 *string    `json:"id"`
	NamaPemda          *string    `json:"nama_pemda"`
	OPDID              *string    `json:"opd_id"` // references opd
	TahunPenilaian     *time.Time `json:"tahun_penilaian"`
	Periode            *string    `json:"periode"`
	UrusanPemerintahan *string    `json:"urusan_pemerintahan"`
	// ProgramKegiatanSubkegiatanID *string    `json:"program_kegiatan_subkegiatan"` // reference to penetapan_konteks_risiko_operasional.id form 2c
	IndikatorKeluaran *string `json:"indikator_keluaran"`
	TahapRisiko       *string `json:"tahap_risiko"`
	KategoriRisikoID  *string `json:"kategori_risiko_id"` // references kategori_resiko.id
	NomorUraianRisiko *int    `json:"nomor_uraian_risiko"`
	UraianRisiko      *string `json:"uraian_risiko"`
	KodeRisiko        *string `json:"kode_risiko"`
	PemilikRisiko     *string `json:"pemilik_risiko"`
	// UraianSebab                *string    `json:"uraian_sebab"` // references rca.akar_penyebab
	// SumberSebab                *string    `json:"sumber_sebab"` // references rca.jenis_penyebab
	Controllable *string   `json:"controllable"` // buat enum
	UraianDampak *string   `json:"uraian_dampak"`
	PihakDampak  *string   `json:"pihak_dampak"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type IdentifikasiRisikoOperasionalOPDGetRes struct {
	IdentifikasiRisikoOperasionalOPD IdentifikasiRisikoOperasionalOPD `json:"identifikasi_risiko_operasional_opd"`
	OPD                              OPD                              `json:"opd"`
}

func (iroopd *IdentifikasiRisikoOperasionalOPD) GenerateKodeRisiko(kodeKategoriRisiko string, kodeOpd string) error {
	if iroopd.TahunPenilaian == nil {
		return fmt.Errorf("TahunPenilaian is nil")
	}

	if iroopd.OPDID == nil {
		return fmt.Errorf("NamaOPD is nil")
	}

	if iroopd.NomorUraianRisiko == nil {
		return fmt.Errorf("NomorUraianRisiko is nil")
	}

	if iroopd.KategoriRisikoID == nil {
		return fmt.Errorf("KategoriRisikoID is nil")
	}

	yearSuffix := fmt.Sprintf("%02d", iroopd.TahunPenilaian.Year()%100)
	iterStr := fmt.Sprintf("%03d", *iroopd.NomorUraianRisiko)

	kodeRisiko := fmt.Sprintf("ROO.%s.%s.%s.%s", yearSuffix, kodeOpd, kodeKategoriRisiko, iterStr)
	iroopd.KodeRisiko = &kodeRisiko
	return nil
}
