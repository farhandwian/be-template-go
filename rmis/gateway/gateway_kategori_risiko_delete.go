// File: gateway/gateway_asset.go

package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type KategoriRisikoDeleteReq struct {
	ID string
}

type KategoriRisikoDeleteRes struct{}

type KategoriRisikoDelete = core.ActionHandler[KategoriRisikoDeleteReq, KategoriRisikoDeleteRes]

func ImplKategoriRisikoDelete(db *gorm.DB) KategoriRisikoDelete {
	return func(ctx context.Context, req KategoriRisikoDeleteReq) (*KategoriRisikoDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.KategoriRisiko{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &KategoriRisikoDeleteRes{}, nil
	}
}
