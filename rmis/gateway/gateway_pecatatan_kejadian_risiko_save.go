package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PencatatanKejadianRisikoSaveReq struct {
	PencatatanKejadianRisiko model.PencatatanKejadianRisiko
}

type PencatatanKejadianRisikoSaveRes struct {
	ID string
}

type PencatatanKejadianRisikoSave = core.ActionHandler[PencatatanKejadianRisikoSaveReq, PencatatanKejadianRisikoSaveRes]

func ImplPencatatanKejadianRisikoSave(db *gorm.DB) PencatatanKejadianRisikoSave {
	return func(ctx context.Context, req PencatatanKejadianRisikoSaveReq) (*PencatatanKejadianRisikoSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.PencatatanKejadianRisiko).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PencatatanKejadianRisikoSaveRes{ID: *req.PencatatanKejadianRisiko.ID}, nil
	}
}
