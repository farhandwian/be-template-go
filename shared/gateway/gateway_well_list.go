package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type GetWellListReq struct {
}

type GetWellListResp struct {
	Wells []model.Sumur
}

type GetWellListGateway = core.ActionHandler[GetWellListReq, GetWellListResp]

func ImplGetWellList(db *gorm.DB) GetWellListGateway {
	return func(ctx context.Context, request GetWellListReq) (*GetWellListResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var weirs []model.Sumur

		err := query.Where("ws_name=?", "CITANDUY").Find(&weirs).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &GetWellListResp{
			Wells: weirs,
		}, err
	}
}
