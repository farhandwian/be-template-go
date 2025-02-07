package gateway

import (
	"context"
	"dashboard/model"
	"fmt"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type AssetGetByIDReq struct {
	ID string
}

type AssetGetByIDRes struct {
	Asset model.Asset
}

type AssetGetByID = core.ActionHandler[AssetGetByIDReq, AssetGetByIDRes]

func ImplAssetGetByID(db *gorm.DB) AssetGetByID {
	return func(ctx context.Context, req AssetGetByIDReq) (*AssetGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var asset model.Asset
		if err := query.First(&asset, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("asset id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &AssetGetByIDRes{Asset: asset}, nil
	}
}
