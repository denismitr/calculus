package calculator

import (
	"calculus/v1/std"
	"github.com/stretchr/testify/assert"
	"testing"
)

type tStruct struct {
	name string
	expr string
	result string
}

func TestCalculusWithValidInputs(t *testing.T) {
	tt := []tStruct{
		{name: "1+1", expr: "1 + 1", result: "2"},
		{name: "20+4", expr: "20 + 4", result: "24"},
		{name: "20009+444", expr: "20009+444", result: "20453"},
		{name: "5 + (4 + 10)", expr: "5 + (4 + 10)", result: "19"},
		{name: "6 * 10", expr: "6 * 10", result: "60"},
		{name: "5 + 6 * 10", expr: "5 + 6 * 10", result: "65"},
		{name: "5 + (6 * 10 + 20)", expr: "5 + (6 * 10 + 20)", result: "85"},
		{name: "5 + (6 * 10 * 2)", expr: "5 + (6 * 10 * 2)", result: "125"},
		{name: "5 + (6 * 10 * 2 * 5)", expr: "5 + (6 * 10 * 2 * 5)", result: "605"},
		{name: "5 + (6 * 10 * 2 * 5) - 1", expr: "5 + (6 * 10 * 2 * 5) - 1", result: "604"},
		{name: "6 / 3", expr: "6 / 3", result: "2"},
		{name: "5++", expr: "5++", result: "6"},
		{name: "5++ + 8", expr: "5++ + 8", result: "14"},
		{name: "5--", expr: "5--", result: "4"},
	}

	c := New(std.Num)

	for _, tc := range tt {
		r, err := c.Calculate(tc.expr)
		assert.NoError(t, err)
		assert.Equal(t, tc.result, r)
	}
}

func TestStdNumFunctionsTest(t *testing.T) {
	tt := []tStruct{
		{name: "lg(2)", expr: "lg(2)", result: "0.30103"},
		{name: "lg(2, 1)", expr: "lg(2, 1)", result: "0.3"},
		{name: "lg(8, 3)", expr: "lg(8, 3)", result: "0.903"},
		{name: "pow(2,2)", expr: "pow(2,2)", result: "4"},
		{name: "pow(2,2) + 3", expr: "pow(2,2) + 3", result: "7"},
		{name: "pow(10,5) + 1", expr: "pow(10,5) + 1", result: "100001"},
		{name: "pow(4.76,3,14)", expr: "pow(4.76,3,14)", result: "107.85017599999999"},
		{name: "pow(4.76,3)", expr: "pow(4.76,3)", result: "107.85018"},
	}

	c := New(std.Num)

	for _, tc := range tt {
		r, err := c.Calculate(tc.expr)
		assert.NoError(t, err)
		assert.Equal(t, tc.result, r)
	}
}
