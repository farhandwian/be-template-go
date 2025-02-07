package controllertest

import (
	"context"
	"encoding/json"
	"iam/model"
	"shared/core"
	"shared/helper"
	"time"

	"github.com/google/uuid"
)

func NewAccessToken(userAccess model.UserAccess) string {

	userID := model.UserID(uuid.New().String())

	now := time.Now()

	userObjForAccessToken, _ := json.Marshal(model.UserTokenPayload{
		Subject:    model.ACCESS_TOKEN,
		UserID:     userID,
		UserAccess: userAccess,
	})

	jwt, _ := helper.NewJWTTokenizer("mock-secret-key")

	token, _ := jwt.CreateToken(userObjForAccessToken, now, 1*time.Minute)

	return token

}

func MockGateway[R any, S any](returnFunc func(R) (*S, error)) core.ActionHandler[R, S] {
	return func(ctx context.Context, request R) (*S, error) {
		return returnFunc(request)
	}
}
