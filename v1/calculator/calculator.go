package calculator

import (
	"calculus/v1/core"
	"calculus/v1/evaluator"
	"calculus/v1/lexer"
	"calculus/v1/parser"
)

type Calculator struct {
	p core.Parser
	e core.Evaluator
}

func New(libraries ...core.Library) *Calculator {
	parserInitializer := parser.InitializeDefaultGrammar()
	l := lexer.New()
	p := parser.New(l, parserInitializer)
	e := evaluator.New()

	libraryInitializer := evaluator.InitializeLibraries(e)
	libraryInitializer(libraries...)

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
