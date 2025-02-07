package model

import (
	"fmt"
	"time"
)

type UserID string

type User struct {
	ID              UserID         `json:"id,omitempty" gorm:"primaryKey"`
	Name            string         `json:"name,omitempty"`
	PhoneNumber     PhoneNumber    `json:"phone_number,omitempty"`
	Email           Email          `json:"email,omitempty"`
	EmailVerifiedAt time.Time      `json:"email_verified,omitempty"`
	Enabled         bool           `json:"enabled,omitempty"`
	UserAccess      UserAccess     `json:"user_access,omitempty"`
	Password        string         `json:"-"`
	Pin             string         `json:"-"`
	OTPValue        string         `json:"-"`
	OTPExpirateAt   time.Time      `json:"-"`
	OTPPurpose      OTPPurposeEnum `json:"-"`
	RefreshTokenID  string         `json:"-"`
	CreatedAt       time.Time      `json:"created_at,omitempty"`
	UpdatedAt       time.Time      `json:"updated_at,omitempty"`
}

func NewUser(userId UserID, email Email, phoneNumber PhoneNumber, name string, now time.Time) User {
	user := User{
		ID:              userId,
		Email:           email,
		PhoneNumber:     phoneNumber,
		Name:            name,
		EmailVerifiedAt: time.Time{},
		Enabled:         true,
		RefreshTokenID:  "",
		CreatedAt:       now,
		UpdatedAt:       now,
		UserAccess:      NewUserAccess(),
	}
	user.ResetOTP()
	return user
}

func NewEmptyUser() User {
	return User{}
}

func (r User) IsEmailVerified() bool {
	return !r.EmailVerifiedAt.IsZero()
}

func (r User) IsExpired(now time.Time) bool {
	return r.OTPExpirateAt.IsZero() || r.OTPExpirateAt.After(now)
}

func (r User) ValidateOTPPurpose(purpose OTPPurposeEnum) error {
	if r.OTPPurpose != purpose {
		return fmt.Errorf("incorrect OTP purpose")
	}
	return nil
}

func (r User) IsOTPExpirate(now time.Time) bool {
	return r.OTPExpirateAt.Before(now)
}

func (r *User) VerifyEmail(now time.Time) {
	r.EmailVerifiedAt = now
}

func (r *User) ResetOTP() {
	r.OTPExpirateAt = time.Time{}
	r.OTPPurpose = ""
	r.OTPValue = ""
}

func (r User) IsValidRefreshToken(refreshTokenID string) bool {
	return r.RefreshTokenID == refreshTokenID
}

func (r *User) SetUpdateAt(now time.Time) {
	r.UpdatedAt = now
}

func (r *User) SetRefreshTokenID(newRefreshTokenID string) {
	r.RefreshTokenID = newRefreshTokenID
}

func (r *User) SetPassword(newPassword string) {
	r.Password = newPassword
}
