package core

import "github.com/pkg/errors"

var ErrNoMoreTokens = errors.New("no more tokens")
var ErrUnexpectedToken = errors.New("unexpected Token")

type Parser interface {
	Parse(input string) (Node, error)
}

