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

type DaftarRisikoPrioritasGetAllReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
}

type DaftarRisikoPrioritasGetAllRes struct {
	DaftarRisikoPrioritas []model.DaftarRisikoPrioritas `json:"daftar_risiko_prioritas"`
	Count                 int64                         `json:"count"`
}

type DaftarRisikoPrioritasGetAll = core.ActionHandler[DaftarRisikoPrioritasGetAllReq, DaftarRisikoPrioritasGetAllRes]

func ImplDaftarRisikoPrioritasGetAll(db *gorm.DB) DaftarRisikoPrioritasGetAll {
	return func(ctx context.Context, req DaftarRisikoPrioritasGetAllReq) (*DaftarRisikoPrioritasGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("risiko_prioritas LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.DaftarRisikoPrioritas{}).
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

		var objs []model.DaftarRisikoPrioritas

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Order(orderClause).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &DaftarRisikoPrioritasGetAllRes{
			DaftarRisikoPrioritas: objs,
			Count:                 count,
		}, nil
	}
}
