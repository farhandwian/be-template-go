// File: gateway/gateway_asset.go

package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type SimpulanKondisiKelemahanLingkunganDeleteReq struct {
	ID string
}

type SimpulanKondisiKelemahanLingkunganDeleteRes struct{}

type SimpulanKondisiKelemahanLingkunganDelete = core.ActionHandler[SimpulanKondisiKelemahanLingkunganDeleteReq, SimpulanKondisiKelemahanLingkunganDeleteRes]

func ImplSimpulanKondisiKelemahanLingkunganDelete(db *gorm.DB) SimpulanKondisiKelemahanLingkunganDelete {
	return func(ctx context.Context, req SimpulanKondisiKelemahanLingkunganDeleteReq) (*SimpulanKondisiKelemahanLingkunganDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.SimpulanKondisiKelemahanLingkungan{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &SimpulanKondisiKelemahanLingkunganDeleteRes{}, nil
	}
}
