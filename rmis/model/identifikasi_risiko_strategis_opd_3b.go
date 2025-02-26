package model

import (
	"fmt"
	"time"
)

// form 3b

type IdentifikasiRisikoStrategisOPD struct {
	ID                 *string    `json:"id"`
	NamaPemda          *string    `json:"nama_pemda"`
	OPDID              *string    `json:"opd_id"` // references opd
	TahunPenilaian     *time.Time `json:"tahun_penilaian"`
	Periode            *string    `json:"periode"`
	UrusanPemerintahan *string    `json:"urusan_pemerintahan"`
	// PenetapanKonteksRisikoStrategisRenstraOPDID *string `json:"penetapan_konteks_risiko_strategis_renstra_opd_id"` // references penetapan_konteks_risiko_strategis_renstra_opd.id / form 2b, references to form 2b or db ?
	IndikatorKinerja  *string `json:"indikator_kinerja"`
	KategoriRisikoID  *string `json:"-"` // references kategori_resiko.id
	NomorUraianRisiko *int    `json:"nomor_uraian_risiko"`
	UraianRisiko      *string `json:"uraian_risiko"`
	KodeRisiko        *string `json:"kode_risiko"`
	PemilikRisiko     *string `json:"pemilik_risiko"`
	// UraianSebab                                 *string   `json:"uraian_sebab"` // references rca.akar_penyebab
	// SumberSebab                                 *string   `json:"sumber_sebab"` // references rca.jenis_penyebab
	Controllable *string   `json:"controllable"`
	UraianDampak *string   `json:"uraian_dampak"`
	PihakDampak  *string   `json:"pihak_dampak"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type IdentifikasiRisikoStrategisOPDGetRes struct {
	IdentifikasiRisikoStrategisOPD IdentifikasiRisikoStrategisOPD `json:"identifikasi_risiko_strategis_opd"`
	OPD                            OPD                            `json:"opd"`
}

func (irsopd *IdentifikasiRisikoStrategisOPD) GenerateKodeRisiko(kodeKategoriRisiko string, kodeOpd string) error {
	fmt.Println("GenerateKodeRisiko")
	if irsopd.TahunPenilaian == nil {
		fmt.Println("TahunPenilaian is nil")
		return fmt.Errorf("TahunPenilaian is nil")
	}

	if irsopd.OPDID == nil {
		fmt.Println("OPDID is nil")
		return fmt.Errorf("NamaOPD is nil")
	}

	if irsopd.NomorUraianRisiko == nil {
		fmt.Println("NomorUraianRisiko is nil")
		return fmt.Errorf("NomorUraianRisiko is nil")
	}

	if irsopd.KategoriRisikoID == nil {
		fmt.Println("KategoriRisiko id is nil")
		return fmt.Errorf("KategoriRisiko id is nil")
	}

	fmt.Println("Kode Kategori Risiko: ", kodeKategoriRisiko)
	fmt.Println("Kode OPD: ", kodeOpd)
	fmt.Println("Tahun Penilaian: ", irsopd.TahunPenilaian)

	yearSuffix := fmt.Sprintf("%02d", irsopd.TahunPenilaian.Year()%100)
	iterStr := fmt.Sprintf("%03d", *irsopd.NomorUraianRisiko)

	kodeRisiko := fmt.Sprintf("RSO.%s.%s.%s.%s", yearSuffix, kodeOpd, kodeKategoriRisiko, iterStr)
	irsopd.KodeRisiko = &kodeRisiko
	return nil
}
