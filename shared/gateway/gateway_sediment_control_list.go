package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type GetSedimentControlListReq struct {
}

type GetSedimentControlListResp struct {
	SedimentControls []model.PengendaliSedimen
}

type GetSedimentControlListGateway = core.ActionHandler[GetSedimentControlListReq, GetSedimentControlListResp]

func ImplGetSedimentControlList(db *gorm.DB) GetSedimentControlListGateway {
	return func(ctx context.Context, request GetSedimentControlListReq) (*GetSedimentControlListResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var sedimentControls []model.PengendaliSedimen

		err := query.Where("ws_name=?", "CITANDUY").Find(&sedimentControls).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &GetSedimentControlListResp{
			SedimentControls: sedimentControls,
		}, err
	}
}
