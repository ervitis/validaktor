package validaktor

import (
	"fmt"
	"reflect"
)

type (
	arrayValidator struct {
		arrayOptions
	}

	arrayError struct {
		message string
	}

	arrayOptions struct {
		min, max    int
		isEmpty     bool
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
		min:         0,
		max:         0,
		isEmpty:     false,
	}
	for _, opt := range options {
		opt.apply(opts)
	}
	return &arrayValidator{
		*opts,
	}
}

func newArrayValidatorError(message string) *arrayError {
	return &arrayError{message: message}
}

func (e *arrayError) Error() string {
	return e.message
}

func (v arrayValidator) validate(data interface{}) (bool, error) {
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
