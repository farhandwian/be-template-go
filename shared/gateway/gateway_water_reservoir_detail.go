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

type WaterReservoirGetDetailReq struct {
	ID string
}
type WaterReservoirGetDetailResp struct {
	WaterReservoir model.Embung
}

type WaterReservoirGetAllGatewayGateway = core.ActionHandler[WaterReservoirGetDetailReq, WaterReservoirGetDetailResp]

func ImplWaterReservoirGetDetailGateway(db *gorm.DB) WaterReservoirGetAllGatewayGateway {
	return func(ctx context.Context, request WaterReservoirGetDetailReq) (*WaterReservoirGetDetailResp, error) {

		query := middleware.GetDBFromContext(ctx, db)

		var waterReservoir model.Embung

		err := query.Where("id=?", request.ID).First(&waterReservoir).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("water reservoir id %v is not found", request.ID)
			}
			return nil, core.NewInternalServerError(err)
		}
		return &WaterReservoirGetDetailResp{
			WaterReservoir: waterReservoir,
		}, err
	}
}
