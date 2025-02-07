package usecase

import (
	"context"
	"dashboard/gateway"
	"dashboard/model"
	"fmt"
	"shared/core"
	"time"
)

type EmployeeCreateUseCaseReq struct {
	Name       string    `json:"name"`
	Role       string    `json:"role"`
	JoinedDate time.Time `json:"joined_date"`
	Now        time.Time
}

type EmployeeCreateUseCaseRes struct {
	ID string `json:"id"`
}

type EmployeeCreateUseCase = core.ActionHandler[EmployeeCreateUseCaseReq, EmployeeCreateUseCaseRes]

func ImplEmployeeCreateUseCase(generateId gateway.GenerateId, createEmployee gateway.EmployeeSave) EmployeeCreateUseCase {
	return func(ctx context.Context, req EmployeeCreateUseCaseReq) (*EmployeeCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		if !req.JoinedDate.Before(req.Now) {
			return nil, fmt.Errorf("tanggal join tidak boleh di masa mendatang")
		}

		obj := model.Employee{
			ID:         genObj.RandomId,
			Name:       req.Name,
			Role:       req.Role,
			JoinedDate: req.JoinedDate,
			Status:     model.EmployeeStatusActive,
		}

		if _, err := createEmployee(ctx, gateway.EmployeeSaveReq{Employee: obj}); err != nil {
			return nil, err
		}

		return &EmployeeCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil

	}
}
