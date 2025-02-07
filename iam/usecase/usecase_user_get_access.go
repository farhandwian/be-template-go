package usecase

import (
	"context"
	"fmt"
	"iam/gateway"
	"iam/model"
	"shared/core"
)

type UserGetAccessReq struct{ gateway.UserGetOneByIDReq }

type UserGetAccessRes struct {
	Accesses []model.MapAccess `json:"accesses"`
}

type UserGetAccess = core.ActionHandler[UserGetAccessReq, UserGetAccessRes]

func ImplUserGetAccess(
	userGetOneByID gateway.UserGetOneByID,
) UserGetAccess {
	return func(ctx context.Context, request UserGetAccessReq) (*UserGetAccessRes, error) {

		userObj, err := userGetOneByID(ctx, gateway.UserGetOneByIDReq{UserID: request.UserID})
		if err != nil {
			return nil, err
		}

		if userObj == nil {
			return nil, fmt.Errorf("user id %v not found", request.UserID)
		}

		var userAccessItem []model.MapAccess
		for _, a := range model.GetMapAccess() {
			mapAccess, err := model.FindMapAccessByID(a.ID)
			if err != nil {
				return nil, err
			}

			userAccessItem = append(userAccessItem, model.MapAccess{
				ID:          mapAccess.ID,
				Access:      mapAccess.Access,
				Description: mapAccess.Description,
				Group:       mapAccess.Group,
				Type:        mapAccess.Type,
				Enabled:     userObj.User.UserAccess.HasAccess(mapAccess.Access),
			})
		}

		return &UserGetAccessRes{
			Accesses: userAccessItem,
		}, nil
	}
}
