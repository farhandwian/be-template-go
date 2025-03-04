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

type RcaGetAllReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
	Status    string
}

type RcaGetAllRes struct {
	Rca   []model.Rca `json:"Rcas"`
	Count int64       `json:"count"`
}

type RcaGetAll = core.ActionHandler[RcaGetAllReq, RcaGetAllRes]

func ImplRcaGetAll(db *gorm.DB) RcaGetAll {
	return func(ctx context.Context, req RcaGetAllReq) (*RcaGetAllRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.Where("nama LIKE ?", keyword)
		}

		var count int64

		if err := query.Model(&model.Rca{}).Count(&count).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		allowedSortBy := map[string]bool{
			"pemilik_risiko": true,
		}

		allowerdForeignSortBy := map[string]string{
			"identifikasi_risiko_strategis_pemda": "identifikasi_risiko_strategis_pemdas.uraian_risiko",
			"penyebab_risiko":                     "penyebab_risikos.nama",
		}

		sortBy, sortOrder, err := helper.ValidateSortParamsWithForeignKey(allowedSortBy, allowerdForeignSortBy, req.SortBy, req.SortOrder, "pemilik_risiko")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.Rca

		// Select with proper JOINs and aliases
		if err := query.
			Joins("LEFT JOIN identifikasi_risiko_strategis_pemdas ON rcas.identifikasi_risiko_strategis_pemda_id = identifikasi_risiko_strategis_pemdas.id").
			Joins("LEFT JOIN penyebab_risikos ON rcas.penyebab_risiko_id = penyebab_risikos.id").
			Select(`
			rcas.*, 
			identifikasi_risiko_strategis_pemdas.uraian_risiko AS identifikasi_risiko_uraian_risiko,
			penyebab_risikos.nama AS penyebab_risiko_nama
		`).
			Offset((page - 1) * size).
			Limit(size).
			Order(orderClause).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		fmt.Println(query)
		return &RcaGetAllRes{
			Rca:   objs,
			Count: count,
		}, nil
	}
}
