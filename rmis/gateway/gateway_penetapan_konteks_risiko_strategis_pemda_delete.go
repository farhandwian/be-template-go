// File: gateway/gateway_asset.go

package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenetepanKonteksRisikoStrategisDeletePemdaReq struct {
	ID string
}

type PenetepanKonteksRisikoStrategisDeletePemdaRes struct{}

type PenetepanKonteksRisikoStrategisDeletePemda = core.ActionHandler[PenetepanKonteksRisikoStrategisDeletePemdaReq, PenetepanKonteksRisikoStrategisDeletePemdaRes]

func ImplPenetepanKonteksRisikoStrategisPemda(db *gorm.DB) PenetepanKonteksRisikoStrategisDeletePemda {
	return func(ctx context.Context, req PenetepanKonteksRisikoStrategisDeletePemdaReq) (*PenetepanKonteksRisikoStrategisDeletePemdaRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.RekapitulasiHasilKuesioner{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenetepanKonteksRisikoStrategisDeletePemdaRes{}, nil
	}
}
