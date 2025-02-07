package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type GetPahAbsahListReq struct {
}

type GetPahAbsahListResp struct {
	PahAbsahs []model.PahAbsah
}

type GetPahAbsahListGateway = core.ActionHandler[GetPahAbsahListReq, GetPahAbsahListResp]

func ImplGetPahAbsahList(db *gorm.DB) GetPahAbsahListGateway {
	return func(ctx context.Context, request GetPahAbsahListReq) (*GetPahAbsahListResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var pahAbsahs []model.PahAbsah

		err := query.Where("ws_name=?", "CITANDUY").Find(&pahAbsahs).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &GetPahAbsahListResp{
			PahAbsahs: pahAbsahs,
		}, err
	}
}
