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

		if input[cursor] == '(' {
			l.tokens = append(l.tokens, token{kind: LPAREN, value: "("})
			cursor++
			continue
		}

		if input[cursor] == ')' {
			l.tokens = append(l.tokens, token{kind: RPAREN, value: ")"})
			cursor++
			continue
		}

		if input[cursor] == '+' {
			l.tokens = append(l.tokens, token{kind: ADD, value: "+"})
			cursor++
			continue
		}

		if input[cursor] == '-' {
			l.tokens = append(l.tokens, token{kind: SUB, value: "-"})
			cursor++
			continue
		}

		if input[cursor] == '*' {
			l.tokens = append(l.tokens, token{kind: MUL, value: "*"})
			cursor++
			continue
		}

		if input[cursor] == '/' {
			l.tokens = append(l.tokens, token{kind: DIV, value: "/"})
			cursor++
			continue
		}

		if input[cursor] == '%' {
			l.tokens = append(l.tokens, token{kind: MOD, value: "%"})
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

		return errors.Errorf("String invalid near character at position %d", cursor)
	}

	return nil
}
