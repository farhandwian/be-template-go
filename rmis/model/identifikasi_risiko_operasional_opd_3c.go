package model

import (
	"fmt"
	sharedModel "shared/model"
	"time"
)

// Form 3C

type IdentifikasiRisikoOperasionalOPD struct {
	ID                                     *string `json:"id"`
	PenetapanKonteksRisikoOperasionalOpdID *string `json:"-" gorm:"type:VARCHAR(255)"`
	KategoriRisikoID                       *string `json:"-" gorm:"type:VARCHAR(255)"` // references kategori_resiko.id
	RcaID                                  *string `json:"-" gorm:"type:VARCHAR(255)"`

	TahapRisiko   *string `json:"tahap_risiko"`
	UraianRisiko  *string `json:"uraian_resiko"`
	NomorUraian   *int    `json:"nomor_uraian"`
	KodeRisiko    *string `json:"kode_resiko"`
	PemilikRisiko *string `json:"pemilik_resiko"`

	UraianSebab       *string `json:"uraian_sebab"` // references rca.akar_penyebab
	SumberSebab       *string `json:"sumber_sebab"` // references rca.jenis_penyebab
	Controllable      *string `json:"controllable"` // could be boolean if desired
	UraianDampak      *string `json:"uraian_dampak"`
	PihakDampak       *string `json:"pihak_dampak"`
	Kegiatan          *string `json:"kegiatan"`
	Indikatorkeluaran *string `json:"indikator_keluaran"`

	Status    sharedModel.Status `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type IdentifikasiRisikoOperasionalOPDGetRes struct {
	IdentifikasiRisikoOperasionalOPD IdentifikasiRisikoOperasionalOPD `json:"identifikasi_risiko_operasional_opd"`
	OPD                              OPD                              `json:"opd"`
}

func (iroopd *IdentifikasiRisikoOperasionalOPD) GenerateKodeRisiko(tahun time.Time, kodeKategoriRisiko string, kodeOpd string) error {

	yearSuffix := fmt.Sprintf("%02d", tahun.Year()%100)
	// iterStr := fmt.Sprintf("%03d", *irsopd.NomorUraian)
	nomor := 4
	iterStr := fmt.Sprintf("%03d", nomor)
	kodeRisiko := fmt.Sprintf("ROO.%s.%s.%s.%s", yearSuffix, kodeOpd, kodeKategoriRisiko, iterStr)
	iroopd.KodeRisiko = &kodeRisiko
	return nil
}
