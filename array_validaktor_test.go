package validaktor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type args struct {
	min, max int
	isEmpty  bool
}

type testArray struct {
	args     args
	isValid  bool
	err      error
	data     interface{}
	testName string
}

func TestArrayValidaktorOk(t *testing.T) {
	testData := []testArray{
		{testName: "testInt", args: args{min: 2, max: 5, isEmpty: false}, isValid: true, err: nil, data: []int{4, 3, 2, 1}},
		{testName: "testString", args: args{min: 1, max: 5, isEmpty: true}, isValid: true, err: nil, data: []string{"what", ""}},
	}

	for _, v := range testData {
		t.Run(v.testName, func(t *testing.T) {
			validator := NewArrayValidator(WithMin(v.args.min), WithMax(v.args.max), WithIsEmpty(v.args.isEmpty))
			isValid, err := validator.validate(v.data)
			if v.isValid != isValid {
				t.Errorf("%+v != %+v it should be valid with data %+v", v.isValid, isValid, v.data)
			}
			if err != v.err {
				t.Errorf("there was an error %s", err)
			}
		})
	}
}

func TestArrayValidaktorKo(t *testing.T) {
	testData := []testArray{
		{testName: "testInt", args: args{min: 2, max: 3, isEmpty: false}, isValid: false, err: &arrayError{message: "the element has invalid length 4 (min=2, max=3)"}, data: []int{4, 3, 2, 1}},
		{testName: "testString", args: args{min: 1, max: 5, isEmpty: false}, isValid: false, err: &arrayError{message: "the element of index 2 should not be empty"}, data: []string{"what", ""}},
	}

	for _, v := range testData {
		t.Run(v.testName, func(t *testing.T) {
			validator := NewArrayValidator(WithMin(v.args.min), WithMax(v.args.max), WithIsEmpty(v.args.isEmpty))
			isValid, err := validator.validate(v.data)
			if v.isValid != isValid {
				t.Errorf("%+v != %+v it should be valid with data %+v", v.isValid, isValid, v.data)
			}

			assert.EqualError(t, err, v.err.Error())
		})
	}
}
