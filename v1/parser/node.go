package parser

import (
	"calculus/v1/core"
	"fmt"
	"strings"
)

type identifier struct {
	t core.Token
}

type prefix struct {
	op  core.Token
	arg core.Node
}

type postfix struct {
	op  core.Token
	arg core.Node
}

type binary struct {
	op    core.Token
	left  core.Node
	right core.Node
}

type integer struct {
	t core.Token
}

type callable struct {
	fn   core.Node
	args []core.Node
}

func (n *identifier) Eval(_ core.Evaluator) (core.Token, error) {
	return n.t, nil
}

func (n *prefix) Eval(e core.Evaluator) (core.Token, error) {
	t, _ := n.arg.Eval(e)
	return e.Unary(n.op, t)
}

func (n *postfix) Eval(e core.Evaluator) (core.Token, error) {
	t, _ := n.arg.Eval(e)
	return e.Unary(n.op, t)
}

func (b *binary) Eval(e core.Evaluator) (core.Token, error)  {
	lTok, err := b.left.Eval(e)
	if err != nil {
		return core.Token{}, err
	}

	rTok, err := b.right.Eval(e)
	if err != nil {
		return core.Token{}, err
	}

	return e.Binary(b.op, lTok, rTok)
}

func (c *callable) Eval(e core.Evaluator) (core.Token, error)  {
	var tokens []core.Token
	l, err := c.fn.Eval(e)
	if err != nil {
		return core.Token{}, nil
	}

	for i := range c.args {
		t, err := c.args[i].Eval(e)
		if err != nil {
			return core.Token{}, nil
		}

		tokens = append(tokens, t)
	}

	return e.Callable(l, tokens...)
}

func (n *integer) Eval(_ core.Evaluator) (core.Token, error) {
	return n.t, nil
}


func (n *identifier) String() string { return n.t.Value }
func (n *integer) String() string    { return n.t.Value }
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
