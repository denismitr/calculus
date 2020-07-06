package num

import (
	"calculus/utils/validator"
	"calculus/v1/core"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

const standardPrecision = 5

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

	if validator.IsRoundFloat(v) {
		v = strings.TrimRight(v, "0")
		v = strings.TrimRight(v, ".")
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
	p, err := strconv.Atoi(t.Value)
	if err != nil {
		return def, errors.Wrapf(core.ErrInvalidInput, "%s is not convertible to integer", t.Value)
	}

	return p, nil
}

func convertTokenToFloatOrDefault(t core.Token, def float64) (float64, error) {
	p, err := strconv.ParseFloat(t.Value, 64)
	if err != nil {
		return def, errors.Wrapf(core.ErrInvalidInput, "%s is not convertible to integer", t.Value)
	}

	return p, nil
}