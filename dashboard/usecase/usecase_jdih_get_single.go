package usecase

import (
	"context"
	"dashboard/gateway"
	"dashboard/model"
	"shared/core"
)

type GetJDIHByIDReq struct {
	ID string
}

type GetJDIHByIDResp struct {
	JDIH model.JDIH `json:"jdih"`
}

type GetJDIHByIDUseCase = core.ActionHandler[GetJDIHByIDReq, GetJDIHByIDResp]

func ImplGetJDIHByIDUseCase(getJDIH gateway.GetJDIHByIDGateway) GetJDIHByIDUseCase {
	return func(ctx context.Context, req GetJDIHByIDReq) (*GetJDIHByIDResp, error) {
		data, err := getJDIH(ctx, gateway.GetJDIHByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &GetJDIHByIDResp{JDIH: data.JDIH}, nil
	}
}
