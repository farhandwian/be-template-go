package model

import (
	"context"
	"errors"
	ketoHelper "shared/helper/ory/keto"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type UserAccessKeto struct {
	UserID string
	Client *ketoHelper.KetoGRPCClient
}

func NewUserAccessKeto(userID string, client *ketoHelper.KetoGRPCClient) *UserAccessKeto {
	return &UserAccessKeto{
		UserID: userID,
		Client: client,
	}
}

// func NewUserAccessAdmin() UserAccess {
// 	userAccess := NewUserAccess()
// 	userAccess.AssignAccess(ADMIN_OPERATION)
// 	return userAccess
// }

func (ua *UserAccessKeto) HasAccess(ctx context.Context, namespace string, relation string, object string) bool {
	ketoRequest := &rts.ListRelationTuplesRequest{
		RelationQuery: &rts.RelationQuery{
			Namespace: &namespace,
			Object:    &object,
			Relation:  &relation,
			Subject: &rts.Subject{
				Ref: &rts.Subject_Id{
					Id: ua.UserID,
				},
			},
		},
	}

	response, err := ua.Client.ReadClient.ListRelationTuples(ctx, ketoRequest)
	if err != nil {
		return false
	}

	return len(response.RelationTuples) > 0
}

func (ua *UserAccessKeto) AssignAccess(ctx context.Context, namespace string, relation string, object string) error {
	relationTuple := &rts.RelationTuple{
		Namespace: namespace,
		Object:    object,
		Relation:  relation,
		Subject: &rts.Subject{
			Ref: &rts.Subject_Id{
				Id: ua.UserID,
			},
		},
	}

	_, err := ua.Client.WriteClient.TransactRelationTuples(ctx, &rts.TransactRelationTuplesRequest{
		RelationTupleDeltas: []*rts.RelationTupleDelta{
			{
				Action:        rts.RelationTupleDelta_ACTION_INSERT,
				RelationTuple: relationTuple,
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (ua *UserAccessKeto) RevokeAccess(ctx context.Context, namespace string, relation string, object string) error {
	deleteRequest := &rts.DeleteRelationTuplesRequest{
		RelationQuery: &rts.RelationQuery{
			Namespace: &namespace,
			Relation:  &relation,
			Object:    &object,
			Subject: &rts.Subject{
				Ref: &rts.Subject_Id{
					Id: ua.UserID,
				},
			},
		},
	}

	_, err := ua.Client.WriteClient.DeleteRelationTuples(ctx, deleteRequest)
	if err != nil {
		return err
	}
	return nil
}

var ErrPermissionDenied = errors.New("permission denied")
