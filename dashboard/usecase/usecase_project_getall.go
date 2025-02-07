package usecase

import (
	"context"
	"dashboard/gateway"
	"dashboard/model"
	"shared/core"
	"shared/usecase"
)

type ProjectGetAllUseCaseReq struct {
	Keyword string
	Page    int
	Size    int
}

type ProjectGetAllUseCaseRes struct {
	Projects []model.Project   `json:"projects"`
	Metadata *usecase.Metadata `json:"metadata"`
}

type ProjectGetAllUseCase = core.ActionHandler[ProjectGetAllUseCaseReq, ProjectGetAllUseCaseRes]

func ImplProjectGetAllUseCase(getAllProjects gateway.ProjectGetAll) ProjectGetAllUseCase {
	return func(ctx context.Context, req ProjectGetAllUseCaseReq) (*ProjectGetAllUseCaseRes, error) {

		res, err := getAllProjects(ctx, gateway.ProjectGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &ProjectGetAllUseCaseRes{
			Projects: res.Projects,
			Metadata: &usecase.Metadata{
				Pagination: usecase.Pagination{
					Page:       req.Page,
					Limit:      req.Size,
					TotalPages: totalPages,
					TotalItems: totalItems,
				},
			},
		}, nil
	}
}
