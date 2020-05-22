package core

import (
	"github.com/pkg/errors"
	"math"
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

func lg(tokens ...token) (result token, err error) {
	var precision int64

	if len(tokens) < 1 {
		err = errors.New("at least one argument is required for lg function")
		return
	}

	if len(tokens) > 2 {
		err = errors.New("lg function takes no more than 2 parameters")
		return
	}


	if len(tokens) == 2 {
		precision, err = strconv.ParseInt(tokens[1].value, 10, 8)
		if err != nil {
			return result, err
		}
	} else {
		precision = 5
	}

	f, err := strconv.ParseFloat(tokens[0].value, 64)
	if err != nil {
		return
	}

	result.kind = FLOAT
	result.value = strconv.FormatFloat(math.Log10(f), 'f', int(precision), 64)
	return
}
