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

type IdentifikasiRisikoStrategisPemdaGetAllReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
	Status    string
	Periode   string
}

type IdentifikasiRisikoStrategisPemdaGetAllRes struct {
	IdentifikasiRisikoStrategisPemda []model.IdentifikasiRisikoStrategisPemda `json:"identifikasi_risiko_strategis_pemda"`
	Count                            int64                                    `json:"count"`
}

type IdentifikasiRisikoStrategisPemdaGetAll = core.ActionHandler[IdentifikasiRisikoStrategisPemdaGetAllReq, IdentifikasiRisikoStrategisPemdaGetAllRes]

func ImplIdentifikasiRisikoStrategisPemdaGetAll(db *gorm.DB) IdentifikasiRisikoStrategisPemdaGetAll {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisPemdaGetAllReq) (*IdentifikasiRisikoStrategisPemdaGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("nama LIKE ?", keyword).
				Or("kode LIKE ?", keyword)
		}

		if req.Status != "" {
			query = query.Where("status = ?", req.Status)
		}

		if req.Periode != "" {
			query = query.Where("periode = ?", req.Periode)
		}

		var count int64
		if err := query.
			Model(&model.IdentifikasiRisikoStrategisPemda{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}
		allowedSortBy := map[string]bool{
			"kode_risiko": true,
		}

		allowerdForeignSortBy := map[string]string{
			"identifikasi_risiko_strategis_pemda": "identifikasi_risiko_strategis_pemdas.uraian_risiko",
			"penyebab_risiko":                     "penyebab_risikos.nama",
		}

		sortBy, sortOrder, err := helper.ValidateSortParamsWithForeignKey(allowedSortBy, allowerdForeignSortBy, req.SortBy, req.SortOrder, "kode_risiko")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.IdentifikasiRisikoStrategisPemda
		if err := query.
			// Preload("KategoriRisiko").
			// Preload("Rca").
			// Preload("PenetapanKonteksRisikoStrategisPemda").
			Offset((page - 1) * size).
			Order(orderClause).
			Limit(size).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &IdentifikasiRisikoStrategisPemdaGetAllRes{
			IdentifikasiRisikoStrategisPemda: objs,
			Count:                            count,
		}, nil
	}
}
