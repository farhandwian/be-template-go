package usecase

import (
	"context"
	"iam/model"
	"shared/core"

	ketoHelper "shared/helper/ory/keto"
)

type AssignAccessKetoReq struct {
	Namespace string
	SubjectID string
	Object    string
	Relation  string
}

type AssignAccessKetoRes struct {
}

type AssignAccessKetoUseCase = core.ActionHandler[AssignAccessKetoReq, AssignAccessKetoRes]

func ImplAssignAccess(ketoClient *ketoHelper.KetoGRPCClient) AssignAccessKetoUseCase {
	return func(ctx context.Context, request AssignAccessKetoReq) (*AssignAccessKetoRes, error) {
		userAccess := model.NewUserAccessKeto(request.SubjectID, ketoClient)

		err := userAccess.AssignAccess(ctx, request.Namespace, request.Relation, request.Object)
		if err != nil {
			return nil, err
		}

		return &AssignAccessKetoRes{}, nil
	}
}
