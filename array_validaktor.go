package validaktor

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type (
	arrayValidator struct {
		arrayOptions
		regx *regexp.Regexp
	}

	arrayError struct {
		message string
	}

	arrayOptions struct {
		min, max int
		isEmpty  bool
	}

	ArrayOption interface {
		apply(*arrayOptions)
	}

	fnArrayOptions func(*arrayOptions)
)

func (f fnArrayOptions) apply(options *arrayOptions) {
	f(options)
}

func WithMin(n int) ArrayOption {
	return fnArrayOptions(func(options *arrayOptions) {
		options.min = n
	})
}

func WithMax(n int) ArrayOption {
	return fnArrayOptions(func(options *arrayOptions) {
		options.max = n
	})
}

func WithIsEmpty(isEmpty bool) ArrayOption {
	return fnArrayOptions(func(options *arrayOptions) {
		options.isEmpty = isEmpty
	})
}

func NewArrayValidator(options ...ArrayOption) *arrayValidator {
	opts := &arrayOptions{
		min:     0,
		max:     0,
		isEmpty: false,
	}
	for _, opt := range options {
		opt.apply(opts)
	}
	return &arrayValidator{
		arrayOptions: *opts,
		regx:         regexp.MustCompile(`(min=(?P<min>\d+))?,?(max=(?P<max>\d+))?,?(contentType=(?P<contentType>(number|string|boolean)))?,?(isEmpty=(?P<isEmpty>(true|false)))?,?`),
	}
}

func mapSubexpNames(m, n []string) map[string]string {
	m, n = m[1:], n[1:]
	r := make(map[string]string, len(m))
	for i := range n {
		if n[i] != "" {
			r[n[i]] = m[i]
		}
	}

	return r
}

func newArrayValidatorError(message string) *arrayError {
	return &arrayError{message: message}
}

func (e *arrayError) Error() string {
	return e.message
}

func (v *arrayValidator) applyValidatorOptions(args ...string) error {
	m := v.regx.FindStringSubmatch(strings.Join(args[1:], ","))
	n := v.regx.SubexpNames()
	d := mapSubexpNames(m, n)

	b, err := strconv.ParseBool(d["isEmpty"])
	if err != nil {
		return fmt.Errorf("arrayValidator: apply options error parse isEmpty: %w", err)
	}

	mx, err := strconv.Atoi(d["max"])
	if err != nil {
		return fmt.Errorf("arrayValidator: apply options error parse max: %w", err)
	}

	mn, err := strconv.Atoi(d["min"])
	if err != nil {
		return fmt.Errorf("arrayValidator: apply options error parse min: %w", err)
	}

	v.min = mn
	v.max = mx
	v.isEmpty = b

	return nil
}

func (v *arrayValidator) validate(data interface{}) (bool, error) {
	var sl reflect.Value

	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		sl = reflect.ValueOf(data)
	default:
		return false, newArrayValidatorError("value is not an array")
	}

	if v.min > sl.Len() || sl.Len() > v.max {
		return false, newArrayValidatorError(fmt.Sprintf("the element has invalid length %d (min=%d, max=%d)",
			sl.Len(), v.min, v.max))
	}

	for i := 0; i < sl.Len(); i++ {
		if !v.isEmpty && isZeroOfUnderlyingType(sl.Index(i).Interface()) {
			return false, newArrayValidatorError(fmt.Sprintf("the element of index %d should not be empty", i+1))
		}
	}

	return true, nil
}

func isZeroOfUnderlyingType(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}
