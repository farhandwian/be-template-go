package gateway

import (
	"context"
	"errors"
	"shared/core"

	"golang.org/x/crypto/bcrypt"
)

type PasswordValidateReq struct {
	PasswordEncrypted string
	PasswordPlain     string
}

type PasswordValidateRes struct {
}

type PasswordValidate = core.ActionHandler[PasswordValidateReq, PasswordValidateRes]

func ImplPasswordValidate() PasswordValidate {
	return func(ctx context.Context, request PasswordValidateReq) (*PasswordValidateRes, error) {

		err := bcrypt.CompareHashAndPassword([]byte(request.PasswordEncrypted), []byte(request.PasswordPlain))
		if err != nil {
			return nil, errors.New(err.Error())
		}

		return &PasswordValidateRes{}, nil
	}
}
