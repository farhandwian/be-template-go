package gateway

import (
	"context"
	"shared/core"
	"shared/helper"
)

type SendWhatsAppReq struct {
	PhoneNumbers []string
	Message      string
}

type SendWhatsAppRes struct{}

type SendWhatsApp = core.ActionHandler[SendWhatsAppReq, SendWhatsAppRes]

func ImplSendWhatsApp() SendWhatsApp { // TODO duplicated
	return func(ctx context.Context, request SendWhatsAppReq) (*SendWhatsAppRes, error) {

		if err := helper.SendMessage(request.Message, request.PhoneNumbers...); err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &SendWhatsAppRes{}, nil
	}
}
