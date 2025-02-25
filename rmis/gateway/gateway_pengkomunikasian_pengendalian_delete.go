// File: gateway/gateway_asset.go

package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PengkomunikasianPengendalianDeleteReq struct {
	ID string
}

type PengkomunikasianPengendalianDeleteRes struct{}

type PengkomunikasianPengendalianDelete = core.ActionHandler[PengkomunikasianPengendalianDeleteReq, PengkomunikasianPengendalianDeleteRes]

func ImplPengkomunikasianPengendalianDelete(db *gorm.DB) PengkomunikasianPengendalianDelete {
	return func(ctx context.Context, req PengkomunikasianPengendalianDeleteReq) (*PengkomunikasianPengendalianDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.PengkomunikasianPengendalian{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PengkomunikasianPengendalianDeleteRes{}, nil
	}
}
