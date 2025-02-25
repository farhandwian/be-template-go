// File: gateway/gateway_asset.go

package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type DaftarRisikoPrioritasDeleteReq struct {
	ID string
}

type DaftarRisikoPrioritasDeleteRes struct{}

type DaftarRisikoPrioritasDelete = core.ActionHandler[DaftarRisikoPrioritasDeleteReq, DaftarRisikoPrioritasDeleteRes]

func ImplDaftarRisikoPrioritasDelete(db *gorm.DB) DaftarRisikoPrioritasDelete {
	return func(ctx context.Context, req DaftarRisikoPrioritasDeleteReq) (*DaftarRisikoPrioritasDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.DaftarRisikoPrioritas{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &DaftarRisikoPrioritasDeleteRes{}, nil
	}
}
