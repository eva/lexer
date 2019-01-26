package ast

import "testing"

func TestRuleChoice(test *testing.T) {
	var fooTokenIdentity TokenIdentity = 1
	var barTokenIdentity TokenIdentity = 2

	fooRule := RuleToken{Target: fooTokenIdentity}
	barRule := RuleToken{Target: barTokenIdentity}

	rule := RuleChoice{
		Rule:  Rule{Identity: 3000},
		Rules: []RuleKind{fooRule, barRule},
	}

	grammar := Grammar{}
	sequence := LexemeSequence{
		Lexeme{Token: barTokenIdentity},
	}

	matched, remaining, node, err := rule.Match(grammar, sequence)

	if matched != true {
		test.Errorf(`Expected match, instead got error: %v`, err)
		return
	}

	if node.CountNodes() != 1 {
		test.Errorf(`Choice should have one child being the token rule, instead got: %d`, node.GetNodes())
		return
	}

	if node.CountLexemes() != 0 {
		test.Errorf(`Choice should not have any direct lexemes matched, instead got: %v`, node.GetLexemes())
		return
	}

	if node.GetNodes()[0].GetLexemes()[0].GetTokenIdentity() != barTokenIdentity {
		test.Errorf(`Token did not match expected token`)
		return
	}

	if len(remaining) != 0 {
		test.Errorf(`Expected remaining sequence to be empty, instead got: %v`, remaining)
		return
	}
}
