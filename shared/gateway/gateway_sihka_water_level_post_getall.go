package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type GetWaterLevelPostListReq struct {
	River string
	Name  string
}

type WaterLevelPostResp struct {
	WaterLevelPost []model.WaterLevelPost
}

type GetWaterLevelPostListGateway = core.ActionHandler[GetWaterLevelPostListReq, WaterLevelPostResp]

func ImplGetWaterLevelPostList(db *gorm.DB) GetWaterLevelPostListGateway {
	return func(ctx context.Context, request GetWaterLevelPostListReq) (*WaterLevelPostResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var waterLevelPosts []model.WaterLevelPost

		if request.River != "" {
			query = query.Where("location LIKE ?", "%"+request.River+"%")
		}
		if request.Name != "" {
			query = query.Where("name LIKE ?", "%"+request.Name+"%")
		}

		// Execute the query
		err := query.Find(&waterLevelPosts).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		// Return the response with water level posts
		return &WaterLevelPostResp{
			WaterLevelPost: waterLevelPosts,
		}, nil
	}
}
