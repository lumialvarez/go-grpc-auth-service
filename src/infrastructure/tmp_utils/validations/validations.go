package validations

import "net/mail"

func ValidPassword(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
