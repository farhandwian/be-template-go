// File: gateway/gateway_asset.go

package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenilaianKegiatanPengendalianDeleteReq struct {
	ID string
}

type PenilaianKegiatanPengendalianDeleteRes struct{}

type PenilaianKegiatanPengendalianDelete = core.ActionHandler[PenilaianKegiatanPengendalianDeleteReq, PenilaianKegiatanPengendalianDeleteRes]

func ImplPenilaianKegiatanPengendalianDelete(db *gorm.DB) PenilaianKegiatanPengendalianDelete {
	return func(ctx context.Context, req PenilaianKegiatanPengendalianDeleteReq) (*PenilaianKegiatanPengendalianDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.PenilaianKegiatanPengendalian{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenilaianKegiatanPengendalianDeleteRes{}, nil
	}
}
