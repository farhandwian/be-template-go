package gateway

import (
	"context"
	"dashboard/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type AssetSaveReq struct {
	Asset model.Asset
}

type AssetSaveRes struct {
	ID string
}

type AssetSave = core.ActionHandler[AssetSaveReq, AssetSaveRes]

func ImplAssetSave(db *gorm.DB) AssetSave {
	return func(ctx context.Context, req AssetSaveReq) (*AssetSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.Asset).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &AssetSaveRes{ID: req.Asset.ID}, nil
	}
}
