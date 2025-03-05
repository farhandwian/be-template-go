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

type RancanganPemantauanGetAllReq struct {
	Keyword         string
	Page            int
	Size            int
	SortBy          string
	SortOrder       string
	Status          string
	PenanggungJawab string
}

type RancanganPemantauanGetAllRes struct {
	RancanganPemantauan []model.RancanganPemantauanResponse `json:"rancangan_pemantauan"`
	Count               int64                               `json:"count"`
}

type RancanganPemantauanGetAll = core.ActionHandler[RancanganPemantauanGetAllReq, RancanganPemantauanGetAllRes]

func ImplRancanganPemantauanGetAll(db *gorm.DB) RancanganPemantauanGetAll {
	return func(ctx context.Context, req RancanganPemantauanGetAllReq) (*RancanganPemantauanGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		query = query.
			Joins("LEFT JOIN penilaian_risikos ON rancangan_pemantauans.penilaian_risiko_id = penilaian_risikos.id").
			Joins("LEFT JOIN daftar_risiko_prioritas ON penilaian_risikos.daftar_risiko_prioritas_id = daftar_risiko_prioritas.id").
			Joins("LEFT JOIN penetapan_konteks_risiko_strategis_pemdas ON daftar_risiko_prioritas.penetapan_konteks_risiko_strategis_pemda_id = penetapan_konteks_risiko_strategis_pemdas.id")

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("metode_pemantauan LIKE ?", keyword)
		}

		if req.Status != "" {
			query = query.Where("rancangan_pemantauans.status =?", req.Status)
		}

		if req.PenanggungJawab != "" {
			query = query.Where("rancangan_pemantauans.penanggung_jawab LIKE ?", req.PenanggungJawab)
		}

		var count int64

		if err := query.
			Model(&model.RancanganPemantauan{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		// Validate sortby
		allowedSortBy := map[string]bool{
			"pemilik_penanggung_jawab": true,
			"status":                   true,
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

		var objs []model.RancanganPemantauanResponse

		if err := query.
			Select(`rancangan_pemantauans.*, 
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

		return &RancanganPemantauanGetAllRes{
			RancanganPemantauan: objs,
			Count:               count,
		}, nil
	}
}
