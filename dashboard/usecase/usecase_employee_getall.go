// File: usecase/usecase_Employee.go

package usecase

import (
	"context"
	"dashboard/gateway"
	"dashboard/model"
	"shared/core"
	"shared/usecase"
)

type EmployeeGetAllUseCaseReq struct {
	Keyword string
	Page    int
	Size    int
}

type EmployeeGetAllUseCaseRes struct {
	Employees []model.Employee  `json:"employees"`
	Metadata  *usecase.Metadata `json:"metadata"`
}

type EmployeeGetAllUseCase = core.ActionHandler[EmployeeGetAllUseCaseReq, EmployeeGetAllUseCaseRes]

func ImplEmployeeGetAllUseCase(getAllEmployees gateway.EmployeeGetAll) EmployeeGetAllUseCase {
	return func(ctx context.Context, req EmployeeGetAllUseCaseReq) (*EmployeeGetAllUseCaseRes, error) {

		res, err := getAllEmployees(ctx, gateway.EmployeeGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &EmployeeGetAllUseCaseRes{
			Employees: res.Employees,
			Metadata: &usecase.Metadata{
				Pagination: usecase.Pagination{
					Page:       req.Page,
					Limit:      req.Size,
					TotalPages: totalPages,
					TotalItems: totalItems,
				},
			},
		}, nil
	}
}
