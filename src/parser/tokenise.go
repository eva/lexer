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

	found, root := grammar.FindRootNamespace()

	if found == false {
		return sequence, index, errors.New("Missing root namespace")
	}

	stack := NamespaceStack{}
	stack.Register(root)

	for {
		if (index + 1) > length {
			break
		}

		namespace := stack.Current()

		fragment := input[index:]
		matched, lexeme := TokeniseFirstLexeme(index, fragment, namespace)

		if matched == false {
			return sequence, index, ErrTokeniserCannotMatchToken
		}

		found, token := namespace.FindToken(lexeme.GetTokenIdentity())

		if found == false {
			panic(`Please make me error`)
		}

		transition, newnamespaceid := token.HasTransition()

		if transition == true {
			if newnamespaceid == ast.NamespaceIdentityShift {
				stack.Shift()
			} else {
				found, newnamespace := grammar.FindNamespace(newnamespaceid)

				if found == false {
					return sequence, index, ErrTokeniserCannotMoveNamespace
				}

				stack.Register(newnamespace)
			}
		}

		offset := lexeme.GetTokenOffset()
		index = offset[0] + offset[1]

		sequence = append(sequence, lexeme)
	}

	if stack.Count() != 1 {
		return sequence, index, ErrTokeniserFinishedNotRoot
	}

	return sequence, index, nil
}

func TokeniseFirstLexeme(index int, input string, namespace ast.NamespaceKind) (bool, ast.LexemeKind) {
	tokens := namespace.GetTokens()

	for _, token := range tokens {
		matched, offset := token.Match(input)

		if matched == false {
			continue
		}

		if offset[0] != 0 {
			continue
		}

		// The offset returned from the token will always be at position zero at this point.
		// We can now normalise this with the index provided.
		offset = ast.TokenOffset{index, offset[1]}

		value := input[0:offset[1]]
		lexeme := ast.Lexeme{
			Namespace: namespace.GetIdentity(),
			Token:     token.GetIdentity(),
			Offset:    offset,
			Value:     value,
		}

		return true, lexeme
	}

	return false, nil
}
