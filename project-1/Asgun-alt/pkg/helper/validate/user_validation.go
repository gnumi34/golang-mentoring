package validate

import (
	"unicode"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// caches struct info
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
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

func Validator(i interface{}) interface{} {
	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	validate = validator.New()
	errCheck := en_translations.RegisterDefaultTranslations(validate, trans)
	if errCheck != nil {
		arrError := []string{}
		return append(arrError, errCheck.Error())
	}
	return TranslateErr(trans, i)
}

func TranslateErr(trans ut.Translator, input interface{}) interface{} {
	err := validate.Struct(input)
	if err != nil {
		arrError := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			arrError = append(arrError, e.Translate(trans))
		}
		if len(arrError) > 0 {
			return arrError
		}
	}
	return nil
}
