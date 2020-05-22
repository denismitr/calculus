package core

import (
	"regexp"
)

var numericRe = regexp.MustCompile("^[0-9]+$")
var nameRe = regexp.MustCompile("^[a-zA-Z_]+$")

func isNumeric(s string) bool {
	return numericRe.MatchString(s)
}

func isName(s string) bool {
	return nameRe.MatchString(s)
}

func isOperator(r uint8) bool {
	switch r {
	case '-','+','*','/':
		return true
	default:
		return false
	}
}

func isParenthesis(r uint8) bool {
	return r == '(' || r == ')'
}
