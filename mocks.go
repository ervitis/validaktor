package validaktor

import "github.com/stretchr/testify/mock"

type (
	mockValidaktor struct {
		mock.Mock
	}
)

func (m *mockValidaktor) initializeValidators(tags ...string) validator {
	args := m.Called(tags)
	return args.Get(0).(validator)
}
