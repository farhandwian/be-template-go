package wiring

import (
	"fmt"
	"iam/model"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type EmailConfig struct {
	MailFrom         string
	EmailSmtpHost    string
	EmailSMTPPort    int
	GmailUsername    string
	GmailAppPassword string
}

func InitEmail() EmailConfig {

	ec := EmailConfig{}

	ec.MailFrom = os.Getenv("GMAIL_FROM")
	ec.EmailSmtpHost = os.Getenv("GMAIL_SMTP_HOST")
	ec.GmailUsername = os.Getenv("GMAIL_USERNAME")
	ec.GmailAppPassword = os.Getenv("GMAIL_APP_PASSWORD")
	emailSMTPPort, err := strconv.Atoi(os.Getenv("GMAIL_SMTP_PORT"))
	if err != nil {
		panic(err)
	}
	ec.EmailSMTPPort = emailSMTPPort

	return ec
}

func CreateAdminIfNotExists(db *gorm.DB) error {
	var user model.User

	id := uuid.New().String()
	email := os.Getenv("ADMIN_EMAIL")
	password, _ := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASSWORD")), bcrypt.DefaultCost)
	pin, _ := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PIN")), bcrypt.DefaultCost)

	result := db.Where(model.User{Email: model.Email(email)}).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		// User doesn't exist, create a new one
		newUser := model.User{
			ID:              model.UserID(id),
			Name:            os.Getenv("ADMIN_NAME"),
			PhoneNumber:     model.PhoneNumber(os.Getenv("ADMIN_PHONE")),
			Email:           model.Email(email),
			Enabled:         true,
			EmailVerifiedAt: time.Now(),
			UserAccess:      model.UserAccess("3"),
			Password:        string(password),
			Pin:             string(pin),
		}

		if err := db.Create(&newUser).Error; err != nil {
			return fmt.Errorf("failed to create admin user: %v", err)
		}

		fmt.Println("Admin user created successfully")
	} else if result.Error != nil {
		// Some other error occurred
		return fmt.Errorf("error checking for existing admin: %v", result.Error)
	} else {
		// User already exists
		fmt.Println("Admin user already exists")
	}

	return nil
}
