// File: gateway/gateway_asset.go

package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type SpipDeleteReq struct {
	ID string
}

type SpipDeleteRes struct{}

type SpipDelete = core.ActionHandler[SpipDeleteReq, SpipDeleteRes]

func ImplSpipDelete(db *gorm.DB) SpipDelete {
	return func(ctx context.Context, req SpipDeleteReq) (*SpipDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.SPIP{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &SpipDeleteRes{}, nil
	}
}
