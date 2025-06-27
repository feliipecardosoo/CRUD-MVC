package validation

import (
	resterr "crud/src/configuration/rest-err"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
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

func BindAndValidateStrict(c *gin.Context, obj interface{}) *resterr.RestErr {
	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields() // rejeita campos extras

	if err := decoder.Decode(obj); err != nil {
		// Pode ser erro de campo extra ou json mal formatado
		return resterr.NewBadRequestError("JSON inválido ou contém campos extras: " + err.Error())
	}

	// Verifica se sobrou algo no body depois do JSON
	if decoder.More() {
		return resterr.NewBadRequestError("JSON contém dados extras após o objeto")
	}

	// Validação do validator/v10
	if err := Validate.Struct(obj); err != nil {
		// Reaproveita sua função que converte erros do validator em restErr
		return ValidadeMembroError(err)
	}

	return nil
}
