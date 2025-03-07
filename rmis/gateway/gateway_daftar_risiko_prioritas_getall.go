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
	Keyword           string
	Page              int
	Size              int
	SortBy            string
	SortOrder         string
	Status            string
	KategoriPenilaian string
}

type DaftarRisikoPrioritasGetAllRes struct {
	DaftarRisikoPrioritas []model.DaftarRisikoPrioritasResponse `json:"daftar_risiko_prioritas"`
	Count                 int64                                 `json:"count"`
}

type DaftarRisikoPrioritasGetAll = core.ActionHandler[DaftarRisikoPrioritasGetAllReq, DaftarRisikoPrioritasGetAllRes]

func ImplDaftarRisikoPrioritasGetAll(db *gorm.DB) DaftarRisikoPrioritasGetAll {
	return func(ctx context.Context, req DaftarRisikoPrioritasGetAllReq) (*DaftarRisikoPrioritasGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

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
					daftar_risiko_prioritas.penetapan_konteks_id = penetapan_konteks.id
				)
			`)
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
			"nama_pemda":        "penetapan_konteks.nama_pemda",
			"tahun":             "penetapan_konteks.tahun",
			"tujuan_strategis":  "penetapan_konteks.penetapan_tujuan",
			"urusan_pemerintah": "penetapan_konteks.urusan_pemerintahan",
		}

		sortBy, sortOrder, err := helper.ValidateSortParamsWithForeignKey(allowedSortBy, allowerdForeignSortBy, req.SortBy, req.SortOrder, "nama_pemda")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.DaftarRisikoPrioritasResponse

		if err := query.
			Select(`
				daftar_risiko_prioritas.*, 
				penetapan_konteks.nama_pemda AS nama_pemda,
				penetapan_konteks.tahun_penilaian AS tahun,
				penetapan_konteks.periode AS periode,
				penetapan_konteks.penetapan_tujuan AS tujuan,
				penetapan_konteks.urusan_pemerintahan AS urusan_pemerintahan
			`).
			Offset((page - 1) * size).
			Limit(size).
			Order(orderClause).
			Scan(&objs). // ✅ Use Find() instead of Scan()
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &DaftarRisikoPrioritasGetAllRes{
			DaftarRisikoPrioritas: objs,
			Count:                 count,
		}, nil
	}
}
