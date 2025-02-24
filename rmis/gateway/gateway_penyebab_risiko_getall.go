package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenyebabRisikoGetAllReq struct {
	Keyword string
	Page    int
	Size    int
}

type PenyebabRisikoGetAllRes struct {
	PenyebabRisiko []model.PenyebabRisiko `json:"PenyebabRisikos"`
	Count          int64                  `json:"count"`
}

type PenyebabRisikoGetAll = core.ActionHandler[PenyebabRisikoGetAllReq, PenyebabRisikoGetAllRes]

func ImplPenyebabRisikoGetAll(db *gorm.DB) PenyebabRisikoGetAll {
	return func(ctx context.Context, req PenyebabRisikoGetAllReq) (*PenyebabRisikoGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("nama LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.PenyebabRisiko{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.PenyebabRisiko

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenyebabRisikoGetAllRes{
			PenyebabRisiko: objs,
			Count:          count,
		}, nil
	}
}
