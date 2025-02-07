package gateway

import (
	"context"
	"dashboard/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type ProjectSaveReq struct {
	Project model.Project
}

type ProjectSaveRes struct {
	Project model.Project
}

type ProjectSave = core.ActionHandler[ProjectSaveReq, ProjectSaveRes]

func ImplProjectSave(db *gorm.DB) ProjectSave {
	return func(ctx context.Context, req ProjectSaveReq) (*ProjectSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.Project).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &ProjectSaveRes{Project: req.Project}, nil
	}
}
