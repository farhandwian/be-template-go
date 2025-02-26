// File: gateway/gateway_asset.go

package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenetepanKonteksRisikoOperasionalDeleteReq struct {
	ID string
}

type PenetepanKonteksRisikoOperasionalDeleteRes struct{}

type PenetepanKonteksRisikoOperasionalDelete = core.ActionHandler[PenetepanKonteksRisikoOperasionalDeleteReq, PenetepanKonteksRisikoOperasionalDeleteRes]

func ImplPenetepanKonteksRisikoOperasionalDelete(db *gorm.DB) PenetepanKonteksRisikoOperasionalDelete {
	return func(ctx context.Context, req PenetepanKonteksRisikoOperasionalDeleteReq) (*PenetepanKonteksRisikoOperasionalDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.PenetapanKonteksRisikoOperasional{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenetepanKonteksRisikoOperasionalDeleteRes{}, nil
	}
}
