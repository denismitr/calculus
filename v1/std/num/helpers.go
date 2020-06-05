package num

import (
	"calculus/v1/core"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

const standardPrecision = 5

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

func convertLeftAndRightToFloats(l, r string) (lf float64, rf float64, err error) {
	lf, err = strconv.ParseFloat(l, 64)
	if err != nil {
		return
	}

	rf, err = strconv.ParseFloat(r, 64)
	if err != nil {
		return
	}

	return
}

func deriveValueAndKindFromFloat(f float64, precision int) (v string, k core.Kind) {
	v = strconv.FormatFloat(f, 'f', precision, 64)

	if strings.Contains(v, ".00000") {
		v = strings.ReplaceAll(v, ".00000", "")
		k = core.INT
	} else {
		k = core.FLOAT
	}

	return
}

func leftAndRightOperation(l, r core.Token, op func(l, r float64) float64) (result core.Token, err error) {
	result = core.Token{Kind: core.ILLEGAL, Value: ""}
	lf, rf, err := convertLeftAndRightToFloats(l.Value, r.Value)
	if err != nil {
		return result, errors.Wrapf(err, "could not add %s to %s", l.Value, r.Value)
	}

	operationResult := op(lf, rf)

	v, k := deriveValueAndKindFromFloat(operationResult, standardPrecision)
	result.Kind = k
	result.Value = v
	return
}

func applyUnaryOperationOnFloat(t core.Token, applier func (float64) float64) (core.Token, error) {
	var result core.Token
	var k core.Kind
	f, err := strconv.ParseFloat(t.Value, 64)
	if err != nil {
		return result, err
	}

	f = applier(f)

	v, k := deriveValueAndKindFromFloat(f, standardPrecision)

	result.Value = v
	result.Kind = k

	return result, nil
}

func convertTokenToIntOrDefault(t core.Token, def int) (int, error) {
	p, err := strconv.ParseInt(t.Value, 10, 8)
	if err != nil {
		return def, err
	}

	return int(p), nil
}
