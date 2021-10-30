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
		validate(interface{}) (bool, error)
		applyValidatorOptions(...string) error
	}

	tag struct {
		validatorType string
		data          interface{}
	}

	validaktor struct {
		v map[string]validator
	}
)

func NewValidaktor() *validaktor {
	return &validaktor{v: validators()}
}

func (vldk *validaktor) getValidator(tag string) (validator, error) {
	tags := strings.Split(tag, ",")
	if len(tags) < 1 {
		return nil, fmt.Errorf("validate: tags structure is not correct")
	}

	if v, ok := vldk.v[tags[0]]; !ok {
		return nil, fmt.Errorf("validator not found")
	} else {
		return v, nil
	}
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

		data := strings.Split(tag, ",")
		if len(data) != 2 {
			errs = append(errs, fmt.Errorf("incorrect use of tag"))
			return errs
		}

		tag = data[0]
		args := data[1]

		validator, err := vldk.getValidator(tag)
		if err != nil {
			errs = append(errs, fmt.Errorf("validate: %w", err))
			return errs
		}

		if err := validator.applyValidatorOptions(args); err != nil {
			errs = append(errs, fmt.Errorf("validate: %w", err))
			return errs
		}

		if _, err := validator.validate(v.Field(i).Interface()); err != nil {
			errs = append(errs, fmt.Errorf("validate: error in %s: %w", v.Type().Field(i).Name, err))
		}
	}

	return errs
}
