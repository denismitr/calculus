package core

type Library string

type Evaluator interface {
	Binary(op, l, r Token) (Token, error)
	Unary(op, t Token) (Token, error)
	Callable(l Token, tokens ...Token) (Token, error)
}

type EvaluatorInitializer func(library ...Library)