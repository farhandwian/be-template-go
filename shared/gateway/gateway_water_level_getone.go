package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type WaterLevelGetDetailReq struct {
	ID string
}
type WaterLevelGetDetailResp struct {
	WaterLevel model.WaterLevelPost
}

type WaterLevelGetDetailGateway = core.ActionHandler[WaterLevelGetDetailReq, WaterLevelGetDetailResp]

func ImplWaterLevelGetDetailGateway(db *gorm.DB) WaterLevelGetDetailGateway {
	return func(ctx context.Context, request WaterLevelGetDetailReq) (*WaterLevelGetDetailResp, error) {

		query := middleware.GetDBFromContext(ctx, db)

		var obj model.WaterLevelPost

		err := query.Where("id = ?", request.ID).First(&obj).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &WaterLevelGetDetailResp{
			WaterLevel: obj,
		}, err
	}
}
