package ast

// RuleChoice represents a choice of many rules in order of definition.
// The node kind returned is dependant of the rule that matched first.
type RuleChoice struct {
	Rule
	Rules RuleSet
}

// Match will attempt a match on all rules in its ruleset and returns the first one that matches.
// The returned node kind is directly passed from the proxied rule, therefore this rule is not represented as a node.
func (rule RuleChoice) Match(grammar GrammarKind, sequence LexemeSequence) (bool, LexemeSequence, NodeKind, error) {
	if rule.Rules.IsEmpty() {
		err := NewErrRuleChoiceEmptyRuleSet(rule)
		return false, sequence, nil, err
	}

	for _, proxy := range rule.Rules {
		matched, remaining, child, err := proxy.Match(grammar, sequence)

		if err != nil {
			switch err.(type) {
			case *ErrRuleReferenceNotFound:
				return false, remaining, nil, err
			}
		}

		if matched == false {
			continue
		}

		node := NodeRule{
			Rule:  proxy.GetIdentity(),
			Nodes: NodeSequence{child},
		}

		return true, remaining, node, nil
	}

	err := NewErrRuleChoiceNoneMatched(rule)
	return false, sequence, nil, err
}
