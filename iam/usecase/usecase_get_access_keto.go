package usecase

import (
	"context"
	"iam/model"
	"shared/core"

	ketoHelper "shared/helper/ory/keto"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type UserGetAccessKetoReq struct {
	UserID    string `json:"user_id"`
	Namespace string `json:"namespace"`
}

type UserGetAccessKetoRes struct {
	ListAccess []*rts.RelationTuple `json:"list_access"`
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
			ListAccess: listAccess,
		}, nil
	}
}
