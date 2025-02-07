package gateway

import (
	"context"
	"dashboard/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type EmployeeSaveReq struct {
	Employee model.Employee
}

type EmployeeSaveRes struct {
	ID string
}

type EmployeeSave = core.ActionHandler[EmployeeSaveReq, EmployeeSaveRes]

func ImplEmployeeSave(db *gorm.DB) EmployeeSave {
	return func(ctx context.Context, req EmployeeSaveReq) (*EmployeeSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.Employee).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &EmployeeSaveRes{ID: req.Employee.ID}, nil
	}
}
