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

type SubjectSet struct {
	Namespace string `json:"namespace"`
	Object    string `json:"object"`
	Relation  string `json:"relation"`
}

func NewUserAccessKeto(userID string, client *ketoHelper.KetoGRPCClient) *UserAccessKeto {
	return &UserAccessKeto{
		UserID: userID,
		Client: client,
	}
}

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

func (ua *UserAccessKeto) AssignAccess(ctx context.Context, namespace string, relation string, object string, isRole bool, subjectSet *SubjectSet) error {
	var ketoRequest *rts.TransactRelationTuplesRequest

	if isRole {
		// Assign role (e.g., `admin`, `member`)
		ketoRequest = &rts.TransactRelationTuplesRequest{
			RelationTupleDeltas: []*rts.RelationTupleDelta{
				{
					Action: rts.RelationTupleDelta_ACTION_INSERT,
					RelationTuple: &rts.RelationTuple{
						Namespace: namespace,
						Object:    object,
						Relation:  relation,
						Subject: &rts.Subject{
							Ref: &rts.Subject_Id{
								Id: ua.UserID,
							},
						},
					},
				},
			},
		}
	} else {
		// assign direct permission to the role (e.g., "admin -> read/write/delete to related object")
		ketoRequest = &rts.TransactRelationTuplesRequest{
			RelationTupleDeltas: []*rts.RelationTupleDelta{
				{
					Action: rts.RelationTupleDelta_ACTION_INSERT,
					RelationTuple: &rts.RelationTuple{
						Namespace: namespace,
						Object:    object,
						Relation:  relation,
						Subject: &rts.Subject{
							Ref: &rts.Subject_Set{
								Set: &rts.SubjectSet{
									Namespace: subjectSet.Namespace,
									Object:    subjectSet.Object,
									Relation:  subjectSet.Relation,
								},
							},
						},
					},
				},
			},
		}
	}

	_, err := ua.Client.WriteClient.TransactRelationTuples(ctx, ketoRequest)
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
