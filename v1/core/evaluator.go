package core

import (
	"github.com/pkg/errors"
)

type binaryHandler func(l, r token) (token, error)
type callableHandler func(toks ...token) (token, error)

type EvaluatorInitializer func(e *evaluator)

type evaluator struct {
	binaryHandlers map[kind]binaryHandler
	callableHandlers map[string]callableHandler
}

func newEvaluator() *evaluator {
	return &evaluator{
		binaryHandlers: make(map[kind]binaryHandler),
		callableHandlers: make(map[string]callableHandler),
	}
}

func (e *evaluator) binary(op, l, r token) (token, error) {
	var result token
	handler, ok := e.binaryHandlers[op.kind]
	if !ok {
		return result, errors.Errorf("handler not found for operator %s", op.String())
	}

	return handler(l, r)
}

func (e *evaluator) callable(l token, tokens ...token) (token, error) {
	var result token
	handler, ok := e.callableHandlers[l.value]
	if !ok {
		return result, errors.Errorf("handler not found")
	}

	return handler(tokens...)
}