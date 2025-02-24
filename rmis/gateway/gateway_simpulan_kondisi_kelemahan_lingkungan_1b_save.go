package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type SimpulanKondisiKelemahanLingkunganSaveReq struct {
	SimpulanKondisiKelemahanLingkungan model.SimpulanKondisiKelemahanLingkungan
}

type SimpulanKondisiKelemahanLingkunganSaveRes struct {
	ID string
}

type SimpulanKondisiKelemahanLingkunganSave = core.ActionHandler[SimpulanKondisiKelemahanLingkunganSaveReq, SimpulanKondisiKelemahanLingkunganSaveRes]

func ImplSimpulanKondisiKelemahanLingkunganSave(db *gorm.DB) SimpulanKondisiKelemahanLingkunganSave {
	return func(ctx context.Context, req SimpulanKondisiKelemahanLingkunganSaveReq) (*SimpulanKondisiKelemahanLingkunganSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.SimpulanKondisiKelemahanLingkungan).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &SimpulanKondisiKelemahanLingkunganSaveRes{ID: *req.SimpulanKondisiKelemahanLingkungan.ID}, nil
	}
}
