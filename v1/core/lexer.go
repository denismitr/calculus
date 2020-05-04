package core

import "github.com/pkg/errors"

func tokenize(input string) ([]token, error) {
	cursor := 0
	var tokens []token

	for cursor < len(input) {
		if input[cursor] == ' ' {
			cursor++
			continue
		}

		if input[cursor] == '(' {
			tokens = append(tokens, token{kind: LPAREN, value: "("})
			cursor++
			continue
		}

		if input[cursor] == ')' {
			tokens = append(tokens, token{kind: RPAREN, value: ")"})
			cursor++
			continue
		}

		if input[cursor] == '+' {
			tokens = append(tokens, token{kind: ADD, value: "+"})
			cursor++
			continue
		}

		if input[cursor] == '-' {
			tokens = append(tokens, token{kind: SUB, value: "-"})
			cursor++
			continue
		}

		if input[cursor] == '*' {
			tokens = append(tokens, token{kind: MUL, value: "*"})
			cursor++
			continue
		}

		if input[cursor] == '/' {
			tokens = append(tokens, token{kind: DIV, value: "/"})
			cursor++
			continue
		}

		if input[cursor] == '%' {
			tokens = append(tokens, token{kind: MOD, value: "%"})
			cursor++
			continue
		}

		if isNumeric(input[cursor:cursor+1]) {
			j := cursor + 1

			for j < len(input) && isNumeric(input[cursor:j+1]) {
				j++
			}

			tokens = append(tokens, token{kind: INT, value: input[cursor:j]})
			cursor = j
			continue
		}

		return nil, errors.Errorf("String invalid near character at position %d", cursor)
	}

	return tokens, nil
}
