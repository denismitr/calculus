package core

import (
	"github.com/pkg/errors"
	"strconv"
)

func sum(l, r token) (token, error) {
	var result token
	if l.kind == INT && r.kind == INT {
		lVal, err := strconv.Atoi(l.value)
		if err != nil {
			return result, errors.Errorf("cannot covert %s to integeer", l.value)
		}
		rVal, err := strconv.Atoi(r.value)
		if err != nil {
			return result, errors.Errorf("cannot covert %s to integeer", r.value)
		}

		return token{value: strconv.Itoa(lVal+rVal), kind: INT}, nil
	}

	return result, errors.Errorf("could not add %s to %s", l.value, r.value)
}
