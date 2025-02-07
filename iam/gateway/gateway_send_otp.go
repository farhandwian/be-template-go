package gateway

import (
	"context"
	"iam/model"
	"shared/core"
	"shared/helper"
)

type SendOTPReq struct {
	PhoneNumber model.PhoneNumber
	Message     string
}

type SendOTPRes struct{}

type SendOTP = core.ActionHandler[SendOTPReq, SendOTPRes]

func ImplSendOTP() SendOTP {
	return func(ctx context.Context, request SendOTPReq) (*SendOTPRes, error) {

		// send using wa, telegram or anything

		return &SendOTPRes{}, nil
	}
}

func ImplSendOTPToWA() SendOTP {
	return func(ctx context.Context, request SendOTPReq) (*SendOTPRes, error) {

		if err := helper.SendMessage(request.Message, string(request.PhoneNumber)); err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &SendOTPRes{}, nil
	}
}
