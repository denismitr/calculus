package core

type ast struct {
	prefixPrecedence map[kind]int
	infixPrecedence map[kind]int
	nodes []node
}

func newAst() *ast {
	a := &ast{
		prefixPrecedence: make(map[kind]int),
		infixPrecedence: make(map[kind]int),
	}

	// todo: add prefixes

	return a
}

func (a *ast) parse(l *lexer) error {
	tok, ok := l.pop()
	if !ok {
		return nil
	}

	if tok.kind == INT {
		a.nodes = append(a.nodes, numericLiteral{token: tok})
	}

	return nil
}
