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

type GetClimatologyDetailByIDReq struct {
	ID string
}

type GetDetailClimatologyByIDResp struct {
	Climatology model.ClimatologyPost
}

type GetDetailClimatologyByIDGateway = core.ActionHandler[GetClimatologyDetailByIDReq, GetDetailClimatologyByIDResp]

func ImplGetClimatologyDetailByID(db *gorm.DB) GetDetailClimatologyByIDGateway {
	return func(ctx context.Context, request GetClimatologyDetailByIDReq) (*GetDetailClimatologyByIDResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var climatology model.ClimatologyPost

		err := query.Where("external_id = ? or id = ?", request.ID, request.ID).First(&climatology).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("climatology id %v is not found", request.ID)
			}
			return nil, core.NewInternalServerError(err)
		}
		return &GetDetailClimatologyByIDResp{
			Climatology: climatology,
		}, err
	}
}
