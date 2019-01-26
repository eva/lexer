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
		test.Errorf(`Expected match, instead got error: %+v`, err)
		return
	}

	if node.CountNodeSequence() != 1 {
		test.Errorf(`Choice should have one child being the token rule, instead got: %+v`, node.GetNodeSequence())
		return
	}

	if node.CountLexemeSequence() != 0 {
		test.Errorf(`Choice should not have any direct lexemes matched, instead got: %+v`, node.GetLexemeSequence())
		return
	}

	if node.GetNodeSequence()[0].GetLexemeSequence()[0].GetTokenIdentity() != barTokenIdentity {
		test.Errorf(`Token did not match expected token`)
		return
	}

	if len(remaining) != 0 {
		test.Errorf(`Expected remaining sequence to be empty, instead got: %+v`, remaining)
		return
	}
}
