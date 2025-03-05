package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenilaianRisikoGetByIDReq struct {
	ID string
}

type PenilaianRisikoGetByIDRes struct {
	PenilaianRisiko model.PenilaianRisiko
}

type PenilaianRisikoGetByID = core.ActionHandler[PenilaianRisikoGetByIDReq, PenilaianRisikoGetByIDRes]

func ImplPenilaianRisikoGetByID(db *gorm.DB) PenilaianRisikoGetByID {
	return func(ctx context.Context, req PenilaianRisikoGetByIDReq) (*PenilaianRisikoGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var PenilaianRisiko model.PenilaianRisiko
		if err := query.
			Joins("LEFT JOIN daftar_risiko_prioritas ON penilaian_risikos.daftar_risiko_prioritas_id = daftar_risiko_prioritas.id").
			Joins("LEFT JOIN penetapan_konteks_risiko_strategis_pemdas ON daftar_risiko_prioritas.penetapan_konteks_risiko_strategis_pemda_id = penetapan_konteks_risiko_strategis_pemdas.id").
			Where("penilaian_risikos.id =?", req.ID).
			First(&PenilaianRisiko).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("PenilaianRisiko id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &PenilaianRisikoGetByIDRes{PenilaianRisiko: PenilaianRisiko}, nil
	}
}
