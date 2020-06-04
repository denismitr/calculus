package core

import "fmt"

type Kind int

const (
	ILLEGAL = iota
	IDENT
	INT
	FLOAT
	ADD
	SUB
	DIV
	MOD
	MUL
	EQL
	INC
	DEC
	LPAREN
	RPAREN

	COMMA
)

var Kinds = [...]string{
	ILLEGAL: "illegal",
	IDENT:   "IDENTIFIER",
	INT:     "integer",
	FLOAT:   "float",
	ADD:     "+",
	SUB:     "-",
	DIV:     "/",
	MUL:     "*",
	MOD:     "%",
	EQL:     "=",
	INC:     "++",
	DEC:     "--",
	LPAREN:  "(",
	RPAREN:  ")",
	COMMA: ".",
}

func (k Kind) String() string {
	if k >= 0 && k < Kind(len(Kinds)) {
		return Kinds[k]
	}

	return fmt.Sprintf("Token(%d)", int(k))
}
