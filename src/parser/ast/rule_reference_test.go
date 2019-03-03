package ast

import "testing"

func TestRuleReference_CannotFindReference(test *testing.T) {
	grammar := Grammar{}
	sequence := LexemeSequence{}

	rule := RuleReference{Target: 1}
	matched, _, _, err := rule.Match(grammar, sequence)

	if matched != false {
		test.Error(`Expected reference rule to fail match as referenced rule does not exist`)
	}

	if err == nil {
		test.Error(`Was expecting an error to be returned when reference is not found`)
	}

	// Remember that errors are always pointers
	cast, casted := err.(*ErrRuleReferenceNotFound)

	if casted == false {
		test.Errorf(`Expected the error to be of type ErrRuleReferenceNotFound, got %#v`, err)
		return
	}

	if cast.RuleIdentity != 1 {
		test.Errorf(`Expected error to contain target rule identity, got %d`, cast.RuleIdentity)
		return
	}
}

func TestRuleReference_FindReferenceRule(test *testing.T) {
	grammar := Grammar{
		Rules: RuleCollection{
			RuleChoice{
				Rule: Rule{Identity: 1},
				Rules: RuleCollection{
					RuleToken{Target: 10},
					RuleToken{Target: 11},
				},
			},
		},
	}

	sequence := LexemeSequence{
		Lexeme{Token: 10},
	}

	rule := RuleReference{Target: 1}
	matched, _, _, err := rule.Match(grammar, sequence)

	if matched == false {
		test.Error(`Expected reference rule to match`)
	}

	if err != nil {
		test.Errorf(`Was not expecting error: %#v`, err)
		return
	}
}

func TestRuleReference_FoundRuleReturnItsNode(test *testing.T) {
	grammar := Grammar{
		Rules: RuleCollection{
			RuleToken{Rule: Rule{Identity: 1}, Target: 10},
		},
	}

	sequence := LexemeSequence{
		Lexeme{Token: 10},
	}

	rule := RuleReference{Target: 1}
	matched, _, node, err := rule.Match(grammar, sequence)

	if matched == false {
		test.Errorf(`Expected reference rule to match, error is: %#v`, err)
		return
	}

	if node.GetNodeType() != NodeTypeLexeme {
		test.Errorf(`Expected the node to be of type NodeTypeLexeme as reference rule is a RuleToken, got %T`, node)
		return
	}

	_, casted := node.(NodeLexeme)

	if casted == false {
		test.Error(`Expected node to be an instance of NodeLexeme`)
		return
	}
}
