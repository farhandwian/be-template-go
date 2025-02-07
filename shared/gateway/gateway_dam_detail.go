package gateway

import (
	"context"
	"errors"
	"fmt"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type GetDamDetailByIDReq struct {
	ID string
}

type GetDetailDamByIDResp struct {
	Dam model.Bendungan
}

type GetDetailDamByIDGateway = core.ActionHandler[GetDamDetailByIDReq, GetDetailDamByIDResp]

func ImplGetDamDetailByID(db *gorm.DB) GetDetailDamByIDGateway {
	return func(ctx context.Context, request GetDamDetailByIDReq) (*GetDetailDamByIDResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var dam model.Bendungan

		err := query.Where("id = ?", request.ID).First(&dam).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("lake id %v is not found", request.ID)
			}
			return nil, core.NewInternalServerError(err)
		}
		return &GetDetailDamByIDResp{
			Dam: dam,
		}, err
	}
}
