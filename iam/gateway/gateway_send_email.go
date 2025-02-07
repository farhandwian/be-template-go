package gateway

import (
	"context"
	"fmt"
	"iam/model"
	"shared/core"
	"shared/helper"
)

type SendEmailReq struct {
	EmailRecipient model.Email
	Subject        string
	Body           string
}

type SendEmailRes struct{}

type SendEmail = core.ActionHandler[SendEmailReq, SendEmailRes]

func ImplSendEmail() SendEmail {
	return func(ctx context.Context, request SendEmailReq) (*SendEmailRes, error) {

		// send email here ...
		fmt.Printf("EMAIL Sent to '%s', with subject '%s' and with body '%s'\n", request.EmailRecipient, request.Subject, request.Body)

		return &SendEmailRes{}, nil
	}
}

func ImplSendEmailUsingGmail() SendEmail {
	return func(ctx context.Context, req SendEmailReq) (*SendEmailRes, error) {

		if err := helper.SendEmail(req.Subject, req.Body, string(req.EmailRecipient)); err != nil {
			return nil, err
		}

		return &SendEmailRes{}, nil
	}
}
