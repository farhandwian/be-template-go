package usecase

import (
	"context"
	"dashboard/gateway"
	"shared/core"
)

type EmployeeDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type EmployeeDeleteUseCaseRes struct{}

type EmployeeDeleteUseCase = core.ActionHandler[EmployeeDeleteUseCaseReq, EmployeeDeleteUseCaseRes]

func ImplEmployeeDeleteUseCase(deleteEmployee gateway.EmployeeDelete) EmployeeDeleteUseCase {
	return func(ctx context.Context, req EmployeeDeleteUseCaseReq) (*EmployeeDeleteUseCaseRes, error) {

		if _, err := deleteEmployee(ctx, gateway.EmployeeDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &EmployeeDeleteUseCaseRes{}, nil
	}
}
