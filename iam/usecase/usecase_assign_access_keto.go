package usecase

import (
	"context"
	"iam/model"
	"shared/core"

	ketoHelper "shared/helper/ory/keto"
)

type AssignAccessKetoReq struct {
	Namespace  string            `json:"namespace"`
	SubjectID  *string           `json:"subject_id"`
	Object     string            `json:"object"`
	Relation   string            `json:"relation"`
	SubjectSet *model.SubjectSet `json:"subject_set"`
}

type AssignAccessKetoRes struct {
}

type AssignAccessKetoUseCase = core.ActionHandler[AssignAccessKetoReq, AssignAccessKetoRes]

func ImplAssignAccess(ketoClient *ketoHelper.KetoGRPCClient) AssignAccessKetoUseCase {
	return func(ctx context.Context, request AssignAccessKetoReq) (*AssignAccessKetoRes, error) {
		subjectID := ""
		if request.SubjectID != nil {
			subjectID = *request.SubjectID
		}

		userAccess := model.NewUserAccessKeto(subjectID, ketoClient)
		isRole := request.SubjectID != nil && request.SubjectSet == nil

		err := userAccess.AssignAccess(ctx, request.Namespace, request.Relation, request.Object, isRole, request.SubjectSet)
		if err != nil {
			return nil, err
		}

		return &AssignAccessKetoRes{}, nil
	}
}
