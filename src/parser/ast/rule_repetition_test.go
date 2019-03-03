package ast

import "testing"

func TestRuleRepetition_MaximumCannotBeZero(test *testing.T) {
	grammar := Grammar{}
	sequence := LexemeSequence{}

	rule := RuleRepetition{
		Target: RuleToken{Target: 1},
	}

	matched, _, _, err := rule.Match(grammar, sequence)

	if matched == true {
		test.Error(`Expected maximum requirement when invalid to cause match failure`)
	}

	if err == nil {
		test.Error(`Expected an error to be returned as maximum is invalid (zero)`)
		return
	}

	// Remember that errors are always pointers
	_, instanceof := err.(*ErrRuleRepetitionMaximumZero)

	if instanceof == false {
		test.Errorf(`Expected error to be related to maximum being zero, instead got: %#v`, err)
		return
	}
}

func TestRuleRepetition_MatchSuccessBecauseMininumZero(test *testing.T) {
	grammar := Grammar{}
	sequence := LexemeSequence{}

	rule := RuleRepetition{
		Target:  RuleToken{Target: 1},
		Maximum: 1,
	}

	matched, _, _, err := rule.Match(grammar, sequence)

	if err != nil {
		test.Errorf(`Was not expecting error: %#v`, err)
		return
	}

	if matched == false {
		test.Error(`Expected minimum requirement to allow no matches to be a match`)
	}
}

func TestRuleRepetition_MatchFailureBecauseMininumExpected(test *testing.T) {
	grammar := Grammar{}
	sequence := LexemeSequence{}

	rule := RuleRepetition{
		Target:  RuleToken{Target: 1},
		Minimum: 2,
		Maximum: 1,
	}

	matched, _, _, err := rule.Match(grammar, sequence)

	if matched == true {
		test.Error(`Expected minimum requirement to cause match to fail`)
	}

	if err == nil {
		test.Error(`Expected specialised error when minimum was not reached`)
	}
}
