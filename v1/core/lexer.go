package core

import "github.com/pkg/errors"

type lexer struct {
	tokens []token
}

func newLexer() *lexer {
	return &lexer{}
}

func (l *lexer) count() int {
	return len(l.tokens)
}

func (l *lexer) pop() (token, bool) {
	var tok token
	if len(l.tokens) == 0 {
		return tok, false
	}

	tok, l.tokens = l.tokens[0], l.tokens[1:]
	return tok, true
}

func (l *lexer) peek() (token, bool) {
	var tok token
	if len(l.tokens) == 0 {
		return tok, false
	}

	tok = l.tokens[0]
	return tok, true
}

func (l *lexer) tokenize(input string) error {
	cursor := 0

	for cursor < len(input) {
		if input[cursor] == ' ' {
			cursor++
			continue
		}

		if isOperator(input[cursor]) {
			token, n, err := identifyOperator(input[cursor : cursor+1])
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
			l.tokens = append(l.tokens, token{kind: COMMA, value: "%"})
			cursor++
			continue
		}

		if isNumeric(input[cursor : cursor+1]) {
			j := cursor + 1

			for j < len(input) && isNumeric(input[cursor:j+1]) {
				j++
			}

			l.tokens = append(l.tokens, token{kind: INT, value: input[cursor:j]})
			cursor = j
			continue
		}

		if isName(input[cursor : cursor+1]) {
			j := cursor + 1

			for j < len(input) && isName(input[cursor:j+1]) {
				j++
			}

			l.tokens = append(l.tokens, token{kind: IDENT, value: input[cursor:j]})
			cursor = j
			continue
		}

		return errors.Errorf("String invalid near character at position %d", cursor)
	}

	return nil
}

func identifyOperator(input string) (token, int, error) {
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

	return token{}, cursor, errors.Errorf("unknown token %s", input)
}

func identify(input string) (token, error) {
	switch input {
	case "+":
		return token{kind: ADD, value: input}, nil
	case "-":
		return token{kind: SUB, value: input}, nil
	case "*":
		return token{kind: MUL, value: input}, nil
	case "/":
		return token{kind: DIV, value: input}, nil
	case "%":
		return token{kind: MOD, value: input}, nil
	case "(":
		return token{kind: LPAREN, value: input}, nil
	case ")":
		return token{kind: RPAREN, value: input}, nil
	case "++":
		return token{kind: INC, value: input}, nil
	case "--":
		return token{kind: DEC, value: input}, nil
	}

	return token{}, errors.Errorf("%s is not a valid operator", input)
}
