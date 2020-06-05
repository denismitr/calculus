package lexer

import (
	"regexp"
)

var intRe = regexp.MustCompile("^[0-9]+$")
var floatRe = regexp.MustCompile("\\d+\\.?\\d*?$")
var nameRe = regexp.MustCompile("^[a-zA-Z_]+$")

func isInt(s string) bool {
	return intRe.MatchString(s)
}

func isFloat(s string) bool {
	return floatRe.MatchString(s)
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
