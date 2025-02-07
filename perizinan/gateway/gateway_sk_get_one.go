package gateway

import (
	"context"
	"perizinan/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type SKPerizinanGetOneReq struct {
	NomorSK model.NomorSK
}

type SKPerizinanGetOneRes struct {
	SKPerizinan *model.SKPerizinan
}

type SKPerizinanGetOne = core.ActionHandler[SKPerizinanGetOneReq, SKPerizinanGetOneRes]

func ImplSKPerizinanGetOne(db *gorm.DB) SKPerizinanGetOne {
	return func(ctx context.Context, req SKPerizinanGetOneReq) (*SKPerizinanGetOneRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		var skPerizinan model.SKPerizinan

		err := query.First(&skPerizinan, "no_sk = ?", req.NomorSK).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return &SKPerizinanGetOneRes{SKPerizinan: nil}, nil
			}
			return nil, core.NewInternalServerError(err)
		}

		return &SKPerizinanGetOneRes{SKPerizinan: &skPerizinan}, nil
	}
}
