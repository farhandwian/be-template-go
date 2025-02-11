package usecase

import (
	"context"
	"log"
	"shared/core"

	ketoHelper "shared/helper/ory/keto"

	relationTuples "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
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

func ImplCheckAccessKeto() CheckAccessKetoUseCase {
	return func(ctx context.Context, request CheckAccessKetoReq) (*CheckAccessKetoRes, error) {
		ketoClient := ketoHelper.SetupKetoGRPCClient().ReadClient

		ketoRequest := &relationTuples.ListRelationTuplesRequest{
			RelationQuery: &relationTuples.RelationQuery{

				Namespace: &request.Namespace,
				Object:    &request.Object,
				Relation:  &request.Relation,
				Subject: &relationTuples.Subject{
					Ref: &relationTuples.Subject_Id{
						Id: request.SubjectID,
					},
				},
			},
		}
		response, err := ketoClient.ListRelationTuples(ctx, ketoRequest)
		if err != nil {
			log.Printf("Error checking permission in Keto: %v", err)
			return nil, err
		}

		canAccess := len(response.RelationTuples) > 0

		return &CheckAccessKetoRes{
			CanAccess: canAccess,
		}, nil
	}
}
