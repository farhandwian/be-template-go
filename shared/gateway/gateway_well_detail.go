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

type WellGetDetailReq struct {
	ID string
}
type WellGetDetailResp struct {
	Well model.Sumur
}

type WellGetDetailGateway = core.ActionHandler[WellGetDetailReq, WellGetDetailResp]

func ImplWellGetDetailGateway(db *gorm.DB) WellGetDetailGateway {
	return func(ctx context.Context, request WellGetDetailReq) (*WellGetDetailResp, error) {

		query := middleware.GetDBFromContext(ctx, db)

		var well model.Sumur

		err := query.Where("id=?", request.ID).First(&well).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("well id %v is not found", request.ID)
			}
			return nil, core.NewInternalServerError(err)
		}
		return &WellGetDetailResp{
			Well: well,
		}, err
	}
}
