package validaktor

import (
	"errors"
	"fmt"
	"regexp"
)

type (
	regexValidator struct {
		regex string
	}

	regexError struct {
		message string
	}
)

func newRegexValidatorError(message string) *regexError {
	return &regexError{message: message}
}

func (e *regexError) Error() string {
	return e.message
}

func (v *regexValidator) validate(data interface{}) (bool, error) {
	s, ok := data.(string)
	if !ok {
		return false, errors.New("data input is not valid")
	}

	rgx, err := regexp.Compile(v.regex)
	if err != nil {
		return false, err
	}

	if ok = rgx.MatchString(s); !ok {
		return false, newRegexValidatorError(fmt.Sprintf("%s not match in %s", s, v.regex))
	}

	return true, nil
}
