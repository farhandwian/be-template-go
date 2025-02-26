package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type RancanganPemantauanGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type RancanganPemantauanGetByIDUseCaseRes struct {
	RancanganPemantauan model.RancanganPemantauan `json:"rancangan_pemantauan"`
}

type RancanganPemantauanGetByIDUseCase = core.ActionHandler[RancanganPemantauanGetByIDUseCaseReq, RancanganPemantauanGetByIDUseCaseRes]

func ImplRancanganPemantauanGetByIDUseCase(getRancanganPemantauanByID gateway.RancanganPemantauanGetByID) RancanganPemantauanGetByIDUseCase {
	return func(ctx context.Context, req RancanganPemantauanGetByIDUseCaseReq) (*RancanganPemantauanGetByIDUseCaseRes, error) {
		res, err := getRancanganPemantauanByID(ctx, gateway.RancanganPemantauanGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &RancanganPemantauanGetByIDUseCaseRes{RancanganPemantauan: res.RancanganPemantauan}, nil
	}
}
