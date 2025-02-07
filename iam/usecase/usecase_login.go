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

type LoginReq struct {
	Email       model.Email
	Password    string
	OTPDuration time.Duration
	Now         time.Time
}

type LoginRes struct{}

type Login = core.ActionHandler[LoginReq, LoginRes]

func ImplLogin(

	userGetAll gateway.UserGetAll,
	sendOTP gateway.SendOTP,
	generateRandom gateway.GenerateRandom,
	passwordValidate gateway.PasswordValidate,
	passwordEncrypt gateway.PasswordEncrypt,
	userSave gateway.UserSave,

) Login {
	return func(ctx context.Context, request LoginReq) (*LoginRes, error) {

		if err := request.Validate(); err != nil {
			return nil, err
		}

		usersObj, err := userGetAll(ctx, gateway.UserGetAllReq{Email: request.Email})
		if err != nil {
			return nil, err
		}

		if usersObj.Count == 0 {
			return nil, fmt.Errorf("password atau email tidak valid") // user with the email given is not exist
		}

		user := usersObj.Items[0]

		// if err := ValidatePasswordValue(ctx, user.Password, request.Password, passwordValidate, "user not found __"); err != nil {
		// 	return nil, err
		// }

		passwordValidateReq := gateway.PasswordValidateReq{
			PasswordEncrypted: user.Password,
			PasswordPlain:     request.Password,
		}

		if _, err := passwordValidate(ctx, passwordValidateReq); err != nil {
			return nil, errors.New("password atau email tidak valid")
		}

		if !user.IsEmailVerified() {
			return nil, fmt.Errorf("user belum mengaktivasi email") // need to do email activation first
		}

		if !user.Enabled {
			return nil, fmt.Errorf("user disabled") // this feature is unused yet
		}

		randomer, err := generateRandom(ctx, gateway.GenerateRandomReq{N: 6})
		if err != nil {
			return nil, err
		}

		nextTime := request.Now.Add(request.OTPDuration)

		user.OTPExpirateAt = nextTime

		encOTP, err := passwordEncrypt(ctx, gateway.PasswordEncryptReq{PasswordPlain: randomer.Random})
		if err != nil {
			return nil, err
		}

		user.OTPValue = encOTP.PasswordEncrypted
		user.OTPPurpose = model.LOGIN

		if _, err := userSave(ctx, gateway.UserSaveReq{User: user}); err != nil {
			return nil, err
		}

		msg := loginOTPMessage(randomer.Random)

		sendOTPReq := gateway.SendOTPReq{
			PhoneNumber: user.PhoneNumber,
			Message:     msg,
		}

		if _, err := sendOTP(ctx, sendOTPReq); err != nil {
			return nil, err
		}

		return &LoginRes{}, nil
	}
}

func (r LoginReq) Validate() error {

	if err := r.Email.Validate(); err != nil {
		return err
	}

	if strings.TrimSpace(r.Password) == "" {
		return errors.New("password must not empty")
	}

	return nil
}

func loginOTPMessage(otp string) string {
	return fmt.Sprintf("*%s* adalah kode OTP anda untuk masuk ke halaman Dashboard Command Center BBWS Citanduy. Demi keamanan, harap jangan bagikan kode ini.", otp)
}
