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

type GetConfigReq struct {
	Name string
}
type GetConfigResp struct {
	Config model.Config
}

const (
	GuestMode = "GUEST_MODE"
)

type GetConfigGateway = core.ActionHandler[GetConfigReq, GetConfigResp]

func ImplGetConfigGateway(db *gorm.DB) GetConfigGateway {
	return func(ctx context.Context, request GetConfigReq) (*GetConfigResp, error) {

		query := middleware.GetDBFromContext(ctx, db)

		var config model.Config

		err := query.Where("name = ?", request.Name).First(&config).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("config name %s is not found", request.Name)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &GetConfigResp{
			Config: config,
		}, err
	}
}
