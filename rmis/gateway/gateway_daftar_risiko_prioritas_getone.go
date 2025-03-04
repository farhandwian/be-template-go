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
				penetapan_konteks_risiko_strategis_pemdas.tahun AS tahun,
				penetapan_konteks_risiko_strategis_pemdas.penetapan_tujuan AS tujuan,
				penetapan_konteks_risiko_strategis_pemdas.urusan_pemerintah AS urusan_pemerintah,
				penetapan_konteks_risiko_strategis_pemdas.penetapan_tujuan AS penetapan_tujuan,
			`).
			Joins("LEFT JOIN penetapan_konteks_risiko_strategis_pemdas ON daftar_risiko_prioritas.penetapan_konteks_risiko_strategis_pemda_id = penetapan_konteks_risiko_strategis_pemdas.id").
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
