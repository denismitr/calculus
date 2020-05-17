package core

import (
	"github.com/pkg/errors"
	"strconv"
)

func sum(l, r token) (token, error) {
	var result token
	if l.kind == INT && r.kind == INT {
		lVal, rVal, err := convertLeftAndRightToIntegers(l.value, r.value)
		if err != nil {
			return result, err
		}

		return token{value: strconv.Itoa(lVal+rVal), kind: INT}, nil
	}

	return result, errors.Errorf("could not add %s to %s", l.value, r.value)
}

func mul(l, r token) (token, error) {
	var result token
	if l.kind == INT && r.kind == INT {
		lVal, rVal, err := convertLeftAndRightToIntegers(l.value, r.value)
		if err != nil {
			return result, err
		}

		return token{value: strconv.Itoa(lVal*rVal), kind: INT}, nil
	}

	return result, errors.Errorf("could not multiply %s by %s", l.value, r.value)
}

func sub(l, r token) (token, error) {
	var result token
	if l.kind == INT && r.kind == INT {
		lVal, rVal, err := convertLeftAndRightToIntegers(l.value, r.value)
		if err != nil {
			return result, err
		}

		return token{value: strconv.Itoa(lVal-rVal), kind: INT}, nil
	}

	return result, errors.Errorf("could not multiply %s by %s", l.value, r.value)
}

func convertLeftAndRightToIntegers(l, r string) (int, int, error) {
	lInt, err := strconv.Atoi(l)
	if err != nil {
		return 0, 0, errors.Errorf("cannot covert %s to integer", l)
	}

	rInt, err := strconv.Atoi(r)
	if err != nil {
		return 0, 0, errors.Errorf("cannot covert %s to integer", r)
	}

	return lInt, rInt, nil
}
