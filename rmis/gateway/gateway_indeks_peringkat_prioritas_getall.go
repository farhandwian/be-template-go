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

type IndeksPeringkatPrioritasGetAllReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
}

type IndeksPeringkatPrioritasGetAllRes struct {
	IndeksPeringkatPrioritas []model.IndeksPeringkatPrioritas `json:"indeks_peringkat_prioritas"`
	Count                    int64                            `json:"count"`
}

type IndeksPeringkatPrioritasGetAll = core.ActionHandler[IndeksPeringkatPrioritasGetAllReq, IndeksPeringkatPrioritasGetAllRes]

func ImplIndeksPeringkatPrioritasGetAll(db *gorm.DB) IndeksPeringkatPrioritasGetAll {
	return func(ctx context.Context, req IndeksPeringkatPrioritasGetAllReq) (*IndeksPeringkatPrioritasGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("intermediate_rank LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.IndeksPeringkatPrioritas{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		// Validate sortby
		allowedSortBy := map[string]bool{
			"intermediate_rank": true,
		}

		sortBy, sortOrder, err := helper.ValidateSortParams(allowedSortBy, req.SortBy, req.SortOrder, "intermediate_rank")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.IndeksPeringkatPrioritas

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Order(orderClause).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &IndeksPeringkatPrioritasGetAllRes{
			IndeksPeringkatPrioritas: objs,
			Count:                    count,
		}, nil
	}
}
