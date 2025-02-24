package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenilaianKegiatanPengendalianGetByIDReq struct {
	ID string
}

type PenilaianKegiatanPengendalianGetByIDRes struct {
	PenilaianKegiatanPengendalian model.PenilaianKegiatanPengendalian
}

type PenilaianKegiatanPengendalianGetByID = core.ActionHandler[PenilaianKegiatanPengendalianGetByIDReq, PenilaianKegiatanPengendalianGetByIDRes]

func ImplPenilaianKegiatanPengendalianGetByID(db *gorm.DB) PenilaianKegiatanPengendalianGetByID {
	return func(ctx context.Context, req PenilaianKegiatanPengendalianGetByIDReq) (*PenilaianKegiatanPengendalianGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var PenilaianKegiatanPengendalian model.PenilaianKegiatanPengendalian
		if err := query.First(&PenilaianKegiatanPengendalian, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("PenilaianKegiatanPengendalian id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &PenilaianKegiatanPengendalianGetByIDRes{PenilaianKegiatanPengendalian: PenilaianKegiatanPengendalian}, nil
	}
}
