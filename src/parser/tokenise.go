package parser

import (
	"errors"

	"./ast"
)

var TokeniseNoMatchError = errors.New("Tokeniser could not match a token in the token set")

func Tokenise(input string, tokenset ast.TokenSet) (LexemeSequence, int, error) {
	sequence := LexemeSequence{}
	length := len(input)
	index := 0

	for {
		if (index + 1) > length {
			break
		}

		fragment := input[index:]
		matched, lexeme := TokeniseFirstLexeme(fragment, tokenset)

		if matched == false {
			return sequence, index, TokeniseNoMatchError
		}

		offset := lexeme.Offset()
		index = index + (offset[0] + offset[1])

		sequence = append(sequence, lexeme)
	}

	return sequence, index, nil
}

func TokeniseFirstLexeme(input string, tokenset ast.TokenSet) (bool, Lexeme) {
	for _, token := range tokenset {
		matched, offset := token.Match(input)

		if matched == false {
			continue
		}

		if offset[0] != 0 {
			continue
		}

		value := input[0:offset[1]]
		lexeme := NewLexeme(token, offset, value)

		return true, lexeme
	}

	return false, nil
}
