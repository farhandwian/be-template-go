package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type GetRainfallPostListReq struct {
	City   string
	Vendor string
	Name   string
}

type GetRainfallPostListResp struct {
	RainFalls []model.RainfallPost
}

type GetRainFallPostListGateway = core.ActionHandler[GetRainfallPostListReq, GetRainfallPostListResp]

func ImplGetRainfallPostList(db *gorm.DB) GetRainFallPostListGateway {
	return func(ctx context.Context, request GetRainfallPostListReq) (*GetRainfallPostListResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var rainfallPosts []model.RainfallPost

		if request.City != "" {
			query = query.Where("city = ?", request.City)
		}
		if request.Vendor != "" {
			query = query.Where("vendor = ?", request.Vendor)
		}
		if request.Name != "" {
			query = query.Where("name LIKE ?", "%"+request.Name+"%")
		}

		err := query.Find(&rainfallPosts).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &GetRainfallPostListResp{
			RainFalls: rainfallPosts,
		}, err
	}
}
