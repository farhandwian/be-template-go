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

type WeirGetDetailReq struct {
	ID string
}
type WeirGetDetailResp struct {
	Weir model.Bendung
}

type WeirGetDetailGateway = core.ActionHandler[WeirGetDetailReq, WeirGetDetailResp]

func ImplWeirGetDetailGateway(db *gorm.DB) WeirGetDetailGateway {
	return func(ctx context.Context, request WeirGetDetailReq) (*WeirGetDetailResp, error) {

		query := middleware.GetDBFromContext(ctx, db)

		var weir model.Bendung

		err := query.Where("id=?", request.ID).First(&weir).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("weir id %v is not found", request.ID)
			}
			return nil, core.NewInternalServerError(err)
		}
		return &WeirGetDetailResp{
			Weir: weir,
		}, err
	}
}
