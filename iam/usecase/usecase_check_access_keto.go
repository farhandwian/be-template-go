package usecase

import (
	"context"
	"iam/model"
	"shared/core"

	ketoHelper "shared/helper/ory/keto"
)

type CheckAccessKetoReq struct {
	Namespace string
	SubjectID string
	Object    string
	Relation  string
}

type CheckAccessKetoRes struct {
	CanAccess bool
}

type CheckAccessKetoUseCase = core.ActionHandler[CheckAccessKetoReq, CheckAccessKetoRes]

func ImplCheckAccessKeto(ketoClient *ketoHelper.KetoGRPCClient) CheckAccessKetoUseCase {
	return func(ctx context.Context, request CheckAccessKetoReq) (*CheckAccessKetoRes, error) {
		userAccess := model.NewUserAccessKeto(request.SubjectID, ketoClient)

		canAccess := userAccess.HasAccess(ctx, request.Namespace, request.Relation, request.Object)

		return &CheckAccessKetoRes{
			CanAccess: canAccess,
		}, nil
	}
}
