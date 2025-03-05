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
	Status    string
}

type HasilAnalisisRisikoGetAllRes struct {
	HasilAnalisisRisiko []model.HasilAnalisisRisiko `json:"hasil_analisis_risikos"`
	Count               int64                       `json:"count"`
}

type HasilAnalisisRisikoGetAll = core.ActionHandler[HasilAnalisisRisikoGetAllReq, HasilAnalisisRisikoGetAllRes]

func ImplHasilAnalisisRisikoGetAll(db *gorm.DB) HasilAnalisisRisikoGetAll {
	return func(ctx context.Context, req HasilAnalisisRisikoGetAllReq) (*HasilAnalisisRisikoGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		query = query.
			Joins("LEFT JOIN penetapan_konteks_risiko_strategis_pemdas ON hasil_analisis_risikos.penetapan_konteks_risiko_strategis_pemda_id = penetapan_konteks_risiko_strategis_pemdas.id").
			Joins("LEFT JOIN identifikasi_risiko_strategis_pemdas ON hasil_analisis_risikos.identifikasi_risiko_strategis_pemda_id = identifikasi_risiko_strategis_pemdas.id")

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("nama LIKE ?", keyword).
				Or("kode LIKE ?", keyword)
		}

		if req.Status != "" {
			query = query.Where("hasil_analisis_risikos.status =?", req.Status)
		}

		var count int64

		if err := query.
			Model(&model.HasilAnalisisRisiko{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}
		// Validate sortby
		allowedSortBy := map[string]bool{}

		allowerdForeignSortBy := map[string]string{
			"nama_pemda":        "penetapan_konteks_risiko_strategis_pemdas.nama_pemda",
			"tahun":             "penetapan_konteks_risiko_strategis_pemdas.tahun",
			"tujuan_strategis":  "penetapan_konteks_risiko_strategis_pemdas.penetapan_tujuan",
			"urusan_pemerintah": "penetapan_konteks_risiko_strategis_pemdas.urusan_pemerintah",
		}

		sortBy, sortOrder, err := helper.ValidateSortParamsWithForeignKey(allowedSortBy, allowerdForeignSortBy, req.SortBy, req.SortOrder, "nama_pemda")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.HasilAnalisisRisiko

		if err := query.
			Select(`hasil_analisis_risikos.*, 
				penetapan_konteks_risiko_strategis_pemdas.nama_pemda AS nama_pemda,
				penetapan_konteks_risiko_strategis_pemdas.tahun_penilaian AS tahun,
				penetapan_konteks_risiko_strategis_pemdas.periode AS periode,
				penetapan_konteks_risiko_strategis_pemdas.penetapan_tujuan AS tujuan,
				penetapan_konteks_risiko_strategis_pemdas.urusan_pemerintahan AS urusan_pemerintahan,
				penetapan_konteks_risiko_strategis_pemdas.penetapan_tujuan AS penetapan_konteks
			`).
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
