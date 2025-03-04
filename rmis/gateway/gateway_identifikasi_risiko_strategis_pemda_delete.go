// File: gateway/gateway_asset.go

package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type IdentifikasiRisikoStrategisPemdaDeleteReq struct {
	ID string
}

type IdentifikasiRisikoStrategisPemdaDeleteRes struct{}

type IdentifikasiRisikoStrategisPemdaDelete = core.ActionHandler[IdentifikasiRisikoStrategisPemdaDeleteReq, IdentifikasiRisikoStrategisPemdaDeleteRes]

func ImplIdentifikasiRisikoStrategisPemdaDelete(db *gorm.DB) IdentifikasiRisikoStrategisPemdaDelete {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisPemdaDeleteReq) (*IdentifikasiRisikoStrategisPemdaDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.IdentifikasiRisikoStrategisPemda{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &IdentifikasiRisikoStrategisPemdaDeleteRes{}, nil
	}
}
