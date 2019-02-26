package ast

import (
	"testing"
)

func TestRuleOptional_Match(test *testing.T) {
	grammar := Grammar{}
	sequence := LexemeSequence{
		Lexeme{Token: 1},
	}

	rule := RuleOptional{Target: RuleToken{Target: 1}}
	matched, _, node, err := rule.Match(grammar, sequence)

	if matched == false {
		test.Errorf(`Expecting rule to have matched, instead got: %#v`, err)
		return
	}

	cast, casted := node.(NodeLexeme)

	if casted == false {
		test.Errorf(`Expected to be parsed to ast.NodeLexeme, instead got: %#v`, node)
		return
	}

	if cast.GetTokenIdentity() != 1 {
		test.Errorf(`Expected to be the token 1, instead got: %d`, cast.GetTokenIdentity())
		return
	}
}

func TestRuleOptional_NoMatchReturnNodeNull(test *testing.T) {
	grammar := Grammar{}
	sequence := LexemeSequence{
		Lexeme{Token: 2},
	}

	rule := RuleOptional{Target: RuleToken{Target: 1}}
	matched, _, node, err := rule.Match(grammar, sequence)

	if matched == false {
		test.Errorf(`Expecting rule to have matched, instead got: %#v`, err)
		return
	}

	_, casted := node.(NodeNull)

	if casted == false {
		test.Errorf(`Expected to be parsed to ast.NodeNull, instead got: %#v`, node)
		return
	}
}
