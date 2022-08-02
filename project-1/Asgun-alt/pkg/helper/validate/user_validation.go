package validate

import (
	"net/mail"
	"unicode"
)

var mustHave = []func(rune) bool{
	unicode.IsUpper,
	unicode.IsLower,
	unicode.IsPunct,
	unicode.IsDigit,
}

func ValidatePassword(password_1 string) bool {

	for _, checkRune := range mustHave {
		valid := false
		for _, r := range password_1 {
			if checkRune(r) {
				valid = true
			}
		}
		if !valid {
			return false
		}
	}
	return true
}

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
