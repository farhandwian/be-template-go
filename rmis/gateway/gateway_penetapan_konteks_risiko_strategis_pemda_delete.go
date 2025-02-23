// File: gateway/gateway_asset.go

package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenetepanKonteksRisikoStrategisPemdaDeleteReq struct {
	ID string
}

type PenetepanKonteksRisikoStrategisPemdaDeleteRes struct{}

type PenetepanKonteksRisikoStrategisPemdaDelete = core.ActionHandler[PenetepanKonteksRisikoStrategisPemdaDeleteReq, PenetepanKonteksRisikoStrategisPemdaDeleteRes]

func ImplPenetepanKonteksRisikoStrategisPemdaDelete(db *gorm.DB) PenetepanKonteksRisikoStrategisPemdaDelete {
	return func(ctx context.Context, req PenetepanKonteksRisikoStrategisPemdaDeleteReq) (*PenetepanKonteksRisikoStrategisPemdaDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.RekapitulasiHasilKuesioner{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenetepanKonteksRisikoStrategisPemdaDeleteRes{}, nil
	}
}
