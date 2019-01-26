package php

import (
	"testing"

	"../../parser"
	"../../parser/ast"
)

func TestTokeniseGrammar(test *testing.T) {
	dataset := []struct {
		input  string
		tokens []ast.TokenIdentity
	}{
		{`+`, []ast.TokenIdentity{TokenAddition}},
		{`+ +`, []ast.TokenIdentity{TokenAddition, TokenWhitespace, TokenAddition}},
		{`( + ]`, []ast.TokenIdentity{TokenSyntaxParenthesisOpen, TokenWhitespace, TokenAddition, TokenWhitespace, TokenSyntaxSquareBracketClose}},
		{`$foo`, []ast.TokenIdentity{TokenDollar, TokenIdentifier}},
	}

	for i, data := range dataset {
		i = i + 1

		sequence, index, err := parser.Tokenise(Grammar, data.input)

		if err != nil {
			test.Errorf(`[%d] Was not expecting to have an error: %v`, i, err)
			return
		}

		if index != len(data.input) {
			test.Errorf(`[%d] Was expecting to have consumed the entire input, index is %d`, i, index)
			return
		}

		if len(data.tokens) != len(sequence) {
			test.Errorf(`[%d] Expecting that length of sequence is %d but got %d`, i, len(data.tokens), len(sequence))
			return
		}

		for position, lexeme := range sequence {
			token := data.tokens[position]

			if lexeme.Token.GetIdentity() != token {
				test.Errorf(`[%d] @%d Token %v was expected, instead got %v`, i, position, token, lexeme.Token.GetIdentity())
				return
			}
		}
	}
}
