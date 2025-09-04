package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_crud/validators"
	"io"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ValidateBody validates request body and provides professional errors
func ValidateBody(obj interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {

		decoder := json.NewDecoder(c.Request.Body)
		decoder.DisallowUnknownFields() // extra fields are errors

		if err := decoder.Decode(obj); err != nil {

			var syntaxErr *json.SyntaxError
			var typeErr *json.UnmarshalTypeError

			switch {
			case err == io.EOF:
				c.JSON(http.StatusBadRequest, gin.H{
					"errors": map[string]string{"error": "Request body cannot be empty"},
				})
				c.Abort()
				return
			case errors.As(err, &syntaxErr):
				c.JSON(http.StatusBadRequest, gin.H{
					"errors": map[string]string{"error": "Invalid JSON format"},
				})
				c.Abort()
				return
			case errors.As(err, &typeErr):
				fieldName := typeErr.Field
				if structField, ok := reflect.TypeOf(obj).Elem().FieldByName(fieldName); ok {
					fieldType := structField.Type.Name()
					c.JSON(http.StatusBadRequest, gin.H{
						"errors": map[string]string{
							fieldName: fmt.Sprintf("%s must be of type %s", fieldName, fieldType),
						},
					})
				} else {
					c.JSON(http.StatusBadRequest, gin.H{
						"errors": map[string]string{
							fieldName: "Invalid field type",
						},
					})
				}
				c.Abort()
				return
			default:
				c.JSON(http.StatusBadRequest, gin.H{
					"errors": map[string]string{"error": err.Error()},
				})
				c.Abort()
				return
			}
		}

		// Use go-playground validator for struct validation
		validate := validator.New()
		if err := validate.Struct(obj); err != nil {
			if ve, ok := err.(validator.ValidationErrors); ok {
				c.JSON(http.StatusBadRequest, gin.H{
					"errors": validators.FormatValidationError(ve),
				})
				c.Abort()
				return
			}
		}

		c.Set("validatedBody", obj)
		c.Next()
	}
}
