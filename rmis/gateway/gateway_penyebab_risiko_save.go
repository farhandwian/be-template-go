package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenyebabRisikoSaveReq struct {
	PenyebabRisiko model.PenyebabRisiko
}

type PenyebabRisikoSaveRes struct {
	ID string
}

type PenyebabRisikoSave = core.ActionHandler[PenyebabRisikoSaveReq, PenyebabRisikoSaveRes]

func ImplPenyebabRisikoSave(db *gorm.DB) PenyebabRisikoSave {
	return func(ctx context.Context, req PenyebabRisikoSaveReq) (*PenyebabRisikoSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.PenyebabRisiko).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenyebabRisikoSaveRes{ID: *req.PenyebabRisiko.ID}, nil
	}
}
