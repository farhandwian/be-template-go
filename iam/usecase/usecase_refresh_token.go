package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"iam/gateway"
	"iam/model"
	"shared/core"
	"time"
)

type RefreshTokenReq struct {
	RefreshToken        string `json:"refresh_token"`
	AccessTokenDuration time.Duration
	Now                 time.Time
}

type RefreshTokenRes struct {
	AccessToken string `json:"access_token"`
}

type RefreshToken = core.ActionHandler[RefreshTokenReq, RefreshTokenRes]

func ImplRefreshToken(

	userGetOneByID gateway.UserGetOneByID,
	generateJWT gateway.GenerateJWT,
	validateJWT gateway.ValidateJWT,

) RefreshToken {
	return func(ctx context.Context, request RefreshTokenReq) (*RefreshTokenRes, error) {

		tokenObj, err := validateJWT(ctx, gateway.ValidateJWTReq{Token: request.RefreshToken})
		if err != nil {
			return nil, err
		}

		var userTokenPayloadInfo model.UserTokenPayload
		if err := json.Unmarshal(tokenObj.Payload, &userTokenPayloadInfo); err != nil {
			return nil, err
		}

		if err := userTokenPayloadInfo.ValidateSubject(model.REFRESH_TOKEN); err != nil {
			return nil, err
		}

		userObj, err := userGetOneByID(ctx, gateway.UserGetOneByIDReq{UserID: userTokenPayloadInfo.UserID})
		if err != nil {
			return nil, err
		}

		if userObj == nil {
			return nil, fmt.Errorf("user id %v not found", userTokenPayloadInfo.UserID)
		}

		user := userObj.User

		if !user.Enabled {
			return nil, fmt.Errorf("user is not enabled") // this feature is unused yet
		}

		if !user.IsValidRefreshToken(userTokenPayloadInfo.TokenID) {
			return nil, fmt.Errorf("user login in other device")
		}

		// accessToken, err := GenerateNewAccessToken(
		// 	ctx,
		// 	generateJWT,
		// 	request.AccessTokenDuration,
		// 	user,
		// )
		// if err != nil {
		// 	return nil, err
		// }

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
			Payload: userObjForAccessToken,
			Now:     request.Now,
		})
		if err != nil {
			return nil, err
		}

		return &RefreshTokenRes{
			AccessToken: accessToken.JWTToken,
		}, nil
	}
}

func ValidateToken_(
	ctx context.Context,
	validateJWT gateway.ValidateJWT,
	refreshToken string,
) (model.UserTokenPayload, error) {

	tokenObj, err := validateJWT(ctx, gateway.ValidateJWTReq{Token: refreshToken})
	if err != nil {
		return model.UserTokenPayload{}, err
	}

	var userTokenPayloadInfo model.UserTokenPayload
	if err := json.Unmarshal(tokenObj.Payload, &userTokenPayloadInfo); err != nil {
		return model.UserTokenPayload{}, err
	}

	if err := userTokenPayloadInfo.ValidateSubject(model.REFRESH_TOKEN); err != nil {
		return model.UserTokenPayload{}, err
	}

	return userTokenPayloadInfo, nil
}
