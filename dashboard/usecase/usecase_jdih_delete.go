package usecase

import (
	"context"
	"dashboard/gateway"
	"shared/core"
)

type DeleteJDIHReq struct {
	ID string `json:"id"`
}

type DeleteJDIHResp struct{}

type DeleteJDIHUseCase = core.ActionHandler[DeleteJDIHReq, DeleteJDIHResp]

func ImplDeleteJDIHUseCase(deleteJDIH gateway.DeleteJDIHGateway) DeleteJDIHUseCase {
	return func(ctx context.Context, req DeleteJDIHReq) (*DeleteJDIHResp, error) {

		if _, err := deleteJDIH(ctx, gateway.DeleteJDIHReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &DeleteJDIHResp{}, nil
	}
}
