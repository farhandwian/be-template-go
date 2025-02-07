// File: gateway/gateway_employee.go

package gateway

import (
	"context"
	"dashboard/model"
	"fmt"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type EmployeeGetByIDReq struct {
	ID string
}

type EmployeeGetByIDRes struct {
	Employee model.Employee
}

type EmployeeGetByID = core.ActionHandler[EmployeeGetByIDReq, EmployeeGetByIDRes]

func ImplEmployeeGetByID(db *gorm.DB) EmployeeGetByID {
	return func(ctx context.Context, req EmployeeGetByIDReq) (*EmployeeGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var employee model.Employee
		if err := query.First(&employee, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("employee id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &EmployeeGetByIDRes{Employee: employee}, nil
	}
}
