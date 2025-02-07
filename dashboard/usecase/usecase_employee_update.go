package usecase

import (
	"context"
	"dashboard/gateway"
	"dashboard/model"
	"shared/core"
)

type EmployeeUpdateUseCaseReq struct {
	ID     string               `json:"id"`
	Name   string               `json:"name"`
	Role   string               `json:"role"`
	Status model.EmployeeStatus `json:"status"`
}

type EmployeeUpdateUseCaseRes struct{}

type EmployeeUpdateUseCase = core.ActionHandler[EmployeeUpdateUseCaseReq, EmployeeUpdateUseCaseRes]

func ImplEmployeeUpdateUseCase(
	getEmployeeById gateway.EmployeeGetByID,
	updateEmployee gateway.EmployeeSave,
) EmployeeUpdateUseCase {
	return func(ctx context.Context, req EmployeeUpdateUseCaseReq) (*EmployeeUpdateUseCaseRes, error) {

		res, err := getEmployeeById(ctx, gateway.EmployeeGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		res.Employee.Name = req.Name
		res.Employee.Role = req.Role
		res.Employee.Status = req.Status

		if _, err = updateEmployee(ctx, gateway.EmployeeSaveReq{Employee: res.Employee}); err != nil {
			return nil, err
		}

		return &EmployeeUpdateUseCaseRes{}, nil
	}
}
