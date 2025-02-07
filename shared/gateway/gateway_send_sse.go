package gateway

import (
	"context"

	"shared/core"
	"shared/helper"
)

type SendSSEMessageReq struct {
	Subject      string
	FunctionName string
	Data         any
}

type SendSSEMessageRes struct {
}

type SendSSEMessage = core.ActionHandler[SendSSEMessageReq, SendSSEMessageRes]

func ImplSendSSEMessage(sse *helper.SSE) SendSSEMessage {
	return func(ctx context.Context, request SendSSEMessageReq) (*SendSSEMessageRes, error) {

		if sse == nil {
			return &SendSSEMessageRes{}, nil
		}

		if err := sse.BroadcastToClients(ctx, helper.Message{
			Subject:      request.Subject,
			FunctionName: request.FunctionName,
			Data:         request.Data,
		}); err != nil {
			return nil, err
		}

		return &SendSSEMessageRes{}, nil
	}
}
