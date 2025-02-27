package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type IdentifikasiRisikoOperasionalOPDSaveReq struct {
	IdentifikasiRisikoOperasionalOPD model.IdentifikasiRisikoOperasionalOPD
}

type IdentifikasiRisikoOperasionalOPDSaveRes struct {
	ID string
}

type IdentifikasiRisikoOperasionalOPDSave = core.ActionHandler[IdentifikasiRisikoOperasionalOPDSaveReq, IdentifikasiRisikoOperasionalOPDSaveRes]

func ImplIdentifikasiRisikoOperasionalOPDSave(db *gorm.DB) IdentifikasiRisikoOperasionalOPDSave {
	return func(ctx context.Context, req IdentifikasiRisikoOperasionalOPDSaveReq) (*IdentifikasiRisikoOperasionalOPDSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.IdentifikasiRisikoOperasionalOPD).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &IdentifikasiRisikoOperasionalOPDSaveRes{ID: *req.IdentifikasiRisikoOperasionalOPD.ID}, nil
	}
}
