package validaktor

func validators() map[string]validator {
	return map[string]validator {
		"regex": &regexValidator{},
		"array": &arrayValidator{},
		"struct": &structValidator{},
	}
}
