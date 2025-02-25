// File: gateway/gateway_asset.go

package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenilaianRisikoDeleteReq struct {
	ID string
}

type PenilaianRisikoDeleteRes struct{}

type PenilaianRisikoDelete = core.ActionHandler[PenilaianRisikoDeleteReq, PenilaianRisikoDeleteRes]

func ImplPenilaianRisikoDelete(db *gorm.DB) PenilaianRisikoDelete {
	return func(ctx context.Context, req PenilaianRisikoDeleteReq) (*PenilaianRisikoDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.PenilaianRisiko{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenilaianRisikoDeleteRes{}, nil
	}
}
