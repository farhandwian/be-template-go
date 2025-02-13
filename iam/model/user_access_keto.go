package model

import (
	"context"
	"errors"
	"fmt"
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
	checkRequest := &rts.CheckRequest{
		Tuple: &rts.RelationTuple{
			Namespace: namespace,
			Object:    object,
			Relation:  relation,
			Subject: &rts.Subject{
				Ref: &rts.Subject_Id{
					Id: ua.UserID,
				},
			},
		},
	}

	response, err := ua.Client.CheckClient.Check(ctx, checkRequest)
	if err != nil {
		return false
	}

	return response.GetAllowed()
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

// func (ua *UserAccessKeto) ListAccessByUser(ctx context.Context, namespace string) ([]*rts.RelationTuple, error) {
// 	// Step 1: Fetch all relation tuples
// 	req := &rts.ListRelationTuplesRequest{
// 		RelationQuery: &rts.RelationQuery{
// 			Namespace: &namespace,
// 		},
// 	}

// 	res, err := ua.Client.ReadClient.ListRelationTuples(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var userRoles []string
// 	var userAccess []*rts.RelationTuple

// 	// Step 2: Find all roles where the user is a direct subject
// 	for _, tuple := range res.RelationTuples {
// 		if tuple.Subject != nil {
// 			if subj, ok := tuple.Subject.Ref.(*rts.Subject_Id); ok && subj.Id == ua.UserID {
// 				userRoles = append(userRoles, tuple.Object) // Collect roles (e.g., "viewer", "admin")
// 				userAccess = append(userAccess, tuple)      // Direct access
// 			}
// 		}
// 	}

// 	// Step 3: Find all permissions linked to these roles
// 	for _, tuple := range res.RelationTuples {
// 		for _, role := range userRoles {
// 			if tuple.Subject.GetSet() != nil &&
// 				tuple.Subject.GetSet().GetObject() == role &&
// 				tuple.Subject.GetSet().GetNamespace() == namespace {
// 				userAccess = append(userAccess, tuple)
// 			}
// 		}
// 	}

// 	return userAccess, nil
// }

func (ua *UserAccessKeto) ListAccessByUser(ctx context.Context, namespace string) ([]*rts.RelationTuple, error) {
	var userAccess []*rts.RelationTuple

	// Step 1: Fetch roles assigned directly to the user
	roleReq := &rts.ListRelationTuplesRequest{
		RelationQuery: &rts.RelationQuery{
			Namespace: &namespace,
			Subject: &rts.Subject{
				Ref: &rts.Subject_Id{Id: ua.UserID},
			},
		},
	}

	roleRes, err := ua.Client.ReadClient.ListRelationTuples(ctx, roleReq)
	if err != nil {
		return nil, err
	}

	// Store user roles to fetch permissions later
	roleSet := make(map[string]struct{})
	for _, tuple := range roleRes.RelationTuples {
		roleSet[tuple.Object] = struct{}{}
		userAccess = append(userAccess, tuple) // Add direct roles
	}

	// Step 2: Fetch permissions for user's roles
	for role := range roleSet {
		permReq := &rts.ListRelationTuplesRequest{
			RelationQuery: &rts.RelationQuery{
				Namespace: &namespace,
				Subject: &rts.Subject{
					Ref: &rts.Subject_Set{
						Set: &rts.SubjectSet{
							Namespace: namespace,
							Object:    role,
							Relation:  "member",
						},
					},
				},
			},
		}

		permRes, err := ua.Client.ReadClient.ListRelationTuples(ctx, permReq)
		if err != nil {
			return nil, err
		}

		fmt.Println(permRes)
		userAccess = append(userAccess, permRes.RelationTuples...)
	}

	return userAccess, nil
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
