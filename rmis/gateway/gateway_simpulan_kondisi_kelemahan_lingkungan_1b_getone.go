package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type SimpulanKondisiKelemahanLingkunganGetByIDReq struct {
	ID string
}

type SimpulanKondisiKelemahanLingkunganGetByIDRes struct {
	SimpulanKondisiKelemahanLingkungan model.SimpulanKondisiKelemahanLingkungan
}

type SimpulanKondisiKelemahanLingkunganGetByID = core.ActionHandler[SimpulanKondisiKelemahanLingkunganGetByIDReq, SimpulanKondisiKelemahanLingkunganGetByIDRes]

func ImplSimpulanKondisiKelemahanLingkunganGetByID(db *gorm.DB) SimpulanKondisiKelemahanLingkunganGetByID {
	return func(ctx context.Context, req SimpulanKondisiKelemahanLingkunganGetByIDReq) (*SimpulanKondisiKelemahanLingkunganGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var SimpulanKondisiKelemahanLingkungan model.SimpulanKondisiKelemahanLingkungan
		if err := query.First(&SimpulanKondisiKelemahanLingkungan, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("SimpulanKondisiKelemahanLingkungan id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &SimpulanKondisiKelemahanLingkunganGetByIDRes{SimpulanKondisiKelemahanLingkungan: SimpulanKondisiKelemahanLingkungan}, nil
	}
}
