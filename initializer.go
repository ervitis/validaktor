package validaktor

import "strings"

type initValidaktor struct{}

func (vldkini *initValidaktor) initializeValidators(tags ...string) map[string]validator {
	if len(tags) > 0 {
		return vldkini.initWithArguments(tags...)
	}
	return vldkini.initWithoutArguments()
}

// Function to use for the user
func (vldkini *initValidaktor) initWithArguments(tags ...string) map[string]validator {
	return map[string]validator{
		"regex":  &regexValidator{regex: strings.Split(tags[0], "=")[1]},
	}
}

// Function to use for the user
func (vldkini *initValidaktor) initWithoutArguments() map[string]validator {
	return map[string]validator{
		"struct": &structValidator{v: NewValidaktor()},
	}
}
