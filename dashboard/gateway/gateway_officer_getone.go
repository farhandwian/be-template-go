package gateway

import (
	"context"
	"iam/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type OfficerGetOneReq struct {
	ID model.UserID
}
type OfficerGetOneRes struct {
	Officer model.User
}

type OfficerGetOneGateway = core.ActionHandler[OfficerGetOneReq, OfficerGetOneRes]

func ImplOfficerGetOneGateway(db *gorm.DB) OfficerGetOneGateway {
	return func(ctx context.Context, request OfficerGetOneReq) (*OfficerGetOneRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		var obj model.User

		if err := query.Where("id = ?", request.ID).First(&obj).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, core.NewInternalServerError(err)
		}
		return &OfficerGetOneRes{
			Officer: obj,
		}, nil
	}
}
