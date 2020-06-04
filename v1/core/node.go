package core

type Node interface {
	Eval(Evaluator) (Token, error)
	String() string
}