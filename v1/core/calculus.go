package core

type Initializer func(p *parser, g *grammar)

type Calculator interface {
	Calculate(string) (string, error)
}

type Calculus struct {
	p *parser
	e *evaluator
}

func New(gInit Initializer, eInit EvaluatorInitializer) *Calculus {
	g := newGrammar()
	l := newLexer()
	p := newParser(l, g)
	e := newEvaluator()

	gInit(p,g)
	eInit(e)

	return &Calculus{p: p, e: e}
}

func (c *Calculus) calculate(in string) (string, error) {
	node, err := c.p.Parse(in)
	if err != nil {
		return "", err
	}

	result, err := node.eval(c.e)

	if err != nil {
		return "", err
	}

	return result.value, nil
}

func InitializeDefaultGrammar() Initializer {
	return func(p *parser, g *grammar) {
		g.prefix(LPAREN, 0, p.resolveParenthesis)
		g.prefix(IDENT, 0, p.resolveIdentifier)

		g.prefix(INT, 0, p.resolveInteger)
		g.prefix(ADD, 6, p.resolvePrefix)

		g.infix(ADD, 4, p.resolveLeftBinary)
		g.infix(INC, 7, p.resolvePostfix)

		g.infix(LPAREN, 8, p.resolveCallable)
	}
}

func InitializeDefaultEvaluator() EvaluatorInitializer {
	return func(e *evaluator) {
		e.binaryHandlers[ADD] = sum
	}
}
