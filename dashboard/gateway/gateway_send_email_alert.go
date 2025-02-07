package gateway

import (
	"context"
	"shared/core"
	"shared/helper"
)

type SendEmailReq struct {
	EmailRecipients []string
	Subject         string
	Body            string
}

type SendEmailRes struct{}

type SendEmail = core.ActionHandler[SendEmailReq, SendEmailRes]

func ImplSendEmailUsingGmail() SendEmail { // TODO duplicated
	return func(ctx context.Context, req SendEmailReq) (*SendEmailRes, error) {

		if err := helper.SendEmail(req.Subject, req.Body, req.EmailRecipients...); err != nil {
			return nil, err
		}

		return &SendEmailRes{}, nil
	}
}
