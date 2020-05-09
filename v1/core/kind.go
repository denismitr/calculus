package core

import "fmt"

type kind int

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
	LPAREN
	RPAREN
)

var kinds = [...]string{
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
	LPAREN:  "(",
	RPAREN:  ")",
}

func (k kind) String() string {
	if k >= 0 && k < kind(len(kinds)) {
		return kinds[k]
	}

	return fmt.Sprintf("token(%d)", int(k))
}
