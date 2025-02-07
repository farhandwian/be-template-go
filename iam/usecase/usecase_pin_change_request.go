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

type PinChangeRequestReq struct {
	UserID            model.UserID
	PinChangeDuration time.Duration
	PinChangePageUrl  string
	Now               time.Time
}

type PinChangeRequestRes struct{}

type PinChangeRequest = core.ActionHandler[PinChangeRequestReq, PinChangeRequestRes]

func ImplPinChangeRequest(

	sendOTP gateway.SendOTP,
	userGetOneByID gateway.UserGetOneByID,
	generateRandom gateway.GenerateRandom,
	saveUser gateway.UserSave,
	passwordEncrypt gateway.PasswordEncrypt,

) PinChangeRequest {
	return func(ctx context.Context, request PinChangeRequestReq) (*PinChangeRequestRes, error) {

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

		nextTime := request.Now.Add(request.PinChangeDuration)

		encOTP, err := passwordEncrypt(ctx, gateway.PasswordEncryptReq{PasswordPlain: randomer.Random})
		if err != nil {
			return nil, err
		}

		userObj.User.OTPExpirateAt = nextTime
		userObj.User.OTPValue = encOTP.PasswordEncrypted
		userObj.User.OTPPurpose = model.PIN_CHANGE

		if _, err := saveUser(ctx, gateway.UserSaveReq{User: userObj.User}); err != nil {
			return nil, err
		}

		msg := changePinOTPMessage(randomer.Random)

		sendOTPReq := gateway.SendOTPReq{
			PhoneNumber: userObj.User.PhoneNumber,
			Message:     msg,
		}

		if _, err := sendOTP(ctx, sendOTPReq); err != nil {
			return nil, err
		}

		return &PinChangeRequestRes{}, nil
	}
}

func (r PinChangeRequestReq) Validate() error {

	if strings.TrimSpace(r.PinChangePageUrl) == "" {
		return errors.New("activation server url must not empty")
	}

	if strings.TrimSpace(string(r.UserID)) == "" {
		return errors.New("user id must not empty")
	}

	if r.PinChangeDuration <= 10*time.Second {
		return errors.New("expiration duration must greater than 10 seconds")
	}

	return nil
}

func changePinOTPMessage(otp string) string {
	return fmt.Sprintf("*%s* adalah kode OTP anda untuk melakukan perubahan PIN pada Dashboard Command Center BBWS Citanduy. Demi keamanan, harap jangan bagikan kode ini.", otp)
}
