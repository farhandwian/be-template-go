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

		query = query.
			Joins("LEFT JOIN daftar_risiko_prioritas ON penilaian_risikos.daftar_risiko_prioritas_id = daftar_risiko_prioritas.id").
			Joins("LEFT JOIN penetapan_konteks_risiko_strategis_pemdas ON daftar_risiko_prioritas.penetapan_konteks_risiko_strategis_pemda_id = penetapan_konteks_risiko_strategis_pemdas.id")
			// Joins("LEFT JOIN ").
			// Joins("LEFT JOIN identifikasi_risiko_strategis_pemdas ON ")

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("penetapan_konteks_risiko_strategis_pemdas.nama_pemda LIKE ?", keyword).
				Or("penetapan_konteks_risiko_strategis_pemdas.tahun_penilaian LIKE ?", keyword).
				Or("penilaian_risikos.pemilik_penanggung_jawab LIKE ?", keyword).
				Or("penilaian_risikos.target_waktu_penyelesaian LIKE ?", keyword)
		}

		var count int64
		if req.Status != "" {
			query = query.Where("penilaian_risikos.status =?", req.Status)
		}

		if err := query.
			Model(&model.PenilaianRisiko{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		// Validate sortby
		allowedSortBy := map[string]bool{
			"pemilik_penanggung_jawab":  true,
			"target_waktu_penyelesaian": true,
			"status":                    true,
		}

		allowerdForeignSortBy := map[string]string{
			"nama_pemda":      "penetapan_konteks_risiko_strategis_pemdas.nama_pemda",
			"tahun_penilaian": "penetapan_konteks_risiko_strategis_pemdas.tahun_penilaian",
		}

		sortBy, sortOrder, err := helper.ValidateSortParamsWithForeignKey(allowedSortBy, allowerdForeignSortBy, req.SortBy, req.SortOrder, "nama_pemda")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.PenilaianRisikoResponse

		if err := query.
			Select(`penilaian_risikos.*, 
                    penetapan_konteks_risiko_strategis_pemdas.nama_pemda AS nama_pemda,
					penetapan_konteks_risiko_strategis_pemdas.tahun_penilaian AS tahun_penilaian
					`).
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
