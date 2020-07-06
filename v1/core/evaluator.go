package core

type Evaluator interface {
	Binary(op, l, r Token) (Token, error)
	Unary(op, t Token) (Token, error)
	Callable(l Token, tokens ...Token) (Token, error)
}