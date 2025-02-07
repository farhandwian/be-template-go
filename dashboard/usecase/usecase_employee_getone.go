package usecase

import (
	"context"
	"dashboard/gateway"
	"dashboard/model"
	"shared/core"
)

type EmployeeGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type EmployeeGetByIDUseCaseRes struct {
	Employee model.Employee `json:"employee"`
}

type EmployeeGetByIDUseCase = core.ActionHandler[EmployeeGetByIDUseCaseReq, EmployeeGetByIDUseCaseRes]

func ImplEmployeeGetByIDUseCase(getEmployeeByID gateway.EmployeeGetByID) EmployeeGetByIDUseCase {
	return func(ctx context.Context, req EmployeeGetByIDUseCaseReq) (*EmployeeGetByIDUseCaseRes, error) {
		res, err := getEmployeeByID(ctx, gateway.EmployeeGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &EmployeeGetByIDUseCaseRes{Employee: res.Employee}, nil
	}
}
