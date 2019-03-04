package json

import (
	"regexp"

	"../../parser/ast"
)

// JSON lexical tokens.
// These are all the valid tokens for the language.
const (
	GrammarTokenRangeLower ast.TokenIdentity = iota + 1

	// 2
	TokenWhitespace

	// 3
	TokenLiteralString
	TokenLiteralFloat
	TokenLiteralInteger
	TokenLiteralBooleanTrue
	TokenLiteralBooleanFalse
	TokenLiteralNull

	// 9
	TokenSyntaxColon
	TokenSyntaxComma
	TokenSyntaxQuoteDouble

	// 12
	TokenSyntaxSquareBracketOpen
	TokenSyntaxSquareBracketClose
	TokenSyntaxCurlyBraceOpen
	TokenSyntaxCurlyBraceClose

	GrammarTokenRangeUpper
)

// PHP lexical namespaces.
const (
	NamespaceRoot   ast.NamespaceIdentity = ast.NamespaceIdentityRoot
	NamespaceString ast.NamespaceIdentity = "string"
)

// PHP lexical rules.
const (
	GrammarRuleRangeLower ast.RuleIdentity = iota + 1

	// 2
	RuleLiteral
	RuleLiteralNull
	RuleLiteralBoolean
	RuleLiteralFloat
	RuleLiteralInteger
	RuleLiteralString

	// 8
	RuleValue
	RuleArray
	RuleObject
	RulePair

	GrammarRuleRangeUpper
)

// OptionalWhitespace quick hand.
var OptionalWhitespace = ast.RuleOptional{Target: ast.RuleToken{Target: TokenWhitespace}}

