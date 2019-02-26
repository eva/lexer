package ast

import "testing"

func TestRuleChoice_EmptyRuleSet(test *testing.T) {
	grammar := Grammar{}
	sequence := LexemeSequence{}

	rule := RuleChoice{}
	matched, _, node, err := rule.Match(grammar, sequence)

	if matched == true {
		test.Error(`Expected empty ruleset to result in no match`)
	}

	if node != nil {
		test.Error(`Expected empty sequence and empty ruleset to not result in a returned node`)
	}

	if err == nil {
		test.Error(`Expected empty ruleset to result in an err`)
	}

	// Remember that errors are always pointers
	_, casted := err.(*ErrRuleChoiceEmptyRuleSet)

	if casted == false {
		test.Errorf(`Expected error to be an instance of ErrRuleChoiceEmptyRuleSet, instead got: %#v`, err)
	}
}

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

	matched, remaining, response, err := rule.Match(grammar, sequence)

	if response.GetNodeType() != NodeTypeRule {
		test.Errorf(`Expected returned node to have a type of ast.NodeTypeRule, instead got: %#v`, response.GetNodeType())
	}

	node, instanceof := response.(NodeRule)

	if instanceof == false {
		test.Errorf(`Expected node to be an instance of ast.NodeRule, instead got error: %#v`, node)
		return
	}

	if matched != true {
		test.Errorf(`Expected match, instead got error: %#v`, err)
		return
	}

	if node.IsEmpty() == false {
		test.Errorf(`Node should not be empty, instead got: %#v`, node.GetNodeSequence())
		return
	}

	childnode := node.GetNodeSequence()[0]

	if childnode.GetNodeType() != NodeTypeLexeme {
		test.Errorf(`Node first sequence node should have a type of ast.NodeTypeLexeme, instead got: %#v`, childnode.GetNodeType())
		return
	}

	child, childinstanceof := childnode.(NodeLexeme)

	if childinstanceof == false {
		test.Errorf(`Expected child node to be an instance of ast.NodeLexeme, instead got error: %#v`, node)
		return
	}

	if child.GetTokenIdentity() != barTokenIdentity {
		test.Errorf(`Token did not match expected token`)
		return
	}

	if len(remaining) != 0 {
		test.Errorf(`Expected remaining sequence to be empty, instead got: %+v`, remaining)
		return
	}
}
