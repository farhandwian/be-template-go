package usecase

import (
	"context"
	"dashboard/gateway"
	"dashboard/model"
	"fmt"
	"shared/core"
)

type ProjectCreateUseCaseReq struct {
	Name   string  `json:"name"`
	Budget float64 `json:"budget"`
}

type ProjectCreateUseCaseRes struct {
	ID string `json:"id"`
}

type ProjectCreateUseCase = core.ActionHandler[ProjectCreateUseCaseReq, ProjectCreateUseCaseRes]

func ImplProjectCreateUseCase(
	generateId gateway.GenerateId,
	createProject gateway.ProjectSave,
) ProjectCreateUseCase {
	return func(ctx context.Context, req ProjectCreateUseCaseReq) (*ProjectCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		if req.Budget <= 0 {
			return nil, fmt.Errorf("budget harus diatas Rp 0")
		}

		obj := model.Project{
			ID:     genObj.RandomId,
			Name:   req.Name,
			Budget: req.Budget,
			Status: model.ProjectStatusNotStarted,
		}

		if _, err = createProject(ctx, gateway.ProjectSaveReq{Project: obj}); err != nil {
			return nil, err
		}

		return &ProjectCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
