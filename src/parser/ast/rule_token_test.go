package ast

import (
	"testing"
)

func TestRuleToken_Match_BasicFailure(test *testing.T) {
	var fooTokenIdentity TokenIdentity = 1
	var barTokenIdentity TokenIdentity = 2

	rule := RuleToken{
		Target: barTokenIdentity,
	}

	grammar := Grammar{}
	sequence := LexemeSequence{
		Lexeme{Token: fooTokenIdentity},
	}

	matched, remaining, node, err := rule.Match(grammar, sequence)

	if matched != false {
		test.Error(`Expected basic lexeme sequence to fail, "tokenfoo" provided instead of "tokenbar"`)
		return
	}

	if err, istype := err.(*ErrRuleTokenMatchFailure); istype == false {
		test.Errorf(`Unexpected error: %#v`, err)
		return
	}

	if node != nil {
		test.Errorf(`Expected failed match to return a nil node, got: %#v`, node)
		return
	}

	if len(remaining) != 1 {
		test.Errorf(`Expected remaining lexeme sequence to be the same as given, instead got: %#v`, remaining)
		return
	}
}

func TestRuleToken_Match_BasicSuccess(test *testing.T) {
	var fooTokenIdentity TokenIdentity = 1

	rule := RuleToken{
		Target: fooTokenIdentity,
	}

	grammar := Grammar{}
	sequence := LexemeSequence{
		Lexeme{Token: fooTokenIdentity},
	}

	matched, remaining, node, err := rule.Match(grammar, sequence)

	if matched != true || err != nil {
		test.Error(`Expected basic lexeme sequence to match`)
		return
	}

	if node.CountLexemeSequence() != 1 {
		test.Errorf(`Expected the node lexeme sequence to be 1, got %+v`, node.GetLexemeSequence())
		return
	}

	if node.GetRuleIdentity() != rule.GetIdentity() {
		test.Error(`Expected the matched rule to be itself`)
		return
	}

	if len(remaining) != 0 {
		test.Errorf(`Expected remaining sequence to be empty, instead got: %+v`, remaining)
		return
	}
}

func TestRuleToken_Match_BasicSuccessLexemeSequenceMutipleMatchSingle(test *testing.T) {
	var fooTokenIdentity TokenIdentity = 1
	var barTokenIdentity TokenIdentity = 2

	rule := RuleToken{
		Target: fooTokenIdentity,
	}

	grammar := Grammar{}
	sequence := LexemeSequence{
		Lexeme{Token: fooTokenIdentity},
		Lexeme{Token: barTokenIdentity},
	}

	matched, remaining, node, err := rule.Match(grammar, sequence)

	if matched != true || err != nil {
		test.Error(`Expected basic lexeme sequence to match`)
		return
	}

	if node.CountLexemeSequence() != 1 {
		test.Errorf(`Expected the node lexeme sequence to be 1, got %+v`, node.CountLexemeSequence())
		return
	}

	if node.GetRuleIdentity() != rule.GetIdentity() {
		test.Error(`Expected the matched rule to be itself`)
		return
	}

	if node.GetLexemeSequence()[0].GetTokenIdentity() != fooTokenIdentity {
		test.Error(`Expected the rule to match the first token`)
		return
	}

	if len(remaining) != 1 {
		test.Errorf(`Expected remaining sequence to be the last remaining, instead got: %v`, remaining)
		return
	}
}
