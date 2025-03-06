package model

import (
	"fmt"
	sharedModel "shared/model"
	"time"
)

// form 3b

type IdentifikasiRisikoStrategisOPD struct {
	ID                                       *string `json:"id" gorm:"primaryKey"`
	PenetapanKonteksRisikoStrategisRenstraID *string `json:"-" gorm:"type:VARCHAR(255)"`
	KategoriRisikoID                         *string `json:"-" gorm:"type:VARCHAR(255)"` // references kategori_resiko.id
	RcaID                                    *string `json:"-" gorm:"type:VARCHAR(255)"`

	UraianRisiko  *string `json:"uraian_resiko"`
	NomorUraian   *int    `json:"nomor_uraian"`
	KodeRisiko    *string `json:"kode_resiko"`
	PemilikRisiko *string `json:"pemilik_resiko"`

	UraianSebab  *string `json:"uraian_sebab"` // references rca.akar_penyebab
	SumberSebab  *string `json:"sumber_sebab"` // references rca.jenis_penyebab
	Controllable *string `json:"controllable"` // could be boolean if desired
	UraianDampak *string `json:"uraian_dampak"`
	PihakDampak  *string `json:"pihak_dampak"`

	Status    sharedModel.Status `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type IdentifikasiRisikoStrategisOPDResponse struct {
	ID                                       *string `json:"id" gorm:"primaryKey"`
	PenetapanKonteksRisikoStrategisRenstraID *string `json:"-" gorm:"type:VARCHAR(255)"`
	KategoriRisikoID                         *string `json:"-" gorm:"type:VARCHAR(255)"` // references kategori_resiko.id
	RcaID                                    *string `json:"-" gorm:"type:VARCHAR(255)"`

	UraianRisiko  *string `json:"uraian_resiko"`
	NomorUraian   *int    `json:"nomor_uraian"`
	KodeRisiko    *string `json:"kode_resiko"`
	PemilikRisiko *string `json:"pemilik_resiko"`

	UraianSebab  *string `json:"uraian_sebab"` // references rca.akar_penyebab
	SumberSebab  *string `json:"sumber_sebab"` // references rca.jenis_penyebab
	Controllable *string `json:"controllable"` // could be boolean if desired
	UraianDampak *string `json:"uraian_dampak"`
	PihakDampak  *string `json:"pihak_dampak"`

	NamaPemda          *string    `json:"nama_pemda"`
	Tahun              *time.Time `json:"tahun"`
	Periode            *string    `json:"periode"`
	PenetapanKonteks   *string    `json:"penetapan_konteks"`
	UrusanPemerintahan *string    `json:"urusan_pemerintahan"`

	Status    sharedModel.Status `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type IdentifikasiRisikoStrategisOPDGetRes struct {
	IdentifikasiRisikoStrategisOPD IdentifikasiRisikoStrategisOPD `json:"identifikasi_risiko_strategis_opd"`
	OPD                            OPD                            `json:"opd"`
}

func (irsopd *IdentifikasiRisikoStrategisOPD) GenerateKodeRisiko(tahun time.Time, kodeKategoriRisiko string, kodeOpd string) error {

	yearSuffix := fmt.Sprintf("%02d", tahun.Year()%100)
	fmt.Println("TESTING", yearSuffix)
	// iterStr := fmt.Sprintf("%03d", *irsopd.NomorUraian)
	nomor := 4
	iterStr := fmt.Sprintf("%03d", nomor)

	fmt.Println("TEST")
	kodeRisiko := fmt.Sprintf("RSO.%s.%s.%s.%s", yearSuffix, kodeOpd, kodeKategoriRisiko, iterStr)
	irsopd.KodeRisiko = &kodeRisiko
	return nil
}
