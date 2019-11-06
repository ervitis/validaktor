package validaktor

import (
	"regexp"
	"strconv"
	"strings"
)

type initValidaktor struct{}

func (vldkini *initValidaktor) initializeValidators(tags ...string) validator {
	switch tags[0] {
	case "regex":
		return &regexValidator{regex: strings.Split(tags[1], "=")[1]}
	case "array":
		r := regexp.MustCompile(`(min=(?P<min>\d+))?,?(max=(?P<max>\d+))?,?(contentType=(?P<contentType>(number|string|boolean)))?,?(isEmpty=(?P<isEmpty>(true|false)))?,?`)
		m := r.FindStringSubmatch(strings.Join(tags[1:], ","))
		n := r.SubexpNames()
		d := mapSubexpNames(m, n)
		b, _ := strconv.ParseBool(d["isEmpty"])
		mx, _ := strconv.Atoi(d["max"])
		mn, _ := strconv.Atoi(d["min"])
		return NewArrayValidator(WithIsEmpty(b), WithMax(mx), WithMin(mn))
	case "struct":
		return &structValidator{v: NewValidaktor()}
	default:
		return &notImplementedValidator{tag: tags[0]}
	}
}

func mapSubexpNames(m, n []string) map[string]string {
	m, n = m[1:], n[1:]
	r := make(map[string]string, len(m))
	for i := range n {
		if n[i] != "" {
			r[n[i]] = m[i]
		}
	}

	return r
}
