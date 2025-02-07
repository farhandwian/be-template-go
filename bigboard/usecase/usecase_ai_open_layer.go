package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
)

type AiOpenLayerReq struct {
	LayersName []AiOpenLayer `json:"layers"`
}

type AiOpenLayer struct {
	LayerName string                 `json:"layer_name"`
	Params    map[string]interface{} `json:"params"`
}

type AiOpenLayerRes struct {
	Layers []AiOpenLayer `json:"layers"`
}

type AiOpenLayerUseCase = core.ActionHandler[AiOpenLayerReq, AiOpenLayerRes]

func ImplAiOpenLayer(sendSSEMessageGateway gateway.SendSSEMessage) AiOpenLayerUseCase {
	return func(ctx context.Context, req AiOpenLayerReq) (*AiOpenLayerRes, error) {

		_, _ = sendSSEMessageGateway(ctx, gateway.SendSSEMessageReq{
			Subject:      "open-layers",
			FunctionName: "openLayers",
			Data:         req.LayersName,
		})
		return &AiOpenLayerRes{Layers: req.LayersName}, nil

	}
}
