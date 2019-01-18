package parser

import (
	"errors"

	"./ast"
)

// ErrTokeniserCannotMatchToken is thrown by the Tokenise() function when the tokenisation
// cannot match any TokenKind within the active namespace.
var ErrTokeniserCannotMatchToken = errors.New("Tokeniser failed to match any tokens in the active namespace")

func Tokenise(grammar ast.GrammarKind, input string) (LexemeSequence, int, error) {
	sequence := LexemeSequence{}
	length := len(input)
	index := 0

	namespace, err := grammar.RootNamespace()

	if err != nil {
		return sequence, index, err
	}

	tokens := namespace.GetTokens()

	for {
		if (index + 1) > length {
			break
		}

		fragment := input[index:]
		matched, lexeme := TokeniseFirstLexeme(fragment, tokens)

		if matched == false {
			return sequence, index, ErrTokeniserCannotMatchToken
		}

		offset := lexeme.Offset
		index = index + (offset[0] + offset[1])

		sequence = append(sequence, lexeme)
	}

	return sequence, index, nil
}

func TokeniseFirstLexeme(input string, tokens ast.TokenSet) (bool, Lexeme) {
	for _, token := range tokens {
		matched, offset := token.Match(input)

		if matched == false {
			continue
		}

		if offset[0] != 0 {
			continue
		}

		value := input[0:offset[1]]
		lexeme := Lexeme{
			Token:  token,
			Offset: offset,
			Value:  value,
		}

		return true, lexeme
	}

	return false, Lexeme{}
}
