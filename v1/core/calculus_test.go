package core

import (
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
	}

	c := New(InitializeDefaultGrammar(), InitializeDefaultEvaluator())

	for _, tc := range tt {
		r, err := c.calculate(tc.expr)

		assert.NoError(t, err)
		assert.Equal(t, tc.result, r)
	}
}
