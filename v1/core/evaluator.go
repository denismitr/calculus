package core

import (
	"github.com/pkg/errors"
	"log"
)

type binaryHandler func(l, r token) (token, error)

type EvaluatorInitializer func(e *evaluator)

type evaluator struct {
	binaryHandlers map[kind]binaryHandler
}

func newEvaluator() *evaluator {
	return &evaluator{
		binaryHandlers: make(map[kind]binaryHandler),
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

func (e *evaluator) multiple(toks ...token) (token, error) {
	log.Fatalf("%#v", toks)
	var result token
	handler, ok := e.binaryHandlers[ADD] // fixme
	if !ok {
		return result, errors.Errorf("handler not found")
	}

	return handler(toks[0], toks[1]) // fixme
}