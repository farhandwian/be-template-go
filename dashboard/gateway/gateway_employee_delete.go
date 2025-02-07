package gateway

import (
	"context"
	"dashboard/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

// EmployeeDeleteReq represents the request to delete an employee
type EmployeeDeleteReq struct {
	ID string
}

// EmployeeDeleteRes represents the response after deleting an employee
type EmployeeDeleteRes struct{}

// EmployeeDelete is the function signature for deleting an employee
type EmployeeDelete = core.ActionHandler[EmployeeDeleteReq, EmployeeDeleteRes]

// ImplEmployeeDelete implements the logic to delete an employee
func ImplEmployeeDelete(db *gorm.DB) EmployeeDelete {
	return func(ctx context.Context, req EmployeeDeleteReq) (*EmployeeDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.Employee{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &EmployeeDeleteRes{}, nil
	}
}
