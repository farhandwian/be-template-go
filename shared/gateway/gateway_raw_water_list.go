package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type GetRawWaterListReq struct {
}

type GetRawWaterListResp struct {
	RawWaters []model.AirBaku
}

type GetRawWaterListGateway = core.ActionHandler[GetRawWaterListReq, GetRawWaterListResp]

func ImplGetRawWaterList(db *gorm.DB) GetRawWaterListGateway {
	return func(ctx context.Context, request GetRawWaterListReq) (*GetRawWaterListResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var groundWaters []model.AirBaku

		err := query.Where("ws_name=?", "CITANDUY").Find(&groundWaters).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &GetRawWaterListResp{
			RawWaters: groundWaters,
		}, err
	}
}
