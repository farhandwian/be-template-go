package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type GetCoastalProtectionListReq struct {
}

type GetCoastalProtectionListResp struct {
	CoastalProtections []model.PengamanPantai
}

type GetCoastalProtectionListGateway = core.ActionHandler[GetCoastalProtectionListReq, GetCoastalProtectionListResp]

func ImplGetCoastalProtectionList(db *gorm.DB) GetCoastalProtectionListGateway {
	return func(ctx context.Context, request GetCoastalProtectionListReq) (*GetCoastalProtectionListResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var coastalProtection []model.PengamanPantai

		err := query.Where("ws_name=?", "CITANDUY").Find(&coastalProtection).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &GetCoastalProtectionListResp{
			CoastalProtections: coastalProtection,
		}, err
	}
}
