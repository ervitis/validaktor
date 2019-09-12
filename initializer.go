package validaktor

import "strings"

type initValidaktor struct{}

// Function to use for the user
func (vldkini *initValidaktor) initializeValidators(tags ...string) map[string]validator {
	return map[string]validator{
		"regex": &regexValidator{regex: strings.Split(tags[0], "=")[1]},
	}
}