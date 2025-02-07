package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type GetWeirListReq struct {
}

type GetWeirListResp struct {
	Weirs []model.Bendung
}

type GetWeirListGateway = core.ActionHandler[GetWeirListReq, GetWeirListResp]

func ImplGetWeirList(db *gorm.DB) GetWeirListGateway {
	return func(ctx context.Context, request GetWeirListReq) (*GetWeirListResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var weirs []model.Bendung

		err := query.Where("ws_name=?", "CITANDUY").Find(&weirs).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &GetWeirListResp{
			Weirs: weirs,
		}, err
	}
}
