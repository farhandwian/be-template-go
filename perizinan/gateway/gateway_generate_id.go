package gateway

import (
	"context"
	"shared/core"

	"github.com/google/uuid"
)

type GenerateIdReq struct{}

type GenerateIdRes struct {
	UUID string
}

type GenerateId = core.ActionHandler[GenerateIdReq, GenerateIdRes]

func ImplGenerateId() GenerateId {
	return func(ctx context.Context, request GenerateIdReq) (*GenerateIdRes, error) {
		return &GenerateIdRes{UUID: uuid.New().String()}, nil
	}
}

func (r GenerateIdReq) Validate() error {
	return nil
}
