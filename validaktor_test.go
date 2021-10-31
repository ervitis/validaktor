package validaktor

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetValidatorType(t *testing.T) {
	validaktor := NewValidaktor()
	v, err := validaktor.getValidator("regex,exp=[A-Z]{3}")
	require.NoError(t, err)

	require.IsType(t, &regexValidator{}, v)
}

func TestGetValidatorNotImplemented(t *testing.T) {
	validaktor := NewValidaktor()
	_, err := validaktor.getValidator("errorCustom,exp=[A-Z]{3}")
	require.Error(t, err)
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
