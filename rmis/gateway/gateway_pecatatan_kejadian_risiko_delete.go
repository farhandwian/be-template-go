// File: gateway/gateway_asset.go

package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PencatatanKejadianRisikoDeleteReq struct {
	ID string
}

type PencatatanKejadianRisikoDeleteRes struct{}

type PencatatanKejadianRisikoDelete = core.ActionHandler[PencatatanKejadianRisikoDeleteReq, PencatatanKejadianRisikoDeleteRes]

func ImplPencatatanKejadianRisikoDelete(db *gorm.DB) PencatatanKejadianRisikoDelete {
	return func(ctx context.Context, req PencatatanKejadianRisikoDeleteReq) (*PencatatanKejadianRisikoDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.PencatatanKejadianRisiko{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PencatatanKejadianRisikoDeleteRes{}, nil
	}
}
