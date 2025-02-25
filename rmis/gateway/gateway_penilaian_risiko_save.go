package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenilaianRisikoSaveReq struct {
	PenilaianRisiko model.PenilaianRisiko
}

type PenilaianRisikoSaveRes struct {
	ID string
}

type PenilaianRisikoSave = core.ActionHandler[PenilaianRisikoSaveReq, PenilaianRisikoSaveRes]

func ImplPenilaianRisikoSave(db *gorm.DB) PenilaianRisikoSave {
	return func(ctx context.Context, req PenilaianRisikoSaveReq) (*PenilaianRisikoSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.PenilaianRisiko).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenilaianRisikoSaveRes{ID: *req.PenilaianRisiko.ID}, nil
	}
}
