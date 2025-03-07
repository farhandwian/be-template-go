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

type IdentifikasiRisikoOperasionalOPDGetAllReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
	Status    string
}

type IdentifikasiRisikoOperasionalOPDGetAllRes struct {
	IdentifikasiRisikoOperasionalOPD []model.IdentifikasiRisikoOperasionalOPDResponse `json:"identifikasi_risiko_operasional_opd"`
	Count                            int64                                            `json:"count"`
}

type IdentifikasiRisikoOperasionalOPDGetAll = core.ActionHandler[IdentifikasiRisikoOperasionalOPDGetAllReq, IdentifikasiRisikoOperasionalOPDGetAllRes]

func ImplIdentifikasiRisikoOperasionalOPDGetAll(db *gorm.DB) IdentifikasiRisikoOperasionalOPDGetAll {
	return func(ctx context.Context, req IdentifikasiRisikoOperasionalOPDGetAllReq) (*IdentifikasiRisikoOperasionalOPDGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		query = query.
			Joins("LEFT JOIN penetapan_konteks_risiko_operasionals ON identifikasi_risiko_operasional_opds.penetapan_konteks_risiko_operasional_opd_id = penetapan_konteks_risiko_operasionals.id").
			Joins("LEFT JOIN kategori_risikos ON identifikasi_risiko_operasional_opds.kategori_risiko_id = kategori_risikos.id")

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("penetapan_konteks_risiko_operasionals.nama_pemda LIKE ?", keyword).
				Or("penetapan_konteks_risiko_operasionals.tahun_penilaian LIKE ?", keyword).
				Or("penetapan_konteks_risiko_operasionals.tujuan_strategis LIKE ?", keyword).
				Or("penetapan_konteks_risiko_operasionals.periode LIKE ?", keyword).
				Or("penetapan_konteks_risiko_operasionals.urusan_pemerintahan LIKE?", keyword)
		}

		if req.Status != "" {
			query = query.Where("identifikasi_risiko_operasional_opds.status =?", req.Status)
		}

		var count int64

		if err := query.
			Model(&model.IdentifikasiRisikoOperasionalOPD{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		allowedSortBy := map[string]bool{
			"status": true,
		}

		allowerdForeignSortBy := map[string]string{
			"nama_pemda":        "penetapan_konteks_risiko_operasionals.nama_pemda",
			"tahun":             "penetapan_konteks_risiko_operasionals.tahun_penilaian",
			"periode":           "penetapan_konteks_risiko_operasionals.periode",
			"penetapan_konteks": "penetapan_konteks_risiko_operasionals.tujuan_strategis",
			"urusan_pemerintah": "penetapan_konteks_risiko_operasionals.urusan_pemerintahan",
		}

		sortBy, sortOrder, err := helper.ValidateSortParamsWithForeignKey(allowedSortBy, allowerdForeignSortBy, req.SortBy, req.SortOrder, "nama_pemda")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.IdentifikasiRisikoOperasionalOPDResponse

		if err := query.
			Select(`identifikasi_risiko_operasional_opds.*, 
				penetapan_konteks_risiko_operasionals.nama_pemda AS nama_pemda,
				penetapan_konteks_risiko_operasionals.tahun_penilaian AS tahun,
				penetapan_konteks_risiko_operasionals.periode AS periode,
				penetapan_konteks_risiko_operasionals.urusan_pemerintahan AS urusan_pemerintahan,
				penetapan_konteks_risiko_operasionals.tujuan_strategis AS penetapan_konteks
			`).
			Offset((page - 1) * size).
			Limit(size).
			Order(orderClause).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &IdentifikasiRisikoOperasionalOPDGetAllRes{
			IdentifikasiRisikoOperasionalOPD: objs,
			Count:                            count,
		}, nil
	}
}
