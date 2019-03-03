package parser

import (
	"testing"

	"./ast"
)

func TestParseAnySequence_ParseReturnNode_WithRuleToken(test *testing.T) {
	grammar := ast.Grammar{
		Namespaces: ast.NamespaceCollection{
			ast.Namespace{
				Tokens: ast.TokenCollection{
					ast.TokenLiteral{Token: ast.Token{Identity: 1}},
					ast.TokenLiteral{Token: ast.Token{Identity: 2}},
				},
			},
		},
		Rules: ast.RuleCollection{
			ast.RuleToken{Target: 1},
		},
	}

	sequence := ast.LexemeSequence{
		ast.Lexeme{Token: 1},
	}

	node, err := ParseAnySequence(grammar, sequence)

	if err != nil {
		test.Errorf(`Expected to have no errors, instead got: %#v`, err)
		return
	}

	cast, casted := node.(ast.NodeLexeme)

	if casted == false {
		test.Errorf(`Expected to be parsed to ast.NodeLexeme, instead got: %#v`, node)
		return
	}

	if cast.GetTokenIdentity() != 1 {
		test.Errorf(`Expected to be the token 1, instead got: %d`, cast.GetTokenIdentity())
		return
	}
}

func TestParseAnySequence_ParseReturnNode_WithRuleConcatenation(test *testing.T) {
	grammar := ast.Grammar{
		Namespaces: ast.NamespaceCollection{
			ast.Namespace{
				Tokens: ast.TokenCollection{
					ast.TokenLiteral{Token: ast.Token{Identity: 1}},
					ast.TokenLiteral{Token: ast.Token{Identity: 2}},
				},
			},
		},
		Rules: ast.RuleCollection{
			ast.RuleConcatenation{
				Rule: ast.Rule{Identity: 10},
				Rules: ast.RuleCollection{
					ast.RuleToken{Target: 1},
					ast.RuleToken{Target: 2},
				},
			},
		},
	}

	sequence := ast.LexemeSequence{
		ast.Lexeme{Token: 1, Offset: ast.TokenOffset{0, 4}, Value: `true`},
		ast.Lexeme{Token: 2, Offset: ast.TokenOffset{4, 5}, Value: `false`},
	}

	if sequence.IsValid() == false {
		test.Error(`Expected test sequence to be valid`)
		return
	}

	node, err := ParseAnySequence(grammar, sequence)

	if err != nil {
		test.Errorf(`Expected to have no errors, instead got: %#v`, err)
		return
	}

	cast, casted := node.(ast.NodeRule)

	if casted == false {
		test.Errorf(`Expected to be parsed to ast.NodeRule, instead got: %#v`, node)
		return
	}

	if cast.GetRuleIdentity() != 10 {
		test.Errorf(`Expected to be the token 1, instead got: %d`, cast.GetRuleIdentity())
		return
	}

	if cast.Count() != 2 {
		test.Errorf(`Expected node to have 2 children, instead got: %d`, cast.Count())
		return
	}
}
