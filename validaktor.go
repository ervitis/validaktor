package validaktor

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	tagName = "validate"
)

type (
	validator interface {
		validate(data interface{}) (bool, error)
	}

	notImplementedValidator struct{
		tag string
	}
	
	validaktor struct{
		initializer initializer
	}

	initializer interface {
		initializeValidators(tags ...string) validator
	}
)

func NewValidaktor() *validaktor {
	return &validaktor{initializer: &initValidaktor{}}
}

func (v *notImplementedValidator) validate(data interface{}) (bool, error) {
	return false, fmt.Errorf("nil validator or not implemented for tag %s", v.tag)
}

func (vldk *validaktor) getValidator(tag string) validator {
	tagArg := strings.Split(tag, ",")

	return vldk.initializer.initializeValidators(tagArg...)
}

func (vldk *validaktor) ValidateData(s interface{}) []error {
	var errs []error

	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get(tagName)
		if tag == "" || tag == "-" {
			continue
		}

		validator := vldk.getValidator(tag)
		if _, err := validator.validate(v.Field(i).Interface()); err != nil {
			errs = append(errs, fmt.Errorf("struct error in %s: %s", v.Type().Field(i).Name, err.Error()))
		}
	}

	return errs
}
