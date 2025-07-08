package errors

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"medical-record-api/internal/constant"
	"medical-record-api/pkg/response"
	"net/http"
	"reflect"
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("json")
		if name == "-" {
			return ""
		}
		return name
	})
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errs := c.Errors
		if len(errs) > 0 {
			err := errs.Last().Err

			switch e := err.(type) {
			case validator.ValidationErrors:
				validationErrors := make(map[string]string)
				for _, ve := range e {
					field := ve.Field()
					validationErrors[field] = getErrorMsg(ve)
				}
				response.Respond(c, http.StatusBadRequest, false, constant.MsgValidationFailed, validationErrors, nil)
				return

			default:
				response.Respond(c, http.StatusInternalServerError, false, err.Error(), nil, nil)
				return
			}
		}
	}
}

func getErrorMsg(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return constant.MsgValidationRequired
	case "email":
		return constant.MsgValidationEmail
	case "min":
		return constant.MsgValidationMin + e.Param()
	case "max":
		return constant.MsgValidationMax + e.Param()
	default:
		return constant.MsgValidationInvalidVal
	}
}
