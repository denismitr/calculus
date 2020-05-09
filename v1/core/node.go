package core

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

type node interface {
	eval() (token, error)
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
	op  kind
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

type call struct {
	fn   node
	args []node
}

func (n *identifier) eval() (token, error) {
	return n.t, nil
}

func (n *prefix) eval() (token, error) {
	return n.op, nil
}

func (n *postfix) eval() (token, error) {
	return token{}, nil
}

func (n *binary) eval() (token, error)  {
	l, lok := n.left.(*integer)
	r, rok := n.right.(*integer)
	if lok && rok {
		lval, _ := strconv.Atoi(l.t.value)
		rval, _ := strconv.Atoi(r.t.value)
		result := lval + rval
		return token{kind: INT, value: strconv.Itoa(result)}, nil
	}

	return token{}, errors.New("BOOOOO")
}

func (c *call) eval() (token, error)  {
	return c.args[0].eval() // fixme
}
func (n *integer) eval() (token, error) {
	return n.t, nil
}


func (n *identifier) String() string { return n.t.value }
func (n *integer) String() string    { return n.t.value }
func (p *prefix) String() string     { return fmt.Sprintf("(prefix %s %s)", p.op, p.arg) }
func (p *postfix) String() string { return fmt.Sprintf("(postfix %s %s)", p.op, p.arg) }
func (b *binary) String() string  { return fmt.Sprintf("(%s %s %s)", b.op, b.left, b.right) }

func (c *call) String() string {
	parts := make([]string, 0, len(c.args) + 1)
	parts = append(parts, c.fn.String())
	for _, arg := range c.args {
		parts = append(parts, arg.String())
	}
	return fmt.Sprintf("(call %s)", strings.Join(parts, " "))
}