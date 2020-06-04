package core

import "fmt"

type Token struct {
	Kind  Kind
	Value string
}

func (t Token) String() string {
	return fmt.Sprintf("[%s]: %s", t.Kind.String(), t.Value)
}
