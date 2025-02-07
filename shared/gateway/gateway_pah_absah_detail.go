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

type PahAbsahGetDetailReq struct {
	ID string
}
type PahAbsahGetDetailResp struct {
	PahAbsah model.PahAbsah
}

type PahAbsahGetDetailGateway = core.ActionHandler[PahAbsahGetDetailReq, PahAbsahGetDetailResp]

func ImplPahAbsahGetDetailGateway(db *gorm.DB) PahAbsahGetDetailGateway {
	return func(ctx context.Context, request PahAbsahGetDetailReq) (*PahAbsahGetDetailResp, error) {

		query := middleware.GetDBFromContext(ctx, db)

		var pahAbsah model.PahAbsah

		err := query.Where("id=?", request.ID).First(&pahAbsah).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("pah absah id %v is not found", request.ID)
			}
			return nil, core.NewInternalServerError(err)
		}
		return &PahAbsahGetDetailResp{
			PahAbsah: pahAbsah,
		}, err
	}
}
