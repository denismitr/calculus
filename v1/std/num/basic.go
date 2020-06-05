package num

import (
	"calculus/v1/core"
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
