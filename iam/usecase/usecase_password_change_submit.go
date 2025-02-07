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

type PasswordChangeSubmitReq struct {
	UserID      model.UserID
	OTPValue    string
	OldPassword string
	NewPassword string
	Now         time.Time
}

type PasswordChangeSubmitRes struct{}

type PasswordChangeSubmit = core.ActionHandler[PasswordChangeSubmitReq, PasswordChangeSubmitRes]

func ImplPasswordChangeSubmit(

	passwordValidate gateway.PasswordValidate,
	passwordEncrypt gateway.PasswordEncrypt,
	userGetOneByID gateway.UserGetOneByID,
	userSave gateway.UserSave,

) PasswordChangeSubmit {
	return func(ctx context.Context, request PasswordChangeSubmitReq) (*PasswordChangeSubmitRes, error) {

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

		if err := userObj.User.ValidateOTPPurpose(model.PASSWORD_CHANGE); err != nil {
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

		// if err := ValidatePasswordValue(ctx, userObj.User.Password, request.OldPassword, passwordValidate, "incorrect old password"); err != nil {
		// 	return nil, err
		// }

		passwordValidateReq := gateway.PasswordValidateReq{
			PasswordEncrypted: userObj.User.Password,
			PasswordPlain:     request.OldPassword,
		}

		if _, err := passwordValidate(ctx, passwordValidateReq); err != nil {
			return nil, errors.New("incorrect old password")
		}

		peObj, err := passwordEncrypt(ctx, gateway.PasswordEncryptReq{PasswordPlain: request.NewPassword})
		if err != nil {
			return nil, err
		}

		userObj.User.SetPassword(peObj.PasswordEncrypted)

		userObj.User.ResetOTP()

		userObj.User.SetUpdateAt(request.Now)

		if _, err := userSave(ctx, gateway.UserSaveReq{User: userObj.User}); err != nil {
			return nil, err
		}

		return &PasswordChangeSubmitRes{}, nil
	}
}

func (r PasswordChangeSubmitReq) Validate() error {

	if r.Now.IsZero() {
		return errors.New("now time must not zero")
	}

	return nil
}
