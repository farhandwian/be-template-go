// File: gateway/gateway_asset.go

package gateway

import (
	"context"
	"dashboard/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type AssetDeleteReq struct {
	ID string
}

type AssetDeleteRes struct{}

type AssetDelete = core.ActionHandler[AssetDeleteReq, AssetDeleteRes]

func ImplAssetDelete(db *gorm.DB) AssetDelete {
	return func(ctx context.Context, req AssetDeleteReq) (*AssetDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.Asset{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &AssetDeleteRes{}, nil
	}
}
