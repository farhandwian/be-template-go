// File: gateway/gateway_asset.go

package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenetepanKonteksRisikoStrategisRenstraOPDDeleteReq struct {
	ID string
}

type PenetepanKonteksRisikoStrategisRenstraOPDDeleteRes struct{}

type PenetepanKonteksRisikoStrategisRenstraOPDDelete = core.ActionHandler[PenetepanKonteksRisikoStrategisRenstraOPDDeleteReq, PenetepanKonteksRisikoStrategisRenstraOPDDeleteRes]

func ImplPenetepanKonteksRisikoStrategisRenstraOPDDelete(db *gorm.DB) PenetepanKonteksRisikoStrategisRenstraOPDDelete {
	return func(ctx context.Context, req PenetepanKonteksRisikoStrategisRenstraOPDDeleteReq) (*PenetepanKonteksRisikoStrategisRenstraOPDDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.PenetapanKonteksRisikoStrategisRenstraOPD{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenetepanKonteksRisikoStrategisRenstraOPDDeleteRes{}, nil
	}
}
