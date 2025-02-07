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

type CoastalProtectionGetDetailReq struct {
	ID string
}
type CoastalProtectionGetDetailResp struct {
	CoastalProtection model.PengamanPantai
}

type CoastalProtectionGetDetailGateway = core.ActionHandler[CoastalProtectionGetDetailReq, CoastalProtectionGetDetailResp]

func ImplCoastalProtectionGetDetailGateway(db *gorm.DB) CoastalProtectionGetDetailGateway {
	return func(ctx context.Context, request CoastalProtectionGetDetailReq) (*CoastalProtectionGetDetailResp, error) {

		query := middleware.GetDBFromContext(ctx, db)

		var coastalProtection model.PengamanPantai

		err := query.Where("id=?", request.ID).First(&coastalProtection).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("coastal protection id %v is not found", request.ID)
			}
			return nil, core.NewInternalServerError(err)
		}
		return &CoastalProtectionGetDetailResp{
			CoastalProtection: coastalProtection,
		}, err
	}
}
