package validaktor

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type (
	regexValidator struct {
		regex *regexp.Regexp
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

func (v *regexValidator) applyValidatorOptions(args ...string) error {
	s := strings.Split(args[0], "=")
	if len(s) != 2 {
		return fmt.Errorf("regexValidator: apply options cannot find regex rules after '='")
	}

	rgx, err := regexp.Compile(s[1])
	if err != nil {
		return fmt.Errorf("regexValidator: apply options err: %w", err)
	}

	v.regex = rgx
	return nil
}

func (v *regexValidator) validate(data interface{}) (bool, error) {
	s, ok := data.(string)
	if !ok {
		return false, errors.New("data input is not valid")
	}

	if ok = v.regex.MatchString(s); !ok {
		return false, newRegexValidatorError(fmt.Sprintf("%s not match in %s", s, v.regex))
	}

	return true, nil
}
