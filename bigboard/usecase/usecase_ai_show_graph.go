package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
)

type AiShowGraphReq struct {
	WaterChannelDoorID int    `json:"water_channel_door_id"`
	GraphType          string `json:"graph_type"`
	StartDate          string `json:"start_date"`
	EndDate            string `json:"end_date"`
}

type AiShowGraphRes struct {
	Params map[string]interface{} `json:"params"`
}

type AiShowGraphResUseCase = core.ActionHandler[AiShowGraphReq, AiShowGraphRes]

func ImplAiShowGraphUseCase(sendSSEMessageGateway gateway.SendSSEMessage) AiShowGraphResUseCase {
	return func(ctx context.Context, req AiShowGraphReq) (*AiShowGraphRes, error) {

		res := &AiShowGraphRes{
			Params: map[string]interface{}{
				"water_channel_door_id": req.WaterChannelDoorID,
				"graph_type":            req.GraphType,
				"start_date":            req.StartDate,
				"end_date":              req.EndDate,
			},
		}
		_, _ = sendSSEMessageGateway(ctx, gateway.SendSSEMessageReq{
			Subject:      "show-graph",
			FunctionName: "showGraph",
			Data:         res,
		})
		return &AiShowGraphRes{Params: res.Params}, nil

	}
}
