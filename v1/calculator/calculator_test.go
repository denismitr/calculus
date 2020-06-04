package calculator

import (
	"calculus/v1/std"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculusWithValidInputs(t *testing.T) {
	tt := []struct{
		name string
		expr string
		result string
	}{
		{name: "one-plus-one", expr: "1 + 1", result: "2"},
		{name: "20+4", expr: "20 + 4", result: "24"},
		{name: "20009+444", expr: "20009+444", result: "20453"},
		{name: "multi-1", expr: "5 + (4 + 10)", result: "19"},
		{name: "mul", expr: "6 * 10", result: "60"},
		{name: "multi-2", expr: "5 + 6 * 10", result: "65"},
		{name: "multi-3", expr: "5 + (6 * 10 + 20)", result: "85"},
		{name: "multi-4", expr: "5 + (6 * 10 * 2)", result: "125"},
		{name: "multi-5", expr: "5 + (6 * 10 * 2 * 5)", result: "605"},
		{name: "multi-6", expr: "5 + (6 * 10 * 2 * 5) - 1", result: "604"},
		{name: "div-1", expr: "6 / 3", result: "2"},
		{name: "increment-1", expr: "5++", result: "6"},
		{name: "increment-2", expr: "5++ + 8", result: "14"},
		{name: "decrement-4", expr: "5--", result: "4"},
		{name: "lg-with-default-precision", expr: "lg(2)", result: "0.30103"},
		{name: "lg-with-precision-1", expr: "lg(2, 1)", result: "0.3"},
		{name: "lg-with-precision-3", expr: "lg(8, 3)", result: "0.903"},
	}

	c := New(std.Num)

	for _, tc := range tt {
		r, err := c.Calculate(tc.expr)

		assert.NoError(t, err)
		assert.Equal(t, tc.result, r)
	}
}
