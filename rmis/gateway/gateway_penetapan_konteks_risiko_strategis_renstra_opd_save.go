package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenetapanKonteksRisikoStrategisRenstraOPDSaveReq struct {
	PenetepanKonteksRisikoStrategisRenstraOPD model.PenetapanKonteksRisikoStrategisRenstraOPD
}

type PenetapanKonteksRisikoStrategisRenstraOPDSaveRes struct {
	ID string
}

type PenetepanKonteksRisikoStrategisRenstraOPDSave = core.ActionHandler[PenetapanKonteksRisikoStrategisRenstraOPDSaveReq, PenetapanKonteksRisikoStrategisRenstraOPDSaveRes]

func ImplPenetepanKonteksRisikoStrategisRenstraOPDSave(db *gorm.DB) PenetepanKonteksRisikoStrategisRenstraOPDSave {
	return func(ctx context.Context, req PenetapanKonteksRisikoStrategisRenstraOPDSaveReq) (*PenetapanKonteksRisikoStrategisRenstraOPDSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.PenetepanKonteksRisikoStrategisRenstraOPD).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenetapanKonteksRisikoStrategisRenstraOPDSaveRes{ID: *req.PenetepanKonteksRisikoStrategisRenstraOPD.ID}, nil
	}
}
