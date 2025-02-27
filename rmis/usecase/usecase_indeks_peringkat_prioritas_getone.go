package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type IndeksPeringkatPrioritasGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type IndeksPeringkatPrioritasGetByIDUseCaseRes struct {
	IndeksPeringkatPrioritas model.IndeksPeringkatPrioritas `json:"indeks_peringkat_prioritas"`
}

type IndeksPeringkatPrioritasGetByIDUseCase = core.ActionHandler[IndeksPeringkatPrioritasGetByIDUseCaseReq, IndeksPeringkatPrioritasGetByIDUseCaseRes]

func ImplIndeksPeringkatPrioritasGetByIDUseCase(getIndeksPeringkatPrioritasByID gateway.IndeksPeringkatPrioritasGetByID) IndeksPeringkatPrioritasGetByIDUseCase {
	return func(ctx context.Context, req IndeksPeringkatPrioritasGetByIDUseCaseReq) (*IndeksPeringkatPrioritasGetByIDUseCaseRes, error) {
		res, err := getIndeksPeringkatPrioritasByID(ctx, gateway.IndeksPeringkatPrioritasGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &IndeksPeringkatPrioritasGetByIDUseCaseRes{IndeksPeringkatPrioritas: res.IndeksPeringkatPrioritas}, nil
	}
}
