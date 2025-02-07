package gateway

import (
	"context"
	"dashboard/model"
	"fmt"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

// ProjectGetByIDReq represents the request to get a project by ID
type ProjectGetByIDReq struct {
	ID string
}

// ProjectGetByIDRes represents the response containing a single project
type ProjectGetByIDRes struct {
	Project model.Project
}

// ProjectGetByID is the function signature for getting a project by ID
type ProjectGetByID = core.ActionHandler[ProjectGetByIDReq, ProjectGetByIDRes]

// ImplProjectGetByID implements the logic to get a project by ID
func ImplProjectGetByID(db *gorm.DB) ProjectGetByID {
	return func(ctx context.Context, req ProjectGetByIDReq) (*ProjectGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var project model.Project
		if err := query.First(&project, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("project id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &ProjectGetByIDRes{Project: project}, nil
	}
}
