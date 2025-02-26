// File: gateway/gateway_asset.go

package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type IdentifikasiRisikoOperasionalOPDDeleteReq struct {
	ID string
}

type IdentifikasiRisikoOperasionalOPDDeleteRes struct{}

type IdentifikasiRisikoOperasionalOPDDelete = core.ActionHandler[IdentifikasiRisikoOperasionalOPDDeleteReq, IdentifikasiRisikoOperasionalOPDDeleteRes]

func ImplIdentifikasiRisikoOperasionalOPDDelete(db *gorm.DB) IdentifikasiRisikoOperasionalOPDDelete {
	return func(ctx context.Context, req IdentifikasiRisikoOperasionalOPDDeleteReq) (*IdentifikasiRisikoOperasionalOPDDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.IdentifikasiRisikoOperasionalOPD{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &IdentifikasiRisikoOperasionalOPDDeleteRes{}, nil
	}
}
