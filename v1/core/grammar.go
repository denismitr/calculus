package core

type grammar struct {
	prefixPrecedence map[kind]int
	infixPrecedence map[kind]int
	prefixParsers map[kind]prefixParser
	infixParsers map[kind]infixParser
}

func (g *grammar) prefix(k kind, precedence int, parsers ...prefixParser) {
	for i := range parsers {
		g.prefixParsers[k] = parsers[i]
		g.prefixPrecedence[k] = precedence
	}
}

func (g *grammar) infix(k kind, precedence int, parsers ...infixParser) {
	for i := range parsers {
		g.infixParsers[k] = parsers[i]
		g.infixPrecedence[k] = precedence
	}
}

func newGrammar() *grammar {
	return &grammar{
		prefixPrecedence: make(map[kind]int),
		infixPrecedence: make(map[kind]int),
		prefixParsers: make(map[kind]prefixParser),
		infixParsers: make(map[kind]infixParser),
	}
}
