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

type GetLakeDetailByIDReq struct {
	ID string
}

type GetLakeDetailByIDResp struct {
	Lake model.Danau
}

type GetLakeDetailByIDGateway = core.ActionHandler[GetLakeDetailByIDReq, GetLakeDetailByIDResp]

func ImplGetLakeDetailByID(db *gorm.DB) GetLakeDetailByIDGateway {
	return func(ctx context.Context, request GetLakeDetailByIDReq) (*GetLakeDetailByIDResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var lake model.Danau

		err := query.Where("id = ?", request.ID).First(&lake).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("lake id %v is not found", request.ID)
			}
			return nil, core.NewInternalServerError(err)
		}
		return &GetLakeDetailByIDResp{
			Lake: lake,
		}, err
	}
}
