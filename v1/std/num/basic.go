package num

import (
	"calculus/v1/core"
	"github.com/pkg/errors"
	"math"
	"strconv"
)

func Inc(t core.Token) (core.Token, error) {
	return applyUnaryOperationOnFloat(t, func(f float64) float64 {
		f++
		return f
	})
}

func Dec(t core.Token) (core.Token, error) {
	return applyUnaryOperationOnFloat(t, func(f float64) float64 {
		f--
		return f
	})
}

func Sum(l, r core.Token) (result core.Token, err error) {
	return leftAndRightOperation(l, r, func(lf, rf float64) float64 {
		return lf + rf
	})
}

func Mul(l, r core.Token) (core.Token, error) {
	return leftAndRightOperation(l, r, func(lf, rf float64) float64 {
		return lf * rf
	})
}

func Sub(l, r core.Token) (core.Token, error) {
	return leftAndRightOperation(l, r, func(lf, rf float64) float64 {
		return lf - rf
	})
}

func Div(l, r core.Token) (core.Token, error) {
	return leftAndRightOperation(l, r, func(lf, rf float64) float64 {
		return lf / rf
	})
}

func Lg(tokens ...core.Token) (result core.Token, err error) {
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
		precision, err = strconv.ParseInt(tokens[1].Value, 10, 8)
		if err != nil {
			return result, err
		}
	} else {
		precision = 5
	}

	f, err := strconv.ParseFloat(tokens[0].Value, 64)
	if err != nil {
		return
	}

	result.Kind = core.FLOAT
	result.Value = strconv.FormatFloat(math.Log10(f), 'f', int(precision), 64)
	return
}
