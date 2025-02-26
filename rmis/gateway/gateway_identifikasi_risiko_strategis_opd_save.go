package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type IdentifikasiRisikoStrategisOPDSaveReq struct {
	IdentifikasiRisikoStrategisOPD model.IdentifikasiRisikoStrategisOPD
}

type IdentifikasiRisikoStrategisOPDSaveRes struct {
	ID string
}

type IdentifikasiRisikoStrategisOPDSave = core.ActionHandler[IdentifikasiRisikoStrategisOPDSaveReq, IdentifikasiRisikoStrategisOPDSaveRes]

func ImplIdentifikasiRisikoStrategisOPDSave(db *gorm.DB) IdentifikasiRisikoStrategisOPDSave {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisOPDSaveReq) (*IdentifikasiRisikoStrategisOPDSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.IdentifikasiRisikoStrategisOPD).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &IdentifikasiRisikoStrategisOPDSaveRes{ID: *req.IdentifikasiRisikoStrategisOPD.ID}, nil
	}
}
