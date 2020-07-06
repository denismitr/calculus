package evaluator

import (
	"calculus/v1/core"
	"github.com/pkg/errors"
)

type binaryHandler func(l, r core.Token) (core.Token, error)
type callableHandler func(toks ...core.Token) (core.Token, error)
type unaryHandler func(k core.Token) (core.Token, error)

type Evaluator struct {
	binaryHandlers map[core.Kind]binaryHandler
	unaryHandler map[core.Kind]unaryHandler
	callableHandlers map[string]callableHandler
}

func New(initializers ...Initializer) *Evaluator {
	e := &Evaluator{
		binaryHandlers: make(map[core.Kind]binaryHandler),
		callableHandlers: make(map[string]callableHandler),
		unaryHandler: make(map[core.Kind]unaryHandler),
	}

	for i := range initializers {
		initializers[i](e)
	}

	return e
}

func (e *Evaluator) Binary(op, l, r core.Token) (core.Token, error) {
	var result core.Token
	handler, ok := e.binaryHandlers[op.Kind]
	if !ok {
		return result, errors.Errorf("handler not found for operator %s", op.String())
	}

	return handler(l, r)
}

func (e *Evaluator) Unary(op, t core.Token) (core.Token, error) {
	var result core.Token

	handler, ok := e.unaryHandler[op.Kind]
	if !ok {
		return result, errors.Errorf("unary handler not found for operator %s", op.String())
	}

	return handler(t)
}

func (e *Evaluator) Callable(l core.Token, tokens ...core.Token) (core.Token, error) {
	var result core.Token
	handler, ok := e.callableHandlers[l.Value]
	if !ok {
		return result, errors.Errorf("handler not found")
	}

	return handler(tokens...)
}