package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type RekapitulasiHasilKuesionerGetAllReq struct {
	Keyword string
	Page    int
	Size    int
}

type RekapitulasiHasilKuesionerRes struct {
	RekapitulasiHasilKuesioner []model.RekapitulasiHasilKuesioner `json:"rekapitulasi_hasil_kuesioners"`
	Count                      int64                              `json:"count"`
}

type RekapitulasiHasilKuesionerGetAll = core.ActionHandler[RekapitulasiHasilKuesionerGetAllReq, RekapitulasiHasilKuesionerRes]

func ImplRekapitulasiHasilkuesionerGetAll(db *gorm.DB) RekapitulasiHasilKuesionerGetAll {
	return func(ctx context.Context, req RekapitulasiHasilKuesionerGetAllReq) (*RekapitulasiHasilKuesionerRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.Joins("JOIN spips ON rekapitulasi_hasil_kuesioners.spip_id = spips.id").
				Where("pertanyaan LIKE ?", keyword).
				Or("spips.nama LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.RekapitulasiHasilKuesioner{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.RekapitulasiHasilKuesioner

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &RekapitulasiHasilKuesionerRes{
			RekapitulasiHasilKuesioner: objs,
			Count:                      count,
		}, nil
	}
}
