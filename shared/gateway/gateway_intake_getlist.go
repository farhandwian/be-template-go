package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type GetIntakeListReq struct {
}

type GetIntakeListResp struct {
	Intakes []model.Intake
}

type GetIntakeListGateway = core.ActionHandler[GetIntakeListReq, GetIntakeListResp]

func ImplGetIntakeList(db *gorm.DB) GetIntakeListGateway {
	return func(ctx context.Context, request GetIntakeListReq) (*GetIntakeListResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var intakes []model.Intake

		err := query.Where("ws_name=?", "CITANDUY").Find(&intakes).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &GetIntakeListResp{
			Intakes: intakes,
		}, err
	}
}
