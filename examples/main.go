package main

import (
	"fmt"
	"github.com/ervitis/validaktor"
)

type (
	User struct {
		Name string `validate:"regex,exp=[A-Z]+" json:"name"`
	}

	Company struct {
		Name string `validate:"regex,exp=[a-zA-Z]+"`
		User *User  `validate:"struct"`
	}

	BusinessSite struct {
		Name         string        `validate:"regex,exp=[a-zA-Z0-9]+"`
		BusinessSite *BusinessSite `validate:"struct"`
	}
)

func main() {
	user := &User{}
	validator := validaktor.NewValidaktor()

	errors := validator.ValidateData(user)

	for _, e := range errors {
		fmt.Println(e)
	}

	user = &User{Name: "ohmy"}

	errors = validator.ValidateData(user)

	for _, e := range errors {
		fmt.Println(e)
	}

	user2 := &User{Name: "OHMY"}

	errors = validator.ValidateData(user2)

	for _, e := range errors {
		fmt.Println(e)
	}

	company := &Company{
		Name: "something",
		User: &User{Name: "WATSON"},
	}

	errors = validator.ValidateData(company)
	fmt.Println(errors)

	businessSite := &BusinessSite{
		Name:         "B1",
		BusinessSite: &BusinessSite{
			Name:         "B12",
			BusinessSite: &BusinessSite{
				Name:         "B13",
				BusinessSite: nil,
			},
		},
	}

	errors = validator.ValidateData(businessSite)
	fmt.Println(errors)
}
