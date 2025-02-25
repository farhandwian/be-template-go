package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PengkomunikasianPengendalianSaveReq struct {
	PengkomunikasianPengendalian model.PengkomunikasianPengendalian
}

type PengkomunikasianPengendalianSaveRes struct {
	ID string
}

type PengkomunikasianPengendalianSave = core.ActionHandler[PengkomunikasianPengendalianSaveReq, PengkomunikasianPengendalianSaveRes]

func ImplPengkomunikasianPengendalianSave(db *gorm.DB) PengkomunikasianPengendalianSave {
	return func(ctx context.Context, req PengkomunikasianPengendalianSaveReq) (*PengkomunikasianPengendalianSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.PengkomunikasianPengendalian).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PengkomunikasianPengendalianSaveRes{ID: *req.PengkomunikasianPengendalian.ID}, nil
	}
}
