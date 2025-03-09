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
	Keyword         string
	Page            int
	Size            int
	SortBy          string
	SortOrder       string
	Status          string
	PenanggungJawab string
}

type PenilaianRisikoGetAllRes struct {
	PenilaianRisiko []model.PenilaianRisikoResponse `json:"penilaian_risiko"`
	Count           int64                           `json:"count"`
}

type PenilaianRisikoGetAll = core.ActionHandler[PenilaianRisikoGetAllReq, PenilaianRisikoGetAllRes]

func ImplPenilaianRisikoGetAll(db *gorm.DB) PenilaianRisikoGetAll {
	return func(ctx context.Context, req PenilaianRisikoGetAllReq) (*PenilaianRisikoGetAllRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		// 🔹 Join `daftar_risiko_prioritas`
		query = query.
			Joins("LEFT JOIN daftar_risiko_prioritas ON penilaian_risikos.daftar_risiko_prioritas_id = daftar_risiko_prioritas.id")

		// 🔹 Create Derived Table for Identifikasi using UNION ALL
		identifikasiUnion := `
			(
				SELECT id, kode_resiko, pemilik_resiko, uraian_resiko FROM identifikasi_risiko_strategis_pemdas
				UNION ALL
				SELECT id, kode_resiko, pemilik_resiko, uraian_resiko FROM identifikasi_risiko_operasional_opds
				UNION ALL
				SELECT id, kode_resiko, pemilik_resiko, uraian_resiko FROM identifikasi_risiko_strategis_opds
			) AS identifikasi_union
		`

		// 🔹 Create Derived Table for Penetapan Konteks using UNION ALL
		penetapanKonteksUnion := `
			(
				SELECT id, nama_pemda, tahun_penilaian, periode, penetapan_tujuan, urusan_pemerintahan 
				FROM penetapan_konteks_risiko_strategis_pemdas
				UNION ALL
				SELECT id, nama_pemda, tahun_penilaian, periode, penetapan_tujuan, urusan_pemerintahan 
				FROM penetapan_konteks_risiko_operasionals
				UNION ALL
				SELECT id, nama_pemda, tahun_penilaian, periode, penetapan_tujuan, urusan_pemerintahan 
				FROM penetapan_konteks_risiko_strategis_renstra_opds
			) AS penetapan_konteks_union
		`

		// 🔹 Join with derived tables
		query = query.
			Joins("LEFT JOIN " + identifikasiUnion + " ON daftar_risiko_prioritas.identifikasi_id = identifikasi_union.id").
			Joins("LEFT JOIN " + penetapanKonteksUnion + " ON daftar_risiko_prioritas.penetapan_konteks_id = penetapan_konteks_union.id")

		// 🔹 Filtering by keyword (Searching)
		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.Where(`
				penetapan_konteks_union.nama_pemda LIKE ? 
				OR identifikasi_union.uraian_resiko LIKE ? 
				OR identifikasi_union.kode_resiko LIKE ?
			`, keyword, keyword, keyword)
		}

		// 🔹 Filtering by Status
		if req.Status != "" {
			query = query.Where("penilaian_risikos.status =?", req.Status)
		}

		// 🔹 Count total records
		var count int64
		if err := query.Model(&model.PenilaianRisiko{}).Count(&count).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		// 🔹 Sorting Configuration
		allowedSortBy := map[string]bool{
			"pemilik_penanggung_jawab":  true,
			"target_waktu_penyelesaian": true,
			"status":                    true,
		}

		allowedForeignSortBy := map[string]string{
			"nama_pemda":      "penetapan_konteks_union.nama_pemda",
			"tahun_penilaian": "penetapan_konteks_union.tahun_penilaian",
			"uraian_resiko":   "identifikasi_union.uraian_resiko",
		}

		sortBy, sortOrder, err := helper.ValidateSortParamsWithForeignKey(allowedSortBy, allowedForeignSortBy, req.SortBy, req.SortOrder, "nama_pemda")
		if err != nil {
			return nil, err
		}

		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		// 🔹 Pagination
		page, size := ValidatePageSize(req.Page, req.Size)

		// 🔹 Fetch Data (SELECT using optimized UNION ALL joins)
		var objs []model.PenilaianRisikoResponse
		if err := query.
			Select(`
				penilaian_risikos.*, 
				penetapan_konteks_union.nama_pemda,
				penetapan_konteks_union.tahun_penilaian,
				penetapan_konteks_union.periode,
				penetapan_konteks_union.penetapan_tujuan AS tujuan,
				penetapan_konteks_union.urusan_pemerintahan,
				identifikasi_union.uraian_resiko,
				identifikasi_union.kode_resiko,
				identifikasi_union.pemilik_resiko
			`).
			Offset((page - 1) * size).
			Limit(size).
			Order(orderClause).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		// 🔹 Return Response
		return &PenilaianRisikoGetAllRes{
			PenilaianRisiko: objs,
			Count:           count,
		}, nil
	}
}
