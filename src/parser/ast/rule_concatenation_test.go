package ast

import "testing"

func TestRuleConcatenation_MissingPart(test *testing.T) {
	var fooTokenIdentity TokenIdentity = 1
	var barTokenIdentity TokenIdentity = 2

	barToken := TokenLiteral{Token: Token{Identity: barTokenIdentity}, Literal: `bar`}

	fooRule := RuleToken{Target: fooTokenIdentity}
	barRule := RuleToken{Target: barTokenIdentity}

	rule := RuleConcatenation{
		Rules: []RuleKind{fooRule, barRule},
	}

	grammar := Grammar{}
	sequence := LexemeSequence{
		Lexeme{Token: barToken},
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

	fooToken := TokenLiteral{Token: Token{Identity: fooTokenIdentity}, Literal: `foo`}
	barToken := TokenLiteral{Token: Token{Identity: barTokenIdentity}, Literal: `bar`}

	fooRule := RuleToken{Target: fooTokenIdentity}
	barRule := RuleToken{Target: barTokenIdentity}

	rule := RuleConcatenation{
		Rules: []RuleKind{fooRule, barRule},
	}

	grammar := Grammar{}
	sequence := LexemeSequence{
		Lexeme{Token: fooToken},
		Lexeme{Token: barToken},
	}

	matched, _, _, err := rule.Match(grammar, sequence)

	if matched != true {
		test.Errorf(`Expected match, instead got: %v`, err)
		return
	}
}
