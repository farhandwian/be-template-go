// File: usecase/usecase_Project.go

package usecase

import (
	"context"
	"dashboard/gateway"
	"dashboard/model"
	"shared/core"
)

type ProjectGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type ProjectGetByIDUseCaseRes struct {
	Project model.Project `json:"project"`
}

type ProjectGetByIDUseCase = core.ActionHandler[ProjectGetByIDUseCaseReq, ProjectGetByIDUseCaseRes]

func ImplProjectGetByIDUseCase(getProjectByID gateway.ProjectGetByID) ProjectGetByIDUseCase {
	return func(ctx context.Context, req ProjectGetByIDUseCaseReq) (*ProjectGetByIDUseCaseRes, error) {
		res, err := getProjectByID(ctx, gateway.ProjectGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &ProjectGetByIDUseCaseRes{Project: res.Project}, nil
	}
}
