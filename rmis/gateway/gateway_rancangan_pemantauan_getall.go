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

type RancanganPemantauanGetAllReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
}

type RancanganPemantauanGetAllRes struct {
	RancanganPemantauan []model.RancanganPemantauan `json:"rancangan_pemantauan"`
	Count               int64                       `json:"count"`
}

type RancanganPemantauanGetAll = core.ActionHandler[RancanganPemantauanGetAllReq, RancanganPemantauanGetAllRes]

func ImplRancanganPemantauanGetAll(db *gorm.DB) RancanganPemantauanGetAll {
	return func(ctx context.Context, req RancanganPemantauanGetAllReq) (*RancanganPemantauanGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("metode_pemantauan LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.RancanganPemantauan{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		// Validate sortby
		allowedSortBy := map[string]bool{
			"metode_pemantauan": true,
		}

		sortBy, sortOrder, err := helper.ValidateSortParams(allowedSortBy, req.SortBy, req.SortOrder, "metode_pemantauan")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.RancanganPemantauan

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Order(orderClause).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &RancanganPemantauanGetAllRes{
			RancanganPemantauan: objs,
			Count:               count,
		}, nil
	}
}
