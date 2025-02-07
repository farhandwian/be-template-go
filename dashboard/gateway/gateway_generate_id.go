package gateway

import (
	"context"
	"shared/core"

	"github.com/google/uuid"
)

type GenerateIdReq struct{}

type GenerateIdRes struct {
	RandomId string
}

type GenerateId = core.ActionHandler[GenerateIdReq, GenerateIdRes]

func ImplGenerateId() GenerateId {
	return func(ctx context.Context, request GenerateIdReq) (*GenerateIdRes, error) {
		return &GenerateIdRes{RandomId: uuid.New().String()}, nil
	}
}

func (r GenerateIdReq) Validate() error {
	return nil
}
