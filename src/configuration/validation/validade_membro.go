package validation

import (
	resterr "crud/src/configuration/rest-err"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		un := ut.New(en, en)
		transl, _ = un.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidadeMembroError(validation_err error) *resterr.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return resterr.NewBadRequestError("Tipos de campos inválidos")
	}

	if errors.As(validation_err, &jsonValidationError) {
		errorCauses := []resterr.Causas{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := resterr.Causas{
				Mensagem: e.Translate(transl),
				Field:    e.Field(),
			}
			errorCauses = append(errorCauses, cause)
		}

		return resterr.NewBadRequestValidationError("Alguns campos estão incorretos", errorCauses)
	}
	return resterr.NewBadRequestError("Error trying to convert fields")
}
