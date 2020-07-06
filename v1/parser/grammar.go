package parser

import "calculus/v1/core"

type Initializer func(p *Parser, g *grammar)

type grammar struct {
	prefixPrecedence map[core.Kind]int
	infixPrecedence map[core.Kind]int
	prefixParsers map[core.Kind]prefixParser
	infixParsers map[core.Kind]infixParser
}

func (g *grammar) prefix(k core.Kind, precedence int, parsers ...prefixParser) {
	for i := range parsers {
		g.prefixParsers[k] = parsers[i]
		g.prefixPrecedence[k] = precedence
	}
}

func (g *grammar) infix(k core.Kind, precedence int, parsers ...infixParser) {
	for i := range parsers {
		g.infixParsers[k] = parsers[i]
		g.infixPrecedence[k] = precedence
	}
}

func newGrammar() *grammar {
	return &grammar{
		prefixPrecedence: make(map[core.Kind]int),
		infixPrecedence: make(map[core.Kind]int),
		prefixParsers: make(map[core.Kind]prefixParser),
		infixParsers: make(map[core.Kind]infixParser),
	}
}

func DefaultGrammar() Initializer {
	return func(p *Parser, g *grammar) {
		g.prefix(core.LPAREN, 0, p.resolveParenthesis)
		g.prefix(core.IDENT, 0, p.resolveIdentifier)

		g.prefix(core.INT, 0, p.resolveInteger)
		g.prefix(core.FLOAT, 0, p.resolveInteger)
		g.prefix(core.ADD, 6, p.resolvePrefix)
		g.prefix(core.SUB, 6, p.resolvePrefix)

		g.infix(core.ADD, 4, p.resolveLeftBinary)
		g.infix(core.SUB, 4, p.resolveLeftBinary)

		g.infix(core.MUL, 5, p.resolveLeftBinary)
		g.infix(core.DIV, 5, p.resolveLeftBinary)

		g.infix(core.INC, 7, p.resolvePostfix)
		g.infix(core.DEC, 7, p.resolvePostfix)

		g.infix(core.LPAREN, 8, p.resolveCallable)
	}
}


