package gateway

import (
	"context"
	"perizinan/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type SKPerizinanGetAllReq struct {
	Keyword string
}

type SKPerizinanGetAllRes struct {
	Items []model.SKPerizinan
	Count int64
}

type SKPerizinanGetAll = core.ActionHandler[SKPerizinanGetAllReq, SKPerizinanGetAllRes]

func ImplSKPerizinanGetAll(db *gorm.DB) SKPerizinanGetAll {
	return func(ctx context.Context, req SKPerizinanGetAllReq) (*SKPerizinanGetAllRes, error) {
		var SKPerizinan []model.SKPerizinan
		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := "%" + req.Keyword + "%"
			query = query.
				Where("no_sk LIKE ?", keyword)
		}

		err := query.Find(&SKPerizinan).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &SKPerizinanGetAllRes{
			Items: SKPerizinan,
		}, nil
	}
}
