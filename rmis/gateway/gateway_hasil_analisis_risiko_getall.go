package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/helper"
	"shared/middleware"

	"gorm.io/gorm"
)

type HasilAnalisisRisikoGetAllReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
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
		// Validate sortby
		allowedSortBy := map[string]bool{
			"risiko_prioritas": true,
		}

		sortBy, sortOrder, err := helper.ValidateSortParams(allowedSortBy, req.SortBy, req.SortOrder, "risiko_prioritas")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.HasilAnalisisRisiko

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Order(orderClause).
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
