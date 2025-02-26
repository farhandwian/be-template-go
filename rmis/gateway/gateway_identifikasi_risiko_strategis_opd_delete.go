// File: gateway/gateway_asset.go

package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type IdentifikasiRisikoStrategisOPDDeleteReq struct {
	ID string
}

type IdentifikasiRisikoStrategisOPDDeleteRes struct{}

type IdentifikasiRisikoStrategisOPDDelete = core.ActionHandler[IdentifikasiRisikoStrategisOPDDeleteReq, IdentifikasiRisikoStrategisOPDDeleteRes]

func ImplIdentifikasiRisikoStrategisOPDDelete(db *gorm.DB) IdentifikasiRisikoStrategisOPDDelete {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisOPDDeleteReq) (*IdentifikasiRisikoStrategisOPDDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.IdentifikasiRisikoStrategisOPD{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &IdentifikasiRisikoStrategisOPDDeleteRes{}, nil
	}
}
