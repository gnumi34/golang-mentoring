package config

import (
	"fmt"

	localEN "github.com/go-playground/locales/en"
	universalTranslator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ENTranslations "github.com/go-playground/validator/v10/translations/en"
)

type CustomValidator struct {
	Validator  *validator.Validate
	Translator universalTranslator.Translator
}

func NewCustomValidator() *CustomValidator {
	validate := validator.New()
	en := localEN.New()
	ut := universalTranslator.New(en, en)

	trans, _ := ut.GetTranslator("en")

	// register default translation
	err := ENTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(fmt.Sprintf("Error register default translations: %v", err))
	}
	return &CustomValidator{Validator: validate, Translator: trans}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
