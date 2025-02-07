package gateway

import (
	"context"
	"dashboard/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type GetWaterLevelListReq struct {
}

type GetWaterLevelListResp struct {
	WaterLevels []model.DugaAir
}

type GetWaterLevelListGateway = core.ActionHandler[GetWaterLevelListReq, GetWaterLevelListResp]

func ImplGetWaterLevelList(db *gorm.DB) GetWaterLevelListGateway {
	return func(ctx context.Context, request GetWaterLevelListReq) (*GetWaterLevelListResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var rainFalls []model.DugaAir

		err := query.Find(&rainFalls).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &GetWaterLevelListResp{
			WaterLevels: rainFalls,
		}, err
	}
}
