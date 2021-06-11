package validations

import "github.com/go-playground/validator"

type ErrorResponse struct {
	FailedField string `json:"field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func ValidateStruct(data interface{}) []*ErrorResponse {
	var errors []*ErrorResponse

	validate := validator.New()

	err := validate.Struct(data)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse

			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()

			errors = append(errors, &element)
		}
	}

	return errors
}
