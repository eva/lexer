package php

import (
	"regexp"

	"../../parser/ast"
)

const (
	GrammarTokenRangeLower ast.TokenIdentity = iota + 1

	TokenWhitespace
	TokenIdentifier

	TokenDollar

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

const (
	NamespaceRoot     ast.NamespaceIdentity = "root"
	NamespaceVariable ast.NamespaceIdentity = "variable"
)

const (
	GrammarRuleRangeLower = iota + 1

	RuleVariable

	GrammarRuleRangeUpper
)

// Grammar represents the PHP grammar.
var Grammar = ast.Grammar{
	Namespaces: ast.NamespaceSet{
		ast.Namespace{
			Identity: NamespaceRoot,
			Tokens: ast.TokenSet{
				ast.TokenRegex{Token: ast.Token{Identity: TokenWhitespace}, Expression: regexp.MustCompile(`\s+`)},

				ast.TokenLiteral{Token: ast.Token{Identity: TokenDollar, TransitionTo: NamespaceVariable}, Literal: `$`},

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
		ast.Namespace{
			Identity: NamespaceVariable,
			Tokens: ast.TokenSet{
				ast.TokenRegex{Token: ast.Token{Identity: TokenIdentifier, TransitionTo: ast.NamespaceIdentityShift}, Expression: regexp.MustCompile(`[a-zA-Z\_]{1}[a-zA-Z0-9\_]*`)},
			},
		},
	},
	Rules: ast.RuleSet{
		ast.RuleConcatenation{
			Rule: ast.Rule{Identity: RuleVariable},
			Rules: ast.RuleSet{
				ast.RuleToken{Token: TokenDollar},
				ast.RuleToken{Token: TokenIdentifier},
			},
		},
	},
}
