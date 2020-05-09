package core

import "fmt"

type token struct {
	kind kind
	value string
}

func (t token) String() string {
	return fmt.Sprintf("[%s]: %s", t.kind.String(), t.value)
}
