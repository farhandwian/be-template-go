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
	Status    string
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

		if req.Status != "" {
			query = query.Where("status =?", req.Status)
		}

		var count int64

		if err := query.
			Model(&model.DaftarRisikoPrioritas{}).
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
			"urusan_pemerintah": "penetapan_konteks_risiko_strategis_pemdas.urusan_pemerintahan",
		}

		sortBy, sortOrder, err := helper.ValidateSortParamsWithForeignKey(allowedSortBy, allowerdForeignSortBy, req.SortBy, req.SortOrder, "nama_pemda")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.DaftarRisikoPrioritas

		if err := query.
			Select(`daftar_risiko_prioritas.*, 
				penetapan_konteks_risiko_strategis_pemdas.nama_pemda AS nama_pemda,
				penetapan_konteks_risiko_strategis_pemdas.tahun_penilaian AS tahun,
				penetapan_konteks_risiko_strategis_pemdas.periode AS periode,
				penetapan_konteks_risiko_strategis_pemdas.penetapan_tujuan AS tujuan,
				penetapan_konteks_risiko_strategis_pemdas.urusan_pemerintahan AS urusan_pemerintahan,
				penetapan_konteks_risiko_strategis_pemdas.penetapan_tujuan AS penetapan_konteks,
				hasil_analisis_risikos.skala_risiko AS skala_risiko
			`).
			Joins("LEFT JOIN penetapan_konteks_risiko_strategis_pemdas ON daftar_risiko_prioritas.penetapan_konteks_risiko_strategis_pemda_id = penetapan_konteks_risiko_strategis_pemdas.id").
			Joins("LEFT JOIN indeks_peringkat_prioritas ON daftar_risiko_prioritas.indeks_peringkat_prioritas_id = indeks_peringkat_prioritas.id").
			Joins("LEFT JOIN hasil_analisis_risikos ON daftar_risiko_prioritas.hasil_analisis_risiko_id = hasil_analisis_risikos.id").
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
