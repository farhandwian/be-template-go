package usecase

import (
	"context"
	"errors"
	"fmt"
	"iam/gateway"
	"iam/model"
	"shared/core"
	"strings"
	"time"
)

type PasswordChangeRequestReq struct {
	UserID                 model.UserID
	PasswordChangeDuration time.Duration
	PasswordChangePageUrl  string
	Now                    time.Time
}

type PasswordChangeRequestRes struct{}

type PasswordChangeRequest = core.ActionHandler[PasswordChangeRequestReq, PasswordChangeRequestRes]

func ImplPasswordChangeRequest(

	sendOTP gateway.SendOTP,
	userGetOneByID gateway.UserGetOneByID,
	generateRandom gateway.GenerateRandom,
	userSave gateway.UserSave,
	passwordEncrypt gateway.PasswordEncrypt,

) PasswordChangeRequest {
	return func(ctx context.Context, request PasswordChangeRequestReq) (*PasswordChangeRequestRes, error) {

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

		randomer, err := generateRandom(ctx, gateway.GenerateRandomReq{N: 6})
		if err != nil {
			return nil, err
		}

		nextTime := request.Now.Add(request.PasswordChangeDuration)

		encOTP, err := passwordEncrypt(ctx, gateway.PasswordEncryptReq{PasswordPlain: randomer.Random})
		if err != nil {
			return nil, err
		}

		userObj.User.OTPExpirateAt = nextTime
		userObj.User.OTPValue = encOTP.PasswordEncrypted
		userObj.User.OTPPurpose = model.PASSWORD_CHANGE

		if _, err = userSave(ctx, gateway.UserSaveReq{User: userObj.User}); err != nil {
			return nil, err
		}

		msg := changePasswordOTPMessage(randomer.Random)

		sendOTPReq := gateway.SendOTPReq{
			PhoneNumber: userObj.User.PhoneNumber,
			Message:     msg,
		}

		if _, err := sendOTP(ctx, sendOTPReq); err != nil {
			return nil, err
		}

		return &PasswordChangeRequestRes{}, nil
	}
}

func (r PasswordChangeRequestReq) Validate() error {

	if strings.TrimSpace(r.PasswordChangePageUrl) == "" {
		return errors.New("activation server url must not empty")
	}

	if strings.TrimSpace(string(r.UserID)) == "" {
		return errors.New("user id must not empty")
	}

	if r.PasswordChangeDuration <= 10*time.Second {
		return errors.New("expiration duration must greater than 10 seconds")
	}

	return nil
}

func changePasswordOTPMessage(otp string) string {
	return fmt.Sprintf("*%s* adalah kode OTP anda untuk melakukan perubahan password pada Dashboard Command Center BBWS Citanduy. Demi keamanan, harap jangan bagikan kode ini.", otp)
}
