package php

import (
	"regexp"

	"../../parser/ast"
)

const (
	GrammarTokenRangeLower ast.TokenIdentity = iota + 1

	TokenWhitespace

	TokenSyntaxQuoteSingle
	TokenSyntaxQuoteDouble
	TokenSyntaxParenthesisOpen
	TokenSyntaxParenthesisClose
	TokenSyntaxSquareBracketOpen
	TokenSyntaxSquareBracketClose

	TokenAddition
	TokenSubtraction
	TokenDivision
	TokenMultiplication

	GrammarTokenRangeUpper
)

var Grammar = ast.Grammar{
	Namespaces: ast.NamespaceSet{
		ast.Namespace{
			Identity: "root",
			Tokens: ast.TokenSet{
				ast.TokenRegex{Token: ast.Token{Identity: TokenWhitespace}, Expression: regexp.MustCompile(`\s+`)},

				ast.TokenLiteral{Token: ast.Token{Identity: TokenSyntaxQuoteSingle}, Literal: `'`},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenSyntaxQuoteDouble}, Literal: `"`},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenSyntaxParenthesisOpen}, Literal: `(`},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenSyntaxParenthesisClose}, Literal: `)`},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenSyntaxSquareBracketOpen}, Literal: `[`},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenSyntaxSquareBracketClose}, Literal: `]`},

				ast.TokenLiteral{Token: ast.Token{Identity: TokenAddition}, Literal: `+`},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenSubtraction}, Literal: `-`},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenDivision}, Literal: `/`},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenMultiplication}, Literal: `*`},
			},
		},
	},
}
