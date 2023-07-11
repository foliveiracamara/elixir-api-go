package validation

import (
	"encoding/json"
	"errors"

	"github.com/foliveiracamara/elixir-api-go/src/config/app_error"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateDonorError(validation_err error) *app_error.AppErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return app_error.NewBadRequestError("Invalid Field Type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorCauses := []app_error.Causes{}
		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := app_error.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}
		return app_error.NewBadRequestValidationError("Some fields are invalid", errorCauses)
	} else {
		return app_error.NewBadRequestError("Error trying to convert fields")
	}
}
