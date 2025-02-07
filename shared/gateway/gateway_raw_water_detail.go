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

type RawWaterGetDetailReq struct {
	ID string
}
type RawWaterGetDetailResp struct {
	RawWater model.AirBaku
}

type RawWaterGetDetailGateway = core.ActionHandler[RawWaterGetDetailReq, RawWaterGetDetailResp]

func ImplRawWaterGetDetailGateway(db *gorm.DB) RawWaterGetDetailGateway {
	return func(ctx context.Context, request RawWaterGetDetailReq) (*RawWaterGetDetailResp, error) {

		query := middleware.GetDBFromContext(ctx, db)

		var rawWater model.AirBaku

		err := query.Where("id=?", request.ID).First(&rawWater).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("raw water id %v is not found", request.ID)
			}
			return nil, core.NewInternalServerError(err)
		}
		return &RawWaterGetDetailResp{
			RawWater: rawWater,
		}, err
	}
}
