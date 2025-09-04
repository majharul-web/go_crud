package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// FormatValidationError converts Gin/validator errors into a professional JSON map
func FormatValidationError(err error) map[string]string {
    errors := make(map[string]string)

    if errs, ok := err.(validator.ValidationErrors); ok {
        for _, e := range errs {
            field := strings.ToLower(e.Field())
            var msg string
            switch e.Tag() {
            case "required":
                msg = field + " is required"
            case "min":
                msg = field + " must be at least " + e.Param() + " characters"
            case "max":
                msg = field + " cannot be longer than " + e.Param() + " characters"
            case "email":
                msg = field + " must be a valid email"
            case "gte":
                msg = field + " must be greater than or equal to " + e.Param()
            case "lte":
                msg = field + " must be less than or equal to " + e.Param()
            default:
                msg = field + " is invalid"
            }
            errors[field] = msg
        }
    } else {
        // generic error
        errors["error"] = err.Error()
    }

    return errors
}


