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

type PengkomunikasianPengendalianGetAllReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
	Status    string
	Media     string
}

type PengkomunikasianPengendalianGetAllRes struct {
	PengkomunikasianPengendalian []model.PengkomunikasianPengendalianResponse `json:"pengkomunikasian_pengendalian"`
	Count                        int64                                        `json:"count"`
}

type PengkomunikasianPengendalianGetAll = core.ActionHandler[PengkomunikasianPengendalianGetAllReq, PengkomunikasianPengendalianGetAllRes]

func ImplPengkomunikasianPengendalianGetAll(db *gorm.DB) PengkomunikasianPengendalianGetAll {
	return func(ctx context.Context, req PengkomunikasianPengendalianGetAllReq) (*PengkomunikasianPengendalianGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		query = query.
			Joins("LEFT JOIN penilaian_risikos ON pengkomunikasian_pengendalians.penilaian_risiko_id = penilaian_risikos.id").
			Joins("LEFT JOIN daftar_risiko_prioritas ON penilaian_risikos.daftar_risiko_prioritas_id = daftar_risiko_prioritas.id").
			Joins("LEFT JOIN penetapan_konteks_risiko_strategis_pemdas ON daftar_risiko_prioritas.penetapan_konteks_risiko_strategis_pemda_id = penetapan_konteks_risiko_strategis_pemdas.id")

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("penilaian_risikos.rencana_tindak_pengendalian LIKE ?", keyword)
		}

		if req.Status != "" {
			query = query.Where("pengkomunikasian_pengendalians.status =?", req.Status)
		}

		if req.Media != "" {
			query = query.Where("pengkomunikasian_pengendalians.media =?", req.Media)
		}

		var count int64

		if err := query.
			Model(&model.PengkomunikasianPengendalian{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		// Validate sortby
		allowedSortBy := map[string]bool{
			"media_komunikasi": true,
			"status":           true,
		}

		allowerdForeignSortBy := map[string]string{
			"nama_pemda":            "penetapan_konteks_risiko_strategis_pemdas.nama_pemda",
			"tahun_penilaian":       "penetapan_konteks_risiko_strategis_pemdas.tahun_penilaian",
			"kegiatan_pengendalian": "penilaian_risikos.rencana_tindak_pengendalian",
		}

		sortBy, sortOrder, err := helper.ValidateSortParamsWithForeignKey(allowedSortBy, allowerdForeignSortBy, req.SortBy, req.SortOrder, "nama_pemda")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.PengkomunikasianPengendalianResponse

		if err := query.
			Select(`pengkomunikasian_pengendalians.*, 
                    penetapan_konteks_risiko_strategis_pemdas.nama_pemda AS nama_pemda,
					penetapan_konteks_risiko_strategis_pemdas.tahun_penilaian AS tahun_penilaian,
					penilaian_risikos.rencana_tindak_pengendalian AS rencana_tindak_pengendalian
					`).
			Offset((page - 1) * size).
			Limit(size).
			Order(orderClause).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PengkomunikasianPengendalianGetAllRes{
			PengkomunikasianPengendalian: objs,
			Count:                        count,
		}, nil
	}
}
