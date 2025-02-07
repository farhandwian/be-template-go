package gateway

import (
	"context"
	"dashboard/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type GetWaterChannelListReq struct {
}

type GetWaterChannelListResp struct {
	WaterChannels []model.WaterChannel
}

type GetWaterChannelListGateway = core.ActionHandler[GetWaterChannelListReq, GetWaterChannelListResp]

func ImplGetWaterChannelList(db *gorm.DB) GetWaterChannelListGateway {
	return func(ctx context.Context, request GetWaterChannelListReq) (*GetWaterChannelListResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var waterChannels []model.WaterChannel

		err := query.Find(&waterChannels).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &GetWaterChannelListResp{
			WaterChannels: waterChannels,
		}, err
	}
}
