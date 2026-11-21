package validation

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type Validator struct{ validate *validator.Validate }

func New() *Validator {
	validate := validator.New()
	if err := validate.RegisterValidation("iyziDate", func(fl validator.FieldLevel) bool {
		dateFormat := "2006-01-02 15:04:05"
		dateStr := fl.Field().String()
		if dateStr == "" {
			return true
		}
		_, err := time.Parse(dateFormat, dateStr)
		return err == nil
	}); err != nil {
		panic(err)
	}
	if err := validate.RegisterValidation("httpsurl", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		parsedURL, err := url.Parse(field)
		if err != nil {
			return false
		}
		return parsedURL.Scheme == "https"
	}); err != nil {
		panic(err)
	}
	if err := validate.RegisterValidation("creditcard", func(fl validator.FieldLevel) bool {
		cardNumber := fl.Field().String()
		cleanNumber := ""
		for _, ch := range cardNumber {
			if ch >= '0' && ch <= '9' {
				cleanNumber += string(ch)
			}
		}

		if len(cleanNumber) < 13 || len(cleanNumber) > 19 {
			return false
		}

		sum := 0
		shouldDouble := false
		for i := len(cleanNumber) - 1; i >= 0; i-- {
			digit := int(cleanNumber[i] - '0')
			if shouldDouble {
				digit *= 2
				if digit > 9 {
					digit -= 9
				}
			}
			sum += digit
			shouldDouble = !shouldDouble
		}

		return sum%10 == 0
	}); err != nil {
		panic(err)
	}

	return &Validator{validate: validate}
}

// Struct, ortak kurallarla doğrulama yapar ve alan hatalarını aynı biçimde döndürür.
func (v *Validator) Struct(request any) error {
	err := v.validate.Struct(request)
	if err == nil {
		return nil
	}
	var fields validator.ValidationErrors
	if !errors.As(err, &fields) {
		return err
	}
	messages := make([]string, 0, len(fields))
	for _, field := range fields {
		messages = append(messages, fmt.Sprintf("Validation error at field '%s' with tag '%s'", field.Field(), field.Tag()))
	}
	return &fieldError{message: "validation errors: " + strings.Join(messages, ", "), cause: err}
}

type fieldError struct {
	message string
	cause   error
}

func (e *fieldError) Error() string { return e.message }
func (e *fieldError) Unwrap() error { return e.cause }
