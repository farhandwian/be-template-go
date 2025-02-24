package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type HasilAnalisisRisikoSaveReq struct {
	HasilAnalisisRisiko model.HasilAnalisisRisiko
}

type HasilAnalisisRisikoSaveRes struct {
	ID string
}

type HasilAnalisisRisikoSave = core.ActionHandler[HasilAnalisisRisikoSaveReq, HasilAnalisisRisikoSaveRes]

func ImplHasilAnalisisRisikoSave(db *gorm.DB) HasilAnalisisRisikoSave {
	return func(ctx context.Context, req HasilAnalisisRisikoSaveReq) (*HasilAnalisisRisikoSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.HasilAnalisisRisiko).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &HasilAnalisisRisikoSaveRes{ID: *req.HasilAnalisisRisiko.ID}, nil
	}
}
