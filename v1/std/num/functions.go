package num

import (
	"calculus/v1/core"
	"github.com/pkg/errors"
	"math"
	"strconv"
)

func Lg(tokens ...core.Token) (result core.Token, err error) {
	var precision int

	if len(tokens) < 1 {
		err = errors.New("at least one argument is required for lg function")
		return
	}

	if len(tokens) > 2 {
		err = errors.New("lg function takes no more than 2 parameters")
		return
	}

	if len(tokens) == 2 {
		precision, err = convertTokenToIntOrDefault(tokens[1], standardPrecision)
		if err != nil {
			return result, err
		}
	} else {
		precision = standardPrecision
	}

	f, err := strconv.ParseFloat(tokens[0].Value, 64)
	if err != nil {
		return
	}

	answer := math.Log10(f)
	v, k := deriveValueAndKindFromFloat(answer, precision)
	result.Kind = k
	result.Value = v
	return
}

func Pow(tokens ...core.Token) (result core.Token, err error) {
	var precision int

	if len(tokens) < 2 {
		err = errors.New("at least two arguments are required for pow() function")
		return
	}

	if len(tokens) > 3 {
		err = errors.New("pow() function takes no more than 3 parameters")
		return
	}

	if len(tokens) == 3 {
		precision, err = convertTokenToIntOrDefault(tokens[2], standardPrecision)
		if err != nil {
			return result, err
		}
	} else {
		precision = standardPrecision
	}

	x, err := strconv.ParseFloat(tokens[0].Value, 64)
	if err != nil {
		return
	}

	y, err := strconv.ParseFloat(tokens[1].Value, 64)
	if err != nil {
		return
	}

	answer := math.Pow(x, y)
	v, k := deriveValueAndKindFromFloat(answer, precision)
	result.Kind = k
	result.Value = v
	return
}

func Root(tokens ...core.Token) (result core.Token, err error) {
	var precision int
	var answer float64
	var base float64 = 2

	if len(tokens) > 3 {
		err = errors.New("root() function takes no more than 3 parameters")
		return
	}

	if len(tokens) == 3 {
		precision, err = convertTokenToIntOrDefault(tokens[2], standardPrecision)
		if err != nil {
			return result, err
		}
	} else {
		precision = standardPrecision
	}

	if len(tokens) >= 2 {
		base, err = strconv.ParseFloat(tokens[1].Value, 64)
		if err != nil {
			return result, errors.Wrapf(core.ErrInvalidInput, "cannot convert %s to float", tokens[1].Value)
		}
	}

	target, err := strconv.ParseFloat(tokens[0].Value, 64)
	if err != nil {
		return
	}

	answer = math.Pow(target, 1 / base)

	v, k := deriveValueAndKindFromFloat(answer, precision)
	result.Kind = k
	result.Value = v
	return
}


