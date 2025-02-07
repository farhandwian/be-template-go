// File: gateway/gateway_project.go

package gateway

import (
	"context"
	"dashboard/model"
	"fmt"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type ProjectGetAllReq struct {
	Keyword string
	Page    int
	Size    int
}

type ProjectGetAllRes struct {
	Projects []model.Project `json:"projects"`
	Count    int64           `json:"count"`
}

type ProjectGetAll = core.ActionHandler[ProjectGetAllReq, ProjectGetAllRes]

func ImplProjectGetAll(db *gorm.DB) ProjectGetAll {
	return func(ctx context.Context, req ProjectGetAllReq) (*ProjectGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("name LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.Project{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var projects []model.Project

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&projects).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &ProjectGetAllRes{
			Count:    count,
			Projects: projects,
		}, nil
	}
}
