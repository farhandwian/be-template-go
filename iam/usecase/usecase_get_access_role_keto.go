package usecase

import (
	"context"
	"iam/model"
	"shared/core"

	ketoHelper "shared/helper/ory/keto"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type GetAccessRoleKetoReq struct {
	Namespace string `json:"namespace"`
}

type GetAccessRoleKetoRes struct {
	ListAccess []model.RolePermissions
}

type GetAccessRoleKetoUseCase = core.ActionHandler[GetAccessRoleKetoReq, GetAccessRoleKetoRes]

func ImplGetAccessRoleKeto(ketoClient *ketoHelper.KetoGRPCClient) GetAccessRoleKetoUseCase {
	return func(ctx context.Context, request GetAccessRoleKetoReq) (*GetAccessRoleKetoRes, error) {
		req := &rts.ListRelationTuplesRequest{
			RelationQuery: &rts.RelationQuery{
				Namespace: &request.Namespace,
			},
			PageSize: 10,
		}

		res, err := ketoClient.ReadClient.ListRelationTuples(ctx, req)
		if err != nil {
			return nil, err
		}

		roleMap := make(map[string][]*rts.RelationTuple)
		for _, tuple := range res.RelationTuples {
			if tuple.Subject.GetSet() != nil {
				roleKey := tuple.Subject.GetSet().GetObject()
				roleMap[roleKey] = append(roleMap[roleKey], tuple)
			}
		}

		var result []model.RolePermissions
		for role, tuples := range roleMap {
			var perms []model.MapAccessKeto
			for _, tuple := range tuples {
				ma := model.MapAccessKeto{
					Namespace: tuple.Namespace,
					Object:    tuple.Object,
					Relation:  tuple.Relation,
					Group:     role,
					Type:      tuple.Relation,
					Enabled:   true,
				}
				perms = append(perms, ma)
			}
			result = append(result, model.RolePermissions{
				Role:        role,
				Permissions: perms,
			})
		}

		return &GetAccessRoleKetoRes{
			ListAccess: result,
		}, nil
	}
}
