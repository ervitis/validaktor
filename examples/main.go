package main

import (
	"fmt"
	"github.com/ervitis/validaktor"
)

type (
	User struct {
		Name string `validate:"regex,exp=[A-Z]+" json:"name"`
	}
)

func main() {
	user := &User{}
	validator := validaktor.NewValidaktor()

	errors := validator.ValidateStruct(user)

	for _, e := range errors {
		fmt.Println(e)
	}

	user = &User{Name: "ohmy"}

	errors = validator.ValidateStruct(user)

	for _, e := range errors {
		fmt.Println(e)
	}

	user2 := User{Name: "OHMY"}

	errors = validator.ValidateStruct(user2)

	for _, e := range errors {
		fmt.Println(e)
	}
}