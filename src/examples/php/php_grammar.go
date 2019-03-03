package php

import (
	"regexp"

	"../../parser/ast"
)

// PHP lexical tokens.
// These are all the valid tokens for the language.
const (
	GrammarTokenRangeLower ast.TokenIdentity = iota + 1

	TokenWhitespace
	TokenIdentifier

	TokenLiteralString
	TokenLiteralFloat
	TokenLiteralInteger
	TokenLiteralBooleanTrue
	TokenLiteralBooleanFalse
	TokenLiteralNull

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

// PHP lexical namespaces.
const (
	NamespaceRoot     ast.NamespaceIdentity = ast.NamespaceIdentityRoot
	NamespaceVariable ast.NamespaceIdentity = "variable"
	NamespaceString   ast.NamespaceIdentity = "string"
)

// PHP lexical rules.
const (
	GrammarRuleRangeLower ast.RuleIdentity = iota + 1

	RuleIdentifierBoolean
	RuleIdentifierNull
	RuleIdentifierFloat
	RuleIdentifierInteger
	RuleIdentifier

	RuleOperator

	RuleVariable
	RuleExpression
	RuleExpressionSide

	GrammarRuleRangeUpper
)

// Grammar represents the PHP grammar.
var Grammar = ast.Grammar{
	Namespaces: ast.NamespaceCollection{
		ast.Namespace{
			Identity: NamespaceRoot,
			Tokens: ast.TokenCollection{
				ast.TokenRegex{Token: ast.Token{Identity: TokenWhitespace}, Expression: regexp.MustCompile(`\s+`)},

				ast.TokenRegex{Token: ast.Token{Identity: TokenLiteralFloat}, Expression: regexp.MustCompile(`[0-9]+\.[0-9]+`)},
				ast.TokenRegex{Token: ast.Token{Identity: TokenLiteralInteger}, Expression: regexp.MustCompile(`[0-9]+`)},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenLiteralBooleanTrue}, Literal: `true`},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenLiteralBooleanFalse}, Literal: `false`},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenLiteralNull}, Literal: `null`},

				ast.TokenLiteral{Token: ast.Token{Identity: TokenDollar, TransitionTo: NamespaceVariable}, Literal: `$`},

				ast.TokenLiteral{Token: ast.Token{Identity: TokenSyntaxQuoteSingle}, Literal: `'`},
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
			Tokens: ast.TokenCollection{
				ast.TokenRegex{Token: ast.Token{Identity: TokenIdentifier, TransitionTo: ast.NamespaceIdentityShift}, Expression: regexp.MustCompile(`[a-zA-Z\_]{1}[a-zA-Z0-9\_]*`)},
			},
		},
		ast.Namespace{
			Identity: NamespaceString,
			Tokens: ast.TokenCollection{
				ast.TokenRegex{Token: ast.Token{Identity: TokenLiteralString}, Expression: regexp.MustCompile(`(\\\"|[^"])+`)},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenSyntaxQuoteDouble, TransitionTo: ast.NamespaceIdentityShift}, Literal: `"`},
			},
		},
	},
	Rules: ast.RuleSet{
		// Expression
		ast.RuleConcatenation{
			Rule: ast.Rule{Identity: RuleExpression},
			Rules: ast.RuleSet{
				ast.RuleReference{Target: RuleExpressionSide},                    // Left
				ast.RuleOptional{Target: ast.RuleToken{Target: TokenWhitespace}}, // Whitespace?
				ast.RuleReference{Target: RuleOperator},                          // Operator
				ast.RuleOptional{Target: ast.RuleToken{Target: TokenWhitespace}}, // Whitespace?
				ast.RuleChoice{ // Right
					Rules: ast.RuleSet{
						ast.RuleReference{Target: RuleExpression},
						ast.RuleReference{Target: RuleExpressionSide},
					},
				},
			},
		},
		// Expresion Side
		//  This represents all the tokens that can be on either side of an expression.
		ast.RuleChoice{
			Rule: ast.Rule{Identity: RuleExpressionSide, Ignore: true},
			Rules: ast.RuleSet{
				ast.RuleReference{Target: RuleVariable},
				ast.RuleReference{Target: RuleIdentifier},
			},
		},
		// Variable
		ast.RuleConcatenation{
			Rule: ast.Rule{Identity: RuleVariable},
			Rules: ast.RuleSet{
				ast.RuleToken{Target: TokenDollar},
				ast.RuleToken{Target: TokenIdentifier},
			},
		},
		// Identifer
		//  A top level rule that includes all types of identifers.
		ast.RuleChoice{
			Rule: ast.Rule{Identity: RuleIdentifier},
			Rules: ast.RuleSet{
				ast.RuleReference{Target: RuleIdentifierBoolean},
				ast.RuleReference{Target: RuleIdentifierNull},
				ast.RuleReference{Target: RuleIdentifierFloat},
				ast.RuleReference{Target: RuleIdentifierInteger},
			},
		},
		// Identifier Boolean
		ast.RuleChoice{
			Rule: ast.Rule{Identity: RuleIdentifierBoolean},
			Rules: ast.RuleSet{
				ast.RuleToken{Target: TokenLiteralBooleanFalse},
				ast.RuleToken{Target: TokenLiteralBooleanTrue},
			},
		},
		// Identifer Null
		ast.RuleChoice{
			Rule: ast.Rule{Identity: RuleIdentifierNull},
			Rules: ast.RuleSet{
				ast.RuleToken{Target: TokenLiteralNull},
			},
		},
		// Identifer Float
		//  A float is technically an integer with trailing parts, so its important
		//  the integer is placed below this rule in heirarchy.
		ast.RuleChoice{
			Rule: ast.Rule{Identity: RuleIdentifierFloat},
			Rules: ast.RuleSet{
				ast.RuleToken{Target: TokenLiteralFloat},
			},
		},
		// Identifer Integer
		ast.RuleChoice{
			Rule: ast.Rule{Identity: RuleIdentifierInteger},
			Rules: ast.RuleSet{
				ast.RuleToken{Target: TokenLiteralInteger},
			},
		},
		// Operators
		//  Some operators can be found as parts of identifers, e.g integers.
		//  For this reason operators are one of the lowest rules, and slowest to get too.
		ast.RuleChoice{
			Rule: ast.Rule{Identity: RuleOperator},
			Rules: ast.RuleSet{
				ast.RuleToken{Target: TokenAddition},
				ast.RuleToken{Target: TokenSubtraction},
				ast.RuleToken{Target: TokenDivision},
				ast.RuleToken{Target: TokenMultiplication},
			},
		},
	},
}
