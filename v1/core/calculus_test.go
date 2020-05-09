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
	}

	c := New(InitializeDefaultGrammar())

	for _, tc := range tt {
		r, err := c.calculate(tc.expr)

		assert.NoError(t, err)
		assert.Equal(t, tc.result, r)
	}
}
