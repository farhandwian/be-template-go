package usecase

import (
	"context"
	"errors"
	"fmt"
	"iam/gateway"
	"iam/model"
	"shared/core"
	"time"
)

type PinChangeSubmitReq struct {
	UserID   model.UserID
	OTPValue string
	NewPIN   string
	Now      time.Time
}

type PinChangeSubmitRes struct{}

type PinChangeSubmit = core.ActionHandler[PinChangeSubmitReq, PinChangeSubmitRes]

func ImplPinChangeSubmit(

	passwordValidate gateway.PasswordValidate,
	passwordEncrypt gateway.PasswordEncrypt,
	userGetOneByID gateway.UserGetOneByID,
	userSave gateway.UserSave,

) PinChangeSubmit {
	return func(ctx context.Context, request PinChangeSubmitReq) (*PinChangeSubmitRes, error) {

		if err := request.Validate(); err != nil {
			return nil, err
		}

		userObj, err := userGetOneByID(ctx, gateway.UserGetOneByIDReq{UserID: request.UserID})
		if err != nil {
			return nil, err
		}

		if userObj == nil {
			return nil, fmt.Errorf("user id %v not found", request.UserID)
		}

		if userObj.User.IsOTPExpirate(request.Now) {

			userObj.User.ResetOTP()

			if _, err = userSave(ctx, gateway.UserSaveReq{User: userObj.User}); err != nil {
				return nil, err
			}

			return nil, fmt.Errorf("OTP expirate")
		}

		if err := userObj.User.ValidateOTPPurpose(model.PIN_CHANGE); err != nil {
			return nil, err
		}

		// if err := ValidatePasswordValue(ctx, userObj.User.OTPValue, request.OTPValue, passwordValidate, "incorrect OTP"); err != nil {
		// 	return nil, err
		// }

		otpValidateReq := gateway.PasswordValidateReq{
			PasswordEncrypted: userObj.User.OTPValue,
			PasswordPlain:     request.OTPValue,
		}

		if _, err := passwordValidate(ctx, otpValidateReq); err != nil {
			return nil, errors.New("incorrect OTP")
		}

		peObj, err := passwordEncrypt(ctx, gateway.PasswordEncryptReq{PasswordPlain: request.NewPIN})
		if err != nil {
			return nil, err
		}

		userObj.User.Pin = peObj.PasswordEncrypted

		userObj.User.ResetOTP()

		userObj.User.SetUpdateAt(request.Now)

		if _, err := userSave(ctx, gateway.UserSaveReq{User: userObj.User}); err != nil {
			return nil, err
		}

		return &PinChangeSubmitRes{}, nil
	}
}

func (r PinChangeSubmitReq) Validate() error {

	if r.Now.IsZero() {
		return errors.New("now time must not zero")
	}

	return nil
}
