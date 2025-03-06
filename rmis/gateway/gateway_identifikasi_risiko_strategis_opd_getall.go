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

type IdentifikasiRisikoStrategisOPDGetAllReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
	Status    string
}

type IdentifikasiRisikoStrategisOPDGetAllRes struct {
	IdentifikasiRisikoStrategisOPD []model.IdentifikasiRisikoStrategisOPDResponse `json:"identifikasi_risiko_strategis_opd"`
	Count                          int64                                          `json:"count"`
}

type IdentifikasiRisikoStrategisOPDGetAll = core.ActionHandler[IdentifikasiRisikoStrategisOPDGetAllReq, IdentifikasiRisikoStrategisOPDGetAllRes]

func ImplIdentifikasiRisikoStrategisOPDGetAll(db *gorm.DB) IdentifikasiRisikoStrategisOPDGetAll {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisOPDGetAllReq) (*IdentifikasiRisikoStrategisOPDGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		query = query.
			Joins("LEFT JOIN penetapan_konteks_risiko_strategis_renstra_opds ON identifikasi_risiko_strategis_opds.penetapan_konteks_risiko_strategis_renstra_id = penetapan_konteks_risiko_strategis_renstra_opds.id").
			Joins("LEFT JOIN kategori_risikos ON identifikasi_risiko_strategis_opds.kategori_risiko_id = kategori_risikos.id")

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("penetapan_konteks_risiko_strategis_renstra_opds.nama_pemda LIKE ?", keyword).
				Or("penetapan_konteks_risiko_strategis_renstra_opds.tahun_penilaian LIKE ?", keyword).
				Or("penetapan_konteks_risiko_strategis_renstra_opds.penetapan_tujuan LIKE ?", keyword).
				Or("penetapan_konteks_risiko_strategis_renstra_opds.periode LIKE ?", keyword).
				Or("penetapan_konteks_risiko_strategis_renstra_opds.urusan_pemerintahan LIKE?", keyword)
		}

		if req.Status != "" {
			query = query.Where("identifikasi_risiko_strategis_opds.status =?", req.Status)
		}

		var count int64

		if err := query.
			Model(&model.IdentifikasiRisikoStrategisOPD{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}
		// Validate sortby
		allowedSortBy := map[string]bool{
			"status": true,
		}

		allowerdForeignSortBy := map[string]string{
			"nama_pemda":        "penetapan_konteks_risiko_strategis_renstra_opds.nama_pemda",
			"tahun":             "penetapan_konteks_risiko_strategis_renstra_opds.tahun_penilaian",
			"periode":           "penetapan_konteks_risiko_strategis_renstra_opds.periode",
			"penetapan_konteks": "penetapan_konteks_risiko_strategis_renstra_opds.penetapan_tujuan",
			"urusan_pemerintah": "penetapan_konteks_risiko_strategis_renstra_opds.urusan_pemerintahan",
		}

		sortBy, sortOrder, err := helper.ValidateSortParamsWithForeignKey(allowedSortBy, allowerdForeignSortBy, req.SortBy, req.SortOrder, "nama_pemda")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.IdentifikasiRisikoStrategisOPDResponse

		if err := query.
			Select(`identifikasi_risiko_strategis_opds.*, 
				penetapan_konteks_risiko_strategis_renstra_opds.nama_pemda AS nama_pemda,
				penetapan_konteks_risiko_strategis_renstra_opds.tahun_penilaian AS tahun,
				penetapan_konteks_risiko_strategis_renstra_opds.periode AS periode,
				penetapan_konteks_risiko_strategis_renstra_opds.urusan_pemerintahan AS urusan_pemerintahan,
				penetapan_konteks_risiko_strategis_renstra_opds.penetapan_tujuan AS penetapan_konteks
			`).
			Offset((page - 1) * size).
			Limit(size).
			Order(orderClause).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &IdentifikasiRisikoStrategisOPDGetAllRes{
			IdentifikasiRisikoStrategisOPD: objs,
			Count:                          count,
		}, nil
	}
}
