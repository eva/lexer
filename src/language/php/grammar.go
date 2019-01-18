package php

import (
	"../../parser/ast"
)

const (
	GrammarTokenRangeLower ast.TokenIdentity = iota + 1

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
				ast.TokenLiteral{Token: ast.Token{Identity: TokenAddition}, Literal: `+`},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenSubtraction}, Literal: `-`},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenDivision}, Literal: `/`},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenMultiplication}, Literal: `*`},
			},
		},
	},
}
