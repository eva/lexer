package ast

type RuleConcatenation struct {
	Rule
	Rules RuleSet
}

func (r RuleConcatenation) Match(grammar GrammarKind, sequence LexemeSequence) (bool, LexemeSequence, NodeKind, error) {
	var nodes NodeSequence

	remaining := sequence

	for _, rule := range r.Rules {
		matched, newremaining, child, _ := rule.Match(grammar, remaining)
		remaining = newremaining

		if matched == false {
			return r.Rule.Match(grammar, sequence)
		}

		nodes = append(nodes, child)
	}

	node := Node{
		Rule:  r,
		Nodes: nodes,
	}

	return true, remaining, node, nil
}
