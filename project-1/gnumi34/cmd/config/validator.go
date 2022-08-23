package config

import (
	"fmt"

	localeEN "github.com/go-playground/locales/en"
	universalTranslator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type CustomValidator struct {
	Validator  *validator.Validate
	Translator universalTranslator.Translator
}

func NewCustomValidator() *CustomValidator {
	validate := validator.New()
	// Instance locale en
	en := localeEN.New()
	// Instance universal translator en
	ut := universalTranslator.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := ut.GetTranslator("en")

	// Register default translation
	err := enTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(fmt.Sprintf("Error register default translations : %v", err))
	}
	return &CustomValidator{Validator: validate, Translator: trans}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
