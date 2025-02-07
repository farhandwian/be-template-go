package usecase

import (
	bigboardGateway "bigboard/gateway"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"iam/gateway"
	"iam/model"
	"shared/constant"
	"shared/core"
	sharedGateway "shared/gateway"
	sharedModel "shared/model"
	"time"

	"github.com/google/uuid"
)

type LoginOTPSubmitAuthenticatorReq struct {
	FCMToken             string
	Email                model.Email
	OTPValue             string
	RefreshTokenDuration time.Duration
	AccessTokenDuration  time.Duration
	Now                  time.Time
}

type LoginOTPSubmitAuthenticatorRes struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type LoginOTPSubmitAuthenticator = core.ActionHandler[LoginOTPSubmitAuthenticatorReq, LoginOTPSubmitAuthenticatorRes]

func ImplLoginOTPSubmitAuthenticator(
	passwordValidate gateway.PasswordValidate,
	userGetAll gateway.UserGetAll,
	generateJWT gateway.GenerateJWT,
	generateId gateway.GenerateId,
	userSave gateway.UserSave,
	createActivityMonitoring sharedGateway.CreateActivityMonitoringGateway,
	FCMTokenSave bigboardGateway.FCMTokenSave,
) LoginOTPSubmitAuthenticator {
	return func(ctx context.Context, request LoginOTPSubmitAuthenticatorReq) (*LoginOTPSubmitAuthenticatorRes, error) {

		if err := request.Validate(); err != nil {
			return nil, err
		}

		usersObj, err := userGetAll(ctx, gateway.UserGetAllReq{Email: request.Email})
		if err != nil {
			return nil, err
		}

		if usersObj.Count == 0 {
			return nil, fmt.Errorf("user not found") // user with the email given is not exist
		}

		user := usersObj.Items[0]

		if !user.IsEmailVerified() {
			return nil, fmt.Errorf("user is not activate the email yet") // need to do email activation first
		}

		if !user.Enabled {
			return nil, fmt.Errorf("user is not enabled") // this feature is unused yet
		}

		if user.IsOTPExpirate(request.Now) {

			user.ResetOTP()

			if _, err = userSave(ctx, gateway.UserSaveReq{User: user}); err != nil {
				return nil, err
			}

			return nil, fmt.Errorf("OTP expirate")
		}

		if err := user.ValidateOTPPurpose(model.LOGIN); err != nil {
			return nil, err
		}

		// if err := ValidatePasswordValue(ctx, user.OTPValue, request.OTPValue, passwordValidate, "incorrect OTP"); err != nil {
		// 	return nil, err
		// }

		passwordValidateReq := gateway.PasswordValidateReq{
			PasswordEncrypted: user.OTPValue,
			PasswordPlain:     request.OTPValue,
		}

		if _, err := passwordValidate(ctx, passwordValidateReq); err != nil {
			return nil, errors.New("incorrect OTP")
		}

		tokenId, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		refreshTokenID := tokenId.RandomId

		// refreshToken, accessToken, err := generateRefreshAndAccessToken(
		// 	ctx,
		// 	generateJWT,
		// 	refreshTokenID,
		// 	request.RefreshTokenDuration,
		// 	request.AccessTokenDuration,
		// 	user,
		// 	request.Now,
		// )
		// if err != nil {
		// 	return nil, err
		// }

		//---
		userObjForRefreshToken, err := json.Marshal(model.UserTokenPayload{
			Subject: model.REFRESH_TOKEN,
			UserID:  user.ID,
			TokenID: refreshTokenID,
		})
		if err != nil {
			return nil, err
		}

		refreshToken, err := generateJWT(ctx, gateway.GenerateJWTReq{
			Expired: request.RefreshTokenDuration,
			Now:     request.Now,
			Payload: userObjForRefreshToken,
		})
		if err != nil {
			return nil, err
		}

		userObjForAccessToken, err := json.Marshal(model.UserTokenPayload{
			Subject:    model.ACCESS_TOKEN,
			UserID:     user.ID,
			UserAccess: user.UserAccess,
		})
		if err != nil {
			return nil, err
		}

		accessToken, err := generateJWT(ctx, gateway.GenerateJWTReq{
			Expired: request.AccessTokenDuration,
			Now:     request.Now,
			Payload: userObjForAccessToken,
		})
		if err != nil {
			return nil, err
		}

		//---

		user.ResetOTP()

		user.SetRefreshTokenID(refreshTokenID)

		user.SetUpdateAt(request.Now)

		if _, err = FCMTokenSave(ctx, bigboardGateway.FCMTokenSaveReq{Token: request.FCMToken}); err != nil {
			return nil, err
		}

		if _, err = userSave(ctx, gateway.UserSaveReq{User: user}); err != nil {
			return nil, err
		}

		//store logging
		_, err = createActivityMonitoring(ctx, sharedGateway.CreateActivityMonitoringReq{
			ActivityMonitor: sharedModel.ActivityMonitor{
				ID:           uuid.NewString(),
				UserName:     user.Name,
				Category:     constant.MONITORING_TYPE_IAM,
				ActivityTime: time.Now(),
				Description:  fmt.Sprintf("%s masuk ke dalam sistem", user.Name),
			},
		})
		if err != nil {
			return nil, err
		}

		return &LoginOTPSubmitAuthenticatorRes{
			RefreshToken: refreshToken.JWTToken,
			AccessToken:  accessToken.JWTToken,
		}, nil
	}
}

func (r LoginOTPSubmitAuthenticatorReq) Validate() error {
	if len(r.FCMToken) == 0 {
		return errors.New("FCM token must not be empty")
	}

	if r.Now.IsZero() {
		return errors.New("now time must not zero")
	}

	return nil
}

// func generateRefreshAndAccessToken_(
// 	ctx context.Context,
// 	generateJWT gateway.GenerateJWT,
// 	refreshTokenID string,
// 	refreshTokenDuration, accessTokenDuration time.Duration,
// 	user model.User,
// 	now time.Time,
// ) (string, string, error) {

// 	userObjForRefreshToken, err := json.Marshal(model.UserTokenPayload{
// 		Subject: model.REFRESH_TOKEN,
// 		UserID:  user.ID,
// 		TokenID: refreshTokenID,
// 	})
// 	if err != nil {
// 		return "", "", err
// 	}

// 	refreshToken, err := generateJWT(ctx, gateway.GenerateJWTReq{
// 		Expired: refreshTokenDuration,
// 		Now:     now,
// 		Payload: userObjForRefreshToken,
// 	})
// 	if err != nil {
// 		return "", "", err
// 	}

// 	userObjForAccessToken, err := json.Marshal(model.UserTokenPayload{
// 		Subject:    model.ACCESS_TOKEN,
// 		UserID:     user.ID,
// 		UserAccess: user.UserAccess,
// 	})
// 	if err != nil {
// 		return "", "", err
// 	}

// 	accessToken, err := generateJWT(ctx, gateway.GenerateJWTReq{
// 		Expired: accessTokenDuration,
// 		Now:     now,
// 		Payload: userObjForAccessToken,
// 	})
// 	if err != nil {
// 		return "", "", err
// 	}

// 	return refreshToken.JWTToken, accessToken.JWTToken, nil

// }
