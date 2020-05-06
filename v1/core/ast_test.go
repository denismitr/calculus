package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAST(t *testing.T) {
	tt := []struct{
		name string
		tokens []token
		nodes []node
	}{
		{
			name: "single number",
			tokens: []token{{kind: INT, value: "1"}},
			nodes: []node{numericLiteral{token: token{kind: INT, value: "1"}}},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			a := newAst()
			l := newLexer()
			l.tokens = tc.tokens

			err := a.parse(l)
			assert.NoError(t, err)
			assert.Equal(t, a.nodes, tc.nodes)
		})
	}
}
