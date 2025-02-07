package model

type PhoneNumber string

func (r PhoneNumber) Validate() error {
	// pattern := `^(\+\d{1,3}[- ]?)?\d{10}$`
	// regex := regexp.MustCompile(pattern)

	// if !regex.MatchString(string(r)) {
	// 	return fmt.Errorf("invalid phone number format")
	// }

	return nil
}
