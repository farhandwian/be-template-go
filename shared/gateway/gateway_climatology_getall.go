package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type GetClimatologyListReq struct {
	Keyword string
}

type GetClimatologyListResp struct {
	Climatology []model.ClimatologyPost
}

type GetClimatologyListGateway = core.ActionHandler[GetClimatologyListReq, GetClimatologyListResp]

func ImplGetClimatologyList(db *gorm.DB) GetClimatologyListGateway {
	return func(ctx context.Context, request GetClimatologyListReq) (*GetClimatologyListResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var climatologyData []model.ClimatologyPost

		if request.Keyword != "" {
			query = query.Where("name LIKE ?", "%"+request.Keyword+"%")
		}

		err := query.Find(&climatologyData).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &GetClimatologyListResp{
			Climatology: climatologyData,
		}, err
	}
}
