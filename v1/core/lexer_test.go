package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLexer(t *testing.T) {
	tt := []struct{
		name string
		input string
		tokens []token
	}{
		{name: "empty-string", input: "     ", tokens: nil},
		{name: "single-digit", input: "1", tokens: []token{{kind: INT, value: "1"}}},
		{name: "triple-digit", input: "123   ", tokens: []token{{kind: INT, value: "123"}}},
		{name: "two-triple-digits", input: "123   567", tokens: []token{
			{kind: INT, value: "123"},
			{kind: INT, value: "567"},
		}},
		{name: "add-two-ints", input: "789 +47", tokens: []token{
			{kind: INT, value: "789"},
			{kind: ADD, value: "+"},
			{kind: INT, value: "47"},
		}},
		{name: "add-and-sub-some-ints", input: "7 - 1+22-334+99", tokens: []token{
			{kind: INT, value: "7"},
			{kind: SUB, value: "-"},
			{kind: INT, value: "1"},
			{kind: ADD, value: "+"},
			{kind: INT, value: "22"},
			{kind: SUB, value: "-"},
			{kind: INT, value: "334"},
			{kind: ADD, value: "+"},
			{kind: INT, value: "99"},
		}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tokens, err := tokenize(tc.input)

			assert.NoError(t, err)
			assert.Equal(t, tc.tokens, tokens)
		})
	}
}
