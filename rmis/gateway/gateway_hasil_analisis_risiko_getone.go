package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type HasilAnalisisRisikoGetByIDReq struct {
	ID string
}

type HasilAnalisisRisikoGetByIDRes struct {
	HasilAnalisisRisiko model.HasilAnalisisRisiko
}

type HasilAnalisisRisikoGetByID = core.ActionHandler[HasilAnalisisRisikoGetByIDReq, HasilAnalisisRisikoGetByIDRes]

func ImplHasilAnalisisRisikoGetByID(db *gorm.DB) HasilAnalisisRisikoGetByID {
	return func(ctx context.Context, req HasilAnalisisRisikoGetByIDReq) (*HasilAnalisisRisikoGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var HasilAnalisisRisiko model.HasilAnalisisRisiko
		if err := query.
			Select(`hasil_analisis_risikos.*, 
			    penetapan_konteks_risiko_strategis_pemdas.nama_pemda AS nama_pemda,
                penetapan_konteks_risiko_strategis_pemdas.tahun AS tahun,
				penetapan_konteks_risiko_strategis_pemdas.penetapan_tujuan AS tujuan,
                penetapan_konteks_risiko_strategis_pemdas.urusan_pemerintah AS urusan_pemerintah,
				penetapan_konteks_risiko_strategis_pemdas.penetapan_tujuan AS penetapan_tujuan,
			`).
			Joins("LEFT JOIN penetapan_konteks_risiko_strategis_pemdas ON hasil_analisis_risikos.penetapan_konteks_risiko_strategis_pemda_id = penetapan_konteks_risiko_strategis_pemdas.id").
			Joins("LEFT JOIN identifikasi_risiko_strategis_pemdas ON hasil_analisis_risikos.identifikasi_risiko_strategis_pemda_id = identifikasi_risiko_strategis_pemdas.id").
			Where("hasil_analisis_risikos.id =?", req.ID).
			First(&HasilAnalisisRisiko).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("HasilAnalisisRisiko id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &HasilAnalisisRisikoGetByIDRes{HasilAnalisisRisiko: HasilAnalisisRisiko}, nil
	}
}
