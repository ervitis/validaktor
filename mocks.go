package validaktor

import "github.com/stretchr/testify/mock"

type (
	mockValidaktor struct {
		mock.Mock
	}
)

func (m *mockValidaktor) initializeValidators(tags ...string) map[string]validator {
	args := m.Called(tags)
	return args.Get(0).(map[string]validator)
}

func (m *mockValidaktor) getValidator(tag string) validator {
	args := m.Called(0)
	return args.Get(0).(validator)
}

func (m *mockValidaktor) ValidateData(s interface{}) []error {
	args := m.Called(0)
	return args.Get(0).([]error)
}