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

type EmailActivationRequestReq struct {
	UserID                  model.UserID
	EmailActivationDuration time.Duration
	EmailActivationPageUrl  string
	ServerUrl               string
	Now                     time.Time
}

type EmailActivationRequestRes struct{}

type EmailActivationRequest = core.ActionHandler[EmailActivationRequestReq, EmailActivationRequestRes]

func ImplEmailActivationRequest(

	userGetOneByID gateway.UserGetOneByID,
	generateJWT gateway.GenerateJWT,
	sendEmail gateway.SendEmail,

) EmailActivationRequest {
	return func(ctx context.Context, request EmailActivationRequestReq) (*EmailActivationRequestRes, error) {

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

		if userObj.User.IsEmailVerified() {
			return &EmailActivationRequestRes{}, nil
		}

		userTokenPayloadInfo, err := json.Marshal(model.UserTokenPayload{
			UserID:  userObj.User.ID,
			Subject: model.EMAIL_ACTIVATION,
		})
		if err != nil {
			return nil, err
		}

		jwtToken, err := generateJWT(ctx, gateway.GenerateJWTReq{
			Payload: userTokenPayloadInfo,
			Now:     request.Now,
			Expired: request.EmailActivationDuration,
		})
		if err != nil {
			return nil, err
		}

		body := generateActivationEmailBody(request.ServerUrl, userObj.User.Name, fmt.Sprintf("%s?token=%s", request.EmailActivationPageUrl, jwtToken.JWTToken))

		sendEmailReq := gateway.SendEmailReq{
			EmailRecipient: userObj.User.Email,
			Subject:        "Aktivasi Akun BBWS Command Center Citanduy",
			Body:           body,
		}

		if _, err := sendEmail(ctx, sendEmailReq); err != nil {
			return nil, err
		}

		return &EmailActivationRequestRes{}, nil
	}
}

func (r EmailActivationRequestReq) Validate() error {

	if strings.TrimSpace(r.EmailActivationPageUrl) == "" {
		return errors.New("activation server url must not empty")
	}

	if strings.TrimSpace(string(r.UserID)) == "" {
		return errors.New("user id must not empty")
	}

	if r.EmailActivationDuration <= 10*time.Second {
		return errors.New("expiration duration must greater than 10 seconds")
	}

	return nil
}

func generateActivationEmailBody(serverUrl, userName, activationUrl string) string {
	// 	// Plain text version
	// 	plainBody := fmt.Sprintf(`Halo %s

	// Email anda telah didaftarkan di Dashboard BBWS Command Center Citanduy. Untuk mengaktivasi akun Anda, silakan kunjungi link berikut:

	// %s

	// Tombol aktivasi di atas hanya berlaku 1 x 24 jam.
	// Terima Kasih

	// *Abaikan email ini jika anda tidak pernah di daftarkan oleh administrator`, userName, activationUrl)

	// HTML version
	htmlBody := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
			<meta charset="UTF-8">
			<style>
					body {
							font-family: Arial, sans-serif;
							line-height: 1.6;
							color: #333333;
					}
					.button {
							display: inline-block;
							padding: 12px 24px;
							background-color: #1a73e8;
							color: #ffffff !important;
							text-decoration: none;
							border-radius: 6px;
							font-weight: 500;
							font-size: 16px;
							text-align: center;
							margin: 20px 0;
							border: 1px solid #1a73e8;
							box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
							transition: all 0.3s ease;
					}
					.button:hover {
							background-color: #1557b0;
							border-color: #1557b0;
							box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
					}
					.container {
							max-width: 600px;
							margin: 0 auto;
							padding: 20px;
					}
					.note {
							font-size: 14px;
							color: #666666;
							font-style: italic;
							margin-top: 30px;
					}
			</style>
	</head>
	<body>
    <div class="banner-wrapper">
        <img src="%s/images/email_logo.png" alt="BBWS Command Center Citanduy" class="banner">
    </div>	
		<div class="container">
				<p>Halo %s</p>
				
				<p>Email anda telah didaftarkan di Dashboard BBWS Command Center Citanduy. Untuk mengaktivasi akun Anda, silakan klik tombol di bawah ini:</p>
				
				<p><a href="%s" class="button">Aktivasi Akun</a></p>
				
				<p>Tombol aktivasi di atas hanya berlaku 1 x 24 jam.</p>
				<p>Terima Kasih</p>
				
				<p class="note">*Abaikan email ini jika anda tidak pernah di daftarkan oleh administrator</p>
		</div>
	</body>
	</html>`, serverUrl, userName, activationUrl)

	return htmlBody
}
