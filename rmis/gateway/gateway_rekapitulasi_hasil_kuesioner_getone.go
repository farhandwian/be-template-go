package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type RekapitulasiHasilKuesionerGetByIDReq struct {
	ID string
}

type RekapitulasiHasilKuesionerGetByIDRes struct {
	RekapitulasiHasilKuesioner model.RekapitulasiHasilKuesioner
}

type RekapitulasiHasilKuesionerGetByID = core.ActionHandler[RekapitulasiHasilKuesionerGetByIDReq, RekapitulasiHasilKuesionerGetByIDRes]

func ImplRekapitulasiHasilKuesionerGetByID(db *gorm.DB) RekapitulasiHasilKuesionerGetByID {
	return func(ctx context.Context, req RekapitulasiHasilKuesionerGetByIDReq) (*RekapitulasiHasilKuesionerGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var RekapitulasiHasilKuesioner model.RekapitulasiHasilKuesioner
		if err := query.First(&RekapitulasiHasilKuesioner, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("RekapitulasiHasilKuesioner id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &RekapitulasiHasilKuesionerGetByIDRes{RekapitulasiHasilKuesioner: RekapitulasiHasilKuesioner}, nil
	}
}
