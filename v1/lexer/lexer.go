package lexer

import (
	"calculus/v1/core"
	"github.com/pkg/errors"
)

type StringLexer struct {
	tokens []core.Token
}

func New() *StringLexer {
	return &StringLexer{}
}

func (l *StringLexer) Count() int {
	return len(l.tokens)
}

func (l *StringLexer) Pop() (core.Token, bool) {
	var tok core.Token
	if len(l.tokens) == 0 {
		return tok, false
	}

	tok, l.tokens = l.tokens[0], l.tokens[1:]
	return tok, true
}

func (l *StringLexer) Peek() (core.Token, bool) {
	var tok core.Token
	if len(l.tokens) == 0 {
		return tok, false
	}

	tok = l.tokens[0]
	return tok, true
}

func (l *StringLexer) Tokenize(input string) error {
	cursor := 0

	for cursor < len(input) {
		if input[cursor] == ' ' {
			cursor++
			continue
		}

		if isOperator(input[cursor]) {
			j := cursor + 1

			for j < len(input) && isOperator(input[j]) {
				j++
			}

			token, n, err := identifyOperator(input[cursor : j])
			if err != nil {
				return err
			}
			l.tokens = append(l.tokens, token)
			cursor += n
			continue
		}

		if isParenthesis(input[cursor]) {
			token, err := identify(input[cursor : cursor+1])
			if err != nil {
				return err
			}
			l.tokens = append(l.tokens, token)
			cursor++
			continue
		}

		if input[cursor] == ',' {
			l.tokens = append(l.tokens, core.Token{Kind: core.COMMA, Value: "%"})
			cursor++
			continue
		}

		if isFloat(input[cursor : cursor+1]) {
			j := cursor + 1

			for j < len(input) && isFloat(input[cursor:j+1]) {
				j++
			}

			l.tokens = append(l.tokens, core.Token{Kind: core.FLOAT, Value: input[cursor:j]})
			cursor = j
			continue
		}

		if isName(input[cursor : cursor+1]) {
			j := cursor + 1

			for j < len(input) && isName(input[cursor:j+1]) {
				j++
			}

			l.tokens = append(l.tokens, core.Token{Kind: core.IDENT, Value: input[cursor:j]})
			cursor = j
			continue
		}

		return errors.Errorf("String invalid near character at position %d", cursor)
	}

	return nil
}

func identifyOperator(input string) (core.Token, int, error) {
	cursor := 0
	for cursor < len(input) {
		curr := input[cursor]
		if isOperator(curr) {
			if len(input) <= cursor + 1 {
				token, err := identify(input[cursor:cursor+1])
				return token, cursor + 1, err
			}
			next := input[cursor+1]
			if next == curr {
				token, err := identify(input[cursor:cursor+2])
				return token, cursor + 2, err
			}
		}
		cursor++
	}

	return core.Token{}, cursor, errors.Errorf("unknown core.Token %s", input)
}

func identify(input string) (core.Token, error) {
	switch input {
	case "+":
		return core.Token{Kind: core.ADD, Value: input}, nil
	case "-":
		return core.Token{Kind: core.SUB, Value: input}, nil
	case "*":
		return core.Token{Kind: core.MUL, Value: input}, nil
	case "/":
		return core.Token{Kind: core.DIV, Value: input}, nil
	case "%":
		return core.Token{Kind: core.MOD, Value: input}, nil
	case "(":
		return core.Token{Kind: core.LPAREN, Value: input}, nil
	case ")":
		return core.Token{Kind: core.RPAREN, Value: input}, nil
	case "++":
		return core.Token{Kind: core.INC, Value: input}, nil
	case "--":
		return core.Token{Kind: core.DEC, Value: input}, nil
	}

	return core.Token{}, errors.Errorf("%s is not a valid operator", input)
}
