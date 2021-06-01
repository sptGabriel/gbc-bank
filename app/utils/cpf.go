package utils

import (
	"bytes"
	"regexp"
	"unicode"
)

func Clean(cpf string) string {
	buf := bytes.NewBufferString("")
	for _, r := range cpf {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

func Format(cpf string) string {
	expr, err := regexp.Compile(`^([\d]{3})([\d]{3})([\d]{3})([\d]{2})$`)
	if err != nil {
		return cpf
	}
	return expr.ReplaceAllString(cpf, "$1.$2.$3-$4")
}
