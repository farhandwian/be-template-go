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

type PenetapanKonteksRisikoStrategisRenstraOPDGetAllReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
	Status    string
}

type PenetapanKonteksRisikoStrategisRenstraOPDGetAllRes struct {
	PenetapanKonteksRisikoStrategisRenstraOPD []model.PenetapanKonteksRisikoStrategisRenstraOPDResponse `json:"penetapan_konteks_risiko_strategis_renstra_opd"`
	Count                                     int64                                                     `json:"count"`
}

type PenetapanKonteksRisikoStrategisRenstraOPDGetAll = core.ActionHandler[PenetapanKonteksRisikoStrategisRenstraOPDGetAllReq, PenetapanKonteksRisikoStrategisRenstraOPDGetAllRes]

func ImplPenetapanKonteksRisikoStrategisRenstraOPDGetAll(db *gorm.DB) PenetapanKonteksRisikoStrategisRenstraOPDGetAll {
	return func(ctx context.Context, req PenetapanKonteksRisikoStrategisRenstraOPDGetAllReq) (*PenetapanKonteksRisikoStrategisRenstraOPDGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		query = query.
			Joins("LEFT JOIN opds ON penetapan_konteks_risiko_strategis_renstra_opds.opd_id = opds.id")

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("penetapan_konteks_risiko_strategis_renstra_opds.nama_pemda LIKE ?", keyword).
				Or("penetapan_konteks_risiko_strategis_renstra_opds.periode LIKE ?", keyword).
				Or("penetapan_konteks_risiko_strategis_renstra_opds.sumber_data LIKE ?", keyword).
				Or("opds.nama LIKE ?", keyword)
		}

		if req.Status != "" {
			query = query.Where("penetapan_konteks_risiko_strategis_renstra_opds.status =?", req.Status)
		}

		var count int64

		if err := query.
			Model(&model.PenetapanKonteksRisikoStrategisRenstraOPD{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}
		allowedSortBy := map[string]bool{
			"nama_pemda":  true,
			"periode":     true,
			"sumber_data": true,
			"status":      true,
		}

		allowerdForeignSortBy := map[string]string{
			"opd": "opds.nama",
		}

		sortBy, sortOrder, err := helper.ValidateSortParamsWithForeignKey(allowedSortBy, allowerdForeignSortBy, req.SortBy, req.SortOrder, "nama_pemda")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.PenetapanKonteksRisikoStrategisRenstraOPDResponse

		if err := query.
			Select(`penetapan_konteks_risiko_strategis_renstra_opds.*, 
                    opds.nama AS opd_nama
					`).
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Order(orderClause).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenetapanKonteksRisikoStrategisRenstraOPDGetAllRes{
			PenetapanKonteksRisikoStrategisRenstraOPD: objs,
			Count: count,
		}, nil
	}
}
