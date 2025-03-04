package model

import (
	"fmt"
	sharedModel "shared/model"
	"time"
)

// Form 3A

type IdentifikasiRisikoStrategisPemda struct {
	ID                                     *string                               `json:"id"`
	PenetapanKonteksRisikoStrategisPemdaID *string                               `json:"-"`
	PenetapanKonteksRisikoStrategisPemda   *PenetapanKonteksRisikoStrategisPemda `json:"penetapan_konteks_risiko_strategis_pemda" gorm:"foreignKey:PenetapanKonteksRisikoStrategisPemdaID;"`
	TujuanStrategis                        *string                               `json:"tujuan_strategis"`
	IndikatorKinerja                       *string                               `json:"indikator_kinerja"`
	KategoriRisikoID                       *string                               `json:"-"` // references kategori_resiko.id
	KategoriRisiko                         *KategoriRisiko                       `json:"kategori_risiko" gorm:"foreignKey:KategoriRisikoID"`
	UraianRisiko                           *string                               `json:"uraian_resiko"`
	NomorUraian                            *int                                  `json:"nomor_uraian"`
	KodeRisiko                             *string                               `json:"kode_resiko"`
	PemilikRisiko                          *string                               `json:"pemilik_resiko"`
	RcaID                                  *string                               `json:"-"`
	Rca                                    *Rca                                  `json:"-" gorm:"foreignKey:RcaID"`
	UraianSebab                            *string                               `json:"uraian_sebab"` // references rca.akar_penyebab
	SumberSebab                            *string                               `json:"sumber_sebab"` // references rca.jenis_penyebab
	Controllable                           *string                               `json:"controllable"` // could be boolean if desired
	UraianDampak                           *string                               `json:"uraian_dampak"`
	PihakDampak                            *string                               `json:"pihak_dampak"`
	Status                                 sharedModel.Status                    `json:"status"`
	CreatedAt                              time.Time                             `json:"created_at"`
	UpdatedAt                              time.Time                             `json:"updated_at"`
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
