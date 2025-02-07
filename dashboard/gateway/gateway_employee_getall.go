package gateway

import (
	"context"
	"dashboard/model"
	"fmt"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type EmployeeGetAllReq struct {
	Keyword string
	Page    int
	Size    int
}

type EmployeeGetAllRes struct {
	Employees []model.Employee `json:"employees"`
	Count     int64            `json:"count"`
}

type EmployeeGetAll = core.ActionHandler[EmployeeGetAllReq, EmployeeGetAllRes]

func ImplEmployeeGetAll(db *gorm.DB) EmployeeGetAll {
	return func(ctx context.Context, req EmployeeGetAllReq) (*EmployeeGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("name LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.Employee{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var employees []model.Employee

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&employees).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &EmployeeGetAllRes{
			Count:     count,
			Employees: employees,
		}, nil
	}
}
