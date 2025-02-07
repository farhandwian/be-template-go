package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type GetLakeListReq struct {
}

type GetLakeListResp struct {
	Lakes []model.Danau
}

type GetLakeListGateway = core.ActionHandler[GetLakeListReq, GetLakeListResp]

func ImplGetLakeList(db *gorm.DB) GetLakeListGateway {
	return func(ctx context.Context, request GetLakeListReq) (*GetLakeListResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var lakes []model.Danau

		err := query.Where("ws_name=?", "CITANDUY").Find(&lakes).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &GetLakeListResp{
			Lakes: lakes,
		}, err
	}
}
