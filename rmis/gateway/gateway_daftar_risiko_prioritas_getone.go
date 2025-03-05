package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type DaftarRisikoPrioritasGetByIDReq struct {
	ID string
}

type DaftarRisikoPrioritasGetByIDRes struct {
	DaftarRisikoPrioritas model.DaftarRisikoPrioritas
}

type DaftarRisikoPrioritasGetByID = core.ActionHandler[DaftarRisikoPrioritasGetByIDReq, DaftarRisikoPrioritasGetByIDRes]

func ImplDaftarRisikoPrioritasGetByID(db *gorm.DB) DaftarRisikoPrioritasGetByID {
	return func(ctx context.Context, req DaftarRisikoPrioritasGetByIDReq) (*DaftarRisikoPrioritasGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var DaftarRisikoPrioritas model.DaftarRisikoPrioritas
		if err := query.
			Select(`daftar_risiko_prioritas.*, 
				penetapan_konteks_risiko_strategis_pemdas.nama_pemda AS nama_pemda,
				penetapan_konteks_risiko_strategis_pemdas.tahun_penilaian AS tahun,
				penetapan_konteks_risiko_strategis_pemdas.periode AS periode,
				penetapan_konteks_risiko_strategis_pemdas.penetapan_tujuan AS tujuan,
				penetapan_konteks_risiko_strategis_pemdas.urusan_pemerintahan AS urusan_pemerintahan,
				penetapan_konteks_risiko_strategis_pemdas.penetapan_tujuan AS penetapan_konteks,
				hasil_analisis_risikos.skala_risiko AS skala_risiko
			`).
			Joins("LEFT JOIN penetapan_konteks_risiko_strategis_pemdas ON daftar_risiko_prioritas.penetapan_konteks_risiko_strategis_pemda_id = penetapan_konteks_risiko_strategis_pemdas.id").
			Joins("LEFT JOIN indeks_peringkat_prioritas ON daftar_risiko_prioritas.indeks_peringkat_prioritas_id = indeks_peringkat_prioritas.id").
			Joins("LEFT JOIN hasil_analisis_risikos ON daftar_risiko_prioritas.hasil_analisis_risiko_id = hasil_analisis_risikos.id").
			Where("daftar_risiko_prioritas.id =?", req.ID).
			First(&DaftarRisikoPrioritas).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("DaftarRisikoPrioritas id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &DaftarRisikoPrioritasGetByIDRes{DaftarRisikoPrioritas: DaftarRisikoPrioritas}, nil
	}
}
