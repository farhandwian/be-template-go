package gateway

import (
	"context"
	"dashboard/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

// ProjectDeleteReq represents the request to delete a project
type ProjectDeleteReq struct {
	ID string
}

// ProjectDeleteRes represents the response after deleting a project
type ProjectDeleteRes struct{}

// ProjectDelete is the function signature for deleting a project
type ProjectDelete = core.ActionHandler[ProjectDeleteReq, ProjectDeleteRes]

// ImplProjectDelete implements the logic to delete a project
func ImplProjectDelete(db *gorm.DB) ProjectDelete {
	return func(ctx context.Context, req ProjectDeleteReq) (*ProjectDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.Project{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &ProjectDeleteRes{}, nil
	}
}
