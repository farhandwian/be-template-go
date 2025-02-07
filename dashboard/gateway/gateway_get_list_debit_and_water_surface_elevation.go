package gateway

import (
	"context"
	"dashboard/model"
	"shared/core"

	"gorm.io/gorm"
)

type GetListDebitAndWaterSurfaceElevationReq struct {
	WaterChannelDoorIDs []int // New field to hold the list of IDs
}

type GetListDebitAndWaterSurfaceElevationResp struct {
	DebitAndWaterSurfaceElevation []model.WaterChannelDoorDebitAndWaterSurfaceElevation
}

type GetListDebitAndWaterSurfaceElevationGateway = core.ActionHandler[GetListDebitAndWaterSurfaceElevationReq, GetListDebitAndWaterSurfaceElevationResp]

func ImplGetListDebitAndWaterSurfaceElevation(tsDB *gorm.DB) GetListDebitAndWaterSurfaceElevationGateway {
	return func(ctx context.Context, request GetListDebitAndWaterSurfaceElevationReq) (*GetListDebitAndWaterSurfaceElevationResp, error) {
		var debitAndWaterSurfaceElevation []model.WaterChannelDoorDebitAndWaterSurfaceElevation

		query := tsDB.Table("actual_debits AS ad").
			Select("ad.water_channel_door_id, ad.actual_debit as debit, wse.water_level as water_surface_elevation")

		subqueryDebit := tsDB.Table("actual_debits").
			Select("water_channel_door_id, MAX(timestamp) as latest_timestamp").
			Group("water_channel_door_id")

		subqueryWaterSurfaceElevation := tsDB.Table("water_surface_elevations").
			Select("water_channel_door_id, MAX(timestamp) as latest_timestamp").
			Group("water_channel_door_id")

		// Add WHERE clause if WaterChannelDoorIDs are provided
		if len(request.WaterChannelDoorIDs) > 0 {
			query = query.Where("ad.water_channel_door_id IN ?", request.WaterChannelDoorIDs)
			subqueryDebit = subqueryDebit.Where("water_channel_door_id IN ?", request.WaterChannelDoorIDs)
			subqueryWaterSurfaceElevation = subqueryWaterSurfaceElevation.Where("water_channel_door_id IN ?", request.WaterChannelDoorIDs)
		}

		err := query.Joins("JOIN (?) AS latest_ad ON latest_ad.water_channel_door_id = ad.water_channel_door_id AND latest_ad.latest_timestamp = ad.timestamp", subqueryDebit).
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
