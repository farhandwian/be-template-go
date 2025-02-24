// File: gateway/gateway_asset.go

package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type HasilAnalisisRisikoDeleteReq struct {
	ID string
}

type HasilAnalisisRisikoDeleteRes struct{}

type HasilAnalisisRisikoDelete = core.ActionHandler[HasilAnalisisRisikoDeleteReq, HasilAnalisisRisikoDeleteRes]

func ImplHasilAnalisisRisikoDelete(db *gorm.DB) HasilAnalisisRisikoDelete {
	return func(ctx context.Context, req HasilAnalisisRisikoDeleteReq) (*HasilAnalisisRisikoDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.HasilAnalisisRisiko{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &HasilAnalisisRisikoDeleteRes{}, nil
	}
}
