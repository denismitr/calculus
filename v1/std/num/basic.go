package num

import (
	"calculus/v1/core"
	"github.com/pkg/errors"
	"math"
	"strconv"
	"strings"
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

func applyUnaryOperationOnFloat(t core.Token, applier func (float64) float64) (core.Token, error) {
	var result core.Token
	var k core.Kind
	f, err := strconv.ParseFloat(t.Value, 64)
	if err != nil {
		return result, err
	}

	f = applier(f)

	v, k := deriveValueAndKind(f, k)

	result.Value = v
	result.Kind = k

	return result, nil
}

func deriveValueAndKind(f float64, k core.Kind) (string, core.Kind) {
	v := strconv.FormatFloat(f, 'f', 5, 64)
	k = core.FLOAT

	if strings.Contains(v, ".00000") {
		v = strings.ReplaceAll(v, ".00000", "")
		k = core.INT
	}
	return v, k
}

func Sum(l, r core.Token) (core.Token, error) {
	var result core.Token
	if l.Kind == core.INT && r.Kind == core.INT {
		lVal, rVal, err := convertLeftAndRightToIntegers(l.Value, r.Value)
		if err != nil {
			return result, err
		}

		return core.Token{Value: strconv.Itoa(lVal+rVal), Kind: core.INT}, nil
	}

	return result, errors.Errorf("could not add %s to %s", l.Value, r.Value)
}

func Mul(l, r core.Token) (core.Token, error) {
	var result core.Token
	if l.Kind == core.INT && r.Kind == core.INT {
		lVal, rVal, err := convertLeftAndRightToIntegers(l.Value, r.Value)
		if err != nil {
			return result, err
		}

		return core.Token{Value: strconv.Itoa(lVal*rVal), Kind: core.INT}, nil
	}

	return result, errors.Errorf("could not multiply %s by %s", l.Value, r.Value)
}

func Sub(l, r core.Token) (core.Token, error) {
	var result core.Token
	if l.Kind == core.INT && r.Kind == core.INT {
		lVal, rVal, err := convertLeftAndRightToIntegers(l.Value, r.Value)
		if err != nil {
			return result, err
		}

		return core.Token{Value: strconv.Itoa(lVal-rVal), Kind: core.INT}, nil
	}

	return result, errors.Errorf("could not multiply %s by %s", l.Value, r.Value)
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
