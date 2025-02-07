package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"iam/gateway"
	"iam/model"
	"shared/core"
	"strings"
	"time"
)

type EmailActivationSubmitReq struct {
	ActivationToken string
	Now             time.Time
	Password        string
	Pin             string
}

type EmailActivationSubmitRes struct{}

type EmailActivationSubmit = core.ActionHandler[EmailActivationSubmitReq, EmailActivationSubmitRes]

func ImplEmailActivationSubmit(

	validateJwt gateway.ValidateJWT,
	userGetOneByID gateway.UserGetOneByID,
	saveUser gateway.UserSave,
	passwordEncrypt gateway.PasswordEncrypt,

) EmailActivationSubmit {
	return func(ctx context.Context, request EmailActivationSubmitReq) (*EmailActivationSubmitRes, error) {

		if err := request.Validate(); err != nil {
			return nil, err
		}

		payload, err := validateJwt(ctx, gateway.ValidateJWTReq{Token: request.ActivationToken})
		if err != nil {
			return nil, err
		}

		var userTokenPayloadInfo model.UserTokenPayload

		if err := json.Unmarshal(payload.Payload, &userTokenPayloadInfo); err != nil {
			return nil, err
		}

		if err := userTokenPayloadInfo.ValidateSubject(model.EMAIL_ACTIVATION); err != nil {
			return nil, err
		}

		userObj, err := userGetOneByID(ctx, gateway.UserGetOneByIDReq{UserID: userTokenPayloadInfo.UserID})
		if err != nil {
			return nil, err
		}

		if userObj == nil {
			return nil, fmt.Errorf("user id %v not found", userTokenPayloadInfo.UserID)
		}

		encryptedPasswordObj, err := passwordEncrypt(ctx, gateway.PasswordEncryptReq{PasswordPlain: request.Password})
		if err != nil {
			return nil, err
		}

		encryptedPinObj, err := passwordEncrypt(ctx, gateway.PasswordEncryptReq{PasswordPlain: request.Pin})
		if err != nil {
			return nil, err
		}

		userObj.User.VerifyEmail(request.Now)
		userObj.User.SetUpdateAt(request.Now)

		userObj.User.Password = encryptedPasswordObj.PasswordEncrypted
		userObj.User.Pin = encryptedPinObj.PasswordEncrypted

		if _, err := saveUser(ctx, gateway.UserSaveReq{User: userObj.User}); err != nil {
			return nil, err
		}

		return &EmailActivationSubmitRes{}, nil
	}
}

func (r EmailActivationSubmitReq) Validate() error {

	if strings.TrimSpace(r.Password) == "" {
		return errors.New("password must not empty")
	}

	if r.Now.IsZero() {
		return errors.New("now time must not zero")
	}

	return nil
}
