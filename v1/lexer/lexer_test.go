package lexer

import (
	"calculus/v1/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLexer(t *testing.T) {
	tt := []struct{
		name string
		input string
		tokens []core.Token
	}{
		{name: "empty-string", input: "     ", tokens: nil},
		{name: "single-digit", input: "1", tokens: []core.Token{{Kind: core.INT, Value: "1"}}},
		{name: "triple-digit", input: "123   ", tokens: []core.Token{{Kind: core.INT, Value: "123"}}},
		{name: "two-triple-digits", input: "123   567", tokens: []core.Token{
			{Kind: core.INT, Value: "123"},
			{Kind: core.INT, Value: "567"},
		}},
		{name: "add-two-ints", input: "789 +47", tokens: []core.Token{
			{Kind: core.INT, Value: "789"},
			{Kind: core.ADD, Value: "+"},
			{Kind: core.INT, Value: "47"},
		}},
		{name: "add-and-sub-some-ints", input: "7 - 1+22-334+99", tokens: []core.Token{
			{Kind: core.INT, Value: "7"},
			{Kind: core.SUB, Value: "-"},
			{Kind: core.INT, Value: "1"},
			{Kind: core.ADD, Value: "+"},
			{Kind: core.INT, Value: "22"},
			{Kind: core.SUB, Value: "-"},
			{Kind: core.INT, Value: "334"},
			{Kind: core.ADD, Value: "+"},
			{Kind: core.INT, Value: "99"},
		}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			l := New()
			err := l.Tokenize(tc.input)

			assert.NoError(t, err)
			assert.Equal(t, tc.tokens, l.tokens)
		})
	}
}
