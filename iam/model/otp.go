package model

type OTP string

func (r OTP) Validate() error {
	return nil
}

type OTPPurposeEnum string

const PASSWORD_CHANGE OTPPurposeEnum = "PASSWORD_CHANGE"
const PIN_CHANGE OTPPurposeEnum = "PIN_CHANGE"
const LOGIN OTPPurposeEnum = "LOGIN"
