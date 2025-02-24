package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type SimpulanKondisiKelemahanLingkunganGetAllReq struct {
	Keyword string
	Page    int
	Size    int
}

type SimpulanKondisiKelemahanLingkunganRes struct {
	SimpulanKondisiKelemahanLingkungan []model.SimpulanKondisiKelemahanLingkungan `json:"rekapitulasi_hasil_kuesioners"`
	Count                              int64                                      `json:"count"`
}

type SimpulanKondisiKelemahanLingkunganGetAll = core.ActionHandler[SimpulanKondisiKelemahanLingkunganGetAllReq, SimpulanKondisiKelemahanLingkunganRes]

func ImplSimpulanKondisiKelemahanLingkunganGetAll(db *gorm.DB) SimpulanKondisiKelemahanLingkunganGetAll {
	return func(ctx context.Context, req SimpulanKondisiKelemahanLingkunganGetAllReq) (*SimpulanKondisiKelemahanLingkunganRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("nama LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.SimpulanKondisiKelemahanLingkungan{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.SimpulanKondisiKelemahanLingkungan

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &SimpulanKondisiKelemahanLingkunganRes{
			SimpulanKondisiKelemahanLingkungan: objs,
			Count:                              count,
		}, nil
	}
}
