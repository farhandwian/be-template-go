package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenilaianKegiatanPengendalianSaveReq struct {
	PenilaianKegiatanPengendalian model.PenilaianKegiatanPengendalian
}

type PenilaianKegiatanPengendalianSaveRes struct {
	ID string
}

type PenilaianKegiatanPengendalianSave = core.ActionHandler[PenilaianKegiatanPengendalianSaveReq, PenilaianKegiatanPengendalianSaveRes]

func ImplPenilaianKegiatanPengendalianSave(db *gorm.DB) PenilaianKegiatanPengendalianSave {
	return func(ctx context.Context, req PenilaianKegiatanPengendalianSaveReq) (*PenilaianKegiatanPengendalianSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.PenilaianKegiatanPengendalian).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenilaianKegiatanPengendalianSaveRes{ID: *req.PenilaianKegiatanPengendalian.ID}, nil
	}
}
