package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type HasilAnalisisRisikoGetAllReq struct {
	Keyword string
	Page    int
	Size    int
}

type HasilAnalisisRisikoGetAllRes struct {
	HasilAnalisisRisiko []model.HasilAnalisisRisiko `json:"hasil_analisis_risikos"`
	Count               int64                       `json:"count"`
}

type HasilAnalisisRisikoGetAll = core.ActionHandler[HasilAnalisisRisikoGetAllReq, HasilAnalisisRisikoGetAllRes]

func ImplHasilAnalisisRisikoGetAll(db *gorm.DB) HasilAnalisisRisikoGetAll {
	return func(ctx context.Context, req HasilAnalisisRisikoGetAllReq) (*HasilAnalisisRisikoGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("nama LIKE ?", keyword).
				Or("kode LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.HasilAnalisisRisiko{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.HasilAnalisisRisiko

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &HasilAnalisisRisikoGetAllRes{
			HasilAnalisisRisiko: objs,
			Count:               count,
		}, nil
	}
}
