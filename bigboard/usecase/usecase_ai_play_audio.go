package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
)

type AiPlayAudioReq struct {
	FileName string `json:"file_name"`
}

type AiPlayAudioResp struct {
	FileName string `json:"file_name"`
}

type AiPlayAudioUseCase = core.ActionHandler[AiPlayAudioReq, AiPlayAudioResp]

func ImplAiPlayAudioUseCase(sendSSEMessageGateway gateway.SendSSEMessage) AiPlayAudioUseCase {
	return func(ctx context.Context, req AiPlayAudioReq) (*AiPlayAudioResp, error) {

		_, _ = sendSSEMessageGateway(ctx, gateway.SendSSEMessageReq{
			Subject:      "play-audio",
			FunctionName: "playAudio",
			Data:         req.FileName,
		})
		return &AiPlayAudioResp{FileName: req.FileName}, nil

	}
}
