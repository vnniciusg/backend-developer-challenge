package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/http/responseshttp"
)

func ValidateDataRequest(data interface{}) []responseshttp.Causes {
	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		var validatorErrors []responseshttp.Causes

		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrors {
				validatorErrors = append(validatorErrors, responseshttp.Causes{
					Field:   e.Field(),
					Message: e.Tag(),
				})
			}
			return validatorErrors
		} else {
			return []responseshttp.Causes{{
				Field:   "error",
				Message: "invalid data",
			}}
		}
	}

	return nil
}
