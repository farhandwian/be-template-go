package gateway

import (
	"context"
	"shared/core"
	"shared/helper"
)

type ValidateJWTReq struct {
	Token string
}

type ValidateJWTRes struct {
	Payload []byte
}

type ValidateJWT = core.ActionHandler[ValidateJWTReq, ValidateJWTRes]

func ImplValidateJWT(jwt helper.JWTTokenizer) ValidateJWT {

	return func(ctx context.Context, request ValidateJWTReq) (*ValidateJWTRes, error) {

		content, err := jwt.VerifyToken(request.Token)
		if err != nil {
			return nil, err
		}

		return &ValidateJWTRes{Payload: content}, nil
	}
}
