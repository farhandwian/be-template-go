package usecase

import (
	"context"
	"fmt"
	"iam/gateway"
	"iam/model"
	"shared/core"
)

type AccessResetReq struct {
	UserID   model.UserID
	Accesses []model.MapAccessID
}

type AccessResetRes struct{}

type AccessReset = core.ActionHandler[AccessResetReq, AccessResetRes]

func ImplAccessReset(

	userGetOneByID gateway.UserGetOneByID,
	saveUser gateway.UserSave,

) AccessReset {
	return func(ctx context.Context, request AccessResetReq) (*AccessResetRes, error) {

		if err := request.Validate(); err != nil {
			return nil, err
		}

		userObj, err := userGetOneByID(ctx, gateway.UserGetOneByIDReq{UserID: request.UserID})
		if err != nil {
			return nil, err
		}

		if userObj == nil {
			return nil, fmt.Errorf("user id %v not found", request.UserID)
		}

		accesses := model.MapAccessIDsToAccess(request.Accesses)

		if err := userObj.User.UserAccess.ResetAccess(accesses...); err != nil {
			return nil, err
		}

		if _, err := saveUser(ctx, gateway.UserSaveReq{User: userObj.User}); err != nil {
			return nil, err
		}

		return &AccessResetRes{}, nil
	}
}

func (r AccessResetReq) Validate() error {

	if r.UserID == "" {
		return fmt.Errorf("user must not empty")
	}

	// if len(r.Accesses) == 0 {
	// 	return fmt.Errorf("access must greater than zero")
	// }

	return nil
}
