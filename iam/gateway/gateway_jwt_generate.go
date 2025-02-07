package gateway

import (
	"context"
	"shared/core"
	"shared/helper"
	"time"
)

type GenerateJWTReq struct {
	Payload []byte
	Now     time.Time
	Expired time.Duration
}

type GenerateJWTRes struct {
	JWTToken string
}

type GenerateJWT = core.ActionHandler[GenerateJWTReq, GenerateJWTRes]

func ImplGenerateJWT(jwt helper.JWTTokenizer) GenerateJWT {

	return func(ctx context.Context, request GenerateJWTReq) (*GenerateJWTRes, error) {

		token, err := jwt.CreateToken(request.Payload, request.Now, request.Expired)
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GenerateJWTRes{JWTToken: token}, nil
	}
}
