package validaktor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type (
	myValidator struct{}
)

func (v *myValidator) validate(d interface{}) (bool, error) {return false, nil}

func TestGetValidatorType(t *testing.T) {
	mockito := &mockValidaktor{}
	mockito.On("initializeValidators", []string{"exp=[A-Z]{3}"}).Return(map[string]validator{"custom": &myValidator{}})

	validaktor := &validaktor{mockito}

	assert.IsType(t, &myValidator{}, validaktor.getValidator("custom,exp=[A-Z]{3}"))
}

func TestGetValidator(t *testing.T) {
	mockito := &mockValidaktor{}
	mockito.On("initializeValidators", []string{"exp=[A-Z]{3}"}).Return(map[string]validator{"custom": &myValidator{}})

	validaktor := &validaktor{mockito}

	assert.NotEqual(t, validaktor.getValidator("custom,exp=[A-Z]{3}"), &notImplementedValidator{})
}

func TestGetValidatorNotImplemented(t *testing.T) {
	mockito := &mockValidaktor{}
	mockito.On("initializeValidators", []string{"exp=[A-Z]{3}"}).Return(map[string]validator{"custom": &myValidator{}})

	validaktor := &validaktor{mockito}

	assert.Equal(t, &notImplementedValidator{tag: "another"}, validaktor.getValidator("another,exp=[A-Z]{3}"))

	_, err := validaktor.getValidator("another,exp=[A-Z]{3}").validate(nil)
	assert.Errorf(t, err, "nil validator or not implemented for tag another")
}

func TestNewValidaktor(t *testing.T) {
	validaktor := NewValidaktor()

	assert.NotNil(t, validaktor)
}

func TestValidateStruct(t *testing.T) {
	type fake struct {
		Name string `json:"name"`
		Age  string `json:"age" validate:"regex,exp=[0-9]{2}"`
	}

	f := fake{Name: "test", Age: "12"}

	validaktor := NewValidaktor()

	errors := validaktor.ValidateData(f)
	assert.Len(t, errors, 0)

	errors = validaktor.ValidateData(&f)
	assert.Len(t, errors, 0)

	f.Age = "3"
	errors = validaktor.ValidateData(f)
	assert.Len(t, errors, 1)
}
