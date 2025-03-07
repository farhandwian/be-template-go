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
	HasilAnalisisRisiko []model.HasilAnalisisRisikoResponse `json:"hasil_analisis_risikos"`
	Count               int64                               `json:"count"`
}

type HasilAnalisisRisikoGetAll = core.ActionHandler[HasilAnalisisRisikoGetAllReq, HasilAnalisisRisikoGetAllRes]

func ImplHasilAnalisisRisikoGetAll(db *gorm.DB) HasilAnalisisRisikoGetAll {
	return func(ctx context.Context, req HasilAnalisisRisikoGetAllReq) (*HasilAnalisisRisikoGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		// 🔹 Efficiently JOIN Identifikasi Risiko Tables
		query = query.Joins(`
					LEFT JOIN identifikasi_risiko_strategis_pemdas ON 
						hasil_analisis_risikos.identifikasi_id = identifikasi_risiko_strategis_pemdas.id 
						AND hasil_analisis_risikos.tipe_identifikasi = 'strategis_pemda'
					LEFT JOIN identifikasi_risiko_operasional_opds ON 
						hasil_analisis_risikos.identifikasi_id = identifikasi_risiko_operasional_opds.id 
						AND hasil_analisis_risikos.tipe_identifikasi = 'operasional_opd'
					LEFT JOIN identifikasi_risiko_strategis_opds ON 
						hasil_analisis_risikos.identifikasi_id = identifikasi_risiko_strategis_opds.id 
						AND hasil_analisis_risikos.tipe_identifikasi = 'strategis_opd'
				`)

		// 🔹 Use UNION ALL for Penetapan Konteks (Only One Will Be Filled)
		query = query.Joins(`
				LEFT JOIN (
					SELECT id, nama_pemda, tahun_penilaian, periode, penetapan_tujuan, urusan_pemerintahan 
					FROM penetapan_konteks_risiko_strategis_pemdas 
					UNION ALL 
					SELECT id, nama_pemda, tahun_penilaian, periode, tujuan_strategis AS penetapan_tujuan, urusan_pemerintahan 
					FROM penetapan_konteks_risiko_operasionals 
					UNION ALL 
					SELECT id, nama_pemda, tahun_penilaian, periode, penetapan_tujuan, urusan_pemerintahan 
					FROM penetapan_konteks_risiko_strategis_renstra_opds
				) AS penetapan_konteks ON (
					hasil_analisis_risikos.penetapan_konteks_id = penetapan_konteks.id
				)
			`)

		// 🔹 Searching Optimization (COALESCE Ensures Non-NULL Values)
		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.Where(`
		COALESCE(penetapan_konteks.nama_pemda, '') LIKE ? 
		OR COALESCE(penetapan_konteks.tahun_penilaian, '') LIKE ? 
		OR COALESCE(penetapan_konteks.periode, '') LIKE ?`,
				keyword, keyword, keyword)
		}

		// 🔹 Status Filtering
		if req.Status != "" {
			query = query.Where("hasil_analisis_risikos.status = ?", req.Status)
		}

		// 🔹 Count Query (Faster)
		var count int64
		if err := query.Model(&model.HasilAnalisisRisiko{}).Count(&count).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		// 🔹 Sorting Configuration
		allowedSortBy := map[string]bool{}
		allowedForeignSortBy := map[string]string{
			"nama_pemda":        "penetapan_konteks.nama_pemda",
			"tahun":             "penetapan_konteks.tahun_penilaian",
			"periode":           "penetapan_konteks.periode",
			"tujuan_strategis":  "penetapan_konteks.tujuan",
			"urusan_pemerintah": "penetapan_konteks.urusan_pemerintahan",
		}

		sortBy, sortOrder, err := helper.ValidateSortParamsWithForeignKey(allowedSortBy, allowedForeignSortBy, req.SortBy, req.SortOrder, "nama_pemda")
		if err != nil {
			return nil, err
		}

		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		// 🔹 Pagination
		page, size := ValidatePageSize(req.Page, req.Size)

		// 🔹 Fetch Data
		var objs []model.HasilAnalisisRisikoResponse
		if err := query.
			Select(`
		hasil_analisis_risikos.*, 
		penetapan_konteks.nama_pemda AS nama_pemda,
		penetapan_konteks.tahun_penilaian AS tahun,
		penetapan_konteks.periode AS periode,
		penetapan_konteks.penetapan_tujuan AS tujuan,
		penetapan_konteks.urusan_pemerintahan AS urusan_pemerintahan
	`).
			Offset((page - 1) * size).
			Limit(size).
			Order(orderClause).
			Scan(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &HasilAnalisisRisikoGetAllRes{
			HasilAnalisisRisiko: objs,
			Count:               count,
		}, nil

	}
}
