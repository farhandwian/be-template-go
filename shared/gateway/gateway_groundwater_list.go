package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type GetGroundWaterListReq struct {
}

type GetGroundWaterListResp struct {
	GroundWaters []model.AirTanah
}

type GetGroundWaterListGateway = core.ActionHandler[GetGroundWaterListReq, GetGroundWaterListResp]

func ImplGetGroundWaterList(db *gorm.DB) GetGroundWaterListGateway {
	return func(ctx context.Context, request GetGroundWaterListReq) (*GetGroundWaterListResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var groundWaters []model.AirTanah

		err := query.Where("ws_name=?", "CITANDUY").Find(&groundWaters).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &GetGroundWaterListResp{
			GroundWaters: groundWaters,
		}, err
	}
}
