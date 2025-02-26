package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type RancanganPemantauanDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type RancanganPemantauanDeleteUseCaseRes struct{}

type RancanganPemantauanDeleteUseCase = core.ActionHandler[RancanganPemantauanDeleteUseCaseReq, RancanganPemantauanDeleteUseCaseRes]

func ImplRancanganPemantauanDeleteUseCase(deleteRancanganPemantauan gateway.RancanganPemantauanDelete) RancanganPemantauanDeleteUseCase {
	return func(ctx context.Context, req RancanganPemantauanDeleteUseCaseReq) (*RancanganPemantauanDeleteUseCaseRes, error) {

		if _, err := deleteRancanganPemantauan(ctx, gateway.RancanganPemantauanDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &RancanganPemantauanDeleteUseCaseRes{}, nil
	}
}
