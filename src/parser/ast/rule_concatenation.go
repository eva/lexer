package ast

type RuleConcatenation struct {
	Rule
	Rules RuleSet
}

func (r RuleConcatenation) Match(grammar GrammarKind, sequence LexemeSequence) (bool, LexemeSequence, NodeKind, error) {
	if r.Rules.IsEmpty() {
		err := NewErrRuleConcatenationEmptyRuleSet(r)
		return false, sequence, nil, err
	}

	var nodes NodeSequence

	remaining := sequence

	for _, rule := range r.Rules {
		matched, newremaining, child, err := rule.Match(grammar, remaining)
		remaining = newremaining

		if matched == false {
			return false, remaining, nil, err
		}

		if child.IsValid() == false {
			continue
		}

		nodes = nodes.Add(child)
	}

	node := NodeRule{
		Rule:  r.GetIdentity(),
		Nodes: nodes,
	}

	return true, remaining, node, nil
}
