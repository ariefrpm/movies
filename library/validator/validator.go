package validator

import (
	"net/http"
	"net/url"
	"regexp"
	"sync"

	"github.com/go-playground/form"
	"github.com/tokopedia/logistic/svc/mp-logistic/infra/errors"
	"github.com/tokopedia/logistic/svc/mp-logistic/infra/locale"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate
var once sync.Once
var decoder *form.Decoder

const (
	requiredLocaleID  = "RequiredValidationMessage"
	minLocaleID       = "MinValidationMessage"
	minStringLocaleID = "MinStringValidationMessage"
	maxLocaleID       = "MaxValidationMessage"
	maxStringLocaleID = "MaxStringValidationMessage"
	regexLocaleID     = "NotValid"
)

func InitValidator() {
	once.Do(func() {
		validate = validator.New()
		validate.RegisterValidation("regex", func(fl validator.FieldLevel) bool {
			matched, err := regexp.MatchString(fl.Param(), fl.Field().String())
			if err != nil || !matched {
				return false
			}
			return true
		})

		decoder = form.NewDecoder()
		decoder.SetTagName("json")

	})
}

//ValidateStruct will validate struct based on `validate` tag
//example of usage is on domain/logistic/confirm_shipping.go
func ValidateStruct(s interface{}, lang string) error {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	msgErr := err.Error()

	//because this library does not support custom locale message, we need to do this manually
	//replace error message with custom locale message
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			switch e.Tag() {
			case "min":
				msgErr = getMinLocalMessage(lang, e)
			case "max":
				msgErr = getMaxLocalMessage(lang, e)
			case "regex":
				msgErr = getRegexLocalMessage(lang, e)
			case "required":
				msgErr = getRequiredLocaleMessage(lang, e)

			}
		}
	}

	return errors.Newr(msgErr, errors.WithHTTPErrResp(http.StatusBadRequest, msgErr))
}

func getRequiredLocaleMessage(lang string, e validator.FieldError) string {
	return locale.TranslateString(requiredLocaleID, lang, map[string]interface{}{
		"field": getFieldName(e.Namespace(), e.Field(), lang),
	})
}
func getMinLocalMessage(lang string, e validator.FieldError) string {
	tid := minLocaleID
	if e.Type().Name() == "string" {
		tid = minStringLocaleID
	}
	return locale.TranslateString(tid, lang, map[string]interface{}{
		"field": getFieldName(e.Namespace(), e.Field(), lang),
		"value": e.Param(),
	})
}
func getMaxLocalMessage(lang string, e validator.FieldError) string {
	tid := maxLocaleID
	if e.Type().Name() == "string" {
		tid = maxStringLocaleID
	}
	return locale.TranslateString(tid, lang, map[string]interface{}{
		"field": getFieldName(e.Namespace(), e.Field(), lang),
		"value": e.Param(),
	})
}
func getRegexLocalMessage(lang string, e validator.FieldError) string {
	return locale.TranslateString(regexLocaleID, lang, map[string]interface{}{
		"field": getFieldName(e.Namespace(), e.Field(), lang),
	})
}

//if want to replace field name with locale text
//put `struct_name.field_name` on locale file i.e `ConfirmShippingRequest.ShippingNumber`
func getFieldName(namespace, defaultField, lang string) string {
	field := locale.TranslateString(namespace, lang)
	if field == "" || field == namespace {
		field = defaultField
	}
	return field
}

//Decode url form and url query data
func Decode(i interface{}, v url.Values) error {
	return decoder.Decode(i, v)
}
