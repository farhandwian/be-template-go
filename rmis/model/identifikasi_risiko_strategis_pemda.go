package model

import (
	"fmt"
	sharedModel "shared/model"
	"time"
)

// Form 3A

type IdentifikasiRisikoStrategisPemda struct {
	ID                                     *string `json:"id" gorm:"primaryKey"`
	PenetapanKonteksRisikoStrategisPemdaID *string `json:"-" gorm:"type:VARCHAR(255)"`
	KategoriRisikoID                       *string `json:"-" gorm:"type:VARCHAR(255)"` // references kategori_resiko.id
	// RcaID                                  *string `json:"-" gorm:"type:VARCHAR(255)"`
	KategoriRisikoNama *string `json:"kategori_resiko_nama" gorm:"-"`
	UraianRisiko       *string `json:"uraian_resiko"`
	NomorUraian        *int    `json:"nomor_uraian"`
	KodeRisiko         *string `json:"kode_resiko"`
	PemilikRisiko      *string `json:"pemilik_resiko"`
	// Rca           *Rca               `json:"-" gorm:"foreignKey:RcaID"`
	UraianSebab  *string `json:"uraian_sebab"` // references rca.akar_penyebab
	SumberSebab  *string `json:"sumber_sebab"` // references rca.jenis_penyebab
	Controllable *string `json:"controllable"` // could be boolean if desired
	UraianDampak *string `json:"uraian_dampak"`
	PihakDampak  *string `json:"pihak_dampak"`

	NamaPemda        *string            `json:"nama_pemda"`
	Tahun            *time.Time         `json:"tahun"`
	Periode          *string            `json:"periode"`
	PenetapanKonteks *string            `json:"penetapan_konteks"`
	UrusanPemerintah *string            `json:"urusan_pemerintah"`
	Status           sharedModel.Status `json:"status"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
}

func (irspd *IdentifikasiRisikoStrategisPemda) GenerateKodeRisiko(tahun time.Time, kategori_risiko string) error {
	fmt.Println("TEST BANG: GenerateKodeRisiko")

	if irspd.NomorUraian == nil {
		defaultNum := 1
		irspd.NomorUraian = &defaultNum
	} else {
		*irspd.NomorUraian = *irspd.NomorUraian + 1
	}
	yearSuffix := fmt.Sprintf("%02d", tahun.Year()%100)
	iterStr := fmt.Sprintf("%03d", *irspd.NomorUraian)

	kodeRisiko := fmt.Sprintf("RSP.%s.%s.%s", yearSuffix, kategori_risiko, iterStr)
	irspd.KodeRisiko = &kodeRisiko

	return nil
}
