package core

type Lexer interface {
	Count() int
	Pop() (Token, bool)
	Peek() (Token, bool)
	Tokenize(input string) error
}