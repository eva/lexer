package ast

import (
	"testing"
)

func TestRuleToken_NoMatch_EmptySequence(test *testing.T) {
	rule := RuleToken{
		Rule:   Rule{Identity: 102030},
		Target: 1,
	}

	grammar := Grammar{}
	sequence := LexemeSequence{}

	matched, remaining, node, err := rule.Match(grammar, sequence)

	if matched == true {
		test.Error(`Expected matching against empty sequence to not match anything`)
	}

	if err == nil {
		test.Error(`Expected matching against empty sequence to return an error`)
	}

	// Remember that errors are always pointers
	casterr, instanceof := err.(*ErrRuleSequenceEmpty)

	if instanceof == false {
		test.Errorf(`Expected error returned to be ast.ErrRuleSequenceEmpty, instead got: %#v`, err)
	}

	if casterr.RuleIdentity != 102030 {
		test.Errorf(`Expected error to contain the offending rule identity, instead got: %#v`, casterr.RuleIdentity)
	}

	if node != nil {
		test.Error(`Expected matching against empty sequence to not return a node`)
	}

	if remaining.Count() != 0 {
		test.Error(`Expected matching against empty sequence to return an empty sequence`)
	}
}

func TestRuleToke_NoMatch(test *testing.T) {
	rule := RuleToken{
		Rule:   Rule{Identity: 102030},
		Target: 2,
	}

	grammar := Grammar{}
	sequence := LexemeSequence{
		Lexeme{Token: 1},
	}

	matched, remaining, node, err := rule.Match(grammar, sequence)

	if matched == true {
		test.Error(`Expected matching against invalid sequence to not match anything`)
	}

	if err == nil {
		test.Error(`Expected matching against invalid sequence to return an error`)
	}

	// Remember that errors are always pointers
	casterr, instanceof := err.(*ErrRuleTokenMatchFailure)

	if instanceof == false {
		test.Errorf(`Expected error returned to be ast.ErrRuleTokenMatchFailure, instead got: %#v`, err)
	}

	if casterr.RuleIdentity != 102030 {
		test.Errorf(`Expected error to contain the offending rule identity, instead got: %#v`, casterr.RuleIdentity)
	}

	if casterr.TargetTokenIdentity != 2 {
		test.Errorf(`Expected error to contain the target token identity, instead got: %#v`, casterr.TargetTokenIdentity)
	}

	if casterr.CurrentTokenIdentity != 1 {
		test.Errorf(`Expected error to contain the current token identity, instead got: %#v`, casterr.CurrentTokenIdentity)
	}

	if node != nil {
		test.Error(`Expected matching against invalid sequence to not return a node`)
	}

	if remaining.Count() != sequence.Count() {
		test.Error(`Expected matching against invalid sequence to return the given sequence`)
	}
}

func TestRuleToke_NoMatch_TokenNotFirst(test *testing.T) {
	rule := RuleToken{
		Rule:   Rule{Identity: 102030},
		Target: 2,
	}

	grammar := Grammar{}
	sequence := LexemeSequence{
		Lexeme{Token: 1},
		Lexeme{Token: 2},
	}

	matched, remaining, node, err := rule.Match(grammar, sequence)

	if matched == true {
		test.Error(`Expected matching against invalid sequence to not match anything`)
	}

	if err == nil {
		test.Error(`Expected matching against invalid sequence to return an error`)
	}

	// Remember that errors are always pointers
	casterr, instanceof := err.(*ErrRuleTokenMatchFailure)

	if instanceof == false {
		test.Errorf(`Expected error returned to be ast.ErrRuleTokenMatchFailure, instead got: %#v`, err)
	}

	if casterr.RuleIdentity != 102030 {
		test.Errorf(`Expected error to contain the offending rule identity, instead got: %#v`, casterr.RuleIdentity)
	}

	if casterr.TargetTokenIdentity != 2 {
		test.Errorf(`Expected error to contain the target token identity, instead got: %#v`, casterr.TargetTokenIdentity)
	}

	if casterr.CurrentTokenIdentity != 1 {
		test.Errorf(`Expected error to contain the current token identity, instead got: %#v`, casterr.CurrentTokenIdentity)
	}

	if node != nil {
		test.Error(`Expected matching against invalid sequence to not return a node`)
	}

	if remaining.Count() != sequence.Count() {
		test.Error(`Expected matching against invalid sequence to return the given sequence`)
	}
}

func TestRuleToken_Match_SingleTokenInSequence(test *testing.T) {
	rule := RuleToken{
		Rule:   Rule{Identity: 102030},
		Target: 1,
	}

	grammar := Grammar{}
	sequence := LexemeSequence{
		Lexeme{Token: 1},
	}

	matched, remaining, node, err := rule.Match(grammar, sequence)

	if matched == false {
		test.Error(`Expected to match against valid sequence`)
	}

	if err != nil {
		test.Error(`Expected to match against valid sequence without error`)
	}

	if node == nil {
		test.Error(`Expected to match against valid sequence with returned node`)
	}

	cast, casted := node.(NodeLexeme)

	if casted == false {
		test.Errorf(`Expected returned node to be instance of ast.NodeLexeme, instead got: %T`, node)
	}

	if cast.GetTokenIdentity() != 1 {
		test.Errorf(`Expected returned node to have token identity from sequence, instead got: %#v`, cast.GetTokenIdentity())
	}

	if remaining.IsEmpty() == false {
		test.Error(`Expected matching against sequence with one entry to have returned a empty remaining sequence`)
	}
}

func TestRuleToken_Match_ManyTokenInSequence(test *testing.T) {
	rule := RuleToken{
		Rule:   Rule{Identity: 102030},
		Target: 1,
	}

	grammar := Grammar{}
	sequence := LexemeSequence{
		Lexeme{Token: 1},
		Lexeme{Token: 2},
		Lexeme{Token: 3},
	}

	matched, remaining, node, err := rule.Match(grammar, sequence)

	if matched == false {
		test.Error(`Expected to match against valid sequence`)
	}

	if err != nil {
		test.Error(`Expected to match against valid sequence without error`)
	}

	if node == nil {
		test.Error(`Expected to match against valid sequence with returned node`)
	}

	cast, casted := node.(NodeLexeme)

	if casted == false {
		test.Errorf(`Expected returned node to be instance of ast.NodeLexeme, instead got: %T`, node)
	}

	if cast.GetTokenIdentity() != 1 {
		test.Errorf(`Expected returned node to have token identity from sequence, instead got: %#v`, cast.GetTokenIdentity())
	}

	if remaining.IsEmpty() == true {
		test.Error(`Expected matching against sequence to return a non empty remaining sequence`)
	}

	if remaining.Count() != 2 {
		test.Errorf(`Expected remaining sequence to be of count 2, instead got: %d`, remaining.Count())
	}
}
