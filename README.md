# Validaktor: a Go library for validations

Validaktor is a struct type validator. It uses a custom tag to validate the structs.

## Getting Started

You can download the project using `go modules`

> go get -u github.com/ervitis/validaktor

Let's see an example of validation

```go
type User struct {
    Name string `validate:"regex,exp=[A-Z]+" json:"name"`
}

func main() {
    user := &User{Name: "HELLO"}

    validator := validaktor.NewValidaktor()

    errors := validator.ValidateData(user)

    for _, e := range errors {
        fmt.Println(e)
    }
}
```

You can see more examples inside the `exmples` folder.

You can add more validators sending a pull request or doing it in your local machine. For more information read the next point in this document.

### Adding more validators

Let's see an example. For example the `regex_validator` file

You should implement two structs:

```go
type (
	regexValidator struct {
		regex string
	}

	regexError struct {
		message string
	}
)
```

The first one should be used for initialization the validator. The second one for custom errors.

Now the error methods.

```go
func newRegexValidatorError(message string) *regexError {
	return &regexError{message: message}
}

func (e *regexError) Error() string {
	return e.message
}
```

Let's implement the `validate` method.

```go
func (v *regexValidator) validate(data interface{}) (bool, error) {
    // implement the validation. You can use here the custom errors created before
    // ...

    return true, nil
}
```

After the validator implementation, let's add it in the `initializer.go` file. Because the `regexValidator` is using a parameter we add it inside the function `initWithArguments`.

If it doesn't need any parameter, you should initialize the constructor inside the `initWithoutArguments` method.

After that, implement the unit tests.

### Installing

> go get -u github.com/ervitis/validaktor

## Running the tests

> go test -v -race ./...

### Break down into end to end tests

You can use any framework for testing. Let's see an example:

```go
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
```

We can prepare the test data using an struct and then iterate it inside the for loop.

## Built With

* [Golang](https://golang.org/) - Golang 1.12

## Contributing

Please read [CONTRIBUTING.md](./CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **ervitis** - *Initial work* - [ervitis](https://github.com/ervitis)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
