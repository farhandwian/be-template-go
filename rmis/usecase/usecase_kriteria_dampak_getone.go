package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type KriteriaDampakGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type KriteriaDampakGetByIDUseCaseRes struct {
	KriteriaDampak model.KriteriaDampak `json:"kriteria_dampak"`
}

type KriteriaDampakGetByIDUseCase = core.ActionHandler[KriteriaDampakGetByIDUseCaseReq, KriteriaDampakGetByIDUseCaseRes]

func ImplKriteriaDampakGetByIDUseCase(getKriteriaDampakByID gateway.KriteriaDampakGetByID) KriteriaDampakGetByIDUseCase {
	return func(ctx context.Context, req KriteriaDampakGetByIDUseCaseReq) (*KriteriaDampakGetByIDUseCaseRes, error) {
		res, err := getKriteriaDampakByID(ctx, gateway.KriteriaDampakGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &KriteriaDampakGetByIDUseCaseRes{KriteriaDampak: res.KriteriaDampak}, nil
	}
}
