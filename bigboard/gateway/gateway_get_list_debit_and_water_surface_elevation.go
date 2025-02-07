package gateway

import (
	"bigboard/model"
	"context"
	"shared/core"

	"gorm.io/gorm"
)

type GetListDebitAndWaterSurfaceElevationReq struct{}

type GetListDebitAndWaterSurfaceElevationResp struct {
	DebitAndWaterSurfaceElevation []model.DebitAndWaterSurfaceElevation
}

type GetListDebitAndWaterSurfaceElevationGateway = core.ActionHandler[GetListDebitAndWaterSurfaceElevationReq, GetListDebitAndWaterSurfaceElevationResp]

func ImplGetListDebitAndWaterSurfaceElevation(tsDB *gorm.DB) GetListDebitAndWaterSurfaceElevationGateway {
	return func(ctx context.Context, request GetListDebitAndWaterSurfaceElevationReq) (*GetListDebitAndWaterSurfaceElevationResp, error) {

		var debitAndWaterSurfaceElevation []model.DebitAndWaterSurfaceElevation

		subqueryDebit := tsDB.Table("actual_debits").
			Select("water_channel_door_id, MAX(timestamp) as latest_timestamp").
			Group("water_channel_door_id")

		subqueryWaterSurfaceElevation := tsDB.Table("water_surface_elevations").
			Select("water_channel_door_id, MAX(timestamp) as latest_timestamp").
			Group("water_channel_door_id")

		err := tsDB.Table("actual_debits AS ad").
			Select("ad.water_channel_door_id, ad.actual_debit as debit, wse.water_level as water_surface_elevation").
			Joins("JOIN (?) AS latest_ad ON latest_ad.water_channel_door_id = ad.water_channel_door_id AND latest_ad.latest_timestamp = ad.timestamp", subqueryDebit).
			Joins("JOIN water_surface_elevations AS wse ON wse.water_channel_door_id = ad.water_channel_door_id").
			Joins("JOIN (?) AS latest_wse ON latest_wse.water_channel_door_id = wse.water_channel_door_id AND latest_wse.latest_timestamp = wse.timestamp", subqueryWaterSurfaceElevation).
			Scan(&debitAndWaterSurfaceElevation).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetListDebitAndWaterSurfaceElevationResp{
			DebitAndWaterSurfaceElevation: debitAndWaterSurfaceElevation,
		}, nil
	}
}
