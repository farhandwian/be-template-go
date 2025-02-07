package gateway

import (
	"context"
	"errors"
	"shared/core"

	"golang.org/x/crypto/bcrypt"
)

type PasswordEncryptReq struct {
	PasswordPlain string
}

type PasswordEncryptRes struct {
	PasswordEncrypted string
}

type PasswordEncrypt = core.ActionHandler[PasswordEncryptReq, PasswordEncryptRes]

func ImplPasswordEncrypt() PasswordEncrypt {
	return func(ctx context.Context, request PasswordEncryptReq) (*PasswordEncryptRes, error) {

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.PasswordPlain), bcrypt.DefaultCost)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		return &PasswordEncryptRes{PasswordEncrypted: string(hashedPassword)}, nil
	}
}
