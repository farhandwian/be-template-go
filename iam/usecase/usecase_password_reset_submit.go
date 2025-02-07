package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"iam/gateway"
	"iam/model"
	"shared/core"
	"time"
)

type PasswordResetSubmitReq struct {
	PasswordResetToken string
	NewPassword        string
	Now                time.Time
}

type PasswordResetSubmitRes struct{}

type PasswordResetSubmit = core.ActionHandler[PasswordResetSubmitReq, PasswordResetSubmitRes]

func ImplPasswordResetSubmit(

	validateJwt gateway.ValidateJWT,
	userGetOneByID gateway.UserGetOneByID,
	saveUser gateway.UserSave,
	passwordEncrypt gateway.PasswordEncrypt,

) PasswordResetSubmit {
	return func(ctx context.Context, request PasswordResetSubmitReq) (*PasswordResetSubmitRes, error) {

		if err := request.Validate(); err != nil {
			return nil, err
		}

		payload, err := validateJwt(ctx, gateway.ValidateJWTReq{Token: request.PasswordResetToken})
		if err != nil {
			return nil, err
		}

		var userTokenPayloadInfo model.UserTokenPayload
		if err := json.Unmarshal(payload.Payload, &userTokenPayloadInfo); err != nil {
			return nil, err
		}

		if err := userTokenPayloadInfo.ValidateSubject(model.PASSWORD_RESET); err != nil {
			return nil, err
		}

		userObj, err := userGetOneByID(ctx, gateway.UserGetOneByIDReq{UserID: userTokenPayloadInfo.UserID})
		if err != nil {
			return nil, err
		}

		if userObj == nil {
			return nil, fmt.Errorf("user id %v not found", userTokenPayloadInfo.UserID)
		}

		peObj, err := passwordEncrypt(ctx, gateway.PasswordEncryptReq{PasswordPlain: request.NewPassword})
		if err != nil {
			return nil, err
		}

		userObj.User.SetPassword(peObj.PasswordEncrypted)

		userObj.User.SetUpdateAt(request.Now)

		if _, err := saveUser(ctx, gateway.UserSaveReq{User: userObj.User}); err != nil {
			return nil, err
		}

		return &PasswordResetSubmitRes{}, nil
	}
}

func (r PasswordResetSubmitReq) Validate() error {

	if r.Now.IsZero() {
		return errors.New("now time must not zero")
	}

	return nil
}
