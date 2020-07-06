package calculator

import (
	"calculus/v1/core"
)

type Calculator struct {
	p core.Parser
	e core.Evaluator
}

func New(e core.Evaluator, p core.Parser) *Calculator {
	return &Calculator{p: p, e: e}
}

func (c *Calculator) Calculate(in string) (string, error) {
	node, err := c.p.Parse(in)
	if err != nil {
		return "", err
	}

	result, err := node.Eval(c.e)

	if err != nil {
		return "", err
	}

	return result.Value, nil
}
