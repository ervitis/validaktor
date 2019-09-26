package validaktor

import (
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

		validator := v.v.getValidator(tag)
		ifdata := dv.Field(i).Interface()
		if (*[2]uintptr)(unsafe.Pointer(&ifdata))[1] == 0 {
			return true, nil
		}

		if _, err := validator.validate(ifdata); err != nil {
			return false, err
		}
	}

	return true, nil
}
