package usecase

import (
	"context"
	"dashboard/gateway"
	"shared/core"
	"sort"
)

type WaterChannelDoorByKeywordReq struct {
	Keyword string `json:"keyword"`
}

type WaterChannelDoorByKeywordRes struct {
	WaterChannelDoors []gateway.WaterChannelDoor `json:"water_channel_doors"`
}

type WaterChannelDoorByKeyword = core.ActionHandler[WaterChannelDoorByKeywordReq, WaterChannelDoorByKeywordRes]

func ImplWaterChannelDoorByKeyword(getByKeyword gateway.WaterChannelDoorByKeyword) WaterChannelDoorByKeyword {
	return func(ctx context.Context, req WaterChannelDoorByKeywordReq) (*WaterChannelDoorByKeywordRes, error) {

		if len(req.Keyword) < 2 {
			return &WaterChannelDoorByKeywordRes{WaterChannelDoors: []gateway.WaterChannelDoor{}}, nil
		}

		res, err := getByKeyword(ctx, gateway.WaterChannelDoorByKeywordReq{Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		sort.Slice(res.Items, func(i, j int) bool {
			return res.Items[i].Name < res.Items[j].Name
		})

		return &WaterChannelDoorByKeywordRes{WaterChannelDoors: res.Items}, nil
	}
}
