package core

import (
	"fmt"
	"strings"
)

type node interface {
	expr()
	String() string
}

type named struct {
	value string
}

type prefix struct {
	op  kind
	arg node
}

type postfix struct {
	op  kind
	arg node
}

type binary struct {
	op    kind
	left  node
	right node
}

type numericLiteral struct {
	token token
}

type call struct {
	fn   node
	args []node
}

func (n *named) expr()   {}
func (p *prefix) expr()  {}
func (p *postfix) expr() {}
func (b *binary) expr()  {}
func (c *call) expr()    {}
func (n numericLiteral) expr()    {}


func (n *named) String() string   { return n.value }
func (n numericLiteral) String() string   { return n.token.value }
func (p *prefix) String() string  { return fmt.Sprintf("(prefix %s %s)", p.op, p.arg) }
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