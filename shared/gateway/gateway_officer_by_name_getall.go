package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type GetOfficerListByNameReq struct {
	Name []string
}

type GetOfficerListByNameResp struct {
	RainFalls []model.PostOfficer
}

type GetOfficerListGateway = core.ActionHandler[GetOfficerListByNameReq, GetOfficerListByNameResp]

func ImplGetOfficerPostByNameList(db *gorm.DB) GetOfficerListGateway {
	return func(ctx context.Context, request GetOfficerListByNameReq) (*GetOfficerListByNameResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var officerDetails []model.PostOfficer

		err := query.Where("name in ?", request.Name).Find(&officerDetails).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &GetOfficerListByNameResp{
			RainFalls: officerDetails,
		}, err
	}
}
