package gateway

import (
	"context"
	"dashboard/model"
	"fmt"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type AssetGetAllReq struct {
	Keyword string
	Page    int
	Size    int
}

type AssetGetAllRes struct {
	Assets []model.Asset `json:"assets"`
	Count  int64         `json:"count"`
}

type AssetGetAll = core.ActionHandler[AssetGetAllReq, AssetGetAllRes]

func ImplAssetGetAll(db *gorm.DB) AssetGetAll {
	return func(ctx context.Context, req AssetGetAllReq) (*AssetGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("name LIKE ?", keyword).
				Or("pic LIKE ?", keyword).
				Or("location LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.Asset{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.Asset

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &AssetGetAllRes{
			Assets: objs,
			Count:  count,
		}, nil
	}
}
