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

type SedimentControlGetDetailReq struct {
	ID string
}
type SedimentControlGetDetailResp struct {
	SedimentControl model.PengendaliSedimen
}

type SedimentControlGetDetailGateway = core.ActionHandler[SedimentControlGetDetailReq, SedimentControlGetDetailResp]

func ImplSedimentControlGetDetailGateway(db *gorm.DB) SedimentControlGetDetailGateway {
	return func(ctx context.Context, request SedimentControlGetDetailReq) (*SedimentControlGetDetailResp, error) {

		query := middleware.GetDBFromContext(ctx, db)

		var sedimentControl model.PengendaliSedimen

		err := query.Where("id=?", request.ID).Find(&sedimentControl).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("sediment control id %v is not found", request.ID)
			}
			return nil, core.NewInternalServerError(err)
		}
		return &SedimentControlGetDetailResp{
			SedimentControl: sedimentControl,
		}, err
	}
}
