package core

import (
	"regexp"
	"strings"
)

var numericRe = regexp.MustCompile("^[0-9]+$")

func isNumeric(s string) bool {
	if strings.Contains(s, " ") {
		return false
	}
	return numericRe.MatchString(s)
}
