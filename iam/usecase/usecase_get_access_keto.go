package usecase

import (
	"context"
	"iam/model"
	"shared/core"

	ketoHelper "shared/helper/ory/keto"
)

type UserGetAccessKetoReq struct {
	UserID    string `json:"user_id"`
	Namespace string `json:"namespace"`
}

type UserGetAccessKetoRes struct {
	Accesses []model.MapAccessKeto `json:"accesses"`
}

type UserGetAccessKetoUseCase = core.ActionHandler[UserGetAccessKetoReq, UserGetAccessKetoRes]

func ImplUserGetAccessKeto(ketoClient *ketoHelper.KetoGRPCClient) UserGetAccessKetoUseCase {
	return func(ctx context.Context, request UserGetAccessKetoReq) (*UserGetAccessKetoRes, error) {
		userAccess := model.NewUserAccessKeto(request.UserID, ketoClient)
		listAccess, err := userAccess.ListAccessByUser(ctx, request.Namespace)
		if err != nil {
			return nil, err
		}

		return &UserGetAccessKetoRes{
			Accesses: listAccess,
		}, nil
	}
}
