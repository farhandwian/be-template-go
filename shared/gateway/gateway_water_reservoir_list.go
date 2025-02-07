package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type GetWaterReservoirListReq struct {
}

type GetWaterReservoirListResp struct {
	WaterReservoirs []model.Embung
}

type GetWaterReservoirListGateway = core.ActionHandler[GetWaterReservoirListReq, GetWaterReservoirListResp]

func ImplGetWaterReservoirList(db *gorm.DB) GetWaterReservoirListGateway {
	return func(ctx context.Context, request GetWaterReservoirListReq) (*GetWaterReservoirListResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var waterReservoirs []model.Embung

		err := query.Where("ws_name=?", "CITANDUY").Find(&waterReservoirs).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &GetWaterReservoirListResp{
			WaterReservoirs: waterReservoirs,
		}, err
	}
}
