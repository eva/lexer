package ast

import "testing"

func TestRuleConcatenation_MissingPart(test *testing.T) {
	var fooTokenIdentity TokenIdentity = 1
	var barTokenIdentity TokenIdentity = 2

	fooRule := RuleToken{Target: fooTokenIdentity}
	barRule := RuleToken{Target: barTokenIdentity}

	rule := RuleConcatenation{
		Rules: []RuleKind{fooRule, barRule},
	}

	grammar := Grammar{}
	sequence := LexemeSequence{
		Lexeme{Token: barTokenIdentity},
	}

	matched, _, _, _ := rule.Match(grammar, sequence)

	if matched != false {
		test.Error(`Expected match failure`)
		return
	}
}

func TestRuleConcatenation_AllPartsMatched(test *testing.T) {
	var fooTokenIdentity TokenIdentity = 1
	var barTokenIdentity TokenIdentity = 2

	fooRule := RuleToken{Target: fooTokenIdentity}
	barRule := RuleToken{Target: barTokenIdentity}

	rule := RuleConcatenation{
		Rules: []RuleKind{fooRule, barRule},
	}

	grammar := Grammar{}
	sequence := LexemeSequence{
		Lexeme{Token: fooTokenIdentity},
		Lexeme{Token: barTokenIdentity},
	}

	matched, _, _, err := rule.Match(grammar, sequence)

	if matched != true {
		test.Errorf(`Expected match, instead got: %v`, err)
		return
	}
}
