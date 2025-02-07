package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type RainfallGetDetailReq struct {
	ID string
}
type RainfallGetDetailResp struct {
	RainFall model.RainfallPost
}

type RainfallGetDetailGateway = core.ActionHandler[RainfallGetDetailReq, RainfallGetDetailResp]

func ImplRainfallGetDetailGateway(db *gorm.DB) RainfallGetDetailGateway {
	return func(ctx context.Context, request RainfallGetDetailReq) (*RainfallGetDetailResp, error) {

		query := middleware.GetDBFromContext(ctx, db)

		var obj model.RainfallPost

		err := query.Where("id = ? or external_id = ?", request.ID, request.ID).First(&obj).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &RainfallGetDetailResp{
			RainFall: obj,
		}, err
	}
}
