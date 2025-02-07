package usecase

import (
	"context"
	"dashboard/gateway"
	"shared/core"
)

type ProjectDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type ProjectDeleteUseCaseRes struct{}

type ProjectDeleteUseCase = core.ActionHandler[ProjectDeleteUseCaseReq, ProjectDeleteUseCaseRes]

func ImplProjectDeleteUseCase(deleteProject gateway.ProjectDelete) ProjectDeleteUseCase {
	return func(ctx context.Context, req ProjectDeleteUseCaseReq) (*ProjectDeleteUseCaseRes, error) {

		if _, err := deleteProject(ctx, gateway.ProjectDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &ProjectDeleteUseCaseRes{}, nil

	}
}
