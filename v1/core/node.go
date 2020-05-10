package core

import (
	"fmt"
	"strings"
)

type node interface {
	eval(*evaluator) (token, error)
	String() string
}

type identifier struct {
	t token
}

type prefix struct {
	op  token
	arg node
}

type postfix struct {
	op  token
	arg node
}

type binary struct {
	op    token
	left  node
	right node
}

type integer struct {
	t token
}

type callable struct {
	fn   node
	args []node
}

func (n *identifier) eval(_ *evaluator) (token, error) {
	return n.t, nil
}

func (n *prefix) eval(_ *evaluator) (token, error) {
	return n.op, nil
}

func (n *postfix) eval(_ *evaluator) (token, error) {
	return token{}, nil
}

func (b *binary) eval(e *evaluator) (token, error)  {
	lTok, err := b.left.eval(e)
	if err != nil {
		return token{}, err
	}

	rTok, err := b.right.eval(e)
	if err != nil {
		return token{}, err
	}

	return e.binary(b.op, lTok, rTok)
}

func (c *callable) eval(e *evaluator) (token, error)  {
	var tokens []token
	l, err := c.fn.eval(e)
	if err != nil {
		return token{}, nil
	}

	tokens = append(tokens, l)
	for i := range c.args {
		t, err := c.args[i].eval(e)
		if err != nil {
			return token{}, nil
		}

		tokens = append(tokens, t)
	}

	return e.multiple(tokens...)
}
func (n *integer) eval(_ *evaluator) (token, error) {
	return n.t, nil
}


func (n *identifier) String() string { return n.t.value }
func (n *integer) String() string    { return n.t.value }
func (p *prefix) String() string     { return fmt.Sprintf("(prefix %s %s)", p.op, p.arg) }
func (p *postfix) String() string { return fmt.Sprintf("(postfix %s %s)", p.op, p.arg) }
func (b *binary) String() string  { return fmt.Sprintf("(%s %s %s)", b.op, b.left, b.right) }

func (c *callable) String() string {
	parts := make([]string, 0, len(c.args) + 1)
	parts = append(parts, c.fn.String())
	for _, arg := range c.args {
		parts = append(parts, arg.String())
	}
	return fmt.Sprintf("(callable %s)", strings.Join(parts, " "))
}