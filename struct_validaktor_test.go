package validaktor

import "testing"

type testStruct struct {
	isValid bool
	err     error
	data    interface{}
}

func TestStructValidatorOk(t *testing.T) {
	type (
		st1 struct{ id int }
	)

	testData := []testStruct{
		{data: &st1{id: 1}, isValid: true, err: nil},
		{data: &st1{}, isValid: true, err: nil},
	}

	for _, v := range testData {
		validator := &structValidator{}
		isValid, err := validator.validate(v.data)
		if v.isValid != isValid {
			t.Errorf("%+v != %+v it should be valid with data %+v", v.isValid, isValid, v.data)
		}
		if err != v.err {
			t.Errorf("there was an error %s", err)
		}
	}
}

func TestStructValidatorKo(t *testing.T) {
	type (
		st1 struct{ st1 *struct{} }
	)

	testData := []testStruct{
		{data: &st1{st1: nil}, isValid: true, err: nil},
	}

	for _, v := range testData {
		validator := &structValidator{}
		isValid, _ := validator.validate(v.data)
		if v.isValid != isValid {
			t.Errorf("%+v != %+v it should be valid with data %+v", v.isValid, isValid, v.data)
		}
	}
}
