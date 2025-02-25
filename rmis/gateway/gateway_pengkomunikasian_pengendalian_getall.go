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

type PengkomunikasianPengendalianGetAllReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
}

type PengkomunikasianPengendalianGetAllRes struct {
	PengkomunikasianPengendalian []model.PengkomunikasianPengendalian `json:"pengkomunikasian_pengendalian"`
	Count                        int64                                `json:"count"`
}

type PengkomunikasianPengendalianGetAll = core.ActionHandler[PengkomunikasianPengendalianGetAllReq, PengkomunikasianPengendalianGetAllRes]

func ImplPengkomunikasianPengendalianGetAll(db *gorm.DB) PengkomunikasianPengendalianGetAll {
	return func(ctx context.Context, req PengkomunikasianPengendalianGetAllReq) (*PengkomunikasianPengendalianGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("kegiatan_pengendalian LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.PengkomunikasianPengendalian{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		// Validate sortby
		allowedSortBy := map[string]bool{
			"kegiatan_pengendalian": true,
		}

		sortBy, sortOrder, err := helper.ValidateSortParams(allowedSortBy, req.SortBy, req.SortOrder, "kegiatan_pengendalian")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.PengkomunikasianPengendalian

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Order(orderClause).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PengkomunikasianPengendalianGetAllRes{
			PengkomunikasianPengendalian: objs,
			Count:                        count,
		}, nil
	}
}
