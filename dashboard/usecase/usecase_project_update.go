package usecase

import (
	"context"
	"dashboard/gateway"
	"dashboard/model"
	"shared/core"
)

type ProjectUpdateUseCaseReq struct {
	ID     string              `json:"id"`
	Name   string              `json:"name"`
	Budget float64             `json:"budget"`
	Status model.ProjectStatus `json:"status"`
}

type ProjectUpdateUseCaseRes struct{}

type ProjectUpdateUseCase = core.ActionHandler[ProjectUpdateUseCaseReq, ProjectUpdateUseCaseRes]

func ImplProjectUpdateUseCase(
	getProjectById gateway.ProjectGetByID,
	updateProject gateway.ProjectSave,
) ProjectUpdateUseCase {
	return func(ctx context.Context, req ProjectUpdateUseCaseReq) (*ProjectUpdateUseCaseRes, error) {

		res, err := getProjectById(ctx, gateway.ProjectGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		res.Project.Name = req.Name
		res.Project.Budget = req.Budget
		res.Project.Status = req.Status

		if _, err = updateProject(ctx, gateway.ProjectSaveReq{Project: res.Project}); err != nil {
			return nil, err
		}

		return &ProjectUpdateUseCaseRes{}, nil
	}
}
