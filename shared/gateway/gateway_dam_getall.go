package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type GetDamListReq struct {
}

type GetDamListResp struct {
	Dams []model.Bendungan
}

type GetDamListGateway = core.ActionHandler[GetDamListReq, GetDamListResp]

func ImplGetDamList(db *gorm.DB) GetDamListGateway {
	return func(ctx context.Context, request GetDamListReq) (*GetDamListResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var dams []model.Bendungan

		err := query.Where("ws_name=?", "CITANDUY").Find(&dams).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &GetDamListResp{
			Dams: dams,
		}, err
	}
}
