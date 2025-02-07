package usecase

import (
	"context"
	"dashboard/gateway"
	"dashboard/model"
	"shared/core"
)

type UpdateJDIHReq struct {
	ID     string           `json:"id"`
	Title  string           `json:"title"`
	Status model.JDIHStatus `json:"status"`
}

type UpdateJDIHResp struct{}

type UpdateJDIHUseCase = core.ActionHandler[UpdateJDIHReq, UpdateJDIHResp]

func ImplUpdateJDIHUseCase(
	getJDIHById gateway.GetJDIHByIDGateway,
	updateJDIH gateway.JDIHSaveGateway,
) UpdateJDIHUseCase {
	return func(ctx context.Context, req UpdateJDIHReq) (*UpdateJDIHResp, error) {

		res, err := getJDIHById(ctx, gateway.GetJDIHByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		res.JDIH.Title = req.Title
		res.JDIH.Status = req.Status

		if _, err := updateJDIH(ctx, gateway.JDIHSaveReq{JDIH: res.JDIH}); err != nil {
			return nil, err
		}
		return &UpdateJDIHResp{}, nil
	}
}
