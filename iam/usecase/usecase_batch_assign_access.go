package usecase

import (
	"context"
	"fmt"
	"iam/model"
	"shared/core"

	ketoHelper "shared/helper/ory/keto"
)

type BatchAssignAccessReq struct {
	Objects    []string          `json:"objects"`
	Operations []string          `json:"operations"`
	Namespace  string            `json:"namespace"`
	SubjectSet *model.SubjectSet `json:"subject_set"`
	SubjectID  *string           `json:"subject_id"`
}

type BatchAssignAccessRes struct {
}

type BatchAssignAccessUseCase = core.ActionHandler[BatchAssignAccessReq, BatchAssignAccessRes]

func ImplBatchAssignAccess(ketoClient *ketoHelper.KetoGRPCClient) BatchAssignAccessUseCase {
	return func(ctx context.Context, request BatchAssignAccessReq) (*BatchAssignAccessRes, error) {
		subjectID := ""
		if request.SubjectID != nil {
			subjectID = *request.SubjectID
		}

		userAccess := model.NewUserAccessKeto(subjectID, ketoClient)
		isRole := request.SubjectID != nil && request.SubjectSet == nil
		for _, object := range request.Objects {
			for _, operation := range request.Operations {
				err := userAccess.AssignAccessPart2(ctx, request.Namespace, operation, object, isRole, request.SubjectSet)
				if err != nil {
					return nil, fmt.Errorf("error assigning %s permission for %s: %w", operation, object, err)
				}

				fmt.Printf("Assigned %s permission for %s to %s\n", operation, object, request.SubjectSet.Object)
			}
		}
		return &BatchAssignAccessRes{}, nil
	}
}
