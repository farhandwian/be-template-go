package usecase

import (
	"context"
	"iam/model"
	"shared/core"

	ketoHelper "shared/helper/ory/keto"
)

type DeleteAccessKetoReq struct {
	Namespace string
	SubjectID string
	Object    string
	Relation  string
}

type DeleteAccessKetoRes struct {
}

type DeleteAccessKetoUseCase = core.ActionHandler[DeleteAccessKetoReq, DeleteAccessKetoRes]

func ImplDeleteAccessKeto(ketoClient *ketoHelper.KetoGRPCClient) DeleteAccessKetoUseCase {
	return func(ctx context.Context, request DeleteAccessKetoReq) (*DeleteAccessKetoRes, error) {
		userAccess := model.NewUserAccessKeto(request.SubjectID, ketoClient)

		err := userAccess.RevokeAccess(ctx, request.Namespace, request.Relation, request.Object)
		if err != nil {
			return nil, err
		}

		return &DeleteAccessKetoRes{}, nil
	}
}
