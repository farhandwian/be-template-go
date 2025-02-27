package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenetapanKonteksRisikoOperasionalSaveReq struct {
	PenetepanKonteksRisikoOperasional model.PenetapanKonteksRisikoOperasional
}

type PenetapanKonteksRisikoOperasionalSaveRes struct {
	ID string
}

type PenetepanKonteksRisikoOperasionalSave = core.ActionHandler[PenetapanKonteksRisikoOperasionalSaveReq, PenetapanKonteksRisikoOperasionalSaveRes]

func ImplPenetepanKonteksRisikoOperasionalSave(db *gorm.DB) PenetepanKonteksRisikoOperasionalSave {
	return func(ctx context.Context, req PenetapanKonteksRisikoOperasionalSaveReq) (*PenetapanKonteksRisikoOperasionalSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.PenetepanKonteksRisikoOperasional).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenetapanKonteksRisikoOperasionalSaveRes{ID: *req.PenetepanKonteksRisikoOperasional.ID}, nil
	}
}
