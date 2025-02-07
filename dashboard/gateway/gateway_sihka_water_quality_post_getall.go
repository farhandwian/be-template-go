package gateway

import (
	"context"
	"dashboard/model"
	"gorm.io/gorm"
	"shared/core"
	"shared/middleware"
)

type GetWaterQualityPostListReq struct {
	Date  string
	River string
	Name  string
}

type WaterLevelQualityResp struct {
	WaterQualityPost []model.WaterQualityPost
}

type GetWaterQualityPostListGateway = core.ActionHandler[GetWaterQualityPostListReq, WaterLevelQualityResp]

func ImplGetWaterQualityPostList(db *gorm.DB) GetWaterQualityPostListGateway {
	return func(ctx context.Context, request GetWaterQualityPostListReq) (*WaterLevelQualityResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var waterQualityPosts []model.WaterQualityPost

		if request.River != "" {
			query = query.Where("location LIKE ?", "%"+request.River+"%")
		}
		if request.Name != "" {
			query = query.Where("name LIKE ?", "%"+request.Name+"%")
		}

		// Execute the query
		err := query.Find(&waterQualityPosts).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		// Return the response with water level posts
		return &WaterLevelQualityResp{
			WaterQualityPost: waterQualityPosts,
		}, nil
	}
}
