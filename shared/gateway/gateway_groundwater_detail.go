package gateway

import (
	"context"
	"errors"
	"fmt"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type GroundWaterGetDetailReq struct {
	ID string
}
type GroundWaterGetDetailResp struct {
	GroundWater model.AirTanah
}

type GroundWaterGetDetailGateway = core.ActionHandler[GroundWaterGetDetailReq, GroundWaterGetDetailResp]

func ImplGroundWaterGetDetailGateway(db *gorm.DB) GroundWaterGetDetailGateway {
	return func(ctx context.Context, request GroundWaterGetDetailReq) (*GroundWaterGetDetailResp, error) {

		query := middleware.GetDBFromContext(ctx, db)

		var groundWater model.AirTanah

		err := query.Where("id=?", request.ID).First(&groundWater).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("ground water id %v is not found", request.ID)
			}
			return nil, core.NewInternalServerError(err)
		}
		return &GroundWaterGetDetailResp{
			GroundWater: groundWater,
		}, err
	}
}
