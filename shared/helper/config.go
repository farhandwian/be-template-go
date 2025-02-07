package helper

// func NewAppConfig() (AppConfig, error) {

// 	emailActivationPageUrl := os.Getenv("EMAIL_ACTIVATION_PAGE_URL")
// 	passwordChangePageUrl := os.Getenv("PASSWORD_CHANGE_PAGE_URL")
// 	passwordResetPageUrl := os.Getenv("PASSWORD_RESET_PAGE_URL")
// 	pinChangePageUrl := os.Getenv("PIN_CHANGE_PAGE_URL")

// 	emailExpirationInSecond, err := strconv.Atoi(os.Getenv("EMAIL_EXPIRATION_IN_SECOND"))
// 	if err != nil {
// 		return AppConfig{}, err
// 	}

// 	otpExpirationInSecond, err := strconv.Atoi(os.Getenv("OTP_EXPIRATION_IN_SECOND"))
// 	if err != nil {
// 		return AppConfig{}, err
// 	}

// 	refreshTokenInSecond, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_IN_SECOND"))
// 	if err != nil {
// 		return AppConfig{}, err
// 	}

// 	accessTokenInSecond, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_IN_SECOND"))
// 	if err != nil {
// 		return AppConfig{}, err
// 	}

// 	// mailFrom := os.Getenv("GMAIL_FROM")
// 	// smtpHost := os.Getenv("GMAIL_SMTP_HOST")
// 	// gmailUsername := os.Getenv("GMAIL_USERNAME")
// 	// gmailAppPassword := os.Getenv("GMAIL_APP_PASSWORD")
// 	// emailSMTPPort, err := strconv.Atoi(os.Getenv("GMAIL_SMTP_PORT"))
// 	// if err != nil {
// 	// 	return AppConfig{}, err
// 	// }

// 	appConfig := AppConfig{
// 		EmailActivationPageUrl:  emailActivationPageUrl,
// 		PasswordChangePageUrl:   passwordChangePageUrl,
// 		PasswordResetPageUrl:    passwordResetPageUrl,
// 		PinChangePageUrl:        pinChangePageUrl,
// 		EmailExpirationInSecond: time.Duration(emailExpirationInSecond) * time.Second,
// 		OTPExpirationInSecond:   time.Duration(otpExpirationInSecond) * time.Second,
// 		RefreshTokenInSecond:    time.Duration(refreshTokenInSecond) * time.Second,
// 		AccessTokenInSecond:     time.Duration(accessTokenInSecond) * time.Second,

// 		// EmailSMTPPort:    emailSMTPPort,
// 		// EmailFrom:        mailFrom,
// 		// EmailSMTPHost:    smtpHost,
// 		// EmailUsername:    gmailUsername,
// 		// EmailAppPassword: gmailAppPassword,
// 	}

// 	return appConfig, nil
// }

// type AppConfig struct {
// 	EmailActivationPageUrl  string
// 	PasswordChangePageUrl   string
// 	PasswordResetPageUrl    string
// 	PinChangePageUrl        string
// 	EmailExpirationInSecond time.Duration
// 	OTPExpirationInSecond   time.Duration
// 	RefreshTokenInSecond    time.Duration
// 	AccessTokenInSecond     time.Duration
// }
