package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
)

type AiBackToHomeReq struct {
}

type AiBackToHomeRes struct {
}

type AiBackToHomeUseCase = core.ActionHandler[AiBackToHomeReq, AiBackToHomeRes]

func ImplAiBackToHome(sendSSEMessageGateway gateway.SendSSEMessage) AiBackToHomeUseCase {
	return func(ctx context.Context, req AiBackToHomeReq) (*AiBackToHomeRes, error) {

		_, err := sendSSEMessageGateway(ctx, gateway.SendSSEMessageReq{
			Subject:      "back-to-home",
			FunctionName: "backToHome",
			Data:         map[string]string{"success": "true"},
		})
		if err != nil {
			return nil, err
		}

		return &AiBackToHomeRes{}, nil

	}
}