// Grammar represents the PHP grammar.
var Grammar = ast.Grammar{
	IgnoreTokens: ast.TokenIdentityCollection{
		TokenWhitespace,
		TokenSyntaxQuoteDouble,
	},
	Namespaces: ast.NamespaceCollection{
		ast.Namespace{
			Identity: NamespaceRoot,
			Tokens: ast.TokenCollection{
				ast.TokenRegex{Token: ast.Token{Identity: TokenWhitespace, Name: `TokenWhitespace`}, Expression: regexp.MustCompile(`\s+`)},

				ast.TokenRegex{Token: ast.Token{Identity: TokenLiteralFloat, Name: `TokenLiteralFloat`}, Expression: regexp.MustCompile(`[0-9]+\.[0-9]+`)},
				ast.TokenRegex{Token: ast.Token{Identity: TokenLiteralInteger, Name: `TokenLiteralInteger`}, Expression: regexp.MustCompile(`[0-9]+`)},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenLiteralBooleanTrue, Name: `TokenLiteralBooleanTrue`}, Literal: `true`},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenLiteralBooleanFalse, Name: `TokenLiteralBooleanFalse`}, Literal: `false`},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenLiteralNull, Name: `TokenLiteralNull`}, Literal: `null`},

				ast.TokenLiteral{Token: ast.Token{Identity: TokenSyntaxColon, Name: `TokenSyntaxColon`}, Literal: `:`},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenSyntaxComma, Name: `TokenSyntaxComma`}, Literal: `,`},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenSyntaxQuoteDouble, Name: `TokenSyntaxQuoteDouble`, TransitionTo: NamespaceString}, Literal: `"`},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenSyntaxCurlyBraceOpen, Name: `TokenSyntaxCurlyBraceOpen`}, Literal: `{`},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenSyntaxCurlyBraceClose, Name: `TokenSyntaxCurlyBraceClose`}, Literal: `}`},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenSyntaxSquareBracketOpen, Name: `TokenSyntaxSquareBracketOpen`}, Literal: `[`},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenSyntaxSquareBracketClose, Name: `TokenSyntaxSquareBracketClose`}, Literal: `]`},
			},
		},
		ast.Namespace{
			Identity: NamespaceString,
			Tokens: ast.TokenCollection{
				ast.TokenRegex{Token: ast.Token{Identity: TokenLiteralString, Name: `TokenLiteralString`}, Expression: regexp.MustCompile(`(\\\"|[^"])+`)},
				ast.TokenLiteral{Token: ast.Token{Identity: TokenSyntaxQuoteDouble, Name: `TokenSyntaxQuoteDouble`, TransitionTo: ast.NamespaceIdentityShift}, Literal: `"`},
			},
		},
	},
	Rules: ast.RuleCollection{
		// Value
		ast.RuleChoice{
			Rule: ast.Rule{Identity: RuleValue},
			Rules: ast.RuleCollection{
				ast.RuleReference{Target: RuleLiteral},
				ast.RuleReference{Target: RuleObject},
				ast.RuleReference{Target: RuleArray},
			},
		},
		// Array
		ast.RuleConcatenation{
			Rule: ast.Rule{Identity: RuleArray},
			Rules: ast.RuleCollection{
				ast.RuleToken{Target: TokenSyntaxSquareBracketOpen}, // Square Open
				OptionalWhitespace, // Whitespace?
				ast.RuleOptional{
					Target: ast.RuleConcatenation{
						Rules: ast.RuleCollection{
							ast.RuleReference{Target: RuleValue},
							ast.RuleRepetition{
								Target: ast.RuleConcatenation{
									Rules: ast.RuleCollection{
										OptionalWhitespace, // Whitespace?
										ast.RuleToken{Target: TokenSyntaxComma},
										OptionalWhitespace, // Whitespace?
										ast.RuleReference{Target: RuleValue},
									},
								},
								Minimum: 0,
								Maximum: -1,
							},
						},
					},
				},
				OptionalWhitespace, // Whitespace?
				ast.RuleToken{Target: TokenSyntaxSquareBracketClose}, // Square Close
			},
		},
		// Object
		ast.RuleConcatenation{
			Rule: ast.Rule{Identity: RuleObject},
			Rules: ast.RuleCollection{
				ast.RuleToken{Target: TokenSyntaxCurlyBraceOpen}, // Curly Open
				OptionalWhitespace, // Whitespace?
				ast.RuleOptional{
					Target: ast.RuleConcatenation{
						Rules: ast.RuleCollection{
							ast.RuleReference{Target: RulePair},
							ast.RuleRepetition{
								Target: ast.RuleConcatenation{
									Rules: ast.RuleCollection{
										OptionalWhitespace, // Whitespace?
										ast.RuleToken{Target: TokenSyntaxComma},
										OptionalWhitespace, // Whitespace?
										ast.RuleReference{Target: RulePair},
									},
								},
								Minimum: 0,
								Maximum: -1,
							},
						},
					},
				},
				OptionalWhitespace, // Whitespace?
				ast.RuleToken{Target: TokenSyntaxCurlyBraceClose}, // Curly Close
			},
		},
		// Pair
		ast.RuleConcatenation{
			Rule: ast.Rule{Identity: RulePair},
			Rules: ast.RuleCollection{
				ast.RuleReference{Target: RuleLiteralString}, // Key
				OptionalWhitespace,                           // Whitespace?
				ast.RuleToken{Target: TokenSyntaxColon},      // Colon
				OptionalWhitespace,                           // Whitespace?
				ast.RuleReference{Target: RuleValue},         // Value
			},
		},
		// Literal
		//  A top level rule that includes all types of identifers.
		ast.RuleChoice{
			Rule: ast.Rule{Identity: RuleLiteral},
			Rules: ast.RuleCollection{
				ast.RuleReference{Target: RuleLiteralBoolean},
				ast.RuleReference{Target: RuleLiteralNull},
				ast.RuleReference{Target: RuleLiteralFloat},
				ast.RuleReference{Target: RuleLiteralInteger},
				ast.RuleReference{Target: RuleLiteralString},
			},
		},
		// Literal String
		ast.RuleConcatenation{
			Rule: ast.Rule{Identity: RuleLiteralString},
			Rules: ast.RuleCollection{
				ast.RuleToken{Target: TokenSyntaxQuoteDouble}, // Double Quote
				ast.RuleToken{Target: TokenLiteralString},     // Characters
				ast.RuleToken{Target: TokenSyntaxQuoteDouble}, // Double Quote
			},
		},
		// Literal Boolean
		ast.RuleChoice{
			Rule: ast.Rule{Identity: RuleLiteralBoolean},
			Rules: ast.RuleCollection{
				ast.RuleToken{Target: TokenLiteralBooleanFalse},
				ast.RuleToken{Target: TokenLiteralBooleanTrue},
			},
		},
		// Literal Null
		ast.RuleChoice{
			Rule: ast.Rule{Identity: RuleLiteralNull},
			Rules: ast.RuleCollection{
				ast.RuleToken{Target: TokenLiteralNull},
			},
		},
		// Literal Float
		//  A float is technically an integer with trailing parts, so its important
		//  the integer is placed below this rule in heirarchy.
		ast.RuleChoice{
			Rule: ast.Rule{Identity: RuleLiteralFloat},
			Rules: ast.RuleCollection{
				ast.RuleToken{Target: TokenLiteralFloat},
			},
		},
		// Literal Integer
		ast.RuleChoice{
			Rule: ast.Rule{Identity: RuleLiteralInteger},
			Rules: ast.RuleCollection{
				ast.RuleToken{Target: TokenLiteralInteger},
			},
		},
	},
}
