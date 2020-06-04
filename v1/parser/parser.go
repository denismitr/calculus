package parser

import "calculus/v1/core"
import "github.com/pkg/errors"

type prefixParser func(core.Token) (core.Node, error)
type infixParser func(left core.Node, tok core.Token) (core.Node, error)

type Parser struct {
	l core.Lexer
	g *grammar
}

func New(l core.Lexer, init Initializer) *Parser {
	g := newGrammar()
	p := &Parser{l: l, g: g}
	init(p, g)
	return p
}


func (p *Parser) Parse(input string) (core.Node, error) {
	if err := p.l.Tokenize(input); err != nil {
		return nil, err
	}

	result, err := p.parse(0)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *Parser) parse(precedence int) (core.Node, error) {
	tok, ok := p.l.Pop()
	if !ok {
		return nil, core.ErrNoMoreTokens
	}

	prefix, ok := p.g.prefixParsers[tok.Kind]
	if !ok {
		return nil, errors.Wrapf(core.ErrUnexpectedToken, "%s Token cannot be parsed", tok.Value)
	}

	left, err := prefix(tok)
	if err != nil {
		return nil, err
	}

	for  {
		next, ok := p.l.Peek();
		if !ok {
			break
		}

		nextPrecedence := p.g.infixPrecedence[next.Kind]
		if precedence >= nextPrecedence {
			break
		}

		tok, ok = p.l.Pop()
		if !ok {
			break
		}

		infix, ok := p.g.infixParsers[tok.Kind]
		if !ok {
			return nil, errors.Wrapf(core.ErrUnexpectedToken, "%s cannot be parsed", tok.Value)
		}

		left , err = infix(left, tok)
		if err != nil {
			return nil, err
		}
	}

	return left, nil
}

func (p *Parser) resolveParenthesis(tok core.Token) (core.Node, error) {
	exp, err := p.parse(0)
	if err != nil {
		return nil, err
	}

	if err :=  p.expect(core.RPAREN); err != nil {
		return nil, err
	}

	return exp, nil
}

func (p *Parser) expect(want core.Kind) error {
	have, ok := p.l.Peek()
	if !ok {
		return core.ErrNoMoreTokens
	}

	if want != have.Kind {
		return errors.Wrapf(core.ErrUnexpectedToken, "want %v, have %v", want, have)
	}

	p.l.Pop()
	return nil
}

func (p *Parser) resolveIdentifier(t core.Token) (core.Node, error) {
	return &identifier{t: t}, nil
}

func (p *Parser) resolveLeftBinary(left core.Node, t core.Token) (core.Node, error) {
	right, err := p.parse(p.g.infixPrecedence[t.Kind])
	if err != nil {
		return nil, err
	}

	return &binary{op: t, left: left, right: right}, nil
}

func (p *Parser) resolveRightBinary(left core.Node, t core.Token) (core.Node, error) {
	right, err := p.parse(p.g.infixPrecedence[t.Kind] - 1)
	if err != nil {
		return nil, err
	}

	return &binary{op: t, left: left, right: right}, nil
}

func (p *Parser) resolvePrefix(t core.Token) (core.Node, error) {
	arg, err := p.parse(p.g.prefixPrecedence[t.Kind])
	if err != nil {
		return nil, err
	}

	return &prefix{op: t, arg: arg}, nil
}

func (p* Parser) resolvePostfix(left core.Node, t core.Token) (core.Node, error) {
	return &postfix{op: t, arg: left}, nil
}

func (p *Parser) resolveInteger(t core.Token) (core.Node, error) {
	return &integer{t}, nil
}

func (p *Parser) resolveCallable(left core.Node, t core.Token) (core.Node, error) {
	next, ok := p.l.Peek()
	if !ok {
		return nil, core.ErrNoMoreTokens
	}

	if next.Kind == core.RPAREN {
		// A call without arguments.
		p.l.Pop()
		return &callable{fn: left}, nil
	}

	var args []core.Node
	for {
		arg, err := p.parse(0)
		if err != nil {
			return nil, err
		}

		args = append(args, arg)

		next, ok := p.l.Peek()
		if !ok || next.Kind != core.COMMA {
			break
		}

		p.l.Pop()
	}

	if err := p.expect(core.RPAREN); err != nil {
		return nil, err
	}

	return &callable{fn: left, args: args}, nil
}

