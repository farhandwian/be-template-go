package flowtest

import (
	"fmt"
	"iam/model"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func returnPasswordCorrect(storedPassword, comparedPassword string) string {
	if comparedPassword == "" {
		return ""
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(comparedPassword))
	if err != nil {
		panic(err.Error())
	}
	return comparedPassword
}

func printUser(message string, user model.User, password, pin, otp string) {
	fmt.Printf("\n   %s\n", message)
	fmt.Printf("     ID              : %v \n", user.ID)
	fmt.Printf("     EmailVerifiedAt : %v \n", user.EmailVerifiedAt.Format(time.DateTime))
	fmt.Printf("     Password        : %v \n", returnPasswordCorrect(user.Password, password))
	fmt.Printf("     PIN             : %v \n", returnPasswordCorrect(user.Pin, pin))
	fmt.Printf("     OTP Purpose     : %v \n", user.OTPPurpose)
	fmt.Printf("     OTP Value       : %v \n", returnPasswordCorrect(user.OTPValue, otp))
	fmt.Printf("     OTP Expired At  : %v \n", user.OTPExpirateAt.Format(time.DateTime))
	fmt.Printf("     RefreshTokenID  : %v \n", user.RefreshTokenID)
	fmt.Printf("     UserAccess      : %v \n\n", user.UserAccess)
}

func resetAllDatabaseForTestingPurpose(db *gorm.DB) {
	fmt.Printf("Reset all data in user table\n")
	db.Unscoped().Where("1 = 1").Delete(&model.User{})
}
