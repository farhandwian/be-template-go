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

type PencatatanKejadianRisikoGetAllReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
	Status    string
}

type PencatatanKejadianRisikoGetAllRes struct {
	PencatatanKejadianRisiko []model.PencatatanKejadianRisikoResponse `json:"pencatatan_kejadian_risiko"`
	Count                    int64                                    `json:"count"`
}

type PencatatanKejadianRisikoGetAll = core.ActionHandler[PencatatanKejadianRisikoGetAllReq, PencatatanKejadianRisikoGetAllRes]

func ImplPencatatanKejadianRisikoGetAll(db *gorm.DB) PencatatanKejadianRisikoGetAll {
	return func(ctx context.Context, req PencatatanKejadianRisikoGetAllReq) (*PencatatanKejadianRisikoGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		query = query.
			Joins("LEFT JOIN identifikasi_risiko_strategis_pemdas ON pencatatan_kejadian_risikos.identifikasi_risiko_strategis_pemda_id = identifikasi_risiko_strategis_pemdas.id").
			Joins("LEFT JOIN penetapan_konteks_risiko_strategis_pemdas ON pencatatan_kejadian_risikos.penetapan_konteks_risiko_strategis_pemda_id = penetapan_konteks_risiko_strategis_pemdas.id")

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("penetapan_konteks_risiko_strategis_pemdas.nama_pemda LIKE ?", keyword).
				Or("penetapan_konteks_risiko_strategis_pemdas.tahun_penilaian LIKE ?", keyword).
				Or("penetapan_konteks_risiko_strategis_pemdas.tujuan_strategis LIKE ?", keyword).
				Or("penetapan_konteks_risiko_strategis_pemdas.urusan_pemerintahan LIKE ?", keyword)
		}

		if req.Status != "" {
			query = query.Where("status =?", req.Status)
		}
		var count int64

		if err := query.
			Model(&model.PencatatanKejadianRisiko{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		// Validate sortby
		allowedSortBy := map[string]bool{
			"status": true,
		}

		allowerdForeignSortBy := map[string]string{
			"nama_pemda":          "penetapan_konteks_risiko_strategis_pemdas.nama_pemda",
			"tahun_penilaian":     "penetapan_konteks_risiko_strategis_pemdas.tahun_penilaian",
			"tujuan_strategis":    "penetapan_konteks_risiko_strategis_pemdas.penetapan_tujuan",
			"urusan_pemerintahan": "penetapan_konteks_risiko_strategis_pemdas.urusan_pemerintahan",
		}

		sortBy, sortOrder, err := helper.ValidateSortParamsWithForeignKey(allowedSortBy, allowerdForeignSortBy, req.SortBy, req.SortOrder, "nama_pemda")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.PencatatanKejadianRisikoResponse

		if err := query.
			Select(`pencatatan_kejadian_risikos.*, 
                    penetapan_konteks_risiko_strategis_pemdas.nama_pemda AS nama_pemda,
					penetapan_konteks_risiko_strategis_pemdas.tahun_penilaian AS tahun_penilaian,
					penetapan_konteks_risiko_strategis_pemdas.penetapan_tujuan AS tujuan_strategis,
					penetapan_konteks_risiko_strategis_pemdas.urusan_pemerintahan AS urusan_pemerintahan,
					identifikasi_risiko_strategis_pemdas.kode_risiko AS kode_risiko
					`).
			Offset((page - 1) * size).
			Limit(size).
			Order(orderClause).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PencatatanKejadianRisikoGetAllRes{
			PencatatanKejadianRisiko: objs,
			Count:                    count,
		}, nil
	}
}
