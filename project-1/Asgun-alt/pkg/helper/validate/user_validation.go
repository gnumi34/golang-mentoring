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

func ValidatePassword(password string) bool {
	for _, checkRune := range mustHave {
		valid := false
		for _, r := range password {
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

func MustNotBeEmpty(value string) bool {
	if value != "" {
		return false
	} else {
		return true
	}
}
