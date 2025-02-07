package usecase

import (
	"context"
	"encoding/json"
	"iam/gateway"
	"iam/model"
	"shared/core"
)

type CheckAccessReq struct {
	AccessToken  string
	FunctionName string
}

type CheckAccessRes struct {
	CanAccess bool `json:"can_access"`
}

type CheckAccess = core.ActionHandler[CheckAccessReq, CheckAccessRes]

func ImplCheckAccess(validateJwt gateway.ValidateJWT) CheckAccess {
	return func(ctx context.Context, req CheckAccessReq) (*CheckAccessRes, error) {

		payload, err := validateJwt(ctx, gateway.ValidateJWTReq{Token: req.AccessToken})
		if err != nil {
			return nil, err
		}

		var utp model.UserTokenPayload
		if err := json.Unmarshal(payload.Payload, &utp); err != nil {
			return nil, err
		}

		a, err := model.FindMapAccessByName(req.FunctionName)
		if err != nil {
			return nil, err
		}

		if !utp.UserAccess.HasAccess(a.Access) {
			return &CheckAccessRes{
				CanAccess: false,
			}, nil
		}

		return &CheckAccessRes{
			CanAccess: true,
		}, nil
	}
}
