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

type PenilaianRisikoGetAllReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
}

type PenilaianRisikoGetAllRes struct {
	PenilaianRisiko []model.PenilaianRisiko `json:"penilaian_risiko"`
	Count           int64                   `json:"count"`
}

type PenilaianRisikoGetAll = core.ActionHandler[PenilaianRisikoGetAllReq, PenilaianRisikoGetAllRes]

func ImplPenilaianRisikoGetAll(db *gorm.DB) PenilaianRisikoGetAll {
	return func(ctx context.Context, req PenilaianRisikoGetAllReq) (*PenilaianRisikoGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("risiko_prioritas LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.PenilaianRisiko{}).
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

		var objs []model.PenilaianRisiko

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Order(orderClause).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenilaianRisikoGetAllRes{
			PenilaianRisiko: objs,
			Count:           count,
		}, nil
	}
}
