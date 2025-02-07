package model

import (
	"fmt"
)

type UserTokenPayload struct {
	Subject    UserTokenSubjectEnum `json:"subject"`
	UserID     UserID               `json:"user_id"`
	UserAccess UserAccess           `json:"user_access,omitempty"`
	TokenID    string               `json:"token_id,omitempty"`
}

func (r UserTokenPayload) ValidateSubject(subject UserTokenSubjectEnum) error {
	if r.Subject != subject {
		return fmt.Errorf("invalid token for %s", subject)
	}
	return nil
}

type UserTokenSubjectEnum string

const EMAIL_ACTIVATION UserTokenSubjectEnum = "EMAIL_ACTIVATION"
const PASSWORD_RESET UserTokenSubjectEnum = "PASSWORD_RESET"
const REFRESH_TOKEN UserTokenSubjectEnum = "REFRESH_TOKEN"
const ACCESS_TOKEN UserTokenSubjectEnum = "ACCESS_TOKEN"
