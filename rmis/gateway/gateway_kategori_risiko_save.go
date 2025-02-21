package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type KategoriRisikoSaveReq struct {
	KategoriRisiko model.KategoriRisiko
}

type KategoriRisikoSaveRes struct {
	ID string
}

type KategoriRisikoSave = core.ActionHandler[KategoriRisikoSaveReq, KategoriRisikoSaveRes]

func ImplKategoriRisikoSave(db *gorm.DB) KategoriRisikoSave {
	return func(ctx context.Context, req KategoriRisikoSaveReq) (*KategoriRisikoSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.KategoriRisiko).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &KategoriRisikoSaveRes{ID: *req.KategoriRisiko.ID}, nil
	}
}
