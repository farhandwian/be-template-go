package model

import (
	"errors"
	"regexp"
)

type Email string

func (r Email) Validate() error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(string(r)) {
		return errors.New("invalid email format")
	}
	return nil
}
