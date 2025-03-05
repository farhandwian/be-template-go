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

type PenilaianKegiatanPengendalianGetAllReq struct {
	Keyword           string
	Page              int
	Size              int
	SortBy            string
	SortOrder         string
	Status            string
	KategoriPenilaian string
}

type PenilaianKegiatanPengendalianGetAllRes struct {
	PenilaianKegiatanPengendalian []model.PenilaianKegiatanPengendalianResponse `json:"penilai_kegiatan_pengendalians"`
	Count                         int64                                         `json:"count"`
}

type PenilaianKegiatanPengendalianGetAll = core.ActionHandler[PenilaianKegiatanPengendalianGetAllReq, PenilaianKegiatanPengendalianGetAllRes]

func ImplPenilaianKegiatanPengendalianGetAll(db *gorm.DB) PenilaianKegiatanPengendalianGetAll {
	return func(ctx context.Context, req PenilaianKegiatanPengendalianGetAllReq) (*PenilaianKegiatanPengendalianGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		query = query.
			Joins("LEFT JOIN spips ON penilaian_kegiatan_pengendalians.spip_id = spips.id")

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("penilaian_kegiatan_pengendalians.nama_pemda LIKE ?", keyword).
				Or("penilaian_kegiatan_pengendalians.tahun_penilaian LIKE ?", keyword).
				Or("spips.nama LIKE ?", keyword).
				Or("penilaian_kegiatan_pengendalians.target_waktu_penyelesaian LIKE ?", keyword)
		}

		if req.Status != "" {
			query = query.Where("penilaian_kegiatan_pengendalians.status =?", req.Status)
		}

		if req.KategoriPenilaian != "" {
			fmt.Println(req.KategoriPenilaian)
			query = query.Where("spips.nama =?", req.KategoriPenilaian)
		}

		var count int64

		if err := query.
			Model(&model.PenilaianKegiatanPengendalian{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}
		// Validate sortby
		allowedSortBy := map[string]bool{
			"nama_pemda":                true,
			"tahun_penilaian":           true,
			"target_waktu_penyelesaian": true,
			"status":                    true,
		}

		allowerdForeignSortBy := map[string]string{
			"kategori_penilaian": "spips.nama",
		}

		sortBy, sortOrder, err := helper.ValidateSortParamsWithForeignKey(allowedSortBy, allowerdForeignSortBy, req.SortBy, req.SortOrder, "nama_pemda")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.PenilaianKegiatanPengendalianResponse

		if err := query.
			Select(`penilaian_kegiatan_pengendalians.*,
			spips.nama AS spip_nama
			`).
			Offset((page - 1) * size).
			Limit(size).
			Order(orderClause).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenilaianKegiatanPengendalianGetAllRes{
			PenilaianKegiatanPengendalian: objs,
			Count:                         count,
		}, nil
	}
}
