package config

import (
	"unicode"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_trans "github.com/go-playground/validator/v10/translations/en"
)

var (
	upp, low, num, sym bool
	tot                uint8
)

func Validator(i interface{}) interface{} {
	en := en.New()
	uni := ut.New(en)
	trans, _ := uni.GetTranslator("en")
	validate := validator.New()

	err := en_trans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return err
	}
	err = validate.Struct(i)
	if err != nil {
		errResult := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errResult = append(errResult, e.Translate(trans))
		}
		if len(errResult) > 0 {
			return errResult
		}
	}
	return nil
}

func ValidationPassword(password string) bool {
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}

	if !upp || !low || !num || !sym || tot < 8 {
		return false
	}

	return true
}
