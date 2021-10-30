package validaktor

import (
	"fmt"
	"reflect"
	"unsafe"
)

type (
	structValidator struct{
		v *validaktor
	}

	structError struct {
		message string
	}
)

func newStructValidatorError(message string) *structError {
	return &structError{message: message}
}

func (e *structError) Error() string {
	return e.message
}

func (v *structValidator) applyValidatorOptions(_ ...string) error {
	return nil
}

func (v *structValidator) validate(data interface{}) (bool, error) {
	dv := reflect.ValueOf(data)

	if dv.Kind() == reflect.Ptr {
		dv = dv.Elem()
	}

	for i := 0; i < dv.NumField(); i++ {
		tag := dv.Type().Field(i).Tag.Get(tagName)
		if tag == "" || tag == "-" {
			continue
		}

		if v.v == nil {
			v.v = NewValidaktor()
		}

		validator, err := v.v.getValidator(tag)
		if err != nil {
			return false, newStructValidatorError(fmt.Sprintf("struct validation: error getting validator: %s", err))
		}

		ifdata := dv.Field(i).Interface()
		if (*[2]uintptr)(unsafe.Pointer(&ifdata))[1] == 0 {
			return false, newStructValidatorError("the struct is empty or nil")
		}

		if _, err := validator.validate(ifdata); err != nil {
			return false, newStructValidatorError(fmt.Sprintf("struct validation: error validating: %s", err))
		}
	}

	return true, nil
}
