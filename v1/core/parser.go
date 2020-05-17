package core

import "github.com/pkg/errors"

var ErrNoMoreTokens = errors.New("no more tokens")
var ErrUnexpectedToken = errors.New("unexpected token")

type prefixParser func(token) (node, error)
type infixParser func(left node, tok token) (node, error)

type parser struct {
	l *lexer
	g *grammar
}

func newParser(l *lexer, g *grammar) *parser {
	return &parser{l: l, g: g}
}

func (p *parser) Parse(input string) (node, error) {
	if err := p.l.tokenize(input); err != nil {
		return nil, err
	}

	result, err := p.parse(0)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *parser) parse(precedence int) (node, error) {
	tok, ok := p.l.pop()
	if !ok {
		return nil, ErrNoMoreTokens
	}

	prefix, ok := p.g.prefixParsers[tok.kind]
	if !ok {
		return nil, errors.Wrapf(ErrUnexpectedToken, "%s token cannot be parsed", tok.value)
	}

	left, err := prefix(tok)
	if err != nil {
		return nil, err
	}

	for  {
		next, ok := p.l.peek();
		if !ok {
			break
		}

		nextPrecedence := p.g.infixPrecedence[next.kind]
		if precedence >= nextPrecedence {
			break
		}

		tok, ok = p.l.pop()
		if !ok {
			break
		}

		infix, ok := p.g.infixParsers[tok.kind]
		if !ok {
			return nil, errors.Wrapf(ErrUnexpectedToken, "%s cannot be parsed", tok.value)
		}

		left , err = infix(left, tok)
		if err != nil {
			return nil, err
		}
	}

	return left, nil
}

func (p *parser) resolveParenthesis(tok token) (node, error) {
	exp, err := p.parse(0)
	if err != nil {
		return nil, err
	}

	if err :=  p.expect(RPAREN); err != nil {
		return nil, err
	}

	return exp, nil
}

func (p *parser) expect(want kind) error {
	have, ok := p.l.peek()
	if !ok {
		return ErrNoMoreTokens
	}

	if want != have.kind {
		return errors.Wrapf(ErrUnexpectedToken, "want %v, have %v", want, have)
	}

	p.l.pop()
	return nil
}

func (p *parser) resolveIdentifier(t token) (node, error) {
	return &identifier{t: t}, nil
}

func (p *parser) resolveLeftBinary(left node, t token) (node, error) {
	right, err := p.parse(p.g.infixPrecedence[t.kind])
	if err != nil {
		return nil, err
	}

	return &binary{op: t, left: left, right: right}, nil
}

func (p *parser) resolveRightBinary(left node, t token) (node, error) {
	right, err := p.parse(p.g.infixPrecedence[t.kind] - 1)
	if err != nil {
		return nil, err
	}

	return &binary{op: t, left: left, right: right}, nil
}

func (p *parser) resolvePrefix(t token) (node, error) {
	arg, err := p.parse(p.g.prefixPrecedence[t.kind])
	if err != nil {
		return nil, err
	}

	return &prefix{op: t, arg: arg}, nil
}

func (p* parser) resolvePostfix(left node, t token) (node, error) {
	return &postfix{op: t, arg: left}, nil
}

func (p *parser) resolveInteger(t token) (node, error) {
	return &integer{t}, nil
}

func (p *parser) resolveCallable(left node, t token) (node, error) {
	next, ok := p.l.peek()
	if !ok {
		return nil, ErrNoMoreTokens
	}

	if next.kind == RPAREN {
		// A call without arguments.
		p.l.pop()
		return &callable{fn: left}, nil
	}

	var args []node
	for {
		arg, err := p.parse(0)
		if err != nil {
			return nil, err
		}

		args = append(args, arg)

		next, ok := p.l.peek()
		if !ok || next.kind != COMMA {
			break
		}

		p.l.pop()
	}

	if err := p.expect(RPAREN); err != nil {
		return nil, err
	}

	return &callable{fn: left, args: args}, nil
}
