package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PengkomunikasianPengendalianGetByIDReq struct {
	ID string
}

type PengkomunikasianPengendalianGetByIDRes struct {
	PengkomunikasianPengendalian model.PengkomunikasianPengendalian
}

type PengkomunikasianPengendalianGetByID = core.ActionHandler[PengkomunikasianPengendalianGetByIDReq, PengkomunikasianPengendalianGetByIDRes]

func ImplPengkomunikasianPengendalianGetByID(db *gorm.DB) PengkomunikasianPengendalianGetByID {
	return func(ctx context.Context, req PengkomunikasianPengendalianGetByIDReq) (*PengkomunikasianPengendalianGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var PengkomunikasianPengendalian model.PengkomunikasianPengendalian
		if err := query.
			Joins("LEFT JOIN penilaian_risikos ON pengkomunikasian_pengendalians.penilaian_risiko_id = penilaian_risikos.id").
			Joins("LEFT JOIN daftar_risiko_prioritas ON penilaian_risikos.daftar_risiko_prioritas_id = daftar_risiko_prioritas.id").
			Joins("LEFT JOIN penetapan_konteks_risiko_strategis_pemdas ON daftar_risiko_prioritas.penetapan_konteks_risiko_strategis_pemda_id = penetapan_konteks_risiko_strategis_pemdas.id").
			Where("pengkomunikasian_pengendalians.id =?", req.ID).
			First(&PengkomunikasianPengendalian).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("PengkomunikasianPengendalian id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &PengkomunikasianPengendalianGetByIDRes{PengkomunikasianPengendalian: PengkomunikasianPengendalian}, nil
	}
}
