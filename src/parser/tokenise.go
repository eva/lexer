package parser

import (
	"errors"

	"./ast"
)

// ErrTokeniserCannotMatchToken is thrown by the Tokenise() function when the tokenisation
// cannot match any TokenKind within the active namespace.
var ErrTokeniserCannotMatchToken = errors.New("Tokeniser failed to match any tokens in the active namespace")

var ErrTokeniserCannotMoveNamespace = errors.New("Tokeniser failed to transition to another namespace")

var ErrTokeniserFinishedNotRoot = errors.New("Tokeniser finished tokenising but was not left in the root namespace.")

func Tokenise(grammar ast.GrammarKind, input string) (ast.LexemeSequence, int, error) {
	sequence := ast.LexemeSequence{}
	length := len(input)
	index := 0

	root, err := grammar.RootNamespace()

	if err != nil {
		return sequence, index, err
	}

	stack := NamespaceStack{}
	stack.Register(root)

	for {
		if (index + 1) > length {
			break
		}

		namespace := stack.Current()
		tokens := namespace.GetTokens()

		fragment := input[index:]
		matched, lexeme := TokeniseFirstLexeme(fragment, tokens)

		if matched == false {
			return sequence, index, ErrTokeniserCannotMatchToken
		}

		found, token := namespace.GetToken(lexeme.GetTokenIdentity())

		if found == false {
			panic(`Please make me error`)
		}

		transition, newnamespaceid := token.HasTransition()

		if transition == true {
			if newnamespaceid == ast.NamespaceIdentityShift {
				stack.Shift()
			} else {
				found, newnamespace := grammar.Namespace(newnamespaceid)

				if found == false {
					return sequence, index, ErrTokeniserCannotMoveNamespace
				}

				stack.Register(newnamespace)
			}
		}

		offset := lexeme.GetTokenOffset()
		index = index + (offset[0] + offset[1])

		sequence = append(sequence, lexeme)
	}

	if stack.Count() != 1 {
		return sequence, index, ErrTokeniserFinishedNotRoot
	}

	return sequence, index, nil
}

func TokeniseFirstLexeme(input string, tokens ast.TokenSet) (bool, ast.LexemeKind) {
	for _, token := range tokens {
		matched, offset := token.Match(input)

		if matched == false {
			continue
		}

		if offset[0] != 0 {
			continue
		}

		value := input[0:offset[1]]
		lexeme := ast.Lexeme{
			Token:  token.GetIdentity(),
			Offset: offset,
			Value:  value,
		}

		return true, lexeme
	}

	return false, nil
}
