package validaktor

import "testing"

type testRegex struct {
	exp     string
	isValid bool
	err     error
	data    interface{}
}

func TestRegexValidate(t *testing.T) {
	testData := []testRegex{
		{exp: "[A-Z]+", isValid: true, err: nil, data: "HELLO"},
		{exp: "[0-9]{4,6}", isValid: true, err: nil, data: "12345"},
		{exp: "\\w+", isValid: true, err: nil, data: "whatever24"},
		{exp: `\w+`, isValid: true, err: nil, data: "iamgood"},
		{exp: "[^A-Z]+", isValid: true, err: nil, data: "123456asdf"},
	}

	for _, v := range testData {
		validator := &regexValidator{regex: v.exp}
		isValid, err := validator.validate(v.data)
		if v.isValid != isValid {
			t.Errorf("%+v != %+v it should be valid with data %+v", v.isValid, isValid, v.data)
		}
		if err != v.err {
			t.Errorf("there was an error %s", err)
		}
	}
}

func TestRegexValidateKo(t *testing.T) {
	testData := []testRegex{
		{exp: "[A-Z]+", isValid: false, data: "1234"},
		{exp: "^[0-9]{4,6}$", isValid: false, data: "123456789"},
		{exp: "^\\w+$", isValid: false, data: "whate ver24"},
		{exp: `^\w+$`, isValid: false, data: "  iamg ood"},
		{exp: "[^A-Z]+", isValid: false, data: "ASDFQWER"},
	}

	for _, v := range testData {
		validator := &regexValidator{regex: v.exp}
		isValid, _ := validator.validate(v.data)
		if v.isValid != isValid {
			t.Errorf("%+v != %+v it should not be valid with data %+v", v.isValid, isValid, v.data)
		}
	}
}
