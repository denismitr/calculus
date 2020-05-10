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
		{name: "callable-1", expr: "5 + (4 + 10)", result: "19"},
	}

	c := New(InitializeDefaultGrammar(), InitializeDefaultEvaluator())

	for _, tc := range tt {
		r, err := c.calculate(tc.expr)

		assert.NoError(t, err)
		assert.Equal(t, tc.result, r)
	}
}
