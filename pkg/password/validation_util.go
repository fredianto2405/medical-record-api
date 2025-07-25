package password

import (
	"errors"
	"medical-record-api/internal/constant"
	"regexp"
)

func Validate(password string) error {
	if len(password) < 8 {
		return errors.New(constant.MsgValidationPasswordMin8)
	}

	lowercase := regexp.MustCompile(`[a-z]`)
	uppercase := regexp.MustCompile(`[A-Z]`)
	number := regexp.MustCompile(`[0-9]`)
	special := regexp.MustCompile(`[@#$_&\-+]`)

	switch {
	case !lowercase.MatchString(password):
		return errors.New(constant.MsgValidationPasswordLower)
	case !uppercase.MatchString(password):
		return errors.New(constant.MsgValidationPasswordUpper)
	case !number.MatchString(password):
		return errors.New(constant.MsgValidationPasswordNumeric)
	case !special.MatchString(password):
		return errors.New(constant.MsgValidationPasswordSpecialChar)
	}

	return nil
}
